package dao

import (
	"log"

	"github.com/go-xorm/xorm"

	"../models"
)

type BlackIpDao struct {
	engine *xorm.Engine
}

func NewBlackIpDao(engine *xorm.Engine) *BlackIpDao {
	return &BlackIpDao{engine: engine}
}

func (d *BlackIpDao) Get(id int) *models.BlackIp {
	data := &models.BlackIp{Id: id}

	ok, err := d.engine.Get(data)

	if ok && err == nil {
		return data
	} else {
		return nil
	}

}

func (d *BlackIpDao) GetAll() []models.BlackIp {
	dataList := make([]models.BlackIp, 0)

	err := d.engine.
		Desc("id").
		Find(&dataList)

	if err != nil {
		log.Println("black_ip_dao.GetAll error=", err)
		return dataList
	} else {
		return dataList
	}
}

func (d *BlackIpDao) CountAll() int64 {
	num, err := d.engine.Count(&models.BlackIp{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

// 软删除
func (d *BlackIpDao) Delete(id int) error {
	data := &models.BlackIp{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *BlackIpDao) Update(data *models.BlackIp, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *BlackIpDao) Insert(data *models.BlackIp) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *BlackIpDao) GetByIp(ip string) *models.BlackIp {
	dataList := make([]models.BlackIp, 0)
	err := d.engine.Where("ip=?", ip).
		Desc("id").
		Limit(1).
		Find(&dataList)

	if err != nil || len(dataList) <= 1 {
		return nil
	} else {
		return &dataList[0]
	}
}
