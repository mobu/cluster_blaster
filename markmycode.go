package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"bufio"
	"io/ioutil"
	"log"
	"runtime"
)

var wg sync.WaitGroup

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

	if *style == ""{
		*style = "none"
	}

	fmt.Printf("Language: %s, style: %s",*progLang,*style)

	fileList, err := ioutil.ReadDir(".")
	if err != nil{
		log.Fatal(err)
	}
	for _,f := range fileList{
		wg.Add(1)
		go fileOpen(f,&wg)
	}

	fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
}

//
// fileOpen function
//

func fileOpen(fileName string,wg *sync.WaitGroup){

	if _,err := os.Stat(fileName); err != nil{
		if os.IsNotExist(err){
			fmt.Println("error: file does not exist")
			os.Exit(1)
		}
	}
	file,err := os.Open(fileName)
	if err != nil{
		fmt.Println("Error opening file")
		return
	}
	defer func(){
		file.Close()
		wg.Done()
	}()

	buf := make([]byte, 32*1024)
	scanner := bufio.NewScanner(file)
	scanner.Buffer(buf,32*1024)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

}