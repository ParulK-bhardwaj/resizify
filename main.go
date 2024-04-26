package main

import (
	"encoding/json"
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
	"github.com/ttacon/chalk"
)

// ResizeParams holds parameters for resizing images
type ResizeParams struct {
	width  uint
	height uint
}

// ImageProcessingResult holds filepath and error
// omitempty: is used here to omit empty errors
type ImageProcessingResult struct {
	FilePath string `json:"file_path"`
	Error    string `json:"error,omitempty"`
}

// resizeImage resizes an image to the specified width and height
// https://spec.oneapi.io/oneipl/0.5/transform/resize_lanczos.html
// This function changes an image size using interpolation with the Lanczos filter.
func resizeImage(img image.Image, params ResizeParams) image.Image {
	return resize.Resize(params.width, params.height, img, resize.Lanczos3)
}

// processImage opens, resizes, and saves the image
func processImage(filePath string, params ResizeParams) (ImageProcessingResult, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return ImageProcessingResult{FilePath: filePath, Error: err.Error()}, err
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		return ImageProcessingResult{FilePath: filePath, Error: err.Error()}, err
	}

	resizedImg := resizeImage(img, params)

	outputFile, err := os.Create(filePath)
	if err != nil {
		return ImageProcessingResult{FilePath: filePath, Error: err.Error()}, err

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
		errorMsg := fmt.Errorf("unsupported image format")
		return ImageProcessingResult{FilePath: filePath, Error: errorMsg.Error()}, errorMsg

	}
	return ImageProcessingResult{FilePath: filePath}, nil
}

func main() {
	path := flag.String("path", "", "Path to the images directory")
	// default 800
	width := flag.String("width", "800", "Width to resize images to")
	// default 600
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

	var successes, failures []ImageProcessingResult
	err = filepath.Walk(*path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			failures = append(failures, ImageProcessingResult{FilePath: path, Error: err.Error()})
			return nil
		}
		if !info.IsDir() {
			result, err := processImage(path, params)
			if err != nil {
				failures = append(failures, result)
				fmt.Printf("%sFailed to process image %s%s: %s%v%s\n", chalk.Red, chalk.Reset, path, chalk.Yellow, err, chalk.Reset)
			} else {
				successes = append(successes, result)
				fmt.Printf("%sProcessed image successfully %s%v%s\n", chalk.Green, chalk.Reset, path, chalk.Reset)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error processing images: %v\n", err)
	}

	// successes to successes.json JSON file
	// 0644: file permission used when creating file
	successFile, _ := json.MarshalIndent(successes, "", "  ")
	_ = os.WriteFile("successes.json", successFile, 0644)

	// failures to a failures.json JSON file
	failureFile, _ := json.MarshalIndent(failures, "", "  ")
	_ = os.WriteFile("failures.json", failureFile, 0644)

	fmt.Println(chalk.Blue, "Image processing complete. Results saved to successes.json and failures.json.")
}
