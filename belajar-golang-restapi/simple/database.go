package simple

type Database struct {
	Name string
}

type DatabasePostgres Database
type DatabaseMongoDB Database

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}

func NewDatabasePostgres() *DatabasePostgres {
	return (*DatabasePostgres)(&Database{Name: "Postgres"})
}

type DatabaseRepository struct {
	DatabaseMongoDB *DatabaseMongoDB
	DatabasePostgres *DatabasePostgres
}

func NewDatabaseRepository(mongo *DatabaseMongoDB, postgres *DatabasePostgres) *DatabaseRepository {
	return &DatabaseRepository{DatabaseMongoDB: mongo, DatabasePostgres: postgres}
}
