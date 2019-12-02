package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const target = 19690720

func add(index int, opList *[]int) {
	inIndex1 := (*opList)[index+1]
	inIndex2 := (*opList)[index+2]
	sum := (*opList)[inIndex1] + (*opList)[inIndex2]
	outIndex := (*opList)[index+3]
	(*opList)[outIndex] = sum
}

func mult(index int, opList *[]int) {
	inIndex1 := (*opList)[index+1]
	inIndex2 := (*opList)[index+2]
	res := (*opList)[inIndex1] * (*opList)[inIndex2]
	outIndex := (*opList)[index+3]
	(*opList)[outIndex] = res
}

func operate(index int, opList *[]int) {
	opCode := (*opList)[index]
	switch opCode {
	case 1:
		add(index, opList)
	case 2:
		mult(index, opList)
	case 99:
		return
	default:
		// log.Fatalln("Invalid input!!")
		(*opList)[0] = -1
		return
	}

	index += 4
	if index < len(*opList) {
		operate(index, opList)
	}
}

func compute(p1 int, p2 int, opList []int) int {
	// fmt.Println("Input: ", p1, p2, opList)
	opList[1] = p1
	opList[2] = p2
	operate(0, &opList)
	return opList[0]
}

func main() {
	file, err := os.Open("./input/input2.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	splitStr := strings.Split(input, ",")
	opList := make([]int, len(splitStr))

	for i, str := range splitStr {
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalln(err)
		}
		opList[i] = num
	}

	// 1202 program alarm, the state it had just before the last computer caught fire
	// opList[1] = 12
	// opList[2] = 2

	searchSpace := 100

	for i := 0; i < searchSpace; i++ {
		for j := 0; j < searchSpace; j++ {
			opListCpy := make([]int, len(opList))
			copy(opListCpy, opList)
			res := compute(i, j, opListCpy)
			if res == target {
				println("Found solution: ", i, j)
			}
		}
	}

	println("Done!")

	// operate(0, &opList)

	// fmt.Println("Output opList: ", opList)
}
