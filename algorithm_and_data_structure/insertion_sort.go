package main

import (
    "fmt"
)

func insertion_sort(array []int) ([]int){
    for i := 1; i < len(array); i++ {
        if array[i] < array[i - 1] {
            tmp := array[i]
            t:= i
            for t > 0 && array[t] < array[t - 1] {
                array[t] = array[t - 1]
                array[t - 1] = tmp
                t--;
            }
        }
    }
    return array
}

func main(){
    array := []int{10, 2, 3, 5, 12, 1}
    sorted_array := insertion_sort(array)
    fmt.Println(sorted_array)
}
