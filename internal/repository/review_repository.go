package repository

import (
	"github.com/tanush-128/openzo_backend/store/internal/models"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	CreateReview(Review models.Review) (models.Review, error)
	GetReviewByID(id string) (models.Review, error)
	GetReviewsByStoreID(id string) ([]models.Review, error)
	GetReviewsByUserID(id string) ([]models.Review, error)
	UpdateReview(Review models.Review) (models.Review, error)
	CheckReviewExists(userId string, storeId string) (bool, error)
	DeleteReview(id string) error

	// Add more methods for other Stor e operations (GetStoreByEmail, UpdateStore, etc.)

}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {

	return &reviewRepository{db: db}
}

func (r *reviewRepository) CreateReview(Review models.Review) (models.Review, error) {
	tx := r.db.Create(&Review)

	if tx.Error != nil {
		return models.Review{}, tx.Error
	}

	return Review, nil
}

func (r *reviewRepository) GetReviewByID(id string) (models.Review, error) {
	var Review models.Review
	tx := r.db.Where("id = ?", id).First(&Review)
	if tx.Error != nil {
		return models.Review{}, tx.Error
	}

	return Review, nil
}

func (r *reviewRepository) GetReviewsByStoreID(id string) ([]models.Review, error) {
	var Reviews []models.Review
	tx := r.db.Where("store_id = ?", id).Find(&Reviews)
	if tx.Error != nil {
		return []models.Review{}, tx.Error
	}
	return Reviews, nil
}

func (r *reviewRepository) GetReviewsByUserID(id string) ([]models.Review, error) {
	var Reviews []models.Review
	tx := r.db.Where("user_id = ?", id).Find(&Reviews)
	if tx.Error != nil {
		return []models.Review{}, tx.Error
	}
	return Reviews, nil
}

func (r *reviewRepository) UpdateReview(Review models.Review) (models.Review, error) {
	tx := r.db.Save(&Review)

	if tx.Error != nil {
		return models.Review{}, tx.Error
	}

	return Review, nil
}

func (r *reviewRepository) CheckReviewExists(userId string, storeId string) (bool, error) {
	var count int64
	tx := r.db.Model(&models.Review{}).Where("user_id = ? AND store_id = ?", userId, storeId).Count(&count)

	if tx.Error != nil {
		return false, tx.Error
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (r *reviewRepository) DeleteReview(id string) error {
	tx := r.db.Delete(&models.Review{}, id)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
