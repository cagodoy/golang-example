package person

type Person struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
	// leveldb doesn't go here
}

// composition over inheritance
//
// type Teacher struct {
// 	Person Person `json:"person"`

// 	Subject string  `json:"subject"`
// }
