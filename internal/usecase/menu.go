package usecase

import (
	"coffee-project/internal/models"
	"coffee-project/internal/repository"
)

type MenuUsecase interface {
	FetchMenu() ([]models.MenuItem, error)
	AddMenuItem(input models.CreateMenuItemInput) (models.MenuItem, error)
}

type menuUsecase struct {
	menuRepo repository.MenuRepository
}

func NewMenuUsecase(repo repository.MenuRepository) MenuUsecase {
	return &menuUsecase{menuRepo: repo}
}

func (u *menuUsecase) FetchMenu() ([]models.MenuItem, error) {
	return u.menuRepo.GetAll()
}

func (u *menuUsecase) AddMenuItem(input models.CreateMenuItemInput) (models.MenuItem, error) {
	// Buisness logic: if no image provided, set default image URL
	imgURL := input.ImageURL
	if imgURL == "" {
		imgURL = "/images/noimage.png" // Default image URL
	}

	item := models.MenuItem{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Volume:      input.Volume,
		ImageURL:    imgURL,
		CategoryID:  input.CategoryID,
		IsAvailable: true,
	}

	err := u.menuRepo.Create(&item)
	return item, err
}
