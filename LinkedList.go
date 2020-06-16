/// Created by Kirk George
/// Copyright: Kirk George

package LinkedList

type Node struct {
	prev  *Node
	next  *Node
	value interface{}
}

type iterator struct {
	index int
	*Node
}

type LinkedList struct {
	head  *Node
	tail  *Node
	count int
}

type ISet interface {
	ToArray() []interface{}
	Count() int
}

type IList interface {
	At(index int) interface{}
	Find(value interface{}) int
	IQueue
	IStack
	Remove(value interface{})
	RemoveAt(index int) interface{}
	Iterator() IIterator
}

type IQueue interface {
	Peek() interface{}
	Enqueue(value interface{})
	Dequeue() interface{}
	ISet
}

type IStack interface {
	Peek() interface{}
	Push(value interface{})
	Pop() interface{}
	ISet
}

type IIterator interface {
	Current() (interface{}, bool)
	MoveNext() (interface{}, bool)
	MovePrev() (interface{}, bool)
	Index() int
}

func List() IList {
	return &LinkedList{head: nil, tail: nil, count: 0}
}

func Queue() IQueue {
	return &LinkedList{head: nil, tail: nil, count: 0}
}

func Stack() IStack {
	return &LinkedList{head: nil, tail: nil, count: 0}
}

func (list *LinkedList) Iterator() IIterator {
	return &iterator{Node: list.head, index: 0}
}

func (iter *iterator) Current() (interface{}, bool) {
	if iter.Node != nil {
		return iter.value, true
	} else {
		return nil, false
	}
}

func (iter *iterator) MoveNext() (interface{}, bool) {
	if iter.Node != nil && iter.next != nil {
		iter.index++
		iter.Node = iter.next
		return iter.value, true
	} else {
		return nil, false
	}
}

func (iter *iterator) MovePrev() (interface{}, bool) {
	if iter.Node != nil && iter.prev != nil {
		iter.index--
		iter.Node = iter.prev
		return iter.value, true
	} else {
		return nil, false
	}
}

func (iter *iterator) Index() int {
	return iter.index
}

func (list *LinkedList) Count() int {
	return list.count
}

func (list *LinkedList) Peek() interface{} {
	if list.tail == nil {
		return nil
	}
	return list.tail.value
}

func (list *LinkedList) Enqueue(value interface{}) {
	toAdd := &Node{prev: list.tail, next: nil, value: value}

	if list.tail != nil {
		list.tail.next = toAdd
	}
	list.tail = toAdd

	if list.head == nil {
		list.head = toAdd
	}

	list.count++
}

func (list *LinkedList) Dequeue() interface{} {
	if list.head == nil {
		return nil
	}

	remove := list.head
	list.head = remove.next
	if list.head != nil {
		list.head.prev = nil
	}

	if list.tail == remove {
		list.tail = nil
	}

	remove.next = nil
	remove.prev = nil

	list.count--

	return remove.value
}

func (list *LinkedList) Push(value interface{}) {
	list.Enqueue(value)
}

func (list *LinkedList) Pop() interface{} {
	if list.tail == nil {
		return nil
	}

	remove := list.tail
	list.tail = remove.prev
	if list.tail != nil {
		list.tail.next = nil
	}

	if list.head == remove {
		list.head = nil
	}

	list.count--

	if list.count == 0 {
		list.head = nil
	}

	return remove.value
}

func (list *LinkedList) ToArray() []interface{} {
	result := make([]interface{}, list.count)

	var cur *Node = list.head
	i := 0
	for cur != nil {
		result[i] = cur.value
		cur = cur.next
		i++
	}

	return result
}

func (list *LinkedList) At(index int) interface{} {
	var cur *Node = list.head
	for i := 0; i < index; i++ {
		cur = cur.next
		if cur == nil {
			return nil
		}
	}

	return cur.value
}

func (list *LinkedList) Find(value interface{}) int {
	found := -1
	var cur *Node = list.head
	i := 0
	for cur != nil {
		if cur.value == value {
			found = i
			break
		}

		i++
		cur = cur.next
	}

	return found
}

func (list *LinkedList) Remove(value interface{}) {
	var cur *Node = list.head
	i := 0
	for cur != nil {
		if cur.value == value {
			prev := cur.prev
			next := cur.next

			if prev != nil {
				prev.next = next
			}
			if next != nil {
				next.prev = prev
			}

			if list.head == cur {
				list.head = nil
			}
			if list.tail == cur {
				list.tail = nil
			}
			break
		}

		i++
		cur = cur.next
	}
}

func (list *LinkedList) RemoveAt(index int) interface{} {
	var cur *Node = list.head
	for i := 0; i <= index; i++ {
		if i == index {
			prev := cur.prev
			next := cur.next

			if prev != nil {
				prev.next = next
			}
			if next != nil {
				next.prev = prev
			}

			if list.head == cur {
				list.head = next
			}
			if list.tail == cur {
				list.tail = prev
			}

			return cur.value
		}

		cur = cur.next
		if cur == nil {
			break
		}
	}

	return nil
}
