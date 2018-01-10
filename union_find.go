package main

import (
    "fmt"
)

var un [20]int

func init(){
    un = [20]int{0,0,0,0,3,5,5,7,7,7,10,10,12,12,12,14,16,16,18,19}
}

//查找顶级boss
func find(x int) int {
    if x<0 || x>= len(un) {
        return -1
    }
    
    for {
        if un[x] == x {
            break
        } else {
            x = un[x]
        }
    }

    return x
}
//合并不相干集合
func union(x,y int){
    if find(x) != find(y){
        un[y] = un[x] 
    }
}

//依赖压缩o(n2)
func compress_1() {
    un_1 := un
    for i:=0; i<len(un_1);i++{
        un_1[i] = find(i)
    }
    fmt.Println(un_1)
}

//依赖压缩o(n2)
//稍微高效，只处理依赖非顶级节点的
func compress_2() {
    m := make(map[int]int)
    un_2 := un
    for i:=0;i<len(un_2);i++ {
        if i == un_2[i] {
            m[i] = i
        }
    }

    for i:=0;i<len(un_2);i++ {
        if _, ok := m[un_2[i]]; !ok {
            un_2[i] = find(i)
        }
    }
    fmt.Println(un_2)
}

/*
//连接所有需要修改(新建)的路径
解决问题：比如已知地区各村之间通路关系，计算怎样建路能使村村联通
1.归一化:双向关系转单向
2.连接任意独立区域的根阶段,连接数即为最少需要新增且连通性最短的路径数
*/
func union_all() int {
    m := make(map[int]int)
    for i:=0;i<len(un);i++ {
        if i == un[i] {
            m[i] = i
        }
    }
    i := un[0]
    for k,_ := range m {
        if i !=k {
            fmt.Println(1)
            union(i,k)
        }
    } 
    fmt.Println(un)
    return i
}

func main(){
    fmt.Println(un)
    fmt.Println("8's boos:",find(8))
    union(7,8)
    fmt.Println("7 union 8")
    fmt.Println(un)
    fmt.Println("7 union 11")
    union(7,11)
    fmt.Println(un)
    fmt.Println("8's boss", find(8))
    fmt.Println(un)
    compress_1()
    compress_2()
    fmt.Println("union all:",union_all())
    compress_1()
    compress_2()
}
