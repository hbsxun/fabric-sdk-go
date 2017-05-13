package handler

import (
	"encoding/json"

	"github.com/gislu/goSocket/server/utils"
	"github.com/hyperledger/fabric-ca/api"
	fabricCAClient "github.com/hyperledger/fabric-sdk-go/fabric-ca-client"
	sdkIgn "github.com/hyperledger/fabric-sdk-go/test/integration"
)

var admin *sdkIgn.Member

const (
	REGISTER = "register"
	ENROLL   = "enroll"
)

type RegisterController struct {
}

func (this *RegisterController) Excute(msg utils.Msg) []byte {
	utils.Log("*********************************************")
	utils.Log(msg.Content)
	/*TODO
	//Type: peer, app, user
	if args["type"] != "user" {
		utils.LogErr("type must be \"user\"")
		return nil
	}
	_, ok = args["maxEnrollments"].(int32)
	if !ok {
		utils.LogErr("maxEnrollments must be integer")
		return nil
	}
	*/
	contentJson, err := json.Marshal(msg.Content)
	if err != nil {
		utils.LogErr("msg.Content marshal err %v", err)
		return nil
	}
	var req fabricCAClient.RegistrationRequest
	err = json.Unmarshal(contentJson, &req)
	if err != nil {
		utils.LogErr("msg.Content Unmarshal err %v", err)
		return nil
	}
	utils.Log("RegistrationRequest struct\n", req)

	name, secret, err := admin.RegisterUser(&req)
	if err != nil {
		utils.LogErr("registerUser err ", err)
		return nil
	}

	utils.Logf("name [%s] secret [%s]", name, secret)
	var retMap = make(map[string]interface{})
	retMap["name"] = name
	retMap["secret"] = secret
	retJson, err := json.Marshal(retMap)
	if err != nil {
		utils.LogErr("retMap marshal to retJson err [%v]", err)
	}
	return retJson
}

//EnrollController
type EnrollController struct {
}

func (this *EnrollController) Excute(msg utils.Msg) []byte {
	utils.Log("*********************************************")
	utils.Log(msg.Content)

	contentJson, err := json.Marshal(msg.Content)
	if err != nil {
		utils.LogErr("msg.Content marshal err %v", err)
		return nil
	}
	var req api.EnrollmentRequest
	err = json.Unmarshal(contentJson, &req)
	if err != nil {
		utils.LogErr("msg.Content Unmarshal err %v", err)
		return nil
	}
	utils.Log("EnrollmentRequest struct\n", req)

	key, cert, err := admin.UserEnrollWithCSR(&req)
	if err != nil {
		utils.LogErr("UserEnrollWithCSR err ", err)
		return nil
	}

	utils.Logf("key\n[%s] cert\n[%s]\n", string(key), string(cert))
	var retMap = make(map[string]interface{})
	//retMap["key"] = base64.StdEncoding.EncodeToString(key)
	//retMap["cert"] = base64.StdEncoding.EncodeToString(cert)
	retMap["key"] = string(key)
	retMap["cert"] = string(cert)
	retJson, err := json.Marshal(retMap)
	if err != nil {
		utils.LogErr("retMap marshal to retJson err [%v]", err)
	}
	return retJson
}

func init() {
	//add RegisterController
	var register RegisterController
	utils.Route(func(entry utils.Msg) bool {
		if entry.Meta["meta"] == REGISTER {
			return true
		}
		return false
	}, &register)

	//add EnrollController
	var enroll EnrollController
	utils.Route(func(entry utils.Msg) bool {
		if entry.Meta["meta"] == ENROLL {
			return true
		}
		return false
	}, &enroll)

	admin = sdkIgn.NewMember()
}
