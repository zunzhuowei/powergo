package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Memberfides0 struct {
	Mid          int `orm:"column(mid);pk"` // 设置主键
	Name         string
	Btd          string
	City         string
	Icon         string
	Invite       int
	Gp           int
	Sta          int
	Mtime        int64
	Email        string
	Bindtime     int64
	Tel          string
	Realname     string
	Identity     string
	Yellowvip    int
	Isyearvip    string
	Passwd       string
	Salt         string
	Lgtm         int
	Lxlg         int
	Qq           string
	Address      string
	Regip        int64
	ActiveAssets int
}

func init() {
	orm.RegisterModel(new(Memberfides0))
}

const (
	MemberfidesTableName string = "memberfides0"
)

func (This Memberfides0) TableName() string {
	return MemberfidesTableName
}

func (This Memberfides0) GetMemberfidesList(start, end int) []Memberfides0 {
	var memberfides []Memberfides0

	o := orm.NewOrm()

	_, err := o.Raw("select * from "+This.TableName()+" limit ?,?", start, end).QueryRows(&memberfides)
	if err != nil {
		beego.Error("GetMemberfidesList err ---::", err)
	}

	o1 := orm.NewOrm()
	o1.Using("log")

	var lists []orm.ParamsList
	sql := "select * from gold_log order by id desc limit ?,?"
	num, err := o1.Raw(sql, start, end).ValuesList(&lists)
	if err == nil && num > 0 {
		/*for k, v := range lists {
			fmt.Println("vvvvvvvvvvvvvvvvvvvvvv --::", k, v)
		}*/
	} else {
		fmt.Println(err.Error())
	}

	return memberfides
}

/*func (This Memberfides0) GetMemberfidesList(start, end int) []orm.ParamsList {
	var lists []orm.ParamsList

	num, err := orm.NewOrm().Raw("select * from "+This.TableName()+" limit ?,?", start, end).ValuesList(&lists)
	if err == nil && num > 0 {
		fmt.Println(lists[0][0]) // slene
	}

	return lists
}*/
