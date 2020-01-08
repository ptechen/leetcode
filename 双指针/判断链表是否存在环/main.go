package main

type Head struct {
	Next *Head
}


func hasCycle(params *Head) bool {
	if params == nil {
		return false
	}
	l1 := params
	l2 := params.Next
	for l1 != nil && l2 != nil && l2.Next != nil {
		if l1 == l2 {
			return true
		}
		l1 = l1.Next
		l2 = l2.Next.Next
	}

	return false
}
