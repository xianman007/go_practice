package dao

import (
	"dbcourse/model"
	"dbcourse/utils"
)

// 找出所有账号的数据
func Getalluser() []*model.User {
	sq := "select username,password,utype,conid from user4794"
	rows, _ := utils.Db.Query(sq)
	var res []*model.User
	for rows.Next() {
		// var temp *model.user
		temp := &model.User{}
		rows.Scan(&temp.Username, &temp.Password, &temp.Utype, &temp.Conid)
		res = append(res, temp)
	}
	return res
}

// 找出某一条账号的数据
func Getoneuser(x string) (*model.User, error) {
	sq := "select username,password,utype,conid from user4794 where username = ?"
	row := utils.Db.QueryRow(sq, x)
	res := &model.User{}
	err := row.Scan(&res.Username, &res.Password, &res.Utype, &res.Conid)
	return res, err
}
