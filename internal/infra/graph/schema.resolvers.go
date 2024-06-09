package graph

import (
	"context"

	"github.com/phelipperibeiro/desafio-clean-architecture/internal/infra/graph/model"
	"github.com/phelipperibeiro/desafio-clean-architecture/internal/usecase"
)

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input *model.OrderInput) (*model.Order, error) {
	dto := usecase.OrderInputDTO{
		ID:    input.ID,
		Price: float64(input.Price),
		Tax:   float64(input.Tax),
	}
	output, err := r.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &model.Order{
		ID:         output.ID,
		Price:      float64(output.Price),
		Tax:        float64(output.Tax),
		FinalPrice: float64(output.FinalPrice),
	}, nil
}

// ListOrders is the resolver for the ListOrders field.
func (r *queryResolver) ListOrders(ctx context.Context) ([]*model.Order, error) {
	output, err := r.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for _, o := range output {
		var order model.Order
		order.ID = o.ID
		order.Tax = o.Tax
		order.Price = o.Price
		order.FinalPrice = o.FinalPrice
		orders = append(orders, &order)
	}
	return orders, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	output, err := r.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for _, o := range output {
		var order model.Order
		order.ID = o.ID
		order.Tax = o.Tax
		order.Price = o.Price
		order.FinalPrice = o.FinalPrice
		orders = append(orders, &order)
	}
	return orders, nil
}