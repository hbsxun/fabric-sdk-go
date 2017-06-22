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
	"strconv"

	"github.com/widuu/gomysql"
)

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
	Mail   string `json:"mail"`
	Cert   string `json:"cert"`
}

var dbUtils *gomysql.Model

func init() {
	var err error
	dbUtils, err = gomysql.SetConfig("./config/db.ini")
	if err != nil {
		panic(err)
	}
}

func AddUser(u *User) (id int, err error) {
	var value = make(map[string]interface{})
	value["name"] = u.Name
	value["passwd"] = u.Passwd
	value["mail"] = u.Mail
	value["cert"] = u.Cert

	t := dbUtils.SetTable("ASSET_USER")
	t.SetPk("id")

	i, err := t.Insert(value)
	if err != nil {
		return -1, err
	}

	return i, nil
}

func GetUser(id int) (u *User, err error) {
	data := dbUtils.SetTable("ASSET_USER").Fileds("id", "name", "passwd", "mail", "cert").Where("id=" + strconv.Itoa(id)).FindOne()
	i, _ := strconv.Atoi(data[1]["id"])
	u = &User{
		Id:     i,
		Name:   data[1]["name"],
		Passwd: data[1]["passwd"],
		Mail:   data[1]["mail"],
		Cert:   data[1]["cert"],
	}
	return u, nil
}

func UpdateUser(id string, u *User) (bool, error) {
	var value = make(map[string]interface{})
	value["name"] = u.Name
	value["passwd"] = u.Passwd
	value["mail"] = u.Mail
	value["cert"] = u.Cert

	_, err := dbUtils.SetTable("ASSET_USER").Where("id=" + id).Update(value)
	if err != nil {
		return false, err
	}

	return true, nil
}

func Login(username, passwd string) bool {
	data := dbUtils.SetTable("ASSET_USER").Fileds("name", "passwd").Where("name=" + username).FindOne()
	if passwd == data[1]["passwd"] {
		return true
	}
	return false
}
