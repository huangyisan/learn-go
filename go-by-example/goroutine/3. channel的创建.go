package main

import "fmt"

type Equip struct {

}



func main() {
    ch1 := make(chan *Equip, 10)
    e1 := &Equip{
    }
    ch1 <- e1

    for {
        i, ok := <- ch1
        if !ok  {
            break
        }
        fmt.Println(i)

    }


}