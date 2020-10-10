package usecase

import (
	"fmt"
	"github.com/kons16/team7-backend/domain/entity"
	"github.com/kons16/team7-backend/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepo    repository.User
	sessionRepo repository.Session
}

// ユーザー登録する際に使用する構造体
type User struct {
	Name     string
	SubmitID string
	Year     int
	Sex      int
	Password string
}

func NewUserUseCase(userRepo repository.User, sessionRepo repository.Session) *UserUseCase {
	return &UserUseCase{userRepo: userRepo, sessionRepo: sessionRepo}
}

func (uc *UserUseCase) CreateUser(user *User) (string, error) {
	var createUserModel entity.User
	createUserModel.SubmitID = user.SubmitID
	createUserModel.Name = user.Name
	createUserModel.Year = user.Year
	createUserModel.Sex = user.Sex

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	createUserModel.PasswordHash = string(passwordHash)

	// MySQL にデータを保存
	id, err := uc.userRepo.Create(&createUserModel)
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	// Redis にセッションを保存
	sessionID, err := uc.sessionRepo.CreateUserSession(id)
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	return sessionID, nil
}
