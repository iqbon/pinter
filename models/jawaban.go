package models

import (
	"errors"

	"pinter/models/database"

	"github.com/astaxie/beego/orm"
)

type Jawaban struct {
	Id        int    `orm:"column(id);pk;auto"`
	IdJawaban int64  `orm:"column(id_jawaban);null"`
	IdSoal    int64  `orm:"column(id_soal);null"`
	Jawaban   string `orm:"column(jawaban);null"`
}

func (t *Jawaban) TableName() string {
	return "jawaban"
}

func init() {
	orm.RegisterModel(new(Jawaban))
}

// AddJawaban insert a new Jawaban into database and returns
// last inserted Id on success.
func AddJawaban(m *Jawaban) (id int64, err error) {
	database.DB.Create(m)
	database.DB.Save(m)
	return
}

// GetJawabanById retrieves Jawaban by Id. Returns error if
// Id doesn't exist
func GetJawabanById(id int) (v *Jawaban, err error) {
	v = &Jawaban{Id: id}
	database.DB.Where("id = ?", id).First(&v)
	return
}

// GetAllJawaban retrieves all Jawaban matches certain condition. Returns empty list if
// no records exist
func GetAllJawaban() (ml []Jawaban, err error) {
	var jawabans []Jawaban
	dbu := database.DB.New()

	dbu.Order("Id, Jawaban").Find(&jawabans)
	ml = jawabans
	return
}

// UpdateJawaban updates Jawaban by Id and returns error if
// the record to be updated doesn't exist
func UpdateJawabanById(mm *Jawaban) (err error) {
	m := Jawaban{}
	dbu := database.DB.New()
	dbu.Where("id = ?", mm.Id).First(&m)

	cate := 0
	if mm.Id != 0 {
		if mm.IdSoal != 0 {
			m.IdSoal = mm.IdSoal
			cate = 1
		}
		if mm.Jawaban != "" {
			m.Jawaban = mm.Jawaban
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

// DeleteJawaban deletes Jawaban by Id and returns error if
// the record to be deleted doesn't exist
func DeleteJawaban(id int) (err error) {
	m := Jawaban{}
	dbu := database.DB.New()
	dbu.Where("id = ?", id).Delete(&m)
	return
}
