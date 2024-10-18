/*
 * @Author: cvenwu yirufeng@foxmail.com
 * @Date: 2024-10-17 09:37:13
 * @LastEditors: cvenwu yirufeng@foxmail.com
 * @LastEditTime: 2024-10-17 09:39:24
 * @Description:
 *
 * Copyright (c) 2024 by yirufeng@foxmail.com, All Rights Reserved.
 */
package main

import "fmt"

func main() {
	buffered_chan := make(chan string, 2)
	buffered_chan <- "buffered"
	buffered_chan <- "channel"

	fmt.Println(<-buffered_chan)
	fmt.Println(<-buffered_chan)˜
}
˜