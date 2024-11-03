package bloom

import (
	"bytes"
	"encoding/gob"
	"gae-backend-crawl/infrastructure/log"
	"github.com/bits-and-blooms/bloom/v3"
	"os"
	"sync"
	"time"
)

type Client interface {
	Add(v string)
	Check(v string) bool
	SaveToFile(filename string) error
	LoadFromFile(filename string) error
	StartAutoSave(filename string, interval time.Duration)
	StopAutoSave()
}

type bloomClient struct {
	filter    *bloom.BloomFilter
	autoSave  bool
	saveMutex sync.Mutex
	stopChan  chan bool
}

func (b *bloomClient) Add(v string) {
	b.filter.Add([]byte(v))
}

func (b *bloomClient) Check(v string) bool {
	return b.filter.Test([]byte(v))
}

// SaveToFile 序列化并存储布隆过滤器到文件
func (b *bloomClient) SaveToFile(filename string) error {
	b.saveMutex.Lock()
	defer b.saveMutex.Unlock()

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(b.filter); err != nil {
		return err
	}
	return os.WriteFile(filename, buffer.Bytes(), 0644)
}

// LoadFromFile 从文件加载布隆过滤器
func (b *bloomClient) LoadFromFile(filename string) error {

	// 检查文件是否存在
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// 如果文件不存在，创建一个新的空布隆过滤器并保存到文件
		if err := b.SaveToFile(filename); err != nil {
			return err
		}
		return nil // 文件不存在时，返回nil表示正常创建
	}

	// 如果文件存在，读取文件数据
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// 解码数据到布隆过滤器
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	return decoder.Decode(&b.filter)
}

// StartAutoSave 开启定时自动保存
func (b *bloomClient) StartAutoSave(filename string, interval time.Duration) {
	b.stopChan = make(chan bool)
	b.autoSave = true
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				b.SaveToFile(filename)
				log.GetTextLogger().Info("auto save successfully")
			case <-b.stopChan:
				return
			}
		}
	}()
}

// StopAutoSave 停止定时自动保存
func (b *bloomClient) StopAutoSave() {
	if b.autoSave {
		b.stopChan <- true
		b.autoSave = false
		close(b.stopChan)
	}
}

// NewBloomClient 初始化布隆过滤器
func NewBloomClient(n uint, falsePositiveRate float64) Client {
	f := bloom.NewWithEstimates(n, falsePositiveRate)
	return &bloomClient{filter: f}
}
