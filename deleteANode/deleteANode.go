package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
	node := &SinglyLinkedListNode{
		next: nil,
		data: nodeData,
	}

	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		singlyLinkedList.tail.next = node
	}

	singlyLinkedList.tail = node
}

func printSinglyLinkedList(node *SinglyLinkedListNode, sep string, writer *bufio.Writer) {
	for node != nil {
		fmt.Fprintf(writer, "%d", node.data)

		node = node.next

		if node != nil {
			fmt.Fprintf(writer, sep)
		}
	}
}

// Delete the node at a given position in a linked list and return a reference to the head node. The head is at
// position 0. The list may be empty after you delete the node. In that case, return a null value.

// **Example**

// llist = 0 -> 1 -> 2 -> 3
// position 2

// After removing the node at position 2, llist = 0 -> 1 -> 3

/*
 * Complete the 'deleteNode' function below.
 *
 * The function is expected to return an INTEGER_SINGLY_LINKED_LIST.
 * The function accepts following parameters:
 *  1. INTEGER_SINGLY_LINKED_LIST llist
 *  2. INTEGER position
 */

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */

func deleteNode(llist *SinglyLinkedListNode, position int32) *SinglyLinkedListNode {
	if llist.next == nil && position == 0 {
		return nil
	}

	if position == 0 {
		head := llist.next

		return head
	}

	head := llist

	for i := 0; int32(i) <= position; i++ {
		previous := llist
		llist = llist.next

		if int32(i) == position-1 {
			previous.next = llist.next
		}
	}

	return head
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	llistCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	llist := SinglyLinkedList{}
	for i := 0; i < int(llistCount); i++ {
		llistItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		llistItem := int32(llistItemTemp)
		llist.insertNodeIntoSinglyLinkedList(llistItem)
	}

	positionTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	position := int32(positionTemp)

	llist1 := deleteNode(llist.head, position)

	printSinglyLinkedList(llist1, " ", writer)
	fmt.Fprintf(writer, "\n")

	writer.Flush()
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
