package main

import (
  "fmt"
)

func main() {
    i:=0
    for i:=0; i<10; i++{
      fmt.Println(i)
      i++
    }
}


func infiniteloop(){
    for {
        fmt.Println(i)
        i+=5
    }
}

func whileloop(){
    x:= 5
    for {
        fmt.Println("what ever", x)
        x+=3
        if x> 25{
          break
        }
    }
}

func simple(){
    for x:=5; x<25; x+=3{
        fmt.Println("Do stuff", x)
        a+=4
    }
}
