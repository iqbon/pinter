package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
	"pinter/models/database"
)

type Subcategory struct {
	Id   int    `orm:"column(id_subcategory);pk"`
	Name string `orm:"column(name)"`
}

func (t *Subcategory) TableName() string {
	return "subcategory"
}

func init() {
	orm.RegisterModel(new(Subcategory))
}

// AddSubcategory insert a new Subcategory into database and returns
// last inserted Id on success.
func AddSubcategory(m *Subcategory) (id int64, err error) {
	database.DB.Create(m)
	database.DB.Save(m)
	return
}

// GetSubcategoryById retrieves Subcategory by Id. Returns error if
// Id doesn't exist
func GetSubcategoryById(id int) (v *Subcategory, err error) {
	v = &Subcategory{Id: id}
	database.DB.Where("id = ?", id).First(&v)
	return
}

// GetAllSubcategory retrieves all Subcategory matches certain condition. Returns empty list if
// no records exist
func GetAllSubcategory() (ml []Subcategory, err error) {
	var subcategories []Subcategory
	dbu := database.DB.New()

	dbu.Order("Id, Name").Find(&subcategories)
	ml = subcategories
	return
}

// UpdateSubcategory updates Subcategory by Id and returns error if
// the record to be updated doesn't exist
func UpdateSubcategoryById(mm *Subcategory) (err error) {
	m := Subcategory{}
	dbu := database.DB.New()
	dbu.Where("id = ?", mm.Id).First(&m)

	cate := 0
	if mm.Id != 0 {
		if mm.Name != "" {
			m.Name = mm.Name
			cate = 1
		}

		if cate == 1 {
			err := dbu.Save(&m).Error
			if err != nil {
				return errors.New("error database")
			}
		}

		return errors.New("ok")
	}

	return nil
}

// DeleteSubcategory deletes Subcategory by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSubcategory(id int) (err error) {
	m := Subcategory{}
	dbu := database.DB.New()
	dbu.Where("id = ?", id).Delete(&m)
	return
}
