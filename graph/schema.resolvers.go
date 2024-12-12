package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.60

import (
	"context"
	"fmt"

	"github.com/yaninyzwitty/gqlgen-eccomerce-project/graph/model"
)

// CreateCustomer is the resolver for the createCustomer field.
func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.NewCustomerInput) (*model.Customer, error) {
	panic(fmt.Errorf("not implemented: CreateCustomer - createCustomer"))
}

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProductInput) (*model.Product, error) {
	panic(fmt.Errorf("not implemented: CreateProduct - createProduct"))
}

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input model.NewOrderInput) (*model.Order, error) {
	panic(fmt.Errorf("not implemented: CreateOrder - createOrder"))
}

// CreateOrderItem is the resolver for the createOrderItem field.
func (r *mutationResolver) CreateOrderItem(ctx context.Context, input model.NewOrderItemInput) (*model.OrderItem, error) {
	panic(fmt.Errorf("not implemented: CreateOrderItem - createOrderItem"))
}

// Customer is the resolver for the customer field.
func (r *orderResolver) Customer(ctx context.Context, obj *model.Order) (*model.Customer, error) {
	panic(fmt.Errorf("not implemented: Customer - customer"))
}

// Items is the resolver for the items field.
func (r *orderResolver) Items(ctx context.Context, obj *model.Order) ([]*model.OrderItem, error) {
	panic(fmt.Errorf("not implemented: Items - items"))
}

// Customers is the resolver for the customers field.
func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	panic(fmt.Errorf("not implemented: Customers - customers"))
}

// Customer is the resolver for the customer field.
func (r *queryResolver) Customer(ctx context.Context, id string) (*model.Customer, error) {
	panic(fmt.Errorf("not implemented: Customer - customer"))
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented: Products - products"))
}

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	panic(fmt.Errorf("not implemented: Product - product"))
}

// OrdersByCustomer is the resolver for the ordersByCustomer field.
func (r *queryResolver) OrdersByCustomer(ctx context.Context, customerID string) ([]*model.Order, error) {
	panic(fmt.Errorf("not implemented: OrdersByCustomer - ordersByCustomer"))
}

// OrderItemsByOrder is the resolver for the orderItemsByOrder field.
func (r *queryResolver) OrderItemsByOrder(ctx context.Context, orderID string) ([]*model.OrderItem, error) {
	panic(fmt.Errorf("not implemented: OrderItemsByOrder - orderItemsByOrder"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Order returns OrderResolver implementation.
func (r *Resolver) Order() OrderResolver { return &orderResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type orderResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
