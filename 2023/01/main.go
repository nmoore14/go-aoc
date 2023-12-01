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

func getDigits(position string) string {
	pos_digits_arr := []string{}

	reg := regexp.MustCompile(`\d+`)

	matches := reg.FindAllString(position, -1)

	for _, match := range matches {
		pos_digits_arr = append(pos_digits_arr, match)
	}

	return strings.Join(pos_digits_arr, "")
}

func main() {
	f, error := os.Open("input-user.txt")
	pos_total := 0

	if error != nil {
		log.Fatal(error)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		pos_total += getTwoDigit(getDigits(scanner.Text()))
	}

	if error := scanner.Err(); error != nil {
		log.Fatal(error)
	}

	fmt.Println(pos_total)
}
