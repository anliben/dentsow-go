package configs

import (
	"os"

	"github.com/mobilemindtec/go-payments/api"
)

type AssasTokenClientHandler struct {
	AsaasToken string        `json:"asaas_token"`
	AsaasMode  api.AsaasMode `json:"asaas_mode"`
}

func GetAsaasToken() AssasTokenClientHandler {
	token := os.Getenv("ASAAS_TOKEN")

	asaasHandler := AssasTokenClientHandler{}

	if token == "" {
		asaasHandler.AsaasToken = "aact_YTU5YTE0M2M2N2I4MTliNzk0YTI5N2U5MzdjNWZmNDQ6OjAwMDAwMDAwMDAwMDAyOTg3Njc6OiRhYWNoXzQxZWVkN2E3LWRkMDgtNGY3Ni1iZGFlLTczYjQzZjVkMmQ2ZA=="
		asaasHandler.AsaasMode = api.AsaasModeProd
	} else {
		asaasHandler.AsaasToken = token
		asaasHandler.AsaasMode = api.AsaasModeProd
	}

	return asaasHandler
}
