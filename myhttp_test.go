package main

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Countstr("aaabb123wZZF")      // 程序输出的结果
	want := map[int]int {
		'1' : 1,
		'2' : 1,
		'3' : 1,
		'F' : 1,
		'Z' : 2,
		'a' : 3,
		'b' : 2,
		'w' : 1,
	} // 期望的结果
	if !reflect.DeepEqual(want, got) {               // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}
