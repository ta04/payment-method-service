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

	paymentMethodPB "github.com/ta04/payment-method-service/proto"
	paymentMethodRepo "github.com/ta04/payment-method-service/repository"
)

// Handler is the handler of payment method service
type Handler struct {
	repository paymentMethodRepo.Repository
}

// NewHandler creates a new payment method service handler
func NewHandler(repo paymentMethodRepo.Repository) *Handler {
	return &Handler{
		repository: repo,
	}
}

// IndexPaymentMethods indexes the payment methods
func (h *Handler) IndexPaymentMethods(ctx context.Context, req *paymentMethodPB.IndexPaymentMethodsRequest, res *paymentMethodPB.Response) error {
	paymentMethods, err := h.repository.Index(req)
	if err != nil {
		return err
	}

	res.PaymentMethods = paymentMethods
	res.Error = nil

	return err
}

// ShowPaymentMethod shows a payment method by ID
func (h *Handler) ShowPaymentMethod(ctx context.Context, req *paymentMethodPB.PaymentMethod, res *paymentMethodPB.Response) error {
	paymentMethod, err := h.repository.Show(req)
	if err != nil {
		return err
	}

	res.PaymentMethod = paymentMethod
	res.Error = nil

	return nil
}

// StorePaymentMethod stores a new payment method
func (h *Handler) StorePaymentMethod(ctx context.Context, req *paymentMethodPB.PaymentMethod, res *paymentMethodPB.Response) error {
	paymentMethod, err := h.repository.Store(req)
	if err != nil {
		return err
	}

	res.PaymentMethod = paymentMethod
	res.Error = nil

	return err
}

// UpdatePaymentMethod updates a payment method
func (h *Handler) UpdatePaymentMethod(ctx context.Context, req *paymentMethodPB.PaymentMethod, res *paymentMethodPB.Response) error {
	paymentMethod, err := h.repository.Update(req)
	if err != nil {
		return err
	}

	res.PaymentMethod = paymentMethod
	res.Error = nil

	return nil
}

// DestroyPaymentMethod destroys an payment methods
func (h *Handler) DestroyPaymentMethod(ctx context.Context, req *paymentMethodPB.PaymentMethod, res *paymentMethodPB.Response) error {
	paymentMethod, err := h.repository.Destroy(req)
	if err != nil {
		return err
	}

	res.PaymentMethod = paymentMethod
	res.Error = nil

	return err
}
