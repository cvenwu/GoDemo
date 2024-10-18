/*
 * @Author: cvenwu yirufeng@foxmail.com
 * @Date: 2024-10-16 09:41:53
 * @LastEditors: cvenwu yirufeng@foxmail.com
 * @LastEditTime: 2024-10-17 09:19:23
 * @Description:
 *
 * Copyright (c) 2024 by yirufeng@foxmail.com, All Rights Reserved.
 */
package main

import (
	"fmt"
	"os"
)

func main() {
	// 当使用 os.Exit 时 defer 将不会 被执行， 所以这里的 fmt.Println 将永远不会被调用。
	defer fmt.Println("456")

	os.Exit(3)
}
