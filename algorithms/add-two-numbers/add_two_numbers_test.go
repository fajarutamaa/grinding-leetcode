package addtwonumbers

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name     string
		l1       []int
		l2       []int
		expected []int
	}{
		{
			name:     "example 1",
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			name:     "zero case",
			l1:       []int{0},
			l2:       []int{0},
			expected: []int{0},
		},
		{
			name:     "different length with carry",
			l1:       []int{9, 9, 9, 9, 9, 9, 9},
			l2:       []int{9, 9, 9, 9},
			expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := buildList(tt.l1)
			l2 := buildList(tt.l2)

			result := AddTwoNumbers(l1, l2)
			resultSlice := listToSlice(result)

			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, resultSlice)
			}
		})
	}
}

func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for _, n := range nums {
		current.Next = &ListNode{Val: n}
		current = current.Next
	}

	return dummy.Next
}

func listToSlice(l *ListNode) []int {
	result := []int{}
	for l != nil {
		result = append(result, l.Val)
		l = l.Next
	}
	return result
}
