package main

import "fmt"

//import "fmt"

type Node struct {
	data int
	left *Node
	right *Node
}

type Btree struct {
	btree *Node
}

type iftree interface{
	Insert(num int)
	//delete(int num)
	show()
}

func (d * Node) init (num int) *Node {
	d.left = nil
	d.right = nil
	d.data = num
	return d
}

func Newnodes(num int) *Node {
	return new(Node).init(num)
}

func NewTree() *Btree {
	return new(Btree).init()
}
func (d* Btree) init() *Btree {
	d.btree = nil
	return d
}
func (d * Node)Insert (num int) (*Node) {
if d == nil {
	return Newnodes(num)
}
	if num < d.data {
	 if d.left == nil {
		d.left = Newnodes(num)
		return d
	} else {
	return d.left.Insert(num)
	}
	} else if (num > d.data){
	         if d.right == nil {
                d.right = Newnodes(num)
                return d
        } else {
		return d.right.Insert(num)
	}
	}
return d

}

func showInOrder(root *Node) {
	if (root != nil) {
		showInOrder(root.left)
		fmt.Println(root.data)
		showInOrder(root.right)
	}
}



func main() {
	 tree := NewTree()

	tree.btree = tree.btree.Insert(5)
	tree.btree.Insert(7)
	tree.btree.Insert(2)
	showInOrder(tree.btree)



}
