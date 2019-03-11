package dao

import (
	"log"

	"github.com/go-xorm/xorm"

	"../models"
)

type CodeDao struct {
	engine *xorm.Engine
}

func NewCodeDao(engine *xorm.Engine) *CodeDao {
	return &CodeDao{engine: engine}
}

func (d *CodeDao) Get(id int) *models.Code {
	data := &models.Code{Id: id}

	ok, err := d.engine.Get(data)

	if ok && err == nil {
		return data
	} else {
		return nil
	}

}

func (d *CodeDao) GetAll() []models.Code {
	dataList := make([]models.Code, 0)

	err := d.engine.
		Desc("id").
		Find(&dataList)

	if err != nil {
		log.Println("code_dao.GetAll error=", err)
		return dataList
	} else {
		return dataList
	}
}

func (d *CodeDao) CountAll() int64 {
	num, err := d.engine.Count(&models.Code{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

// 软删除
func (d *CodeDao) Delete(id int) error {
	data := &models.Code{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *CodeDao) Update(data *models.Code, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *CodeDao) Insert(data *models.Code) error {
	_, err := d.engine.Insert(data)
	return err
}
