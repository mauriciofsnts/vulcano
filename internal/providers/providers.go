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
	Cache    cache.Valkey
)

func New(cfg config.Config, db *gorm.DB) {
	cache := cache.New(cfg.Valkey.Address)
	cron := cron.New()
	shorten := shorten.New(cfg)

	DB = db
	Cache = cache
	Cron = cron
	Shorten = shorten
	News = news.New(cfg, cache, shorten)
	Services = service.New(db)
	Football = footballdata.New(cfg, cron, cache)
}
