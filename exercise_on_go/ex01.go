package main

import "fmt"

func main() {
    
    number_list := []int{510, 63, 7, 837, 36, 98, 9}

    max := number_list[0]

    for _, number := range number_list {
        if number > max {
            max = number
        }
    }

    fmt.Println("O maior número na lista é:", max)
}