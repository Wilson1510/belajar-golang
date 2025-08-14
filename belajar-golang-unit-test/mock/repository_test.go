// main.go
package main

import (
	"errors"
	"fmt"
	"testing"
)

// --- Repositori dan Model Data ---
// User merepresentasikan model data pengguna
type User struct {
	ID   int
	Name string
}

// UserRepository adalah antarmuka untuk operasi database pengguna
type UserRepository interface {
	GetUserByID(id int) (*User, error)
}

// --- Implementasi Mock Database ---
// MockDB adalah implementasi mock dari UserRepository
type MockDB struct {
	Users map[int]*User
}

// NewMockDB membuat instance MockDB baru dengan data yang telah ditentukan
func NewMockDB() *MockDB {
	return &MockDB{
		Users: map[int]*User{
			1: {ID: 1, Name: "Alice"},
			2: {ID: 2, Name: "Bob"},
		},
	}
}

// GetUserByID mengambil data dari map Users kita, bukan dari database
func (m *MockDB) GetUserByID(id int) (*User, error) {
	if user, ok := m.Users[id]; ok {
		return user, nil
	}
	return nil, errors.New("user not found")
}

// --- Layanan Aplikasi ---
// UserService adalah layanan yang menggunakan UserRepository
type UserService struct {
	Repo UserRepository
}

// GetUserNameByID mengambil nama pengguna berdasarkan ID
func (s *UserService) GetUserNameByID(id int) (string, error) {
	user, err := s.Repo.GetUserByID(id)
	fmt.Println("user", user)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}

// --- Fungsi Pengujian ---
// TestGetUserNameByID adalah fungsi pengujian untuk UserService
func TestGetUserNameByID(t *testing.T) {
	// Buat instance mock database
	mockDB := NewMockDB()

	// Buat instance UserService dengan mock database
	service := &UserService{Repo: mockDB}
	fmt.Println(service)

	// Kasus uji: pengguna ditemukan
	t.Run("pengguna ditemukan", func(t *testing.T) {
		name, err := service.GetUserNameByID(1)
		if err != nil {
			t.Errorf("expected no error, but got: %v", err)
		}
		if name != "Alice" {
			t.Errorf("expected name 'Alice', but got: %s", name)
		}
	})

	// Kasus uji: pengguna tidak ditemukan
	t.Run("pengguna tidak ditemukan", func(t *testing.T) {
		_, err := service.GetUserNameByID(99)
		if err == nil {
			t.Errorf("expected an error, but got none")
		}
	})
}

// --- Fungsi utama (opsional) ---
// main adalah fungsi utama untuk menjalankan contoh ini
func main() {
	// Contoh penggunaan dengan mock database
	mockDB := NewMockDB()
	service := &UserService{Repo: mockDB}

	name, err := service.GetUserNameByID(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Nama pengguna dengan ID 1:", name) // Output: Nama pengguna dengan ID 1: Alice
	}

	name, err = service.GetUserNameByID(3)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: user not found
	} else {
		fmt.Println("Nama pengguna dengan ID 3:", name)
	}
}