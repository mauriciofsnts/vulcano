package providers

import (
	"github.com/mauriciofsnts/bot/internal/config"
	"github.com/mauriciofsnts/bot/internal/database/service"
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
)

func New(db *gorm.DB, cfg config.Config) {
	DB = db
	Shorten = shorten.New(cfg)
	News = news.New(cfg)
	Services = service.New(db)
	Football = footballdata.New(cfg)
}
