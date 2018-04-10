package repositories

type IRepositoryFactory interface {
	CreateRepository() IBaseRepository
}

//

type OrderPathRepositoryFactory struct {
}

//
func (factory *OrderPathRepositoryFactory) CreateRepository() IBaseRepository {
	re := &OrderPathRepository{}
	re.Create("order_path")
	return re

}

//
type PassengerLBSRepositoryFactory struct {
}

//
func (factory *PassengerLBSRepositoryFactory) CreateRepository() IBaseRepository {
	re := &PassengerLBSRepository{}
	re.Create("passenger_lbs")
	return re
}

//
func BuilderOrderPathRepositoryFactory() *OrderPathRepositoryFactory {
	return &OrderPathRepositoryFactory{}
}


