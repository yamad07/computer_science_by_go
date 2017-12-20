// 並列処理を行うプログラム 一つのコアの上で複数のプログラムが動く

package main

import (
    "fmt"
    "time"
)

func sub(){
    fmt.Println("sub thread is starting")
    // printされたあとに二秒間をあける
    time.Sleep(3 * time.Second)
    fmt.Println("sub thread is finishing")
}

func main(){
    go sub()
    time.Sleep(time.Second)
    fmt.Println("sub thread is running")
    time.Sleep(3 * time.Second)
}
