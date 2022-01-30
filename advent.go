package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func problem2() {
	file, err := os.Open("input2.txt")
	check(err)
	defer file.Close()

	aim, depth, horizontal := 0, 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		temp, err := strconv.Atoi(split[1])
		check(err)
		if split[0] == "up" {
			aim -= temp
		} else if split[0] == "down" {
			aim += temp
		} else if split[0] == "forward" {
			horizontal += temp
			depth += aim * temp
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(depth * horizontal)
}

func problem3() {
	file, err := os.Open("input3.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var inputArr []string

	for scanner.Scan() {
		inputArr = append(inputArr, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	oxygenArr := make([]string, 1000)
	co2Arr := make([]string, 1000)
	copy(oxygenArr, inputArr)
	copy(co2Arr, inputArr)

	index := 0
	for len(oxygenArr) > 1 {
		countOnes := 0
		for _, inputStr := range oxygenArr {
			if inputStr[index:index+1] == "1" {
				countOnes += 1
			}
		}
		var tempArr []string
		var checkVal string
		fmt.Println(countOnes)
		if countOnes >= len(oxygenArr)/2 {
			checkVal = "1"
		} else {
			checkVal = "0"
		}
		for _, inputStr := range oxygenArr {
			if inputStr[index:index+1] == checkVal {
				tempArr = append(tempArr, inputStr)
			}
		}
		oxygenArr = tempArr
		index += 1
	}

	index = 0
	for len(co2Arr) > 1 {
		countOnes := 0
		for _, inputStr := range co2Arr {
			if inputStr[index:index+1] == "1" {
				countOnes += 1
			}
		}
		var tempArr []string
		var checkVal string
		if countOnes >= len(co2Arr)/2 {
			checkVal = "0"
		} else {
			checkVal = "1"
		}
		for _, inputStr := range co2Arr {
			if inputStr[index:index+1] == checkVal {
				tempArr = append(tempArr, inputStr)
			}
		}
		co2Arr = tempArr
		index += 1
	}

	fmt.Println(oxygenArr)
	fmt.Println(co2Arr)

	oxygen, err := strconv.ParseInt(oxygenArr[0], 2, 64)
	check(err)

	co2, err := strconv.ParseInt(co2Arr[0], 2, 64)
	check(err)

	fmt.Println(co2 * oxygen)
}

func checkForBingo(bingoSheet [][]string, indexRow int, indexCol int) [][]string {
	// check the horizontal/vertical bingo for current sheet
	boolHorz := true
	boolVert := true
	for i := 0; i < 5; i += 1 {
		boolHorz = boolHorz && (bingoSheet[indexRow][i] == "-1")
		boolVert = boolVert && (bingoSheet[i][indexCol] == "-1")
	}
	if boolHorz || boolVert {
		return bingoSheet
	} else {
		return nil
	}
}

func problem4() {
	file, err := os.Open("input4.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var inputArr []string

	for scanner.Scan() {
		inputArr = append(inputArr, scanner.Text())
	}

	bingoInput := strings.Split(inputArr[0], ",")

	var bingoCharts [][][]string

	for i := 2; i < len(inputArr); i += 6 {
		var appendArr [][]string
		for j := i; j < i+5 && j < len(inputArr); j += 1 {
			appendArr = append(appendArr, strings.Split(strings.TrimSpace(inputArr[j]), " "))
		}
		bingoCharts = append(bingoCharts, appendArr)
	}

	inputIndex := 0
	for inputIndex < len(inputArr) {
		for indexSheet, bingoSheet := range bingoCharts {
			if bingoSheet[0][0] == "W" {
				continue
			}
			for indexRow, bingoRow := range bingoSheet {
				for indexCol, value := range bingoRow {
					if bingoInput[inputIndex] == value {
						prevVal := value
						bingoCharts[indexSheet][indexRow][indexCol] = "-1"
						if checkForBingo(bingoSheet, indexRow, indexCol) != nil {
							sum := 0
							for _, thingRow := range bingoSheet {
								for _, thingVal := range thingRow {
									if thingVal != "-1" {
										temp, err := strconv.Atoi(thingVal)
										check(err)
										sum += temp
									}
								}
							}
							temp, err := strconv.Atoi(prevVal)
							check(err)
							fmt.Println(temp, sum)
							bingoSheet[0][0] = "W"
						}
					}

				}
			}
		}
		inputIndex += 1
	}
}

func readFile(input string) *os.File {
	file, err := os.Open(input)
	check(err)
	return file
}

func problem5() {
	file := readFile("input5.txt")

	scanner := bufio.NewScanner(file)

	const length = 1000

	var inputMatrix [length][length]int

	for scanner.Scan() {
		unformattedPoints := strings.Split(scanner.Text(), " -> ")
		tempOne := strings.Split(unformattedPoints[0], ",")
		tempTwo := strings.Split(unformattedPoints[1], ",")

		tempOneX, err := strconv.Atoi(tempOne[0])
		check(err)
		tempOneY, err := strconv.Atoi(tempOne[1])
		check(err)

		tempTwoX, err := strconv.Atoi(tempTwo[0])
		check(err)
		tempTwoY, err := strconv.Atoi(tempTwo[1])
		check(err)

		if tempOneX == tempTwoX && tempOneY == tempTwoY {
			inputMatrix[tempOneX][tempOneY] += 1
		} else {
			if tempOneX == tempTwoX {
				var greater int
				var lesser int
				if tempOneY > tempTwoY {
					greater = tempOneY
					lesser = tempTwoY
				} else {
					greater = tempTwoY
					lesser = tempOneY
				}

				for i := lesser; i <= greater; i++ {
					inputMatrix[tempOneX][i] += 1
				}
			} else if tempOneY == tempTwoY {
				var greater int
				var lesser int
				if tempOneX > tempTwoX {
					greater = tempOneX
					lesser = tempTwoX
				} else {
					greater = tempTwoX
					lesser = tempOneX
				}
				for i := lesser; i <= greater; i++ {
					inputMatrix[i][tempOneY] += 1
				}
			} else if tempOneX-tempTwoX == tempOneY-tempTwoY || -(tempOneX-tempTwoX) == tempOneY-tempTwoY {
				var greaterX int // second point
				var greaterY int
				var lesserX int // first point
				var lesserY int
				if tempOneX > tempTwoX {
					greaterX = tempOneX
					greaterY = tempOneY
					lesserX = tempTwoX
					lesserY = tempTwoY
				} else {
					greaterX = tempTwoX
					greaterY = tempTwoY
					lesserX = tempOneX
					lesserY = tempOneY
				}
				if lesserY < greaterY {
					for i, j := lesserX, lesserY; i <= greaterX && j <= greaterY; i, j = i+1, j+1 {
						inputMatrix[i][j] += 1
					}
				} else {
					for i, j := lesserX, lesserY; i <= greaterX && j >= greaterY; i, j = i+1, j-1 {
						inputMatrix[i][j] += 1
					}
				}
			}
		}
	}
	var count int
	for i := 0; i < length; i += 1 {
		for j := 0; j < length; j += 1 {
			if inputMatrix[i][j] >= 2 {
				count += 1
			}
		}
	}
	fmt.Println(count)
}

func nextCycle(fishy *int) int {
	if *fishy == 0 {
		*fishy = 6
		return 1
	} else {
		*fishy -= 1
		return 0
	}
}

func problem6() {
	const totalCycles = 256

	file := readFile("text6.txt")
	scanner := bufio.NewScanner(file)
	var rawFishs []string
	for scanner.Scan() {
		rawFishs = strings.Split(scanner.Text(), ",")
	}

	var fishies []int
	for _, strFish := range rawFishs {
		intFish, err := strconv.Atoi(strFish)
		check(err)
		fishies = append(fishies, intFish)
	}

	for i := 0; i < totalCycles; i += 1 {
		countNewFishies := 0
		for fishIndex := 0; fishIndex < len(fishies); fishIndex += 1 {
			addFish := nextCycle(&fishies[fishIndex])
			countNewFishies += addFish
		}

		for j := 0; j < countNewFishies; j++ {
			fishies = append(fishies, 8)
		}
	}

	fmt.Println(len(fishies))
}

func problem6part2() {
	const totalCycles = 256

	file := readFile("input6.txt")
	scanner := bufio.NewScanner(file)
	var rawFishs []string
	for scanner.Scan() {
		rawFishs = strings.Split(scanner.Text(), ",")
	}

	fishDict := make(map[int]int)
	fishDict[0] = 0
	fishDict[1] = 0
	fishDict[2] = 0
	fishDict[3] = 0
	fishDict[4] = 0
	fishDict[5] = 0
	fishDict[6] = 0
	fishDict[7] = 0
	fishDict[8] = 0
	for _, strFish := range rawFishs {
		intFish, err := strconv.Atoi(strFish)
		check(err)

		fishDict[intFish] += 1
	}

	for i := 0; i < totalCycles; i += 1 {
		temp := fishDict[0]
		fishDict[0] = fishDict[1]
		fishDict[1] = fishDict[2]
		fishDict[2] = fishDict[3]
		fishDict[3] = fishDict[4]
		fishDict[4] = fishDict[5]
		fishDict[5] = fishDict[6]
		fishDict[6] = fishDict[7]
		fishDict[7] = fishDict[8]
		fishDict[8] = temp
		fishDict[6] += temp
	}
	totalCount := 0
	for i := 0; i < 9; i += 1 {
		totalCount += fishDict[i]
	}

	fmt.Println(totalCount)
}

func sumUpTo(num int) int {
	retVal := 0
	for i := 1; i <= num; i += 1 {
		retVal += i
	}
	return retVal
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func problem7() {
	file := readFile("input7.txt")
	scanner := bufio.NewScanner(file)
	var strCrabbies []string
	for scanner.Scan() {
		strCrabbies = strings.Split(scanner.Text(), ",")
	}

	var crabbies []int
	for _, crab := range strCrabbies {
		newCrab, err := strconv.Atoi(crab)
		check(err)
		crabbies = append(crabbies, newCrab)
	}

	minFuel := math.MaxInt32
	crabMin, crabMax := MinMax(crabbies)
	for i := crabMin; i < crabMax; i += 1 {
		currentFuel := 0
		for _, crabs := range crabbies {
			currentFuel += sumUpTo(int(math.Abs(float64(i) - float64(crabs))))
		}
		minFuel = int(math.Min(float64(minFuel), float64(currentFuel)))
	}

	fmt.Println(minFuel)
}

func main() {
	problem7()
}
