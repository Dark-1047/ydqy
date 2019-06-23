package models

import (
	"fmt"

	_ "github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Customer struct {
	Id           int     //客户ID
	CustomerName string  //客户姓名
	CPhone       string  //客户电话
	CAdr         string  //客户地址
	CProductId   int     //产品信息ID
	BuyDate      int64   //购买时间
	CSum         float64 //产品总额
	Deposit      float64 //定金
	Tail         float64 //尾款
	Remark       string  //备注
	CreateDate   int64   //创建时间
	UpdateDate   int64   //修改时间
}

func (c *Customer) TableName() string {
	return "customer"
}

//查询所有用户信息
func FindCustomerAll() ([]*Customer, int64) {
	list := make([]*Customer, 0)
	query := orm.NewOrm().QueryTable("customer")
	query.OrderBy("-id").All(&list)
	total, _ := query.Count()
	fmt.Println("查询客户信息数量：%d", total)
	return list, total
}

func AddCustomer(c *Customer) (int64, error) {
	return orm.NewOrm().Insert(c)
}

func DelCustomer(int id) (int64, error) {
	
	query := orm.NewOrm().QueryTable("customer")
	return query.Filter("id", id).Delete()
}
