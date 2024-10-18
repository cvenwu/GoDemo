/*
 * @Author: cvenwu yirufeng@foxmail.com
 * @Date: 2024-10-17 09:47:06
 * @LastEditors: cvenwu yirufeng@foxmail.com
 * @LastEditTime: 2024-10-17 10:03:23
 * @Description:
 *
 * Copyright (c) 2024 by yirufeng@foxmail.com, All Rights Reserved.
 */
package main

import "fmt"

func readOnlyChan(ch <-chan string) string {
	return <-ch
}

func writeOnlyChan(ch chan<- string, msg string) {
	ch <- msg
}

func main() {
	fmt.Println("234")
	ch := make(chan string, 2)
	writeOnlyChan(ch, "msg")
	writeOnlyChan(ch, "msg2")
	fmt.Println(readOnlyChan(ch))
	fmt.Println(<-ch)
}
