package avito_interview

import "fmt"

type Node struct {
	title    string
	children *Tree
}

type Tree = []Node

func printTree(rootNodes Tree) {
	var traverse func(nodes Tree, path string)

	traverse = func(nodes Tree, path string) {
		for _, node := range nodes {
			currentPath := path + node.title
			if node.children == nil || len(*node.children) == 0 {
				// Leaf node, print the full path
				fmt.Println(currentPath)
			} else {
				// Non-leaf node, continue traversing
				traverse(*node.children, currentPath+" => ")
			}
		}
	}

	traverse(rootNodes, "")
}
