package main

import (
	"time"
	"fmt"
	"strconv"
	"flag"
	"os"
	"rickb777/timestamp/util"
	"strings"
)

var version = flag.Bool("V", false,
	"Prints the build version information.")

var uppercase = flag.Bool("uppercase", false,
	"Prints the timestamp using 0-9A-Z (default is 0-9a-z).")

var precision = flag.String("precision", "min",
	"The precision of the generated timestamp, one of: day, hour, min, sec, s, ms, us, µs, ns")

var layout = flag.String("layout", "2006-01-02",
	"The format string for the -zero date/time. See http://golang.org/pkg/time/#pkg-constants for more info.")

var zero = flag.String("zero", "",
			"A date/time to be used as the zero point for time calculations. " +
					"For consistent results, choose a fixed data and do not choose a future date! " +
				"The default is the Unix epoch (1st January 1970). Format: yyyy-mm-dd unless you use -layout")

var base = flag.Int("base", 36,
	"The number base used for the output string, 2 to 36. The default is 36, which is a short string.")

func main() {
	flag.Parse()
	if *version {
		fmt.Printf("tip %s, branch %s, date %s\n", util.HgTip, util.HgBranch, util.BuildDate)
		fmt.Println(util.HgPath)
		os.Exit(0)

	}

	var divisor int64
	switch *precision {
		case "day": divisor = int64(24 * 60 * 60 * 1000 * 1000 * 1000)
		case "hour": divisor = int64(60 * 60 * 1000 * 1000 * 1000)
		case "min": divisor = int64(60 * 1000 * 1000 * 1000)
		case "sec": divisor = int64(1000 * 1000 * 1000)
		case "s": divisor = int64(1000 * 1000 * 1000)
		case "ms": divisor = int64(1000 * 1000)
		case "us": divisor = int64(1000)
		case "µs": divisor = int64(1000)
		case "ns": divisor = int64(1)
	default:
		flag.Usage()
		os.Exit(1)
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
	nowS := nowNs / divisor
	tstamp := strconv.FormatInt(nowS, *base)
	if *uppercase {
		tstamp = strings.ToUpper(tstamp)
	}
	fmt.Println(tstamp)
}
