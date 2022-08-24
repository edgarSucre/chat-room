package usecase

import (
	"context"

	"github.com/edgarSucre/chat/internal/domain"
)

type AdminUsecase struct {
	repo   AdminRepository
	hasher Secure
}

func NewAdminUsecase(repository AdminRepository, opt ...AdminUsecaseOption) *AdminUsecase {
	uc := &AdminUsecase{repo: repository, hasher: Hasher{}}

	for _, v := range opt {
		v(uc)
	}

	return uc
}

func (uc *AdminUsecase) CreateUser(ctx context.Context, params domain.UserParam) (domain.UserResponse, error) {
	hashed, err := uc.hasher.SecurePassword(params.Password)
	if err != nil {
		return domain.UserResponse{}, domain.ErrBadParamInput
	}

	params.Password = hashed

	return uc.repo.CreateUser(ctx, params)
}