package main

import "fmt"

// It is New Year's Day and people are in line for the Wonderland rollercoaster ride.
// Each person wears a sticker indicating their initial position in the queue from  to .
// Any person can bribe the person directly in front of them to swap positions, but they still wear their original sticker.
// One person can bribe at most two others.

// Determine the minimum number of bribes that took place to get to a given queue order.
// Print the number of bribes, or, if anyone has bribed more than two people, print Too chaotic.

func minimumBribes(q []int32) {
	swaps := 0
	var chaotic = false

	for i := (len(q) - 1); i >= 0; i-- {
		moved := false
		atual := int(q[i])

		if atual != i+1 {
			moved = true
		}

		if moved {
			if (i-1) >= 0 && int(q[i-1]) == i+1 {
				var temp int32 = q[i-1]
				q[i-1] = q[i]
				q[i] = temp
				swaps++
			} else if (i-2) >= 0 && int(q[i-2]) == i+1 {
				q[i-2] = q[i-1]
				q[i-1] = q[i]
				q[i] = q[i-2]
				swaps += 2
			} else {
				chaotic = true
				break
			}
		}
	}

	if chaotic {
		fmt.Println("Too chaotic")
	} else {
		fmt.Println(swaps)
	}
}

func main() {

}
