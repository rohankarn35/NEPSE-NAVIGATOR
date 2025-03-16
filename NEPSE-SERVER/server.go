package main

import (
	"context"
	"nepse-sever-graphql/applog"
	"nepse-sever-graphql/graph"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/ast"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const defaultPort = "8080"

func main() {

	err := applog.InitLogger("app.log", applog.INFO) // Set minimum level to INFO
	if err != nil {
		applog.Log(applog.ERROR, "Failed to initialize logger: %v", err)
		return
	}
	defer applog.CloseLogger()
	err = godotenv.Load()
	if err != nil {
		applog.Log(applog.ERROR, "Error loading .env file: %v", err)
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URL"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		applog.Log(applog.ERROR, "Failed to connect to MongoDB: %v", err)
		return
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		applog.Log(applog.ERROR, "Failed to ping MongoDB: %v", err)
		return
	}

	db := client.Database("nepsedata")

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{MongoClient: db}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	applog.Log(applog.INFO, "connect to http://localhost:%s/ for GraphQL playground", port)
	applog.Log(applog.ERROR, "%v", http.ListenAndServe(":"+port, nil))
}
