package main

// Given a 6x6 2D Array, arr :

// 1 1 1 0 0 0
// 0 1 0 0 0 0
// 1 1 1 0 0 0
// 0 0 0 0 0 0
// 0 0 0 0 0 0
// 0 0 0 0 0 0

// An hourglass in A is a subset of values with indices falling in this pattern in arr's graphical representation:

// a b c
//   d
// e f g

// There are 16 hourglasses in arr. An hourglass sum is the sum of an hourglass' values. Calculate the
// hourglass sum for every hourglass in arr, then print the maximum hourglass sum. The array will always be
// 6x6.

func hourglassSum(arr [][]int32) int32 {
	var maxSum int32 = -99999

	for i := 0; i < len(arr)-2; i++ {
		var sum int32

		for y := 0; y < len(arr[i])-2; y++ {
			sum = arr[i][y] + arr[i][y+1] + arr[i][y+2]
			sum += arr[i+1][y+1]
			sum += arr[i+2][y] + arr[i+2][y+1] + arr[i+2][y+2]

			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum
}

func main() {

}
