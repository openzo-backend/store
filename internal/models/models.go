package models

import "time"

type StoreType string

const (
	StoreTypeGeneralStores StoreType = "general_store"
	StoreTypeStationery    StoreType = "stationery"
	StoreTypeElectronics   StoreType = "electronics"
	StoreTypeFashion       StoreType = "fashion"
	StoreTypeFootwear      StoreType = "footwear"
	StoreTypeAccessories   StoreType = "accessories"
	StoreTypeRental        StoreType = "rental"
	StoreTypeGrocery       StoreType = "grocery"
	StoreTypeBeauty        StoreType = "beauty"
	StoreTypeSports        StoreType = "sports"
)

type Store struct {
	StorePublic

	StorePrivate
}

type StorePublic struct {
	ID          string `json:"id" gorm:"primaryKey"` // The ID field is also the UUID for the store
	Name        string `json:"name"`
	Image       string `json:"image"`
	Pincode     string `json:"pincode"`
	Location    string `json:"location"`
	Address     string `json:"address"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`

	//To be removed
	StoreType StoreType `json:"store_type"`

	Category    string `json:"category" gorm:"default:store"`
	SubCategory string `json:"sub_category" gorm:"default:general_store'"`

	Description string  `json:"description"`
	Rating      float64 `json:"rating" gorm:"default:0"`
	ReviewCount int     `json:"review_count" gorm:"default:0"`

	CreatedAt           time.Time `json:"created_at" gorm:"autoCreateTime"`
	SelfDeliveryService bool      `json:"self_delivery_service" gorm:"default:false"`

	RestaurantDetails
}

type StorePrivate struct {
	FCMToken        string `json:"fcm_token,omitempty" args:"private"`
	UserID          string `json:"user_id,omitempty" args:"private" gorm:"unique"`
	UserEmail       string `json:"user_email,omitempty" args:"private" gorm:"unique"`
	Phone           string `json:"phone,omitempty" gorm:"unique" args:"private"`
	DetailsComplete bool   `json:"details_complete,omitempty" gorm:"default:false"`
	OnlineDiscovery bool   `json:"online_discovery,omitempty" gorm:"default:true"`
}

type RestaurantDetails struct {
	Cuisine         string `json:"cuisine,omitempty" gorm:"default:multi cuisine"`
	PureVeg         bool   `json:"pure_veg,omitempty" gorm:"default:false"`
	Alcohol         bool   `json:"alcohol,omitempty" gorm:"default:false"`
	SeatingType     string `json:"seating_type,omitempty" gorm:"default:indoor"`
	TableCount      int    `json:"table_count,omitempty"`
	SeatingCapacity int    `json:"seating_capacity,omitempty"`
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
