package db

import (
	"fmt"

	"github.com/pcranaway/biblioteka/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(environment *env.Environment) *gorm.DB {
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s port=%d dbname=%s",

        environment.Postgres.Host,
        environment.Postgres.Username,
        environment.Postgres.Password,

        // default values (might need to make them into environment
        // variables at some point)
        5432,
        "biblioteka",
    )

    db, err := gorm.Open(postgres.Open(dsn))

    if err != nil {
        panic(err)
    }

    return db
}
