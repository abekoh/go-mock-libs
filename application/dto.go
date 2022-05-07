package application

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/abekoh/go-mock-libs/domain/model/examination"
	"github.com/abekoh/go-mock-libs/domain/model/user"
	"github.com/abekoh/go-mock-libs/domain/types"
	"github.com/google/uuid"
)

type UserExamGetRequest struct {
	ID string `json:"id"`
}

func (r UserExamGetRequest) UUID() (uuid.UUID, error) {
	return uuid.Parse(r.ID)
}

type UserAddRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Birthday  string `json:"birthday"`
}

func (r UserAddRequest) NewUser() (user.User, error) {
	name, err := user.NewName(r.FirstName, r.LastName)
	if err != nil {
		return user.User{}, err
	}
	year, month, day, err := dateInts(r.Birthday)
	if err != nil {
		return user.User{}, err
	}
	birthday, err := types.NewDate(year, month, day)
	if err != nil {
		return user.User{}, err
	}
	return user.NewUser(name, birthday), nil
}

type ExamResponse struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Score int    `json:"score"`
	Date  string `json:"date"`
}

func NewExamResponse(exam examination.Examination) ExamResponse {
	return ExamResponse{
		ID:    exam.ID().String(),
		Type:  exam.Type().String(),
		Score: exam.Score().Int(),
		Date:  dateString(exam.Date()),
	}
}

type ExamResponseList []ExamResponse

func NewExamListReponse(examList examination.ExaminationList) ExamResponseList {
	result := make(ExamResponseList, 0, len(examList))
	for _, exam := range examList {
		result = append(result, NewExamResponse(exam))
	}
	return result
}

type UserExamResponse struct {
	ID       string           `json:"id"`
	FullName string           `json:"full_name"`
	Birthday string           `json:"birthday"`
	Exams    ExamResponseList `json:"exams"`
}

func NewUserExamResponse(user user.User, examList examination.ExaminationList) UserExamResponse {
	return UserExamResponse{
		ID:       user.ID().String(),
		FullName: user.Name().FullName(),
		Birthday: dateString(user.Birthday()),
		Exams:    NewExamListReponse(examList),
	}
}

func dateInts(s string) (int, int, int, error) {
	invalidErr := func() (int, int, int, error) {
		return 0, 0, 0, errors.New("invalid birthday")
	}
	ss := strings.Split(s, "/")
	if len(ss) != 3 {
		return invalidErr()
	}
	year, err := strconv.Atoi(ss[0])
	if err != nil {
		return invalidErr()
	}
	month, err := strconv.Atoi(ss[1])
	if err != nil {
		return invalidErr()
	}
	day, err := strconv.Atoi(ss[2])
	if err != nil {
		return invalidErr()
	}
	return year, month, day, nil
}

func dateString(b types.Date) string {
	return fmt.Sprintf("%04d/%02d/%02d", b.Year(), b.Month(), b.Day())
}
