package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
	Disclaimer:
		I got pretty close at doing this just in bash. But figured a little Golang gives me extra easy freedom for checking/mutating values when comparing

	Aim:
		To have a script that takes 2 files:
			1 a list of everything
			1 a list of items that are specifically needed to be removed

	Result:
		Spit out the URL's that are left in scope to be piped elsewhere

	Usage:
		./scopepurge allurls.txt outofscope.txt | ./anothertool
		./scopepurge allurls.txt outofscope.txt > urls.txt
*/

func main() {
	allUrlsFilePath := os.Args[1]
	outOfScopeUrlsFilePath := os.Args[2]

	allUrls, err := readLines(string(allUrlsFilePath))
	if err != nil {
		log.Fatalln(err)
		return
	}

	outOfScopeUrls, err := readLines(string(outOfScopeUrlsFilePath))
	if err != nil {
		log.Fatalln(err)
		return
	}

	finalUrls := []string{}

	for _, u := range allUrls {
		if !exists(outOfScopeUrls, u) {
			// This assumes the URLs identified are the EXACT same as the out of scope URLs.
			// To help, this should generate mutations and possible options for the scope (http://, https:// specifically)
			finalUrls = append(finalUrls, u)
		}
	}

	for _, u := range finalUrls {
		fmt.Println(u)
	}
}

func exists(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
