package handler

import (
	"gae-backend-analysis/constant/common"
	"gae-backend-analysis/constant/dao"
	"gae-backend-analysis/domain"
	"gae-backend-analysis/infrastructure/log"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/reader"
	"github.com/xitongsys/parquet-go/writer"
	"strconv"
	"sync"
	"sync/atomic"
)

var (
	fileMutex sync.Mutex
	fileIndex int64
)

func NewIntermediateHandler() IntermediateHandler {
	return &ParquetHandler{}
}

// IntermediateHandler 中间存储的实现方式是Parquet
type IntermediateHandler interface {
	// WriteData 写入函数
	WriteData(talent []domain.Talent)
	ReadData(filePath string) []domain.Talent
	GetIndex() int
}

type ParquetHandler struct {
}

func (p *ParquetHandler) WriteData(any []domain.Talent) {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	index := atomic.AddInt64(&fileIndex, 1)
	filePath := dao.OutputFilepath + common.UnderScore + strconv.FormatInt(index, 10) + dao.ParquetSuffix

	// 创建文件
	f, err2 := local.NewLocalFileWriter(filePath)
	if err2 != nil {
		log.GetTextLogger().Error("Error creating file %s: %v", filePath, err2)
		return
	}
	defer f.Close()

	// 创建 Parquet Writer
	pw, err := writer.NewParquetWriter(f, any, dao.ParquetDefaultGoroutine)
	if err != nil {
		log.GetTextLogger().Warn("Error creating Parquet writer for %s: %v", filePath, err)
		return
	}
	defer pw.WriteStop()
	defer pw.Flush(true)

	pw.CompressionType = parquet.CompressionCodec_SNAPPY

	// 写入数据
	for _, record := range any {
		if err := pw.Write(record); err != nil {
			log.GetTextLogger().Error("Write error for file %s: %v", filePath, err)
		}
	}

	log.GetTextLogger().Info("Finished writing to %s", filePath)
}

func (p *ParquetHandler) ReadData(filePath string) []domain.Talent {
	//filePath := dao.OutputFilepath + common.UnderScore + strconv.Itoa(getParquetIndex()) + dao.ParquetSuffix
	fileReader, _ := local.NewLocalFileReader(filePath)
	defer fileReader.Close()
	parquetReader, err := reader.NewParquetReader(fileReader, new(domain.Talent), 4)
	if err != nil {
		log.GetTextLogger().Error("Error creating Parquet reader for %s: %v", filePath, err)
	}
	defer parquetReader.ReadStop()

	var talents []domain.Talent
	num := int(parquetReader.GetNumRows())
	for j := 0; j < num; j++ {
		var talentBatch []domain.Talent
		if err := parquetReader.Read(&talentBatch); err != nil {
			log.GetTextLogger().Error("Read error for file %s: %v", filePath, err)
			break
		}
		talents = append(talents, talentBatch...)
	}
	return talents
}

func (p *ParquetHandler) GetIndex() int {
	return int(fileIndex)
}
