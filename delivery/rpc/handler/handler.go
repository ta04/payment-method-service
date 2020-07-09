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
	"errors"

	proto "github.com/ta04/payment-method-service/model/proto"
	"github.com/ta04/payment-method-service/usecase"
)

// Handler is the handler of payment method service
type Handler struct {
	Usecase usecase.Usecase
}

// NewHandler will create a new payment method handler
func NewHandler(usecase usecase.Usecase) *Handler {
	return &Handler{
		Usecase: usecase,
	}
}

// GetAllPaymentMethods will get all payment methods
func (handler *Handler) GetAllPaymentMethods(ctx context.Context, req *proto.GetAllPaymentMethodsRequest, res *proto.Response) error {
	paymentMethods, err := handler.Usecase.GetAll(req)
	if err != nil {
		res.PaymentMethods = nil
		res.Error = err

		return errors.New(err.Message)
	}

	res.PaymentMethods = paymentMethods
	res.Error = nil

	return nil
}

// GetOnePaymentMethod will get a payment method
func (handler *Handler) GetOnePaymentMethod(ctx context.Context, req *proto.GetOnePaymentMethodRequest, res *proto.Response) error {
	paymentMethod, err := handler.Usecase.GetOne(req)
	if err != nil {
		res.PaymentMethods = nil
		res.Error = err

		return errors.New(err.Message)
	}

	res.PaymentMethod = paymentMethod
	res.Error = nil

	return nil
}

// CreateOnePaymentMethod will create a new payment method
func (handler *Handler) CreateOnePaymentMethod(ctx context.Context, req *proto.PaymentMethod, res *proto.Response) error {
	paymentMethod, err := handler.Usecase.CreateOne(req)
	if err != nil {
		res.PaymentMethods = nil
		res.Error = err

		return errors.New(err.Message)
	}

	res.PaymentMethod = paymentMethod
	res.Error = nil

	return nil
}

// UpdateOnePaymentMethod will update a payment method
func (handler *Handler) UpdateOnePaymentMethod(ctx context.Context, req *proto.PaymentMethod, res *proto.Response) error {
	paymentMethod, err := handler.Usecase.UpdateOne(req)
	if err != nil {
		res.PaymentMethods = nil
		res.Error = err

		return errors.New(err.Message)
	}

	res.PaymentMethod = paymentMethod
	res.Error = nil

	return nil
}

