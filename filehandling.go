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
	if err != nil {
		panic(err)
	}
	// data := "harold wilson (cap, 2) : ' I’m a optimist ,but a optimist who carries a raincoat . '"
	result := strings.Fields(string(data))

	// runs range loop to modify result
	for i, v := range result {
		// replaces the word before with its decimal version
		if v == "(hex)" {
			j, _ := strconv.ParseInt(result[i-1], 16, 64)
			result[i-1] = fmt.Sprint(j)
			result[i] = ""

		}
		// replaces the word before with its decimal version
		if v == "(bin)" {
			j, _ := strconv.ParseInt(result[i-1], 2, 64)
			result[i-1] = fmt.Sprint(j)
			result[i] = ""

		}

		// converts previous word upper
		if v == "(up)" {
			result[i-1] = strings.ToUpper(result[i-1])
			result[i] = ""
		}
		// num of previous words upper
		if v == "(up," {
			result[i-1] = strings.ToUpper(result[i-1])

			le := len(result[i+1])
			numb := result[i+1][:le-1]

			nu, err := strconv.Atoi(numb)
			if err != nil {
				panic(err)
			}

			for j := 1; j <= nu; j++ {
				result[i-j] = strings.ToUpper(result[i-j])
			}
			result[i], result[i+1] = "", ""
		}

		// converts the word before to lowercase
		if v == "(low)" {
			result[i-1] = strings.ToLower(result[i-1])
			result[i] = ""
		}
		// converts the number of words before to lowercase
		if v == "(low," {
			result[i-1] = strings.ToLower(result[i-1])

			le := len(result[i+1])
			numb := result[i+1][:le-1]

			nu, err := strconv.Atoi(numb)
			if err != nil {
				panic(err)
			}

			for j := 1; j <= nu; j++ {
				result[i-j] = strings.ToLower(result[i-j])
			}
			result[i], result[i+1] = "", ""
		}
		if v == "(cap)" {
			result[i-1] = capitalise(result[i-1])
			result[i] = ""
		}
		// capitalises the number of words before
		if v == "(cap," {
			result[i-1] = capitalise(result[i-1])

			le := len(result[i+1])
			numb := result[i+1][:le-1]
			nu, err := strconv.Atoi(numb) // changed to integer
			if err != nil {
				panic(err)
			}

			for j := 1; j <= nu; j++ {
				result[i-j] = capitalise(result[i-j])
			}
			result[i], result[i+1] = "", ""
		}
		strr := ""
		for _, v := range result {
			strr = strr + v + " "
		}
		result1 := strings.Fields(strr)

		// converts a to an
		if v == "a" && first_rune(result[i+1]) == "a" || v == "a" && first_rune(result[i+1]) == "e" || v == "a" && first_rune(result[i+1]) == "i" || v == "a" && first_rune(result[i+1]) == "o" || v == "a" && first_rune(result[i+1]) == "u" || v == "a" && first_rune(result[i+1]) == "h" {
			result[i] = "an"
		}
		str := ""
		for _, v := range result1 {
			str += v + " "
		}

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

		output := []byte(quotes(word))
		OurData := os.Args[2]
		words := os.WriteFile(OurData, output, 0644)
		if words != nil {
			panic(words)
		}
	}
	// fmt.Println(word)
}

// quits if there's an error

// gets the first rune of a string
func first_rune(s string) string {
	a := []rune(s)
	return string(a[0])
}

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
