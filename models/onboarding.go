package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
	"pinter/models/database"
)

type Onboarding struct {
	Id          int    `orm:"column(id_board);pk"`
	Image       string `orm:"column(image);null"`
	Description string `orm:"column(description);null"`
}

func (t *Onboarding) TableName() string {
	return "onboarding"
}

func init() {
	orm.RegisterModel(new(Onboarding))
}

// AddOnboarding insert a new Onboarding into database and returns
// last inserted Id on success.
func AddOnboarding(m *Onboarding) (id int64, err error) {
	database.DB.Create(m)
	database.DB.Save(m)
	return
}

// GetOnboardingById retrieves Onboarding by Id. Returns error if
// Id doesn't exist
func GetOnboardingById(id int) (v *Onboarding, err error) {
	v = &Onboarding{Id: id}
	database.DB.Where("id = ?", id).First(&v)
	return
}

// GetAllOnboarding retrieves all Onboarding matches certain condition. Returns empty list if
// no records exist
func GetAllOnboarding() (ml []Onboarding, err error) {
	var onboardings []Onboarding
	dbu := database.DB.New()

	dbu.Order("Id, Onboarding").Find(&onboardings)
	ml = onboardings
	return
}

// UpdateOnboarding updates Onboarding by Id and returns error if
// the record to be updated doesn't exist
func UpdateOnboardingById(mm *Onboarding) (err error) {
	m := Onboarding{}
	dbu := database.DB.New()
	dbu.Where("id = ?", mm.Id).First(&m)

	cate := 0
	if mm.Id != 0 {
		if mm.Image != "" {
			m.Image = mm.Image
			cate = 1
		}
		if mm.Description != "" {
			m.Description = mm.Description
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

// DeleteOnboarding deletes Onboarding by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOnboarding(id int) (err error) {
	m := Onboarding{}
	dbu := database.DB.New()
	dbu.Where("id = ?", id).Delete(&m)
	return
}
