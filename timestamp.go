package main

import (
	"time"
	"fmt"
	"flag"
	"os"
)

var version = flag.Bool("V", false,
	"Prints the build version information.")

var precision = flag.String("precision", "second",
	"The precision of the generated timestamp, one of: day, hour, minute|min, second|sec|s, ms, us, µs, ns")

var layout = flag.String("layout", "2006-01-02",
	"The format string for the -zero date/time. See http://golang.org/pkg/time/#pkg-constants for more info.")

var zero = flag.String("zero", "",
	"A date/time to be used as the zero point for time calculations. " +
		"For consistent results, choose a fixed data and do not choose a future date! " +
		"The default is the Unix epoch (1st January 1970). Format: yyyy-mm-dd unless you use -layout")

var base = flag.Int("base", 10,
	"The number base used for the output string, 2 to 66." +
		"\n\tHigh bases are particularly useful in that a short string results." +
		"\n\tBases up to 62 contain only alphanumeric characters; higher bases contain URL-safe punctuation." +
		"\n\tThe characters used are " + digits)

var value = flag.Int("value", 0,
	"Sets the actual number printed. Value must be a base-10 number and will be converted to the specified output base.")

func main() {
	flag.Parse()
	if *version {
		fmt.Printf("Version %s, date %s\n", Version, BuildDate)
		os.Exit(0)
	}

	var divisor int64
	switch *precision {
	case "day": divisor = int64(24 * 60 * 60 * 1000 * 1000 * 1000)
	case "hour": divisor = int64(60 * 60 * 1000 * 1000 * 1000)
	case "minute": divisor = int64(60 * 1000 * 1000 * 1000)
	case "min": divisor = int64(60 * 1000 * 1000 * 1000)
	case "second": divisor = int64(1000 * 1000 * 1000)
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

	nowS := int64(*value)

	if *value == 0 {
		sub := int64(0)
		if *zero != "" {
			t, err := time.Parse(*layout, *zero)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
			sub = t.UnixNano()
		}
		nowNs := time.Now().UnixNano() - sub
		nowS = nowNs / divisor
	}

	tstamp, err := FormatNumber(nowS, *base)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println(tstamp)
}
