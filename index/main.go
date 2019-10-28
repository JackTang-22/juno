package main

import (
	"fmt"
	"time"
)

/*
#include<stdlib.h>
int test() {
    return rand() % (50000000 - 0 + 1) + 0;
}
int test1() {
	return rand() % (10000000 - 0 + 1) + 0;
}
*/
import "C"

/**
 * @author: tangye
 * @Date: 2019/10/27 13:53
 * @Description:
 */


func main() {
	// fmt.Println(C.test())
	var Int Func = func(a, b interface{}) int {
		if a.(int) > b.(int) {
			return 1
		} else if  a.(int) < b.(int) {
			return -1
		}
		return 0
	}

	sl := NewSKipListIterator(DEFAULT_MAX_LEVEL, Int)
	t1 := time.Now()
	for i := 0; i < 100; i++ {
		sl.Add(i, i)
		//	fmt.Println(sl.level)
	}
	fmt.Println(time.Since(t1))
    fmt.Println(sl.GetGE(10000))
	for sl.HasNext() {
		fmt.Printf("%d ", sl.Next())
	}
	fmt.Println()

	fmt.Println(sl.GetGE(46))
	sl.Del(46)
	sl.Del(77)
	fmt.Println(sl.GetGE(46))
	fmt.Println(sl.GetGE(77))
	fmt.Println(sl.GetGE(78))
	fmt.Println(sl.GetGE(37))

	//
	//fmt.Println()
	//
	//for j := 0; j < 10; j++ {
	//	t := time.Now()
	//	for i := 0; i < 15000; i++ {
	//		sl.GetK(int(C.test()))
	//	}
	//	fmt.Println(time.Since(t))
	//}
	//
	//fmt.Println()
	//
	//for j := 0; j < 10; j++ {
	//	t := time.Now()
	//	for i := 0; i < 100000; i++ {
	//		sl.GetK(int(C.test()))
	//	}
	//	fmt.Println(time.Since(t))
	//}
	//
	//fmt.Println()
	//for j := 0; j < 10; j++ {
	//	t := time.Now()
	//	for i := 0; i < 200000; i++ {
	//		sl.GetK(int(C.test()))
	//	}
	//	fmt.Println(time.Since(t))
	//}
	//
	//fmt.Println()
	//for j := 0; j < 10; j++ {
	//	t := time.Now()
	//	for i := 0; i < 300000; i++ {
	//		sl.GetK(int(C.test()))
	//	}
	//	fmt.Println(time.Since(t))
	//}
	//
	//fmt.Println()
	//for j := 0; j < 10; j++ {
	//	t := time.Now()
	//	for i := 0; i < 400000; i++ {
	//		sl.GetK(int(C.test()))
	//	}
	//	fmt.Println(time.Since(t))
	//}
	//
	//fmt.Println()
	//for j := 0; j < 10; j++ {
	//	t := time.Now()
	//	for i := 0; i < 400000; i++ {
	//		sl.GetK(int(C.test()))
	//	}
	//	fmt.Println(time.Since(t))
	//}

	//fmt.Println()
	//fmt.Println(sl.Len())
	//s := sl.findElements(3)
	//for i, v := range s {
	//	fmt.Println(i, v.next)
	//}
	//fmt.Println(sl.GetK(3))
	//fmt.Println(sl.Del(3))
	//fmt.Println(sl.GetK(4))
	//fmt.Println(sl.GetV(4))
	//c := 0
	//for sl.HasNext() {
	//	fmt.Printf("%d ",sl.Next())
	//	c++
	//}
	//fmt.Println()
	//fmt.Println(c)


}
