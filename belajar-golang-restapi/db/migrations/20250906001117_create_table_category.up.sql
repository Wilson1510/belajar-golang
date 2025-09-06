CREATE TABLE category
(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
)ENGINE=InnoDB;

-- migrate -database "mysql://root:password@tcp(localhost:3306)/golang" -path db/migrations up
-- migrate create -ext sql -dir db/migrations create_table_category   