package models

import (
	"fmt"
	"testing"
)

func TestProblemCheck_Create(t *testing.T) {
	InitAllInTest()

	problemCheck := &ProblemCheck{Titile: "sadas", Description: "1111", ProblemId: "1", ProblemUserId: "1", CaseData: "123456"}
	if _, err := problemCheck.Create(problemCheck); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}

func TestProblemCheck_Remove(t *testing.T) {
	InitAllInTest()

	var problem ProblemCheck
	if err := problem.Remove(1); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}

func TestProblemCheck_Update(t *testing.T) {
	InitAllInTest()

	problemCheck := &ProblemCheck{Titile: "sadas", Description: "11221111", ProblemId: "1", ProblemUserId: "1"}
	if err := problemCheck.Update(problemCheck); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}

func TestProblemCheck_GetById(t *testing.T) {
	InitAllInTest()

	problemCheck := &ProblemCheck{Titile: "sadas", Description: "11221111", ProblemId: "1", ProblemUserId: "1"}
	ProblemCheck{}.Create(problemCheck)

	getProblemCheck, err := ProblemCheck{}.GetById(problemCheck.Id)
	if err != nil {
		t.Error("GetById() failed:", err)
	}

	if *getProblemCheck != *problemCheck {
		t.Error("GetById() failed:", "%v != %v", problemCheck, getProblemCheck)
	}
}

func TestProblemCheckQueryByCaseData(t *testing.T) {
	InitAllInTest()
	problemCheck := &ProblemCheck{Titile: "sadas", Description: "fffff", CaseData: "123456"}

	getProblemCheck, err := ProblemCheck{}.QueryByCaseData(problemCheck.CaseData)

	if err != nil {
		t.Error("GetByCaseData() failed:", err.Error())
	}

	fmt.Println(len(getProblemCheck))

}
