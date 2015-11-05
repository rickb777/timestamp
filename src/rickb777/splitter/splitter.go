package main

import (
	"fmt"
	"flag"
	"os"
)

var version = flag.Bool("V", false,
	"Prints the build version information.")

var par = flag.Bool("par", false,
	"Parallel reading of the source files.")

var lines = flag.Int("lines", 1000,
	"Prints the timestamp using 0-9A-Z (default is 0-9a-z).")

var tmpdir = flag.String("tmpdir", "~/splitter-tmp",
	"The precision of the generated timestamp, one of: day, hour, min, sec, s, ms, us, µs, ns")

func main() {
	flag.Parse()
	if *version {
		fmt.Printf("tip %s, branch %s, date %s\n", HgTip, HgBranch, BuildDate)
		fmt.Println(HgPath)
		os.Exit(0)

	}

	files := flag.Args()
	if len(files) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	for _, arg := range files {
		//		if *par {
		//			go splitFile(arg)
		//		} else {
		readFile(arg, func(s string) {
			fmt.Println(s)
		})
		//		}
	}
}
