package hackerrank

import "fmt"

/*
	1234
	1243
	1324
	1342
	2134
	2143
	2314
	2341
	2413
	2431
	3214

*/
func Permutations(pivot int, perm []int32, end int) {
	if pivot == end {
		fmt.Println("perm:", perm)
		return
	}
	for i := pivot; i < end; i++ {
		perm[i], perm[pivot] = perm[pivot], perm[i]
		Permutations(pivot+1, perm, end)
		perm[i], perm[pivot] = perm[pivot], perm[i]
	}

}
