package simplecms

import (
   "database/sql"

   // Use the postgres SQL drive
   _ "github.com/lib/pq"
)

type PgStore struct {

   DB *sql.DB

}

func newDB() *PgStore {
     db, err := sql.Open("postgres", "user=dbadmin dbname=simplecms sslmode=disable")

     if err != nil {
      panic(err)
     }

     return &PgStore{
        DB: db,
     }

}
