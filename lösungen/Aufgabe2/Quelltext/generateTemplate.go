package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
)

func GenerateTemplate(path string, generateSpeed func(int, int) (Vector, Vector), mult float64) {
	imageReader, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	image, err := png.Decode(imageReader)
	if err != nil {
		log.Panic(err)
	}
	imageConfigReader, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	imageConfig, err := png.DecodeConfig(imageConfigReader)
	if err != nil {
		log.Panic(err)
	}
	var templateData = fmt.Sprintf("%v;%v", int(float64(imageConfig.Width)*mult), int(float64(imageConfig.Height)*mult))
	for x := 0; x < imageConfig.Width; x++ {
		for y := 0; y < imageConfig.Height; y++ {
			r, g, b, _ := image.At(x, y).RGBA()
			r, g, b = r/255, g/255, b/255
			if int((r+g+b)/3) == 0 {
				continue
			}
			fmt.Printf("%v, %v, %v\n", r, g, b)
			plusSpeed, minusSpeed := generateSpeed(x, y)
			templateData += fmt.Sprintf("\n%v;%v;%v;%v;%v;%v;%v", int(float64(x)*mult), int(float64(y)*mult), plusSpeed.x, plusSpeed.y, minusSpeed.x, minusSpeed.y, (r+g+b)/3)
		}
	}
	file, err := os.Create("template.txt")
	if err != nil {
		log.Panic(err)
	}
	_, err = file.Write([]byte(templateData))
	if err != nil {
		log.Panic(err)
	}
}
