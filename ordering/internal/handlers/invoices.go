package handlers

import (
	"eda-in-go/internal/ddd"
	"eda-in-go/ordering/internal/application"
	"eda-in-go/ordering/internal/domain"
)

func RegisterInvoiceHandlers(invoiceHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.OrderReadied{}, invoiceHandlers.OnOrderReadied)
}
