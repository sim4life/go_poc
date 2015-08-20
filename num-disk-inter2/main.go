package main

import (
	"fmt"
	"sort"
)

type int64arr []int64

func (a int64arr) Len() int           { return len(a) }
func (a int64arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a int64arr) Less(i, j int) bool { return a[i] < a[j] }

func Solution(A []int) int {
	// write your code in Go 1.4
	l := len(A)
	A1 := make([]int64, len(A))
	A2 := make([]int64, len(A))

	for ind, val := range A {
		A1[ind] = int64(val + ind)
		A2[ind] = int64(-(val - ind))
	}

	fmt.Println("A1 b4 sorting:   ", A1)
	fmt.Println("A2 b4 sorting:   ", A2)
	sort.Sort(int64arr(A1))
	sort.Sort(int64arr(A2))
	fmt.Println("A1 a3 sorting:   ", A1)
	fmt.Println("A2 a3 sorting:   ", A2)

	var count int64 = 0
	pos := -1

	for i := l - 1; i >= 0; i-- {
		pos = sort.Search(len(A2), func(pos int) bool { return A2[pos] >= A1[i] })
		if pos < len(A2) && A2[pos] == A1[i] {
			for pos < l && A2[pos] == A1[i] {
				pos++
			}
			count += int64(pos)
		} else {
			count += int64(pos)
		}
	}

	var sub int64 = (int64(l) * int64(l+1)) / 2
	count = count - sub

	if count > 1e7 {
		return -1
	}

	return int(count)
}

func main() {
	fmt.Println("The Solution is: ", Solution([]int{1, 5, 2, 1, 4, 0}))
}
