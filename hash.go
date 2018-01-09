package main

import (
    "fmt"
)

func h_1(str string) int {
    bytes:= []byte(str)

    hash := 5381
    for i:=len(bytes)-1; i>=0; i--{
        hash = ((hash << 5) + hash) + int(bytes[i])
    }
    return hash & 1<<31
}

func h_2(x int) int {
    x += ^(x << 15) + 1
    x ^= (x >> 10)
    x += ^(x << 3) + 1
    x ^= (x >> 6)
    x += ^(x << 11) + 1
    x ^= (x >> 16)
    return x
}

func main() {
    fmt.Println(h_1("hello world"))
    fmt.Println(h_2(12341))
}
