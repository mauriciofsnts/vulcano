package providers

import (
	"github.com/mauriciofsnts/bot/internal/providers/news"
	"github.com/mauriciofsnts/bot/internal/providers/shorten"
	"gorm.io/gorm"
)

type Provider struct {
	Shorten shorten.URLShortener
	News    news.NewsProvider
	DB      *gorm.DB
}

var Providers *Provider
