package appctx

import (
	"go-api/component/tokenpvd"
	"go-api/component/uploadpvd"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetUploadProvider() uploadpvd.UploadProvider
	GetTokenProvider() tokenpvd.TokenProvider
}

type appCtx struct {
	db             *gorm.DB
	tokenProvider  tokenpvd.TokenProvider
	uploadProvider uploadpvd.UploadProvider
}

func NewAppContext(db *gorm.DB, tokenProvider tokenpvd.TokenProvider, uploadProvider uploadpvd.UploadProvider) *appCtx {
	return &appCtx{
		db:             db,
		tokenProvider:  tokenProvider,
		uploadProvider: uploadProvider,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetTokenProvider() tokenpvd.TokenProvider {
	return ctx.tokenProvider
}

func (ctx *appCtx) GetUploadProvider() uploadpvd.UploadProvider {
	return ctx.uploadProvider
}
