package main

import (
	"image"
	"log"

	"github.com/disintegration/imaging"
)

func main() {
	cropImg(49, 583, 396, 672, "express.png", "1") //express
	cropImg(49, 560, 558, 658, "regular.png", "2") //regular

}

func cropImg(x0, y0, x1, y1 int, imgname string, number string) {

	img, err := imaging.Open("img" + number + ".png")
	if err != nil {
		log.Printf("Error on image cropping: %v", err)
		return
	}
	cropped := imaging.Crop(img, image.Rect(x0, y0, x1, y1))
	err = imaging.Save(cropped, imgname)
}
