package util

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// var (
// 	mockChunks = [][]byte{
// 		{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32},
// 		{5, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32},
// 		{2, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32},
// 		{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32},
// 		{100, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32},
// 		{66, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32},
// 		{130, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32},
// 		{67, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32},
// 		{54, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32},
// 		{24, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32},
// 		{14, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32},
// 	}
// )
var generateRandom32Bytes = func(n int) [][]byte {
	result := make([][]byte, n)
	for i := 0; i < n; i++ {
		rand.Seed(rand.Int63())
		token := make([]byte, 32)
		rand.Read(token)
		result[i] = token
	}
	return result
}
var generateRandomNodeIDs = func(n int) []int64 {
	result := make([]int64, n)
	for i := 0; i < n; i++ {
		result[i] = rand.Int63()
	}
	return result
}

func TestChunk_ShardChunk(t *testing.T) {
	t.Run("6 shardBit-n", func(t *testing.T) {
		const n = 1000000
		startPrepareData := time.Now()
		fmt.Printf("preparing %d random data\n", n)
		mockChunks := generateRandom32Bytes(n)
		fmt.Printf("data prepared in: %v ms\n", time.Since(startPrepareData).Milliseconds())
		fmt.Printf("start sharding data\n")
		startSharding := time.Now()
		chunk := &Chunk{
			chunkSize: sha256.Size,
		}
		var chunks []byte
		for _, mockChunk := range mockChunks {
			chunks = append(chunks, mockChunk...)
		}
		result := chunk.ShardChunk(chunks, 6)
		fmt.Printf("finish sharding in : %v ms\n", time.Since(startSharding).Milliseconds())
		for u, i := range result {
			fmt.Printf("shardN: %d\tcontent: %d\n", u, len(i))
		}
	})
}

func TestChunk_AssignShard(t *testing.T) {
	t.Run("assignShard - 1000 nodes", func(t *testing.T) {
		const n = 1000000
		startPrepareData := time.Now()
		fmt.Printf("preparing %d random data\n", n)
		mockChunks := generateRandom32Bytes(n)
		fmt.Printf("data prepared in: %v ms\n", time.Since(startPrepareData).Milliseconds())
		fmt.Printf("start sharding data\n")
		chunk := &Chunk{
			chunkSize: sha256.Size,
		}
		var chunks []byte
		for _, mockChunk := range mockChunks {
			chunks = append(chunks, mockChunk...)
		}
		startSharding := time.Now()
		shards := chunk.ShardChunk(chunks, 6)
		fmt.Printf("time sharding chunks: %v ms\n", time.Since(startSharding).Milliseconds())
		nodeIDs := generateRandomNodeIDs(1000)
		startAssignChunk := time.Now()
		shard, err := chunk.AssignShard(shards, nodeIDs)
		if err != nil {
			t.Errorf("error-assigning-shard: %v", err)
		}
		fmt.Printf("time assigning shard: %v ms\n", time.Since(startAssignChunk).Milliseconds())
		for i, s := range shard {
			fmt.Printf("nodeID: %d\tnumShard: %d\n", i, len(s))
		}
	})
}
