package services

import (
	"goapi/internal/modules/tokens/dto"
	"goapi/pkg/permissions"
	"goapi/pkg/errors"
	"goapi/pkg/jwt"
	"os"

	"github.com/google/uuid"
)

type TokenServices interface {
	CreateToken() (*dto.TokenResponse, error)
}

type TokenServicesImpl struct{}

func NewTokenService() (*TokenServicesImpl, error) {
	return &TokenServicesImpl{}, nil
}

func (serv TokenServicesImpl) CreateToken(body *dto.TokenRequest) (*dto.TokenResponse, error) {
	id := uuid.New().String()

	SECRET := os.Getenv("JWT_SECRET_KEY")

	if body.Secret != SECRET {

		credentials, _ := permissions.GetCredentialsByRole(permissions.ReadOnlyRole)

		token, err := jwt.GenerateNewTokens(id, permissions.ReadOnlyRole, credentials)

		if err != nil {
			return nil, errors.InternalServerError(errors.Message{"error": true, "msg": err.Error()})
		}

		return &dto.TokenResponse{
			Token:       token.Access,
			Role:        permissions.ReadOnlyRole,
			Credentials: credentials,
		}, nil
	}

	credentials, err := permissions.GetCredentialsByRole(body.Role)

	if err != nil {
		return nil, errors.NotFound(errors.Message{"error": true, "msg": err.Error()})
	}

	token, err := jwt.GenerateNewTokens(id, body.Role, credentials)

	if err != nil {
		return nil, errors.InternalServerError(errors.Message{"error": true, "msg": err.Error()})
	}

	return &dto.TokenResponse{
		Token:       token.Access,
		Role:        body.Role,
		Credentials: credentials,
	}, nil
}
