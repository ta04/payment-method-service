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

	proto "github.com/ta04/payment-method-service/model/proto"
)

// Postgres is the implementor of Postgres interface
type Postgres struct {
	DB *sql.DB
}

// NewPostgres will create a new postgres instance
func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{
		DB: db,
	}
}

// GetAll will get all payment methods
func (postgres *Postgres) GetAll(request *proto.GetAllPaymentMethodsRequest) ([]*proto.PaymentMethod, error) {
	var id int32
	var name, accountNumber, accountHolderName, status string
	var paymentMethods []*proto.PaymentMethod

	query := fmt.Sprintf("SELECT * FROM payment_methods WHERE status = '%s'", request.Status)
	rows, err := postgres.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &name, &accountNumber, &accountHolderName, &status)
		if err != nil {
			return nil, err
		}
		paymentMethod := &proto.PaymentMethod{
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

// GetOne will get a payment method by id
func (postgres *Postgres) GetOne(request *proto.GetOnePaymentMethodRequest) (*proto.PaymentMethod, error) {
	var id int32
	var name, accountNumber, accountHolderName, status string

	query := fmt.Sprintf("SELECT * FROM payment_methods WHERE id = %d", request.Id)
	err := postgres.DB.QueryRow(query).Scan(&id, &name, &accountNumber, &accountHolderName, &status)
	if err != nil {
		return nil, err
	}

	return &proto.PaymentMethod{
		Id:                id,
		Name:              name,
		AccountNumber:     accountNumber,
		AccountHolderName: accountHolderName,
		Status:            status,
	}, err
}
