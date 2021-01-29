// Command bech32-decode decodes bech32 from arguments or lines of
// stdin.
//
// Usage:
//
//     bech32-decode BECH32..
//     bech32-decode <LINES_OF_BECH32
//
// The tag ("human readable part") is not output.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"eagain.net/go/bech32"
)

var prog = filepath.Base(os.Args[0])

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "  %s BECH32..\n", prog)
	fmt.Fprintf(os.Stderr, "  %s <LINES_OF_BECH32\n", prog)
	flag.PrintDefaults()
}

func process(input string) error {
	tag, decoded, err := bech32.Decode(input)
	if err != nil {
		return fmt.Errorf("decoding input: %q: %v", input, err)
	}
	// in the future, we can offer a flag to e.g. enforce known tag
	_ = tag
	if _, err := os.Stdout.Write(decoded); err != nil {
		return err
	}
	return nil
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(prog + ": ")

	flag.Usage = usage
	flag.Parse()

	if flag.NArg() > 0 {
		for _, input := range flag.Args() {
			if err := process(input); err != nil {
				log.Fatal(err)
			}
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			if err := process(input); err != nil {
				log.Fatal(err)
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("reading standard input: %v", err)
		}
	}
}
