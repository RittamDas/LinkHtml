package main

import (
	"fmt"
	"io"
	"link"
	"os"
)

func main() {
	var r io.Reader
	var e error
	r, e = os.Open("e.html")
	if e != nil {
		panic(e)
	}
	li, err := link.Connect(r)
	if err != nil {
		fmt.Println("p")
		panic(err)
	}
	for _, a := range li {
		fmt.Println("Href : ", a.Href)
		fmt.Println("Text : ", a.Text)
	}
}
