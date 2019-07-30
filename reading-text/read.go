package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	fileHandle, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	regex := regexp.MustCompile(`class=\"([a-zA-Z ]*)\"`)

	i := 0
	for fileScanner.Scan() {
		i++
		line := fileScanner.Text()

		if regex.MatchString(line) {
			fmt.Printf("* ")
			e := strings.Fields(line)
			fmt.Print(e)
		}
		fmt.Println(i, line)

	}
}
