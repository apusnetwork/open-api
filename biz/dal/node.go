package dal

import (
	"context"
	"time"

	"github.com/apusnetwork/open-api/biz/model"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
)

func GetNodes(ctx context.Context, conn *gorm.DB, onlyActive bool) ([]*model.Node, error) {
	sl := make([]*model.Node, 0)

	if onlyActive {
		m, _ := time.ParseDuration("-60s")
		conn = conn.Where("updated_at > ?", time.Now().Add(m))
	}

	err := conn.Find(&sl).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "conn `Find` failed, detail: %s", err)
		return nil, err
	}
	return sl, nil
}
