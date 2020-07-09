package v1

import (
	"net/http"

	paymentMethodPB "github.com/ta04/payment-method-service/model/proto"
	"github.com/ta04/payment-method-service/repository"
)

//Usecase is the struct of payment method usecase
type Usecase struct {
	Repository repository.Repository
}

// NewUsecase will create a new order usecase
func NewUsecase(repository repository.Repository) *Usecase {
	return &Usecase{
		Repository: repository,
	}
}

func (usecase *Usecase) Index(request *paymentMethodPB.IndexPaymentMethodsRequest) ([]*paymentMethodPB.PaymentMethod, error) {
	if request == nil {
		return nil, &paymentMethodPB.Error{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}
	}

	if request.Status == "" {
		request.Status = "waiting for payment"
	}

	var paymentMethod []*paymentMethodPB.PaymentMethod
	var err error
	if request.UserId != "" {
		paymentMethod, err = usecase.Repository.GetAllByUserID(request)
		if err != nil {
			return nil, &paymentMethodPB.Error{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			}
		}
	} else {
		paymentMethod, err = usecase.Repository.GetAll(request)
		if err != nil {
			return nil, &paymentMethodPB.Error{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			}
		}
	}

	return paymentMethod, nil
}
func (usecase *Usecase) Show(request *paymentMethodPB.PaymentMethod) (*paymentMethodPB.PaymentMethod, error) {
	if request == nil || (request.Id == 0) {
		return nil, &paymentMethodPB.Error{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}
	}

	paymentMethod, err := usecase.Repository.GetOne(request)
	if err != nil {
		return nil, &paymentMethodPB.Error{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	return paymentMethod, nil
}
func (usecase *Usecase) Store(paymentMethod *paymentMethodPB.PaymentMethod) (*paymentMethodPB.PaymentMethod, error) {
	paymentMethod, err := usecase.Repository.Store(paymentMethod)
	if err != nil {
		return nil, &paymentMethodPB.Error{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	return paymentMethod, nil
}
func (usecase *Usecase) Update(paymentMethod *paymentMethodPB.PaymentMethod) (*paymentMethodPB.PaymentMethod, error) {
	paymentMethod, err := usecase.Repository.Update(paymentMethod)
	if err != nil {
		return nil, &paymentMethodPB.Error{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	return paymentMethod, nil
}
func (usecase *Usecase) Destroy(paymentMethod *paymentMethodPB.PaymentMethod) (*paymentMethodPB.PaymentMethod, error) {
	paymentMethod, err := usecase.Repository.Destroy(paymentMethod)
	if err != nil {
		return nil, &paymentMethodPB.Error{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	return paymentMethod, nil
}
