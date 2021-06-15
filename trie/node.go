package trie

import (
	"sync"
)

type Node struct {
	Rank     int
	children map[string]*Node
	lock     sync.RWMutex
}

func NewTrieNode(rank int) *Node {
	return &Node{
		Rank:     rank,
		children: make(map[string]*Node),
	}
}

func (n *Node) Add(key string, rank int) {
	n.lock.Lock()
	defer n.lock.Unlock()

	if !n.hasChild(key) {
		n.children[key] = NewTrieNode(rank)
	}

}

func (n *Node) hasChild(key string) bool {
	_, ok := n.children[key]
	return ok
}

func (n *Node) Child(key string) (*Node, bool) {
	n.lock.RLock()
	defer n.lock.RUnlock()

	child, ok := n.children[key]
	return child, ok
}
