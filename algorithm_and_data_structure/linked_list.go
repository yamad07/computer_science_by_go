package main

import (
    "fmt"
)

type linked_list struct{
    top_node *node
}

type node struct{
    value int
    next_node *node
}

func (l *linked_list) printlist(current_node *node) (int){
    if current_node == nil{
        return 0
    }
    fmt.Println(current_node.value)
    l.printlist(current_node.next_node)
    return 0
}

func (l *linked_list) insert(current_node *node, num int) (*node){
    if current_node == nil{
        new_node := &node{num, nil}
        return new_node
    }
    current_node.next_node = l.insert(current_node.next_node, num)
    return current_node
}

func main(){
    top_node := &node{10, nil}
    ll := &linked_list{top_node}
    ll.insert(top_node, 10)
    ll.insert(top_node, 11)
    ll.insert(top_node, 13)
    ll.insert(top_node, 1)
    ll.printlist(top_node)
}
