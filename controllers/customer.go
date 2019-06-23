package controllers

import (
	"fmt"
	"strings"
	"time"

	 _"github.com/astaxie/beego"
	"github.com/george518/PPGo_ApiAdmin/models"
)

type CustomerController struct {
	BaseController
}

//显示所有用户信息
func (self *CustomerController) Table() {
	//列表分页
	//page, err := self.GetInt("page")
	// if err != nil {
	// 	page = 1
	// }
	// limit, err := self.GetInt("limit")
	// if err != nil {
	// 	limit = 30
	// }

	//sourceList := sourceLists()

	//self.pageSize = limit
	result, count := models.FindCustomerAll()

	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["customerName"] = v.CustomerName
		row["cPhone"] = v.CPhone
		row["cAdr"] = v.CAdr
		row["buyDate"] = v.BuyDate
		row["cSum"] = v.CSum
		row["remark"] = v.Remark
		list[k] = row
	}

	self.ajaxList("成功", MSG_OK, count, list)
}

//添加用户信息
func (self *CustomerController) Add() {
	fmt.Println("add")
	self.display()
}

func (self *CustomerController) Delete(){
	fmt.Println("delete")
	Customer_id, _ := self.GetInt("id")
	_,err:=models.DelCustomer(Customer_id)
	if err !=nil {
		self.ajaxMsg("",MSG_OK)
	}else{
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	
}

//添加和修改用户信息ajax请求
func (self *CustomerController) AjaxSave() {
	fmt.Println("添加用户")
	//Customer_id, _ := self.GetInt("id")
	//if Customer_id == 0 {
	Customer := new(models.Customer)
	Customer.CustomerName = strings.TrimSpace(self.GetString("customer_name"))
	Customer.CAdr = strings.TrimSpace(self.GetString("c_adr"))
	Customer.CPhone = strings.TrimSpace(self.GetString("c_phone"))
	//Customer.CSum = self.GetFloat("c_sum")
	//Customer.Deposit = self.GetFloat("deposit")
	Customer.BuyDate = time.Now().Unix()
	Customer.UpdateDate = time.Now().Unix()
	Customer.Remark = strings.TrimSpace(self.GetString("remark"))
	_, err := models.AddCustomer(Customer)
	if err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
	//}
	// //修改
	// ApiDetail, _ := models.ApiDetailGetById(Api_id)
	// ApiDetail.SourceId, _ = self.GetInt("source_id")
	// ApiDetail.Id, _ = self.GetInt("id")
	// ApiDetail.Method, _ = self.GetInt("method")
	// ApiDetail.ApiName = strings.TrimSpace(self.GetString("api_name"))
	// ApiDetail.ApiUrl = strings.TrimSpace(self.GetString("api_url"))
	// ApiDetail.Detail = strings.TrimSpace(self.GetString("detail"))

	// ApiDetail.UpdateId = self.userId
	// ApiDetail.UpdateTime = time.Now().Unix()
	// ApiDetail.Status, _ = self.GetInt("status")

	// ApiDetail.Status = 1
	// if err := ApiDetail.Update(); err != nil {
	// 	self.ajaxMsg(err.Error(), MSG_ERR)
	// }
	// self.ajaxMsg("", MSG_OK)
}
