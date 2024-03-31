package hashStorage

import (
	"time"
)

type Node struct {
	Next            *Node
	timeOfExistence int
	Value           int
}

func (n *Node) PushBack(node *Node) {
	if n == nil {
		n = node
		n.Next = nil
		return
	}

	for n.Next != nil {
		n = n.Next
	}

	n.Next = node

}

func (n *Node) SetTimeOfExistence(duration int) {
	if duration != 0 {
		n.timeOfExistence = int(time.Now().Unix()) + duration
	}
}
