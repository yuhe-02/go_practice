package main

// TODO create godoc
// TODO create user definition

import (
	"convert/imgconv"
	"fmt"
	"os"
)

// execute Error
func main() {
	const bExtension = ".jpg"
	const aExtension = ".png"
	if len(os.Args) != 2 {
		fmt.Printf("error: invalid argument\n")
		return
	}
	var imagePath = os.Args[1]
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		fmt.Printf("error: %s: no such file or directory\n", imagePath)
		return
	}
	ic := &converter.ImageConverter{}
	err := ic.ConvertImg(imagePath, bExtension, aExtension)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}
	return
}
