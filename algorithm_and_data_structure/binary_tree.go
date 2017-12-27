// 二分木を構成するプログラム

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

func (t *tree) _insert(current_node *node, num int) (*node){
    if current_node == nil {
        n := &node{num, nil, nil}
        return n
    }

    if current_node.value < num {
        current_node.right_node = t._insert(current_node.right_node, num)
    }
    if current_node.value > num {
        current_node.left_node = t._insert(current_node.left_node, num)
    }
    return current_node
}

func (t *tree) insert(num int) (){
    t._insert(t.top_node, num)
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

func(t *tree) _max(current_node *node) (int) {
    if (current_node.right_node == nil){
        return current_node.value
    }
    t_max := t._max(current_node.right_node)
    return t_max;
}
func(t *tree) max() (int){
    return t._max(t.top_node)
}

func (t *tree) _min(current_node *node) (int) {
    if (current_node.left_node == nil){
        return current_node.value
    }
    t_min := t._min(current_node.left_node)
    return t_min;
}
func (t *tree) min()(int){
    return t._min(t.top_node)
}

func main(){
    top_node := &node{10, nil, nil}
    t := &tree{top_node}
    t.insert(11)
    t.insert(9)
    t.insert(13)
    t.insert(8)
    t_max := t.max()
    t_min := t.min()
    fmt.Println(t_max)
    fmt.Println(t_min)
}
