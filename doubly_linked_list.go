package doubly_linked_list

type Item struct {
	list  *List
	value interface{}
	next  *Item
	prev  *Item
}

func (item *Item) Value() interface{} {
	return item.value
}

func (item *Item) Next() *Item {
	return item.next
}

func (item *Item) Prev() *Item {
	return item.prev
}

func (item *Item) Remove() {

	if item.next == nil && item.prev == nil {
		item.list.first = nil
	}

	if item.prev != nil {
		item.prev.next = item.next
	}

	if item.next != nil {
		item.next.prev = item.prev
	}

}

type List struct {
	first *Item
}

func (list *List) PushFront(value interface{}) {

	if list.first != nil {
		oldFirst := list.first
		list.first = &Item{list: list, value: value, next: oldFirst}
		oldFirst.prev = list.first
	} else {
		list.first = &Item{list: list, value: value}
	}
}

func (list *List) PushBack(value interface{}) {
	item := list.Last()
	if item != nil {
		item.next = &Item{list: list, value: value, prev: item}
	} else {
		list.PushFront(value)
	}
}

func (list *List) First() *Item {
	return list.first
}

func (list *List) Last() *Item {
	elem := list.First()

	for elem != nil && elem.Next() != nil {
		elem = elem.Next()
	}
	return elem
}

func (list *List) Len() int {
	if list.first == nil {
		return 0
	}
	counter := 1
	elem := list.First()
	for elem.Next() != nil {
		counter += 1
		elem = elem.Next()
	}
	return counter
}
