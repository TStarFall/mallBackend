package model

type ShoppingCartItem struct {
	CartItemId string `json:"cartItemId"`
	UserId string `json:"userId"`
	ProductId string `json:"productId"`
	ProductCount int `json:"productCount"`
	IsDeleted bool `json:"isDeleted"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}
