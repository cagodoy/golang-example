package person

type Person struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name" validate:"required"`
	Age  int64  `json:"age" validate:"required"`
}

func New(name string, age int64) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

// composition over inheritance
//
// type Teacher struct {
// 	Person Person `json:"person"`

// 	Subject string  `json:"subject"`
// }
