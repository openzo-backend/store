package models

import "time"

type Store struct {
	StorePublic

	StorePrivate
}

type StorePublic struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Phone       string `json:"phone,omitempty" gorm:"unique" args:"private"`
	Image       string `json:"image"`
	Pincode     string `json:"pincode"`
	Location    string `json:"location"`
	Address     string `json:"address"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`

	//To be removed
	StoreType string `json:"store_type"`

	Category    string `json:"category" gorm:"default:store"`
	SubCategory string `json:"sub_category" gorm:"default:general_store'"`

	MetaDescription string `json:"meta_description,omitempty"`
	MetaTags        string `json:"meta_tags,omitempty"`

	Description string  `json:"description"`
	Rating      float64 `json:"rating" gorm:"default:0"`
	ReviewCount int     `json:"review_count" gorm:"default:0"`

	CreatedAt           time.Time `json:"created_at" gorm:"autoCreateTime"`
	SelfDeliveryService bool      `json:"self_delivery_service" gorm:"default:false"`
	DeliveryCharge      int       `json:"delivery_charge" gorm:"default:0"`
	PackagingCharge     int       `json:"packaging_charge" gorm:"default:0"`
	Busy                bool      `json:"busy" gorm:"default:false"`

	RestaurantDetails
	Ranking int `json:"ranking" gorm:"default:1"`
}

type StorePrivate struct {
	FCMToken  string `json:"fcm_token,omitempty" args:"private"`
	UserID    string `json:"user_id,omitempty" args:"private" gorm:"unique"`
	UserEmail string `json:"user_email,omitempty" args:"private" gorm:"unique"`

	DetailsComplete bool `json:"details_complete,omitempty" gorm:"default:false"`
	OnlineDiscovery bool `json:"online_discovery,omitempty" gorm:"default:true"`
}

type RestaurantDetails struct {
	// Cuisine         string `json:"cuisine,omitempty" gorm:"default:multi cuisine"`
	PrimaryCuisine   string `json:"primary_cuisine,omitempty"`
	SecondaryCuisine string `json:"secondary_cuisine,omitempty"`

	AvgPricePerPerson int `json:"avg_price_per_person,omitempty"`

	PureVeg bool `json:"pure_veg,omitempty"`
	Alcohol bool `json:"alcohol,omitempty"`

	TableCount         int  `json:"table_count,omitempty"`
	SeatingCapacity    int  `json:"seating_capacity,omitempty"`
	ReserveTableOnline bool `json:"reserve_table_online,omitempty" gorm:"default:false"`
}

type Review struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	StoreID   string    `json:"store_id" gorm:"size:36"`
	UserID    string    `json:"user_id" gorm:"size:36"`
	UserName  string    `json:"user_name"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	Review    string    `json:"review"`
}

type ResTable struct {
	ID        string `json:"id" gorm:"primaryKey"`
	StoreID   string `json:"store_id" gorm:"size:36"`
	Name      string `json:"table_name"`
	Seats     int    `json:"seats"`
	Available bool   `json:"available"`
	Section   string `json:"section"`
	Shape     string `json:"shape"`
}

//to store the views of the store per day
type View struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	StoreID   string    `json:"store_id" gorm:"size:36"`
	ViewCount int       `json:"view_count"`
	Date      time.Time `json:"date" gorm:"autoCreateTime"`
}
