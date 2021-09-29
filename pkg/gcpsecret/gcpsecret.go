package gcpsecret

import (
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/sirupsen/logrus"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

func Get(projectID, secretID, secretVersion string) ([]byte, error) {

	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		logrus.Fatalf("failed to setup client: %v", err)
	}
	defer client.Close()

	// Build the request.
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/%s", projectID, secretID, secretVersion),
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		logrus.Fatalf("failed to access secret version: %v", err)
	}

	return result.Payload.Data, nil

}
