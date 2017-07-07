package AssetApp

import (
	"github.com/astaxie/beego/orm"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/cert"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/user"
)

type CertificateImpl struct{}
type Identity struct {
	Id          int    `json:"id"`
	Key         string `json:"key"`
	Certificate string `json:"cert"`
	User_id     int    `json:"user_id"`
}

func (this *CertificateImpl) GetIdentity(name string) (string, string) {
	var attribute []cert.Attribute
	o := orm.NewOrm()
	u := user.User{}
	err := o.Raw("SELECT id from user WHERE name = ?", name).QueryRow(&u)
	if err != nil {
		appLogger.Debugf("name does not exist  err [%v]\n", err)
		return "", ""
	}
	registerAction, err := cert.NewRegisterAction(&cert.RegisterArgs{name, "", "", "", attribute})
	if err != nil {
		appLogger.Debugf("create registeraction err [%v]\n", err)
		return "", ""
	}
	secret, err := registerAction.Execute()
	if err != nil {
		appLogger.Debugf("registeraction execute err [%v]\n", err)
		return "", ""
	}
	enrollAction, err := cert.NewEnrollAction(&cert.EnrollArgs{name, secret})
	if err != nil {
		appLogger.Debugf("create enrollaction err [%v]\n", err)
		return "", ""
	}
	key, cert, err := enrollAction.Execute()
	if err != nil {
		appLogger.Debugf("enrollaction execute err [%v]\n", err)
		return "", ""
	}
//	this.SaveToDB(key, cert, u.Id)
	return key, cert
}

//Deprecated, Not used now
func (this *CertificateImpl) SaveToDB(key string, cert string, userId int) error {
	o := orm.NewOrm()
	i := &Identity{
		Id:          1,
		Key:         key,
		Certificate: cert,
		User_id:     userId,
	}
	_, err := o.Insert(i)
	if err != nil {
		appLogger.Debugf("insert failed err [%v]\n", err)
		return err
	}
	return nil
}
func init() {
	orm.RegisterModel(new(Identity))
}
