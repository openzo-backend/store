package service

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/store/internal/models"
	"github.com/tanush-128/openzo_backend/store/internal/pb"
	"github.com/tanush-128/openzo_backend/store/internal/repository"
	"github.com/tanush-128/openzo_backend/store/internal/utils"
)

type StoreService interface {

	//CRUD
	CreateStore(ctx *gin.Context, req models.Store) (models.Store, error)
	GetStoreByID(ctx *gin.Context, id string) (models.Store, error)
	GetStoreBasicDetailsByID(ctx *gin.Context, id string) (repository.StoreBasicDetails, error)
	GetStoresByPincode(ctx *gin.Context, pincode string) ([]models.StorePublic, error)
	GetStoresByPincodeAndCategory(ctx *gin.Context, pincode string, category string) ([]models.StorePublic, error)
	GetStoresByPincodeAndSubCategory(ctx *gin.Context, pincode string, category string) ([]models.StorePublic, error)
	GetStoreByPhoneNo(ctx *gin.Context, phoneNo string) (models.Store, error)
	GetStoreByUserID(ctx *gin.Context, userID string) (models.Store, error)
	GetCategories(ctx *gin.Context) ([]string, error)
	UpdateStore(ctx *gin.Context, req models.Store) (models.Store, error)
}

type storeService struct {
	storeRepository repository.StoreRepository
	imageClient     pb.ImageServiceClient
	kafkaProducer   *kafka.Producer
}

func NewStoreService(storeRepository repository.StoreRepository,
	imageClient pb.ImageServiceClient, p *kafka.Producer,
) StoreService {
	return &storeService{storeRepository: storeRepository, imageClient: imageClient, kafkaProducer: p}
}

func (s *storeService) GetStoreByID(ctx *gin.Context, id string) (models.Store, error) {
	store, err := s.storeRepository.GetStoreByID(id)
	if err != nil {
		return models.Store{}, err
	}

	return store, nil
}

func (s *storeService) GetStoreBasicDetailsByID(ctx *gin.Context, id string) (repository.StoreBasicDetails, error) {
	store, err := s.storeRepository.GetStoreBasicDetailsByID(id)
	if err != nil {
		return repository.StoreBasicDetails{}, err
	}

	return store, nil
}

func (s *storeService) GetStoresByPincode(ctx *gin.Context, pincode string) ([]models.StorePublic, error) {

	return s.storeRepository.GetStoresByPincode(pincode)
}

func (s *storeService) GetStoreByPhoneNo(ctx *gin.Context, phoneNo string) (models.Store, error) {
	store, err := s.storeRepository.GetStoreByPhoneNo(phoneNo)
	if err != nil {
		return models.Store{}, err
	}

	return store, nil
}

func (s *storeService) GetStoresByPincodeAndCategory(ctx *gin.Context, pincode string, category string) ([]models.StorePublic, error) {
	stores, err := s.storeRepository.GetStoresByPincodeAndCategory(pincode, category)
	if err != nil {
		return []models.StorePublic{}, err
	}

	return stores, nil
}

func (s *storeService) GetStoresByPincodeAndSubCategory(ctx *gin.Context, pincode string, category string) ([]models.StorePublic, error) {
	stores, err := s.storeRepository.GetStoresByPincodeAndSubCategory(pincode, category)
	if err != nil {
		return []models.StorePublic{}, err
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

func (s *storeService) CreateStore(ctx *gin.Context, req models.Store) (models.Store, error) {

	err := ctx.Request.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		return models.Store{}, err
	}

	file, err := ctx.FormFile("image")
	if err == nil {

		imageBytes, err := utils.FileHeaderToBytes(file)
		if err != nil {
			return models.Store{}, err
		}

		Image, err := s.imageClient.UploadImage(ctx, &pb.ImageMessage{
			ImageData: imageBytes,
		})
		if err != nil {
			return models.Store{}, err
		}

		req.Image = Image.Url
	}
	createdStore, err := s.storeRepository.CreateStore(req)
	if err != nil {
		return models.Store{}, err // Propagate error
	}
	go writeStoreToKafka(createdStore, s.kafkaProducer)
	return createdStore, nil
}

func writeStoreToKafka(store models.Store, p *kafka.Producer) {
	// Write store to Kafka
	topic := "stores"
	storeJson, err := json.Marshal(store)
	if err != nil {
		return
	}

	p.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          storeJson,
			Key:            []byte(store.ID),
		}, nil)

	p.Flush(15 * 1000)

}

func (s *storeService) GetStoreByUserID(ctx *gin.Context, id string) (models.Store, error) {
	store, err := s.storeRepository.GetStoreByUserID(id)
	if err != nil {
		return models.Store{}, err
	}

	return store, nil
}

func (s *storeService) UpdateStore(ctx *gin.Context, req models.Store) (models.Store, error) {

	err := ctx.Request.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		return models.Store{}, err
	}

	store, err := s.storeRepository.GetStoreByID(req.ID)

	if err != nil {
		return models.Store{}, err
	}

	req.Image = store.Image
	req.Rating = store.Rating
	req.ReviewCount = store.ReviewCount
	req.SelfDeliveryService = store.SelfDeliveryService

	file, err := ctx.FormFile("image")
	if err == nil {

		imageBytes, err := utils.FileHeaderToBytes(file)
		if err != nil {
			return models.Store{}, err
		}

		Image, err := s.imageClient.UploadImage(ctx, &pb.ImageMessage{
			ImageData: imageBytes,
		})
		if err != nil {
			return models.Store{}, err
		}

		req.Image = Image.Url
	}

	updatedStore, err := s.storeRepository.UpdateStore(req)
	if err != nil {
		return models.Store{}, err
	}

	go writeStoreToKafka(updatedStore, s.kafkaProducer)

	return updatedStore, nil
}
