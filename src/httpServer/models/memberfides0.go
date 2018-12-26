package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
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

func (This *Memberfides0) GetMemberfidesList(start, end int) map[string]interface{} {
	mem := make(map[string]interface{})

	orm.NewOrm().Raw("select * from " + This.TableName() + " limit " + strconv.Itoa(start) + "," + strconv.Itoa(end)).QueryRow(&mem)
	return mem
}
