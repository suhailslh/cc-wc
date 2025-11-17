package cmd

import (
	"testing"
	"regexp"
	"os"
	"bytes"
	"bufio"
	"io"
)

// TestRunFile calls cmd.Run with a filename, checking
// for a valid return value.
func TestRunFileName(t *testing.T) {
	fileName := "test.txt"
	want := regexp.MustCompile(`7145\t58164\t342190\t` + fileName)
	
	result, err := Run(nil, false, false, false, false, fileName)
	
	if !want.MatchString(result) || err != nil {
		t.Errorf(`Run() = %q, %v, want match for %#q, nil`, result, err, want)
	}
}

// TestRunStdin calls cmd.Run with standard input, checking
// for a valid return value.
func TestRunStdin(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		t.Errorf(`%v`, err)
		return
	}
	defer file.Close()
	
	stat, err := file.Stat()
	if err != nil {
		t.Errorf(`%v`, err)
		return
	}

	data := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(data)
	if err != nil && err != io.EOF {
		t.Errorf(`%v`, err)
		return
	}

	want := regexp.MustCompile(`7145\t58164\t342190\t`)
	
	result, err := Run(bytes.NewBuffer(data), false, false, false, false, "")
	
	if !want.MatchString(result) || err != nil {
		t.Errorf(`Run() = %q, %v, want match for %#q, nil`, result, err, want)
	}
}
