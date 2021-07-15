package ginyy

import "strings"

type node struct {
	part string
	children []*node
	isWild bool //是否是精准匹配
	pattern string //完整的路径，只有叶子节点才存
}

//part是将要匹配的路由项
//第0层是一个根节点
func (n *node) insert(pattern string, parts []string, height int){
	//n总是father，child是当前节点
	part := parts[height-1]
	child := n.matchChild(part)

	if child == nil {
		child = &node{part: part, isWild: n.isPartWild(part) }
		if len(parts) == height {
			child.pattern = pattern
			n.children = append(n.children, child)
			return
		}
		n.children = append(n.children, child)
	}
	if len(parts) != height {
		child.insert(pattern, parts, height + 1)
	}
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	if n.children == nil {
		n.children = []*node{}
	}
	if len(n.children) == 0 {
		return nil
	}
	if n.isPartWild(part) {
		for _, child := range n.children {
			if child.isWild {
				return child
			}
		}
	} else {
		for _, child := range n.children {
			if child.part == part {
				return child
			}
		}
	}
	return nil
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) + 1 == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height-1]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}


	return nil
}

//所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) isPartWild(part string) bool{
	if part[0] == ':' || part[0] == '*' {
		return true
	}
	return false
}