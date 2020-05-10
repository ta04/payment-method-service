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

package repository

import (
	paymentMethodPB "github.com/ta04/payment-method-service/proto"
)

// Repository is the interface of repositories.
// As there are number of repositories can be used.
type Repository interface {
	Index(*paymentMethodPB.IndexPaymentMethodsRequest) ([]*paymentMethodPB.PaymentMethod, error)
	Show(*paymentMethodPB.PaymentMethod) (*paymentMethodPB.PaymentMethod, error)
	Store(*paymentMethodPB.PaymentMethod) (*paymentMethodPB.PaymentMethod, error)
	Update(*paymentMethodPB.PaymentMethod) (*paymentMethodPB.PaymentMethod, error)
	Destroy(*paymentMethodPB.PaymentMethod) (*paymentMethodPB.PaymentMethod, error)
}
