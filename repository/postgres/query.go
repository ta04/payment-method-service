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
	"database/sql"
	"fmt"

	paymentMethodPB "github.com/ta04/payment-method-service/proto"
)

// Postgres is the implementor of Postgres interface
type Postgres struct {
	DB *sql.DB
}

// Index indexes all active payment methods
func (repo *Postgres) Index(req *paymentMethodPB.IndexPaymentMethodsRequest) (paymentMethods []*paymentMethodPB.PaymentMethod, err error) {
	var id int32
	var name, accountNumber, accountHolderName, status string

	query := "SELECT * FROM payment_methods WHERE status = 'active'"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &name, &accountNumber, &accountHolderName, &status)
		if err != nil {
			return nil, err
		}
		paymentMethod := &paymentMethodPB.PaymentMethod{
			Id:                id,
			Name:              name,
			AccountNumber:     accountNumber,
			AccountHolderName: accountHolderName,
			Status:            status,
		}
		paymentMethods = append(paymentMethods, paymentMethod)
	}

	return paymentMethods, err
}

// Show shows an active payment method by id
func (repo *Postgres) Show(paymentMethod *paymentMethodPB.PaymentMethod) (*paymentMethodPB.PaymentMethod, error) {
	var id int32
	var name, accountNumber, accountHolderName, status string

	query := fmt.Sprintf("SELECT * FROM payment_methods WHERE id = %d AND status = 'active'", paymentMethod.Id)
	err := repo.DB.QueryRow(query).Scan(&id, &name, &accountNumber, &accountHolderName, &status)
	if err != nil {
		return nil, err
	}

	return &paymentMethodPB.PaymentMethod{
		Id:                id,
		Name:              name,
		AccountNumber:     accountNumber,
		AccountHolderName: accountHolderName,
		Status:            status,
	}, err
}
