package imp

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type DoPaymentImp struct {
}

type Health struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Time    string `json:"time"`
}

type Authorisation struct {
	Authorised bool   `json:"authorised"`
	Message    string `json:"message"`
}

type Service struct {
	declineOverAmount float32
}

func (imp *DoPaymentImp) EchoHello(name string, greeting *string) (int32, error) {
	*greeting = "hello " + name
	return 0, nil
}

func (imp *DoPaymentImp) Health() (string, error) {
	var health []Health
	app := Health{"payment", "OK", time.Now().String()}
	health = append(health, app)
	retHealth, _ := json.Marshal(health)
	return string(retHealth), nil
}
func (imp *DoPaymentImp) Authorise(amount float32) (string, error) {
	var ErrInvalidPaymentAmount = errors.New("Invalid payment amount")
	if amount == 0 {
		return "", ErrInvalidPaymentAmount
	}
	if amount < 0 {
		return "", ErrInvalidPaymentAmount
	}
	authorised := false
	message := "Payment declined"
	service := Service{99.99}
	fmt.Println(amount)
	if amount <= service.declineOverAmount {
		authorised = true
		message = "Payment authorised"
	} else {
		message = fmt.Sprintf("Payment declined: amount exceeds %.2f", service.declineOverAmount)
	}
	authrisation := Authorisation{authorised, message}
	retAuthrisation, _ := json.Marshal(authrisation)
	return string(retAuthrisation), nil
}
