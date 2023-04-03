package cache

type GoogleSheetsCache interface {
	Set(key string, value []string)
	Get(key string) []string
}
