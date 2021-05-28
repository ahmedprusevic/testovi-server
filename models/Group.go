package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Group struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:1000;not null;" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (g *Group) FillFields() {
	g.ID = 0
	g.Name = html.EscapeString(strings.TrimSpace(g.Name))
	g.CreatedAt = time.Now()
}

func (g *Group) Validate() error {
	if g.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

func (g *Group) SaveGroup(db *gorm.DB) (*Group, error) {
	err := db.Debug().Create(&g).Error
	if err != nil {
		return &Group{}, err
	}
	return g, nil
}

func (g *Group) FindAllGroups(db *gorm.DB) (*[]Group, error) {
	groups := []Group{}
	err := db.Debug().Model(&Group{}).Limit(100).Find(&groups).Error
	if err != nil {
		return &[]Group{}, err
	}
	return &groups, nil
}

func (g *Group) FindGroupById(db *gorm.DB, gid uint32) (*Group, error) {
	err := db.Debug().Model(&Group{}).Where("id = ?", gid).Take(&g).Error
	if err != nil {
		return &Group{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Group{}, errors.New("group not found")
	}
	return g, nil
}

func (g *Group) UpdateGroup(db *gorm.DB, gid uint32) (*Group, error) {
	db = db.Debug().Model(&Group{}).Where("id = ?", gid).UpdateColumns(map[string]interface{}{"name": g.Name})

	if db.Error != nil {
		return &Group{}, db.Error
	}

	err := db.Debug().Model(&Group{}).Where("id = ?", gid).Take(&g).Error

	if err != nil {
		return &Group{}, err
	}

	return g, nil

}

func (g *Group) DeleteGroup(db *gorm.DB, gid uint32) (int64, error) {
	db = db.Debug().Model(&Group{}).Where("id = ?", gid).Take(&Group{}).Delete(&Group{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
