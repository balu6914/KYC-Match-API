package database

type Database interface {
	Connect() error
	Close() error
}

type HarperDB struct {
	// Placeholder for HarperDB client connection
}

func (h *HarperDB) Connect() error {
	// Implement HarperDB connection logic here
	return nil
}

func (h *HarperDB) Close() error {
	// Implement HarperDB close logic here
	return nil
}
