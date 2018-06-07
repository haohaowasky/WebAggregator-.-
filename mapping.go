package main

import ("fmt"
        )

func main(){
    grades:= make(map[string]float32)

    grades["Timmy"] = 42
    grades["Jess"] = 92
    grades["Sam"] = 22

//    fmt.Println(grades)

//    TimsGrade := grades["Timmy"]
//    fmt.Println(TimsGrade)

//    delete(grades, "Timmy")
//    fmt.Println(grades)
    a,b := grades["Timmy"]

    fmt.Println(a,b)
    // for  j,i:= range grades{
    //     fmt.Println(i,j)
    // }
}
