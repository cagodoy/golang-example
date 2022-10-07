package person

type Person struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func New(name string, age int64) *Person {
	return &Person{
		Name: name,
		Age: age,
	}	
}

// composition over inheritance
//
// type Teacher struct {
// 	Person Person `json:"person"`

// 	Subject string  `json:"subject"`
// }
