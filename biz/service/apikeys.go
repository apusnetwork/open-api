package service

import (
	"context"
	"errors"

	"github.com/apusnetwork/open-api/biz/dal"
	"github.com/apusnetwork/open-api/biz/model"
)

func GenerateSecretKey() {

}

func CheckSecretKey(ctx context.Context, accessKey string, message string, signature string) (*model.Authz, error) {
	authz, err := dal.GetAuthzByAccessKey(ctx, dal.GetDBConn(ctx), accessKey)
	if err != nil {
		return nil, err
	}
	if authz == nil {
		return nil, errors.New("authentication failed: access key doesn't exist")
	}
	// TODO 更改为签名校验
	if authz.SecretKey != signature {
		return nil, errors.New("authentication failed: signature is illegal")
	}

	return authz, nil
}
