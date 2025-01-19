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

// execute printf   
func main() {
	const bExtension = "jpg"
	const aExtension = "png"
	if (len(os.Args) != 2) {
		fmt.Println("error: invalid argument")
		return
	}
	var imagePath = os.Args[1]
	var files []string
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		fmt.Printf("error: %s: no such file or directory\n", imagePath)
		return
	}
	err := filepath.Walk(imagePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return nil
		}
		if info.IsDir() {
			fmt.Println("Directory:", path)
		} else {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking the path %v\n", err)
		return
	}
	for _, filePath := range files {
		file, err := os.Open(filePath)
		if (err != nil) {
			fmt.Printf("error: %v \n", err)
			return 
		}
		img, _, err := image.Decode(file)
		if (err != nil) {
			fmt.Printf("error: %s is not a valid file\n", filePath)
			return 
		}
		if (filepath.Ext(filePath) == ".jpg") {
			outputPath := filePath[:len(filePath)-len(filepath.Ext(filePath))] + ".png"
			outFile, _ := os.Create(outputPath)
			png.Encode(outFile, img)
		}
	}
}