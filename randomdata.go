// package goranddatagen generates random string, uint64, and datetime data
package main

//package goranddatagen

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// randSeq(n int)
// generate a random string length n with lower, upper case letters and digits
func randSeq(slen int, rlen bool) string {
	b := make([]rune, slen)
	ll := len(letters)
	for i := range b {
		b[i] = letters[rand.Intn(ll)]
	}
	if rlen == true {
		rl := rand.Intn(slen)
		if rl != 0 {
			b = b[:rl]
		}
	}
	return string(b)
}

// randomstrings(n int64, slen int)
// generate n random strings with length slen
// return a slice containing the strings
func randomstrings(n int64, slen int, rlen bool, emit bool) []string {
	ssl := make([]string, 0)
	if emit == true {
		for _ = range n {
			fmt.Println(randSeq(slen, rlen))
		}
	} else {
		for _ = range n {
			ssl = append(ssl, randSeq(slen, rlen))
		}
	}
	return ssl
}

// randomints(ņ int)
// generate n random int64 values
// return a slice containing the int64 values
func randomuints(n int64, emit bool) []uint64 {
	usl := make([]uint64, 0)
	if emit == true {
		for _ = range n {
			fmt.Println(rand.Uint64())
		}
	} else {
		for _ = range n {
			usl = append(usl, rand.Uint64())
		}
	}
	return usl
}

// randomdates(n int64, format string)
// generate n random dates with format
// return a slice containing the random date strings
func randomdates(n int64, format string, emit bool) []string {
	now := time.Now().Unix()
	var mod = int64(now)
	var s string
	dsl := make([]string, 0)
	if emit == true {
		for _ = range n {
			ri := rand.Int63() % mod
			tm := time.Unix(int64(ri), int64(0))

			switch format {
			case "DateTime":
				fmt.Println(tm.Format(time.DateTime))
			case "Layout":
				fmt.Println(tm.Format(time.Layout))
			case "RubyDate":
				fmt.Println(tm.Format(time.RubyDate))
			case "UnixDate":
				fmt.Println(tm.Format(time.UnixDate))
			case "RFC3339":
				fmt.Println(tm.Format(time.RFC3339))
			default:
				fmt.Println(tm)
			}
		}
	} else {
		for _ = range n {
			ri := rand.Int63() % mod
			tm := time.Unix(int64(ri), int64(0))

			switch format {
			case "DateTime":
				s = fmt.Sprint(tm.Format(time.DateTime))
			case "Layout":
				s = fmt.Sprint(tm.Format(time.Layout))
			case "RubyDate":
				s = fmt.Sprint(tm.Format(time.RubyDate))
			case "UnixDate":
				s = fmt.Sprint(tm.Format(time.UnixDate))
			case "RFC3339":
				s = fmt.Sprint(tm.Format(time.RFC3339))
			default:
				s = fmt.Sprint(tm)
			}
			dsl = append(dsl, s)
		}
	}
	return dsl
}
