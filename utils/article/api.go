package article

type Article interface {
	Get(url string) map[string]interface{}
}
