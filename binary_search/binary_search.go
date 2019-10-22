package main

import (
	"fmt"
)

type Node struct {
	key		int
	left	*Node
	right	*Node
}

func (n *Node) search(key int) bool {
	// This is our base case. If n == nil, `key`
    // doesn't exist in our binary search tree.
	if n == nil {
		return false
	}

	if n.key < key { // move right
		return n.right.search(key)
	} else if n.key > key { // move left
		return n.left.search(key)
	}

	// if not < or > than, then its found
	return true
}

func (n *Node) insert(key int) {
	// if node being visited value's smaller than value we are trying to insert
	// as value is larger, we should intert into the right node, so we check if the right is nil
	// if it is, then we set the right value of visiting node to a new node with key
	// else then we recursively go to the next right node,
	// as the current node is smaller, we must go right
	if n.key < key {
		if n.right == nil {
			n.right = &Node{key: key}
		} else {
			n.right.insert(key)
		}
	} else if n.key > key {
		if n.left == nil {
			n.left = &Node{key: key}
		} else {
			n.left.insert(key)
		}
	}
}

func (n *Node) delete(key int) *Node {
    // search for `key`
    if n.key < key {
        n.right = n.right.delete(key)
    } else if n.key > key {
        n.left = n.left.delete(key)   
    // n.Key == `key`
    } else {
        if n.left == nil { // just point to opposite node 
            return n.right
        } else if n.right == nil { // just point to opposite node 
            return n.left
        }

        // if `n` has two children, you need to 
        // find the next highest number that 
        // should go in `n`'s position so that
        // the BST stays correct 
        min := n.right.min()
 
        // we only update `n`'s key with min
        // instead of replacing n with the min
        // node so n's immediate children aren't orphaned
        n.key = min
        n.right = n.right.delete(min)
    }
    return n
}

// keep searching left until you find a nil pointer
func (n *Node) min() int {
	if n.left == nil {
		return n.key
	}
	return n.left.min()
}

func (n *Node) max() int {
	if n.right == nil {
		return n.key
	}
	return n.right.max()
}

func main() {
	// default type for pointer is nil
	tree := &Node{
		key: 6,
		left: nil,
		right: nil,
	}
	fmt.Println(tree)
}