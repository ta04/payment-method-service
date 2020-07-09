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

package usecase

import (
	paymentMethodPB "github.com/ta04/payment-method-service/model/proto"
)

// Repository is the interface of repositories.
// As there are number of repositories can be used.
type Repository interface {
	Index(request *paymentMethodPB.IndexPaymentMethodsRequest) ([]*paymentMethodPB.PaymentMethod, error)
	Show(request *paymentMethodPB.PaymentMethod) (*paymentMethodPB.PaymentMethod, error)
	Store(paymentMethod *paymentMethodPB.PaymentMethod) (*paymentMethodPB.PaymentMethod, error)
	Update(paymentMethod *paymentMethodPB.PaymentMethod) (*paymentMethodPB.PaymentMethod, error)
	Destroy(paymentMethod *paymentMethodPB.PaymentMethod) (*paymentMethodPB.PaymentMethod, error)
}
