package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
	"strings"
	"regexp"
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

	if *progLang == ""{
		fmt.Println("A valid programming language is required.")
		os.Exit(1)
	}

	data,err := os.Open(*fileName)
	check(err)
	defer data.Close()
	fmt.Println(*progLang,*style)

	scanner := bufio.NewScanner(data)
	// buf := make([]byte, 0, 1024*1024)
	// scanner.Buffer(buf, 10*1024*1024)

	re := regexp.MustCompile("(\\/\\/)")
	for scanner.Scan(){
		trimmed := strings.TrimSpace(scanner.Text())
		if(re.Match([]byte(trimmed))){
			split := re.Split(trimmed,-1)
			fmt.Println(split[0])
		}
	}
}
