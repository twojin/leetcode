package main

import (
	"fmt"
	"strconv"
)

var MinInt = -1 << 31
var MaxInt = 1<<31 - 1

func main() {
	head := new(ListNode)
	n1 := new(ListNode)
	n2 := new(ListNode)

	head.Val = 3
	n1.Val = 4
	n2.Val = 5

	head.Next = n1
	n1.Next = n2

	ShowAll(head)
	h := reverseList(head)
	ReverseShowAll(h)
	fmt.Println("--------------------------------")
}

func foo(n int) {
	if n <= 0 {
		return
	}

	n--
	fmt.Println(n)
	foo(n)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var i int
	for ; i < len(strs[0]); i++ {
		for j := 0; j < len(strs); j++ {
			if strs[j] == "" {
				return ""
			}

			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0][:i]
}

func countAndSay(n int) string {
	var say string = "1"
	if n <= 1 {
		return say
	}

	for n > 1 {
		var newStr string
		for i := 0; i < len(say); {
			count := 1
			num := string(say[i])
			j := i + 1
			for ; j < len(say) && say[i] == say[j]; j++ {
				count++
			}
			newStr = newStr + strconv.Itoa(count) + num
			i = j
		}
		say = newStr
		n--
	}

	return say
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre, cur := head, head
	for n > 0 { // 使用双指针 一个指针先行n步
		fmt.Println(pre.Val)
		pre = pre.Next
		n--
	}

	if pre == nil { // 删除的是链表的头部
		return head.Next
	}

	for pre.Next != nil { // 当先行指针下一个节点为空时 当前指指针的下一个节点就是倒数第n个节点
		cur = cur.Next
		pre = pre.Next
	}

	cur.Next = cur.Next.Next
	return head
}

func ShowAll(head *ListNode) {
	if head == nil {
		fmt.Printf("\n")
		return
	}
	fmt.Printf("%d", head.Val)
	ShowAll(head.Next)
}

func ReverseShowAll(head *ListNode) *ListNode {
	if head.Next == nil {
		fmt.Printf("%d \n", head.Val)
		return head
	}

	newHead := ReverseShowAll(head.Next)
	fmt.Printf("%d \n", newHead.Val)
	return newHead
}

// 翻转链表
func reverseList(head *ListNode) *ListNode {
	p := head.Next

	// 类似头插法
	for p.Next != nil {
		q := p.Next
		p.Next = q.Next
		q.Next = head.Next
		head.Next = q
	}

	p.Next = head // 形成环
	head = head.Next
	p.Next.Next = nil // 断掉环

	return head
}
