package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type plist []string

func getInput(marker int) string {
	reader := bufio.NewReader(os.Stdin)
	if marker == 0 {
		fmt.Print("Please enter the name of the file in which your proxies are located (e.g. proxies.txt): ")
		input, _ := reader.ReadString('\n')
		return strings.TrimSpace(input)
	} else if marker == 1 {
		fmt.Print("Please enter the range which you want to filter (e.g. 145): ")
		input, _ := reader.ReadString('\n')
		return strings.TrimSpace(input)
	} else {
		return ("proxy.txt")
	}
}

func readFile(filename string, d string) plist {

	l := plist{}

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Split(scanner.Text(), ".")[0] != d {
			l = append(l, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return plist(l)
}

func (l plist) dumpToFile(filename string) {
	file, err := os.Create(filename)

	w := bufio.NewWriter(file)

	for _, s := range l {
		w.WriteString(s + "\n")
	}

	w.Flush()

	if err != nil {
		log.Fatal(err)
	}
}
