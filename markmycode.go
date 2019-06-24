package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fileName := flag.String("file", "", "File to be parsed (required)")
	progLang := flag.String("lang", "", "Programming language of the source file (required)")
	style := flag.String("css", "", "stylesheet (optional)")

	if len(os.Args) < 3 {
		fmt.Println("Number of arguments is invalid. Usage is -file <filename> -lang <language>")
		os.Exit(1)
	}
	flag.Parse()
	if *fileName == "" {
		fmt.Println("A valid file name is required.")
		os.Exit(1)
	}
	data, err := ioutil.ReadFile(*fileName)
	check(err)
	fmt.Println(progLang,style)
	fmt.Print(string(data))
}
