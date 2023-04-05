package utility

import (
	"reflect"
)

// FindDuplicates finds and returns duplicates in a given slice of generic items.
// O(n) complexity
func FindDuplicates(items interface{}) []interface{} {
	slice := reflect.ValueOf(items)

	if slice.Kind() != reflect.Slice {
		panic("FindDuplicates: items must be a slice")
	}

	duplicates := make([]interface{}, 0)
	seen := make(map[interface{}]bool)

	for i := 0; i < slice.Len(); i++ {
		item := slice.Index(i).Interface()
		if _, ok := seen[item]; ok {
			duplicates = append(duplicates, item)
		} else {
			seen[item] = true
		}
	}

	return duplicates
}

// FindMissingItems finds and returns the missing items in two given slices of generic items.
// Given SliceA and SliceB - this method returns items in A but not in B
// O(n + m) complexity
func FindMissingItems(sliceA, sliceB interface{}) []interface{} {
	slice1 := reflect.ValueOf(sliceA)
	slice2 := reflect.ValueOf(sliceB)

	if slice1.Kind() != reflect.Slice || slice2.Kind() != reflect.Slice {
		panic("FindMissingItems: both inputs must be slices")
	}

	missing := make([]interface{}, 0)
	itemsMap := make(map[interface{}]bool)

	for i := 0; i < slice2.Len(); i++ {
		itemsMap[slice2.Index(i).Interface()] = true
	}

	for i := 0; i < slice1.Len(); i++ {
		item := slice1.Index(i).Interface()
		if _, ok := itemsMap[item]; !ok {
			missing = append(missing, item)
		}
	}

	return missing
}
