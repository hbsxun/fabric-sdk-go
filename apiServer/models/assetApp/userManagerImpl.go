package assetApp

import (
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/assetApp/auth"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/user"
)

type UserManagerImpl struct{}

func (this *UserManagerImpl) Register(u *user.User) (int64, error) {
	id, err := user.AddUser(u)
	if err != nil {
		appLogger.Debugf("User Register err [%v]\n", err)
		return -1, err
	}
	return id, nil
}

//Login signedToken is for SSO and authorization
func (this *UserManagerImpl) Login(ss *user.Secret) (signedToken string, err error) {
	user, err := user.Login(ss)
	if err != nil {
		appLogger.Debugf("User Login failed [%v]\n", err)
		return "", err
	}

	var isAdmin bool = false
	if user.Type == 0 {
		isAdmin = true
	}
	signedToken = auth.CreateToken(user.Id, user.Name, isAdmin)
	return signedToken, nil
}

func (this *UserManagerImpl) VerifyUser(ss *user.Secret) error {
	_, err := user.Login(ss)
	if err != nil {
		appLogger.Debugf("Verify user failed [%v]\n", err)
		return err
	}
	return nil
}

func (this *UserManagerImpl) UpdateInfo(u *user.UpdateUserArgs) error {
	err := user.UpdateUser(u)
	if err != nil {
		appLogger.Debugf("User UpdateInfo err [%v]\n", err)
		return err
	}
	return nil
}

func (this *UserManagerImpl) UpdatePwd(name string, oldPwd string, newPwd string) error {
	err := user.UpdatePasswd(name, oldPwd, newPwd)
	if err != nil {
		appLogger.Debugf("User UpdatePwd err [%v]\n", err)
		return err
	}
	return nil
}

func (this *UserManagerImpl) GetUserInfoByName(username string) (*user.User, error) {
	userInfo, err := user.GetUserByName(username)
	if err != nil {
		appLogger.Debugf("User GetUserByName err [%v]\n", err)
		return nil, err
	}
	return userInfo, nil
}

func (this *UserManagerImpl) GetUserInfoById(userid int) (*user.User, error) {
	userInfo, err := user.GetUserById(userid)
	if err != nil {
		appLogger.Debugf("User GetUserById err [%v]\n", err)
		return nil, err
	}
	return userInfo, nil
}

/*func (this *UserManagerImpl) Logout() error {

}*/
