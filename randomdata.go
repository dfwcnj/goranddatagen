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

func randSeq(n uint) string {
	b := make([]rune, n)
	ll := len(letters)
	for i := range b {
		b[i] = letters[rand.Intn(ll)]
	}
	return string(b)
}

func randomstrings(n uint, slen uint) []string {
	ssl := make([]string, 0)
	for _ = range n {
		ssl = append(ssl, randSeq(slen))
	}
	return ssl
}

func randomuints(n uint) []uint64 {
	usl := make([]uint64, 0)
	for _ = range n {
		fmt.Println(rand.Uint64())
		usl = append(usl, rand.Uint64())
	}
	return usl
}

func randomdates(n uint, format string) []string {
	now := time.Now().Unix()
	var mod = int64(now)
	var s string
	dsl := make([]string, 0)
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
	return dsl
}
