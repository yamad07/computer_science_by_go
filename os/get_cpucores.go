// 複数のコアを使うための並列処理の、コア数を取得するプログラム

package main

import(
    "runtime"
    "fmt"
)

func main(){
    cpus := runtime.NumCPU()
    fmt.Printf("Num of Multi Core %d \n", cpus)
}
