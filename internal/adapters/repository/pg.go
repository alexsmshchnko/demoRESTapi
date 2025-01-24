package repository

import (
	"database/sql"
	"fmt"
	"os"

	"demorestapi/internal/entity"

	_ "github.com/lib/pq"
)

type Repository struct {
	*sql.DB
}

func ConnectDB() (db Repository, err error) {
	const (
		host     = "localhost"
		port     = "5430"
		user     = "postgres_user"
		password = "postgres_password"
		dbname   = "postgres_db"
	)

	pghost, f := os.LookupEnv("POSTGRES_HOST")
	if !f {
		pghost = host
	}
	pgport, f := os.LookupEnv("POSTGRES_PORT")
	if !f {
		pgport = port
	}
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", pghost, pgport, user, password, dbname)
	fmt.Println(psqlconn)

	if db.DB, err = sql.Open("postgres", psqlconn); err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	return
}

func (pg *Repository) GetUser(id string) *entity.User {
	row := pg.QueryRow("SELECT id, firstname, lastname, email, age, created FROM public.user WHERE id = $1", id)

	u := &entity.User{}
	if err := row.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email, &u.Age, &u.Created); err != nil {
		fmt.Println(err)
	}

	return u
}

func (pg *Repository) AddUser(u *entity.User) (err error) {
	if err = pg.QueryRow("INSERT INTO public.user (id, firstname, lastname, email, age, created) VALUES (gen_random_uuid(), $1, $2, $3, $4, now()) RETURNING id, created",
		u.Firstname, u.Lastname, u.Email, u.Age).Scan(&u.ID, &u.Created); err != nil {
		fmt.Println(err)
	}

	return
}

func (pg *Repository) UpdateUser(u *entity.User) (err error) {
	if err = pg.QueryRow("UPDATE public.user SET firstname=$1, lastname=$2, email=$3, age=$4, created=now() WHERE id = $5 RETURNING created",
		u.Firstname, u.Lastname, u.Email, u.Age, u.ID).Scan(&u.Created); err != nil {
		fmt.Println(err)
	}

	return
}
