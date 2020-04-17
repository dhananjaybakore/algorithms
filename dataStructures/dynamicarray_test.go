package dataStructures

import (
	"testing"
)

func TestCreatArr(t *testing.T) {
	arr := DArr{}
	n := len(arr.arr)
	if n != 0 {
		t.Errorf("Did not get expected result. Wanted: %v, got: %v", 1, n)
	}
}

func TestAddItem(t *testing.T) {
	arr := DArr{}
	arr.add(2)
	arr.add(3)
	arr.add(4)
	n := len(arr.arr)
	if n != 3 {
		t.Errorf("Did not get expected result. Wanted: %v, got: %v", 3, n)
	}
}

func TestAddItem2(t *testing.T) {
	arr := DArr{}
	arr.add(2)
	n := len(arr.arr)
	if n != 1 {
		t.Errorf("Did not get expected result. Wanted: %v, got: %v", 1, n)
	}
}

func TestAddMultipleItems(t *testing.T) {
	arr := DArr{}
	arr.add(2, 3, 4, 5, 7)
	n := len(arr.arr)

	firstEle, _ := arr.get(0)
	if n != 5 {
		t.Errorf("Did not get expected result. Wanted: %v, got: %v", 5, n)
	}

	if firstEle != 2 {
		t.Errorf("Did not get expected result FirstEle. Wanted: %v, got: %v", 2, n)
	}
}

func TestAddAll(t *testing.T) {
	arr := DArr{}

	arr.addAll([]int{1, 2, 3, 4, 5, 6})
	n := len(arr.arr)

	firstEle, _ := arr.get(0)
	if n != 6 {
		t.Errorf("Did not get expected result. Wanted: %v, got: %v", 6, n)
	}

	if firstEle != 1 {
		t.Errorf("Did not get expected result FirstEle. Wanted: %v, got: %v", 1, n)
	}
}

func BenchmarkAddMultiple(b *testing.B) {
	arr := DArr{}
	data := []int{1, 2, 3, 4, 5, 6}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		arr.addAll(data)
	}
}
