package ezpay

type Config struct {
	MerchantID string `mapstructure:"merchantId"`
	URL        string `mapstructure:"url"`
	APIVersion string `mapstructure:"apiVersion"`
	Key        string `mapstructure:"key"`
	IV         string `mapstructure:"iv"`
}
