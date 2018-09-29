package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jtscuba/wikipedia_corpus/bzip2_fork"
)

func main() {
	// filename := os.Args[1]

	// in, err := os.Open(filename)
	//
	// if err != nil {
	// 	log.Fatal("Unable to read zip", err)
	// }
	// defer in.Close()
	in, _ := os.Open("bzip2_fork/testdata/enwiki-latest-pages-articles-multistream.xml.bz2")
	// bufReader := bufio.NewReaderSize(in, 1<<20)

	decompressed := bzip2.NewReader(in)
	// xmlStream := xml.NewDecoder(in)

	fmt.Println("processing the entire file file")

	// currentElement := ""
	// currentTitle := ""
	//
	// for {
	// 	_, err := xmlStream.RawToken()
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			log.Println("End of file.")
	// 			break
	// 		}
	//
	// 		log.Fatal("Error reading zip", err)
	// 	}
	//
	// 	fmt.Println(&token)
	//
	// 	if start, ok := token.(xml.StartElement); ok {
	// 		currentElement = start.Name.Local
	//
	// 		if currentElement == "page" {
	// 			currentTitle = ""
	// 		}
	// 	} else if _, ok := token.(xml.EndElement); ok {
	// 		currentElement = ""
	// 	}
	//
	// 	if char, ok := token.(xml.CharData); ok {
	// 		if currentElement == "title" {
	// 			currentTitle = string(char)
	// 		} else if currentElement == "text" {
	// 			text := string(char)
	// 			if strings.Contains(text, "{{Infobox musical artist") {
	// 				fmt.Print(currentTitle)
	// 				fmt.Print("\n")
	// 			}
	// 		}
	// 	}
	// }

	// read raw bytes from in
	buf := make([]byte, 1024*1024*4) // 1M buffer
	var totalBytes int64
	for {
		n, err := decompressed.Read(buf)
		totalBytes += int64(n)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
	}
	fmt.Println(totalBytes)
}
