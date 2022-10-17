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
	input := string(data)
	result := split_white_spaces(input)

	// runs range loop to modify result
	for i, v := range result {
		// replaces the word before with its decimal version
		if compare(v, "(hex)") == 0 {
			j, _ := strconv.ParseInt(result[i-1], 16, 64)
			result[i-1] = fmt.Sprint(j)
		}
		// replaces the word before with its decimal version
		if compare(v, "(bin)") == 0 {
			j, _ := strconv.ParseInt(result[i-1], 2, 64)
			result[i-1] = fmt.Sprint(j)
		}
		// converts the word before to lowercase
		if compare(v, "(low)") == 0 {
			result[i-1] = to_lower(result[i-1])
		}
		// converts the number of words before to lowercase
		if compare(v, "(low,") == 0 {
			result[i-1] = to_lower(result[i-1])
			fmt.Println(result)

			le := len(result[i+1])
			numb := result[i+1][:le-1]
			nu, err := strconv.Atoi(numb)
			check(err)
			fmt.Println(le)
			fmt.Println(numb)
			fmt.Println(nu)

			for j := 1; j <= nu; j++ {
				result[i-j] = to_lower(result[i-j])
			}
		}
		// converts the word before to uppercase
		if compare(v, "(up)") == 0 {
			result[i-1] = to_upper(result[i-1])
		}
		// converts the number of words before to uppercase
		if compare(v, "(up,") == 0 {
			result[i-1] = to_upper(result[i-1])

			le := len(result[i+1])
			numb := result[i+1][:le-1]
			nu, err := strconv.Atoi(numb)
			check(err)

			for j := 1; j <= nu; j++ {
				result[i-j] = to_upper(result[i-j])
			}
		}
		// capitalises the word before
		if compare(v, "(cap)") == 0 {
			result[i-1] = capitalise(result[i-1])
		}
		// capitalises the number of words before  //maybe remove 70 to 73
		if compare(v, "(cap,") == 0 {
			result[i-1] = capitalise(result[i-1])

			le := len(result[i+1])        // le=2
			numb := result[i+1][:le-1]    // remove parenthesis
			nu, err := strconv.Atoi(numb) // changed ascii to integer
			check(err)

			for j := 1; j <= nu; j++ {
				result[i-j] = capitalise(result[i-j])
			}
		}
		// converts 'a' into 'an' when the next word begins with a vowel or 'h'.//maybe shorten with loop
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

// converts word to uppercase
func to_upper(s string) string {
	sentence := []rune(s)
	for i := 0; i < len(sentence); i++ {

		currentLetter := sentence[i]

		if currentLetter >= 'a' && currentLetter <= 'z' {
			sentence[i] = sentence[i] - 32
		}
	}
	return string(sentence)
}

// converts word to lowercase
func to_lower(s string) string {
	sentence := []rune(s)
	for i := 0; i < len(sentence); i++ {

		currentLetter := sentence[i]

		if currentLetter >= 'A' && currentLetter <= 'Z' {
			sentence[i] = sentence[i] + 32
		}
	}
	return string(sentence)
}

// gets the first rune of a string
func first_rune(s string) string {
	a := []rune(s)
	return string(a[0])
}

// capitalises the first character of a word
func capitalise(s string) string {
	runes := []rune(s)

	strlen := 0
	for i := range runes {
		strlen = i + 1
	}

	for i := 0; i < strlen; i++ {
		if i != 0 && (!((runes[i-1] >= 'a' && runes[i-1] <= 'z') || (runes[i-1] >= 'A' && runes[i-1] <= 'Z'))) {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] = rune(runes[i] - 32)
			}
		} else if i == 0 {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] = rune(runes[i] - 32)
			}
		} else {
			if runes[i] >= 'A' && runes[i] <= 'Z' {
				runes[i] = rune(runes[i] + 32)
			}
		}
	}
	return string(runes)
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
