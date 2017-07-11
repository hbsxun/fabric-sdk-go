package AssetApp

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestChaincode(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	modelName := "Asset" + strconv.Itoa(r.Int())
	invoke(modelName, t)
	time.Sleep(time.Second * 3)
	queryModelByXX(modelName, t)
}

func invoke(modelName string, t *testing.T) {
	//add asset
	if _, err := AddModel(&AddModelArgs{
		Name:  modelName,
		Desc:  "desc",
		Price: "13.14",
		Owner: "alice",
	}); err != nil {
		t.Fatal(err)
	}

	//move asset
	if _, err := TransferModel(&TransferModelArgs{
		Name:     modelName,
		NewOwner: "bob",
	}); err != nil {
		t.Fatal(err)
	}
}

func queryModelByXX(modelName string, t *testing.T) {
	//query model
	asset, err := QueryModel(modelName)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("Model: ", asset)
	}
	//query model history
	history, err := GetHistoryForModel(modelName)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("Model History: ", history)
	}

	//query owner's models
	models, err := QueryModelsByOwner("bob")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("Owner's models: ", models)
	}
}

/*
func TestDelModel(modelName string, t *testing.T) {
	//del asset
	if _, err := AddModel(modelName); err != nil {
		t.Fatal(err)
	}
	//query model
	asset, err := QueryModel(modelName)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("Model: ", asset)
	}

}
*/
