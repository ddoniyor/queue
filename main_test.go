package main

import (
	"testing"
)

func Test_empty_queue(t *testing.T) {
	queueList := queue{}
	if queueList.len() != 0 {
		t.Error("new list length must be zero, got: ", queueList.len())
	}
	if queueList.first() != nil {
		t.Error("first element of queue must have nil value, got: ", queueList.first(), "want", nil)
	}
	if queueList.last() != nil {
		t.Error("last element of queue must have nil value, got: ", queueList.last(), "want", nil)
	}
	if queueList.first() != queueList.last() {
		t.Error("first and last elements must have the same values, got: ", queueList.first(), ";", queueList.last())
	}

}

func Test_queue_with_one_result(t *testing.T) {
	queueList := queue{}
	queueList.equeue(1)
	if queueList.len() != 1 {
		t.Error("after adding one element size must be 1, got: ", queueList.len(), "want", 1)
	}
	if queueList.first() != 1 {
		t.Error("first element of queue must be 1, got:", queueList.first(), "want", 1)
	}
	if queueList.last() != 1 {
		t.Error("last element of queue must be 1, got:", queueList.last(), "want", 1)
	}
	if queueList.last() != queueList.first() {
		t.Error("first and last elements must have the same values, got:", queueList.last(), ";", queueList.first())
	}

}

func Test_queue_with_multiple_result(t *testing.T) {
	queueList := queue{}
	queueList.equeue(1)
	queueList.equeue(2)
	queueList.equeue(3)
	if queueList.len() != 3 {
		t.Error("after adding 3 elements size must be 3, got: ", queueList.len())
	}
	if queueList.first() != 1 {
		t.Error("first element of queue must be 1, got:", queueList.first(), "want", 1)
	}
	if queueList.last() != 3 {
		t.Error("last element of queue must be 3, got:", queueList.last(), "want", 3)
	}
}

func Test_dequeue_for_function(t *testing.T) {
	queueList := queue{}
	queueList.equeue(1)
	queueList.equeue(2)
	queueList.equeue(3)
	queueList.dequeue()

	if queueList.len() != 2 {
		t.Error("after removing firstEl element size must be 2, got: ", queueList.len())
	}
	if queueList.first() != 2 {
		t.Error("first element of queue must be 2, got:", queueList.first(), "want", 2)
	}
	if queueList.last() != 3 {
		t.Error("last element of queue must be 3, got:", queueList.last(), "want", 3)
	}

}
func Test_dequeue_removed_single_element(t *testing.T) {
	queueList := queue{}
	queueList.equeue(1)

	queueList.dequeue()

	if queueList.len() != 0 {
		t.Error("after removing firstEl element size must be 0, got: ", queueList.len())
	}
	if queueList.first() != nil {
		t.Error("first element of queue must be 1, got:", queueList.first(), "want", nil)
	}
	if queueList.last() != nil {
		t.Error("last element of queue must be 1, got:", queueList.last(), "want", nil)
	}
	if queueList.last() != queueList.first() {
		t.Error("first and last elements must have the same values, got:", queueList.last(), ";", queueList.first())
	}
}

func Test_dequeue_for_no_result(t *testing.T) {
	queueList := queue{}

	queueList.dequeue()

	if queueList.len() != 0 {
		t.Error("after removing firstEl element size must be 0, got: ", queueList.len())
	}
	if queueList.first() != nil {
		t.Error("first element of queue must be nil, got:", queueList.first(), "want", nil)
	}
	if queueList.last() != nil {
		t.Error("last element of queue must be nil, got:", queueList.last(), "want", nil)
	}
	if queueList.last() != queueList.first() {
		t.Error("first and last elements must have the same values, got:", queueList.last(), ";", queueList.first())
	}
}
