package test

import (
	"testing"
)

// TestHelloWorld 测试用例
func TestHelloWorld(t *testing.T) {
	t.Log("hello world")
}

// TestA 测试用例
func TestA(t *testing.T) {
	t.Log("TestA")
}

// TestB 测试用例
func TestB(t *testing.T) {
	t.Log("TestB")
}

func TestFailNow(t *testing.T) {
	t.Log("before fail: 标记测试错误,终止该案例")
	t.FailNow()
	t.Log("after fail")
}

func TestFail(t *testing.T) {
	t.Log("before fail: 标记测试错误,继续执行")
	t.Fail()
	t.Log("after fail")
}
