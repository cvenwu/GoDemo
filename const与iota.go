/*
 * @Author: cvenwu yirufeng@foxmail.com
 * @Date: 2024-10-15 10:56:59
 * @LastEditors: cvenwu yirufeng@foxmail.com
 * @LastEditTime: 2024-10-15 11:03:33
 * @Description:
 *
 * Copyright (c) 2024 by yirufeng@foxmail.com, All Rights Reserved.
 */
package main

import "fmt"

// const可以用来定义枚举类型
const (
	BEIJING   = 0
	SHANGHAI  = 1
	GUANGZHOU = 2
	CHONGQING = 3
	SHENZHEN  = 4
	HAERBIN   = 5
)

// const与iota搭配使用
const (
	a, b = iota + 1, iota + 2
	c, d
	e, f

	g, h = iota * 2, iota * 3
	i, k
)

func main() {

	// 定义一个常量
	const length int = 32
	fmt.Println("length = ", length)

	// 报错：常量不允许修改
	// length = 33

	fmt.Println(a, b, c, d, e, f)
	fmt.Println(g, h, i, k)

}
