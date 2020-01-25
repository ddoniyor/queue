package main

import (
	"testing"
)

func Test_empty_queue(t *testing.T) {
	queueList := queue{}
	if queueList.len() != 0 {
		t.Error("new list length must be zero, got: ", queueList.len())
	}

}

func Test_queue_with_one_result(t *testing.T) {
	queueList := queue{}
	queueList.equeue(1)
	if queueList.len() != 1 {
		t.Error("after adding one element size must be 1, got: ", queueList.len())
	}
}

func Test_queue_with_multiple_result(t *testing.T) {
	queueList := queue{}
	queueList.equeue(1)
	queueList.equeue(1)
	queueList.equeue(1)
	if queueList.len() != 3 {
		t.Error("after adding 3 elements size must be 3, got: ", queueList.len())
	}
}

func Test_dequeue_for_function(t *testing.T) {
	queueList := queue{}
	queueList.equeue(1)
	queueList.equeue(1)
	queueList.equeue(1)
	queueList.dequeue()

	if queueList.len() != 2 {
		t.Error("after removing firstEl element size must be 2, got: ", queueList.len())
	}
}
func Test_dequeue_removed_single_element(t *testing.T) {
	queueList := queue{}
	queueList.equeue(1)

	queueList.dequeue()

	if queueList.len() != 0 {
		t.Error("after removing firstEl element size must be 0, got: ", queueList.len())
	}
}

func Test_dequeue_for_no_result(t *testing.T) {
	queueList := queue{}

	queueList.dequeue()

	if queueList.len() != 0 {
		t.Error("after removing firstEl element size must be 0, got: ", queueList.len())
	}
}
