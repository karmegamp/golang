package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var fileArg = flag.String("f", "fileread.go", "lookup file name")

func ReadFile1(f *os.File, lines map[string]int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		lines[text]++
	}
}

func ReadFile2(f *os.File, lines map[string]int) {
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

	ReadFile1(f, lines)

	for k, v := range lines {
		fmt.Printf("\n")
		fmt.Print(k, ":", v)
	}

	fmt.Println("\n\n---- end of read ------")

	lines = make(map[string]int)
	ReadFile2(f, lines)

	for k, v := range lines {
		fmt.Printf("\n")
		fmt.Print(k, ":", v)
	}
}
