package main

type queue struct {
	firstEl *queueNode
	lastEl  *queueNode
	size    int
}

type queueNode struct {
	next     *queueNode
	prev     *queueNode
	value    interface{}
	priority int
}

type ByQueue []queueNode

func (receiver ByQueue) Len() int {
	return len(receiver)
}

func (receiver ByQueue) Less(i, j int) bool {
	return receiver[i].priority > receiver[j].priority
}

func (receiver ByQueue) Swap(i, j int) {
	receiver[i].priority, receiver[j].priority = receiver[j].priority, receiver[i].priority
}

func (receiver *queue) len() int {
	return receiver.size
}

func (receiver *queue) first() *queueNode {
	return receiver.firstEl
}

func (receiver *queue) last() *queueNode {
	return receiver.lastEl
}

func (receiver *queue) equeue(elementPtr interface{}) {
	if receiver.len() == 0 {
		receiver.firstEl = &queueNode{
			next:  nil,
			prev:  nil,
			value: elementPtr,
		}
		receiver.lastEl = receiver.firstEl
		receiver.size++
		return
	}
	receiver.size++

	nextValuePtr := receiver.firstEl
	for {
		if nextValuePtr.next == nil {
			nextValuePtr.next = &queueNode{
				next:  nil,
				prev:  nextValuePtr,
				value: elementPtr,
			}
			return
		}
		nextValuePtr = nextValuePtr.next
	}

}

func (receiver *queue) dequeue() {
	if receiver.len() == 0 {
		return
	}
	if receiver.len() == 1 {
		receiver.size = 0
		return
	}
	receiver.firstEl = receiver.firstEl.next
	current := receiver.firstEl
	receiver.size = 1

	for current.next != nil {
		current = current.next
		receiver.size++
	}

	if receiver.len() == 0 {
		receiver.lastEl = receiver.firstEl

	}

}

func main() {

}
