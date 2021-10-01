package admin

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Review struct {
	Id        int
	Name      string
	Review    string    `orm:"size(500)"`
	Reply     string    `orm:"size(500)"`
	Site      string    `orm:"size(500)"`
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now;type(datetime)"`
	Status    int       `orm:"default(1)"`
	ArticleId int
	Customer  *Customer `orm:"rel(fk)"`
	Like      int
	Star      int
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Review))
}

func AddReview(m []Review, aid int) error {

	o := orm.NewOrm()
	err := o.Begin()
	// 更新评论数量
	article := Article{Id: aid}
	o.Read(&article)
	article.Review = article.Review + len(m)
	o.Update(&article)

	num, err := o.InsertMulti(len(m), m)

	fmt.Println(num)
	if err != nil {
		err = o.Rollback()

	} else {
		err = o.Commit()
	}

	return err
}
