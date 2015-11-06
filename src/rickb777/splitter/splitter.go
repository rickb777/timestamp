package main

import (
	"fmt"
	"flag"
	"os"
	"strings"
	"path"
	"log"
	"io"
)

var version = flag.Bool("V", false,
	"Prints the build version information.")

var gz = flag.Bool("gz", false,
	"Gzip the output files.")

//var par = flag.Bool("par", false,
//	"Parallel reading of the source files.")

var chunkSizeP = flag.Int("chunk", 0,
	"Chunk the output into many files, each of the number of lines specified here. Disabled if less than one.")
var chunkSize int

var outdirP = flag.String("outdir", "~/splitter-tmp",
	"Output directory; prefix with '~/' for home folder.")
var outdir string

//-------------------------------------------------------------------------------------------------

var chunkNum = 0

func checkErrFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	if *version {
		fmt.Printf("tip %s, branch %s, date %s\n", HgTip, HgBranch, BuildDate)
		fmt.Println(HgPath)
		os.Exit(0)
	}

	chunkSize = *chunkSizeP

	outdir = path.Clean(*outdirP)
	if strings.HasPrefix(outdir, "~/") {
		outdir = os.Getenv("HOME") + outdir[1:]
	}
	os.MkdirAll(outdir, 0755)

	files := flag.Args()
	if len(files) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	var out io.WriteCloser

	for _, arg := range files {

		line := 0
		readFile(arg, func(s string) {
			line++
			if out == nil {
				out = openNextOutputFile(arg)
			}
			fmt.Fprintln(out, s)
			if line >= chunkSize && chunkSize > 0 {
				checkErrFatal(out.Close())
				out = nil
				line = 0
			}
		})
	}

	checkErrFatal(out.Close())
}
