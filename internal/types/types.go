package types

type Student struct {
	Id    int64
	Name  string `validate:"required,min=5,max=20"`
	Email string `validate:"required"`
	Age   int    `validate:"required"`
}
