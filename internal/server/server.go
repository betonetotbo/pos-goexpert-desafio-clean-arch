package server

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/graph"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/database"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/pb"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/rest"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ListenAndServe(ctx context.Context, httpPort, grpcPort int) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := ctx.Value(database.DBContextKey).(*gorm.DB)

	ordersSvc := service.NewOrderService(db)

	configureRoutes(r, ordersSvc, db)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", httpPort),
		Handler: r,
	}

	go func() {
		log.Printf("Servidor HTTP executando na porta %d...", httpPort)
		log.Printf("Playground do GraphQL: http://localhost:%d/graph", httpPort)
		e := httpServer.ListenAndServe()
		if e != nil && e != http.ErrServerClosed {
			log.Fatalf("Falha ao iniciar o servidor: %v\n", e)
		}
	}()

	grpcServer := newGrpcServer(ordersSvc)
	go func() {
		log.Printf("Servidor gRPC executando na porta %d...", grpcPort)
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
		if err != nil {
			log.Fatalf("Falha ao configurar o servidor gRPC: %v\n", err)
		}
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatalf("Falha ao iniciar o servidor gRPC: %v\n", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	log.Println("Servidores executando, aguardando sinal de desligamento...")
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	log.Println("Desligando o servidores...")
	e := httpServer.Shutdown(ctx)
	if e != nil {
		log.Fatalf("Falha ao desligar o servidor HTTP: %v\n", e)
	}
	grpcServer.Stop()

	log.Println("Servidores desligado")
}

func configureRoutes(router *chi.Mux, ordersSvc pb.OrderServiceServer, db *gorm.DB) {
	// GraphQL
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Service: ordersSvc, DB: db}}))
	router.Post("/graph/query", h.ServeHTTP)
	router.Get("/graph", playground.Handler("GraphQL", "/graph/query"))

	// REST
	ordersHandler := rest.NewHandler(ordersSvc)
	router.Get("/rest/orders", ordersHandler.ListOrdersHandler)
	router.Post("/rest/orders", ordersHandler.CreateOrderHandler)
}

func newGrpcServer(ordersSvc pb.OrderServiceServer) *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterOrderServiceServer(server, ordersSvc)
	reflection.Register(server)
	return server
}
