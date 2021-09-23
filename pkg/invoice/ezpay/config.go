package ezpay

type Config struct {
	MerchantID string `json:"merchant_id"`
	URL        string `json:"url"`
	APIVersion string `json:"api_version"`
	Key        string `json:"key"`
	IV         string `json:"iv"`
}
