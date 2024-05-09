package service

import (
	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/store/internal/models"
)

func (s *storeService) GetStoreByID(ctx *gin.Context, id string) (models.Store, error) {
	store, err := s.storeRepository.GetStoreByID(id)
	if err != nil {
		return models.Store{}, err
	}

	return store, nil
}

func (s *storeService) GetStoresByPincode(ctx *gin.Context, pincode string) ([]models.Store, error) {
	stores, err := s.storeRepository.GetStoresByPincode(pincode)
	if err != nil {
		return []models.Store{}, err
	}

	return stores, nil
}

func (s *storeService) GetStoreByPhoneNo(ctx *gin.Context, phoneNo string) (models.Store, error) {
	store, err := s.storeRepository.GetStoreByPhoneNo(phoneNo)
	if err != nil {
		return models.Store{}, err
	}

	return store, nil
}

func (s *storeService) GetStoresByPincodeAndCategory(ctx *gin.Context, pincode string, category string) ([]models.Store, error) {
	stores, err := s.storeRepository.GetStoresByPincodeAndCategory(pincode, category)
	if err != nil {
		return []models.Store{}, err
	}

	return stores, nil
}

func (s *storeService) GetCategories(ctx *gin.Context) ([]string, error) {
	categories, err := s.storeRepository.GetCategories()
	if err != nil {
		return []string{}, err
	}

	return categories, nil
}
