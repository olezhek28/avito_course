package config

// DSN ...
func (s *DB) DSN() string {
	return s.dsn
}

// MaxOpenConnections ...
func (s *DB) MaxOpenConnections() int32 {
	return s.maxOpenConnections
}
