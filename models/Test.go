package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Test struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:1000;not null" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

type IntersectionTestQuestion struct {
	ID         uint32
	QuestionID uint32 `gorm:"foreignkey:question_id"`
	TestID     uint32 `gorm:"foreignkey:test_id"`
}

func (t *Test) FillFields() {
	t.ID = 0
	t.Name = html.EscapeString(strings.TrimSpace(t.Name))
	t.CreatedAt = time.Now()
}

func (t *Test) Validate() error {
	if t.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

func (t *Test) SaveTest(db *gorm.DB) (*Test, error) {

	err := db.Debug().Create(&t).Error

	if err != nil {
		return &Test{}, err
	}

	return t, nil
}

func (t *Test) FindAllTests(db *gorm.DB) (*[]Test, error) {

	tests := []Test{}

	err := db.Debug().Model(&Test{}).Limit(100).Find(&tests).Error

	if err != nil {
		return &[]Test{}, nil
	}

	return &tests, nil
}

func (t *Test) FindTestById(db *gorm.DB, tid uint32) (*[]*Question, error) {

	question := Question{}

	questions := []*Question{}

	qids := []IntersectionTestQuestion{}

	err := db.Debug().Table("test_questions").Where("id = ?", tid).Limit(40).Find(&qids).Error

	if err != nil {
		return &[]*Question{}, err
	}

	for _, i := range qids {
		q, err := question.FindQuestionById(db, i.QuestionID)

		if err != nil {
			return &questions, err
		}

		questions = append(questions, q)
	}

	return &questions, nil

}

func (t *Test) UpdateTest(db *gorm.DB, tid uint32) (*Test, error) {

	db = db.Debug().Model(&Test{}).Where("id = ?", tid).Take(&Test{}).UpdateColumns(map[string]interface{}{
		"name": t.Name,
	})

	if db.Error != nil {
		return &Test{}, db.Error
	}

	err := db.Debug().Model(&Test{}).Where("id = ?", tid).Take(&t).Error

	if err != nil {
		return &Test{}, nil
	}

	return t, nil

}

func (t *Test) DeleteTest(db *gorm.DB, tid uint32) (int64, error) {

	db = db.Debug().Model(&Test{}).Where("id = ?", tid).Take(&Test{}).Delete(&Test{})

	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil

}

func (t *Test) AddQuestionToTest(db *gorm.DB, tid uint32, qid uint32) (*IntersectionTestQuestion, error) {

	intersection := IntersectionTestQuestion{TestID: tid, QuestionID: qid}

	err := db.Debug().Table("test_questions").Where("question_id = ?", qid).Where("test_id = ?", tid).Find(&IntersectionTestQuestion{}).Error

	if err == nil {
		return &IntersectionTestQuestion{}, nil
	}

	err = db.Debug().Table("test_questions").Create(&intersection).Error

	if err != nil {
		return &IntersectionTestQuestion{}, err
	}

	return &intersection, nil

}
