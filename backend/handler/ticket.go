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

const ticketColumns = `id, title, content, author_id, status, created_at, updated_at`

func statusFromDB(s string) v1.Status {
	switch s {
	case "open":
		return v1.Status_STATUS_OPEN
	case "in_progress":
		return v1.Status_STATUS_IN_PROGRESS
	case "resolved":
		return v1.Status_STATUS_RESOLVED
	case "closed":
		return v1.Status_STATUS_CLOSED
	default:
		return v1.Status_STATUS_UNSPECIFIED
	}
}

type scanner interface {
	Scan(dest ...any) error
}

func scanTicket(row scanner) (*v1.Ticket, error) {
	var (
		ticket    v1.Ticket
		status    string
		createdAt time.Time
		updatedAt time.Time
	)

	err := row.Scan(
		&ticket.Id,
		&ticket.Title,
		&ticket.Content,
		&ticket.AuthorId,
		&status,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	ticket.Status = statusFromDB(status)
	ticket.CreatedAt = createdAt.Format(time.RFC3339)
	ticket.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &ticket, nil
}

func (t TicketHandler) CreateTicket(ctx context.Context, request *v1.CreateTicketRequest) (*v1.CreateTicketResponse, error) {
	ticket, err := scanTicket(t.DB.QueryRow(
		ctx,
		`INSERT INTO tickets (title, content, status, author_id)
        VALUES ($1, $2, $3, $4)
        RETURNING `+ticketColumns,
		request.Title,
		request.Content,
		"open",
		1,
	))

	if err != nil {
		return nil, fmt.Errorf("failed to create ticket: %w", err)
	}

	return &v1.CreateTicketResponse{
		Ticket: ticket,
	}, nil
}

func (t TicketHandler) GetTicket(ctx context.Context, request *v1.GetTicketRequest) (*v1.GetTicketResponse, error) {
	ticket, err := scanTicket(t.DB.QueryRow(
		ctx,
		`SELECT `+ticketColumns+`FROM tickets WHERE id = $1`,
		request.Id,
	))

	if err != nil {
		return nil, fmt.Errorf("failed to get ticket: %w", err)
	}

	return &v1.GetTicketResponse{
		Ticket: ticket,
	}, nil
}

func (t TicketHandler) GetAllTickets(ctx context.Context, request *v1.GetAllTicketsRequest) (*v1.GetAllTicketsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TicketHandler) DeleteTicket(ctx context.Context, request *v1.DeleteTicketRequest) (*v1.DeleteTicketResponse, error) {
	var id int64

	err := t.DB.QueryRow(
		ctx,
		`DELETE FROM tickets WHERE id = $1 RETURNING id`,
		request.Id,
	).Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("failed to delete ticket: %w", err)
	}

	return &v1.DeleteTicketResponse{}, nil
}

func (t TicketHandler) UpdateStatus(ctx context.Context, request *v1.UpdateStatusRequest) (*v1.UpdateStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}
