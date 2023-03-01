package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserInfo struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	OSName      string             `json:"os_name" bson:"os_name"`
	OSVersion   string             `json:"os_version" bson:"os_version"`
	BrowserName string             `json:"browser_name" bson:"browser_name"`
	BrowserVer  string             `json:"browser_ver" bson:"browser_ver"`
	IPAddress   string             `json:"ip_address" bson:"ip_address"`
	PhoneBrand  string             `json:"phone_brand" bson:"phone_brand"`
	PhoneModel  string             `json:"phone_model" bson:"phone_model"`
	ScreenRes   string             `json:"screen_res" bson:"screen_res"`
}

type TopResult struct {
	ID    string `bson:"_id"`
	Count int    `bson:"count"`
}
