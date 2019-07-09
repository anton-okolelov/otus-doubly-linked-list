package doubly_linked_list

import "testing"
import "github.com/stretchr/testify/assert"

func TestList_PushFront(t *testing.T) {
	list := List{}
	list.PushFront(1)
	list.PushFront(2)
	list.PushFront(3)

	assert.Equal(t, 3, list.First().Value())
	assert.Equal(t, 1, list.Last().Value())
}

func TestList_PushBack(t *testing.T) {
	list := List{}
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	assert.Equal(t, 1, list.First().Value())
	assert.Equal(t, 3, list.Last().Value())
}

func TestList_EmptyList(t *testing.T) {
	list := List{}
	assert.Nil(t, list.First())
	assert.Nil(t, list.Last())
}

func TestList_Len(t *testing.T) {
	list := List{}
	assert.Equal(t, 0, list.Len())
	list.PushFront(1)
	list.PushFront(1)
	assert.Equal(t, 2, list.Len())
}

func TestItem_Prev(t *testing.T) {
	list := List{}
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	assert.Equal(t, 2, list.Last().Prev().Value())
}

func TestItem_Remove(t *testing.T) {
	list := List{}
	list.PushBack(1)
	list.First().Remove()
	assert.Equal(t, 0, list.Len())

	list.PushBack(1)
	list.PushBack(2)
	list.Last().Remove()
	assert.Equal(t, 1, list.Last().value)

	list.PushBack(2)
	list.PushBack(3)
	list.First().Next().Remove()
	assert.Equal(t, 1, list.First().Value())
	assert.Equal(t, 3, list.Last().Value())
	assert.Equal(t, 2, list.Len())

}
