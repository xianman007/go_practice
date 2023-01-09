package dao

import (
	"dbcourse/model"
	"dbcourse/utils"
	"fmt"
	"strconv"
)

// 插入一条员工
func Insertstaff(x *model.Staff) {
	sq := "insert into staff4794(age,name,username,password,phone,sex) values(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sq, x.Age, x.Name, x.Username, x.Password, x.Phone, x.Sex)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("添加员工成功")
}

// 在数据库中查询到所有的供应商的信息
func Getallstaff() []*model.Staff {
	sq := "select staffid,age,name,username,password,phone,sex from staff4794"
	rows, _ := utils.Db.Query(sq)
	var res []*model.Staff
	for rows.Next() {
		// var temp *model.Staff
		temp := &model.Staff{}
		rows.Scan(&temp.Staffid, &temp.Age, &temp.Name, &temp.Username, &temp.Password, &temp.Phone, &temp.Sex)
		res = append(res, temp)
	}
	return res
}

// 根据id找到某一条员工的数据
func Getonestaff(x int) (*model.Staff, error) {
	sq := "select staffid,age,name,username,password,phone,sex from staff4794 where staffid = ?"
	row := utils.Db.QueryRow(sq, x)
	res := &model.Staff{}
	err := row.Scan(&res.Staffid, &res.Age, &res.Name, &res.Username, &res.Password, &res.Phone, &res.Sex)
	return res, err
}

// 根据id修改staff的某一项
func Updatestaff(x *model.Staff) {
	sq := "update staff4794 set age = ?,name = ?,username = ?,password = ?,phone = ?,sex = ? where staffid = ?"
	_, err := utils.Db.Exec(sq, &x.Age, &x.Name, &x.Username, &x.Password, &x.Phone, &x.Sex, x.Staffid)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("修改成功\n")
	fmt.Printf("x: %v\n", x)
}

// 根据id删除staff的某一项
func Deletestaff(x int) {
	sq := "delete from staff4794 where staffid = ?"
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

// 获得切片包含所有的staffid(员工名名)
func Getallstaffidname() []string {
	var res []string
	stafflist := Getallstaff()
	for _, val := range stafflist {
		temp := strconv.Itoa(val.Staffid)
		temp = temp + "(" + val.Name + ")"
		res = append(res, temp)
	}
	return res
}

// 获取员工绩效
func Getallstaffperformance() []*model.Staffperformance {
	sq := "select * from staff_performance_view4794"
	rows, _ := utils.Db.Query(sq)
	var res []*model.Staffperformance
	for rows.Next() {
		// var temp *model.Staff
		temp := &model.Staffperformance{}
		rows.Scan(&temp.Staffid, &temp.Staffname, &temp.Opcount, &temp.Allgoodsamount)
		res = append(res, temp)
	}
	return res
}
