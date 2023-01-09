package dao

import (
	"dbcourse/model"
	"dbcourse/utils"
	"fmt"
	"strconv"
)

// 插入一条供应商数据
func Insertsupplier(x *model.Supplier) error {
	sq := "insert into supplier4794(address,name,phone,head) values(?,?,?,?)"
	_, err := utils.Db.Exec(sq, x.Address, x.Name, x.Phone, x.Head)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	fmt.Printf("插入供应商成功")
	return nil
}

// 在数据库中查询到所有的供应商的信息
func Getallsupplier() []*model.Supplier {
	sq := "select id,address,name,phone,head from supplier4794"
	rows, _ := utils.Db.Query(sq)
	var res []*model.Supplier
	for rows.Next() {
		// var temp *model.Supplier
		temp := &model.Supplier{}
		rows.Scan(&temp.Id, &temp.Address, &temp.Name, &temp.Phone, &temp.Head)
		res = append(res, temp)
	}
	return res
}

func Getonesupplier(x int) (*model.Supplier, error) {
	sq := "select id,address,name,phone,head from supplier4794 where id = ?"
	row := utils.Db.QueryRow(sq, x)
	res := &model.Supplier{}
	err := row.Scan(&res.Id, &res.Address, &res.Name, &res.Phone, &res.Head)
	return res, err
}

// 根据id修改supplier的某一项
func Updatesupplier(x *model.Supplier) {
	sq := "update supplier4794 set address = ?,name = ?,phone = ?,head = ? where id = ?"
	_, err := utils.Db.Exec(sq, x.Address, x.Name, x.Phone, x.Head, x.Id)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("修改成功\n")
	fmt.Printf("x: %v\n", x)
}

// 根据id删除supplier的某一项
func Deletesupplier(x int) {
	sq := "delete from supplier4794 where id = ?"
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

// 获得切片包含所有的supplierid(供应商名)
func Getallsupplieridname() []string {
	var res []string
	supplierlist := Getallsupplier()
	for _, val := range supplierlist {
		temp := strconv.Itoa(val.Id)
		temp = temp + "(" + val.Name + ")"
		res = append(res, temp)
	}
	return res
}

// 获取供应商存货
func Getallsuppliergoods() []*model.Suppliergoods {
	sq := "select * from supplier_goods_count4794"
	rows, _ := utils.Db.Query(sq)
	var res []*model.Suppliergoods
	for rows.Next() {
		// var temp *model.Supplier
		temp := &model.Suppliergoods{}
		rows.Scan(&temp.Supplierid, &temp.Suppliername, &temp.Goodsspeciesamount, &temp.Allgoodsamount)
		res = append(res, temp)
	}
	return res
}
