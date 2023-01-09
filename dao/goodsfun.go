package dao

import (
	"dbcourse/model"
	"dbcourse/utils"
	"fmt"
	"strconv"
)

// 插入货物
func Insertgoods(x *model.Goods) {
	sq := "insert into goods4794(supplierid,price,amount,name) values(?,?,?,?)"
	_, err := utils.Db.Exec(sq, x.Supplierid, x.Price, x.Amount, x.Name)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("添加商品成功")
}

// 在数据库中查询到所有的货物的信息
func Getallgoods() []*model.Goods {
	sq := "select id,supplierid,price,amount,name from goods4794"
	rows, _ := utils.Db.Query(sq)
	var res []*model.Goods
	for rows.Next() {
		// var temp *model.Goods
		temp := &model.Goods{}
		rows.Scan(&temp.Id, &temp.Supplierid, &temp.Price, &temp.Amount, &temp.Name)
		res = append(res, temp)
	}
	return res
}

// 根据id获取到某一项货物信息
func Getonegoods(x int) (*model.Goods, error) {
	sq := "select id,supplierid,price,amount,name from goods4794 where id = ?"
	row := utils.Db.QueryRow(sq, x)
	res := &model.Goods{}
	err := row.Scan(&res.Id, &res.Supplierid, &res.Price, &res.Amount, &res.Name)
	return res, err
}

// 根据id修改goods的某一项
func Updategoods(x *model.Goods) {
	sq := "update goods4794 set supplierid = ?,price = ?,amount = ?,name = ? where id = ?"
	_, err := utils.Db.Exec(sq, x.Supplierid, x.Price, x.Amount, x.Name, x.Id)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("修改成功\n")
	fmt.Printf("x: %v\n", x)
}

// 根据id删除goods的某一项
func Deletegoods(x int) {
	sq := "delete from goods4794 where id = ?"
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

// 获得切片包含所有的goodsid(商品名)
func Getallgoodsidname() []string {
	var res []string
	goodslist := Getallgoods()
	for _, val := range goodslist {
		temp := strconv.Itoa(val.Id)
		temp = temp + "(" + val.Name + ")"
		res = append(res, temp)
	}
	return res
}
