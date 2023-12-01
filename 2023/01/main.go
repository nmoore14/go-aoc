package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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
