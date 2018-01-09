package main

import (
    "fmt"
)
var ha [20]int
func h_1(x int) int {
    return x%20
}

func h_2(x int) int {
    x += ^(x << 15) + 1
    x ^= (x >> 10)
    x += ^(x << 3) + 1
    x ^= (x >> 6)
    x += ^(x << 11) + 1
    x ^= (x >> 16)
    return x%20
}
func set(x int){
    ha[h_1(x)] = 1
    ha[h_2(x)] = 1
    fmt.Println("set ",x, ":",h_1(x), h_2(x))
}

func isSet(x int) bool {
    h1 := h_1(x)
    h2 := h_2(x)
    fmt.Println("isSet ", x, ":", h1, h2)

    return ha[h1]==1 && ha[h2]==1
}
func main() {
    fmt.Println(ha)
    set(5)
    fmt.Println(ha)
    set(7)
    fmt.Println(ha)
    set(134)
    fmt.Println(ha)
    help := `
    ***************************
    false一定正确,true存在一定概率错误
    ***************************
    `
    fmt.Println(help)
    fmt.Println(isSet(5))
    fmt.Println(isSet(11))
    fmt.Println(isSet(14), "一定概率判断失误")
}
