package idgen

import (
	"log"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
	once sync.Once
)

func NewSnowflakeNode() {
	once.Do(func() {
		nodeId := 1
		snowflake.Epoch = time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli()

		n, err := snowflake.NewNode(int64(nodeId))
		if err != nil {
			log.Fatalf("Failed to initialize snowflake node: %v", err)
		}
		node = n
	})
}

func GenerateID() int64 {
	if node == nil {
		NewSnowflakeNode()
	}
	return node.Generate().Int64()
}
