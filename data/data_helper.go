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

// GetShoppingLists returns a list of shoppingList
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

// CreateShoppingList returns a ShoppingList
func CreateShoppingList(list types.ShoppingList) (types.ShoppingList, error) {
	lists, _ := GetShoppingLists()
	newID := lists[len(lists)-1].ID + 1
	list.ID = newID
	list.Items = []types.ListItem{}
	lists = append(lists, list)
	writeListsToFile(lists)
	return list, nil
}

// GetShoppingList returns a ShoppingList
func GetShoppingList(listID string) (types.ShoppingList, error) {
	lists, err := GetShoppingLists()
	if err != nil {
		fmt.Println(err)
		return types.ShoppingList{}, err
	}
	var found types.ShoppingList
	for _, list := range lists {
		listIDInt, _ := strconv.Atoi(listID)
		if list.ID == listIDInt {
			found = list
		}
	}
	return found, nil
}

// RemoveShoppingList returns a ShoppingList
func RemoveShoppingList(listID string) ([]types.ShoppingList, error) {
	lists, err := GetShoppingLists()
	listIDInt, _ := strconv.Atoi(listID)
	if err != nil {
		fmt.Println(err)
		return []types.ShoppingList{}, err
	}
	for i, list := range lists {
		if list.ID == listIDInt {
			lists = append(lists[:i], lists[i+1:]...)
			break
		}
	}
	writeListsToFile(lists)
	return lists, nil
}

// AddItemToShoppingList returns a ListItem
func AddItemToShoppingList(listID string, item types.ListItem) (types.ListItem, error) {
	lists, err := GetShoppingLists()
	listIDInt, _ := strconv.Atoi(listID)
	item.ShoppingListID = listIDInt
	if err != nil {
		fmt.Println(err)
		return types.ListItem{}, err
	}
	for i, list := range lists {
		if list.ID == listIDInt {
			items := list.Items
			var lastID int
			if len(items) > 0 {
				lastID = items[len(items)-1].ID
			} else {
				lastID = 0
			}
			newID := lastID + 1
			item.ID = newID
			list.Items = append(list.Items, item)
			lists[i] = list
			break
		}
	}
	writeListsToFile(lists)
	return item, nil
}

// RemoveItemFromShoppingList returns a ShoppingList
func RemoveItemFromShoppingList(listID string, itemID string) (types.ShoppingList, error) {
	lists, err := GetShoppingLists()
	listIDInt, _ := strconv.Atoi(listID)
	itemIDInt, _ := strconv.Atoi(itemID)
	if err != nil {
		fmt.Println(err)
		return types.ShoppingList{}, err
	}
	var modifiedList types.ShoppingList
	for i, list := range lists {
		if list.ID == listIDInt {
			items := list.Items
			for j, item := range items {
				if item.ID == itemIDInt {
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

// UpdateShoppingListItem returns a ListItem
func UpdateShoppingListItem(listID string, itemID string, updateFields types.ListItem) (types.ListItem, error) {
	lists, err := GetShoppingLists()
	listIDInt, _ := strconv.Atoi(listID)
	itemIDInt, _ := strconv.Atoi(itemID)
	updateFields.ShoppingListID = listIDInt
	if err != nil {
		fmt.Println(err)
		return types.ListItem{}, err
	}
	var modifiedItem types.ListItem
	for i, list := range lists {
		if list.ID == listIDInt {
			items := list.Items
			for j, item := range items {
				if item.ID == itemIDInt {
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
