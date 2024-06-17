package main

import (
	"encoding/json"
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

	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
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

	conf := ReadConfig()
	p, _ := kafka.NewProducer(&conf)
	// topic := "notification"

	// go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

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
	storeRepository3 := repository.NewStoreRepository(db)
	reviewRepository := repository.NewReviewRepository(db)

	tableRespository := repository.NewTableRepository(db)

	StoreService := service.NewStoreService(storeRepository, imageClient, p)
	ReviewService := service.NewReviewService(reviewRepository, storeRepository3)

	TableService := service.NewTableService(tableRespository)

	go consumeKafka(storeRepository2, p)

	// Initialize HTTP server with Gin
	router := gin.Default()

	handler := handlers.NewHandler(&StoreService)
	reviewHandler := handlers.NewReviewHandler(&ReviewService)

	tableHandler := handlers.NewTableHandler(&TableService)

	router.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// router.Use(middlewares.JwtMiddleware(c))
	router.GET("/pincode/:pincode", handler.GetStoresByPincode)
	router.GET("/pincode/:pincode/category/:category", handler.GetStoresByPincodeAndCategory)
	router.GET("/getCategories", handler.GetCategories)

	router.GET("/reviews/:store_id", reviewHandler.GetReviewsByStoreID)
	router.GET("/review/:id", reviewHandler.GetReviewByID)

	router.POST("/table", tableHandler.CreateTable)
	router.GET("/table/:store_id", tableHandler.GetTablesByStoreID)
	router.PUT("/table", tableHandler.UpdateTable)
	router.DELETE("/table/:id", tableHandler.DeleteTable)

	router.Use(middlewares.NewMiddleware(c).JwtMiddleware)

	router.GET("/:id", handler.GetStoreByID)
	router.GET("/user/:user_id", handler.GetStoreByUserID)
	router.POST("/", handler.CreateStore)
	router.PUT("/:id", handler.UpdateStore)

	router.POST("/review", reviewHandler.CreateReview)
	router.GET("/reviews/user/:user_id", reviewHandler.GetReviewsByUserID)
	router.PUT("/review/:id", reviewHandler.UpdateReview)
	router.DELETE("/review/:id", reviewHandler.DeleteReview)

	// router.Use(middlewares.JwtMiddleware)

	router.Run(fmt.Sprintf(":%s", cfg.HTTPPort))

}

type Notification struct {
	Message  string `json:"message"`
	FCMToken string `json:"fcm_token"`
}

func consumeKafka(storeRepo repository.StoreRepository, notificationProducer *kafka.Producer) {
	conf := ReadConfig()

	topic := "sales"

	// sets the consumer group ID and offset
	conf["group.id"] = "go-group-1"
	conf["auto.offset.reset"] = "earliest"

	// creates a new consumer and subscribes to your topic
	consumer, _ := kafka.NewConsumer(&conf)
	consumer.SubscribeTopics([]string{topic}, nil)
	var order struct {
		StoreId     string `json:"store_id"`
		OrderStatus string `json:"status"`
		Type        string `json:"type"`
	}
	run := true
	for run {
		// consumes messages from the subscribed topic and prints them to the console
		e := consumer.Poll(1000)
		switch ev := e.(type) {
		case *kafka.Message:
			// application-specific processing
			err := json.Unmarshal(ev.Value, &order)
			if err != nil {
				fmt.Println("Error unmarshalling JSON: ", err)
			}
			fmt.Println("Order received: ", order)

			if order.OrderStatus == "placed" && order.Type == "online_order" {
				fcm, err := storeRepo.GetFCMTokenByStoreID(order.StoreId)
				if err != nil {
					fmt.Println("Error getting FCM token: ", err)
				}
				fmt.Println("FCM token: ", fcm)

				notificationMsg, _ := json.Marshal(Notification{
					Message:  "You have a new order",
					FCMToken: fcm,
				})
				notificationTopic := "notification"

				// send a notification to the store
				notificationProducer.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{Topic: &notificationTopic, Partition: kafka.PartitionAny},
					Value:          notificationMsg,
				}, nil)

				notificationProducer.Flush(15 * 1000)
			}

		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", ev)
			run = false
		}
	}

	// closes the consumer connection
	consumer.Close()

}
