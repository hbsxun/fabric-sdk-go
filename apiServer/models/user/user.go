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
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id     int    `json:"id"`   //primary key
	Type   int    `json:"type"` //0: admin, 1: user
	Name   string `json:"name"` //unique
	Passwd string `json:"passwd"`
	Email  string `json:"mail"`
	Phone  string `json:"phone"`
}
type Secret struct {
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
}

/*
type Identity struct {
	Id          int    `json:"id"`
	Key         string `json:"key"`
	Certificate string `json:"certificate"`
	UserId      int    `json:"userId"`
}
*/

func AddUser(user *User) (id int64, err error) {
	o := orm.NewOrm()

	id64, err := o.Insert(user)
	fmt.Println(id64, err)
	if err != nil {
		return -1, err
	}
	return id64, nil
}

func GetUser(username string) (*User, error) {
	o := orm.NewOrm()
	u := User{}

	err := o.Raw("SELECT type, name, passwd, phone, email FROM user WHERE name = ?", username).QueryRow(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func UpdateUser(newU *User) error {
	o := orm.NewOrm()
	oldU := User{}

	err := o.Raw("SELECT id from user WHERE name = ?", newU.Name).QueryRow(&oldU)
	if err != nil {
		return err
	}
	newU.Id = oldU.Id

	_, err = o.Update(newU)
	if err != nil {
		return err
	}

	return nil
}

func Login(ss *Secret) (*User, error) {
	o := orm.NewOrm()
	u := User{}

	err := o.Raw("SELECT * FROM user WHERE name = ?", ss.Name).QueryRow(&u)
	if err != nil {
		return nil, err
	}

	//	fmt.Println(reflect.TypeOf(passwd), reflect.TypeOf(u.Passwd))

	if u.Passwd == ss.Passwd {
		return &u, nil
	}
	return nil, errors.New("Invalid password")
}

func init() {
	//register model
	orm.RegisterModel(new(User))
	//register driver
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//set default database
	orm.RegisterDataBase("default", "mysql", "hxy:hxy@tcp(10.0.48.50:3306)/hxydb?charset=utf8", 30)

	/*
		//max idle connections
		orm.SetMaxIdleConns("default", 30)
		//max opened connections
		orm.SetMaxOpenConns("default", 30)
	*/
}
