package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "server/pb/query"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is running"))
	})
	r.Get("/query-response/{message}", func(w http.ResponseWriter, r *http.Request) {
		message := chi.URLParam(r, "message")
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		c := pb.NewResponseServiceClient(conn)

		response, err := c.GetResponse(context.Background(), &pb.RequestBody{Message: message})

		if err != nil {
			log.Fatalf("Error when calling GetResponse: %s", err)

			w.Write([]byte("Error when calling GetResponse"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write([]byte(response.Message))
	})

	http.ListenAndServe(":5000", r)
}
