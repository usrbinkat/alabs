package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ./sha1 <file>")
		os.Exit(1)
	}

	f := os.Args[1]

	// TODO: evaluate compression type and add support for both gzip, bzip2, and xz
	s, err := sha1Sum(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", s)
}

func sha1Sum(filename string) (string, error) {
	file, err := os.Open(filename)
	logErr(err)
	defer file.Close()

	read, err := gzip.NewReader(file)
	logErr(err)
	defer read.Close()

	sha1 := sha1.New()
	if _, err := io.Copy(sha1, read); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", sha1.Sum(nil)), nil
}

func logErr(err error) (string, error) {
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return "", err
}
