# Golang Utility Package

This utility package provides basic functions for finding duplicates in a slice of generic items and finding missing items in two slices. This will be extended.

## Functions

- `FindDuplicates(items interface{}) []interface{}`: Finds and returns duplicates in a given slice of generic items.
- `FindMissingItems(sliceA, sliceB interface{}) []interface{}`: Finds and returns the missing items in two given slices of generic items.

## Usage

To use the utility package in your Go project, first import it:

```go
import "github.com/golang-utility"

```

Then you can use it below

```
intSlice := []int{1, 2, 3, 4, 3, 5, 1, 6}
duplicates := utility.FindDuplicates(intSlice)
fmt.Printf("Duplicates: %v\n", duplicates)

stringSliceA := []string{"apple", "banana", "cherry", "date", "fig"}

stringSliceB := []string{"apple", "banana", "date", "fig"}

missingItems := utility.FindMissingItems(stringSliceA, stringSliceB)

fmt.Printf("Missing items: %v\n", missingItems)
```
