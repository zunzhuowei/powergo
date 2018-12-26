package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Memberfides0 struct {
	Mid          int
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

const (
	MemberfidesTableName string = "memberfides0"
)

func (This *Memberfides0) TableName() string {
	return MemberfidesTableName
}

func (This *Memberfides0) GetMemberfidesList(start, end int) []Memberfides0 {
	var memberfides []Memberfides0

	num, err := orm.NewOrm().Raw("select * from "+This.TableName()+" limit ?,?", start, end).QueryRows(&memberfides)
	if err == nil {
		fmt.Println("user nums: ", num)
	}

	return memberfides
}

/*func (This *Memberfides0) GetMemberfidesList(start, end int) []orm.ParamsList {
	var lists []orm.ParamsList

	num, err := orm.NewOrm().Raw("select * from "+This.TableName()+" limit ?,?", start, end).ValuesList(&lists)
	if err == nil && num > 0 {
		fmt.Println(lists[0][0]) // slene
	}

	return lists
}*/
