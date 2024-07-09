package main

//package goranddatagen

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	var dtype string
	var format string
	var slen int
	var rlen bool
	var n int

	flag.BoolVar(&rlen, "rlen", true, "random lengtha")
	flag.StringVar(&dtype, "datatype", "string", "type of data to sort - string, uint64, datetime")
	flag.StringVar(&format, "format", "RFC3339", "if datatype is datetime, what date format")
	flag.IntVar(&slen, "len", 32, "max length for random strings")
	flag.IntVar(&n, "n", 1<<20, "number of random data elements to generate")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	if dtype == "string" {
		ssl := randomstrings(n, slen, rlen)
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
