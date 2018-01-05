package oj_getter

import (
	"self/commons"
	"testing"
)

func TestGetProblem(t *testing.T) {
	commons.InitAll("../../../cfg/cfg.toml.debug")
	h := HDUGetter{}
	h.getProblem("http://acm.hdu.edu.cn/showproblem.php?pid=1000")
	//
	//for i := 1000; i < 1500; i++ {
	//	h.getProblem("")
	//}
}
