package handler

import (
	"context"
	v1 "teacket/gen/api/v1"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TicketHandler struct {
	DB *pgxpool.Pool
}

func (t TicketHandler) CreateTicket(ctx context.Context, request *v1.CreateTicketRequest) (*v1.CreateTicketResponse, error) {
	return &v1.CreateTicketResponse{Ticket: &v1.Ticket{
		Title:       request.Title,
		Description: request.Description,
	}}, nil
}

func (t TicketHandler) GetTicket(ctx context.Context, request *v1.GetTicketRequest) (*v1.GetTicketResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TicketHandler) GetAllTickets(ctx context.Context, request *v1.GetAllTicketsRequest) (*v1.GetAllTicketsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TicketHandler) DeleteTicket(ctx context.Context, request *v1.DeleteTicketRequest) (*v1.DeleteTicketResponse, error) {
	//TODO implement me
	panic("implement me")
}
