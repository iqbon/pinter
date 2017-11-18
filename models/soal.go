package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"pinter/models/database"

	"github.com/astaxie/beego/orm"
)

type Soal struct {
	Id            int    `orm:"column(id);pk;auto"`
	IdCategory    int64  `orm:"column(id_category)"`
	Soal          string `orm:"column(soal)"`
	JawabanA      string `orm:"column(jawaban_a)"`
	JawabanB      string `orm:"column(jawaban_b)"`
	JawabanC      string `orm:"column(jawaban_c)"`
	JawabanD      string `orm:"column(jawaban_d)"`
	JawabanE      string `orm:"column(jawaban_e)"`
	IdJawaban     int64  `orm:"column(id_jawaban)"`
	IdSubcategory int64  `orm:"column(id_subcategory)"`
}

func (t *Soal) TableName() string {
	return "soal"
}

func init() {
	orm.RegisterModel(new(Soal))
}

// AddSoal insert a new Soal into database and returns
// last inserted Id on success.
func AddSoal(m *Soal) (id int64, err error) {
	database.DB.Create(m)
	database.DB.Save(m)
	return
}

// GetSoalById retrieves Soal by Id. Returns error if
// Id doesn't exist
func GetSoalById(id int) (v *Soal, err error) {
	o := orm.NewOrm()
	v = &Soal{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetSoalByCategory(id int) (v []*Soal, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Soal))
	var questions []*Soal
	_, err = qs.Filter("id_category", id).All(&questions)
	if err == nil {
		return nil, err
	}
	return questions, nil

}

// GetAllSoal retrieves all Soal matches certain condition. Returns empty list if
// no records exist
func GetAllSoal(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Soal))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Soal
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateSoal updates Soal by Id and returns error if
// the record to be updated doesn't exist
func UpdateSoalById(m *Soal) (err error) {
	o := orm.NewOrm()
	v := Soal{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSoal deletes Soal by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSoal(id int) (err error) {
	o := orm.NewOrm()
	v := Soal{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Soal{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
