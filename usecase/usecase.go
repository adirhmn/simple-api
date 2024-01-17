package usecase

import (
	"context"
	"simpleapi/models"
	"simpleapi/repository"
)

type UserUseCaseInt interface {
	GetAllUser(ctx context.Context) []models.User
	GetUser(ctx context.Context, userId string) models.User
	UpdateUser(ctx context.Context, user models.User, userId string) models.User
	CreateUser(ctx context.Context, user models.User) models.User
	DeleteUser(ctx context.Context, userId string)
}

type UserUseCase struct {
	userRepo repository.UserRepoInt
}

func NewUserUseCase(userRepo repository.UserRepoInt) UserUseCaseInt {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (uc UserUseCase) GetAllUser(ctx context.Context) []models.User {
	users := uc.userRepo.FindAll(ctx)
	return users
}

func (uc UserUseCase) GetUser(ctx context.Context, userId string) models.User {
	user := uc.userRepo.FindById(ctx, userId)
	return user
}

func (uc UserUseCase) UpdateUser(ctx context.Context, user models.User, userId string) models.User {
	updatedUser := uc.userRepo.Update(ctx, user, userId)
	return updatedUser
}

func (uc UserUseCase) CreateUser(ctx context.Context, user models.User) models.User {
	addedUser := uc.userRepo.Create(ctx, user)
	return addedUser
}

func (uc UserUseCase) DeleteUser(ctx context.Context, userId string) {
	uc.userRepo.Delete(ctx, userId)
}
