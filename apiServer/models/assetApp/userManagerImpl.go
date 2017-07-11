package assetApp

import "github.com/hyperledger/fabric-sdk-go/apiServer/models/user"

type UserManagerImpl struct{}

func (this *UserManagerImpl) Register(u *user.User) (int64, bool) {
	id, err := user.AddUser(u)
	if err != nil {
		appLogger.Debugf("User Register err [%v]\n", err)
		return -1, false
	}
	return id, true
}

func (this *UserManagerImpl) Login(name, passwd string) bool {
	ok, err := user.Login(name, passwd)
	if err != nil {
		appLogger.Debugf("User Login err [%v]\n", err)
	}
	if !ok {
		appLogger.Debugf("User Login failed, Username or Password is incorrect")
		return false
	}
	return true
}

func (this *UserManagerImpl) UpdateInfo(u *user.User) error {
	err := user.UpdateUser(u)
	if err != nil {
		appLogger.Debugf("User UpdateInfo err [%v]\n", err)
		return err
	}
	return nil
}

func (this *UserManagerImpl) UpdatePasswd(username, passwd string) bool {
	appLogger.Debug("Not implemented")
	return false
}
