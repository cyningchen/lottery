package dao

import (
	"github.com/go-xorm/xorm"
	"log"

	"../models"
)

type UserDayDao struct {
	engine *xorm.Engine
}

func NewUserDayDao(engine *xorm.Engine) *UserDayDao {
	return &UserDayDao{engine: engine}
}

func (d *UserDayDao) Get(id int) *models.UserDay {
	data := &models.UserDay{Id: id}

	ok, err := d.engine.Get(data)

	if ok && err == nil {
		return data
	} else {
		return nil
	}

}

func (d *UserDayDao) GetAll() []models.UserDay {
	dataList := make([]models.UserDay, 0)

	err := d.engine.
		Desc("id").
		Find(&dataList)

	if err != nil {
		log.Println("black_user_dao.GetAll error=", err)
		return dataList
	} else {
		return dataList
	}
}

func (d *UserDayDao) CountAll() int64 {
	num, err := d.engine.Count(&models.UserDay{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

// 软删除
func (d *UserDayDao) Delete(id int) error {
	data := &models.UserDay{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *UserDayDao) Update(data *models.UserDay, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *UserDayDao) Insert(data *models.UserDay) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *UserDayDao) GetByUid(uid int) *models.UserDay {
	dataList := make([]models.UserDay, 0)
	err := d.engine.Where("uid=?", uid).
		Desc("id").
		Limit(1).
		Find(&dataList)

	if err != nil || len(dataList) <= 1 {
		return nil
	} else {
		return &dataList[0]
	}
}

func (d *UserDayDao) Search(uid int, day string) *models.UserDay {
	dataList := make([]models.UserDay, 0)
	err := d.engine.Where("uid=?", uid).
		Where("day=", day).
		Limit(1).
		Find(&dataList)
	if err != nil || len(dataList) == 1 {
		return nil
	} else {
		return &dataList[0]
	}
}
