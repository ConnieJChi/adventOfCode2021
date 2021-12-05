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

func main() {
    problem3()
}
