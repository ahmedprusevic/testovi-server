package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;" json:"name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	Role      string    `gorm:"size:100;not null;" json:"role"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) FillFields() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Role = html.EscapeString(strings.TrimSpace(u.Role))
	u.CreatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("name is required")
		}
		if u.Password == "" {
			return errors.New("password is required")
		}
		if u.Email == "" {
			return errors.New("email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("please input valid email")
		}
		return nil
	case "login":
		if u.Password == "" {
			return errors.New("password is required")
		}
		if u.Email == "" {
			return errors.New("email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("please input valid email")
		}
		return nil
	default:
		if u.Name == "" {
			return errors.New("name is required")
		}
		if u.Password == "" {
			return errors.New("password is required")
		}
		if u.Email == "" {
			return errors.New("email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("please input valid email")
		}
		return nil
	}

}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	err := u.HashPassword()

	if err != nil {
		return &User{}, err
	}

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	users := []User{}
	err := db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err

}

func (u *User) FindUserById(db *gorm.DB, uid uint32) (*User, error) {
	err := db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("user not found")
	}
	return u, err
}

func (u *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {
	err := u.HashPassword()
	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password": u.Password,
			"nickname": u.Name,
			"email":    u.Email,
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteUser(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
