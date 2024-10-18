/*
 * @Author: cvenwu yirufeng@foxmail.com
 * @Date: 2024-10-17 10:04:01
 * @LastEditors: cvenwu yirufeng@foxmail.com
 * @LastEditTime: 2024-10-17 10:14:35
 * @Description:
 *
 * Copyright (c) 2024 by yirufeng@foxmail.com, All Rights Reserved.
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("ch1 msg:", msg1)
		case msg2 := <-ch2:
			fmt.Println("ch2 msg:", msg2)
		}
	}

	fmt.Println("123")
}
