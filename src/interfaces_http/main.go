package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error is ", err)
		os.Exit(1)
	}
	//fmt.Println(resp)

	/*
		body := make([]byte, 99999)
		n, error := resp.Body.Read(body)

		if error == nil {
			fmt.Println("Could not read the response body")
			os.Exit(1)
		}
		fmt.Println("Total Bytes Read are ", n)
		fmt.Println(string(body))
	*/
	io.Copy(os.Stdout, resp.Body)
}
