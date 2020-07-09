/*
Dear Programmers,

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*                                                 *
*	This file belongs to Kevin Veros Hamonangan   *
*	and	Fandi Fladimir Dachi and is a part of     *
*	our	last project as the student of Del        *
*	Institute of Technology, Sitoluama.           *
*	Please contact us via Instagram:              *
*	sleepingnext and fandi_dachi                  *
*	before copying this file.                     *
*	Thank you, buddy. 😊                          *
*                                                 *
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*/

package handler

import (
	"context"

	paymentMethodPB "github.com/ta04/payment-method-service/model/proto"
	"github.com/ta04/payment-method-service/usecase"
)

// Handler is the handler of payment method service
type Handler struct {
	UseCase usecase.Usecase
}

// NewHandler creates a new payment method service handler
func NewHandler(usecase usecase.Usecase) *Handler {
	return &Handler{
		UseCase: usecase,
	}
}

// IndexPaymentMethods indexes the payment methods
func (handler *Handler) IndexPaymentMethods(ctx context.Context, req *paymentMethodPB.IndexPaymentMethodsRequest, res *paymentMethodPB.Response) error {
	paymentMethods, err := handler.Usecase.Index(req)
	if err != nil {
		return err
	}

	res.PaymentMethods = paymentMethods
	res.Error = nil

	return err
}

// ShowPaymentMethod shows a payment method by ID
func (handler *Handler) ShowPaymentMethod(ctx context.Context, req *paymentMethodPB.PaymentMethod, res *paymentMethodPB.Response) error {
	paymentMethod, err := handler.Usecase.Show(req)
	if err != nil {
		return err
	}

	res.PaymentMethod = paymentMethod
	res.Error = nil

	return nil
}

// StorePaymentMethod stores a new payment method
func (handler *Handler) StorePaymentMethod(ctx context.Context, req *paymentMethodPB.PaymentMethod, res *paymentMethodPB.Response) error {
	paymentMethod, err := handler.Usecase.Store(req)
	if err != nil {
		return err
	}

	res.PaymentMethod = paymentMethod
	res.Error = nil

	return err
}

// UpdatePaymentMethod updates a payment method
func (handler *Handler) UpdatePaymentMethod(ctx context.Context, req *paymentMethodPB.PaymentMethod, res *paymentMethodPB.Response) error {
	paymentMethod, err := handler.Usecase.Update(req)
	if err != nil {
		return err
	}

	res.PaymentMethod = paymentMethod
	res.Error = nil

	return nil
}

// DestroyPaymentMethod destroys an payment methods
func (handler *Handler) DestroyPaymentMethod(ctx context.Context, req *paymentMethodPB.PaymentMethod, res *paymentMethodPB.Response) error {
	paymentMethod, err := handler.Usecase.Destroy(req)
	if err != nil {
		return err
	}

	res.PaymentMethod = paymentMethod
	res.Error = nil

	return err
}
