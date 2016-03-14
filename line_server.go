package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var a []int64
var cache []string
var file *os.File

func handler(w http.ResponseWriter, r *http.Request) {
	var value string

	fmt.Printf("%s\n", r.URL.Path)
	tokens := strings.Split(r.URL.Path, "/")
	index, err := strconv.ParseInt(tokens[2], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Looking for index %d.\n", index)

	if int64(len(cache)) >= index && cache[index] != "" {
		value = cache[index]
		fmt.Printf("Found value in cache")
	} else {
		s := make([]byte, a[index]-a[index-1], a[index]-a[index-1])
		file.ReadAt(s, a[index-1])
		value = string(s)
	}

	fmt.Fprintf(w, value)

}

func main() {
	var current_pos int64 = 0
	var err error
	var filename string

	if len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		fmt.Printf("usage %s filename", os.Args[0])
		log.Fatal("No filename")
	}

	file, err = os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		a = append(a, current_pos)
		current_pos += int64(len(scanner.Bytes()) + 1)
		//Uncomment this to fully cache the file
		//cache = append(cache, scanner.Text())
	}
	file.Close()

	file, err = os.Open(filename)

	fmt.Printf("Filename: %s\n", filename)
	fmt.Printf("Number of Lines: %d\n", len(a))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/lines/", handler)
	http.ListenAndServe(":4567", nil)

}
