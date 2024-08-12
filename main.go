package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	anyascii "github.com/anyascii/go"
	"github.com/docopt/docopt-go"
)

const version = "v0.1"

const usage = `
Transliterate file(s) containing Unicode characters into ASCII equivalents.

Reads from given files, or stdin by default.

Outputs to stdout.

Usage:
  asciify [FILE...]
  asciify -h | --help
  asciify --version

Global Options:
  -h, --help             Show this screen.
  --version              Show version.

Homepage: https://github.com/alasdairmorris/asciify
`

type Config struct {
	infiles []string
}

func exitOnError(e error) {
	if e != nil {
		panic(e)
	}
}

// Parse and validate command-line arguments
func getConfig() Config {

	var (
		retval Config
		opts   docopt.Opts
		err    error
	)

	opts, err = docopt.ParseArgs(usage+" ", nil, version)
	exitOnError(err)

	// Input file(s)
	if infiles, ok := opts["FILE"].([]string); ok {
		for _, i := range infiles {
			if _, err := os.Stat(i); err != nil {
				log.Fatal(fmt.Sprintf("Error accessing file %s - %s", i, err))
			}
			retval.infiles = append(retval.infiles, i)
		}
	}

	return retval
}

func main() {
	log.SetFlags(0)
	config := getConfig()
	if len(config.infiles) > 0 {

		for _, path := range config.infiles {
			fp, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer fp.Close()

			scanner := bufio.NewScanner(fp)
			for scanner.Scan() {
				fmt.Println(anyascii.Transliterate(scanner.Text()))
			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}

	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(anyascii.Transliterate(scanner.Text()))
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
