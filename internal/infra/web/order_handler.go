package web

import (
	"encoding/json"
	"net/http"

	"github.com/phelipperibeiro/desafio-clean-architecture/internal/entity"
	"github.com/phelipperibeiro/desafio-clean-architecture/internal/usecase"
	"github.com/phelipperibeiro/desafio-clean-architecture/pkg/events"
)

type WebOrderHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	OrderRepository   entity.OrderRepositoryInterface
	OrderCreatedEvent events.EventInterface
}

func NewWebOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreatedEvent events.EventInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		EventDispatcher:   EventDispatcher,
		OrderRepository:   OrderRepository,
		OrderCreatedEvent: OrderCreatedEvent,
	}
}

func (h *WebOrderHandler) Create(responseWriter http.ResponseWriter, request *http.Request) {

	var dto usecase.OrderInputDTO

	err := json.NewDecoder(request.Body).Decode(&dto)

	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)

	output, err := createOrder.Execute(dto)

	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(responseWriter).Encode(output)

	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebOrderHandler) List(responseWriter http.ResponseWriter, request *http.Request) {

	listOrder := usecase.NewListOrdersUseCase(h.OrderRepository)

	output, err := listOrder.Execute()

	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(responseWriter).Encode(output)

	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
