package main

import (
	"time"
	"fmt"
	"strconv"
	"flag"
	"os"
	"rickb777/timestamp/util"
)

var version = flag.Bool("V", false,
	"Prints the build version information.")

var layout = flag.String("layout", "2006-01-02",
	"The format string for the -zero date/time. See http://golang.org/pkg/time/#pkg-constants for more info.")

var zero = flag.String("zero", "",
	"A date/time to be used as the zero point for time calculations. "+
			"For consistent results, choose a fixed data and do not choose a future date! "+
			"The default is the Unix epoch (1st January 1970). Format: yyyy-mm-dd unless you use -layout")

var base = flag.Int("base", 36,
	"The number base used for the output string, 2 to 36. The default is 36, which is a short string.")

func main() {
	flag.Parse()
	if *version {
		fmt.Println(util.HgPath)
	}
	sub := int64(0)
	if *zero != "" {
		t, err := time.Parse(*layout, *zero)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		sub = t.UnixNano()
	}
	nowNs := time.Now().UnixNano() - sub
	nowS := nowNs / 1000000000
	fmt.Println(strconv.FormatInt(nowS, *base))
}
