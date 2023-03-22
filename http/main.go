package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Print("An Error occured;", err)
		os.Exit(1)
	}
	bs := make([]byte, 99999)
	n, err := resp.Body.Read(bs)
	if err != nil {
		fmt.Println("Error occured")
	}
	fmt.Print("n:", n)
	fmt.Print("Byte slice:", string(bs))

	io.Copy(os.Stdout, resp.Body)
	//fmt.Printf("res:%v. err:%v", resp, err)
}
