package mahasiswa

import (
	"context"
	"fmt"
	"api-crud-go/config"
	"api-crud-go/models"
	"log"
	"time"
)

const (
	table          = "go_mahasiswa"
	layoutDateTime = "2006-01-02 15:04:05"
)



// GetAll
func GetAll(ctx context.Context) ([]models.Mahasiswa, error) {

	var mahasiswas []models.Mahasiswa

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var mahasiswa models.Mahasiswa
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&mahasiswa.ID,
			&mahasiswa.NIM,
			&mahasiswa.Name,
			&mahasiswa.Semester,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		mahasiswa.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		mahasiswa.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		mahasiswas = append(mahasiswas, mahasiswa)
	}
	return mahasiswas, nil
}

// Insert
func Insert(ctx context.Context, mhs models.Mahasiswa) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	fmt.Println("query")

	queryText := fmt.Sprintf("INSERT INTO %v (nim, name, semester, created_at, updated_at) values(%v,'%v',%v,'%v','%v')", table,
		mhs.NIM,
		mhs.Name,
		mhs.Semester,
		time.Now().Format(layoutDateTime),
		time.Now().Format(layoutDateTime))

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}