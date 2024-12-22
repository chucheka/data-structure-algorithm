package linked_list

import "testing"

func TestLinkedList(t *testing.T) {

	linkedList := &LinkedList{}

	linkedList.InsertBeginning(2)
	linkedList.InsertEnd(6)
	_ = linkedList.Insert(2, 4)

	t.Run("Check LinkedList size", func(t *testing.T) {
		want := 3
		got := linkedList.size

		if got != want {
			t.Errorf("expected linked list size to be %d but got %d", want, got)
		}
	})

	t.Run("Check LinkedList head value", func(t *testing.T) {
		want := 2
		got := linkedList.head.data

		if got != want {
			t.Errorf("expected first value in linked list to be %d but got %d", want, got)
		}
	})

	t.Run("Delete the first element of linked list", func(t *testing.T) {
		want := 2
		got, _ := linkedList.DeleteFirst() // Type assertion i.(T)

		if got != want {
			t.Errorf("expected deleted value from the linked list to be %d but got %d", want, got)
		}
	})
}
