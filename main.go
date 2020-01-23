package main

type queue struct {
	first *queueNode
	last  *queueNode
	size  int


}

type queueNode struct {
	next  *queueNode
	prev  *queueNode
	value interface{}
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

func (receiver *queue) firstP() *queueNode {
	return receiver.first
}

func (receiver *queue) lastP() *queueNode {
	return receiver.last
}

func (receiver *queue) equeue(elementPtr interface{}) {
	if receiver.len() == 0 {
		receiver.first = &queueNode{
			next:  nil,
			prev:  nil,
			value: elementPtr,
		}
		receiver.last = receiver.first
		receiver.size++
		return
	}
	receiver.size++

	nextValuePtr := receiver.first
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
	receiver.first = receiver.first.next
	current := receiver.first
	receiver.size = 1

	for current.next != nil {
		current = current.next
		receiver.size++
	}

	if receiver.len() == 0 {
		receiver.last = receiver.first

	}

}

func main() {

}
