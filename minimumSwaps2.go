package main

// You are given an unordered array consisting of consecutive integers  [1, 2, 3, ..., n] without any duplicates.
// You are allowed to swap any two elements. Find the minimum number of swaps required to sort the array in
// ascending order.

// **Example**

// Perform the following steps:

// i   arr                         swap (indices)
// 0   [7, 1, 3, 2, 4, 5, 6]   swap (0,3)
// 1   [2, 1, 3, 7, 4, 5, 6]   swap (0,1)
// 2   [1, 2, 3, 7, 4, 5, 6]   swap (3,4)
// 3   [1, 2, 3, 4, 7, 5, 6]   swap (4,5)
// 4   [1, 2, 3, 4, 5, 7, 6]   swap (5,6)
// 5   [1, 2, 3, 4, 5, 6, 7]

// It took  swaps to sort the array.

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	var swaps int32

	for i := 0; i < len(arr)-1; i++ {
		current := arr[i]
		rightPosition := int32(i + 1)

		if current == rightPosition {
			continue
		}

		t := i
		for {
			t++
			searchRightValue := arr[t]

			if searchRightValue == rightPosition {
				break
			}
		}

		rightValue := arr[t]
		arr[t] = current
		arr[i] = rightValue

		swaps++
	}

	return swaps
}
