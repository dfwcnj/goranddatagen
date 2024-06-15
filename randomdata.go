package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randomstrings(n int64, len int) {
	for _ = range n {
		fmt.Println(randSeq(len))
	}
}

func randomuints(n int64) {
	for _ = range n {
		fmt.Println(rand.Uint64())
	}
}

func randomdates(n int64, format string) {
	var mod = int64(306327747)
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
}
