package ds

import "fmt"

type LinkedList struct {
	data int
	next *LinkedList
}

/*
	o(n) time and o(1) space
	Tail Recursive
*/
func search(list *LinkedList, data int) *LinkedList {
	if list == nil {
		return nil
	}

	if list.data == data {
		return list
	} else {
		return search(list.next, data)
	}
}

/*
	O(1)
*/
func insertToList(list **LinkedList, data int) {
	item := LinkedList{data: data, next: *list}
	*list = &item
}

// O(n) time, O(1) space
func deleteFromList(list **LinkedList, data int) {
	// O(n) time, O(1) space
	item := search(*list, data)
	if item != nil {
		//O(n) time, O(1) space
		pred := predecessor(*list, data)
		if pred == nil {
			//remove first element from the head
			*list = item.next
			fmt.Printf("Value of list: %v", list)
		} else {
			pred.next = item.next
		}
	}
}

// O(n) time, O(1) space
func predecessor(list *LinkedList, data int) *LinkedList {
	if list == nil || list.next == nil {
		return nil
	}
	if list.next.data == data {
		return list
	} else {
		return predecessor(list.next, data)
	}
}

func length(list *LinkedList) int {
	if list == nil {
		return 0
	}
	return 1 + length(list.next)
}
