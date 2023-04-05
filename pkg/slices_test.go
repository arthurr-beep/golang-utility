package utility

import (
	"reflect"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected []interface{}
	}{
		{"Integers", []int{1, 2, 3, 4, 3, 5, 1, 6}, []interface{}{3, 1}},
		{"Strings", []string{"apple", "banana", "cherry", "apple", "banana", "date"}, []interface{}{"apple", "banana"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FindDuplicates(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, got: %v", tc.expected, result)
			}
		})
	}
}

func TestFindMissingItems(t *testing.T) {
	testCases := []struct {
		name     string
		sliceA   interface{}
		sliceB   interface{}
		expected []interface{}
	}{
		{"Integers", []int{1, 2, 3, 4, 5}, []int{1, 3, 5}, []interface{}{2, 4}},
		{"Strings", []string{"apple", "banana", "cherry", "date", "fig"}, []string{"apple", "banana", "date", "fig"}, []interface{}{"cherry"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FindMissingItems(tc.sliceA, tc.sliceB)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, got: %v", tc.expected, result)
			}
		})
	}
}
