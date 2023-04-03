package cache

import "time"

type GoogleSheetsCache interface {
	Set(key string, value []string)
	SetPollingKey(key string)
	Get(key string) []string
	Ttl(key string) time.Duration
}
