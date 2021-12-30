package main

import "testing"

type searchTest struct {
	name         string
	inputArr     []string
	searchItem   string
	actualResult bool
}

func TestSearching(t *testing.T) {
	tests := []searchTest{
		{
			name:         "successful search",
			inputArr:     []string{"a", "b", "cd", "ef"},
			searchItem:   "cd",
			actualResult: true,
		},
		{
			name:         "unsuccessful search",
			inputArr:     []string{"a", "b", "cd", "ef"},
			searchItem:   "gh",
			actualResult: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			expectedResult := linearSearch(tt.inputArr, tt.searchItem)
			if expectedResult != tt.actualResult {
				t.Error("Not equal")
			}

		})
	}
}
