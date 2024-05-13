package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

type position int

type CalibrationValue struct {
	first_digit, last_digit int
	sum                     int
}

func (c *CalibrationValue) sumValues(reset bool) {
	c.sum = c.sum + (c.first_digit + c.last_digit)

	if reset {
		c.first_digit, c.last_digit = 0, 0
	}
}

var (
	ErrorArgsLen            = "Must be 1 single arg: The file name"
	ErrorOpeningFile        = "Error opening file"
	ErrorStartReadingBuffer = "Error start reading buffer"
	ErrorReadingBuffer      = "Error reading buffer"
)

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

	// Open file
	buff := bufio.NewScanner(file)
	if buff.Err() != nil {
		log.Fatalln(ErrorStartReadingBuffer)
	}

	c := &CalibrationValue{}
	// Read line by line
	for buff.Scan() {
		if buff.Err() != nil {
			log.Fatalln(ErrorReadingBuffer)
		}

		line := buff.Text()
		max_index := len(line) - 1

		for i, v := range line {
			// if all set; break
			if c.first_digit != 0 && c.last_digit != 0 {
				break
			}

			// Read from start
			first_digit := rune(v)
			if unicode.IsDigit(first_digit) && c.first_digit == 0 {
				// Any digit rune subtracted from rune 0 results an integer
				c.first_digit = int(first_digit-'0') * 10
			}

			// Read from end
			last_digit := rune(line[max_index-i])
			if unicode.IsDigit(last_digit) && c.last_digit == 0 {
				// Any digit rune subtracted from rune 0 results an integer
				c.last_digit = int(last_digit - '0')
			}
		}

		c.sumValues(true)
	}
	fmt.Printf("Value: %v\n", c.sum)
}
