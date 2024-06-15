
package main

import (
	"flag"
//	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	var dtype string
	var format string
	var len int
	var n int64
	flag.StringVar(&dtype, "datatype", "string", "type of data to sort - string, uint64, datetime")
	flag.StringVar(&format, "format", "RFC3339", "if datatype is datetime, what date format")
	flag.IntVar(&len, "len", 32, "lengthfor random strings")
	flag.Int64Var(&n, "n", 1<<20, "number of random data elements to generate")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	if dtype == "string" {
		randomstrings(n, len)
	} else if dtype == "uint64" {
		randomuints(n)
	} else if dtype == "datetime" {
		randomdates(n, format)
	} else {
		log.Fatal("only string, int, or datetime")
	}
}
