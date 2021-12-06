package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strings"
    "strconv"
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
        if countOnes >= len(oxygenArr) / 2 {
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
        if countOnes >= len(co2Arr) / 2 {
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

func checkForBingo(bingoSheet [][] string, indexRow int, indexCol int) [][] string {
    // check the horizontal/vertical bingo for current sheet
    boolHorz := true
    boolVert := true
    for i := 0; i < 5; i+= 1 {
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

    var bingoCharts [][][] string

    for i := 2; i < len(inputArr); i += 6 {
        var appendArr [][]string
        for j := i; j < i + 5 && j < len(inputArr); j += 1 {
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
    // fmt.Println(bingoCharts)
}

func main() {
    problem4()
}
