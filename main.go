package main

import (
	"flag"
	"fmt"

	//	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	var dtype string
	var format string
	var slen uint
	var n uint
	flag.StringVar(&dtype, "datatype", "string", "type of data to sort - string, uint64, datetime")
	flag.StringVar(&format, "format", "RFC3339", "if datatype is datetime, what date format")
	flag.UintVar(&slen, "len", 32, "lengthfor random strings")
	flag.UintVar(&n, "n", 1<<20, "number of random data elements to generate")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	if dtype == "string" {
		ssl := randomstrings(n, slen)
		for _, l := range ssl {
			fmt.Println(l)
		}
	} else if dtype == "uint64" {
		usl := randomuints(n)
		for _, u := range usl {
			fmt.Println(u)
		}
	} else if dtype == "datetime" {
		ssl := randomdates(n, format)
		for _, l := range ssl {
			fmt.Println(l)
		}
	} else {
		log.Fatal("datatype may only be string, uint64, or datetime")
	}
}
