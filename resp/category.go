package resp

type Category struct {
	CategoryId	string 				  `json:"categoryId"`
	Name		string 			      `json:"name"`
	Desc		string 				  `json:"desc"`
	Order		int	   				  `json:"order"`
	ParentId	string 				  `json:"parentId"`
	Children	map[string]*Category2 `json:"children"`
}

type Category2 struct {
	CategoryId	string 				  `json:"categoryId"`
	Name		string 			      `json:"name"`
	Desc		string 				  `json:"desc"`
	Order		int	   				  `json:"order"`
	ParentId	string 				  `json:"parentId"`
	Children	map[string]*Category3 `json:"children"`
}

type Category3 struct {
	Key			string `json:"key"`
	Id			string `json:"id"`
	CategoryId	string `json:"categoryId"`
	Name		string `json:"name"`
	Desc		string `json:"desc"`
	Order		int	   `json:"order"`
	ParentId	string `json:"parentId"`
	IsDeleted	bool   `json:"isDeleted"`
}

