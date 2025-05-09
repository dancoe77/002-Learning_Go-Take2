package main

import "testing"

func TestGetUser(t *testing.T) {
	md := &MockDatastore{
		Users: map[int]User{
			2: {ID: 2, First: "Jenny"},
		},
	}
	s := &Service{
		ds: md,
	}

	u, err := s.GetUser(2)
	if err != nil {
		t.Errorf("got error: %v", err)
	}

	if u.First != "Jenny" {
		t.Errorf("got: %s, want: %s", u.First, "Jenny")
	}
}

/*
Pointers are used in method receivers in Go for two main reasons:

1. Efficiency: When a large struct is passed by value to a method (i.e., without using a pointer), the entire struct must be copied, which can be inefficient in terms of memory and performance. When a pointer to the struct is passed instead, only the memory address of the struct is passed, which is typically much more efficient.

2. Modification: If you want your method to modify the state of the receiver, the receiver must be a pointer. This is because Go is a pass-by-value language, meaning that when a value is passed to a function, a copy of that value is made in the function's scope. If the function modifies the copy, the original value is not affected. But if a pointer to the value is passed, the function can modify the original value.

In the case of your code, both `Service` and `MockDatastore` methods are defined with pointer receivers because the methods may need to modify the state of the receiver. For example, `MockDatastore.SaveUser` needs to add a user to the `Users` map of the `MockDatastore`. If `md` were not a pointer, `SaveUser` would receive a copy of the `MockDatastore`, and changes would only be made to that copy, not the original `MockDatastore`.

```go
func (md *MockDatastore) SaveUser(u User) error {
    _, ok := md.Users[u.ID]
    if ok {
        return fmt.Errorf("User %v already exists", u.ID)
    }
    md.Users[u.ID] = u
    return nil
}
```

Also, note that even when methods only read the receiver's data, you might still want to use a pointer receiver for efficiency reasons, especially if the struct is large.

*/

/*
A map is just a hash table.
The data is arranged into an array of buckets.
...
When the hashtable grows, we allocate a new array
of buckets twice as big. Buckets are incrementally
copied from the old bucket array to the new bucket array.
*/
// https://cs.opensource.google/go/go/+/refs/tags/go1.20.4:src/runtime/map.go
/*
// A header for a Go map.
type hmap struct {
	// Note: the format of the hmap is also encoded in cmd/compile/internal/reflectdata/reflect.go.
	// Make sure this stays in sync with the compiler's definition.
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	extra *mapextra // optional fields
}
*/
