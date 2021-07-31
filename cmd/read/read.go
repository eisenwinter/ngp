package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"image"
	_ "image/png"
	"io/ioutil"
	"os"

	"github.com/eisenwinter/ngp/pkg/greenpass"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func main() {
	var fileFlag = flag.String("file", "", "specify a text file to be read from")
	flag.Parse()
	if fileFlag != nil && *fileFlag != "" {
		file, err := loadFile(*fileFlag)
		if err != nil {
			panic(err)
		}
		decode(file)
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		decode(text)
	}
}

func loadFile(path string) (string, error) {
	fmt.Printf("Loading: %v\n", path)
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

func decode(input string) {
	gp := greenpass.New()
	result, err := gp.Decode([]byte(input))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", result)
}
