package models

import (
	"errors"

	"pinter/models/database"

	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id   int    `orm:"column(id_category);pk"`
	Name string `orm:"column(name)"`
}

func (t *Category) TableName() string {
	return "category"
}

func init() {
	orm.RegisterModel(new(Category))
}

// AddCategory insert a new Category into database and returns
// last inserted Id on success.
func AddCategory(m *Category) (id int64, err error) {
	database.DB.Create(m)
	database.DB.Save(m)
	return
}

// GetCategoryById retrieves Category by Id. Returns error if
// Id doesn't exist
func GetCategoryById(id int) (v *Category, err error) {
	v = &Category{Id: id}
	database.DB.Where("id = ?", id).First(&v)
	return
}

// GetAllCategory retrieves all Category matches certain condition. Returns empty list if
// no records exist
func GetAllCategory() (ml []Category, err error) {
	var categories []Category
	dbu := database.DB.New()

	dbu.Order("Id, Name").Find(&categories)
	ml = categories
	return
}

// UpdateCategory updates Category by Id and returns error if
// the record to be updated doesn't exist
func UpdateCategoryById(mm *Category) (err error) {
	m := Category{}
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

// DeleteCategory deletes Category by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCategory(id int) (err error) {
	m := Category{}
	dbu := database.DB.New()
	dbu.Where("id = ?", id).Delete(&m)
	return
}
