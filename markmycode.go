package main

import (
	"flag"
	"fmt"
	"os"
	//"regexp"
	"strings"
	"bufio"
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

	if _,err := os.Stat(*fileName); err != nil{
		if os.IsNotExist(err){
			fmt.Println("error: file does not exist")
			os.Exit(1)
		}
	}

	data,err := os.Open(*fileName)
	check(err)
	defer data.Close()

	fmt.Println(*progLang,*style)

	scanner := bufio.NewScanner(data)

	//re := regexp.MustCompile("(\\/\\*([^*]|[\r\n]|(\\*+([^*\\/]|[\r\n])))*\\*+\\/|\\/\\/.*\n?)")

	for scanner.Scan(){
		line := strings.TrimSpace(scanner.Text())
		fmt.Println(scanner.Text())
	}

	
}

func parse(html string)string{
	return ""
}