package repository

import (
	"log"

	"github.com/google/uuid"

	"github.com/tanush-128/openzo_backend/store/internal/models"

	"gorm.io/gorm"
)

type StoreBasicDetails struct {
	ID    string `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Image string `json:"image"`

	Address string `json:"address"`

	Category    string `json:"category" gorm:"default:store"`
	SubCategory string `json:"sub_category" gorm:"default:general_store'"`

	Description string  `json:"description"`
	Rating      float64 `json:"rating" gorm:"default:0"`
	ReviewCount int     `json:"review_count" gorm:"default:0"`
}

type StoreRepository interface {
	CreateStore(Store models.Store) (models.Store, error)
	GetStoreByID(id string) (models.Store, error)
	GetStoreBasicDetailsByID(id string) (StoreBasicDetails, error)

	GetStoreByEmail(email string) (models.Store, error)
	GetStoreByUserID(userID string) (models.Store, error)
	GetStoreByPhoneNo(phoneNo string) (models.Store, error)
	GetStoresByPincode(pincode string) ([]models.StorePublic, error)
	GetStoresByPincodeAndCategory(pincode string, category string) ([]models.StorePublic, error)
	GetStoresByPincodeAndSubCategory(pincode string, category string) ([]models.StorePublic, error)
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

func (r *storeRepository) GetStoreBasicDetailsByID(id string) (StoreBasicDetails, error) {
	var Store StoreBasicDetails
	tx := r.db.Model(models.Store{}).Where("id = ?", id).First(&Store)
	if tx.Error != nil {
		return StoreBasicDetails{}, tx.Error
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

func (r *storeRepository) GetStoreByUserID(userID string) (models.Store, error) {
	var Store models.Store
	tx := r.db.Where("user_id = ?", userID).First(&Store)
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

func (r *storeRepository) GetStoresByPincode(pincode string) ([]models.StorePublic, error) {
	var Stores []models.StorePublic
	tx := r.db.Model(&models.Store{}).
		Where("pincode = ?", pincode).
		Order("ranking ASC").
		Find(&Stores)
	if tx.Error != nil {
		return []models.StorePublic{}, tx.Error
	}

	return Stores, nil
}

func (r *storeRepository) GetStoresByPincodeAndCategory(pincode string, category string) ([]models.StorePublic, error) {
	var Stores []models.StorePublic

	tx := r.db.Model(&models.Store{}).Where("category = ? AND pincode = ?", category, pincode).Order("ranking ASC").Find(&Stores)
	if tx.Error != nil {
		return []models.StorePublic{}, tx.Error
	}

	return Stores, nil
}

func (r *storeRepository) GetStoresByPincodeAndSubCategory(pincode string, category string) ([]models.StorePublic, error) {
	var Stores []models.StorePublic

	tx := r.db.Model(&models.Store{}).Where("sub_category = ? AND pincode = ?", category, pincode).Order("ranking ASC").Find(&Stores)
	if tx.Error != nil {
		return []models.StorePublic{}, tx.Error
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

func (r *storeRepository) UpdateStore(store models.Store) (models.Store, error) {
	// Retrieve the existing store record from the database
	var existingStore models.Store
	if err := r.db.First(&existingStore, store.ID).Error; err != nil {
		return models.Store{}, err
	}

	// For integer fields
	if store.Ranking == 0 {
		existingStore.Ranking = 1
		// existingStore.Ranking = store.Ranking
	}

	// Save the updated store back to the database
	if err := r.db.Save(&store).Error; err != nil {
		return models.Store{}, err
	}

	return existingStore, nil
}

// Implement other repository methods (GetStoreByID, GetStoreByEmail, UpdateStore, etc.) with proper error handling
