package main

import (
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"testing"
)

func TestResizeImage(t *testing.T) {
	tests := []struct {
		name         string
		img          image.Image
		params       ResizeParams
		expectedSize func(width, height int) bool
	}{
		{
			name:         "Resize PNG",
			img:          image.NewRGBA(image.Rect(0, 0, 100, 100)),
			params:       ResizeParams{width: 50, height: 50},
			expectedSize: func(width, height int) bool { return width == 50 && height == 50 },
		},
		{
			name:         "Resize JPEG",
			img:          image.NewRGBA(image.Rect(0, 0, 1000, 1000)),
			params:       ResizeParams{width: 500, height: 800},
			expectedSize: func(width, height int) bool { return width == 500 && height == 800 },
		},
		{
			name:         "Should not Resize file to zero dimensions and keep it as is",
			img:          image.NewRGBA(image.Rect(0, 0, 1000, 1000)),
			params:       ResizeParams{width: 0, height: 0},
			expectedSize: func(width, height int) bool { return width == 1000 && height == 1000 },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resizedImg := resizeImage(test.img, test.params)
			bounds := resizedImg.Bounds()
			if !test.expectedSize(bounds.Dx(), bounds.Dy()) {
				t.Errorf("Failed %s: expected size %dx%d, got %dx%d", test.name, test.params.width, test.params.height, bounds.Dx(), bounds.Dy())
			}
		})
	}
}

func TestProcessImagePredefined(t *testing.T) {
	tests := []struct {
		name          string
		filePath      string
		params        ResizeParams
		expectedError bool
	}{
		{
			name:          "Process JPEG image",
			filePath:      filepath.Join("test_images", "valid.jpeg"),
			params:        ResizeParams{width: 100, height: 100},
			expectedError: false,
		},
		{
			name:          "Process PNG image",
			filePath:      filepath.Join("test_images", "valid.png"),
			params:        ResizeParams{width: 100, height: 100},
			expectedError: false,
		},
		{
			name:          "Process unsupported image",
			filePath:      filepath.Join("test_images", "invalid.jpg"),
			params:        ResizeParams{width: 100, height: 100},
			expectedError: true,
		},
		{
			name:          "File does not exist",
			filePath:      filepath.Join("test_images", ""),
			params:        ResizeParams{width: 100, height: 100},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := processImage(test.filePath, test.params, "processed_images_tests", 700)
			if (err != nil) != test.expectedError {
				t.Errorf("Test %s failed, expected error: %v, got: %v", test.name, test.expectedError, err)
			}
		})
	}
	defer os.RemoveAll("processed_images_tests") // clean up
}

func BenchmarkResizeImage(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
	params := ResizeParams{width: 100, height: 100}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resizeImage(img, params)
	}
}

func BenchmarkProcessImage(b *testing.B) {
	tempFile, err := os.CreateTemp("", "test.*.jpg")
	if err != nil {
		b.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // clean up

	img := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
	jpeg.Encode(tempFile, img, nil) //save to the temp file
	tempFile.Close()

	params := ResizeParams{width: 100, height: 100}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := processImage(tempFile.Name(), params, "processed_images_tests", 700)
		if err != nil {
			b.Fatalf("Failed to process image: %v", err)
		}
	}
	defer os.RemoveAll("processed_images_tests") // clean up
}
