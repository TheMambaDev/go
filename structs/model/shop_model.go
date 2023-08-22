package model

type Product struct {
	Name        string
	Description string
	Categories  []Category
	Price       int32
	Stock       int32
	ID          int8
}

type Category struct {
	Name string
	ID   int8
}
