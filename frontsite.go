package main

import (
	"fmt"
	"html/template"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"strconv"

	"github.com/nfnt/resize"
)

// ResizeParams holds parameters for resizing images
type ResizeParams struct {
	Width  uint
	Height uint
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *ResizeParams) {
	fullPath := "/Users/parulbhardwaj/Desktop/ACS-4210-strongly-typed-language-golang/resizify/" + tmpl + ".html"

	fmt.Println("Attempting to load template from:", fullPath)
	t, err := template.ParseFiles(fullPath)
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}

// Upload and resize image
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderTemplate(w, "upload", nil)
		return
	}

	// Handle POST
	// 10 MB max file size
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("uploadFile")
	if err != nil {
		http.Error(w, "Could not read file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	width, err := strconv.Atoi(r.FormValue("width"))
	if err != nil {
		http.Error(w, "Invalid width: "+err.Error(), http.StatusBadRequest)
		return
	}

	height, err := strconv.Atoi(r.FormValue("height"))
	if err != nil {
		http.Error(w, "Invalid height: "+err.Error(), http.StatusBadRequest)
		return
	}

	img, format, err := image.Decode(file)
	if err != nil {
		http.Error(w, "Unsupported image format or decode error: "+err.Error(), http.StatusBadRequest)
		return
	}

	resizedImg := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	w.Header().Set("Content-Type", "image/"+format)
	switch format {
	case "jpeg":
		jpeg.Encode(w, resizedImg, nil)
	case "png":
		png.Encode(w, resizedImg)
	case "gif":
		gif.Encode(w, resizedImg, nil)
	default:
		http.Error(w, "Unsupported image format", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", uploadHandler)
	fmt.Println("Server started on :3000")
	http.ListenAndServe(":3000", nil)
}
