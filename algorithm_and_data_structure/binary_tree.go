package main

import (
    "fmt"
)

type tree struct{
    top_node *node
}
type node struct{
    value int
    left_node *node
    right_node *node

}

func (t *tree) insert(current_node *node, num int) (*node){
    if current_node == nil {
        n := &node{num, nil, nil}
        return n
    }

    if current_node.value < num {
        current_node.right_node = t.insert(current_node.right_node, num)
    }
    if current_node.value > num {
        current_node.left_node = t.insert(current_node.left_node, num)
    }
    return current_node
}

func (t *tree) print_node(current_node *node) (int){
    if (current_node == nil){
        return 0;
    }
    fmt.Println(current_node.value)
    t.print_node(current_node.left_node)
    t.print_node(current_node.right_node)
    return 0;
}

func main(){
    top_node := &node{10, nil, nil}
    t := &tree{top_node}
    t.insert(t.top_node, 11)
    t.insert(t.top_node, 8)
    t.print_node(top_node)
}
