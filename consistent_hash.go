package main

//顺时针定位节点
import(
    "fmt"
    "hash/crc32"
    "sort"
    "sync"
    "strconv"
    "strings"
    "errors"
)

//虚拟节点
const (
    REPLICAS uint8 = 2
)
type Ring []int
func (c Ring) Len() int {
    return len(c)
}

func (c Ring) Less(i, j int) bool {
    return c[i] < c[j]
}

func (c Ring) Swap(i, j int) {
    c[i], c[j] = c[j], c[i]
}

//物理节点
//权重1-255
type Node struct {
    Ip string
    Port    int
    Weight  uint8
}

type ConRes struct {
    ring Ring
    nodes map[int]*Node //key:hash(ip_replica#weight)
    members map[int]bool //key:ipv4
    numRes int
    sync.RWMutex
}

//对于key冲突可以不处理，一般只影响虚拟节点数量
//适当告警即可
func (c *ConRes) Add(node *Node) error {
    c.Lock()
    defer c.Unlock()

    convertErr,ip := ipInt32(node.Ip)
    if convertErr != nil {
        return convertErr
    }
    if _,ok:=c.members[ip]; ok {
        return errors.New("tht ip has been existed,check it")
    }
    
    genErr, keys := genKeys(node.Ip, REPLICAS, node.Weight)
    if genErr != nil {
        return genErr
    }
    
    keyMap := make(map[string]int)
    for _,key := range keys {
        keyMap[key] = hash(key)
        if _,ok := c.nodes[keyMap[key]]; ok {
            return errors.New("key conflict")
        }
    }
    if len(keys) != len(keyMap) {
        return errors.New("key conflict")
    }
    
    for _,v := range keyMap{
        c.nodes[v] = node
        c.ring = append(c.ring, v)
    }
    sort.Sort(c.ring)
    c.members[ip] = true
    c.numRes++
    return nil
}

func (c *ConRes) search1(hash int) int{
    var i,j,n int
    j = len(c.ring) - 1
    if hash > c.ring[j] || hash <= c.ring[0] {
        return 0
    }
    for {
        n = (i+j) / 2
        if hash == c.ring[n] {
            break
        } else if hash > c.ring[n] {
            i = n + 1
        } else {
            j = n - 1
        }
        if i > j {
            break
        }
    }
    if n >= len(c.ring) {
        return 0
    } else {
        if hash <= c.ring[n] {
            return n
        } else {
            return n+1
        }
    }
}

func (c *ConRes) search2(hash int) int{
    f := func(i int) bool { return c.ring[i] >= hash }
    i := sort.Search(len(c.ring), f)
    if i >= len(c.ring) {
        return 0
    } else {
        return i
    }
}

func hash(key string) int {
    return int(crc32.ChecksumIEEE([]byte(key)))
}

func ipInt32(str string) (err error, ip int) {
    strArr := strings.Split(str, ".")
    if len(strArr) != 4 {
        err = errors.New("invalid ip address, must be format:x.x.x.x ")
    }
    var n int
    for i:=0;i<4;i++ {
        ip = ip << 8
        n,err = strconv.Atoi(strArr[i])
        if err != nil {
            return
        }
        ip += int(n)
    }
    return
}

func genKeys(ip string, replicas, weight uint8) (err error, keys []string) {
    if 0 == weight || 0 == replicas {
        err = errors.New("invalid value in replicas or weight")
        return 
    }
    
    var str string
    var i,j uint8
    for i=0; i < replicas; i++ {
        for j=0; j < weight; j++ {
            str = ip + "_" + strconv.Itoa(int(i)) + "#" + strconv.Itoa(int(j))
            keys = append(keys, str)
        }
    }
    
    return
}

func main() {
    con := &ConRes{nodes:make(map[int]*Node),members:make(map[int]bool)}
    con.Add(&Node{Ip:"127.0.0.1", Port:80, Weight:1})
    con.Add(&Node{Ip:"127.0.0.2", Port:80, Weight:2})
    con.Add(&Node{Ip:"127.0.0.3", Port:80, Weight:4})
    con.Add(&Node{Ip:"127.0.0.4", Port:80, Weight:8})

    k := hash("hello word!")
    fmt.Println(k, con.search1(k), con.search2(k))
    fmt.Println(0, con.search1(0), con.search2(0))
    fmt.Println(209498633, con.search1(209498633),con.search2(209498633))
    fmt.Println(209498634, con.search1(209498634), con.search2(209498634))
    fmt.Println(230611005, con.search1(230611005), con.search2(230611005))
    fmt.Println(230611006, con.search1(230611006), con.search2(230611006))
    fmt.Println(230611007, con.search1(230611007), con.search2(230611007))
    fmt.Println(4294967296, con.search1(4294967296), con.search2(4294967296))
    k = hash("dongping!an")
    fmt.Println(k, con.search1(k), con.search2(k))
    k = hash("psxf")
    fmt.Println(k, con.search1(k), con.search2(k))
    k = hash("kjgaldsjgjeiortjwertgdlksgnowi4tnasognoiw4t36n;lkn;sdg;46s")
    fmt.Println(k, con.search1(k), con.search2(k))
}
