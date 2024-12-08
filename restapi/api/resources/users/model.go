package users

type User struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
}
