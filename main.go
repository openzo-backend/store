package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tanush-128/openzo_backend/store/config"
	handlers "github.com/tanush-128/openzo_backend/store/internal/api"
	"github.com/tanush-128/openzo_backend/store/internal/middlewares"
	"github.com/tanush-128/openzo_backend/store/internal/pb"
	"github.com/tanush-128/openzo_backend/store/internal/repository"
	"github.com/tanush-128/openzo_backend/store/internal/service"
	"google.golang.org/grpc"
)

var UserClient pb.UserServiceClient

type User2 struct {
}

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to load config: %w", err))
	}

	db, err := connectToDB(cfg) // Implement database connection logic
	if err != nil {
		log.Fatal(fmt.Errorf("failed to connect to database: %w", err))
	}

	// // Initialize gRPC server
	// grpcServer := grpc.NewServer()
	// Storepb.RegisterStoreServiceServer(grpcServer, service.NewGrpcStoreService(StoreRepository, StoreService))
	// reflection.Register(grpcServer) // Optional for server reflection

	//Initialize gRPC client
	conn, err := grpc.Dial(cfg.UserGrpc, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)
	UserClient = c

	imageConn, err := grpc.Dial(cfg.ImageGrpc, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer imageConn.Close()
	imageClient := pb.NewImageServiceClient(imageConn)

	storeRepository := repository.NewStoreRepository(db)
	storeRepository2 := repository.NewStoreRepository(db)
	StoreService := service.NewStoreService(storeRepository, imageClient)

	go service.GrpcServer(cfg, &service.Server{
		StoreRepository: storeRepository2,
	})

	// Initialize HTTP server with Gin
	router := gin.Default()
	handler := handlers.NewHandler(&StoreService)

	router.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// router.Use(middlewares.JwtMiddleware(c))
	router.POST("/", handler.CreateStore)
	router.GET("/:id", handler.GetStoreByID)
	router.GET("/pincode/:pincode", handler.GetStoresByPincode)
	router.GET("/pincode/:pincode/category/:category", handler.GetStoresByPincodeAndCategory)
	router.GET("/phone/:phone_no", handler.GetStoreByPhoneNo)
	router.GET("/getCategories", handler.GetCategories)
	router.Use(middlewares.NewMiddleware(c).JwtMiddleware)
	router.GET("/:id", handler.GetStoreByID)
	router.PUT("/:id", handler.UpdateStore)

	// router.Use(middlewares.JwtMiddleware)

	router.Run(fmt.Sprintf(":%s", cfg.HTTPPort))

}
