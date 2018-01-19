package main

import (
    "fmt"
)

func main(){
    n:=25
    m:=0
    for {
        if n == 0 {
            break
        }
        n=n/5
        m+=n
    }

    fmt.Println(m)
}
