package main

import (
	"dbcourse/dao"
	"dbcourse/model"
	"dbcourse/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static")

	//后台首页
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台首页",
		})
	})

	//将供应商表列出来并操作
	r.GET("/admin/supplierop", func(c *gin.Context) {
		Supplist := dao.Getallsupplier()
		c.HTML(http.StatusOK, "admin/supplierop.html", gin.H{
			"Supplist": Supplist,
		})
	})
	//删除供应商后跳转到的路由
	r.GET("/supplier/deletesupplier", func(c *gin.Context) {
		p := c.Query("Supplierid")
		x, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}
		dao.Deletesupplier(x)
		c.HTML(http.StatusOK, "supplier/deletesupplier.html", gin.H{})
	})
	//接收数据后添加供应商信息
	r.POST("/supplier/addsupplier", func(c *gin.Context) {
		id := c.PostForm("id")
		address := c.PostForm("address")
		name := c.PostForm("name")
		phone := c.PostForm("phone")
		head := c.PostForm("head")
		x := &model.Supplier{
			Address: address,
			Name:    name,
			Phone:   phone,
			Head:    head,
		}
		flag := 1
		if id != "" {
			x.Id, _ = strconv.Atoi(id)
		}
		fmt.Printf("x: %#v\n", x)
		err := dao.Insertsupplier(x)
		fmt.Printf("err: %v,%T\n", err, err)
		c.HTML(http.StatusOK, "supplier/addsupplier.html", gin.H{
			"flag": flag,
		})
	})
	//获取到图书的id后更改图书信息
	r.GET("/supplier/updatesupplier", func(c *gin.Context) {
		upidstr := c.Query("Supplierid")
		upid, _ := strconv.Atoi(upidstr)
		suff, _ := dao.Getonesupplier(upid)
		fmt.Printf("upid: %v\n", upid)
		c.HTML(http.StatusOK, "supplier/updatesupplier.html", gin.H{
			"suff": suff,
		})
	})
	//更改后最后的跳转
	r.POST("/supplier/updatesupplierfinal", func(c *gin.Context) {
		iid := c.PostForm("id")
		id, _ := strconv.Atoi(iid)
		address := c.PostForm("address")
		name := c.PostForm("name")
		phone := c.PostForm("phone")
		head := c.PostForm("head")

		suff := &model.Supplier{
			Id:      id,
			Address: address,
			Name:    name,
			Phone:   phone,
			Head:    head,
		}
		// fmt.Printf("suff: %v!!!!!!!!!!!!!!\n", suff)
		dao.Updatesupplier(suff)
		c.HTML(http.StatusOK, "supplier/updatesupplierfinal.html", gin.H{})
	})

	//将货物表列出来并进行操作
	r.GET("/admin/goodsop", func(c *gin.Context) {
		Goodslist := dao.Getallgoods()
		// for _, val := range Goodslist {
		// 	fmt.Printf("val: %v\n", val)
		// }
		Supplierlist := dao.Getallsupplier()
		var suppidname []string
		for _, val := range Supplierlist {
			res := strconv.Itoa(val.Id)
			res = res + "(" + val.Name + ")"
			suppidname = append(suppidname, res)
		}
		// fmt.Printf("suppidname: %v\n", suppidname)
		c.HTML(http.StatusOK, "admin/goodsop.html", gin.H{
			"Goodslist":  Goodslist,
			"Suppidname": suppidname,
		})
	})
	//删除货物信息后转跳到的路由
	r.GET("/goods/deletegoods", func(c *gin.Context) {
		p := c.Query("goodsid")
		x, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}
		dao.Deletegoods(x)
		c.HTML(http.StatusOK, "goods/deletegoods.html", gin.H{})
	})
	//接收到数据后添加货物信息
	r.POST("/goods/addgoods", func(c *gin.Context) {
		supplierid := c.PostForm("supplierid")
		price := c.PostForm("price")
		amount := c.PostForm("amount")
		name := c.PostForm("name")
		var pos int = 0
		for ; pos < len(supplierid); pos++ {
			if supplierid[pos] == '(' {
				break
			}
		}
		isupplierid, _ := strconv.Atoi(string([]byte(supplierid)[:pos]))
		iprice, _ := strconv.Atoi(price)
		iamount, _ := strconv.Atoi(amount)
		iname := name
		fmt.Println(supplierid, price, amount, name)
		goods := &model.Goods{
			Supplierid: isupplierid,
			Price:      iprice,
			Amount:     iamount,
			Name:       iname,
		}
		fmt.Printf("goods: %v\n", goods)
		dao.Insertgoods(goods)
		c.HTML(http.StatusOK, "goods/addgoods.html", gin.H{})
	})
	//获取到图书的id后更改图书信息
	r.GET("/goods/updategoods", func(c *gin.Context) {
		upidstr := c.Query("goodsid")
		upid, _ := strconv.Atoi(upidstr)
		goods, _ := dao.Getonegoods(upid)
		// fmt.Printf("goods: %v\n", goods)
		Supplierlist := dao.Getallsupplier()
		var suppidname []string
		for _, val := range Supplierlist {
			res := strconv.Itoa(val.Id)
			res = res + "(" + val.Name + ")"
			suppidname = append(suppidname, res)
		}
		c.HTML(http.StatusOK, "goods/updategoods.html", gin.H{
			"goods":      goods,
			"Suppidname": suppidname,
		})
	})
	//修改货物最后的跳转
	r.POST("/goods/updategoodsfinal", func(c *gin.Context) {
		id := c.PostForm("id")
		supplierid := c.PostForm("supplierid")
		price := c.PostForm("price")
		amount := c.PostForm("amount")
		name := c.PostForm("name")
		var pos int = 0
		for ; pos < len(supplierid); pos++ {
			if supplierid[pos] == '(' {
				break
			}
		}
		iid, _ := strconv.Atoi(id)
		isupplierid, _ := strconv.Atoi(string([]byte(supplierid)[:pos]))
		iprice, _ := strconv.Atoi(price)
		iamount, _ := strconv.Atoi(amount)
		iname := name
		fmt.Println(supplierid, price, amount, name)
		goods := &model.Goods{
			Id:         iid,
			Supplierid: isupplierid,
			Price:      iprice,
			Amount:     iamount,
			Name:       iname,
		}
		fmt.Printf("goods: %v\n", goods)
		dao.Updategoods(goods)
		c.HTML(http.StatusOK, "goods/updategoodsfinal.html", gin.H{})
	})

	//将员工表列出来并进行操作
	r.GET("/admin/staffop", func(c *gin.Context) {
		Stafflist := dao.Getallstaff()
		// for _, val := range Stafflist {
		// 	fmt.Printf("val: %v\n", val)
		// }
		c.HTML(http.StatusOK, "admin/staffop.html", gin.H{
			"Stafflist": Stafflist,
		})
	})
	//删除员工信息后跳转到的路由
	r.GET("/staff/deletestaff", func(c *gin.Context) {
		p := c.Query("staffid")
		x, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}
		dao.Deletestaff(x)
		c.HTML(http.StatusOK, "staff/deletestaff.html", gin.H{})
	})
	//接收到数据后添加货物信息
	r.POST("/staff/addstaff", func(c *gin.Context) {
		name := c.PostForm("name")
		age := c.PostForm("age")
		sex := c.PostForm("sex")
		phone := c.PostForm("phone")
		username := c.PostForm("username")
		password := c.PostForm("password")
		iage, _ := strconv.Atoi(age)
		x := &model.Staff{
			Name:     name,
			Age:      iage,
			Sex:      sex,
			Phone:    phone,
			Username: username,
			Password: password,
		}
		dao.Insertstaff(x)
		c.HTML(http.StatusOK, "goods/addstaff.html", gin.H{})
	})
	//收到员工的staffid并修改
	r.GET("/staff/updatestaff", func(c *gin.Context) {
		upidstr := c.Query("staffid")
		upid, _ := strconv.Atoi(upidstr)
		staff, _ := dao.Getonestaff(upid)
		// fmt.Printf("upidstr: %v\n", upidstr)
		// fmt.Printf("upid: %v\n", upid)
		// fmt.Printf("staff: %v\n", staff)
		c.HTML(http.StatusOK, "staff/updatestaff.html", gin.H{
			"staff": staff,
		})
	})
	//修改员工信息的最后跳转
	r.POST("/staff/updatestafffinal", func(c *gin.Context) {
		id := c.PostForm("id")
		name := c.PostForm("name")
		age := c.PostForm("age")
		sex := c.PostForm("sex")
		phone := c.PostForm("phone")
		username := c.PostForm("username")
		password := c.PostForm("password")
		iid, _ := strconv.Atoi(id)
		iage, _ := strconv.Atoi(age)
		staff := &model.Staff{
			Staffid:  iid,
			Name:     name,
			Sex:      sex,
			Phone:    phone,
			Username: username,
			Password: password,
			Age:      iage,
		}
		fmt.Printf("staff: %v!!!\n", staff)
		dao.Updatestaff(staff)
		c.HTML(http.StatusOK, "goods/updatestafffinal.html", gin.H{})
	})

	//将for列出来并进行操作
	r.GET("/admin/formop", func(c *gin.Context) {
		Formlist := dao.Getallform()
		Supplieridname := dao.Getallsupplieridname()
		Goodsidname := dao.Getallgoodsidname()
		Staffidname := dao.Getallstaffidname()
		c.HTML(http.StatusOK, "admin/formop.html", gin.H{
			"Formlist":       Formlist,
			"Supplieridname": Supplieridname,
			"Goodsidname":    Goodsidname,
			"Staffidname":    Staffidname,
		})
	})
	//删除员工信息后跳转到的路由
	r.GET("/form/deleteform", func(c *gin.Context) {
		p := c.Query("formnumber")
		dao.Deleteform(p)
		c.HTML(http.StatusOK, "form/deleteform.html", gin.H{})
	})
	//接收到数据后添加货物信息
	r.POST("form/addform", func(c *gin.Context) {
		opstaffid := dao.Removebucket(c.PostForm("opstaffid"))
		goodsid := dao.Removebucket(c.PostForm("goodsid"))
		amount := c.PostForm("amount")
		price := c.PostForm("price")
		inorout := c.PostForm("inorout")
		datee := c.PostForm("datee")
		iopstaffid, _ := strconv.Atoi(opstaffid)
		igoodsid, _ := strconv.Atoi(goodsid)
		iamount, _ := strconv.Atoi(amount)
		iprice, _ := strconv.Atoi(price)
		p := &model.Inoutform{
			Formnumber: dao.Proformnum(),
			Opstaffid:  iopstaffid,
			Goodsid:    igoodsid,
			Amount:     iamount,
			Price:      iprice,
			Inorout:    inorout,
			Datee:      datee,
		}
		dao.Insertinoutform(p)
		c.HTML(http.StatusOK, "form/addform.html", gin.H{})
	})
	//收到formnumber并进行修改
	r.GET("/form/updateform", func(c *gin.Context) {
		upnumber := c.Query("formnumber")
		form, _ := dao.Getoneform(upnumber)
		Staffidname := dao.Getallstaffidname()
		Goodsidname := dao.Getallgoodsidname()
		c.HTML(http.StatusOK, "form/updateform.html", gin.H{
			"form":        form,
			"Staffidname": Staffidname,
			"Goodsidname": Goodsidname,
		})
	})
	//修改form后最后的跳转
	r.POST("/form/updateformfinal", func(c *gin.Context) {
		Formnumber := c.PostForm("formnumber")
		opstaffid := dao.Removebucket(c.PostForm("opstaffid"))
		goodsid := dao.Removebucket(c.PostForm("goodsid"))
		amount := c.PostForm("amount")
		price := c.PostForm("price")
		inorout := c.PostForm("inorout")
		datee := c.PostForm("datee")
		iopstaffid, _ := strconv.Atoi(opstaffid)
		igoodsid, _ := strconv.Atoi(goodsid)
		iamount, _ := strconv.Atoi(amount)
		iprice, _ := strconv.Atoi(price)
		p := &model.Inoutform{
			Formnumber: Formnumber,
			Opstaffid:  iopstaffid,
			Goodsid:    igoodsid,
			Amount:     iamount,
			Price:      iprice,
			Inorout:    inorout,
			Datee:      datee,
		}
		dao.Updateform(p)
		c.HTML(http.StatusOK, "form/updateformfinal.html", gin.H{})
	})

	//用户登录界面
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"Nowloginuser": model.Nowloginuser,
		})
	})
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		// fmt.Println(username, password)
		p, _ := dao.Getoneuser(username)
		if p.Password == password {
			model.Nowloginuser.Username = username
			model.Nowloginuser.Status = 1
			model.Nowloginuser.Conid = p.Conid
			model.Nowloginuser.Utype = p.Utype
		}
		c.HTML(http.StatusOK, "login_judge.html", gin.H{
			"Nowloginuser": model.Nowloginuser,
		})
	})
	//员工进货出货
	r.GET("/work/gowork", func(c *gin.Context) {
		goods := dao.Getallgoodsidname()
		Suppidname := dao.Getallsupplieridname()
		c.HTML(http.StatusOK, "work/gowork.html", gin.H{
			"Nowloginuser": model.Nowloginuser,
			"goods":        goods,
			"Suppidname":   Suppidname,
		})
	})
	//处理出货进货
	r.POST("/work/addform", func(c *gin.Context) {
		supplierid := dao.Removebucket(c.PostForm("supplierid"))
		goodsname := c.PostForm("goodsname")
		amount := c.PostForm("amount")
		price := c.PostForm("price")
		inorout := c.PostForm("inorout")
		isupplierid, _ := strconv.Atoi(supplierid)
		iamount, _ := strconv.Atoi(amount)
		iprice, _ := strconv.Atoi(price)
		fmt.Printf("supplierid: %v\n", supplierid)
		fmt.Printf("goodsname: %v\n", goodsname)
		fmt.Printf("amount: %v\n", amount)
		fmt.Printf("price: %v\n", price)
		fmt.Printf("inorout: %v\n", inorout)
		sq := "select id,supplierid,price,amount,name from goods4794 where supplierid = ? and name = ? and price = ?"
		row := utils.Db.QueryRow(sq, supplierid, goodsname, price)
		res := &model.Goods{}
		err := row.Scan(&res.Id, &res.Supplierid, &res.Price, &res.Amount, &res.Name)
		fmt.Printf("err: %v\n", err)
		fmt.Printf("res: %v\n", res)
		if err != nil {
			newgoods := &model.Goods{
				Supplierid: isupplierid,
				Price:      iprice,
				Amount:     iamount,
				Name:       goodsname,
			}
			fmt.Printf("newgoods: %v!!!\n", newgoods)
			dao.Insertgoods(newgoods)
			fmt.Printf("isupplierid: %v!!!\n", isupplierid)
			fmt.Printf("goodsname: %v!!!\n", goodsname)
			fmt.Printf("iamount: %v!!!\n", iamount)
			row = utils.Db.QueryRow(sq, isupplierid, goodsname, iprice)
			err = row.Scan(&res.Id, &res.Supplierid, &res.Price, &res.Amount, &res.Name)
			fmt.Printf("err: %v\n", err)
			fmt.Printf("res: %v!!!\n", res)
		} else {
			if inorout == "入" {
				res.Amount += iamount
			} else {
				res.Amount -= iamount
			}
			dao.Updategoods(res)
		}
		fmt.Printf("res: %v\n", res)
		newform := &model.Inoutform{
			Formnumber: dao.Proformnum(),
			Opstaffid:  model.Nowloginuser.Conid,
			Goodsid:    res.Id,
			Amount:     iamount,
			Price:      iprice,
			Inorout:    inorout,
			Datee:      time.Now().Format("2006-01-02"),
		}
		dao.Insertinoutform(newform)
		c.HTML(http.StatusOK, "work/addform.html", gin.H{})
	})
	//注销
	r.GET("/logout", func(c *gin.Context) {
		model.Nowloginuser.Status = 0
		c.HTML(http.StatusOK, "logout.html", gin.H{})
	})
	//展示当前登录员工所参与的form
	r.GET("/work/myopform", func(c *gin.Context) {
		formslist := dao.Getallformbystaffid(model.Nowloginuser.Conid)
		fmt.Printf("formslist: %v\n", formslist)
		c.HTML(http.StatusOK, "work/myopform.html", gin.H{
			"formslist": formslist,
		})
	})
	//统计员工绩效
	r.GET("/admin/staffperformance", func(c *gin.Context) {
		staffperformancelist := dao.Getallstaffperformance()
		c.HTML(http.StatusOK, "/admin/staffperformance", gin.H{
			"staffperformancelist": staffperformancelist,
		})
	})
	// 统计供应商供货数
	r.GET("/admin/suppliergoods", func(c *gin.Context) {
		suppliergoods := dao.Getallsuppliergoods()
		c.HTML(http.StatusOK, "/admin/suppliergoods", gin.H{
			"suppliergoods": suppliergoods,
		})
	})
	r.Run(":8080")
}

//https://max.book118.com/html/2022/0519/6232140051004150.shtm
