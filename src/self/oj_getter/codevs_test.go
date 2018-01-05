package oj_getter

import "testing"

func TestCodeVSGetter_GetProblem(t *testing.T) {
	c := CodeVSGetter{}
	c.GetProblem("http://www.codevs.cn/problem/1001/")
}
