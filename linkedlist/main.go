package main

import "fmt"

// Start with the basics, then practice problems 1–10 in order.
// Try writing each function yourself before reading its solution.
// Run all examples with: go run ./linkedlist
// Unless stated otherwise, these functions expect a list without a cycle.

// A node stores a value and a pointer to the next node.
// head -> [10 | next] -> [20 | nil]; nil means the list has ended.
// An empty list has head == nil. *ListNode is a pointer; &ListNode creates a node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Basic question: Build a list from values. O(n) time, O(n) new nodes.
func buildList(values ...int) *ListNode {
	var head, tail *ListNode
	for _, value := range values {
		node := &ListNode{Val: value}
		if head == nil {
			head = node // The first node becomes the head.
		} else {
			tail.Next = node
		}
		tail = node
	}
	return head
}

// Basic question: Traverse a list and collect its values for printing.
// O(n) time and O(n) space. Do not call this on a cyclic list.
func toSlice(head *ListNode) []int {
	values := []int{}
	for current := head; current != nil; current = current.Next {
		values = append(values, current.Val)
	}
	return values
}

// Basic question: Find the first node with a given value. O(n) time, O(1) space.
func search(head *ListNode, value int) *ListNode {
	for current := head; current != nil; current = current.Next {
		if current.Val == value {
			return current
		}
	}
	return nil
}

// Basic question: Insert at the beginning. O(1) time and extra space.
// Save the returned head: head = insertFront(head, 10).
func insertFront(head *ListNode, value int) *ListNode {
	return &ListNode{Val: value, Next: head}
}

// Basic question: Insert at the end. O(n) time, O(1) extra space.
func insertEnd(head *ListNode, value int) *ListNode {
	node := &ListNode{Val: value}
	if head == nil {
		return node
	}
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = node
	return head
}

// Basic question: Insert at a zero-based index (0 through length inclusive).
// Invalid indices leave the list unchanged. O(n) time, O(1) extra space.
func insertAt(head *ListNode, index, value int) *ListNode {
	if index < 0 {
		return head
	}
	if index == 0 {
		return insertFront(head, value)
	}
	if head == nil {
		return head
	}
	previous := head
	for i := 0; i < index-1; i++ {
		if previous.Next == nil {
			return head
		}
		previous = previous.Next
	}
	// Connect the new node to the rest BEFORE changing the previous node's link.
	previous.Next = &ListNode{Val: value, Next: previous.Next}
	return head
}

// Basic question: Delete at a zero-based index, including the head or tail.
// Invalid indices leave the list unchanged. O(n) time, O(1) extra space.
func deleteAt(head *ListNode, index int) *ListNode {
	if index < 0 || head == nil {
		return head
	}
	if index == 0 {
		return head.Next // Deleting the head changes where the list starts.
	}
	previous := head
	for i := 0; i < index-1; i++ {
		if previous.Next == nil {
			return head
		}
		previous = previous.Next
	}
	if previous.Next != nil {
		previous.Next = previous.Next.Next // Bypass the node being deleted.
	}
	return head
}

// Basic question: Delete the FIRST node with a given value.
// A missing value leaves the list unchanged. O(n) time, O(1) extra space.
func deleteValue(head *ListNode, value int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == value {
		return head.Next
	}
	previous := head
	for previous.Next != nil {
		if previous.Next.Val == value {
			previous.Next = previous.Next.Next
			break
		}
		previous = previous.Next
	}
	return head
}

// 1. 876. Middle of the Linked List — Easy
// Return the middle node; for an even length, return the SECOND middle.
// https://leetcode.com/problems/middle-of-the-linked-list/
// Slow moves one step, fast moves two. O(n) time, O(1) space.
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 2. 206. Reverse Linked List — Easy
// Reverse all links and return the new head. Try iteration before recursion.
// https://leetcode.com/problems/reverse-linked-list/
// Iterative: O(n) time, O(1) space.
func reverseList(head *ListNode) *ListNode {
	var previous *ListNode
	current := head
	for current != nil {
		next := current.Next // Save the rest so we don't lose it.
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}

// Recursive: O(n) time, O(n) call stack space.
func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListRecursive(head.Next)
	head.Next.Next = head // The next node now points back to this node.
	head.Next = nil       // Remove the old forward link to prevent a cycle.
	return newHead
}

// 3. 21. Merge Two Sorted Lists — Easy
// Connect nodes from two ascending, separate lists into one sorted list.
// https://leetcode.com/problems/merge-two-sorted-lists/
// Reuses input nodes. O(m+n) time, O(1) extra space.
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	// Choose the first node, then attach the remaining nodes after it.
	var head *ListNode
	if list1.Val <= list2.Val {
		head = list1
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}
	tail := head
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}
	return head
}

// 4. 141. Linked List Cycle — Easy
// Determine whether following Next can revisit a node.
// https://leetcode.com/problems/linked-list-cycle/
// Floyd's algorithm: in a cycle, fast eventually catches slow.
// O(n) time, O(1) space. Compare pointers, not values.
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// 5. 160. Intersection of Two Linked Lists — Easy
// Return the first SHARED NODE, or nil. Equal values do not mean intersection.
// https://leetcode.com/problems/intersection-of-two-linked-lists/
// Switching heads gives both pointers the same total distance to travel.
// Inputs must have no cycles. O(m+n) time, O(1) space.
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}

// 6. 234. Palindrome Linked List — Easy
// Check whether the values read the same forward and backward.
// https://leetcode.com/problems/palindrome-linked-list/
// Find the first half's end, reverse the second half, compare, then restore.
// O(n) time, O(1) extra space. The original list is preserved.
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	secondHalf := reverseList(slow.Next)
	left, right := head, secondHalf
	result := true
	for right != nil {
		if left.Val != right.Val {
			result = false
			break
		}
		left, right = left.Next, right.Next
	}
	slow.Next = reverseList(secondHalf)
	return result
}

// 7. 19. Remove Nth Node From End — Medium
// Remove the nth node counting from the end (n=1 removes the tail).
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// Move fast n steps first. Handle removing the head, then move both pointers.
// Invalid n leaves the list unchanged. O(length) time, O(1) space.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 || head == nil {
		return head
	}
	slow, fast := head, head
	for i := 0; i < n; i++ {
		if fast == nil {
			return head
		}
		fast = fast.Next
	}
	if fast == nil {
		return head.Next // n equals the length, so remove the head.
	}
	// Stop with slow just before the node to remove.
	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}
	slow.Next = slow.Next.Next
	return head
}

// 8. 24. Swap Nodes in Pairs — Medium
// Swap adjacent nodes without changing their values; leave an odd tail alone.
// https://leetcode.com/problems/swap-nodes-in-pairs/
// O(n) time, O(1) space.
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// Swap the first pair and update the head explicitly.
	first := head
	second := head.Next
	first.Next = second.Next
	second.Next = first
	head = second
	previous := first

	// The remaining pairs have a previous node to reconnect.
	for previous.Next != nil && previous.Next.Next != nil {
		first := previous.Next
		second := first.Next
		// previous -> first -> second -> rest becomes previous -> second -> first -> rest.
		first.Next = second.Next
		second.Next = first
		previous.Next = second
		previous = first
	}
	return head
}

// 9. 328. Odd Even Linked List — Medium
// Group odd POSITION nodes before even position nodes, preserving order.
// Positions start at 1; node values do not determine the groups.
// https://leetcode.com/problems/odd-even-linked-list/
// O(n) time, O(1) space.
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd, even := head, head.Next
	evenHead := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

// 10. 2. Add Two Numbers — Medium
// Digits are stored in reverse order: [2,4,3] represents 342.
// Each node contains a digit 0–9; inputs represent nonnegative integers.
// https://leetcode.com/problems/add-two-numbers/
// O(max(m,n)) time and output space; O(1) auxiliary space.
func addTwoNumbers(list1, list2 *ListNode) *ListNode {
	var head, tail *ListNode
	carry := 0
	for list1 != nil || list2 != nil || carry != 0 {
		sum := carry
		if list1 != nil {
			sum += list1.Val
			list1 = list1.Next
		}
		if list2 != nil {
			sum += list2.Val
			list2 = list2.Next
		}
		node := &ListNode{Val: sum % 10}
		if head == nil {
			head = node // The first result digit becomes the head.
		} else {
			tail.Next = node
		}
		tail = node
		carry = sum / 10
	}
	return head
}

func main() {
	fmt.Println("Basics (indices start at 0):")
	head := buildList(10, 20, 30)
	fmt.Println("Build:", toSlice(head)) // [10 20 30]
	head = insertFront(head, 5)
	fmt.Println("Insert front:", toSlice(head)) // [5 10 20 30]
	head = insertEnd(head, 40)
	fmt.Println("Insert end:", toSlice(head)) // [5 10 20 30 40]
	head = insertAt(head, 2, 15)
	fmt.Println("Insert 15 at index 2:", toSlice(head)) // [5 10 15 20 30 40]
	fmt.Println("Search 20:", search(head, 20) != nil)  // true
	head = deleteAt(head, 0)
	fmt.Println("Delete head:", toSlice(head)) // [10 15 20 30 40]
	head = deleteAt(head, 4)
	fmt.Println("Delete tail:", toSlice(head)) // [10 15 20 30]
	head = deleteValue(head, 20)
	fmt.Println("Delete value 20:", toSlice(head)) // [10 15 30]

	// Use fresh lists because many solutions change the links in their inputs.
	fmt.Println("\nPractice problems:")
	fmt.Println("1. Middle (even length):", toSlice(middleNode(buildList(1, 2, 3, 4, 5, 6)))) // [4 5 6]
	fmt.Println("2. Reverse iterative:", toSlice(reverseList(buildList(1, 2, 3))))            // [3 2 1]
	fmt.Println("2. Reverse recursive:", toSlice(reverseListRecursive(buildList(1, 2, 3))))   // [3 2 1]
	fmt.Println("3. Merge:", toSlice(mergeTwoLists(buildList(1, 2, 4), buildList(1, 3, 4))))  // [1 1 2 3 4 4]

	cycle := buildList(3, 2, 0, -4)
	cycle.Next.Next.Next.Next = cycle.Next                 // Tail points back to the node containing 2.
	fmt.Println("4. Cycle:", hasCycle(cycle))              // true; never pass cycle to toSlice.
	fmt.Println("4. No cycle:", hasCycle(buildList(1, 2))) // false

	shared := buildList(8, 4, 5)
	listA := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: shared}}
	listB := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: shared}}}
	fmt.Println("5. Intersection:", toSlice(getIntersectionNode(listA, listB)))                            // [8 4 5]
	fmt.Println("5. Same values, separate nodes:", getIntersectionNode(buildList(8), buildList(8)) == nil) // true

	palindrome := buildList(1, 2, 2, 1)
	fmt.Println("6. Palindrome:", isPalindrome(palindrome))                                           // true
	fmt.Println("   List after checking:", toSlice(palindrome))                                       // [1 2 2 1]
	fmt.Println("6. Not a palindrome:", isPalindrome(buildList(1, 2)))                                // false
	fmt.Println("7. Remove 2nd from end:", toSlice(removeNthFromEnd(buildList(1, 2, 3, 4, 5), 2)))    // [1 2 3 5]
	fmt.Println("8. Swap pairs:", toSlice(swapPairs(buildList(1, 2, 3, 4, 5))))                       // [2 1 4 3 5]
	fmt.Println("9. Odd even:", toSlice(oddEvenList(buildList(1, 2, 3, 4, 5))))                       // [1 3 5 2 4]
	fmt.Println("10. Add 342 + 465:", toSlice(addTwoNumbers(buildList(2, 4, 3), buildList(5, 6, 4)))) // [7 0 8]
	fmt.Println("10. Add 99 + 1:", toSlice(addTwoNumbers(buildList(9, 9), buildList(1))))             // [0 0 1]

	// Extra practice: empty lists, one node, invalid indices, and a missing value.
	fmt.Println("\nEmpty list deletion:", toSlice(deleteAt(nil, 0)))             // []
	fmt.Println("Remove only node:", toSlice(removeNthFromEnd(buildList(7), 1))) // []
	fmt.Println("Invalid insertion:", toSlice(insertAt(buildList(1), 3, 9)))     // [1]
	fmt.Println("Missing value:", toSlice(deleteValue(buildList(1), 9)))         // [1]
}
