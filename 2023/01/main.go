package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// NOTE: part 1 => 54450
// BUG: part 2 => 53789 is too low
// BUG: part 2 => 60750 is too high

// Part 1
func getTwoDigit(digits string) int {
	digits_len := len(digits)
	digit_buf := []byte(digits)
	temp_buf := []byte{}

	if digits_len > 1 {
		temp_buf = append(temp_buf, digit_buf[0])
		temp_buf = append(temp_buf, digit_buf[digits_len-1])
	} else {
		temp_buf = append(temp_buf, digit_buf[0])
		temp_buf = append(temp_buf, digit_buf[0])
	}

	two_digit, error := strconv.Atoi(string(temp_buf))

	if error != nil {
		log.Fatal(error)
	}

	return two_digit
}

// Part 1
func getDigits(position string) string {
	pos_digits_arr := []string{}

	reg := regexp.MustCompile(`\d+`)

	matches := reg.FindAllString(position, -1)

	for _, match := range matches {
		pos_digits_arr = append(pos_digits_arr, match)
	}

	return strings.Join(pos_digits_arr, "")
}

// Part 2
var NUM_STRINGS = [...]string{"ZERO", "ONE", "TWO", "THREE", "FOUR", "FIVE", "SIX", "SEVEN", "EIGHT", "NINE"}
var STARTING_CHARS = "OTFSEN"

func checkDigitStringValue(digit_string []string) int {
	var str string = strings.Join(digit_string, "")

	for i, NUM_STRING := range NUM_STRINGS {
		if str == NUM_STRING {
			return i
		}
	}

	return 0
}

// TODO: Figure out if I need to just need to check no matter what as long as I find a starting char or not
func getFirstDigit(position string) string {
	var digit string
	var digit_string []string

	for i, c := range position {
		if unicode.IsDigit(c) {
			digit = string(c)
			break
		}

		if strings.Contains(STARTING_CHARS, string(c)) {
			digit_string = append(digit_string, string(c))
			for j := i + 1; j < (i + 5); j++ {
				if len(digit_string) < 3 {
					digit_string = append(digit_string, digit_string[j])
				}
			}
		}

		digit_string = nil
	}

	return digit
}

// TODO: Build this out like above but reverse
func getLastDigit(position string) string {}

func getDigitsFromString(position string) int {
	position = strings.ToUpper(position)
	position_norm := position
	position_rev := position

	for i := 0; i < 10; i++ {
		position_norm = strings.Replace(position_norm, NUM_STRINGS[i], strconv.Itoa(i), -1)
	}

	for j := 9; j >= 0; j-- {
		position_rev = strings.Replace(position_rev, NUM_STRINGS[j], strconv.Itoa(j), -1)
	}

	norm_digits := getDigits(position_norm)
	rev_digits := getDigits(position_rev)

	fmt.Println(norm_digits)
	fmt.Println(rev_digits)

	norm := getTwoDigit(norm_digits)
	rev := getTwoDigit(rev_digits)

	if norm > rev {
		return norm
	}

	return rev
}

func main() {
	f, error := os.Open("input-user.txt")
	var pos_total int = 0

	var isPart2 bool = true

	if error != nil {
		log.Fatal(error)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if !isPart2 {
			pos_total += getTwoDigit(getDigits(scanner.Text()))
		} else {
			pos_total += getDigitsFromString(scanner.Text())
		}
	}

	if error := scanner.Err(); error != nil {
		log.Fatal(error)
	}

	fmt.Println(pos_total)
}
