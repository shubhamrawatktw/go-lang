package sqllite

import (
	"database/sql"
	"go-license/internal/config"

	_ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

type Sqllite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqllite, error) {

	db, err := sql.Open("sqlite", cfg.StoragePath)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT,
		age INTEGER
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqllite{
		Db: db,
	}, nil

}


func (s *Sqllite) CreateStudent(name string, email string, age int ) (int64,error){

	stmt,err :=s.Db.Prepare("INSERT INTO students (name, email, age) VALUES(?, ?, ?)")
  
	if err != nil {
		return 0,err
	}

	defer stmt.Close()

  result ,err :=stmt.Exec(name,email,age)

  if err != nil {
	return 0,err
}

id , err := result.LastInsertId()
if err != nil {
	return 0,err
}

 return id ,nil
}
