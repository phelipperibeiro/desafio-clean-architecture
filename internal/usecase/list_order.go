package usecase

import "github.com/phelipperibeiro/desafio-clean-architecture/internal/entity"

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() ([]OrderOutputDTO, error) {

	orders, err := c.OrderRepository.List()

	if err != nil {
		return nil, err
	}

	var ordersOutput []OrderOutputDTO

	for _, order := range orders {
		ordersOutput = append(ordersOutput, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		})
	}

	return ordersOutput, nil
}
