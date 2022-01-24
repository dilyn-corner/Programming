/*
Package tree is an implementation of a tree building algorithm.  It defines a
Record and a Node, the Record containing simple information about any give node,
the Node containing relevant information about the structure of our tree at any
given vertex.  This tree is built out of pointers to various children, and then
the list of vertices is sorted.
*/
package tree

import (
	"errors"
	"sort"
)

// Define the Record type
type Record struct {
	ID, Parent int
}

// Define the Node type
type Node struct {
	ID       int
	Children []*Node
}

// We define nodePart to pass a valid type to the following three functions
type nodePart []*Node

// These three functions only exist to define what the sort.Interface looks like
// on our incredibly odd structs. Without these, sorting will fail (if we opt to
// use "sort" instead of... implementing our own sorting algorithm...
// sort.Sort() requires Len(), Swap(i, j int), and Less(i, j int) bool
func (node nodePart) Len() int {
	return len(node)
}
func (node nodePart) Swap(i, j int) {
	node[i], node[j] = node[j], node[i]
}
func (node nodePart) Less(i, j int) bool {
	return node[i].ID < node[j].ID
}

// Build takes an array of records and sorts them into a proper tree
func Build(records []Record) (*Node, error) {

	if len(records) <= 0 {
		return nil, nil
	}

	parentNode := make([]*Node, len(records)) // Why does this have to be a pointer...
	childNode := make([]Node, len(records))   // and this doesn't...
	alreadyKnown := make([]bool, len(records))

	// error handling requires ranging _, err := range records { ... }
	for _, record := range records {

		// We can't have more IDs than we have records
		if record.ID >= len(records) {
			return nil, errors.New("record list is missing a node")
		}

		// We can't have a ID with Parent <= ID UNLESS ID == 0
		if record.ID != 0 && record.ID <= record.Parent {
			return nil, errors.New("a nonroot record has self or child as own ancestor")
		}

		// We can't have duplicate IDs
		if alreadyKnown[record.ID] {
			return nil, errors.New("can't have the same ID multiple times")
		}

		// Track if we are aware of any given node
		alreadyKnown[record.ID] = true

		// Populate our parentNode at any given ID
		// The record.Parent of any given parentNode[i] is
		// the childNode[j]'s record.Parent; don't copy this information into
		// the array; use a pointer. When in doubt, use a pointer.
		if record.ID != 0 {
			parentNode[record.ID] = &childNode[record.Parent]
		} else if record.Parent != 0 {
			// A kill switch on a malformed root record
			return nil, errors.New("root record has nonzero Parent")
		}
	}

	// Iterate over the records, appending child nodes to our list of children
	// It's important that we start at not the first record! What does
	// childNode[0] even mean!
	for i := 1; i < len(records); i++ {
		parentNode[i].Children = append(parentNode[i].Children, &childNode[i])
	}

	// We have now a Node which contains the arbitrarily organized children
	// nodes of any given record. We have to now sort this node to generate a
	// proper, top-down tree.

	for i, node := range childNode {
		childNode[i].ID = i // for node i, set its ID to i
		sort.Sort(nodePart(node.Children))
	}
	return &childNode[0], nil // childNode[0] == whole tree
}

// https://ieftimov.com/post/golang-datastructures-trees/
