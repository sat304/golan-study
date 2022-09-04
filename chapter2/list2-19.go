package main

import "fmt"

func main() {
	a := []int{10, 20, 30}
	fmt.Println(a)
	a = push(a, 1000)
	fmt.Println(a)
	a = pop(a)
	fmt.Println(a)
	a = unshift(a, 1000)
	fmt.Println(a)
	a = shift(a)
	fmt.Println(a)
	a = insert(a, 1000, 2)
	fmt.Println(a)
	a = remove(a, 2)
	fmt.Println(a)
}

func push(a []int, v int) []int {
	// 最後に追加
	return append(a, v)
}

func pop(a []int) []int {
	// 最後を除去
	return a[:len(a)-1]
}

func unshift(a []int, v int) []int {
	// 最初に追加
	return append([]int{v}, a...)
}

func shift(a []int) []int {
	//最初を削除
	return a[1:]
}

func insert(a []int, v int, p int) []int {
	// 指定の位置に追加
	a = append(a, 0)
	a = append(a[:p+1], a[p:len(a)-1]...)
	a[p] = v
	return a
}

func remove(a []int, p int) []int {
	// 指定の位置を削除
	return append(a[:p], a[p+1:]...)
}
