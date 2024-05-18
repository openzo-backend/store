<<<<<<< HEAD
package service

import (
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
	GetStoresByPincode(ctx *gin.Context, pincode string) ([]models.Store, error)
	GetStoresByPincodeAndCategory(ctx *gin.Context, pincode string, category string) ([]models.Store, error)
	GetStoreByPhoneNo(ctx *gin.Context, phoneNo string) (models.Store, error)
	GetCategories(ctx *gin.Context) ([]string, error)
	UpdateStore(ctx *gin.Context, req models.Store) (models.Store, error)
}

type storeService struct {
	storeRepository repository.StoreRepository
	imageClient     pb.ImageServiceClient
}

func NewStoreService(storeRepository repository.StoreRepository,
	imageClient pb.ImageServiceClient,
) StoreService {
	return &storeService{storeRepository: storeRepository, imageClient: imageClient}
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

	return createdStore, nil
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

	return updatedStore, nil
}
=======
package service

import (
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
	GetStoresByPincode(ctx *gin.Context, pincode string) ([]models.Store, error)
	GetStoresByPincodeAndCategory(ctx *gin.Context, pincode string, category string) ([]models.Store, error)
	GetStoreByPhoneNo(ctx *gin.Context, phoneNo string) (models.Store, error)
	GetCategories(ctx *gin.Context) ([]string, error)
	UpdateStore(ctx *gin.Context, req models.Store) (models.Store, error)
}

type storeService struct {
	storeRepository repository.StoreRepository
	imageClient     pb.ImageServiceClient
}

func NewStoreService(storeRepository repository.StoreRepository,
	imageClient pb.ImageServiceClient,
) StoreService {
	return &storeService{storeRepository: storeRepository, imageClient: imageClient}
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

	return createdStore, nil
}

func (s *storeService) UpdateStore(ctx *gin.Context, req models.Store) (models.Store, error) {

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

	updatedStore, err := s.storeRepository.UpdateStore(req)
	if err != nil {
		return models.Store{}, err
	}

	return updatedStore, nil
}
>>>>>>> 76ef4cf (kafka added)
