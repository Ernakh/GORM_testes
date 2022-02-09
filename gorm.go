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

	fmt.Println("Inserindo...")
	pessoa := Pessoa{Nome: "Fabrício", Email: "fabriciotonettolondero@gmail.com"}
	db.Create(&pessoa)

	//db.Select("Name", "Age", "CreatedAt").Create(&user)//insert apenas em campos específicos

	var pessoas []Pessoa
	fmt.Println("Buscando...")
	db.Find(&pessoas)

	fmt.Println("Listando...")
	fmt.Println(pessoas)
	fmt.Println("Filtrando...")

	var p []Pessoa
	db.Where("email = ?", "fabriciotonettolondero@gmail.com").Find(&p)
	fmt.Println(p)

	/*

		result.RowsAffected // returns found records count, equals `len(users)`
		result.Error        // returns error

		// Get the first record ordered by primary key
		db.First(&user)
		// SELECT * FROM users ORDER BY id LIMIT 1;

		// Get one record, no specified order
		db.Take(&user)
		// SELECT * FROM users LIMIT 1;

		// Get last record, ordered by primary key desc
		db.Last(&user)
		// SELECT * FROM users ORDER BY id DESC LIMIT 1;

		result := db.First(&user)
		result.RowsAffected // returns count of records found
		result.Error        // returns error or nil

		// check error ErrRecordNotFound
		errors.Is(result.Error, gorm.ErrRecordNotFound)

		db.First(&user, 10)
		// SELECT * FROM users WHERE id = 10;

		db.First(&user, "10")
		// SELECT * FROM users WHERE id = 10;

		db.Find(&users, []int{1,2,3})
		// SELECT * FROM users WHERE id IN (1,2,3);


		// Get first matched record
		db.Where("name = ?", "jinzhu").First(&user)
		// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

		// Get all matched records
		db.Where("name <> ?", "jinzhu").Find(&users)
		// SELECT * FROM users WHERE name <> 'jinzhu';

		// IN
		db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
		// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

		// LIKE
		db.Where("name LIKE ?", "%jin%").Find(&users)
		// SELECT * FROM users WHERE name LIKE '%jin%';

		// AND
		db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
		// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

		// Time
		db.Where("updated_at > ?", lastWeek).Find(&users)
		// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

		// BETWEEN
		db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
		// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

		db.Where(&User{Name: "jinzhu", Age: 0}).Find(&users)
		// SELECT * FROM users WHERE name = "jinzhu";

		db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
		// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

		db.Order("age desc, name").Find(&users)
		// SELECT * FROM users ORDER BY age desc, name;

		Group By & Having
		type result struct {
		Date  time.Time
		Total int
		}

		db.Model(&User{}).Select("name, sum(age) as total").Where("name LIKE ?", "group%").Group("name").First(&result)
		// SELECT name, sum(age) as total FROM `users` WHERE name LIKE "group%" GROUP BY `name` LIMIT 1


		db.Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "group").Find(&result)
		// SELECT name, sum(age) as total FROM `users` GROUP BY `name` HAVING name = "group"

		rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Rows()
		for rows.Next() {
		...
		}

		rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Rows()
		for rows.Next() {
		...
		}

		type Result struct {
		Date  time.Time
		Total int64
		}
		db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Scan(&results)

		Joins Preloading
		You can use Joins eager loading associations with a single SQL, for example:

		db.Joins("Company").Find(&users)
		// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id`;
		Join with conditions

		db.Joins("Company", DB.Where(&Company{Alive: true})).Find(&users)
		// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id` AND `Company`.`alive` = true;
	*/

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
