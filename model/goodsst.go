package model

type Goods struct {
	Id         int    `db:"id"`
	Supplierid int    `db:"supplierid"`
	Price      int    `db:"price"`
	Amount     int    `db:"amount"`
	Name       string `db:"name"`
}

// p := &model.Goods{
// 	Id:         1,
// 	Supplierid: 1,
// 	Price:      20,
// 	Amount:     100,
// 	Name:       "恐龙玩具",
// }
