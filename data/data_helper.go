package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"shopaholic-service/types"
)

func writeListsToFile(lists []types.ShoppingList) {
	bytes, _ := json.Marshal(lists)
	filename := "./data/sample.json"
	ioutil.WriteFile(filename, bytes, 0600)
}

func GetShoppingLists() ([]types.ShoppingList, error) {
	filename := "./data/sample.json"
	jsonFile, err := os.Open(filename)
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

func CreateShoppingList(list types.ShoppingList) (types.ShoppingList, error) {
	lists, _ := GetShoppingLists()
	newId := lists[len(lists)-1].ID + 1
	list.ID = newId
	list.Items = []types.ListItem{}
	lists = append(lists, list)
	writeListsToFile(lists)
	return list, nil
}

func GetShoppingList(listId string) (types.ShoppingList, error) {
	lists, err := GetShoppingLists()
	if err != nil {
		fmt.Println(err)
		return types.ShoppingList{}, err
	}
	var found types.ShoppingList
	for _, list := range lists {
		listIdInt, _ := strconv.Atoi(listId)
		if list.ID == listIdInt {
			found = list
		}
	}
	return found, nil
}

func RemoveShoppingList(listId string) ([]types.ShoppingList, error) {
	lists, err := GetShoppingLists()
	listIdInt, _ := strconv.Atoi(listId)
	if err != nil {
		fmt.Println(err)
		return []types.ShoppingList{}, err
	}
	for i, list := range lists {
		if list.ID == listIdInt {
			lists = append(lists[:i], lists[i+1:]...)
			break
		}
	}
	writeListsToFile(lists)
	return lists, nil
}

func AddItemToShoppingList(listId string, item types.ListItem) (types.ListItem, error) {
	lists, err := GetShoppingLists()
	listIdInt, _ := strconv.Atoi(listId)
	item.ShoppingListID = listIdInt
	if err != nil {
		fmt.Println(err)
		return types.ListItem{}, err
	}
	for i, list := range lists {
		if list.ID == listIdInt {
			items := list.Items
			var lastID int
			if len(items) > 0 {
				lastID = items[len(items)-1].ID
			} else {
				lastID = 0
			}
			newId := lastID + 1
			item.ID = newId
			list.Items = append(list.Items, item)
			lists[i] = list
			break
		}
	}
	writeListsToFile(lists)
	return item, nil
}

func RemoveItemFromShoppingList(listId string, itemId string) (types.ShoppingList, error) {
	lists, err := GetShoppingLists()
	listIdInt, _ := strconv.Atoi(listId)
	itemIdInt, _ := strconv.Atoi(itemId)
	if err != nil {
		fmt.Println(err)
		return types.ShoppingList{}, err
	}
	var modifiedList types.ShoppingList
	for i, list := range lists {
		if list.ID == listIdInt {
			items := list.Items
			for j, item := range items {
				if item.ID == itemIdInt {
					items = append(items[:j], items[j+1:]...)
					break
				}
			}
			list.Items = items
			lists[i] = list
			modifiedList = list
			break
		}
	}
	writeListsToFile(lists)
	return modifiedList, nil
}

func UpdateShoppingListItem(listId string, itemId string, updateFields types.ListItem) (types.ListItem, error) {
	lists, err := GetShoppingLists()
	listIdInt, _ := strconv.Atoi(listId)
	itemIdInt, _ := strconv.Atoi(itemId)
	updateFields.ShoppingListID = listIdInt
	if err != nil {
		fmt.Println(err)
		return types.ListItem{}, err
	}
	var modifiedItem types.ListItem
	for i, list := range lists {
		if list.ID == listIdInt {
			items := list.Items
			for j, item := range items {
				if item.ID == itemIdInt {
					item.IsChecked = updateFields.IsChecked
					item.Quantity = updateFields.Quantity
					item.Name = updateFields.Name
					modifiedItem = item
					items[j] = modifiedItem
					break
				}
			}
			list.Items = items
			lists[i] = list
			break
		}
	}
	writeListsToFile(lists)
	return modifiedItem, nil
}
