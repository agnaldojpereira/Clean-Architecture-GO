package main

import (
	"log"
	"net"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/seu-usuario/projeto-pedidos/internal/handler"
	"github.com/seu-usuario/projeto-pedidos/internal/service"
	"github.com/seu-usuario/projeto-pedidos/pkg/db"
	"google.golang.org/grpc"
)

func main() {
	// Inicializar conexão com o banco de dados
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Falha ao inicializar o banco de dados: %v", err)
	}
	defer database.Close()

	// Inicializar serviço de pedidos
	orderService := service.NewOrderService(database)

	// Inicializar handler HTTP
	httpHandler := handler.NewHTTPHandler(orderService)

	// Inicializar servidor gRPC
	grpcServer := grpc.NewServer()
	handler.RegisterOrderServiceServer(grpcServer, handler.NewGRPCHandler(orderService))

	// Inicializar servidor GraphQL
	graphqlHandler := handler.NewGraphQLHandler(orderService)

	// Iniciar servidor HTTP
	go func() {
		http.Handle("/graphql", graphqlHandler)
		http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
		http.Handle("/order", httpHandler)
		log.Printf("Servidor HTTP rodando na porta 8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// Iniciar servidor gRPC
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Falha ao escutar: %v", err)
		}
		log.Printf("Servidor gRPC rodando na porta 50051")
		log.Fatal(grpcServer.Serve(lis))
	}()

	// Manter a goroutine principal viva
	select {}
}
