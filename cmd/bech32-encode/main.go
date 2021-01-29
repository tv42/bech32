// Command bech32-encode encodes its standard input as bech32.
//
// Usage:
//
//     bech32-encode TAG <FILE
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"eagain.net/go/bech32"
)

var prog = filepath.Base(os.Args[0])

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "  %s TAG <FILE\n", prog)
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(prog + ": ")

	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 1 {
		usage()
		os.Exit(1)
	}

	tag := flag.Arg(0)
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	encoded, err := bech32.Encode(tag, data)
	if err != nil {
		log.Fatal(err)
	}
	encoded = encoded + "\n"
	if _, err := os.Stdout.WriteString(encoded); err != nil {
		log.Fatal(err)
	}
}
