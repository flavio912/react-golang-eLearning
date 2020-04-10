package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/handler"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/loader"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/models"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/schema"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/uploads"

	"github.com/golang/glog"
	graphql "github.com/graph-gophers/graphql-go"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/database/migration"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/helpers"
	"gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/resolvers"
)

var (
	// We can pass an option to the schema so we don’t need to
	// write a method to access each type’s field:
	opts = []graphql.SchemaOpt{graphql.UseFieldResolvers()}
)

// Reads and parses the schema from file. Associates resolver. Panics if can't read.
func parseSchema(schemaString string, resolver interface{}) *graphql.Schema {
	parsedSchema, err := graphql.ParseSchema(schemaString, resolver, opts...)
	if err != nil {
		panic(err)
	}
	return parsedSchema
}

// If admin exists with current email, update its fields, otherwise create a new one
func updateOrCreateDevAdmin() {
	adminUser := database.GormDB.Where("email = ?", helpers.Config.DevAdmin.Email).First(&models.Admin{})
	newAdmin := &models.Admin{
		Email:     helpers.Config.DevAdmin.Email,
		Password:  helpers.Config.DevAdmin.Password,
		FirstName: helpers.Config.DevAdmin.FirstName,
		LastName:  helpers.Config.DevAdmin.LastName,
	}
	if !adminUser.RecordNotFound() {
		// If found update with new values
		adminUser.Update(newAdmin)

		glog.Info("DevAdmin account found and updated")
		return
	}
	err := database.GormDB.Create(newAdmin)
	if err != nil {
		glog.Error("Unable to create admin user")
	}
	glog.Info("Successfully created admin user")
}

func usage() {
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "INFO")
	flag.Parse()

	// Load in the config.yaml file
	if err := helpers.LoadConfig("config.yml"); err != nil {
		panic(err)
	}
	errDb := database.SetupDatabase()
	if errDb != nil {
		panic(errDb)
	}
	defer database.GormDB.Close()

	migration.InitMigrations()
	uploads.Initialize()

	loaders := loader.Init()

	// Setup DevAdmin user
	updateOrCreateDevAdmin()

	_schema := parseSchema(schema.String(), &resolvers.RootResolver{})

	handle := handler.GraphQL{
		Schema:  _schema,
		Loaders: loaders,
	}
	http.Handle("/graphql", handler.AuthHandler(handle.Serve()))

	log.Println("serving on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	glog.Flush()
}
