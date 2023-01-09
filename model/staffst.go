package model

type Staff struct {
	Staffid  int    `db:"staffid"`
	Name     string `db:"name"`
	Age      int    `db:"age"`
	Sex      string `db:"sex"`
	Phone    string `db:"phone"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type Staffperformance struct {
	Staffid        int
	Staffname      string
	Opcount        int
	Allgoodsamount int
}
