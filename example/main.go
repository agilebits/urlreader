package main

import (
	"fmt"
	"io/ioutil"

	"github.com/agilebits/urlreader"
)

func main() {
	// url := "https://twitter.com/1Password"
	url := "file://./main.go"

	reader, err := urlreader.Open(url)
	if err != nil {
		panic(err)
	}

	defer reader.Close()
	result, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))
}
