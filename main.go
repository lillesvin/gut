package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const VERSION = `0.1`

var showVersion bool
var delim string
var delimDefault *regexp.Regexp = regexp.MustCompile(`\s+`)
var field string
var fieldDefault string = ""

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func parseFieldString(param string) []int {
	var fields []int

	substr := strings.Split(param, `,`)

	isRange := regexp.MustCompile(`^[0-9]+(-|\.{2,3}|:)[0-9]+$`)

	for _, v := range substr {
		if s, err := strconv.Atoi(v); err == nil {
			fields = append(fields, s-1)
			continue
		}

		if isRange.MatchString(v) {
			subsubstr := strings.Split(v, `-`)
			min, _ := strconv.Atoi(subsubstr[0])
			max, _ := strconv.Atoi(subsubstr[1])
			fields = append(fields, makeRange(min-1, max-1)...)
		}
	}

	return fields
}

func init() {
	flag.StringVar(&delim, `delimiter`, delimDefault.String(), `Regex to use as delimiter`)
	flag.StringVar(&delim, `d`, delimDefault.String(), `Short form of -delimiter`)

	flag.StringVar(&field, `fields`, fieldDefault, `Field(s) to display`)
	flag.StringVar(&field, `f`, fieldDefault, `Short form of -fields`)

	flag.BoolVar(&showVersion, `version`, false, `Show version and exit`)

	flag.Parse()
}

func main() {
	if showVersion {
		fmt.Println("gut", VERSION)
		os.Exit(0)
	}

	re := regexp.MustCompile(delim)
	fields := parseFieldString(field)

	scanner := bufio.NewScanner(os.Stdin)
	var tmp []string
	var out []string

	for scanner.Scan() {
		tmp = re.Split(scanner.Text(), -1)

		if len(fields) == 0 {
			fields = makeRange(0, len(tmp)-1)
		}

		for _, i := range fields {
			// Does field i even exist?
			if len(tmp) > i {
				out = append(out, tmp[i])
			}
		}
		fmt.Printf("%s\n", strings.Join(out, ` `))
	}
}
