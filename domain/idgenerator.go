package domain

import "github.com/bwmarrin/snowflake"

// Interface that can generate a unique id
type IDGenerator interface {
	GenerateID() int64
}

// Class that can generate a unique id using twitter snowflake algorithm
type SnowFlake struct {
	Node *snowflake.Node
}

// Generate a unique id using twitter snowflake algorithm
func (s *SnowFlake) GenerateID() int64 {
	id := s.Node.Generate()
	return id.Int64()
}
