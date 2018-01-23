package main

import (
    "fmt"
)

/*
判断给定数值是否是2的n次方，要求时间复杂度O(1)
*/
func main(){
    i:=32
    fmt.Println(i,((i-1)&i)==0)
    i=36
    fmt.Println(i,((i-1)&i)==0)
}
