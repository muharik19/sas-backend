package models

// ItemModel struct
type ItemModel struct {
	ID         int         `json:"id"`
	ItemName   string      `json:"item_name"`
	Price      float64     `json:"price"`
	WeightML   interface{} `json:"weight_ml"`
	WeightMG   interface{} `json:"weight_mg"`
	CreatedBy  interface{} `json:"created_by"`
	CreatedAt  interface{} `json:"created_at"`
	ModifiedBy interface{} `json:"modified_by"`
	ModifiedAt interface{} `json:"modified_at"`
	IsDeleted  interface{} `json:"is_deleted"`
}

// ItemsModel struct
type ItemsModel struct {
	Data  []ItemModel `json:"data"`
	Total int         `json:"total"`
}
