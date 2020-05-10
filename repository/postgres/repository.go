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
	"fmt"

	paymentMethodPB "github.com/ta04/payment-method-service/proto"
)

// Store stores a new payment method
func (repo *Postgres) Store(paymentMethod *paymentMethodPB.PaymentMethod) (*paymentMethodPB.PaymentMethod, error) {
	query := fmt.Sprintf("INSERT INTO payment_methods (name, accountNumber, accountHolderName, status)"+
		" VALUES ('%s', '%s', '%s', 'active')", paymentMethod.Name, paymentMethod.AccountNumber, paymentMethod.AccountHolderName)
	_, err := repo.DB.Exec(query)

	return paymentMethod, err
}

// Update updates a payment method
func (repo *Postgres) Update(paymentMethod *paymentMethodPB.PaymentMethod) (*paymentMethodPB.PaymentMethod, error) {
	query := fmt.Sprintf("UPDATE payment_methods SET name = '%s', accountNumber = '%s', accountHolderName = '%s', status = 'active'"+
		" WHERE id = %d", paymentMethod.Name, paymentMethod.AccountNumber, paymentMethod.AccountHolderName, paymentMethod.Id)
	_, err := repo.DB.Exec(query)

	return paymentMethod, err
}

// Destroy updates an existing payment method's status to inactive
func (repo *Postgres) Destroy(paymentMethod *paymentMethodPB.PaymentMethod) (*paymentMethodPB.PaymentMethod, error) {
	query := fmt.Sprintf("UPDATE payment_methods SET status = 'inactive' WHERE id = %d", paymentMethod.Id)
	_, err := repo.DB.Exec(query)

	return paymentMethod, err
}
