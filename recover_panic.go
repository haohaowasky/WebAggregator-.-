
package main

import (
  "time"
  "fmt"
  "sync"
)

var wg sync.WaitGroup // wait signal

func cleanup(){
    defer wg.Done()
    if r:= recover(); r != nil{
        fmt.Println("Recoverd in cleanup:", r)
    }
}

func say(s string){
  defer cleanup()
  for i := 0; i < 3; i++ {
    fmt.Println(s);
    time.Sleep(time.Millisecond*100)
    if i == 2{
        panic("oh, wrong")
    }
  }
}

func main(){
  wg.Add(1)
  go say("Hey") // run parallel, so it will happen
  wg.Add(1)
  go say("There")
  wg.Wait()

  go say("hihi")
}
