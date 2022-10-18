package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	GoReloaded()
}

// reads file and saves content to 'data' var
func GoReloaded() {
	data, err := os.ReadFile(os.Args[1])
	check(err)
	// data := "harold wilson (cap, 2) : ' I’m a optimist ,but a optimist who carries a raincoat . '"
	result := strings.Fields(string(data))

	// runs range loop to modify result
	for i, v := range result {
		// replaces the word before with its decimal version
		if v == "(hex)" {
			j, _ := strconv.ParseInt(result[i-1], 16, 64)
			result[i-1] = fmt.Sprint(j)

		}
		// replaces the word before with its decimal version
		if compare(v, "(bin)") == 0 {
			j, _ := strconv.ParseInt(result[i-1], 2, 64)
			result[i-1] = string(rune(j))

		}
		// converts the word before to lowercase
		if v == "(low)" {
			result[i-1] = strings.ToLower(result[i-1])
		}
		// converts the number of words before to lowercase
		if v == "(low," {
			result[i-1] = strings.ToLower(result[i-1])

			le := len(result[i+1])
			numb := result[i+1][:le-1]
			nu, err := strconv.Atoi(numb)
			check(err)

			for j := 1; j <= nu; j++ {
				result[i-j] = strings.ToLower(result[i-j])
			}
		}
		// converts the word before to uppercase
		if compare(v, "a") == 0 && first_rune(result[i+1]) == "a" {
			result[i] = "an"
		}
		if compare(v, "a") == 0 && first_rune(result[i+1]) == "e" {
			result[i] = "an"
		}
		if compare(v, "a") == 0 && first_rune(result[i+1]) == "i" {
			result[i] = "an"
		}
		if compare(v, "a") == 0 && first_rune(result[i+1]) == "o" {
			result[i] = "an"
		}
		if compare(v, "a") == 0 && first_rune(result[i+1]) == "u" {
			result[i] = "an"
		}
		if compare(v, "a") == 0 && first_rune(result[i+1]) == "h" {
			result[i] = "an"
		}
	}

	// calls remove_tags() and split_white_spaces() and gets a new result variable
	notagResult := remove_tags(result)
	result2 := split_white_spaces(notagResult)

	str := ""
	for _, word := range result2 {
		str = str + word + " "
	}
	// remove spaces from string
	str = remove_spaces(str)

	word := ""
	for i, char := range str {
		if i == len(str)-1 {
			if char == '.' || char == ',' || char == '!' || char == '?' || char == ';' || char == ':' {
				if str[i-1] == ' ' {
					word = word[:len(word)-1] // end of paragraph avoidance of space after the full stop
					word = word + string(char)
				} else {
					word = word + string(char)
				}
			} else {
				word = word + string(char)
			}
		} else if char == '.' || char == ',' || char == '!' || char == '?' || char == ';' || char == ':' {
			if str[i-1] == ' ' {
				word = word[:len(word)-1] // removes blank space prior to character
				word = word + string(char)
			} else {
				word = word + string(char)
			}
			if str[i+1] != ' ' && str[i+1] != '.' && str[i+1] != ',' && str[i+1] != '!' && str[i+1] != '?' && str[i+1] != ';' && str[i+1] != ':' {
				word = word + " " // adds a space after character
			}
		} else {
			word = word + string(char)
		}
	}
	// calls quotes() to remove additional spaces
	word = quotes(word)
	output := []byte(word)
	OurData := os.Args[2]
	words := os.WriteFile(OurData, output, 0644)
	check(words)
	// fmt.Println(word)
}

// quits if there's an error
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func compare(a, b string) int {
	if a == b {
	} else if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// gets the first rune of a string
func first_rune(s string) string {
	a := []rune(s)
	return string(a[0])
}

// seperate string by spaces and appends to string list
func split_white_spaces(s string) []string {
	return strings.Split(s, " ")

	// var str []string
	// var word string
	// l := len(s) - 1

	// for i, v := range s {
	// 	if i == l {
	// 		word = word + string(v)
	// 		str = append(str, word)
	// 	} else if v == 32 || v == 15 || v == 10 {
	// 		if s[i+1] == ' ' || s[i+1] == '	' || s[i+1] == 10 {
	// 		} else {
	// 			str = append(str, word)
	// 			word = ""
	// 		}
	// 	} else {
	// 		word = word + string(v)
	// 	}
	// }
	// return str
}

func quotes(s string) string {
	str := ""
	var removeSpace bool // default false
	for i, char := range s {
		if char == 39 && s[i-1] == ' ' {
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(char)
				removeSpace = false
			} else {
				str = str + string(char)
				removeSpace = true
			}
		} else if i > 1 && s[i-2] == 39 && s[i-1] == ' ' {
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(char)
			} else {
				str = str + string(char)
			}
		} else {
			str = str + string(char)
		}
	}
	return str
}

func remove_tags(s []string) string {
	str := ""

	for i, tag := range s {
		if tag == "(cap," || tag == "(low," || tag == "(up," {
			s[i] = ""
			s[i+1] = ""
		} else if tag != "(up)" && tag != "(hex)" && tag != "(bin)" && tag != "(cap)" && tag != "(low)" && tag != "" {
			if i == 0 {
				str = str + tag
			} else {
				str = str + " " + tag
			}
		}
	}
	return str
}

func remove_spaces(s string) string {
	len := len(s) - 1
	if s[len-1] == ' ' {
		return remove_spaces(s[:len])
	}
	return s[:len]
}
