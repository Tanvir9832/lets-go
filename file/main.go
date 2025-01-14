package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	//! Create file
	file, error := os.Create("fileByGo.txt")
	if error != nil {
		fmt.Println("File create failed! ", error)
	}
	defer file.Close()

	//! Write file
	content := "Hi I am Tanvir writing from golang ok"
	contentSize, _ := io.WriteString(file, content)
	fmt.Println(contentSize)

	//! Read file (data chunk with add the buffer array and buffer will be printed)
	files, _ := os.Open("fileByGo.txt")
	// fmt.Println(file)
	defer files.Close()
	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			// fmt.Println(n)
			break
		}
		// fmt.Println(n)
		// fmt.Println(buffer[:n])
		fmt.Println(string(buffer[:n]))
	}

	//! Read file through byte slice (complete buffer adds)
	contents, _ := os.ReadFile("fileByGo.txt")
	fmt.Println(string(contents))
}
