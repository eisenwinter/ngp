package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	"github.com/eisenwinter/ngp/pkg/euspec"
	"github.com/eisenwinter/ngp/pkg/greenpass"
	tp "github.com/eisenwinter/ngp/pkg/template"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

var mainTemplate *template.Template
var errorTemplate *template.Template

type errorOutput struct {
	Error string
}

func readQr(data []byte) (string, error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)
	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		return "", err
	}
	return result.String(), nil
}

func generateNgp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "Sorry, POST method is supported.")
	}
	r.ParseMultipartForm(1 << 20)
	file, handler, err := r.FormFile("qr")
	if err != nil {
		errorTemplate.Execute(w, &errorOutput{Error: "Keine oder ungültige Datei"})
		return
	}
	defer file.Close()
	ct := handler.Header.Get("Content-Type")
	if ct != "image/png" {
		errorTemplate.Execute(w, &errorOutput{Error: "QR Code ist keine .png Datei"})
		return
	}
	template := r.FormValue("format")
	templateToLoad := "templates/default.html"
	switch template {
	case "portrait":
		templateToLoad = "templates/portrait.html"
		break
	default:
		templateToLoad = "templates/default.html"
		break
	}
	lang := r.FormValue("lang")
	if lang != "en" {
		lang = "de"
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		errorTemplate.Execute(w, &errorOutput{Error: "QR Code konnte nicht gelesen werden"})
		return
	}
	qr, err := readQr(fileBytes)
	if err != nil {
		errorTemplate.Execute(w, &errorOutput{Error: "Inhalt von QR Code konnte nicht gelesen werden"})
		return
	}
	gp := greenpass.New()
	result, err := gp.Decode([]byte(qr))
	if err != nil {
		errorTemplate.Execute(w, &errorOutput{Error: "Inhalt von QR Code konnte nicht verarbeitet werden"})
		return
	}
	view := euspec.ToRenderView(*result, "data:image/png;base64,"+base64.StdEncoding.EncodeToString(fileBytes), lang)
	tr := tp.New()
	rendered, err := tr.Render(templateToLoad, view)
	if err != nil {
		errorTemplate.Execute(w, &errorOutput{Error: "ngp konnte nicht generiert werden"})
		return
	}
	fmt.Fprintf(w, rendered)
}

func showPage(w http.ResponseWriter, r *http.Request) {
	mainTemplate.Execute(w, nil)
}

func routes() {
	http.HandleFunc("/gen", generateNgp)
	http.HandleFunc("/", showPage)

}

func getPortVariable() string {
	if value, ok := os.LookupEnv("PORT"); ok {
		return value
	}
	return "8080"
}

func main() {
	mainTemplate = template.Must(template.ParseFiles("./views/index.html"))
	errorTemplate = template.Must(template.ParseFiles("./views/error.html"))
	routes()
	http.ListenAndServe(fmt.Sprintf(":%s", getPortVariable()), nil)
}
