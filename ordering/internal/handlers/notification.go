package handlers

import (
	"eda-in-go/internal/ddd"
	"eda-in-go/ordering/internal/application"
	"eda-in-go/ordering/internal/domain"
)

func RegisterNotifcationHandlers(notificationsHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.OrderCreated{}, notificationsHandlers.OnOrderCreated)
	domainSubscriber.Subscribe(domain.OrderCanceled{}, notificationsHandlers.OnOrderCanceled)
	domainSubscriber.Subscribe(domain.OrderReadied{}, notificationsHandlers.OnOrderReadied)
}
