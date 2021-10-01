package admin

import (
	"github.com/astaxie/beego/orm"
)

type Fans struct {
	Id         int `orm:"column(id);auto" description:"主键"`
	CustomerId int `orm:"column(customer_id)" description:"用户ID 关注者 ID"`
	FansId     int `orm:"column(fans_id)" description:"用户ID 被关注者ID"`
}

func (t *Fans) TableName() string {
	return "fans"
}

func init() {
	orm.RegisterModel(new(Fans))
}

// 关注用户customer_id关注这,fans_id被关注者
func AddFocus(customer_id, fans_id int) (int, bool) {
	o := orm.NewOrm()

	flag := o.QueryTable(new(Fans)).Filter("customer_id", customer_id).Filter("fans_id", fans_id).Exist()
	if !flag {
		o.Begin()
		_, err := o.Insert(&Fans{
			CustomerId: customer_id,
			FansId:     fans_id,
		})

		if err != nil {
			o.Rollback()
			return 0, false
		}

		if SetFans(fans_id, true) && SetFocus(customer_id, true) {
			o.Commit()
			return 1, true
		}
		o.Rollback()
		return 0, false
	} else {
		return unsetFans(customer_id, fans_id)
	}
}

func unsetFans(customer_id, fans_id int) (int, bool) {
	o := orm.NewOrm()
	o.Begin()
	_, err := o.QueryTable(new(Fans)).Filter("customer_id", customer_id).Filter("fans_id", fans_id).Delete()

	if err != nil {
		o.Rollback()
		return 0, false
	}

	if SetFans(fans_id, false) && SetFocus(customer_id, false) {
		o.Commit()
		return -1, true
	}
	o.Rollback()
	return 0, false
}

func IsFans(customer_id, fans_id int) bool {
	o := orm.NewOrm()
	flag := o.QueryTable(new(Fans)).Filter("customer_id", customer_id).Filter("fans_id", fans_id).Exist()
	return flag
}

func GetFocus(customer_id int) []Customer {
	o := orm.NewOrm()
	var focus []*Fans
	o.QueryTable(new(Fans)).Filter("customer_id", customer_id).All(&focus)

	var customer []Customer
	for _, v := range focus {
		c := Customer{Id: v.FansId}
		err := o.Read(&c)
		if err == nil {
			customer = append(customer, c)
		}
	}
	return customer
}

func GetFans(customer_id int) []Customer {
	o := orm.NewOrm()
	var fans []*Fans
	o.QueryTable(new(Fans)).Filter("fans_id", customer_id).All(&fans)

	var customer []Customer
	for _, v := range fans {
		c := Customer{Id: v.CustomerId}
		err := o.Read(&c)
		if err == nil {
			customer = append(customer, c)
		}
	}
	return customer
}
