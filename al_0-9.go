package main

import (
    "fmt"
)

/*
一组不重复数据，范围0-9，只有9个元素，计算出缺少的数
*/

func main(){
    a:=[]int{1,2,7,8,3,4,5,9,0}
    n:=(0+9)*10/2
    m:=0
    for i :=0;i<len(a);i++ {
        fmt.Println(a[i])
        m+=a[i]
    }
    fmt.Println("缺少0-9中的数字:",n,m,n-m)
}
