package users

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type UserService struct{}

func (us UserService) GetUsers(ctx context.Context, conn *pgx.Conn) ([]User, bool) {
	fmt.Println("Getting all users...")
	rows, _ := conn.Query(ctx, "select * from users.user")
	listofusers := []User{}
	for rows.Next() {
		// from models.User{}
		var (
			id   int
			name string
			age  int
		)
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			fmt.Printf("error happened when getting all users %v\n", err)
			fmt.Printf("error happened")
			return []User{}, false
		}
		usertoadd := User{
			UserId: id,
			Name:   name,
			Age:    age,
		}
		listofusers = append(listofusers, usertoadd)
	}
	return listofusers, true
}

func (us UserService) FindUser(ctx context.Context, conn *pgx.Conn, id int) (User, bool) {
	fmt.Printf("Finding user with id of %s\n", id)
	var userId int
	var username string
	var age int
	err := conn.QueryRow(ctx, "select * from users.user where user_id = $1", id).Scan(&userId, &username, &age)
	if err == pgx.ErrNoRows {
		fmt.Println("No rows were found - skipping")
		return User{}, false
	}
	if err != nil {
		fmt.Printf("Error happened when trying to find specific user %v\n", err)
		return User{}, false
	}

	return User{
		UserId: userId,
		Name:   username,
		Age:    age,
	}, true
}

func (us UserService) CreateUser(ctx context.Context, conn *pgx.Conn, user User) bool {
	fmt.Printf("Creating user with details of %v\n", user)
	query := `INSERT INTO users.user (name, age) VALUES (@userName, @userAge)`
	args := pgx.NamedArgs{
		"userName": user.Name,
		"userAge":  user.Age,
	}
	result, err := conn.Exec(ctx, query, args)
	if err != nil {
		fmt.Printf("Error happened when trying to insert user %v\n", user)
		return false
	}

	fmt.Printf("result of insertion is %v\n", result)
	return true
}

func (us UserService) UpdateUser(ctx context.Context, conn *pgx.Conn, userId int, user User) bool {
	fmt.Printf("Updating user with id of %v\n", userId)
	query := `UPDATE users.user set name = @userName, age = @userAge where user_id = @userId`
	args := pgx.NamedArgs{
		"userName": user.Name,
		"userAge":  user.Age,
		"userId":   userId,
	}
	_, err := conn.Exec(ctx, query, args)
	if err != nil {
		fmt.Printf("Error happened when updating user %v\n", err)
	}
	fmt.Println("Successfully updated user")
	return true
}

func (us UserService) DeleteUser(ctx context.Context, conn *pgx.Conn, userId int) bool {
	fmt.Printf("Deleting user with user_id of %v\n", userId)
	query := "DELETE FROM users.user where user_id = @userId"
	args := pgx.NamedArgs{
		"userId": userId,
	}
	_, err := conn.Exec(ctx, query, args)
	if err != nil {
		fmt.Printf("Error happened when deleting user %\v\n", err)
		return false
	}
	return true
}
