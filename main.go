package main

import (
	"fmt"

	gm "github.com/rbrahul/go-migration/migrator"
)

func main() {
	fmt.Println("Hello world")
	schema := &gm.Schema{}
	schema.Create("users", func(table *gm.TableManager) {
		table.String("user_name", 10).UseCurrent().Default(20).Comment("Important data")
		table.String("amount", 10).Default(1000).Comment("Expensive data")
		table.ENUM("category", []string{"Xx", "Xl", "M"}).Default("Xl").Comment("Shirt Size")
		table.Foreign("user_id").Referrences("id").On("users").OnDelete("cascade").OnUpdate("restrict")
		//table.FindExistingTuple()
		table.Timestamps()
		table.AllCommands()
	})
}
