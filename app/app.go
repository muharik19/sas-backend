package app

import (
	"fmt"
	"net/http"
	"os"
	dbconnect "sas-backend/app/database/connectors/postgres"
	"sas-backend/app/graphql/schema"
	"sas-backend/app/utils"

	"github.com/graphql-go/handler"
)

//App struct
type App struct{}

//Initialize function
func (a *App) Initialize() {
	fmt.Println("Init...")
	db, err := dbconnect.ConnectDb()
	if err == nil {
		defer db.Close()
		fmt.Println("Database Ready...")
	}

	graphqlSchema := schema.InitGraphQLSchema()
	gqlHandler := handler.New(&handler.Config{
		Schema:     &graphqlSchema,
		Pretty:     true,
		Playground: true,
	})
	handlerNonSubscription := utils.PassContext(gqlHandler)
	http.Handle("/", handlerNonSubscription)
}

//Run function
func (a *App) Run() {
	port := ":" + os.Getenv("APP_PORT")
	fmt.Println("GQL SERVICE RUN AT", port)
	fmt.Println(http.ListenAndServe(port, nil))
}
