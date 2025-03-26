package providers

import (
	"github.com/mauriciofsnts/bot/internal/config"
	"github.com/mauriciofsnts/bot/internal/database/service"
	"github.com/mauriciofsnts/bot/internal/providers/cache"
	"github.com/mauriciofsnts/bot/internal/providers/cron"
	footballdata "github.com/mauriciofsnts/bot/internal/providers/football_data"
	"github.com/mauriciofsnts/bot/internal/providers/news"
	"github.com/mauriciofsnts/bot/internal/providers/shorten"
	"gorm.io/gorm"
)

var (
	Shorten  shorten.URLShortener
	News     news.NewsProvider
	DB       *gorm.DB
	Services service.IService
	Football footballdata.FootballDataProvider
	Cron     cron.Cron
	Cache    cache.Cache
)

func New(cfg config.Config, db *gorm.DB) {
	DB = db
	Shorten = shorten.New(cfg)
	News = news.New(cfg)
	Services = service.New(db)
	Football = footballdata.New(cfg)
	Cron = cron.New()
	Cache = cache.New(cfg.Valkey.Address)
}
