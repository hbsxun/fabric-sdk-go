package assetApp

import (
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/hjwt"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/user"
)

type UserManagerImpl struct{}

func (this *UserManagerImpl) Register(u *user.User) (int64, bool) {
	id, err := user.AddUser(u)
	if err != nil {
		appLogger.Debugf("User Register err [%v]\n", err)
		return -1, false
	}
	return id, true
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
	signedToken = hjwt.CreateToken(user.Id, user.Name, isAdmin)
	return signedToken, nil
}

func (this *UserManagerImpl) UpdateInfo(u *user.User) error {
	err := user.UpdateUser(u)
	if err != nil {
		appLogger.Debugf("User UpdateInfo err [%v]\n", err)
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
func (this *UserManagerImpl) UpdatePasswd(username, passwd string) bool {
	appLogger.Debug("Not implemented")
	return false
}
