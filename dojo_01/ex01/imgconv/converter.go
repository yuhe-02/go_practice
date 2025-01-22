// Package converter provides functionality to convert images from one format to another.
package imgconv

// module imports
import (
	"fmt"
	"image"
	_ "image/gif"  // Import GIF decoder
	_ "image/jpeg" // Import JPEG decoder
	"image/png"
	"os"
	"path/filepath"
)

// ImageConverter is a struct that holds methods for converting images.
type ImageConverter struct{}

// execConvert is a private method that handles the actual image conversion process.
// It takes a list of file paths and converts images of a given extension to the target extension.
func (ic *ImageConverter) execConvert(files []string, bExtension string, aExtension string) error {
	// Loop through each file and process the conversion
	for _, filePath := range files {
		// Open the image file
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		defer file.Close()

		// Decode the image (it will automatically detect the format)
		img, _, err := image.Decode(file)
		if err != nil {
			return fmt.Errorf("%s is not a valid file", filePath)
		}

		// Check if the file extension matches the expected source extension
		if filepath.Ext(filePath) == bExtension {
			// Create the output file with the new extension
			outputPath := filePath[:len(filePath)-len(filepath.Ext(filePath))] + aExtension
			outFile, err := os.Create(outputPath)
			if err != nil {
				return fmt.Errorf("%v", err)
			}
			defer outFile.Close()

			// Encode the image into PNG format
			err = png.Encode(outFile, img)
			if err != nil {
				return fmt.Errorf("%v", err)
			}
		}
	}
	return nil
}

// ConvertImg is a public method that initiates the image conversion process.
// It walks through the given directory, collects all image files, and converts them using the execConvert method.
func (ic *ImageConverter) ConvertImg(imagePath string, bExtension string, aExtension string) error {
	var files []string

	// Walk through the directory to find all files
	err := filepath.Walk(imagePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		// If the file is not a directory, add it to the list of files
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("walking the path %v", err)
	}

	// Call the execConvert method to perform the actual conversion
	return ic.execConvert(files, bExtension, aExtension)
}
