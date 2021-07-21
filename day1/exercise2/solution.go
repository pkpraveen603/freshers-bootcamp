package main

import "fmt"

type Tree struct {
	nodeVal string
	left *Tree
	right *Tree
}

func CreateNode(val string) *Tree{
	node := Tree{val,nil,nil}
	return &node
}

func PreOrder(head *Tree){
	if head==nil {
		return
	}
	fmt.Printf(" %s ",head.nodeVal)
	if head.left!=nil {
		PreOrder(head.left)
	}
	if head.right!=nil {
		PreOrder(head.right)
	}
}

func PostOrder(head *Tree){
	if head==nil{
		return
	}
	if head.left!=nil {
		PostOrder(head.left)
	}
	if head.right!=nil {
		PostOrder(head.right)
	}
	fmt.Printf(" %s ",head.nodeVal)
}

func main(){
	root:=CreateNode("+")
	root.left=CreateNode("a")
	root.right=CreateNode("-")
	root.right.left=CreateNode("b")
	root.right.right=CreateNode("c")
	fmt.Println("Pre Order Traversal:")
	PreOrder(root)
	fmt.Println("\n Post Order Traversal:")
	PostOrder(root)
}
