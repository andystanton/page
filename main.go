package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := []string{}

	pageSizePtr := flag.Int("size", 8, "page size")
	pageCountPtr := flag.Int("num", 1, "page number")
	flag.Parse()

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin: %s", err.Error())
	}

	startIndex := *pageSizePtr * (*pageCountPtr - 1)
	endIndex := *pageSizePtr * *pageCountPtr

	if endIndex > len(input)-1 {
		endIndex = len(input)
	}

	if startIndex > endIndex {
		fmt.Fprintln(os.Stderr, "Invalid page range requested")
		os.Exit(2)
	}

	for _, line := range input[startIndex:endIndex] {
		fmt.Println(line)
	}
}
