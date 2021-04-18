package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"shopaholic-service/types"
)

func GetShoppingLists() ([]types.ShoppingList, error) {
	jsonFile, err := os.Open("./data/sample.json")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data []types.ShoppingList
	json.Unmarshal([]byte(byteValue), &data)
	fmt.Println(data)
	return data, nil
}
