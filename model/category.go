package model

type Category struct {
	CategoryId string `json:"category" gorm:"column:category_id"`
	Name string `json:"name" gorm:"column:name"`
	Desc string `json:"desc" gorm:"column:desc"`
	Order int `json:"order" gorm:"column:order"`
	ParentId string `json:"parentId" gorm:"column:parent_id"`
	IsDeleted bool `json:"isDeleted" gorm:"column:is_deleted"`
}

type CategoryResult struct {
	C1CategoryID string `gorm:"column:c1_category_id"`
	C1Name string `gorm:"column:c1_name"`
	C1Desc string `gorm:"column:c1_desc"`
	C1Order int `gorm:"column:c1_order"`
	C1ParentId string `gorm:"column:c1_parent_id"`

	C2CategoryID string `gorm:"column:c2_category_id"`
	C2Name string `gorm:"column:c2_name"`
	C2Order int `gorm:"column:c2_order"`
	C2ParentId string `gorm:"column:c2_parent_id"`

	Key string `json:"key"`
	Id string `json:"id"`
	C3CategoryID string `gorm:"column:c3_category_id"`
	C3Name string `gorm:"column:c3_name"`
	C3Order int `gorm:"column:c3_order"`
	C3ParentId string `gorm:"column:c3_parent_id"`

	IsDeleted bool `json:"isDeleted" gorm:"column:c3_is_deleted"`
}






