package main

import (
"fmt"
"os"
"runtime/debug"
"log"
)

func ErrorHandler() {
    if er := recover(); er != nil {
        log.Println(er, string(debug.Stack()))
    }
}

func Divide(dividend float32, divisor float32) (quotient float32) {
    defer ErrorHandler()
    if divisor == 0 {
        panic("Can't Divide by 0")
    }
    quotient = dividend / divisor
    return
}

func main() {
    f, err := os.OpenFile("ErrLogs.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        log.Println(err)
    }
    log.SetOutput(f)
    log.Println("Starting Program")
    no := Divide(20, 0)
    fmt.Println(no)

    no = Divide(20, 4)
    fmt.Println(no)

    f.Close()
}