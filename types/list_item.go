package types

type ListItem struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Quantity       int    `json:"quantity"`
	ShoppingListID int    `json:"shopping_list_id"`
	AddedAt        string `json:"added_at"`
	IsChecked      bool   `json:"is_checked"`
}
