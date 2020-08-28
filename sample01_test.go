package sample01

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	actual := HelloWorld("hoge")
	expected := "hello world, hoge6"
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func failTestHelloWorld(t *testing.T) {
	actual := HelloWorld("hoge1")
	expected := "hello world, hoge"
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
