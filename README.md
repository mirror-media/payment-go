# PAYMENT-GO

`PAYMENT-GO` provides packages for payments which should be easy to used. Upon the packages, it also contain payment implementation leveraging GCP cloud functions.

## Cloud Functions

1. create invoice

### Invoice Creation API

Please refer to `EZPAY_INVOICE_1_2_1.pdf` in `doc` folder. The payload response should be a JSON object, which has key as the property name and value as the property value defined in the doc.

You may use any api version, but please bear in mind that the default version and implementation is `1.4`.

## Deployment of invoice function

**Mirror Weekly**

[Dev env](https://console.cloud.google.com/functions/details/asia-east1/weekly-subscribtion-ezpay-invoice-issuer?authuser=2&project=mirrormedia-1470651750304)
[Prod env](https://console.cloud.google.com/functions/details/asia-east1/weekly-subscribtion-ezpay-invoice-issuer-prod?authuser=2&project=mirrormedia-1470651750304)

Conveniently, cloud function features testing in the console. You may test with API payload to see the result in `Dev env`.

### Environmental Variables

The names are self-explanatory, and the config should be defined as in `ENVIRONMENT variables` style for the implementation. For now it supports [EZPAY configs](https://github.com/mirror-media/payment-go/blob/1db4f4e099f617a62daad373987b244b4fd105ed/pkg/invoice/ezpay/config.go) only.

1. `MM_PROJECT_ID`
2. `MM_CONFIG_SECRET`
3. `MM_CONFIG_SECRET_VERSION`

## CI/CD

It not configured! Please refer to the [GCP doc](https://cloud.google.com/functions/docs/deploying) for deployment.
