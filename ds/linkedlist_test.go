package ds

import (
	"testing"
)

func TestInsertToListIntoList(t *testing.T) {
	var list *LinkedList
	insertToList(&list, 3)
	if list == nil {
		t.Errorf("No elements were insertToListed to the list, got nil")
	}
	if list.data != 3 {
		t.Errorf("List first element data doesnt match. Wanted: %v, got: %v", 3, list.data)
	}
}

func TestInsertToListMultipleElements(t *testing.T) {
	var list *LinkedList
	insertToList(&list, 4)
	insertToList(&list, 5)
	insertToList(&list, 6)
	insertToList(&list, 7)
	insertToList(&list, 8)
	if list == nil {
		t.Errorf("No elements were insertToListed to the list, got nil")
	}
	n := length(list)
	if n != 5 {
		t.Errorf("Something went wrong with insertToListion. Wanted Element count: %v, got: %v", 5, n)
	}
}

func TestListDeleteFromList(t *testing.T) {
	var list *LinkedList
	insertToList(&list, 4)
	insertToList(&list, 5)
	insertToList(&list, 6)
	insertToList(&list, 7)
	insertToList(&list, 8)

	if &list == nil {
		t.Errorf("No elements were insertToListed to the list, got nil")
	}
	deleteFromList(&list, 7)
	n := length(list)
	if n != 4 {
		t.Errorf("Something went wrong with Deletion. Wanted Element count: %v, got: %v", 5, n)
	}
	ele := search(list, 7)
	if ele != nil {
		t.Errorf("Something went wrong with Deletion.Input Element still exists and was not deleteFromListd")
	}
}

func TestListDeleteFromList2(t *testing.T) {
	var list *LinkedList
	deleteFromList(&list, 0)
	n := length(list)
	if n != 0 {
		t.Errorf("Something went wrong with Deletion. Wanted Element count: %v, got: %v", 0, n)
	}
}

func TestListDeleteFromList3(t *testing.T) {
	var list *LinkedList
	insertToList(&list, 8)
	deleteFromList(&list, 7)
	n := length(list)
	if n != 1 {
		t.Errorf("Something went wrong with Deletion. Wanted Element count: %v, got: %v", 1, n)
	}
	ele := search(list, 8)
	if ele == nil {
		t.Errorf("Something went wrong with Deletion. It should not have deleteFromListd ele with value:8")
	}
}

func TestListDeleteFromList4(t *testing.T) {
	var list *LinkedList
	insertToList(&list, 8)
	deleteFromList(&list, 8)
	n := length(list)
	if n != 0 {
		t.Errorf("Something went wrong with Deletion. Wanted Element count: %v, got: %v", 0, n)
	}
}

func TestListdeleteFromList5(t *testing.T) {
	var list *LinkedList
	insertToList(&list, 8)
	insertToList(&list, 9)
	deleteFromList(&list, 9)
	n := length(list)
	if n != 1 {
		t.Errorf("Something went wrong with Deletion. Wanted Element count: %v, got: %v", 1, n)
	}
}

func TestSearchList2(t *testing.T) {
	var list *LinkedList
	insertToList(&list, 4)
	insertToList(&list, 5)
	insertToList(&list, 6)
	insertToList(&list, 7)
	insertToList(&list, 8)

	item := search(list, 7)
	if item == nil {
		t.Errorf("Expected to get %v and got nil", item)
	}
	n := length(list)
	if n != 5 {
		t.Errorf("Something went wrong with Deletion. Wanted Element count: %v, got: %v", 5, n)
	}

}

func TestSearchList3(t *testing.T) {
	var list *LinkedList
	insertToList(&list, 4)
	insertToList(&list, 5)
	insertToList(&list, 6)
	insertToList(&list, 7)
	insertToList(&list, 8)

	item := search(list, 7)
	if item == nil {
		t.Errorf("Expected to get %v and got nil", item)
	}
	n := length(list)
	if n != 5 {
		t.Errorf("Something went wrong with Deletion. Wanted Element count: %v, got: %v", 5, n)
	}

}

func TestListLength(t *testing.T) {
	var list *LinkedList
	n := length(list)
	if n != 0 {
		t.Errorf("Something went wrong with calculating length. Wanted Element count: %v, got: %v", 0, n)
	}
}

func TestListLength2(t *testing.T) {
	var list *LinkedList
	insertToList(&list, 4)
	n := length(list)
	if n != 1 {
		t.Errorf("Something went wrong with calculating length. Wanted Element count: %v, got: %v", 1, n)
	}
}

func TestListLength3(t *testing.T) {
	var list *LinkedList
	insertToList(&list, 4)
	insertToList(&list, 5)
	insertToList(&list, 6)
	insertToList(&list, 7)
	insertToList(&list, 8)
	n := length(list)
	if n != 5 {
		t.Errorf("Something went wrong with calculating length. Wanted Element count: %v, got: %v", 5, n)
	}
}
