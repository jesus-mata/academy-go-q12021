package repository

import "github.com/jesus-mata/academy-go-q12021/domain"

type NewsArticleRepository interface {
	FindAll() ([]*domain.NewsArticle, error)
	FindByID(id string) (*domain.NewsArticle, error)
	FetchCurrent() (string, error)
}