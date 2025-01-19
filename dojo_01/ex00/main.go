package main

import (
	"fmt"
	"io"
	// "strings"
	"os"
	"bufio"
)

func ft_cat(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		_, err := writer.Write([]byte(scanner.Text() + "\n"))
		if err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func main() {
	if (len(os.Args) == 1) {
		err := ft_cat(os.Stdin, os.Stdout)
		if err != nil {
			fmt.Printf("ft_cat: %v\n", err)
		}
		return
	}
	for _, arg := range os.Args[1:] {
		// check if directory
		info, err := os.Stat(arg)
		if (err != nil) {
			fmt.Printf("ft_cat: %s: ", arg)
			if os.IsNotExist(err) {
				fmt.Println("No such file or directory")
			} else {
				fmt.Println("unknown error")
			}
			continue
		}
		file, err := os.Open(arg)
		if (err != nil) {
			fmt.Printf("ft_cat: %s: ", arg)
			if (os.IsPermission(err)) {
				fmt.Println("Permission denied")
			} else {
				fmt.Println("unknown error")
			}
			continue
		}
		defer file.Close()
		if (info.IsDir()) {
			fmt.Printf("ft_cat: %s: ", arg)
			fmt.Println("Is a directory")
			continue
		}
		err = ft_cat(file, os.Stdout)
		if (err != nil) {
			fmt.Printf("ft_cat: %s: ", arg)
			fmt.Println("unknown error")
		}
	}
}