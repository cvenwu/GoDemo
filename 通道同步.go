/*
 * @Author: cvenwu yirufeng@foxmail.com
 * @Date: 2024-10-17 09:40:34
 * @LastEditors: cvenwu yirufeng@foxmail.com
 * @LastEditTime: 2024-10-17 09:44:57
 * @Description:
 *
 * Copyright (c) 2024 by yirufeng@foxmail.com, All Rights Reserved.
 */
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("worker is running")
	time.Sleep(time.Second * 5)
	fmt.Println("done")
	done <- true
}

func main() {
	doneNotify := make(chan bool, 1)
	fmt.Println("456")
	go worker(doneNotify)
	fmt.Println("4567")
	<-doneNotify
	fmt.Println("45678")
}
