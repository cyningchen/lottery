package dao

import (
	"log"

	"github.com/go-xorm/xorm"

	"../models"
)

type BlackUserDao struct {
	engine *xorm.Engine
}

func NewBlackUserDao(engine *xorm.Engine) *BlackUserDao {
	return &BlackUserDao{engine: engine}
}

func (d *BlackUserDao) Get(id int) *models.BlackUser {
	data := &models.BlackUser{Id: id}

	ok, err := d.engine.Get(data)

	if ok && err == nil {
		return data
	} else {
		return nil
	}

}

func (d *BlackUserDao) GetAll() []models.BlackUser {
	dataList := make([]models.BlackUser, 0)

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

func (d *BlackUserDao) CountAll() int64 {
	num, err := d.engine.Count(&models.BlackUser{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

// 软删除
func (d *BlackUserDao) Delete(id int) error {
	data := &models.BlackUser{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *BlackUserDao) Update(data *models.BlackUser, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *BlackUserDao) Insert(data *models.BlackUser) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *BlackUserDao) GetByUid(uid int) *models.BlackUser {
	dataList := make([]models.BlackUser, 0)
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
