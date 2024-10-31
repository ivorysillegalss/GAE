package bloom

import (
	"github.com/bits-and-blooms/bloom/v3"
)

type Client interface {
	Add(v string)
	Check(v string) bool
}

type bloomClient struct {
	filter *bloom.BloomFilter
}

func (b *bloomClient) Add(v string) {
	b.filter.Add([]byte(v))
}

func (b *bloomClient) Check(v string) bool {
	return b.filter.Test([]byte(v))
}

func NewBloomClient() Client {
	// 初始化布隆过滤器
	n := uint(1000)           // 预计要插入的元素数量
	falsePositiveRate := 0.01 // 允许的误判率
	f := bloom.NewWithEstimates(n, falsePositiveRate)
	return &bloomClient{filter: f}
}
