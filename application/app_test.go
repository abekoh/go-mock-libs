package application

import (
	"github.com/abekoh/go-mock-libs/domain/model/examination"
	"github.com/abekoh/go-mock-libs/domain/model/user"
	"github.com/abekoh/go-mock-libs/domain/types"
	"github.com/google/uuid"
)

func testUser() user.User {
	userID := uuid.New()
	name, _ := user.NewName("Taro", "Yamada")
	birthday, _ := types.NewDate(1990, 12, 31)
	return user.NewUserWithID(userID, name, birthday)
}

func testExams(userID uuid.UUID) examination.ExaminationList {
	examID1 := uuid.New()
	examDate1, _ := types.NewDate(2022, 4, 25)
	exam1 := examination.NewExaminationWithID(examID1, userID, examination.ExaminationTypeEnglish, examDate1, examination.Score(85))
	examID2 := uuid.New()
	examDate2, _ := types.NewDate(2022, 5, 1)
	exam2 := examination.NewExaminationWithID(examID2, userID, examination.ExaminationTypeMath, examDate2, examination.Score(53))
	return examination.ExaminationList{exam1, exam2}
}
