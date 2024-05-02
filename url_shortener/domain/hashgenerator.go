package domain

import "math/big"

// Interface that can generate a unique hash
type HashGenerator interface {
	GenerateHash() string
}

// Class that can generate and transform a snowflake id to hash by base 62
type SnowFlakeHashGenerator struct {
	IDGenerator SnowFlake
}

// Generate a hash from a snowflake id
func (s *SnowFlakeHashGenerator) GenerateHash() string {
	id := s.IDGenerator.GenerateID()
	var base62Int big.Int
	base62Int.SetInt64(id)
	return base62Int.Text(62)
}
