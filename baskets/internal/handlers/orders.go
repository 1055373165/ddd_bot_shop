package handlers

import (
	"eda-in-go/baskets/internal/application"
	"eda-in-go/baskets/internal/domain"
	"eda-in-go/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.BasketCheckedOut{}, orderHandlers.OnBasketCheckedOut)
}
