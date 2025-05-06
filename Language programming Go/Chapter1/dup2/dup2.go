package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	names := make(map[string]int)
	counts := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "stdin", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			names[arg]++
			countLines(f, arg, counts)
			f.Close()
		}
	}

	for line, name := range counts {
		fmt.Printf("%s\t%s\n", name, line)
	}
}

func countLines(f *os.File, name string, counts map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line] = name
	}
}
