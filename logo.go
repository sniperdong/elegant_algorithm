package main

import(
    "fmt"
)

//http://www.figlet.org/fontdb.cgi
const str = ` _______ 
(  ____ \
| (    \/
| (__    
|  __)    ___  ___ 
| (      / __|/ __|
| (____/\\__ \ (__ 
(_______/|___/\___|
`

func main(){
    fmt.Println(fmt.Sprintf("%s%s%s%s","\x1b[35m", "\x1b[1m",str,"\x1b[0m"))
}
