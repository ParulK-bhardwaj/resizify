package main

import (
	"flag"
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

// ResizeParams holds parameters for resizing images
type ResizeParams struct {
	width  uint
	height uint
}

// resizeImage resizes an image to the specified width and height
// https://spec.oneapi.io/oneipl/0.5/transform/resize_lanczos.html
// This function changes an image size using interpolation with the Lanczos filter.
func resizeImage(img image.Image, params ResizeParams) image.Image {
	return resize.Resize(params.width, params.height, img, resize.Lanczos3)
}

// processImage opens, resizes, and saves the image
func processImage(filePath string, params ResizeParams) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		return err
	}

	resizedImg := resizeImage(img, params)

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
		return fmt.Errorf("unsupported image format")
	}
	return nil
}

func main() {
	path := flag.String("path", "", "Path to the images directory")
	width := flag.String("width", "800", "Width to resize images to")
	height := flag.String("height", "600", "Height to resize images to")

	flag.Parse()

	if *path == "" {
		fmt.Println("Usage: go run main.go -path [path_to_images] -width [width] -height [height]")
		return
	}

	// string to uint: base 10, 64-bit: most commanly used
	widthUint, err := strconv.ParseUint(*width, 10, 64)
	if err != nil {
		fmt.Println("Invalid width value")
		return
	}

	heightUint, err := strconv.ParseUint(*height, 10, 64)
	if err != nil {
		fmt.Println("Invalid height value")
		return
	}

	params := ResizeParams{width: uint(widthUint), height: uint(heightUint)}

	err = filepath.Walk(*path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			err = processImage(path, params)
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
