package models

import (
	errors2 "errors"
	"html"
	"strings"
	"time"

	"github.com/Pcorner0/corpmin-api/utils/errors"
	"github.com/Pcorner0/corpmin-api/utils/token"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Users struct {
	ID        int            `json:"idvendedor" gorm:"primaryKey; autoIncrement:true"`
	Nombre    string         `json:"nombre"`
	Apellidop string         `json:"apellidop" `
	Apellidom string         `json:"apellidom"`
	Password  string         `json:"password"`
	Email     string         `json:"email" binding:"required"`
	Rol       string         `json:"rol"`
	Office    string         `json:"office"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (u Users) ValidateUser() *errors.RestErr {
	u.Email = strings.TrimSpace(u.Email)
	u.Password = strings.TrimSpace(u.Password)
	if u.Email == "" {
		return errors.NewBadRequestError("dirección email invalida")
	}
	if u.Password == "" {
		return errors.NewBadRequestError("contraseña invalida")
	}
	return nil
}

func RegisterUser(db *gorm.DB, user *Users) (err error) {
	err = db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(db *gorm.DB, user *Users, email string) (err error) {
	db.Raw("SELECT * FROM users WHERE email = ? AND deleted_at IS NULL LIMIT 1", email).Scan(user)
	if err != nil {
		return err
	}
	return nil
}

// ########################################################################################

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(db *gorm.DB, email string, password string) (string, error, Users) {
	user := Users{}

	//db.Raw("SELECT * FROM users WHERE email = ? AND deleted_at IS NULL LIMIT 1", email).Scan(user)
	err := db.Model(Users{}).Where("email = ?", email).Take(&user).Error

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err, user
	}

	token, err := token.GenerateToken(user.ID)

	if err != nil {
		return "", err, user
	}

	return token, nil, user
}

func (u *Users) BeforeSave(*gorm.DB) error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return nil

}

func GetUserByID(db *gorm.DB, uid uint) (Users, error) {

	var user Users

	if err := db.First(&user, uid).Error; err != nil {
		return user, errors2.New("usuario no encontrado")
	}

	user.PrepareGive()

	return user, nil

}

func (u *Users) PrepareGive() {
	u.Password = ""
}

func (u *Users) SaveUser(db *gorm.DB) (*Users, error) {

	err := db.Create(&u).Error

	if err != nil {
		return &Users{}, err
	}
	return u, nil
}
