// Package addrset provides a specialized set data structure for managing unique Gno addresses.
//
// It is built on top of an AVL tree for efficient operations and maintains addresses in sorted order.
// This package is particularly useful when you need to:
//   - Track a collection of unique addresses (e.g., for whitelists, participants, etc.)
//   - Efficiently check address membership
//   - Support pagination when displaying addresses
//
// Example usage:
//
//	import (
//	    "std"
//	    "gno.land/p/moul/addrset"
//	)
//
//	func MyHandler() {
//	    // Create a new address set
//	    var set addrset.Set
//
//	    // Add some addresses
//	    addr1 := std.Address("g1jg8mtutu9khhfwc4nxmuhcpftf0pajdhfvsqf5")
//	    addr2 := std.Address("g1sss5g0rkqr88k4u648yd5d3l9t4d8vvqwszqth")
//
//	    set.Add(addr1)  // returns true (newly added)
//	    set.Add(addr2)  // returns true (newly added)
//	    set.Add(addr1)  // returns false (already exists)
//
//	    // Check membership
//	    if set.Has(addr1) {
//	        // addr1 is in the set
//	    }
//
//	    // Get size
//	    size := set.Size()  // returns 2
//
//	    // Iterate with pagination (10 items per page, starting at offset 0)
//	    set.IterateByOffset(0, 10, func(addr std.Address) bool {
//	        // Process addr
//	        return false  // continue iteration
//	    })
//
//	    // Remove an address
//	    set.Remove(addr1)  // returns true (was present)
//	    set.Remove(addr1)  // returns false (not present)
//	}
package addrset

import (
	"std"

	"gno.land/p/demo/avl"
)

type Set struct {
	tree avl.Tree
}

// Add inserts an address into the set.
// Returns true if the address was newly added, false if it already existed.
func (s *Set) Add(addr std.Address) bool {
	return !s.tree.Set(string(addr), nil)
}

// Remove deletes an address from the set.
// Returns true if the address was found and removed, false if it didn't exist.
func (s *Set) Remove(addr std.Address) bool {
	_, removed := s.tree.Remove(string(addr))
	return removed
}

// Has checks if an address exists in the set.
func (s *Set) Has(addr std.Address) bool {
	return s.tree.Has(string(addr))
}

// Size returns the number of addresses in the set.
func (s *Set) Size() int {
	return s.tree.Size()
}

// IterateByOffset walks through addresses starting at the given offset.
// The callback should return true to stop iteration.
func (s *Set) IterateByOffset(offset int, count int, cb func(addr std.Address) bool) {
	s.tree.IterateByOffset(offset, count, func(key string, _ any) bool {
		return cb(std.Address(key))
	})
}

// ReverseIterateByOffset walks through addresses in reverse order starting at the given offset.
// The callback should return true to stop iteration.
func (s *Set) ReverseIterateByOffset(offset int, count int, cb func(addr std.Address) bool) {
	s.tree.ReverseIterateByOffset(offset, count, func(key string, _ any) bool {
		return cb(std.Address(key))
	})
}

// Tree returns the underlying AVL tree for advanced usage.
func (s *Set) Tree() avl.ITree {
	return &s.tree
}
