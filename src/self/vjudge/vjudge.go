package vjudge

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"self/models"

	log "github.com/sirupsen/logrus"
)

const (
	NONE   = 1
	C      = 2
	CPP    = 3
	JAVA   = 4
	GOLANG = 5
)

type VJudge struct {
	SubmitType string         `json:"submit_type"` //提交类型
	SubmitId   int64          `json:"submit_id"`   //提交id
	Problem    models.Problem `json:"problem"`     //题目信息
	Submit     models.Submit  `json:"submit"`      //提交信息
}

func (this *VJudge) DoJudge() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("Judge Error : %v", err)
		}
	}()

	this.getSubmitData()
	this.getProblemData(this.Submit.ProblemType)

	this.doJudge()
}

func (this *VJudge) doJudge() {
	user, _ := models.User{}.GetById(this.Problem.UserId)
	judge := newJudger(user.NickName)
	runid := newJudger(user.NickName).Submit(strconv.FormatInt(this.Submit.ProblemId, 10), this.Submit.Language, this.Submit.Code)
	res := judge.GetResult(runid)
	fmt.Println(res)
	//this.getCode()

}

//func (this *VJudge) getCode() {
//	minioCli := components.NewMinioCli()
//
//	path := this.WorkDir + "/code" + codeSuffixMap[this.Submit.Language]
//	minioCli.DownloadCode(this.Submit.Code, path)
//}

func (this *VJudge) saveSubmit() {
	switch this.SubmitType {
	case "submit":
		submit, err := models.Submit{}.GetById(this.SubmitId)
		if err != nil {
			panic(err)
		}
		submit.Result = this.Submit.Result
		submit.ResultDes = this.Submit.ResultDes
		submit.RunningTime = this.Submit.RunningTime
		submit.RunningMemory = this.Submit.RunningMemory

		err = models.Submit{}.Update(submit)
		if err != nil {
			panic(err)
		}

		log.Infof("saveSubmit: %#v", submit)

	case "submit_contest":
	case "submit_user":
	case "submit_test":
	}
}

func (this *VJudge) getProblemData(problemType string) {
	var problemJson []byte

	switch problemType {
	case "real":
		{
			problem, err := models.Problem{}.GetById(this.Submit.ProblemId)
			if err != nil {
				panic("getProblemData-Problem: " + err.Error())
			}
			problemJson, err = json.Marshal(problem)
			break
		}
	case "user":
		{
			problemUser, err := models.ProblemUser{}.GetById(this.Submit.ProblemId)
			if err != nil {
				panic("getProblemData-ProblemUser: " + err.Error())
			}
			problemJson, err = json.Marshal(problemUser)
			break
		}
	default:
		panic("getProblemData: not recognized ProblemType " + this.Submit.ProblemType)
	}

	if err := json.Unmarshal(problemJson, &this.Problem); err != nil {
		panic("getProblemData: " + err.Error())
	}

	log.Infof("getProblemData: %#v\n", this.Problem)
}

func (this *VJudge) getSubmitData() {
	var submitJson []byte

	switch this.SubmitType {
	case "submit":
		{
			submit, err := models.Submit{}.GetById(this.SubmitId)
			if err != nil {
				panic("getSubmitData-Submit: " + err.Error())
			}
			submitJson, err = json.Marshal(submit)
			break
		}
	case "submit_user":
		{
			submitUser, err := models.SubmitUser{}.GetById(this.SubmitId)
			if err != nil {
				panic("getSubmitData-SubmitUser: " + err.Error())
			}
			submitJson, err = json.Marshal(submitUser)
			break
		}
	case "submit_contest":
		{
			submitContest, err := models.SubmitContest{}.GetById(this.SubmitId)
			if err != nil {
				panic("getSubmitData-SubmitContest: " + err.Error())
			}
			submitJson, err = json.Marshal(submitContest)
			break
		}
	case "submit_test":
		{

		}
	default:
		panic("getSubmitData: not recognized submitType " + this.SubmitType)
	}

	if err := json.Unmarshal(submitJson, &this.Submit); err != nil {
		panic("getSubmitData: " + err.Error())
	}

	log.Infof("getSubmitData: %#v\n", this.Submit)
}

func getCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic("getCurrentPath: " + err.Error())
	}
	return dir
}
