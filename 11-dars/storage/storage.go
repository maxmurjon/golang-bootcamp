package storage

type StorageI interface{
	Get(key string) ([]byte, error)
}