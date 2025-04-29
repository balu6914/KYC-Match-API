package database

type Database interface {
	Connect() error
	Close() error
}

type HarperDB struct {
}

func (h *HarperDB) Connect() error {

	return nil
}

func (h *HarperDB) Close() error {

	return nil
}
