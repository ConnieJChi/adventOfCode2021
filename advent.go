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
	file, err := os.Open("input.txt")
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

func main() {
	fmt.Println("Welcome to the playground!")
    problem2()
}
