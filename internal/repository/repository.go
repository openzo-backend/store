package repository

import (
	"log"

	"github.com/google/uuid"
	"github.com/tanush-128/openzo_backend/store/internal/models"

	"gorm.io/gorm"
)

type StoreRepository interface {
	CreateStore(Store models.Store) (models.Store, error)
	GetStoreByID(id string) (models.Store, error)
	GetStoreByEmail(email string) (models.Store, error)
	GetStoreByPhoneNo(phoneNo string) (models.Store, error)
	GetStoresByPincode(pincode string) ([]models.Store, error)
	GetStoresByPincodeAndCategory(pincode string, category string) ([]models.Store, error)
	GetCategories() ([]string, error)
	GetFCMTokenByStoreID(storeID string) (string, error)
	UpdateStore(Store models.Store) (models.Store, error)
	// Add more methods for other Store operations (GetStoreByEmail, UpdateStore, etc.)

}

type storeRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) StoreRepository {

	return &storeRepository{db: db}
}

func (r *storeRepository) CreateStore(Store models.Store) (models.Store, error) {
	Store.ID = uuid.New().String()
	log.Println("Phone no", Store.Phone)
	tx := r.db.Create(&Store)

	if tx.Error != nil {
		return models.Store{}, tx.Error
	}

	return Store, nil
}

func (r *storeRepository) GetStoreByID(id string) (models.Store, error) {
	var Store models.Store
	tx := r.db.Where("id = ?", id).First(&Store)
	if tx.Error != nil {
		return models.Store{}, tx.Error
	}

	return Store, nil
}

func (r *storeRepository) GetStoreByEmail(email string) (models.Store, error) {
	var Store models.Store
	tx := r.db.Where("email = ?", email).First(&Store)
	if tx.Error != nil {
		return models.Store{}, tx.Error
	}

	return Store, nil
}

func (r *storeRepository) GetStoreByPhoneNo(phoneNo string) (models.Store, error) {
	var Store models.Store
	tx := r.db.Where("phone = ?", phoneNo).First(&Store)
	if tx.Error != nil {
		return models.Store{}, tx.Error
	}

	return Store, nil
}

func (r *storeRepository) GetStoresByPincode(pincode string) ([]models.Store, error) {
	var Stores []models.Store
	tx := r.db.Where("pincode = ?", pincode).Find(&Stores)
	if tx.Error != nil {
		return []models.Store{}, tx.Error
	}

	return Stores, nil
}

func (r *storeRepository) GetStoresByPincodeAndCategory(pincode string, category string) ([]models.Store, error) {
	var Stores []models.Store
	tx := r.db.Find(&Stores, "store_type = ? AND pincode = ?", category, pincode)
	if tx.Error != nil {
		return []models.Store{}, tx.Error
	}

	return Stores, nil
}

func (r *storeRepository) GetCategories() ([]string, error) {
	var categories []string
	tx := r.db.Model(&models.Store{}).Distinct("store_type").Pluck("store_type", &categories)
	if tx.Error != nil {
		return []string{}, tx.Error
	}

	return categories, nil
}

func (r *storeRepository) GetFCMTokenByStoreID(storeID string) (string, error) {
	var fcmToken string
	tx := r.db.Model(&models.Store{}).Where("id = ?", storeID).Pluck("fcm_token", &fcmToken)
	if tx.Error != nil {
		return "", tx.Error
	}

	return fcmToken, nil
}

func (r *storeRepository) UpdateStore(Store models.Store) (models.Store, error) {
	tx := r.db.Save(&Store)
	if tx.Error != nil {
		return models.Store{}, tx.Error
	}

	return Store, nil
}

// Implement other repository methods (GetStoreByID, GetStoreByEmail, UpdateStore, etc.) with proper error handling
