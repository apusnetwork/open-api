package dal

import (
	"context"

	"github.com/apusnetwork/open-api/biz/model"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
)

func GetAuthzByAccessKey(ctx context.Context, conn *gorm.DB, accessKey string) (*model.Authz, error) {
	s := &model.Authz{}

	err := conn.First(s, "access_key = ?", accessKey).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		hlog.CtxErrorf(ctx, "conn `Find` failed, detail: %s", err)
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return s, nil
}

func CreateOrUpdate(ctx context.Context, conn *gorm.DB, accessKey string, secretKey string, role string) (*model.Authz, error) {
	s := &model.Authz{
		AccessKey: accessKey,
		SecretKey: secretKey,
		Role:      role,
	}

	err := conn.First(s, "access_key = ?", accessKey).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		hlog.CtxErrorf(ctx, "conn `Find` failed, detail: %s", err)
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		err = conn.Model(&model.Authz{}).Create(s).Error
		if err != nil {
			hlog.CtxErrorf(ctx, "conn `Create` query error: %s", err)
			return nil, err
		}
		return s, nil
	}

	s.SecretKey = secretKey
	s.Role = role

	err = conn.Where("id = ?", s.ID).Save(s).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "conn `Updates` error: %s", err)
		return nil, err
	}
	return s, nil
}
