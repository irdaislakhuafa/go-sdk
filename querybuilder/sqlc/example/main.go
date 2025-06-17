package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/irdaislakhuafa/go-sdk/db"
	"github.com/irdaislakhuafa/go-sdk/querybuilder/sqlc"
	"github.com/irdaislakhuafa/go-sdk/querybuilder/sqlc/example/gen"
)

type (
	ListUserParams struct {
		IDs       []int64
		Search    string
		IsDeleted int
		Ages      []int64
		MinAge    int64
		MaxAge    int64
		Limit     int
		Page      int
		OrderBy   string
		OrderType string
	}
)

func main() {
	ctx := context.Background()

	// connect to db
	db, err := db.Init(db.Config{
		Driver:  db.DriverSQLite,
		User:    "example",
		Options: db.Options{InMemory: true},
	})
	if err != nil {
		panic(err)
	}

	// execute schema query [only for this example, don't write it here in production]
	schema, err := os.ReadFile("./schema.sql")
	if err != nil {
		panic(err)
	}
	_, err = db.ExecContext(ctx, string(schema))
	if err != nil {
		panic(err)
	}

	// wrap sqlc generated code with query builder. query builder won't work without this
	queries := gen.New(sqlc.Wrap(db, sqlc.WrappedOpts{
		ShowQuery: true, // disable in production
		ShowArgs:  true, // disable in production
	}))

	// create data
	for i := 1; i <= 5; i++ {
		_, err := queries.CreateUser(ctx, gen.CreateUserParams{
			Name:  "John doe " + strconv.Itoa(i),
			Email: fmt.Sprintf("jhondoe[%d]@gmail.com", i),
			Age:   int64(20 + i),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println()
	}

	// querying data with dynamic filter
	params := ListUserParams{
		IDs:       []int64{},
		Search:    "doe 4",
		IsDeleted: 0,
		Ages:      []int64{},
		MinAge:    20,
		MaxAge:    25,
		Limit:     1,
		Page:      0,
		OrderBy:   "id",
		OrderType: "desc",
	}
	users, err := queries.ListUser(sqlc.Build(ctx, func(b *sqlc.Builder) {
		b.GroupBy("id")

		// if params.IDs is not empty then filter by params.IDs
		if len(params.IDs) > 0 {
			_, args := sqlc.GenQueryArgs(ctx, params.IDs...)
			b.In("id", args...)
		}

		// AND if params.Ages is not empty then filter by params.Ages
		if len(params.Ages) > 0 {
			_, args := sqlc.GenQueryArgs(ctx, params.Ages...)
			b.In("age", args...)
		}

		// AND filter by `age` that greater than or equal params.MinAge
		b.And("age >= ?", params.MinAge)

		// AND if params.MaxAge is greater than 0 then filter by `age` that lower than or equal params.MaxAge
		if params.MaxAge > 0 {
			b.And("age <= ?", params.MaxAge)
		}

		// AND filter user deleted/undeleted
		b.And("is_deleted = ?", params.IsDeleted)

		// AND if params.Search is not empty then search by name/email
		if params.Search != "" {
			params.Search = "%" + params.Search + "%"
			b.And("(name LIKE ? OR email LIKE ?)", params.Search, params.Search)
		}

		b.Limit(params.Limit)
		b.Offset(params.Page)
		b.Order(params.OrderBy + " " + params.OrderType)
	}))
	if err != nil {
		panic(err)
	}
	fmt.Println()

	fmt.Printf("users: %+v\n", users) // expected: [{ID:4 Name:John doe 4 Email:jhondoe[4]@gmail.com Age:24}]
}
