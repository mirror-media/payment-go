package payment

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	formatter "github.com/bcgodev/logrus-formatter-gke"
	"github.com/mirror-media/payment-go/pkg/gcpsecret"
	"github.com/mirror-media/payment-go/pkg/invoice"
	"github.com/mirror-media/payment-go/pkg/invoice/ezpay"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	logrus.SetFormatter(&formatter.GKELogFormatter{})
	logrus.SetReportCaller(true)
}

func loadConfig() (ezpay.Config, error) {
	projectID := os.Getenv("MM_PROJECT_ID")
	secretID := os.Getenv("MM_CONFIG_SECRET")
	secretVersion := os.Getenv("MM_CONFIG_SECRET_VERSION")

	configValue, err := gcpsecret.Get(projectID, secretID, secretVersion)
	if err != nil {
		return ezpay.Config{}, err
	}

	config := ezpay.Config{}
	b := bytes.NewBuffer(configValue)
	viper.SetConfigType("env")
	viper.ReadConfig(b)
	viper.Unmarshal(&config)

	return config, err
}

func CreateInvoice(w http.ResponseWriter, r *http.Request) {

	config, err := loadConfig()
	if err != nil {
		logrus.Errorf("load config encounter error:%+v", err)
		http.Error(w, "", http.StatusInternalServerError)
	}

	var payload map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		err = errors.Wrap(err, "decoding payload error")
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	provider, _ := invoice.NewEzPayInvoiceProvider(config, payload)

	resp, err := provider.Create()
	if err != nil {
		err = errors.Wrap(err, "invoice creation error")
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		logrus.Errorf("json encode resp(%+v) error:%+v", resp, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
