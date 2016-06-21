package utils

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //导入驱动
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dbcon := "saltfish:saltfish@tcp(114.215.145.118:3306)/hello?charset=utf8&loc=Asia%2fShanghai"
	orm.RegisterDataBase("default", "mysql", dbcon, 30)
}

// Result 数据返回格式
type Result struct {
	Code int
	Msg  string
	Data interface{}
}

// ReturnSuccess 返回成功结果
func ReturnSuccess(data interface{}) Result {
	return Result{Code: 0, Msg: "success", Data: data}
}

// ReturnFailure 返回失败结果
func ReturnFailure(code int) Result {
	return Result{Code: code, Msg: "failed", Data: nil}
}
