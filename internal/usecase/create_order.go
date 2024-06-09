package usecase

import (
	"github.com/phelipperibeiro/desafio-clean-architecture/internal/entity"
	"github.com/phelipperibeiro/desafio-clean-architecture/pkg/events"
)

type OrderInputDTO struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type OrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type CreateOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderCreated    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewCreateOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: OrderRepository,
		OrderCreated:    OrderCreated,
		EventDispatcher: EventDispatcher,
	}
}

func (createOrderUseCase *CreateOrderUseCase) Execute(input OrderInputDTO) (OrderOutputDTO, error) {

	order := entity.Order{
		ID:    input.ID,
		Price: input.Price,
		Tax:   input.Tax,
	}

	err := order.CalculateFinalPrice()

	if err != nil {
		return OrderOutputDTO{}, err
	}

	if err := createOrderUseCase.OrderRepository.Save(&order); err != nil {
		return OrderOutputDTO{}, err
	}

	dto := OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.Price + order.Tax,
	}

	createOrderUseCase.OrderCreated.SetPayload(dto)

	err = createOrderUseCase.EventDispatcher.Dispatch(createOrderUseCase.OrderCreated)

	if err != nil {
		return OrderOutputDTO{}, err
	}

	return dto, nil
}
