# Resizify: A Batch Image Resizer Utility Project

## The Problem
In web development and digital content management, images are essential but can also be a source of inefficiency. Managing and resizing numerous images to meet specific size requirements for various platforms is time-consuming and resource-intensive. This manual resizing process not only slows down project timelines but also risks inconsistency in image quality and dimensions, which can negatively impact website performance and user experience.

## Goals
### Proposed Solution
"Resizify" is a proposed utility designed to automate the process of resizing images in bulk. This tool will allow users to specify target dimensions and resize multiple images simultaneously, ensuring that all images meet the necessary specifications.

### Features
- Batch Processing: Resize hundreds of images in one go, saving time and effort.
- Support for Multiple Formats: Compatible with popular image formats such as JPEG, PNG, and GIF.
- Easy-to-Use Interface: Simple command-line interface for quick operations, with potential for a graphical interface in future versions.
- Might add more later

## Audience
- Web designer/developers
- Digital Marketers 
- Small Businesses

## Benefits
- Efficiency: Significantly reduces the amount of time and effort required to prepare images for various platforms.
- Consistency: Ensures all images are uniformly resized, improving the visual consistency across digital assets.
- Accessibility: Easy enough for non-technical users, making it a versatile tool across different departments.
- Resource Optimization: Helps in optimizing images for web use, leading to faster page load times and improved SEO performance.

## Used Existing Libraries
- fmt : Package fmt used to format I/O with functions analogous to C's printf and scanf. Helps in printing and formatting success or failure messages.
- os : provides a platform-independent interface to operating system functionality. Used for opening and creating files and handling command line arguments.
- image : To decode and encode images in GIF, JPEG, PNG formats.
- path/filepath: To iterate over files in a directory
- strconv: To convert str to integer for command line arguments for width and height

### Third-party package:
- "github.com/nfnt/resize": To resize an image to the specified width and height. Resize supports many interpolation methods. We used resize.Lanczos3, which is well-regarded for producing high-quality resized images.

## Conclusions