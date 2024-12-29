package types

type Student struct {
	Id    string
	Name  string `validate:"required,min=5,max=20"`
	Email string `validate:"required"`
	Age   int    `validate:"required"`
}
