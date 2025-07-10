package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var fileArg = flag.String("f", "fileread.go", "lookup file name")

func ReadLineByLine(f *os.File, lines map[string]int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		lines[text]++
	}
}

func ReadWholeFile(f *os.File, lines map[string]int) {
	file, _ := os.ReadFile(f.Name())
	data := strings.SplitSeq(string(file), "\n")
	for v := range data {
		k := strings.TrimSpace(v)
		if k == "" {
			continue
		}
		lines[k]++
	}

}
func main() {

	flag.Parse()

	lines := make(map[string]int)
	f, _ := os.Open(*fileArg)

	ReadLineByLine(f, lines)

	for k, v := range lines {
		fmt.Printf("\n")
		fmt.Print(k, ":", v)
	}

	fmt.Println("\n\n---- end of read ------")

	lines = make(map[string]int)
	ReadWholeFile(f, lines)

	for k, v := range lines {
		fmt.Printf("\n")
		fmt.Print(k, ":", v)
	}
}
