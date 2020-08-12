package main

import (
	test "github.com/lianan-234rf/module-test"
)

func TestA() string {
	res := test.Get("AAA")
	return res
}
