package model

import (
	"github.com/google/uuid"
)

type Examination struct {
	id       uuid.UUID
	userId   uuid.UUID
	examType ExaminationType
	examDate Birthday
}

func NewExamination(userId uuid.UUID, examType ExaminationType, examDate Birthday) Examination {
	return NewExaminationWithID(uuid.New(), userId, examType, examDate)
}

func NewExaminationWithID(id, userId uuid.UUID, examType ExaminationType, examDate Birthday) Examination {
	return Examination{
		id:       id,
		userId:   userId,
		examType: examType,
		examDate: examDate,
	}
}

type ExaminationType int

const (
	ExaminationTypeEnglish ExaminationType = iota
	ExaminationTypeMath
)
