// cmd/server/main.go
package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/agnaldojpereira/Clean-Architecture-GO/internal/database"
	"github.com/agnaldojpereira/Clean-Architecture-GO/internal/repository"
	"github.com/agnaldojpereira/Clean-Architecture-GO/internal/service"
	"github.com/agnaldojpereira/Clean-Architecture-GO/pkg/graphql"
	"github.com/agnaldojpereira/Clean-Architecture-GO/pkg/grpc"
	"github.com/agnaldojpereira/Clean-Architecture-GO/pkg/rest"
	"google.golang.org/grpc"
)

func main() {
	// Conexão com o banco de dados
	db, err := database.NewConnection()
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	// Inicialização dos componentes
	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)

	// Configuração do servidor REST
	restHandler := rest.NewHandler(orderService)
	http.HandleFunc("/order", restHandler.ListOrders)

	// Configuração do servidor gRPC
	grpcServer := grpc.NewServer(orderService)
	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Falha ao iniciar o listener gRPC: %v", err)
	}

	// Configuração do servidor GraphQL
	graphqlResolver := graphql.NewResolver(orderService)
	graphqlServer := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: graphqlResolver}))
	http.Handle("/graphql", graphqlServer)
	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))

	// Iniciar servidores em goroutines separadas
	go func() {
		log.Println("Iniciando servidor REST na porta 8080")
		if err := http.ListenAndServe(":8080", nil); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Erro ao iniciar servidor HTTP: %v", err)
		}
	}()

	go func() {
		log.Println("Iniciando servidor gRPC na porta 50051")
		if err := grpcServer.Serve(grpcListener); err != nil {
			log.Fatalf("Erro ao iniciar servidor gRPC: %v", err)
		}
	}()

	// Configuração para graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	log.Println("Desligando servidores...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := http.Server{}.Shutdown(ctx); err != nil {
		log.Printf("Erro ao desligar servidor HTTP: %v", err)
	}

	grpcServer.GracefulStop()

	log.Println("Servidores desligados com sucesso")
}