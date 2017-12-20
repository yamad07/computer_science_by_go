package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

func main(){
    service := ":7777"
    // portのtcpAddrを取得
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)
    // tcpAddrの中でデータを受信できる状態にする
    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        // 通信が確立されたらdaytimeをStringにしたもののbyte列を送信する。
        daytime := time.Now().String()
        conn.Write([]byte(daytime))
        conn.Close()
    }
}

func checkError(err error){
    if err != nil{
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
