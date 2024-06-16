package main
 import (
    "fmt"
    "sync"
    )
    
func f(n int , wg *sync .WaitGroup){
    for i:=0 ; i<n ; i++{
        fmt.Println(i)
    }
}

func main(){
    var n int
    var wg sync.WaitGroup
    
    fmt.Println("Enter an integer : ")
    fmt.Scanln(&n)
    
    wg.Add(1)
    
    go f(n , &wg)
    wg.Wait()
}
