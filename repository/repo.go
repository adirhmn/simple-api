package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"simpleapi/helper"
	"simpleapi/models"
	"simpleapi/tools/dbpool"
)

type UserRepoInt interface {
	Create(ctx context.Context, user models.User) models.User
	Update(ctx context.Context, user models.User, userId string) models.User
	Delete(ctx context.Context, userId string)
	FindById(ctx context.Context, userId string) models.User
	FindAll(ctx context.Context) []models.User
}

type UserRepo struct {
	pools *dbpool.DBPools
}

func NewUserRepo(pools *dbpool.DBPools) UserRepoInt {
	return &UserRepo{
		pools: pools,
	}
}

func (repo UserRepo) Create(ctx context.Context, user models.User) models.User {
	newUser := user
	if err := newUser.Insert(ctx, repo.pools.GetWriteDB(), boil.Infer()); err != nil {
		b, _ := json.Marshal(err)
		fmt.Println(b)
	}
	return newUser
}

func (repo UserRepo) Update(ctx context.Context, user models.User, userId string) models.User {
	userToUpdate, err := models.FindUser(ctx, repo.pools.GetReadDB(), userId)
	if err != nil {
		helper.PanicIfError(err)
	}

	userToUpdate.FirstName = user.FirstName
	userToUpdate.LastName = user.LastName
	userToUpdate.Email = user.Email
	userToUpdate.Active = user.Active
	userToUpdate.Gender = user.Gender
	userToUpdate.BirthPlace = user.BirthPlace
	userToUpdate.DateOfBirth = user.DateOfBirth
	userToUpdate.RegisterProvider = user.RegisterProvider
	userToUpdate.Photo = user.Photo
	userToUpdate.Profession = user.Profession
	userToUpdate.MobilephonePrefix = user.MobilephonePrefix

	_, err = userToUpdate.Update(ctx, repo.pools.GetWriteDB(), boil.Infer())
	if err != nil {
		helper.PanicIfError(err)
	}

	return user
}

func (repo UserRepo) Delete(ctx context.Context, userId string) {
	user, err := models.FindUser(ctx, repo.pools.GetReadDB(), userId)
	helper.PanicIfError(err)
	userToDelete := user
	if _, err := userToDelete.Delete(ctx, repo.pools.GetWriteDB()); err != nil {
		helper.PanicIfError(err)
	}
}

func (repo UserRepo) FindById(ctx context.Context, userId string) models.User {
	user, err := models.FindUser(ctx, repo.pools.GetReadDB(), userId)
	helper.PanicIfError(err)

	userResp := models.User{
		ID:                  user.ID,
		Email:               user.Email,
		EmailVerified:       user.EmailVerified,
		Mobilephone:         user.Mobilephone,
		MobilephoneVerified: user.MobilephoneVerified,
		Active:              user.Active,
		FirstName:           user.FirstName,
		LastName:            user.LastName,
		Gender:              user.Gender,
		BirthPlace:          user.BirthPlace,
		DateOfBirth:         user.DateOfBirth,
		RegisterProvider:    user.RegisterProvider,
		Photo:               user.Photo,
		Deleted:             user.Deleted,
		CreatedAt:           user.CreatedAt,
		UpdatedAt:           user.UpdatedAt,
		Profession:          user.Profession,
		MobilephonePrefix:   user.MobilephonePrefix,
		IsAdmin:             user.IsAdmin,
		IsStaff:             user.IsStaff,
		MustUpdatePassword:  user.MustUpdatePassword,
	}

	return userResp
}

func (repo UserRepo) FindAll(ctx context.Context) []models.User {
	users, err := models.Users().All(ctx, repo.pools.GetReadDB())
	if err != nil {
		helper.PanicIfError(err)
	}

	var allUser []models.User
	for _, user := range users {
		allUser = append(allUser, *user)
	}
	return allUser
}
