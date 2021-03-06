package usecase

import (
	"fmt"
	"github.com/kons16/team7-backend/domain/entity"
	"github.com/kons16/team7-backend/domain/repository"
)

type CordinateUseCase struct {
	cordiRepo   repository.Cordinate
	clothRepo   repository.Cloth
	sessionRepo repository.Session
}

func NewCordinateUseCase(cordiRepo repository.Cordinate, sessionRepo repository.Session, clothRepo repository.Cloth) *CordinateUseCase {
	return &CordinateUseCase{cordiRepo: cordiRepo, sessionRepo: sessionRepo, clothRepo: clothRepo}
}

// コーディネートを新規追加
func (cdu *CordinateUseCase) CreateCordinate(cordinate *entity.Cordinate, sessionID string) error {
	// sessionID から userID を取得
	userID := cdu.sessionRepo.CheckBySession(sessionID)
	cordinate.UserID = userID

	err := cdu.cordiRepo.Create(cordinate)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

// すべてのコーディネート情報を取得
func (cdu *CordinateUseCase) GetAll(sessionID string) *[]entity.Cordinate {
	// sessionID から userID を取得
	userID := cdu.sessionRepo.CheckBySession(sessionID)
	cordinates := cdu.cordiRepo.Get(userID)
	return cordinates
}

// 服IDs の服情報を全て取得
func (cdu *CordinateUseCase) GetClothes(ids []int) *[]entity.Cloth {
	clothes := cdu.clothRepo.GetByIDs(ids)
	return clothes
}
