package main

import (
    "fmt"
    "math/rand"
    "sort"
    //"time"
)

/*
n个随机任意数据，从中快速找出3个和是0的3元组
*/

const(
    Max = 100
)

type splitInt []int

func (s splitInt) Len() int { return len(s)}
func (s splitInt) Less(i,j int) bool { return s[i]<s[j]}
func (s splitInt) Swap(i, j int) {s[i],s[j] = s[j], s[i]}

func main() {
    iArr := make([]int, Max)
    for i:=0;i<Max;i++{
        //rand.Seed(time.Now().UnixNano())
        iArr[i] = rand.Intn(100) - 50
        //time.Sleep(7)
    }
    fmt.Println(iArr)
    sort.Sort(splitInt(iArr))
    fmt.Println(iArr)
    
    for i:=0;i<Max && iArr[i]<0;i++{
        for j:=Max-2;j>i;j-- {
            if (iArr[i]+iArr[Max-1-i]+iArr[j]) == 0 {
                fmt.Println(iArr[i],iArr[Max-1-i],iArr[j])
                continue
            } 
        }
    }
}
