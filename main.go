package main

import (
	"fmt"
	"io"
	"link"
	"os"
)

func main() {
	var r io.Reader
	r, _ = os.Open("e.html")
	li, err := link.Connect(r)
	if err != nil {
		panic(err)
	}
	for _, a := range li {
		fmt.Println("Href : ", a.Href)
		fmt.Println("Text : ", a.Text)
	}
}
