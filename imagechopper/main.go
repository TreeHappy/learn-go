package imagechopper

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

// ImageBuffer represents a buffer of image lines
type ImageBuffer struct {
	Lines []image.Image
}

// ChopImage divides the image into buffers of N lines
func ChopImage(img image.Image, n int) []ImageBuffer {
	var buffers []ImageBuffer
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	for y := 0; y < height; y += n {
		var buffer ImageBuffer
		for i := 0; i < n && (y+i) < height; i++ {
			line := image.NewRGBA(image.Rect(0, 0, width, 1))
			for x := 0; x < width; x++ {
				line.Set(x, 0, img.At(x, y+i))
			}
			buffer.Lines = append(buffer.Lines, line)
		}
		buffers = append(buffers, buffer)
	}

	return buffers
}

// GetImageBytes returns the bytes of the image at the specified index
func (ib *ImageBuffer) GetImageBytes(index int) ([]byte, error) {
	if index < 0 || index >= len(ib.Lines) {
		return nil, fmt.Errorf("index out of bounds")
	}

	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, ib.Lines[index], nil); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// LoadImage loads an image from a file
func LoadImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// SaveImage saves an image to a file
func SaveImage(filename string, img image.Image) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return jpeg.Encode(file, img, nil)
}

func Dodido() {
	// Load the image
	img, err := LoadImage("/tmp/11106.jpg")
	if err != nil {
		fmt.Println("Error loading image:", err)
		return
	}

	// Define the number of lines per buffer
	n := 64 // Change this value as needed

	// Chop the image into buffers
	buffers := ChopImage(img, n)

	// Example: Get bytes of the first image in the first buffer
	if len(buffers) > 0 {
		imageBytes, err := buffers[0].GetImageBytes(0)
		if err != nil {
			fmt.Println("Error getting image bytes:", err)
		} else {
			// You can now use imageBytes as needed
			fmt.Printf("Bytes of the first image in the first buffer: %d bytes\n", len(imageBytes))
		}
	}

	// Save each buffer as a separate image
	for i, buffer := range buffers {
		for j, line := range buffer.Lines {
			outputFilename := fmt.Sprintf("output_buffer_%d_line_%d.jpg", i, j)
			if err := SaveImage(outputFilename, line); err != nil {
				fmt.Println("Error saving image:", err)
			}
		}
	}

	fmt.Println("Image processing complete.")
}
