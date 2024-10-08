package main

import (
	"flag"
	"log"
	"math/rand"
	"time"
	"github.com/dfwcnj/goranddatagen/randomdata"
)

func main() {
	var emit bool
	var rlen bool
	var dtype string
	var format string
	var slen int
	var n int64

	flag.BoolVar(&emit, "emit", false, "emit instead of return")
	flag.BoolVar(&rlen, "rlen", false, "random lengths")
	flag.StringVar(&dtype, "datatype", "string", "type of data to sort - string, uint64, datetime")
	flag.StringVar(&format, "format", "RFC3339", "if datatype is datetime, what date format")
	flag.IntVar(&slen, "len", 32, "max length for random strings")
	flag.Int64Var(&n, "n", 1<<20, "number of random data elements to generate")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	if dtype == "string" {
		randomstrings(n, slen, rlen, emit)
	} else if dtype == "uint64" {
		randomuints(n, emit)
	} else if dtype == "datetime" {
		randomdates(n, format, emit)
	} else {
		log.Fatal("datatype may only be string, uint64, or datetime")
	}
}
