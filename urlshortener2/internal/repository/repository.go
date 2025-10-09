package repository

import (
	"urlshortner/internal/models"

	"gorm.io/gorm"
)

type URLRepository interface {
}

type urlRepository struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) URLRepository {
	return &urlRepository{db: db}
}

func (r *urlRepository) Create(url *models.URL) error {
	return r.db.Create(url).Error
}

func (r *urlRepository) FindByShortCode(code string) (*models.URL, error) {
	var url models.URL
	if err := r.db.Where("short_code = ? and is_active = ?", code, true).Error; err != nil {
		return &models.URL{}, err
	}
	return &url, nil
}

func (r *urlRepository) FindByURL(urlstr string) (*models.URL, error) {
	var url models.URL
	if err := r.db.Where("original_url = ? and is_active = ?", urlstr, true).Error; err != nil {
		return nil, err
	}
	return &url, nil
}

func (r *urlRepository) IncrementCounter(code string) error {
	return r.db.Model(&models.URL{}).Where("short_code = ?", code).
		Update("click_count", gorm.Expr("click_count + ?", 1)).Error
}

func (r *urlRepository) Delete(code string) error {
	return r.db.Model(&models.URL{}).Where("short_code = ?", code).
		Update("is_active", false).Error
}
