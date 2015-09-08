// Go's `sort` package implements sorting for builtins
// and user-defined types. We'll look at sorting for
// builtins first.

package main

import (
	"fmt"
	"sort"
)

func main() {

	// Sort methods are specific to the builtin type;
	// here's an example for strings. Note that sorting is
	// in-place, so it changes the given slice and doesn't
	// return a new one.
	strs := []string{"can", "abs", "bus", "calli"}
	sort.Sort(sort.Reverse(sort.StringSlice(strs)))
	fmt.Println("Strings:", strs)
	fmt.Println("StringsAreSorted: ", sort.StringsAreSorted(strs))
	pos := sort.SearchStrings(strs, "bus") //doesn't work
	fmt.Printf("pos of \"bus\" in descending sorting is: %d\n", pos)
	pos = sort.SearchStrings(strs, "truck") //doesn't work
	fmt.Printf("pos of \"truck\" in descending sorting is: %d\n", pos)
	sort.Strings(strs)
	fmt.Println("Strings:", strs)
	fmt.Println("StringsAreSorted: ", sort.StringsAreSorted(strs))
	pos = sort.SearchStrings(strs, "bus")
	fmt.Printf("pos of \"bus\" in ascending sorting is: %d\n", pos)
	pos = sort.SearchStrings(strs, "truck")
	fmt.Printf("pos of \"truck\" in ascending sorting is: %d\n", pos)
	pos = sort.SearchStrings(strs, "call")
	fmt.Printf("pos of \"call\" in ascending sorting is: %d\n", pos)

	// An example of sorting `int`s.
	ints := []int{7, 2, 4}
	x := 4
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)
	pos = sort.SearchInts(ints, x)
	fmt.Printf("pos of %d in slice is: %d\n", x, pos)
	// We can also use `sort` to check if a slice is
	// already in sorted order.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
