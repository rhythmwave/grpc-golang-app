package utils

import "sort"

// LessFunc defines the comparison function for sorting elements.
type LessFunc func(i, j interface{}) bool

// Sorter is an interface that any type can implement to be sortable.
type Sorter interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
}

// multiSorter implements the Sorter interface, sorting elements within.
type multiSorter struct {
	data      []interface{}
	lessFuncs []LessFunc
}

// Sort sorts the data according to the provided LessFunc functions.
func (ms *multiSorter) Sort(data []interface{}) {
	ms.data = data // Type assertion to slice of interface{}
	sort.Sort(ms)
}

// OrderedBy returns a Sorter that sorts using the provided LessFunc functions.
func OrderedBy(lessFuncs ...LessFunc) *multiSorter {
	return &multiSorter{
		lessFuncs: lessFuncs,
	}
}

// Len is part of the Sorter interface.
func (ms *multiSorter) Len() int {
	return len(ms.data)
}

// Swap is part of the Sorter interface.
func (ms *multiSorter) Swap(i, j int) {
	ms.data[i], ms.data[j] = ms.data[j], ms.data[i]
}

// Less is part of the Sorter interface.
func (ms *multiSorter) Less(i, j int) bool {
	k := 0
	// Try all but the last comparison function.
	for k = 0; k < len(ms.lessFuncs)-1; k++ {
		less := ms.lessFuncs[k]
		switch less(ms.data[i], ms.data[j]) {
		case true:
			return true
		case false:
			return false
		}
	}
	// All comparisons to here said "equal", so use the final comparison.
	return ms.lessFuncs[k](ms.data[i], ms.data[j])
}
