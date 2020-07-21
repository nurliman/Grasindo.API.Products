package protos

import (
	"context"
	"log"
)

// Server struct
type Server struct{}

// Create Create a Product
func (server *Server) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	log.Printf("Receive request from")

	return &CreateResponse{Api: "v1", Id: 0}, nil
}
