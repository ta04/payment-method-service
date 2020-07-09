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
	proto "github.com/ta04/payment-method-service/model/proto"
)

// Repository is the interface of repositories.
// As there are number of repositories can be used.
type Usecase interface {
	GetAll(request *proto.GetAllPaymentMethodsRequest) ([]*proto.PaymentMethod, *proto.Error)
	GetOne(request *proto.GetOnePaymentMethodRequest) (*proto.PaymentMethod, *proto.Error)
	CreateOne(paymentMethod *proto.PaymentMethod) (*proto.PaymentMethod, *proto.Error)
	UpdateOne(paymentMethod *proto.PaymentMethod) (*proto.PaymentMethod, *proto.Error)
}
