package handler

import (
	"context"
	"fmt"
	v1 "teacket/gen/api/v1"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TicketHandler struct {
	DB *pgxpool.Pool
}

func (t TicketHandler) CreateTicket(ctx context.Context, request *v1.CreateTicketRequest) (*v1.CreateTicketResponse, error) {
	var (
		ticket    v1.Ticket
		status    string
		createdAt time.Time
		updatedAt time.Time
	)

	err := t.DB.QueryRow(
		ctx,
		`
        INSERT INTO tickets (title, description, status)
        VALUES ($1, $2, $3)
        RETURNING id, title, description, status, created_at, updated_at
        `,
		request.Title,
		request.Description,
		"open",
	).Scan(
		&ticket.Id,
		&ticket.Title,
		&ticket.Description,
		&status,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create ticket: %w", err)
	}

	switch status {
	case "open":
		ticket.Status = v1.Status_STATUS_OPEN
	case "in_progress":
		ticket.Status = v1.Status_STATUS_IN_PROGRESS
	case "resolved":
		ticket.Status = v1.Status_STATUS_RESOLVED
	case "closed":
		ticket.Status = v1.Status_STATUS_CLOSED
	default:
		ticket.Status = v1.Status_STATUS_UNSPECIFIED
	}

	ticket.CreatedAt = createdAt.Format(time.RFC3339)
	ticket.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &v1.CreateTicketResponse{
		Ticket: &ticket,
	}, nil
}

func (t TicketHandler) GetTicket(ctx context.Context, request *v1.GetTicketRequest) (*v1.GetTicketResponse, error) {
	var (
		ticket    v1.Ticket
		status    string
		createdAt time.Time
		updatedAt time.Time
	)

	err := t.DB.QueryRow(
		ctx,
		`
		SELECT id, title, description, status, created_at, updated_at
		FROM tickets
		WHERE id = $1
		`,
		request.Id,
	).Scan(
		&ticket.Id,
		&ticket.Title,
		&ticket.Description,
		&status,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get ticket: %w", err)
	}

	switch status {
	case "open":
		ticket.Status = v1.Status_STATUS_OPEN
	case "in_progress":
		ticket.Status = v1.Status_STATUS_IN_PROGRESS
	case "resolved":
		ticket.Status = v1.Status_STATUS_RESOLVED
	case "closed":
		ticket.Status = v1.Status_STATUS_CLOSED
	default:
		ticket.Status = v1.Status_STATUS_UNSPECIFIED
	}

	ticket.CreatedAt = createdAt.Format(time.RFC3339)
	ticket.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &v1.GetTicketResponse{
		Ticket: &ticket,
	}, nil
}

func (t TicketHandler) GetAllTickets(ctx context.Context, request *v1.GetAllTicketsRequest) (*v1.GetAllTicketsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TicketHandler) DeleteTicket(ctx context.Context, request *v1.DeleteTicketRequest) (*v1.DeleteTicketResponse, error) {
	var (
		ticket v1.Ticket
	)

	err := t.DB.QueryRow(
		ctx,
		`
        DELETE FROM tickets 
        WHERE id = $1
        RETURNING id
        `,
		request.Id,
	).Scan(&ticket.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to delete ticket: %w", err)
	}

	return &v1.DeleteTicketResponse{}, nil
}
