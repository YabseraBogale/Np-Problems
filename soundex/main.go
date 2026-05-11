package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Soundex(str string) string {

	var result string
	result += string(str[0])
	for _, i := range strings.ToLower(str) {
		if i == 0 {
			continue
		} else if string(i) == "b" || string(i) == "f" || string(i) == "p" || string(i) == "v" {
			result += "1"
		} else if string(i) == "c" || string(i) == "g" || string(i) == "j" || string(i) == "k" || string(i) == "q" || string(i) == "s" || string(i) == "x" || string(i) == "z" {
			result += "2"
		} else if string(i) == "d" || string(i) == "t" {
			result += "3"
		} else if string(i) == "l" {
			result += "4"
		} else if string(i) == "m" || string(i) == "n" {
			result += "5"

		} else if string(i) == "r" {
			result += "6"
		} else if string(i) == "a" || string(i) == "e" || string(i) == "i" || string(i) == "o" || string(i) == "u" || string(i) == "h" || string(i) == "w" || string(i) == "y" {
			result += "0"
		}
		if len(result) == 4 {
			return result
		}

	}

	for i := 0; len(result) < 4; i++ {
		result += "0"
	}

	return result
}

func main() {
	file, err := os.Open("last-names.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := make(map[string][]string)

	for scanner.Scan() {
		if _, ok := result[Soundex(scanner.Text())]; ok {
			result[Soundex(scanner.Text())] = append(result[Soundex(scanner.Text())], scanner.Text())
		} else {
			result[Soundex(scanner.Text())] = []string{scanner.Text()}
		}

	}
	for i, j := range result {
		fmt.Println(i, j)
	}
}
