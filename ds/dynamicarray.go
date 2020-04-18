package ds

import "fmt"

type DArr struct {
	arr []int
}

func (dr *DArr) get(i int) (int, error) {
	if i < len(dr.arr) {
		return dr.arr[i], nil
	}
	return -1, fmt.Errorf("input index %v is out of bounds for the given array", i)
}

func (dr *DArr) add(data ...int) {
	dr.addAll(data)
}

func (dr *DArr) addAll(data []int) {
	m := len(dr.arr)
	n := m + len(data)

	// Increase allocation if needed
	if n > cap(dr.arr) {
		//Always allocate double of whats needed
		newSlice := make([]int, n, (n+1)*2)
		copy(newSlice, dr.arr)
		dr.arr = newSlice
	}
	//Update the slice since
	dr.arr = dr.arr[0:n]
	copy(dr.arr[m:n], data)
}
