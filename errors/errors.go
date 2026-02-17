package ers

import "errors"

var (
	ErrIndexerAlreadyRunning = errors.New("indexer: already running")
	ErrIndexerNotRunning     = errors.New("indexer: not running")
	ErrIndexerUnknownNetwork = errors.New("indexer: unknown network")
	ErrIndexerNoPeers        = errors.New("indexer: no peers available")
)
var (
	ErrDbClosed               = errors.New("db: database is closed")
	ErrDbColumnFamilyNotFound = errors.New("db: column family not found")
	ErrDbKeyNotFound          = errors.New("db: key not found")
	ErrDbNilKey               = errors.New("db: key must not be nil")
	ErrDbBatchClosed          = errors.New("db: batch is closed")
)
