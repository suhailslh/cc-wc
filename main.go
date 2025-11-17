package main

import (
	"fmt"
	"github.com/suhailslh/cc-wc/cmd"
	"os"
	"flag"
)

func main() {
	byteFlag := flag.Bool("c", false, "display number of bytes")
	lineFlag := flag.Bool("l", false, "display number of lines")
	wordFlag := flag.Bool("w", false, "display number of words")
	charFlag := flag.Bool("m", false, "display number of characters")
	flag.Parse()

	result, err := cmd.Run(os.Stdin, *byteFlag, *lineFlag, *wordFlag, *charFlag, flag.Arg(0))
	if err != nil {
		return
	}

	fmt.Print(result)
}
