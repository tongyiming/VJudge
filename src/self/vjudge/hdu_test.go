package vjudge

import "testing"

func TestHDUJudger_Init(t *testing.T) {

	h := HDUJudger{}
	h.Init()
	h.Login()
	h.Submit()
}
