package main

import (
	"context"

	auth "github.com/zjhM3l/go-e-commerce/kitex_gen/auth"
	cart "github.com/zjhM3l/go-e-commerce/kitex_gen/cart"
	checkout "github.com/zjhM3l/go-e-commerce/kitex_gen/checkout"
	order "github.com/zjhM3l/go-e-commerce/kitex_gen/order"
	payment "github.com/zjhM3l/go-e-commerce/kitex_gen/payment"
	product "github.com/zjhM3l/go-e-commerce/kitex_gen/product"
	user "github.com/zjhM3l/go-e-commerce/kitex_gen/user"
)

// CartServiceImpl implements the cart service interface defined in the IDL.
type CartServiceImpl struct{}

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// DeliverTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, request *auth.DeliverTokenReq) (resp *auth.DeliverTokenResp, err error) {
	// TODO: Your code here...
	return
}

// VerifyTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyTokenByRPC(ctx context.Context, request *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// TODO: Your code here...
	return
}

// AddItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddItem(ctx context.Context, request *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// TODO: Your code here...
	return
}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, request *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// TODO: Your code here...
	return
}

// EmptyCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) EmptyCart(ctx context.Context, request *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	// TODO: Your code here...
	return
}

// UserServiceImpl implements the user service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, request *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, request *user.LoginReq) (resp *user.LoginResp, err error) {
	// TODO: Your code here...
	return
}

// ProductCatalogServiceImpl implements the product catalog service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, request *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// TODO: Your code here...
	return
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, request *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// TODO: Your code here...
	return
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, request *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// TODO: Your code here...
	return
}

// PaymentServiceImpl implements the payment service interface defined in the IDL.
type PaymentServiceImpl struct{}

// Charge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Charge(ctx context.Context, request *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// TODO: Your code here...
	return
}

// OrderServiceImpl implements the order service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, request *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// TODO: Your code here...
	return
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, request *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// TODO: Your code here...
	return
}

// MarkOrderPaid implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, request *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// TODO: Your code here...
	return
}

// CheckoutServiceImpl implements the checkout service interface defined in the IDL.
type CheckoutServiceImpl struct{}

// Checkout implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) Checkout(ctx context.Context, request *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// TODO: Your code here...
	return
}
