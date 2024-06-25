package main

type Employee struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Subscription struct {
	UserID     int `json:"user_id"`
	EmployeeID int `json:"employee_id"`
}
