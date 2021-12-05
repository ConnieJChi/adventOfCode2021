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

    var countOnes[12]int

    for scanner.Scan() {
        for index, character := range scanner.Text() {
            if character == '1' {
                countOnes[index] += 1
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    strGamma := ""
    strEpsilon := ""
    for _, count := range countOnes {
        if count > 500 {
            strGamma += "1"
            strEpsilon += "0"
        } else {
            strGamma += "0"
            strEpsilon += "1"
        }
    }
    gamma, err := strconv.ParseInt(strGamma, 2, 64)
    check(err)

    epsilon, err := strconv.ParseInt(strEpsilon, 2, 64)
    check(err)

    fmt.Println(gamma * epsilon)
}

func main() {
    problem3()
}
