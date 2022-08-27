package model

type SourceType int64

const (
	DbSource SourceType = iota
	MemorySource
)
