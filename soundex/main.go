package main

import "fmt"

func Soundex(str string) string {

	var result string
	result += string(str[0])
	for i, _ := range str {
		if i == 0 {
			continue
		} else if string(str[i]) == "b" || string(str[i]) == "f" || string(str[i]) == "p" || string(str[i]) == "v" {
			result += "1"
		} else if string(str[i]) == "c" || string(str[i]) == "g" || string(str[i]) == "j" || string(str[i]) == "k" || string(str[i]) == "q" || string(str[i]) == "s" || string(str[i]) == "x" || string(str[i]) == "z" {
			result += "2"
		} else if string(str[i]) == "d" || string(str[i]) == "t" {
			result += "3"
		} else if string(str[i]) == "l" {
			result += "4"
		} else if string(str[i]) == "m" || string(str[i]) == "n" {
			result += "5"
		} else if string(str[i]) == "r" {
			result += "6"
		} else if string(str[i]) == "a" || string(str[i]) == "e" || string(str[i]) == "i" || string(str[i]) == "o" || string(str[i]) == "u" || string(str[i]) == "h" || string(str[i]) == "w" || string(str[i]) == "y" {
			result += "0"
		}
		if len(result) == 4 {
			return result
		}
	}
	return result
}

func main() {
	fmt.Println(Soundex("Hello"))
}
