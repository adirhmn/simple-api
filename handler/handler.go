package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"net/http"
	"simpleapi/app"
	"simpleapi/dto"
	"simpleapi/models"
	"time"
)

type UserHandler struct {
	Containers *app.ServiceContainers
}

func (uh *UserHandler) GetUsers(c echo.Context) error {
	ctx := c.Request().Context()

	users := uh.Containers.UserUseCase.GetAllUser(ctx)

	return c.JSON(http.StatusOK, users)
}

func (uh *UserHandler) GetUser(c echo.Context) error {
	ctx := c.Request().Context()

	userId := c.Param("id")
	users := uh.Containers.UserUseCase.GetUser(ctx, userId)

	return c.JSON(http.StatusOK, users)
}

func (uh *UserHandler) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()
	user := new(dto.UserRequest)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON data"})
	}

	dateOfBirth, _ := time.Parse("2006-01-02", user.DateOfBirth)

	newUser := models.User{
		ID:                  uuid.New().String(),
		Email:               null.String{String: user.Email, Valid: true},
		EmailVerified:       false,
		Mobilephone:         null.String{String: user.Mobilephone, Valid: true},
		MobilephoneVerified: false,
		Active:              true,
		FirstName:           null.String{String: user.FirstName, Valid: true},
		LastName:            null.String{String: user.LastName, Valid: true},
		Gender:              null.String{String: user.Gender, Valid: true},
		BirthPlace:          null.String{String: user.BirthPlace, Valid: true},
		DateOfBirth:         null.Time{Time: dateOfBirth, Valid: true},
		RegisterProvider:    null.String{String: user.RegisterProvider, Valid: true},
		Photo:               null.String{String: user.Photo, Valid: true},
		Deleted:             false,
		Profession:          null.String{String: user.Profession, Valid: true},
		MobilephonePrefix:   null.String{String: user.MobilephonePrefix, Valid: true},
		IsAdmin:             false,
		IsStaff:             false,
		MustUpdatePassword:  false,
	}
	addedUser := uh.Containers.UserUseCase.CreateUser(ctx, newUser)
	return c.JSON(http.StatusOK, addedUser)
}

func (uh *UserHandler) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Param("id")
	user := new(dto.UserRequest)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON data"})
	}

	dateOfBirth, _ := time.Parse("2006-01-02", user.DateOfBirth)

	updatedUser := models.User{
		ID:                  uuid.New().String(),
		Email:               null.String{String: user.Email, Valid: true},
		EmailVerified:       false,
		Mobilephone:         null.String{String: user.Mobilephone, Valid: true},
		MobilephoneVerified: false,
		Active:              true,
		FirstName:           null.String{String: user.FirstName, Valid: true},
		LastName:            null.String{String: user.LastName, Valid: true},
		Gender:              null.String{String: user.Gender, Valid: true},
		BirthPlace:          null.String{String: user.BirthPlace, Valid: true},
		DateOfBirth:         null.Time{Time: dateOfBirth, Valid: true},
		RegisterProvider:    null.String{String: user.RegisterProvider, Valid: true},
		Photo:               null.String{String: user.Photo, Valid: true},
		Deleted:             false,
		Profession:          null.String{String: user.Profession, Valid: true},
		MobilephonePrefix:   null.String{String: user.MobilephonePrefix, Valid: true},
		IsAdmin:             false,
		IsStaff:             false,
		MustUpdatePassword:  false,
	}
	addedUser := uh.Containers.UserUseCase.UpdateUser(ctx, updatedUser, userId)
	return c.JSON(http.StatusOK, addedUser)
}

func (uh *UserHandler) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()

	userId := c.Param("id")
	uh.Containers.UserUseCase.DeleteUser(ctx, userId)
	return c.JSON(http.StatusOK, map[string]string{"status": "oke"})
}

//func FetchAllUsers(c echo.Context) error {
//
//	db, err := sql.Open("postgres", "user=santika password=my_password dbname=testaja sslmode=disable")
//	if err != nil {
//		fmt.Println("Gagal terhubung ke database:", err)
//		return c.JSON(http.StatusInternalServerError, err.Error())
//	}
//	defer db.Close()
//	// Query data
//	users, err := models.Users().All(context.Background(), db)
//	if err != nil {
//		fmt.Println("Gagal mengambil data dari database:", err)
//		return c.JSON(http.StatusInternalServerError, err.Error())
//	}
//
//	return c.JSON(http.StatusOK, users)
//}
//
//func StoreUser(c echo.Context) error {
//	db, err := sql.Open("postgres", "user=santika password=my_password dbname=testaja sslmode=disable")
//	if err != nil {
//		fmt.Println("Gagal terhubung ke database:", err)
//		return c.JSON(http.StatusInternalServerError, err.Error())
//	}
//	defer db.Close()
//
//	date := time.Date(2023, time.October, 17, 0, 0, 0, 0, time.UTC)
//	newUser := &models.User{
//		ID:                  uuid.New().String(),
//		Email:               null.String{String: "al@gmail.com", Valid: true},
//		EmailVerified:       false,
//		Mobilephone:         null.String{String: "8821238488", Valid: true},
//		MobilephoneVerified: false,
//		Active:              true,
//		FirstName:           null.String{String: "John", Valid: true},
//		LastName:            null.String{String: "Doe", Valid: true},
//		Gender:              null.String{String: "M", Valid: true},
//		BirthPlace:          null.String{String: "Jakarta", Valid: true},
//		DateOfBirth:         null.Time{Time: date, Valid: true},
//		RegisterProvider:    null.String{},
//		Photo:               null.String{},
//		Deleted:             false,
//		Profession:          null.String{},
//		MobilephonePrefix:   null.String{String: "+62", Valid: true},
//		IsAdmin:             false,
//		IsStaff:             false,
//		MustUpdatePassword:  false,
//	}
//	err = newUser.Insert(context.Background(), db, boil.Infer())
//	if err != nil {
//		fmt.Println("Gagal melakukan operasi penambahan data:", err)
//		return err
//	}
//
//	return c.String(http.StatusOK, "Oke")
//}
//
//func UpdateUser(c echo.Context) error {
//	fmt.Println("here")
//	user := new(models.User)
//	if err := c.Bind(user); err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON data"})
//	}
//	fmt.Println(user.FirstName.String)
//	return c.String(http.StatusOK, "Update User")
//}
