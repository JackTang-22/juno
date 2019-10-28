package main
//
//import "C"
//import (
//	"fmt"
//	"math/rand"
//	"testing"
//	"time"
//)
//
//
//var sl = NewSKipListIterator(DEFAULT_MAX_LEVEL)
//
//func init() {
//
//	t1 := time.Now()
//	for i := 0; i < 10000000; i++ {
//		//sl.Add(C.test(), i)
//		//	fmt.Println(sl.level)
//	}
//	fmt.Println(time.Since(t1))
//}
//
//func BenchmarkNewSkipList(b *testing.B) {
//
//	//fmt.Println(sl.GetGE(3))
//	fmt.Println(sl.Del(3))
//	fmt.Println(sl.GetK(5))
//	fmt.Println(sl.GetV(5))
//	for i := 0; i < b.N; i++ {
//		for j := 0; j < 100000; j++ {
//			sl.GetGE(num[j])
//		}
//	}
//
//	//t2 := time.Now()
//	//for i := 0; i < 100000; i++ {
//	//	sl.GetGE(i)
//	//}
//	//fmt.Println(time.Since(t2))
//
//}
