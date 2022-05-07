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
	birthday, err := parseDate(r.Birthday)
	if err != nil {
		return user.User{}, err
	}
	return user.NewUser(name, birthday), nil
}

type ExamAddRequest struct {
	UserID string `json:"user_id"`
	Type   string `json:"type"`
	Score  int    `json:"score"`
	Date   string `json:"date"`
}

func (r ExamAddRequest) NewExam() (examination.Examination, error) {
	userID, err := uuid.Parse(r.UserID)
	if err != nil {
		return examination.Examination{}, err
	}
	examType, err := parseExamType(r.Type)
	if err != nil {
		return examination.Examination{}, err
	}
	date, err := parseDate(r.Date)
	if err != nil {
		return examination.Examination{}, err
	}
	score, err := examination.NewScore(r.Score)
	if err != nil {
		return examination.Examination{}, err
	}
	return examination.NewExamination(userID, examType, date, score), nil
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
		Type:  examTypeString(exam.Type()),
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

func parseDate(s string) (types.Date, error) {
	invalidErr := errors.New("invalid birthday")
	ss := strings.Split(s, "/")
	if len(ss) != 3 {
		return types.Date{}, invalidErr
	}
	year, err := strconv.Atoi(ss[0])
	if err != nil {
		return types.Date{}, invalidErr
	}
	month, err := strconv.Atoi(ss[1])
	if err != nil {
		return types.Date{}, invalidErr
	}
	day, err := strconv.Atoi(ss[2])
	if err != nil {
		return types.Date{}, invalidErr
	}
	return types.NewDate(year, month, day)
}

func dateString(b types.Date) string {
	return fmt.Sprintf("%04d/%02d/%02d", b.Year(), b.Month(), b.Day())
}

func parseExamType(s string) (examination.ExaminationType, error) {
	switch s {
	case "English":
		return examination.ExaminationTypeEnglish, nil
	case "Math":
		return examination.ExaminationTypeMath, nil
	default:
		return examination.ExaminationTypeEnglish, errors.New("invalid type")
	}
}

func examTypeString(t examination.ExaminationType) string {
	switch t {
	case examination.ExaminationTypeEnglish:
		return "English"
	case examination.ExaminationTypeMath:
		return "Math"
	default:
		return ""
	}
}
