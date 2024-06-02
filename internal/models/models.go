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

var storeToProductCategory = map[StoreType][]string{
	StoreTypeGeneralStores: []string{
		"household_essentials",
		"cleaning_supplies",
		"snacks_and_beverages",
		"personal_care_products",
		"pet_care_items",
		"office_supplies",
		"toys_and_games",
		"kitchenware",
		"seasonal_items",
		"health_and_wellness_products",
	},
	StoreTypeStationery: []string{
		"writing_instruments",
		"notebooks_and_pads",
		"art_supplies",
		"desk_accessories",
		"calendars_and_planners",
		"filing_and_organization",
		"crafting_materials",
		"school_supplies",
		"presentation_tools",
	},
	StoreTypeElectronics: []string{
		"smartphones_and_accessories",
		"laptops_and_computers",
		"audio_devices",
		"cameras_and_photography_equipment",
		"home_appliances",
		"gaming_consoles_and_accessories",
		"wearable_technology",
		"cables_and_connectors",
		"batteries_and_chargers",
		"virtual_reality_products",
	},
	StoreTypeFashion: []string{
		"men’s_clothing",
		"women’s_clothing",
		"kids’_clothing",
		"shoes_and_footwear",

		"accessories",
		"jewelery",
		"handbags_and_purses",
		"sunglasses",
		"watches",
		"formal_wear",
	},
	StoreTypeFootwear: []string{
		"men’s_shoes",
		"women’s_shoes",

		"athletic_footwear",
		"casual_shoes",
		"boots_and_booties",
		"sandals",
		"slippers",
		"kids’_shoes",

		"work_and_safety_shoes",
		"specialty_footwear",
	},
	StoreTypeAccessories: []string{
		"handbags_and_clutches",
		"jewelry",
		"hats_and_caps",
		"scarves_and_wraps",

		"sunglasses",

		"watches",
		"belts",
		"gloves_and_mittens",
		"hair_accessories",

		"tech_accessories",
	},

	StoreTypeRental: []string{
		"event_equipment",

		"tools_and_machinery",
		"party_supplies",

		"formal_attire_rentals",
		"outdoor_gear",
		"furniture_rentals",
		"home_appliance_rentals",

		"vehicle_rentals",
		"audio-visual_equipment",
		"recreational_equipment",
	},
	StoreTypeGrocery: {
		"fresh_produce",
		"dairy_and_eggs",
		"meat_and_seafood",

		"canned_and_packaged_goods",
		"bakery_items",

		"frozen_foods",
		"beverages",
		"snacks_and_sweets",

		"organic_and_natural_products",
		"international_foods",
	},
	StoreTypeBeauty: []string{
		"skincare",
		"makeup",
		"haircare",

		"fragrances",

		"bath_and_body_products",
		"nail_care",

		"men’s_grooming_products",
		"beauty_tools_and_accessories",
		"wellness_products",

		"specialty_beauty_items",
	},
	StoreTypeSports: []string{
		"athletic_apparel",
		"footwear_for_sports",
		"exercise_equipment",

		"team_sports_gear",
		"outdoor_and_adventure_gear",
		"yoga_and_fitness_accessories",

		"swimwear",
		"sports_nutrition",
		"recovery_and_injury_prevention",

		"camping_and_hiking_equipment",
	},
}

type Store struct {
	StorePublic

	StorePrivate
}

type StorePublic struct {
	ID          string    `json:"id" gorm:"primaryKey"` // The ID field is also the UUID for the store
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	Pincode     string    `json:"pincode"`
	StoreType   StoreType `json:"store_type"`
	Location    string    `json:"location"`
	Address     string    `json:"address"`
	OpeningTime string    `json:"opening_time"`
	ClosingTime string    `json:"closing_time"`

	Rating      float64 `json:"rating" gorm:"default:0"`
	Description string  `json:"description"`
	ReviewCount int     `json:"review_count" gorm:"default:0"`

	CreatedAt           time.Time `json:"created_at" gorm:"autoCreateTime"`
	SelfDeliveryService bool      `json:"self_delivery_service" gorm:"default:false"`
}

type StorePrivate struct {
	FCMToken        string `json:"fcm_token,omitempty" args:"private"`
	UserID          string `json:"user_id,omitempty" args:"private" gorm:"unique"`
	UserEmail       string `json:"user_email,omitempty" args:"private" gorm:"unique"`
	Phone           string `json:"phone,omitempty" gorm:"unique" args:"private"`
	DetailsComplete bool   `json:"details_complete,omitempty" gorm:"default:false"`
	OnlineDiscovery bool   `json:"online_discovery,omitempty" gorm:"default:true"`
}

type Review struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	StoreID   string    `json:"store_id"`
	UserID    string    `json:"user_id"`
	UserName  string    `json:"user_name"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	Review    string    `json:"review"`
}

type RestaurantDetails struct {
	ID              string `json:"id" gorm:"primaryKey"`
	StoreID         string `json:"store_id" gorm:"unique"`
	Cuisine         string `json:"cuisine" gorm:"default:multi cuisine"`
	PureVeg         bool   `json:"pure_veg" gorm:"default:false"`
	Alcohol         bool   `json:"alcohol" gorm:"default:false"`
	SeatingType     string `json:"seating_type" gorm:"default:indoor"`
	TableCount      int    `json:"table_count"`
	SeatingCapacity int    `json:"seating_capacity"`
}

type Restaurant struct {
	Store
	RestaurantDetails
}

type ResTable struct {
	ID        string `json:"id" gorm:"primaryKey"`
	StoreID   string `json:"store_id"`
	Name      string `json:"table_name"`
	Seats     int    `json:"seats"`
	Available bool   `json:"available"`
	Section   string `json:"section"`
	Shape     string `json:"shape"`
}
