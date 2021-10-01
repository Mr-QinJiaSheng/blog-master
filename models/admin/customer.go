package admin

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Customer struct {
	Id        int       `orm:"column(id);auto" description:"主键"`
	Uid       string    `orm:"column(uid);size(50)" description:"用户ID"`
	Username  string    `orm:"column(username);size(255);null" description:"用户名"`
	Password  string    `orm:"column(password);size(255);null" description:"密码"`
	Nickname  string    `orm:"column(nickname);size(255);null" description:"昵称"`
	Image     string    `orm:"column(image);size(255);null" description:"头像"`
	Url       string    `orm:"column(url);size(255);null" description:"博客地址"`
	Signature string    `orm:"column(signature);size(255);null" description:"个性签名"`
	Email     string    `orm:"column(email);size(50);null" description:"邮箱"`
	Phone     string    `orm:"column(phone);size(50);null" description:"电话"`
	Wishlist  int       `orm:"column(wishlist);null" description:"收藏"`
	Review    int       `orm:"column(review);null" description:"评论"`
	Like      int       `orm:"column(like);null" description:"点赞"`
	Status    int       `orm:"column(status);null" description:"1可用，2禁用，0删除"`
	Integral  int       `orm:"column(integral);null" description:"积分"`
	Fans      int       `orm:"column(fans);null" description:"粉丝数量"`
	Focus     int       `orm:"column(focus);null" description:"关注数量"`
	Created   time.Time `orm:"column(created);type(datetime);null" description:"创建时间"`
	Updated   time.Time `orm:"column(updated);type(datetime);null" description:"修改时间"`
	IsFans    bool      `orm:"-"`
}

func (t *Customer) TableName() string {
	return "customer"
}

func init() {
	orm.RegisterModel(new(Customer))
}

// AddCustomer insert a new Customer into database and returns
// last inserted Id on success.
func AddCustomer(m *Customer) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCustomerById retrieves Customer by Id. Returns error if
// Id doesn't exist
func GetCustomerById(id int) (v *Customer, err error) {
	o := orm.NewOrm()
	v = &Customer{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCustomer retrieves all Customer matches certain condition. Returns empty list if
// no records exist
func GetAllCustomer(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Customer))
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

	var l []Customer
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

// UpdateCustomer updates Customer by Id and returns error if
// the record to be updated doesn't exist
func UpdateCustomerById(m *Customer) (err error) {
	o := orm.NewOrm()
	v := Customer{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		v.Nickname = m.Nickname
		v.Username = m.Username
		v.Phone = m.Phone
		v.Image = m.Image
		v.Signature = m.Signature
		v.Url = m.Url
		if num, err = o.Update(&v); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCustomer deletes Customer by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCustomer(id int) (err error) {
	o := orm.NewOrm()
	v := Customer{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Customer{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func CustomerLogin(username, password string) (*Customer, bool) {
	o := orm.NewOrm()
	var (
		customer Customer
		err      error
	)
	ok := false
	o.Using("default")
	cond := orm.NewCondition()
	cond = cond.And("status", 1).And("Username", username).Or("Email", username).Or("Uid", username).Or("Phone", username)
	qs := o.QueryTable(&customer)
	qs = qs.SetCond(cond)
	if err = qs.One(&customer); err == nil {
		if customer.Password == password {

			updateTime := customer.Updated.Local().Format("2006-01-02")
			integral := customer.Integral

			if updateTime != time.Now().Format("2006-01-02") {

				// 如果是当天第一次登录则修改积分
				integral += 10
			}

			o := orm.NewOrm()
			v := Customer{Id: customer.Id}
			// ascertain id exists in the database
			if err = o.Read(&v); err == nil {
				//var num int64
				v.Updated = time.Now()
				v.Integral = integral
				if _, err = o.Update(&v); err != nil {
					return nil, false
				}
			}

			ok = true
		}
	}
	return &customer, ok
}

func UpdateIntegral(uid, integral int) bool {
	o := orm.NewOrm()
	v := Customer{Id: uid}
	// ascertain id exists in the database
	if err := o.Read(&v); err == nil {
		//var num int64
		//v.Updated = time.Now()
		v.Integral = v.Integral + integral
		if _, err = o.Update(&v); err != nil {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

func SetFans(uid int, flag bool) bool {
	o := orm.NewOrm()
	v := Customer{Id: uid}
	// ascertain id exists in the database
	if err := o.Read(&v); err == nil {
		//var num int64
		//v.Updated = time.Now()
		if flag == true {
			v.Fans = v.Fans + 1
		} else {
			v.Fans = v.Fans - 1
		}
		if _, err = o.Update(&v); err != nil {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

func SetFocus(uid int, flag bool) bool {
	o := orm.NewOrm()
	v := Customer{Id: uid}
	// ascertain id exists in the database
	if err := o.Read(&v); err == nil {
		//var num int64
		//v.Updated = time.Now()
		if flag == true {
			v.Focus = v.Focus + 1
		} else {
			v.Focus = v.Focus - 1
		}
		if _, err = o.Update(&v); err != nil {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}
