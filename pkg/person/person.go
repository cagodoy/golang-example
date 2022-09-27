package person

// import (
// 	"github.com/syndtr/goleveldb/leveldb"
// )

type Person struct {
	Name string `json:"name"`
	Age int64 `json:"age"`
	
	// DB *leveldb.DB
}

func NewPerson(name string, age int64) *Person {
	// db, err := leveldb.OpenFile("./storage.db", nil)
	// if err != nil {
	// 	return nil, err
	// }

	return &Person{
		Name: name,
		Age: age,
		// DB: db,
	}
}

func (p *Person) GetName() string {
	return p.Name
}

// composition over inheritance
//
// type Teacher struct {
// 	Person Person `json:"person"`

// 	Subject string  `json:"subject"`
// }