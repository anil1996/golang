package main

import "testing"

type tests struct {
	name          string
	searchArray   []int
	searchItem    int
	expectedIndex int
}

func TestBinarySearch(t *testing.T) {
	binaryTests := []tests{
		{
			name:          "search item found",
			searchArray:   []int{1, 2, 3, 4, 9},
			searchItem:    2,
			expectedIndex: 1,
		},
		{
			name:          "search item not found",
			searchArray:   []int{1, 2, 3, 4, 9},
			searchItem:    12,
			expectedIndex: -1,
		},
	}
	for _, tc := range binaryTests {
		t.Run(tc.name, func(t *testing.T) {
			actualIndex := binarySearch(tc.searchArray, tc.searchItem)
			if actualIndex != tc.expectedIndex {
				t.Error("not equal")
			}
		})
	}
}
