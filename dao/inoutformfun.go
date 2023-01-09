package dao

import (
	"dbcourse/model"
	"dbcourse/utils"
	"fmt"
	"math/rand"
	"time"
)

// 随机生成单号
func Proformnum() string {
	var res string
	rand.Seed(time.Now().Unix())
	for i := 1; i <= 8; i++ {
		c := '0'
		c += rune(rand.Intn(10))
		res = res + string(c)
	}
	return res
}

// 插入表单
func Insertinoutform(x *model.Inoutform) {
	sq := "insert into inoutform4794(formnumber,opstaffid,goodsid,amount,price,inorout,datee) values(?,?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sq, x.Formnumber, x.Opstaffid, x.Goodsid, x.Amount, x.Price, x.Inorout, x.Datee)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("添加表单成功")
}

// 获取所有表单
func Getallform() []*model.Inoutform {
	sq := "select formnumber,opstaffid,goodsid,amount,price,inorout,datee from inoutform4794"
	rows, _ := utils.Db.Query(sq)
	var res []*model.Inoutform
	for rows.Next() {
		// var temp *model.Goods
		temp := &model.Inoutform{}
		rows.Scan(&temp.Formnumber, &temp.Opstaffid, &temp.Goodsid, &temp.Amount, &temp.Price, &temp.Inorout, &temp.Datee)
		res = append(res, temp)
	}
	return res
}

// 根据单号获取某一条
func Getoneform(x string) (*model.Inoutform, error) {
	sq := "select formnumber,opstaffid,goodsid,amount,price,inorout,datee from inoutform4794 where formnumber = ?"
	row := utils.Db.QueryRow(sq, x)
	res := &model.Inoutform{}
	err := row.Scan(&res.Formnumber, &res.Opstaffid, &res.Goodsid, &res.Amount, &res.Price, &res.Inorout, &res.Datee)
	return res, err
}

// 根据id获得表单数据
func Getallformbystaffid(x int) []*model.Inoutform {
	sq := "select formnumber,opstaffid,goodsid,amount,price,inorout,datee from inoutform4794 where opstaffid = ?"
	fmt.Printf("sq: %v\n", sq)
	fmt.Printf("x: %v\n", x)
	rows, _ := utils.Db.Query(sq, x)
	var res []*model.Inoutform
	for rows.Next() {
		// var temp *model.Goods
		temp := &model.Inoutform{}
		rows.Scan(&temp.Formnumber, &temp.Opstaffid, &temp.Goodsid, &temp.Amount, &temp.Price, &temp.Inorout, &temp.Datee)
		res = append(res, temp)
	}
	return res
}

// 根据单号修改form的某一项
func Updateform(x *model.Inoutform) {
	sq := "update inoutform4794 set formnumber = ?,opstaffid = ?,goodsid = ?,amount = ?,price = ?,inorout = ?,datee = ? where formnumber = ?"
	_, err := utils.Db.Exec(sq, x.Formnumber, x.Opstaffid, x.Goodsid, x.Amount, x.Price, x.Inorout, x.Datee, x.Formnumber)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("修改成功\n")
	fmt.Printf("x: %v\n", x)
}

// 根据单号删除form的某一项
func Deleteform(x string) {
	sq := "delete from inoutform4794 where formnumber = ?"
	res, err := utils.Db.Exec(sq, x)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("删除了%d行\n", row)
}
