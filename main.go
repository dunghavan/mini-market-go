package main

import (
	"flag"
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

func setDBToUTF8() {
	o := orm.NewOrm()
	userFields := []string{"name", "first_name", "last_name"}
	for _, field := range userFields {
		queryStr := fmt.Sprintf("alter table user modify %s varchar(128) character SET utf8;", field)
		_, err := o.Raw(queryStr).Exec()
		if err != nil {
			glog.Errorf("Change column %s to UTF-8 err: %s", field, err.Error())
		}
	}
	itemFields := []string{"name", "desc", "delivery_way", "address", "note"}
	for _, field := range itemFields {
		queryStr := fmt.Sprintf("alter table item modify `%s` varchar(255) character SET utf8;", field)
		_, err := o.Raw(queryStr).Exec()
		if err != nil {
			glog.Errorf("Change column %s to UTF-8 err: %s", field, err.Error())
		}
	}
}

func init() {
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")

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
	setDBToUTF8()
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
	//au := models.Authority{Id: constant.RoleAdminId}
	//user := models.User{Login: "dunghavan", Email: "dunghavan@gmail.com", FacebookId: "12345", Authorities: []*models.Authority{&au}}
}
