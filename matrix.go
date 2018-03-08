package main

import "fmt"

const (
    oriRight = 0
    oriDown = 1
    oriLeft = 2
    oriUp   = 3
)

type Leader struct {
    Ori int
    Next *Leader
}

func main(){
    var head,r,d,l,u *Leader
    r=&Leader{Ori:oriRight}
    d=&Leader{Ori:oriDown}
    l=&Leader{Ori:oriLeft}
    u=&Leader{Ori:oriUp}
    head=r
    r.Next=d
    d.Next=l
    l.Next=u
    u.Next=r
    mn := [6][6]int{
        {1,2,3,4,5,6},
        {20,21,22,23,24,7},
        {19,32,33,34,25,8},
        {18,31,36,35,26,9},
        {17,30,29,28,27,10},
        {16,15,14,13,12,11},
    }
    var rm,dm,lm,um,i,j int
    rm = len(mn[0])-1
    dm = len(mn)-1
    lm = 0
    um = 1//因为已经处于横行，um直接指向下一行
    for {
        fmt.Println(mn[i][j])
        switch head.Ori {
            case oriRight:
                j++
                if j>rm {
                    rm--
                    head=head.Next
                    j--
                    i++
                    if i>dm{
                        goto END
                    }
                }
            case oriDown:
                i++
                if i>dm{
                    dm--
                    head=head.Next
                    i--
                    j--
                    if j<lm{
                        goto END
                    }
                }
            case oriLeft:
                j--
                if j<lm {
                    lm++
                    head=head.Next
                    j++
                    i--
                    if i<um{
                        goto END
                    }
                }
            case oriUp:
                i--
                if i<um{
                    um++
                    head=head.Next
                    i++
                    j++
                    if j>rm{
                        goto END
                    }
                }
        }
    }
END:
}
