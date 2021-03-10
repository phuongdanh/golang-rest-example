package model

type User struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string
	BirthDay string `json:"birth_day"`
	Password string `json:"-"`
}

func UserExamples() []*User {
	return []*User{
		{ID: 1, Name: "Nguyen Van A", Email: "nguyena@gmail.com", BirthDay: "2000-01-01", Password: "****************"},
		{ID: 2, Name: "Nguyen Van B", BirthDay: "2000-01-01", Password: "****************"},
		{ID: 3, Name: "Nguyen Van C", Email: "nguyenc@gmail.com", BirthDay: "2000-01-01", Password: "****************"},
		{ID: 4, Name: "Nguyen Van D", Email: "nguyend@gmail.com", BirthDay: "2000-01-01", Password: "****************"},
		{ID: 5, Name: "Nguyen Van E", Email: "nguyene@gmail.com", BirthDay: "2000-01-01", Password: "****************"},
	}
}
