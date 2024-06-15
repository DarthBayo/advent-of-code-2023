package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

var (
	ErrorArgsLen            = "It must have a single arg: The file name"
	ErrorOpeningFile        = "Error opening the file"
	ErrorStartReadingBuffer = "Error when starting to read the buffer"
	ErrorReadingBuffer      = "Error reading the buffer"

	Empty int = -1
)

type calibrationValue struct {
	firstDigit, lastDigit, totalSum int
}

func NewCalibrationValue() *calibrationValue {
	return &calibrationValue{firstDigit: Empty, lastDigit: Empty, totalSum: 0}
}

func (c calibrationValue) IsFullfilled() bool {
	return c.firstDigit != Empty && c.lastDigit != Empty
}

func (c calibrationValue) IsFirstDigitSet() bool {
	return c.firstDigit != Empty
}

func (c calibrationValue) IsLastDigitSet() bool {
	return c.lastDigit != Empty
}

func (c *calibrationValue) SetFirstDigit(digit rune) {
	// Any digit rune subtracted from rune 0 results an integer
	c.firstDigit = int(digit-'0') * 10
}

func (c *calibrationValue) SetLastDigit(digit rune) {
	// Any digit rune subtracted from rune 0 results an integer
	c.lastDigit = int(digit - '0')
}

func (c *calibrationValue) SumValues(reset bool) {
	c.totalSum = c.totalSum + (c.firstDigit + c.lastDigit)

	if reset {
		c.firstDigit, c.lastDigit = Empty, Empty
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln(ErrorArgsLen)
	}

	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(ErrorOpeningFile)
	}
	defer file.Close()

	simple(file)
}

func simple(file *os.File) {
	// Open file
	buff := bufio.NewScanner(file)
	if buff.Err() != nil {
		log.Fatalln(ErrorStartReadingBuffer)
	}

	c := NewCalibrationValue()
	// Read line by line
	for buff.Scan() {
		if buff.Err() != nil {
			log.Fatalln(ErrorReadingBuffer)
		}

		line := buff.Text()
		max_index := len(line) - 1

		for i, v := range line {
			// if all set; break
			if c.IsFullfilled() {
				break
			}

			// Reads from the beginning to the end of the string
			first_digit := rune(v)
			if unicode.IsDigit(first_digit) && !c.IsFirstDigitSet() {
				c.SetFirstDigit(first_digit)
			}

			// Reads from the end to the beginning of the string
			last_digit := rune(line[max_index-i])
			if unicode.IsDigit(last_digit) && !c.IsLastDigitSet() {
				c.SetLastDigit(last_digit)
			}
		}

		c.SumValues(true)
	}
	fmt.Printf("Value: %v\n", c.totalSum)
}
