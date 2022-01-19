package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Consider the following version of Bubble Sort:

// for (int i = 0; i < n; i++) {

//     for (int j = 0; j < n - 1; j++) {
//         // Swap adjacent elements if they are in decreasing order
//         if (a[j] > a[j + 1]) {
//             swap(a[j], a[j + 1]);
//         }
//     }

// }

/*
 * Complete the 'countSwaps' function below.
 *
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func countSwaps(a []int32) {
	n := len(a)
	swaps := 0

	for i := 0; i < n; i++ {

		for j := 0; j < n-1; j++ {
			// Swap adjacent elements if they are in decreasing order
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp

				swaps++
			}
		}

	}

	fmt.Printf("Array is sorted in %d swaps. \n", swaps)
	fmt.Println("First Element:", a[0])
	fmt.Println("Last Element:", a[n-1])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	countSwaps(a)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
