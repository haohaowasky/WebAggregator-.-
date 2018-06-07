package main
import (
    "fmt"
    "sync"
  )

var wg sync.WaitGroup // wait signal

func foo(c chan int, someValue int){
    defer wg.Done()
    c <- someValue * 5
}

func main(){
    fooVal := make(chan int, 10) // 10 is the buffering,
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go foo(fooVal, i)
    }
    wg.Wait()
    close(fooVal)

    for item := range fooVal{
        fmt.Println(item)
    }
}
