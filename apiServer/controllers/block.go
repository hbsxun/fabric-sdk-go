package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/test/apiServer/models"
)

// Operations about Blocks
type BlockController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Blocks
// @Success 200 {object} models.Block
// @router / [get]
func (u *BlockController) GetAll() {
	blocks, err := models.GetBlocks()
	fmt.Println(blocks)
	if err != nil {
		u.Data["json"] = err
	} else {
		u.Data["json"] = blocks
	}
	fmt.Println("blocks: ", blocks)
	u.ServeJSON()
}

// @Title Get
// @Description get block by number
// @Param	number		path 	int	true		"The key for block"
// @Success 200 {object} models.Block
// @Failure 403 :number is out of bound
// @router /:number [get]
func (u *BlockController) Get() {
	uid := u.GetString(":number")
	fmt.Println("number: ", uid)
	if uid != "" {
		i, err := strconv.Atoi(uid)
		if err != nil {
			u.Data["json"] = err
		} else {
			block, err := models.GetBlockByNumber(i)
			if err != nil {
				u.Data["json"] = err
			} else {
				u.Data["json"] = block
			}
		}
	}
	u.ServeJSON()
}
