package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/dfwcnj/randomdata"
)

func main() {
	var rlen, nl bool
	var dtype, format string
	var slen int
	var n int64

	flag.BoolVar(&nl, "nl", false, "emit strings with newlines")
	flag.BoolVar(&rlen, "rlen", false, "random lengths")
	flag.StringVar(&dtype, "datatype", "string", "type of data to sort - string, uint64, datetime")
	flag.StringVar(&format, "format", "RFC3339", "if datatype is datetime, what date format")
	flag.IntVar(&slen, "len", 32, "max length for random strings")
	flag.Int64Var(&n, "n", 1<<20, "number of random data elements to generate")
	flag.Parse()
	if rlen {
		nl = true
	}

	rand.Seed(time.Now().UnixNano())

	if dtype == "string" {
		lns := randomdata.Randomstrings(n, slen, rlen)
		if nl {
			for _, l := range lns {
				fmt.Println(l)
			}
		} else {
			for _, l := range lns {
				fmt.Print(l)
			}
		}
	} else if dtype == "uint64" {
		lns := randomdata.Randomuints(n)
		for _, l := range lns {
			fmt.Print(l)
		}
	} else if dtype == "datetime" {
		lns := randomdata.Randomdates(n, format, false)
		for _, l := range lns {
			fmt.Print(l)
		}
	} else {
		log.Fatal("datatype may only be string, uint64, or datetime")
	}
}
