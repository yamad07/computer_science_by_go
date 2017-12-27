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

func (l *linked_list) _printList(current_node *node) (int){
    if current_node == nil{
        return 0
    }
    fmt.Println(current_node.value)
    l.printlist(current_node.next_node)
    return 0
}
func (l *linked_list) printList()(int){
    return l._printList(l.top_node)
}

func (l *linked_list) _insert(current_node *node, num int) (*node){
    if current_node == nil{
        new_node := &node{num, nil}
        return new_node
    }
    current_node.next_node = l.insert(current_node.next_node, num)
    return current_node
}

func (l *linked_list) insert()(*node){
    l._insert(l.top_node)
}

func main(){
    top_node := &node{10, nil}
    ll := &linked_list{top_node}
    ll.insert(10)
    ll.insert(11)
    ll.insert(13)
    ll.insert(1)
    ll.printList()
}
