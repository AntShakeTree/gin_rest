package repositories

import (
	"gopkg.in/mgo.v2"
	"gin_rest/common"
)

type BaseRepository struct {
	C       *mgo.Collection
	Name    string
	Context *common.Context
}
//订单
type OrderPathRepository struct {
	BaseRepository
}
//乘客LBS定位
type PassengerLBSRepository struct {
	BaseRepository
}
/**
 * 操作Mongo的基础仓库
 */
type IBaseRepository interface{
	CreateContext()
	Insert(docs interface{})
	UpdateId(id interface{}, data interface{})
	Get(id interface{}, res interface{})
	GetByQuery(vs map[string]interface{}, res interface{})
	Update(vs interface{},data interface{})
	FindAll(res interface{})
}

func (reps *BaseRepository) CreateContext() {
	reps.Context = common.NewContext()
	col := reps.Context.DbCollection(reps.Name)
	reps.C = col
}


//必须是引用类型
func (reps *BaseRepository) Insert(docs interface{}) {
	defer reps.Context.Close()
	reps.C.Insert(docs)
}

func (reps *BaseRepository) UpdateId(id interface{}, data interface{}) {
	defer reps.Context.Close()
	reps.C.UpdateId(id, data)
}
func (reps *BaseRepository) Get(id interface{}, res interface{}) {
	qu := reps.C.FindId(id)
	qu.One(res)
}

//
func (reps *BaseRepository) GetByQuery(vs map[string]interface{}, res interface{}) {
	qu := reps.C.Find(vs)
	qu.One(res)
}

//
func (reps *BaseRepository) FindAll(res interface{}) {
	defer reps.Context.Close()
	reps.C.Find(nil).All(res)

}

//
func (reps *BaseRepository) Create(name string)  {
	reps.Name=name
	reps.CreateContext()
}
//修改
func (reps *BaseRepository) Update(vs interface{},data interface{})  {
	defer reps.Context.Close()
	reps.C.Update(vs,data)
}