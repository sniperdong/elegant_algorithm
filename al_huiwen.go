package main

import(
    "fmt"
)
/*
给定一字符串，能否通过增加一个字符成为回文
*/
func main() {
    //str := []byte("abceca")
    //str := []byte("abcecb")
    //str := []byte("abcecca")
    //str := []byte("abcca")
    //str := []byte("abccba")
    str := []byte("abcicba")

    ret := true
    firstWrong := false
    i := 0
    for j := len(str)-1;j>i;j-- {
        if str[i] != str[j] {
            if firstWrong == false {
                    if str[i+1] == str[j] {
                        i++
                        firstWrong = true
                    } else if str[i] == str[j-1] {
                        j--
                        firstWrong = false
                    } else {
                        ret = false
                        break
                    }
            } else {
                ret = false
                break
            }
        }
        i++
    }
    
    fmt.Println(ret)
}
