// Package converter
package converter

// module import
import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// ImageConverter
type ImageConverter struct {
}

// execConvert
func (ic *ImageConverter) execConvert(files []string, bExtension string, aExtension string) error {
	for _, filePath := range files {
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		defer file.Close()
		img, _, err := image.Decode(file)
		if err != nil {
			return fmt.Errorf("%s is not a valid file", filePath)
		}
		if filepath.Ext(filePath) == bExtension {
			outputPath := filePath[:len(filePath)-len(filepath.Ext(filePath))] + aExtension
			outFile, err := os.Create(outputPath)
			if err != nil {
				return fmt.Errorf("%v", err)
			}
			defer outFile.Close()
			err = png.Encode(outFile, img)
			if err != nil {
				return fmt.Errorf("%v", err)
			}
		}
	}
	return nil
}

// ConvertImg
func (ic *ImageConverter) ConvertImg(imagePath string, bExtension string, aExtension string) error {
	var files []string

	err := filepath.Walk(imagePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("walking the path %v", err)
	}
	return ic.execConvert(files, bExtension, aExtension)
}
