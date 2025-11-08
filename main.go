package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

func read(r io.Reader) (fileName string, data []byte, err error)  {
	if flag.NArg() == 0 {
		data, err = io.ReadAll(r)
		if err != nil {
			return "", nil, err
		}
		return "", data, nil
	} 
	
	fileName = flag.Arg(0)

	file, err := os.Open(fileName)
	if err != nil {
		return fileName, nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return fileName, nil, err
	}

	data = make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(data)
	if err != nil && err != io.EOF {
		return fileName, nil, err
	}

	return fileName, data, nil
}

func main() {
	byteFlag := flag.Bool("c", false, "display number of bytes")
	lineFlag := flag.Bool("l", false, "display number of lines")
	wordFlag := flag.Bool("w", false, "display number of words")
	charFlag := flag.Bool("m", false, "display number of characters")
	flag.Parse()

	noFlag := !(*byteFlag || *lineFlag || *wordFlag || *charFlag)
	
	fileName, data, err := read(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	if *byteFlag || noFlag {
		fmt.Printf("%d\t", len(data))
	}

	if *lineFlag || noFlag {
		lineSep := []byte{'\n'}
		fmt.Printf("%d\t", bytes.Count(data, lineSep))
	}

	if *wordFlag || noFlag {
		fmt.Printf("%d\t", len(bytes.Fields(data)))
	}

	if *charFlag {
		fmt.Printf("%d\t", utf8.RuneCount(data))
	}

	if fileName != "" {
		fmt.Print(fileName)
	}

	fmt.Println()
}
