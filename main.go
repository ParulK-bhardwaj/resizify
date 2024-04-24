package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strconv"

	"github.com/nfnt/resize"
)

// resizeImage resizes an image to the specified width and height
// https://spec.oneapi.io/oneipl/0.5/transform/resize_lanczos.html
// This function changes an image size using interpolation with the Lanczos filter.
func resizeImage(img image.Image, width, height uint) image.Image {
	return resize.Resize(width, height, img, resize.Lanczos3)
}

// processImage opens, resizes, and saves the image
func processImage(filePath string, width, height uint) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		return err
	}

	resizedImg := resizeImage(img, width, height)

	outputFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	switch format {
	case "jpeg":
		jpeg.Encode(outputFile, resizedImg, nil)
	case "png":
		png.Encode(outputFile, resizedImg)
	case "gif":
		gif.Encode(outputFile, resizedImg, nil)
	default:
		return fmt.Errorf("unsupported Image format")
	}

	return nil
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Please use: go run main.go [path_to_images] [width] [height]")
		return
	}

	path := os.Args[1]
	width := uint(atoui(os.Args[2]))
	height := uint(atoui(os.Args[3]))

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			err = processImage(path, width, height)
			if err != nil {
				fmt.Printf("Failed to process image %s: %v\n", path, err)
			} else {
				fmt.Printf("Processed image %s successfully\n", path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error processing images: %v\n", err)
	}
}

// atoui converts a string to uint
func atoui(str string) uint {
	val, _ := strconv.ParseUint(str, 10, 64)
	return uint(val)
}
