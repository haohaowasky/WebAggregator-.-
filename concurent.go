
package main

import (
  "time"
  "fmt"
  "sync"
)

var wg sync.WaitGroup

func say(s string){
  defer wg.Done()
  for i:= 0; i<3; i++ {
    fmt.Println(s);
    time.Sleep(time.Millisecond*100)
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
