package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args
	img_path := args[1]
	percentage, err := strconv.Atoi(args[2])
	check(err)

	if percentage < 0 {
		fmt.Printf("Invalid percentage value:%d.\n", percentage)
		return
	}

	file, err := os.Open(img_path)

	check(err)
	defer file.Close()

	img, err := jpeg.Decode(file)
	check(err)

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	fmt.Printf("Original width: %d.\n", width)
	fmt.Printf("Original height: %d.\n", height)

	new_width := width * percentage / 100
	new_height := height * percentage / 100

	fmt.Printf("New width: %d.\n", new_width)
	fmt.Printf("New height: %d.\n", new_height)

	upLeft := image.Point{0, 0}
	lowRight := image.Point{new_width, new_height}

	new_img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for w := 0; w < new_width; w++ {
		for h := 0; h < new_height; h++ {
			new_img.Set(w, h, img.At(w*100/percentage, h*100/percentage))
		}
	}

	new_file, err := os.Create("resized-image.jpg")
	check(err)

	jpeg.Encode(new_file, new_img, nil)
	check(err)

	file.Close()
	new_file.Close()
}
