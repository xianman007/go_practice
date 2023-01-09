package model

type Supplier struct {
	Id      int    `db:"id"`
	Address string `db:"address"`
	Name    string `db:"name"`
	Phone   string `db:"phone"`
	Head    string `db:"head"`
}

type Suppliergoods struct {
	Supplierid         int
	Suppliername       string
	Goodsspeciesamount int
	Allgoodsamount     int
}

// p := &model.Supplier{
// 	Id:      ,
// 	Address: "",
// 	Name:    "",
// 	Phone:   "",
// 	Head:    "",
// }

// func (Supplier) Getsupplier() []*Supplier {

// }
