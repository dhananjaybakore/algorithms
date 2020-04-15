package dataStructures

const SIZE = 10
type DArr struct {
	arr [SIZE] int
	resizeCount int
	currentLen int
}




func (dr * DArr) length() int  {
	return dr.currentLen
}

func (dr *DArr) add(item int)  {
	if dr.currentLen < len(dr.arr) {
		dr.arr[dr.currentLen+1] = item
		dr.currentLen++
	}else{
		// Increase the array size
		updArr := make([]int, dr.length(), dr.length()*2)

	}
}

