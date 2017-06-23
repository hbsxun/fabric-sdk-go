/*
Copyright Beijing Sansec Technology Development Co., Ltd. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at


      http://www.apache.org/licenses/LICENSE-2.0


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package user

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
	Mail   string `json:"mail"`
	Cert   string `json:"cert"`
}

func AddUser(user *User) (id int, err error) {
	o := orm.NewOrm()

	id64, err := o.Insert(user)
	fmt.Println(id64, err)
	if err != nil {
		return -1, err
	}
	return int(id64), nil
}

func GetUser(idStr string) (u *User, err error) {
	o := orm.NewOrm()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	user := User{
		Id: id,
	}
	err = o.Read(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(idStr string, u *User) (*User, error) {
	o := orm.NewOrm()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}

	user := User{
		Id:     id,
		Name:   u.Name,
		Mail:   u.Mail,
		Passwd: u.Passwd,
		Cert:   u.Cert,
	}
	_, err = o.Update(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func Login(username, passwd string) (bool, error) {
	o := orm.NewOrm()
	u := User{}

	err := o.Raw("SELECT name, passwd FROM user WHERE name = ?", username).QueryRow(&u)
	if err != nil {
		return false, err
	}
	if passwd == u.Passwd {
		return true, nil
	}
	return false, nil
}

func init() {
	//register model
	orm.RegisterModel(new(User))
	//register driver
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//set default database
	orm.RegisterDataBase("default", "mysql", "hxy:hxy@tcp(localhost:3306)/hxydb?charset=utf8", 30)
	/*
		//max idle connections
		orm.SetMaxIdleConns("default", 30)
		//max opened connections
		orm.SetMaxOpenConns("default", 30)
	*/
}
