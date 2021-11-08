package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dorakueyon/GraphQL-gRPC-demo/article/client"
	"github.com/dorakueyon/GraphQL-gRPC-demo/graph"
	"github.com/dorakueyon/GraphQL-gRPC-demo/graph/generated"
)

const defaultgRPCHost = "localhost"
const defaultgRPCPort = "50051"
const defaultGraphqlHost = "localhost"
const defaultGraphqlPort = "8080"

func main() {
	graphqlPort := os.Getenv("GRAPHQL_PORT")
	if graphqlPort == "" {
		graphqlPort = defaultGraphqlPort
	}

	// articleClientを生成
	grpcHost := os.Getenv("GRPC_HOST")
	if grpcHost == "" {
		grpcHost = defaultgRPCHost
	}
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = defaultgRPCPort
	}

	gRPCDns := grpcHost + ":" + grpcPort
	articleClient, err := client.NewClient(gRPCDns)
	if err != nil {
		articleClient.Close()
		log.Fatalf("Failed to create article client: %v\n", err)
	}
	log.Printf("request to  http://%s/ for gRPC server", gRPCDns)

	// GraphQLサーバーにResolverを登録
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					ArticleClient: articleClient,
				}}))

	// GraphQL playgroundのエンドポイント
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	//　実装したqueryが実行可能なサーバーのエンドポイント
	http.Handle("/query", srv)

	// GraphQLサーバーを起動
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", graphqlPort)
	log.Fatal(http.ListenAndServe(":"+graphqlPort, nil))
}
