package types

type ShoppingList struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	CreatedAt string     `json:"created_at"`
	Items     []ListItem `json:"items"`
}
