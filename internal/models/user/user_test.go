package user

import "testing"

func TestCreateUser(t *testing.T) {
	t.Run("Create user", func(t *testing.T) {
		user := CreateUser("test", "test@test.com", "test", map[string]string{
			"name": "Test",
		})
		if user.Username != "test" {
			t.Error("Username is not test")
		}
		if user.Email != "test@test.com" {
			t.Error("Email is not test@test.com")
		}
		if user.Password != "test" {
			t.Error("Password is not test")
		}
		if user.Name != "Test" {
			t.Error("Name is not Test")
		}
		if user.CreatedAt.IsZero() {
			t.Error("CreatedAt is zero")
		}
		if user.UpdatedAt.IsZero() {
			t.Error("UpdatedAt is zero")
		}
	})
}

func TestInsertUser(t *testing.T) {
	t.Run("Insert user", func(t *testing.T) {
		user := CreateUser("test", "test@test.com", "test", map[string]string{
			"name": "Test",
		})
		err := InsertUser(user)
		if err != nil {
			t.Error("Failed to insert user")
		}
	})
}
