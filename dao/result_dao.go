package dao

import (
	"log"

	"github.com/go-xorm/xorm"

	"../models"
)

type ResultDao struct {
	engine *xorm.Engine
}

func NewResultDao(engine *xorm.Engine) *ResultDao {
	return &ResultDao{engine: engine}
}

func (d *ResultDao) Get(id int) *models.Result {
	data := &models.Result{Id: id}

	ok, err := d.engine.Get(data)

	if ok && err == nil {
		return data
	} else {
		return nil
	}

}

func (d *ResultDao) GetAll() []models.Result {
	dataList := make([]models.Result, 0)

	err := d.engine.
		Desc("id").
		Find(&dataList)

	if err != nil {
		log.Println("Result_dao.GetAll error=", err)
		return dataList
	} else {
		return dataList
	}
}

func (d *ResultDao) CountAll() int64 {
	num, err := d.engine.Count(&models.Result{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *ResultDao) CountByGift(giftId int) int64 {
	num, err := d.engine.
		Where("gift_di=?", giftId).
		Count(&models.Result{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *ResultDao) CountByUser(uid int) int64 {
	num, err := d.engine.
		Where("uid=?", uid).
		Count(&models.Result{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

// 软删除
func (d *ResultDao) Delete(id int) error {
	data := &models.Result{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *ResultDao) Update(data *models.Result, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *ResultDao) Insert(data *models.Result) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *ResultDao) SearchByGift(giftId, page, size int) []models.Result {
	dataList := make([]models.Result, 0)

	err := d.engine.
		Where("gift_id=?", giftId).
		Limit(page, (page-1)*size).
		Desc("id").
		Find(&dataList)

	if err != nil {
		log.Println("Result_dao.GetAll error=", err)
		return dataList
	} else {
		return dataList
	}
}

func (d *ResultDao) SearchByUser(uid, page, size int) []models.Result {
	dataList := make([]models.Result, 0)

	err := d.engine.
		Where("uid=?", uid).
		Limit(page, (page-1)*size).
		Desc("id").
		Find(&dataList)

	if err != nil {
		log.Println("Result_dao.GetAll error=", err)
		return dataList
	} else {
		return dataList
	}
}
