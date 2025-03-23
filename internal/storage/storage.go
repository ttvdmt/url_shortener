package storage

type Storager interface {
	Save(url string) (string, error)
	Get(code string) (string, error)
}
