package main

import "testing"

func TestFormatNumber(t *testing.T) {
	var cases = []struct {
		v int64
		b int
		expected string
	}{
		{0, 10, "0"},
		{1, 10, "1"},
		{1234567890, 10, "1234567890"},
		{1024, 16, "400"},
		{1023, 16, "3ff"},
		{-1023, 16, "-3ff"},
		{1024, 20, "2b4"},
		{123456, 20, "f8cg"},
		{123456, 36, "2n9c"},
		{64, 32, "20"},
		{64, 64, "10"},
		{65, 65, "10"},
		{4225, 65, "100"},
		{4224, 65, "--"},
		{4223, 65, "-_"},
		{4222, 65, "-."},
		{4221, 65, "-Z"},
		{238328, 62, "1000"},
		{238327, 62, "ZZZ"},
		{-238327, 62, "-ZZZ"},
	}
	for _, c := range cases {
		actual, err := FormatNumber(c.v, c.b)
		if err != nil {
			t.Fatal(err)
		} else if actual != c.expected {
			t.Errorf("%12d base %2d: want %s but got %s", c.v, c.b, c.expected, actual)
		}
	}
}

