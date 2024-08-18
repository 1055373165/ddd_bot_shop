package handlers

import (
	"eda-in-go/depot/internal/application"
	"eda-in-go/depot/internal/domain"
	"eda-in-go/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.ShoppingListCompleted{}, orderHandlers.OnShoppingListCompleted)
}
