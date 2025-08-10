package database

var connection string

func init() {
	connection = "MySQL"
}

func Connect() string {
	return connection
}