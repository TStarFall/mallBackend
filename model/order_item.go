package model

type OrderItem struct {
	OrderItemId		string `json:"OrderItemId"`
	OrderId			string `json:"orderId"`
	ProductId		string `json:"productId"`
	ProductName		string `json:"productName"`
	ProductCoverImg string `json:"productCoverImg"`
	SellingPrice	int	   `json:"sellingPrice"`
	ProductCount	int    `json:"productCount"`
	CreateAt		string `json:"createAt"`
	UpdateAt		string `json:"updateAt"`
}
