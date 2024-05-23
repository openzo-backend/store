package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/store/internal/models"
	"github.com/tanush-128/openzo_backend/store/internal/repository"
)

type ReviewService interface {
	CreateReview(ctx *gin.Context, req models.Review) (models.Review, error)
	GetReviewByID(ctx *gin.Context, id string) (models.Review, error)
	GetReviewsByStoreID(ctx *gin.Context, id string) ([]models.Review, error)
	GetReviewsByUserID(ctx *gin.Context, id string) ([]models.Review, error)
	UpdateReview(ctx *gin.Context, req models.Review) (models.Review, error)
	DeleteReview(ctx *gin.Context, id string) error
}

type reviewService struct {
	reviewRepository repository.ReviewRepository
	storeRepository  repository.StoreRepository
}

func NewReviewService(reviewRepository repository.ReviewRepository, storeRepository repository.StoreRepository) ReviewService {
	return &reviewService{
		reviewRepository: reviewRepository,
		storeRepository:  storeRepository,
	}
}

func (s *reviewService) CreateReview(ctx *gin.Context, req models.Review) (models.Review, error) {

	if req.Rating < 1 || req.Rating > 5 {
		return models.Review{}, errors.New("rating should be between 1 and 5")
	}

	exists, err := s.reviewRepository.CheckReviewExists(req.UserID, req.StoreID)
	if err != nil {
		return models.Review{}, err
	}

	if exists {
		return models.Review{}, errors.New("review already exists")
	}

	createdReview, err := s.reviewRepository.CreateReview(req)
	if err != nil {
		return models.Review{}, err
	}

	store, err := s.storeRepository.GetStoreByID(req.StoreID)
	if err != nil {
		return models.Review{}, err
	}

	store.Rating = store.Rating*float64(store.ReviewCount) + float64(req.Rating)
	store.ReviewCount = store.ReviewCount + 1
	store.Rating = store.Rating / float64(store.ReviewCount)

	_, err = s.storeRepository.UpdateStore(store)
	if err != nil {
		return models.Review{}, err
	}

	return createdReview, nil
}

func (s *reviewService) GetReviewByID(ctx *gin.Context, id string) (models.Review, error) {
	review, err := s.reviewRepository.GetReviewByID(id)
	if err != nil {
		return models.Review{}, err
	}
	return review, nil
}

func (s *reviewService) GetReviewsByStoreID(ctx *gin.Context, id string) ([]models.Review, error) {
	reviews, err := s.reviewRepository.GetReviewsByStoreID(id)
	if err != nil {
		return []models.Review{}, err
	}
	return reviews, nil
}

func (s *reviewService) GetReviewsByUserID(ctx *gin.Context, id string) ([]models.Review, error) {
	reviews, err := s.reviewRepository.GetReviewsByUserID(id)
	if err != nil {
		return []models.Review{}, err
	}
	return reviews, nil
}

func (s *reviewService) UpdateReview(ctx *gin.Context, req models.Review) (models.Review, error) {
	updatedReview, err := s.reviewRepository.UpdateReview(req)
	if err != nil {
		return models.Review{}, err
	}
	return updatedReview, nil
}

func (s *reviewService) DeleteReview(ctx *gin.Context, id string) error {
	err := s.reviewRepository.DeleteReview(id)
	if err != nil {
		return err
	}
	return nil
}
