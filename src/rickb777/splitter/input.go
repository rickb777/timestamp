package main

import (
	"bufio"
	"os"
	"log"
	"strings"
	"compress/gzip"
	"compress/bzip2"
	"io"
)

func readFile(name string, consume func (string)) {

	in, err := os.Open(name) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()
	var r io.Reader = in

	if strings.HasSuffix(name, ".gz") {
		r, err = gzip.NewReader(r)
		if err != nil {
			log.Fatal(err)
		}
	} else if strings.HasSuffix(name, ".bz2") {
		r = bzip2.NewReader(r)
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		consume(scanner.Text())
	}
}
