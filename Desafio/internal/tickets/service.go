package tickets

import (
	"context"

	"github.com/bootcamp-go/desafio-go-web/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
	GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error)
	AverageDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
}
type serviceRepo struct {
	r Repository
}

func NewService(repo Repository) Service {
	return &serviceRepo{repo}
}

func (s *serviceRepo) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	l, err := s.r.GetAll(ctx)
	return l, err

}

// revisar metodos que hacen solo copie el mismo codigo
func (s *serviceRepo) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	l, err := s.r.GetTicketByDestination(ctx, destination)
	return l, err
}

func (s *serviceRepo) GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error) {
	l, err := s.r.GetTicketByDestination(ctx, destination)
	return l, err
}

func (s *serviceRepo) AverageDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	l, err := s.r.GetTicketByDestination(ctx, destination)
	return l, err
}
