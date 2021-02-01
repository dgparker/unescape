package main

import (
	"fmt"
	"html"
	"os"
)

func main() {
	fmt.Println("🏃🏃🏃🏃🏃🏃🏃")
	args := os.Args[1:]
	res := []string{}
	for _, v := range args {
		res = append(res, html.UnescapeString(v))
	}

	for _, v := range res {
		fmt.Println(v)
	}

	os.Exit(1)
}
