package providers

import (
	"github.com/mauriciofsnts/bot/internal/config"
	"github.com/mauriciofsnts/bot/internal/providers/news"
	"github.com/mauriciofsnts/bot/internal/providers/shorten"
	"github.com/mauriciofsnts/bot/internal/service"
	"gorm.io/gorm"
)

var (
	Shorten  shorten.URLShortener
	News     news.NewsProvider
	DB       *gorm.DB
	Services service.IService
)

func New(db *gorm.DB, cfg config.Config) {
	DB = db
	Shorten = shorten.New(cfg)
	News = news.New(cfg)
	Services = service.New(db)
}
