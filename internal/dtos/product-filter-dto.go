package dtos

type ProductFilterDto struct {
	Name  *string `form:"title"`
	Stock *int    `form:"stock"`
}
