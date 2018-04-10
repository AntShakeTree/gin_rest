package service

import (
	"gin_rest/entity"
	"gin_rest/common"
	"gopkg.in/mgo.v2/bson"
	"gin_rest/model/geoJson"
	"gin_rest/repositories"
)

type LbsService struct {
	b repositories.IBaseRepository
}

func CreateInstance(f repositories.IRepositoryFactory) LbsService {
	lss := LbsService{}
	lss.b = f.CreateRepository()
	return lss
}

//
func (reps *LbsService) CreatePassengerLbs(pareq *entity.PassengerLbsReq) {
	pa := entity.PassengerLbs{}
	pa.PassengerId = pareq.PassengerId
	//随机生成避免重复
	pa.ClientId = bson.NewObjectId().String()
	pa.Locations = *geoJson.NewPoint(geoJson.Coordinate{geoJson.Coord(pareq.Longitude), geoJson.Coord(pareq.Latitude)})
	reps.b.Insert(&pa)
}

//
func (reps *LbsService) CreateOrderPath(paq entity.OrderPathReq) {
	var entity = entity.OrderPath{}
	entity.OrderId = paq.OrderId
	entity.ClientId = common.GetBrokerId()
	entity.PassengerId = paq.PassengerId
	entity.DriverId = paq.DriverId
	entity.Points = *geoJson.NewGeoPointCollection(nil)
	reps.b.Insert(&entity)
}

//
func (reps *LbsService) GetOrderPath(orderid int) entity.OrderPath {
	//
	var entity entity.OrderPath
	reps.b.GetByQuery(bson.M{"order_id": orderid}, &entity)
	return entity
}
func (reps *LbsService) FindOrderPathAll() []entity.OrderPath {
	//
	var results []entity.OrderPath
	reps.b.FindAll(&results)
	return results
}
func (reps *LbsService) UpdateByOrderId(topic entity.OrderPathTopic) {
	entity := reps.GetOrderPath(topic.OrderId)

	ps := &entity.Points
	ps.Appends(topic.Points)
	reps.b.Update(bson.D{{"order_id", topic.OrderId}}, &entity)
}
