package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]
	/*fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error occured while reading file-%v", fileName)
	}
	fmt.Println(string(fileContent))*/
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error occured while opening the file.")
	}
	/*data := make([]byte, 1024*2)
	fileLength, _ := file.Read(data)
	fmt.Println("Total length:", fileLength)
	fmt.Println("File contents", string(data))*/
	//data := make([]byte, 1024*2)
	//fileLength, _ := file.Read(data)
	io.Copy(os.Stdout, file)
	//fmt.Print(fileLength)
}
