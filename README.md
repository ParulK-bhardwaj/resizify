[![Go Report Card](https://goreportcard.com/badge/github.com/ParulK-bhardwaj/resizify)](https://goreportcard.com/report/github.com/ParulK-bhardwaj/resizify)

# Resizify - Bulk Image Resizer Utility

## Description
Resizify is a command-line utility designed to simplify the process of resizing images in bulk. It takes a directory containing images as input, resizes each image to the specified dimensions, and saves the resized images to a new directory.

## Features
- Bulk Image Resizing: Resizify allows you to resize multiple images at once, saving you time and effort.
- Customizable Dimensions: You can specify the width and height to which you want to resize the images, providing flexibility for various use cases.
- Supports Multiple Formats: Resizify supports popular image formats including JPEG, PNG, and GIF.
- Error Handling: This provides feedback on successful and failed image processing, making it easy to identify any issues during resizing.
- Preventing Data Loss: Saving the resized images in a new directory
- Customizable Compression: Compress the size of the image.
- Drag and Drop - Front End: (Future Feature) Users will be able to drag and drop images into a window for resizing

## Installation
To install Resizify, you need to have Go installed on your system. To install go follow the instructions from [Go's official site](https://go.dev/doc/install).

### Follow these steps to install and run the program:

1. Clone the repository:

```bash
git clone https://github.com/ParulK-bhardwaj/resizify.git 
```

2. Navigate to the project directory:
```bash
cd <path_to_project_directory>
```

3. Build the project
```bash
go build -o resizify
```

## Usage

```go
go run main.go -path [path_to_images] -width [width] -height [height] -size[max_size_in_kb]
```

## Test the Utility
To test this program you can use the images directory that has few sample images.
- To Only resize the images command: 
```go
go run main.go -path images -width 600 -height 800
```

- To only compress the images command:
```go
go run main.go -path images -size 2000
```

- To do both resize and compress:
```go
go run main.go -path images -width 600 -height 800 -size 800
```

## Example Output in CLI

```bash
Failed to process image images/poha.jpg: image: unknown format
Failed to process image images/rista.jpg: image: unknown format
Processed image successfully images/vindaloo-actual.jpg
Processed image successfully images/vindaloo.jpeg
 Image processing complete. Results saved to successes.json and failures.json.
```

## Screenshots:

<img width="699" alt="image" src="https://github.com/ParulK-bhardwaj/resizify/assets/111934039/a4882d77-8f45-42e2-a47f-91453f3d5116">

<img width="579" alt="image" src="https://github.com/ParulK-bhardwaj/resizify/assets/111934039/3868dc12-a7b1-42ad-bc4e-d72f5dadd54c">

<img width="667" alt="image" src="https://github.com/ParulK-bhardwaj/resizify/assets/111934039/ea914bc7-26fa-4f83-aef2-fdb9ae565d10">

