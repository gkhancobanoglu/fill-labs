package user

import (
	"database/sql"
	"errors"
	"filllabs-task/user-management-system/backend/db"
)

// Fetch all users from the database
func GetUsersFromDB() ([]User, error) {
	database, err := db.InitDatabase()
	if err != nil {
		return nil, err
	}
	defer database.Close()

	rows, err := database.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Fetch a specific user by ID
func GetUserFromDB(id int) (*User, error) {
	database, err := db.InitDatabase()
	if err != nil {
		return nil, err
	}
	defer database.Close()

	var user User
	// Query the user by ID
	err = database.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

// Save a new user to the database
func SaveUserToDB(user *User) error {
	database, err := db.InitDatabase()
	if err != nil {
		return err
	}
	defer database.Close()

	// Insert user data into the users table
	_, err = database.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	return err
}

// Update an existing user in the database
func UpdateUserInDB(id int, user *User) error {
	database, err := db.InitDatabase()
	if err != nil {
		return err
	}
	defer database.Close()

	// Update user data by ID
	_, err = database.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, id)
	return err
}

// Delete a user by ID
func DeleteUserFromDB(id int) error {
	database, err := db.InitDatabase()
	if err != nil {
		return err
	}
	defer database.Close()

	// Delete user data from the users table
	_, err = database.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
