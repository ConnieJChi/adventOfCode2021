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

    x, y := 0, 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        split := strings.Split(scanner.Text(), " ")
        temp, err := strconv.Atoi(split[1])
        check(err)
        if split[0] == "up" {  
            y -= temp
        } else if split[0] == "down" {
            y += temp
        } else if split[0] == "forward" {
            x += temp
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    fmt.Println(x, y, x * y)
}

func main() {
	fmt.Println("Welcome to the playground!")
    problem2()
}
