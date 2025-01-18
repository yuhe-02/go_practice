package main
// TODO create godoc
// TODO create user definition
// TODO create gomodules

import  (
		"fmt"
		"os"
		"path/filepath"
)

// execute printf
func main() {
	const b_extension = "jpg"
	const a_extension = "png"
	if (len(os.Args) <= 1) {
		fmt.Println("error: invalid argument")
		return
	}
	var image_path = os.Args[1]
	err := filepath.Walk(image_path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return nil
		}
		if info.IsDir() {
			fmt.Println("Directory:", path)
		} else {
			fmt.Println("File:", path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking the path %v\n", err)
		return
	}
	fmt.Println("Directory walk complete.")
}