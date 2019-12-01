package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func getRequiredFuel(weight int) int {
	fuel := int(math.Floor(float64(weight)/3.0)) - 2
	if fuel > 2 {
		addedFuel := getRequiredFuel(fuel)
		if addedFuel > 0 {
			fuel += addedFuel
		}
	}
	return fuel
}

func day1() {
	file, err := os.Open("./input/input1.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		weight, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}
		sum += getRequiredFuel(weight)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sum: ", sum)
}

func main() {
	fmt.Println("Hello, playground")
	day1()
}
