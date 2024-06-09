package service

import (
	"context"

	"github.com/phelipperibeiro/desafio-clean-architecture/internal/infra/grpc/pb"
	"github.com/phelipperibeiro/desafio-clean-architecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase   usecase.ListOrdersUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrderUseCase usecase.ListOrdersUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrderUseCase:   listOrderUseCase,
	}
}

func (orderService *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {

	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}

	output, err := orderService.CreateOrderUseCase.Execute(dto)

	if err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (orderService *OrderService) ListOrders(ctx context.Context, in *pb.Blank) (*pb.OrdersList, error) {

	output, err := orderService.ListOrderUseCase.Execute()

	if err != nil {
		return nil, err
	}

	var orders []*pb.ListOrdersResponse

	for _, order := range output {
		orders = append(orders, &pb.ListOrdersResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}

	return &pb.OrdersList{Orders: orders}, nil
}
