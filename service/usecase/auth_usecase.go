package usecase

import (
	"context"
	"example/data/entity"
	"example/data/repository"
	"example/data/uow"
	"example/db"
	"example/service/request"
	"example/service/response"
)

type AuthUsecase interface {
	Register(ctx context.Context, req request.RegisterRequest) *response.UserResponse
}

type authUsecase struct {
	uow      uow.UnitOfWork
	userRepo repository.UserRepository
}

func NewAuthUsecase(db *db.Database) AuthUsecase {
	usecase := &authUsecase{
		uow: uow.NewUnitOfWork(db),
	}
	usecase.userRepo = usecase.uow.Users()
	return usecase
}

func (uc authUsecase) Register(
	ctx context.Context,
	req request.RegisterRequest,
) *response.UserResponse {
	data := &entity.User{}
	req.Map(data)

	uc.userRepo.Create(ctx, data)

	return response.ToUserResponse(data)
}
