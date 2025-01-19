package main
// TODO create godoc
// TODO create user definition
// TODO create gomodules

import  (
		"fmt"
		"os"
		"path/filepath"
		"image"
		_ "image/jpeg" 
		"image/png"    
)

func execConvert(files []string, bExtension string, aExtension string) error {
	for _, filePath := range files {
		file, err := os.Open(filePath)
		if (err != nil) {
			return fmt.Errorf("%v", err)
		}
		img, _, err := image.Decode(file)
		if (err != nil) {
			return fmt.Errorf("%s is not a valid file", filePath)
			
		}
		if (filepath.Ext(filePath) == bExtension) {
			outputPath := filePath[:len(filePath)-len(filepath.Ext(filePath))] + aExtension
			outFile, err := os.Create(outputPath)		
			if (err != nil) {
				return fmt.Errorf("%v", err)
				 
			}
			err = png.Encode(outFile, img)
			if (err != nil) {
				return fmt.Errorf("%v", err)
			}
		}
	}
	return nil
}

func convertImg(imagePath string, bExtension string, aExtension string) error {
	var files []string

	err := filepath.Walk(imagePath, func(path string, info os.FileInfo, err error) error {
		if (err != nil) {
			return fmt.Errorf("%v", err)
		}
		if (!info.IsDir()) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("walking the path %v", err)
	}
	return execConvert(files, bExtension, aExtension)
}

// execute Error   
func main() {
	const bExtension = ".jpg"
	const aExtension = ".png"
	if (len(os.Args) != 2) {
		fmt.Printf("error: invalid argument")
		return
	}
	var imagePath = os.Args[1]
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		fmt.Printf("error: %s: no such file or directory\n", imagePath)
		return
	}
	err := convertImg(imagePath, aExtension, bExtension)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}
	return 
}