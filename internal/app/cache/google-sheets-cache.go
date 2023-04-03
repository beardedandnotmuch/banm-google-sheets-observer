package cache

type GoogleSheetsCache interface {
	Set(key string, value []string)
	SetPollingKey(key string)
	Get(key string) []string
}
