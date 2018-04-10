package entity

import (
	"gopkg.in/mgo.v2/bson"
	"gin_rest/model/geoJson"
)

//请求
type (
	PassengerLbsReq struct {
		PassengerId int     `json:"passengerId"`
		OrderId     int     `json:"orderId"`
		Longitude   float64 `json:"longitude"`
		Latitude    float64 `json:"latitude"`
		OrderState  int     `json:"orderState"`
	}

	OrderPathReq struct {
		PassengerId int    `json:"passengerId"`
		OrderId     int    `json:"orderId"`
		DriverId    int    `json:"driverId"`
		ClientId    string `json:"clientId"`
		Area        int    `json:"area"`
	}
)

//topic内容格式
//orders/1/1/{orderid}
type (
	OrderPathTopic struct {
		OrderId    int                 `json:"orderId"`
		StartPoint geoJson.Coordinate  `json:"startPoint"`
		EndPoint   geoJson.Coordinate  `json:"endPoint"`
		Points      geoJson.GeoPointCollection `json:"multiPoint"`
	}
)




//

type (
	OrderPath struct {
		Id          bson.ObjectId      `bson:"_id,omitempty" json:"id"`

		PassengerId int                `json:"passengerId" bson:"passenger_id"`
		DriverId    int                `json:"driverId" bson:"driver_id"`
		OrderId    int                 `json:"orderId" bson:"order_id"`
		StartPoint geoJson.Coordinate  `json:"startPoint" bson:"start_point"`
		EndPoint   geoJson.Coordinate  `json:"endPoint" bson:"end_point"`
		Points      geoJson.GeoPointCollection `json:"multiPoint" bson:"points"` 
		ClientId    string             `json:"clientId" bson:"client_id"`
	}

	PassengerLbs struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"clientId"`
		ClientId    string        `bson:"client_id"`
		PassengerId int           `bson:"passenger_id" json:"passenger_id"`
		Locations   geoJson.Point `bson:"locations" json:"locations"`
		OrderId     int           `bson:"order_id" json:"order_id"`
		OrderState  int           `bson:"order_state" json:"order_state"`
	}
)
