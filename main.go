package main

import (
	gm "github.com/rbrahul/go-migration/migrator"
)

func main() {
	schema := &gm.Schema{CharSet: "utf8mb4", Collation: "utf8mb4_unicode_ci"}
	schema.Create("employees", func(table *gm.TableManager) {
		table.UnsignedBigInteger("id").Autoincrement()
		table.UnsignedBigInteger("user_id").SetSize(20)
		table.String("name", 10).Comment("Full Name")
		table.Float("salary", 10, 2)
		table.ENUM("gender", []string{"Male", "Female", "Others"})
		//table.Foreign("user_id").Referrences("id").On("users").OnDelete("cascade").OnUpdate("restrict")
		//table.Index([]string{"name"})
		table.String("designation", 50)
		table.Timestamps()
		//table.RenameColumn("salary", "salary_amount")
		table.Foreign("user_id").Referrences("id").On("users").OnDelete("cascade").OnUpdate("restrict")
		table.Index([]string{"name"})
	})
}
