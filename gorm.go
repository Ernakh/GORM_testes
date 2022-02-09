package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Pessoa struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

var db = gorm.DB{}

func main() {
	//go mod init gorm/m
	//go mod tidy

	fmt.Println("Iniciando...")

	dbURL := "postgres://postgres:postgres@localhost:5432/go"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		fmt.Printf("erro ---------------------")
	}

	db.AutoMigrate(&Pessoa{}) //cria a tabela

	var pessoas []Pessoa

	fmt.Println("Buscando...")

	db.Find(&pessoas)

	fmt.Println("Listando...")

	fmt.Println(pessoas)

	fmt.Println("Finalizando...")
}

// func buscaPessoas() {
// 	var pessoas []Pessoa

// 	fmt.Println("Buscando...")

// 	db.Find(&pessoas)
// 	fmt.Println("{}", pessoas)
// }

// func buscaPessoasPeloID() {

// }

// func cadastraPessoa() {

// }
