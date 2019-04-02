package main

import (
	"fmt"

	gm "github.com/rbrahul/go-migration/migrator"
)

func main() {
	fmt.Println("Hello world")
	schema := &gm.Schema{}

	schema.Table("users", func(table *gm.TableManager) {
		table.String("user_name", 10).UseCurrent().Default("20").Comment("Important data")
		table.String("amount", 10).Default("1000").Comment("Expensive data")
		table.ENUM("category", []string{"Xx", "Xl", "M"}).Default("Xl").Comment("Shirt Size")
		table.Foreign("user_id").Referrences("id").On("Products").OnDelete("cascade").OnUpdate("restrict")
		//table.FindExistingTuple()
		table.BigInteger("ID").Autoincrement().Unsigned()
		table.String("comment").Collation("acci_general_ci")
		table.Index([]string{"Name", "Email"})
		table.DropColumns([]string{"Name", "Email"})
		table.DropForeign("user_id")
		table.Timestamps()
		//table.AllCommands()
	})
}
