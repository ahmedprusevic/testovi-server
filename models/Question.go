package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Question struct {
	ID        uint32         `gorm:"primary_key;auto_increment" json:"id"`
	Name      string         `gorm:"size:1000;not null;" json:"name"`
	Answers   pq.StringArray `gorm:"type:text[];not null;" json:"answers"`
	Correct   pq.StringArray `gorm:"type:text[];not null;" json:"correct"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	GroupID   uint32         `gorm:"type:bigserial;not null;" json:"group_id"`
	TestID    pq.Int64Array  `gorm:"type:bigserial[];" json:"test_id"`
}

func (q *Question) FillFields() {
	q.ID = 0
	q.Name = html.EscapeString(strings.TrimSpace(q.Name))
	q.CreatedAt = time.Now()
}

func (q *Question) Validate() error {

	if q.Name == "" {
		return errors.New("name is required")
	}
	if len(q.Answers) == 0 {
		return errors.New("you need at least one answer in order to create question")
	}
	if len(q.Correct) == 0 {
		return errors.New("you need at least one correct answer in order to create question")
	}

	return nil

}

func (q *Question) SaveQuestion(db *gorm.DB) (*Question, error) {

	err := db.Debug().Create(&q).Error

	if err != nil {
		return &Question{}, err
	}

	return q, nil
}

func (q *Question) FindAllQuestions(db *gorm.DB) (*[]Question, error) {

	questions := []Question{}

	err := db.Debug().Model(&Question{}).Limit(100).Find(&questions).Error

	if err != nil {
		return &[]Question{}, nil
	}

	return &questions, nil
}

func (q *Question) FindQuestionById(db *gorm.DB, qid uint32) (*Question, error) {

	err := db.Debug().Model(Question{}).Where("id = ?", qid).Take(&q).Error

	if err != nil {
		return &Question{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &Question{}, err
	}

	return q, nil
}

func (q *Question) UpdateQuestion(db *gorm.DB, qid uint32) (*Question, error) {

	db = db.Debug().Model(&Question{}).Where("id = ?", qid).Take(&Question{}).UpdateColumns(map[string]interface{}{
		"name":     q.Name,
		"answers":  q.Answers,
		"correct":  q.Correct,
		"group_id": q.GroupID,
		"test_id":  q.TestID,
	})

	if db.Error != nil {
		return &Question{}, db.Error
	}

	err := db.Debug().Model(&Question{}).Where("id = ?", qid).Take(&q).Error

	if err != nil {
		return &Question{}, nil
	}

	return q, nil
}

func (q *Question) DeleteQuestion(db *gorm.DB, qid uint32) (int64, error) {

	db = db.Debug().Model(&Question{}).Where("id = ?", qid).Take(&Question{}).Delete(&Question{})

	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil

}
