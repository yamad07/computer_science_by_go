package main
import (
    "fmt"
)

func quick_sort(array []int) ([]int){
    if len(array) < 1{
        return array
    }
    pivot := array[0]
    var right = []int{}
    var left = []int{}
    for i := 1; i < len(array); i++ {
        if pivot > array[i] {
            left = append(left, array[i])
        } else{
            right = append(right, array[i])
        }
    }
    left = quick_sort(left)
    right = quick_sort(right)
    tmp := []int{pivot}
    left = append(left, tmp...)
    sorted_array := append(left, right...)
    return sorted_array
}

func main(){
    array := []int{6, 10, 1, 2, 11, 1, 3, 8}
    sorted_array := quick_sort(array)
    fmt.Println(sorted_array)
}
