package model

type User struct {
	Username string
	Password string
	Utype    int
	Conid    int
}

type Nowuser struct {
	Username string
	Status   int
	Utype    int
	Conid    int
}

var Nowloginuser Nowuser
