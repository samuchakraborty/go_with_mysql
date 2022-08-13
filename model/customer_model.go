package model

type Customers struct {
	Id       int    `json:"id" query:"id"`
	Fullname string `json:"name" query:"name"`
	Mobile   string `json:"mobile" query:"mobile"`
}
