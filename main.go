package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("document.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	err = file.Close()
	if err != nil {
		return
	}

	totalSum := 0

	for _, eachline := range txtlines {
		var firstDigit string = "0"
		var lastDigit string = "0"
		for _, char := range eachline {
			if isNumber(string(char)) {
				if firstDigit == "0" {
					firstDigit = string(char)
					lastDigit = string(char)
				} else {
					lastDigit = string(char)
				}
			}
		}

		number := firstDigit + lastDigit

		finalNumber, _ := strconv.Atoi(number)

		totalSum += finalNumber
	}

	fmt.Println(totalSum)
}

func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
