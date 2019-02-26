package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"mini-market-go/constant"
	"mini-market-go/models"
	_ "mini-market-go/models"
	_ "mini-market-go/routers"
)

var (
	aliasDBName = "default"
)

func initAuthorities() {
	aa := models.Authority{Id: constant.RoleAdminId, Name: "ROLE_ADMIN"}
	models.AddAuthority(&aa)
	au := models.Authority{Id: constant.RoleUserId, Name: "ROLE_USER"}
	models.AddAuthority(&au)
}

func init() {
	driverName := "mysql"
	orm.RegisterDriver(driverName, orm.DRMySQL)
	conf := beego.AppConfig.String("MySQLConnConfig")
	orm.RegisterDataBase(aliasDBName, driverName, conf)
	orm.Debug = true
	err := orm.RunSyncdb(aliasDBName, false, true)

	if err != nil {
		glog.Errorf("Run sync database error: %s", err.Error())
	}
	initAuthorities()
}

func main() {
	o := orm.NewOrm()
	o.Using(aliasDBName)

	beego.Run()
}

func testQuery() {
	//user, err := models.GetUserById(5)
	//if err == nil {
	//	fmt.Println(user.Login)
	//}

	au, err := models.GetAuthorityById(1)
	if err == nil {
		fmt.Println(au.Name)
	}
}

func testAddAuthority() {
	user := models.User{Id: 0, Login: "nghianv"}
	user2 := models.User{Id: 0, Login: "user2"}
	au := models.Authority{Id: 0, Name: "Test", Users: []*models.User{&user, &user2}}
	id, err := models.AddAuthority(&au)
	if err == nil {
		fmt.Printf("%v", id)
	}
}

func testAddUser() {
	au := models.Authority{Id: constant.RoleAdminId}
	user := models.User{Id: 0, Login: "dunghv", Authorities: []*models.Authority{&au}}
	id, err := models.AddUser(&user)
	if err != nil {
		glog.Errorf("Add user error: %s", err.Error())
	}

	o := orm.NewOrm()
	user.Id = id
	m2m := o.QueryM2M(&user, "Authorities")
	_, err = m2m.Add(user.Authorities[0])
	if err != nil {
		glog.Errorf("Add m2m error: %s", err.Error())
	}
}
