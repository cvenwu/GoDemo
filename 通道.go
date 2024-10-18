/*
 * @Author: cvenwu yirufeng@foxmail.com
 * @Date: 2024-10-17 09:31:22
 * @LastEditors: cvenwu yirufeng@foxmail.com
 * @LastEditTime: 2024-10-17 09:36:42
 * @Description:
 *
 * Copyright (c) 2024 by yirufeng@foxmail.com, All Rights Reserved.
 */
package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() { ch <- "ping" }()

	msg := <-ch
	fmt.Println("接收到消息：", msg)
}
