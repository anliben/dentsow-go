package configs

import (
	"os"

	"github.com/mobilemindtec/go-payments/api"
)

type AssasTokenClientHandler struct {
	AsaasToken string        `json:"asaas_token"`
	AsaasMode  api.AsaasMode `json:"asaas_token"`
}

func GetAsaasToken() AssasTokenClientHandler {
	token := os.Getenv("ASAAS_TOKEN")

	asaasHandler := AssasTokenClientHandler{}

	if token == "" {
		asaasHandler.AsaasToken = "$aact_YTU5YTE0M2M2N2I4MTliNzk0YTI5N2U5MzdjNWZmNDQ6OjAwMDAwMDAwMDAwMDAwNTIxMTY6OiRhYWNoX2QwMzI3MDA4LTlmZDktNDM4OS1hODJiLTI3MTczN2U5YTQxNQ=="
		asaasHandler.AsaasMode = api.AsaasModeTest
	} else {
		asaasHandler.AsaasToken = token
		asaasHandler.AsaasMode = api.AsaasModeProd
	}

	return asaasHandler
}
