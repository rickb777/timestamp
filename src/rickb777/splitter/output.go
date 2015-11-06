package main

import (
	"strings"
	"path"
	"fmt"
	"io"
	"strconv"
	"os"
	"compress/gzip"
)

func outputName(file, chunkStr string) string {
	name := path.Base(file)
	parts := strings.Split(name, ".")
	if parts[len(parts)-1] == "gz" {
		parts = parts[:len(parts)-1]
	}
	andGz := ""
	if *gz {
		andGz = ".gz"
	}
	return fmt.Sprintf("%s/%s-%06s.%s%s", outdir, parts[0], chunkStr, parts[len(parts)-1], andGz)
}

func openNextOutputFile(name string) (out io.WriteCloser) {
	var err error
	chunkNum++
	file := outputName(name, strconv.FormatInt(int64(chunkNum), 36))
	fmt.Printf("Writing %s\n", file)
	out, err = os.Create(file)
	checkErrFatal(err)
	if *gz {
		out, err = gzip.NewWriterLevel(out, gzip.BestCompression)
		checkErrFatal(err)
	}
	return out
}

