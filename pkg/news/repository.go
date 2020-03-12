package news

type Repository interface {
	Find(key string) (News, error)
	FindByQuery(query string) ([]News, error)
	Store(news News) error
}
