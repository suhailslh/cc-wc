package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

func Run(r io.Reader, byteFlag bool, lineFlag bool, wordFlag bool, charFlag bool, fileName string) (result string, err error) {
	noFlag := !(byteFlag || lineFlag || wordFlag || charFlag)
	
	data, err := readData(r, fileName)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	
	result = ""

	if lineFlag || noFlag {
		lineSep := []byte{'\n'}
		result += fmt.Sprintf("%d\t", bytes.Count(data, lineSep))
	}

	if wordFlag || noFlag {
		result += fmt.Sprintf("%d\t", len(bytes.Fields(data)))
	}

	if byteFlag || noFlag {
		result += fmt.Sprintf("%d\t", len(data))
	}

	if charFlag {
		result += fmt.Sprintf("%d\t", utf8.RuneCount(data))
	}

	result += fileName + "\n"

	return result, nil
}

func readData(r io.Reader, fileName string) (data []byte, err error)  {
	if fileName == "" {
		data, err = io.ReadAll(r)
		if err != nil {
			return nil, err
		}
		return data, nil
	} 

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	data = make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(data)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return data, nil
}
