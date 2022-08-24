package __0

import (
	"bytes"
	"context"
	"net/http"
	"toscactl/entity"
)

const (
	APICreateExecution = "execution"
)

type Provider struct {
	config *entity.ApplicationConfig
}

func (t *Provider) triggerExecution(ctx context.Context) (*http.Request, error) {
	serverURL := "http://<Tosca Server Gateway IP address or hostname>:<Gateway port>/automationobjectservice"
	body := []byte("{\n  \"projectName\": \"string\",\n  \"executionEnvironment\": \"string\",\n  \"events\": [\n    {\n      \"eventId\": \"string\",\n      \"parameters\": {\n        \"additionalProp1\": \"string\",\n        \"additionalProp2\": \"string\",\n        \"additionalProp3\": \"string\"\n      },\n      \"characteristics\": {\n        \"additionalProp1\": \"string\",\n        \"additionalProp2\": \"string\",\n        \"additionalProp3\": \"string\"\n      }\n    }\n  ],\n  \"importResult\": true,\n  \"creator\": \"string\"\n}")
	req, err := http.NewRequestWithContext(ctx, "POST", serverURL, bytes.NewBuffer(body))
	return req, err
}
