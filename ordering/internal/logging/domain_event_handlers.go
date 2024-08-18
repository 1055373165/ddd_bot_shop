package logging

import (
	"context"
	"eda-in-go/internal/ddd"
	"eda-in-go/ordering/internal/application"

	"github.com/rs/zerolog"
)

type DomainEventHandlers struct {
	application.DomainEventHandlers
	logger zerolog.Logger
}

var _ application.DomainEventHandlers = (*DomainEventHandlers)(nil)

func LogDomainEventHandlerAccess(handlers application.DomainEventHandlers, logger zerolog.Logger) DomainEventHandlers {
	return DomainEventHandlers{
		DomainEventHandlers: handlers,
		logger:              logger,
	}
}

func (h DomainEventHandlers) OnOrderCreated(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msg("--> Ordering.OnOrderCreated")
	defer func() { h.logger.Info().Err(err).Msg("<-- Ordering.OnOrderCreated") }()
	return h.DomainEventHandlers.OnOrderCreated(ctx, event)
}

func (h DomainEventHandlers) OnOrderReadied(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msg("--> Ordering.OnOrderReadied")
	defer func() { h.logger.Info().Err(err).Msg("<-- Ordering.OnOrderReadied") }()
	return h.DomainEventHandlers.OnOrderReadied(ctx, event)
}

func (h DomainEventHandlers) OnOrderCanceled(ctx context.Context, event ddd.Event) (err error) {
	h.logger.Info().Msg("--> Ordering.OnOrderCanceled")
	defer func() { h.logger.Info().Err(err).Msg("<-- Ordering.OnOrderCanceled") }()
	return h.DomainEventHandlers.OnOrderCanceled(ctx, event)
}
