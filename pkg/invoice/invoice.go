package invoice

import "github.com/mirror-media/payment-go/pkg/invoice/ezpay"

// Provider is the interface each invoice service has to implement
type Provider interface {
	Create() (resp []byte, err error)
	Validate() error
}

// NewEzPayInvoiceProvider returns a provider from ezPay
func NewEzPayInvoiceProvider(config ezpay.Config, data map[string]interface{}) (p Provider, err error) {
	return &ezpay.InvoiceClient{Payload: data, Config: config}, nil
}
