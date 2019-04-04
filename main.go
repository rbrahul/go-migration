package main

import (
	"fmt"

	gm "github.com/rbrahul/go-migration/migrator"
)

func main() {
	fmt.Println("Hello world")
	db := &gm.Database{User: "root", Password: "mysql", Host: "localhost", DatabaseName: "laravel_blog"}
	db.New()
	tupleInf, err := db.TupleDefinition("users", "Email")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", tupleInf)
	fmt.Println("****************")
	schema := &gm.Schema{CharSet: "utf8mb4", Collation: "utf8mb4_unicode_ci", Database: db}

	schema.Table("users", func(table *gm.TableManager) {
		table.BigInteger("Email").SetSize(10, 2).Default("Hello Rahul").Unsigned().Change()
		table.String("user_name", 10).UseCurrent().Default("20").Comment("Important data")
		table.String("amount", 10).Default("1000").Comment("Expensive data")
		table.ENUM("category", []string{"Xx", "Xl", "M"}).Default("Xl").Comment("Shirt Size")
		table.Foreign("user_id").Referrences("id").On("Products").OnDelete("cascade").OnUpdate("restrict")
		table.BigInteger("ID").Autoincrement().Unsigned()
		table.String("comment").Collation("acci_general_ci")
		table.Index([]string{"Name", "Email"})
		table.DropColumns([]string{"Name", "Email"})
		table.DropForeign("user_id")
		table.Timestamps()
	})
}
