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

package postgres

import (
	"errors"
	"fmt"

	proto "github.com/ta04/payment-method-service/model/proto"
)

// CreateOne will create a new payment method
func (postgres *Postgres) CreateOne(paymentMethod *proto.PaymentMethod) (*proto.PaymentMethod, error) {
	query := fmt.Sprintf("INSERT INTO payment_methods (name, account_number, account_holder_name, status)"+
		" VALUES ('%s', '%s', '%s', 'active')", paymentMethod.Name, paymentMethod.AccountNumber, paymentMethod.AccountHolderName)
	_, err := postgres.DB.Exec(query)
	if err != nil {
		return nil, err
	}

	return paymentMethod, err
}

// UpdateOne will update a payment method
func (postgres *Postgres) UpdateOne(paymentMethod *proto.PaymentMethod) (*proto.PaymentMethod, error) {
	query := fmt.Sprintf("UPDATE payment_methods SET name = '%s', account_number = '%s', account_holder_name = '%s', status = '%s'"+
		" WHERE id = %d", paymentMethod.Name, paymentMethod.AccountNumber, paymentMethod.AccountHolderName, paymentMethod.Status, paymentMethod.Id)
	res, err := postgres.DB.Exec(query)
	if err != nil {
		return nil, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if count <= 0 {
		return nil, errors.New("sql: no rows found")
	}

	return paymentMethod, err
}
