package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"image"
	_ "image/png"
	"io/ioutil"
	"os"

	"github.com/eisenwinter/ngp/pkg/euspec"
	"github.com/eisenwinter/ngp/pkg/greenpass"
	"github.com/eisenwinter/ngp/pkg/template"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func main() {
	var fileFlag = flag.String("f", "", "specify a text file to be read from")
	var templateFlag = flag.String("t", "", "specify a template file to be read from")
	var outFlag = flag.String("o", "out.html", "name of output file")
	flag.Parse()
	if fileFlag != nil && *fileFlag != "" && templateFlag != nil && *templateFlag != "" {
		file, err := loadFile(*fileFlag)
		if err != nil {
			panic(err)
		}
		res, err := generate(file, fileAsBase64(*fileFlag), *templateFlag)
		if err != nil {
			panic(err)
		}
		f, err := os.Create(*outFlag)

		if err != nil {
			panic(err)
		}

		defer f.Close()

		_, err = f.WriteString(res)
		if err != nil {
			panic(err)
		}

	} else {
		fmt.Println("Please specify -f and -t flags.")
	}
}

func fileAsBase64(path string) string {
	imgdata, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(imgdata)
}

func loadFile(path string) (string, error) {
	imgdata, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	img, _, err := image.Decode(bytes.NewReader(imgdata))
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

func generate(input string, qr string, templatePath string) (string, error) {
	gp := greenpass.New()
	result, err := gp.Decode([]byte(input))
	if err != nil {
		panic(err)
	}
	view := euspec.ToRenderView(*result, "data:image/png;base64,"+qr, "de")
	tr := template.New()
	return tr.Render(templatePath, view)
}
