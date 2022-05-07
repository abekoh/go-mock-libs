package examination

import (
	"fmt"

	"github.com/abekoh/go-mock-libs/domain/types"
	"github.com/google/uuid"
)

type Examination struct {
	id       uuid.UUID
	userId   uuid.UUID
	examType ExaminationType
	examDate types.Date
	score    Score
}

func NewExamination(userId uuid.UUID, examType ExaminationType, examDate types.Date, score Score) Examination {
	return NewExaminationWithID(uuid.New(), userId, examType, examDate, score)
}

func NewExaminationWithID(id, userId uuid.UUID, examType ExaminationType, examDate types.Date, score Score) Examination {
	return Examination{
		id:       id,
		userId:   userId,
		examType: examType,
		examDate: examDate,
		score:    score,
	}
}

func (e Examination) ID() uuid.UUID {
	return e.id
}

func (e Examination) Type() ExaminationType {
	return e.examType
}

func (e Examination) Date() types.Date {
	return e.examDate
}

func (e Examination) Score() Score {
	return e.score
}

type ExaminationType int

const (
	ExaminationTypeEnglish ExaminationType = iota
	ExaminationTypeMath
)

type ExaminationList []Examination

type Score int

func NewScore(score int) (Score, error) {
	if score > 100 || score < 0 {
		return Score(0), fmt.Errorf("invalid score")
	}
	return Score(score), nil
}

func (s Score) Int() int {
	return int(s)
}
