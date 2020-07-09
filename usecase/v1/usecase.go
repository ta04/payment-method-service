package v1

import (
	"net/http"

	proto "github.com/ta04/payment-method-service/model/proto"
	"github.com/ta04/payment-method-service/repository"
)

//Usecase is the struct of payment method usecase
type Usecase struct {
	Repository repository.Repository
}

// NewUsecase will create a new payment method usecase
func NewUsecase(repository repository.Repository) *Usecase {
	return &Usecase{
		Repository: repository,
	}
}

func (usecase *Usecase) GetAll(request *proto.GetAllPaymentMethodsRequest) ([]*proto.PaymentMethod, *proto.Error) {
	if request == nil {
		return nil, &proto.Error{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}
	}

	var paymentMethod []*proto.PaymentMethod
	var err error
	paymentMethod, err = usecase.Repository.GetAll(request)
	if err != nil {
		return nil, &proto.Error{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	return paymentMethod, nil
}

func (usecase *Usecase) GetOne(request *proto.GetOnePaymentMethodRequest) (*proto.PaymentMethod, *proto.Error) {
	if request == nil || (request.Id == 0) {
		return nil, &proto.Error{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}
	}

	paymentMethod, err := usecase.Repository.GetOne(request)
	if err != nil {
		return nil, &proto.Error{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	return paymentMethod, nil
}

func (usecase *Usecase) CreateOne(paymentMethod *proto.PaymentMethod) (*proto.PaymentMethod, *proto.Error) {
	paymentMethod, err := usecase.Repository.CreateOne(paymentMethod)
	if err != nil {
		return nil, &proto.Error{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	return paymentMethod, nil
}

func (usecase *Usecase) UpdateOne(paymentMethod *proto.PaymentMethod) (*proto.PaymentMethod, *proto.Error) {
	paymentMethod, err := usecase.Repository.UpdateOne(paymentMethod)
	if err != nil {
		return nil, &proto.Error{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	return paymentMethod, nil
}