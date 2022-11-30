package domain

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       int
	Username string `binding:"required"`
	Password string `binding:"required"`
	Role     string
	Name     string `binding:"required"`
}

type Login struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u User) HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}
