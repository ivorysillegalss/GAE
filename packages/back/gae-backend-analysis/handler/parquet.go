package handler

import (
	"gae-backend-analysis/constant/common"
	"gae-backend-analysis/constant/dao"
	"gae-backend-analysis/constant/mq"
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
	repoMutex sync.Mutex
	repoIndex int64
	userMutex sync.Mutex
	userIndex int64
)

func NewIntermediateHandler() IntermediateHandler {
	return &ParquetHandler{}
}

// IntermediateHandler 中间存储的实现方式是Parquet
type IntermediateHandler interface {
	WriteData(talent domain.Talent, businessId int)
	//ReadContributorData(filePath string) []domain.Talent
	//ReadRepoData(filePath string) []domain.Talent
	ReadData(filePath string, businessId int) *domain.Talent
	GetIndex(businessId int) int
}

type ParquetHandler struct {
}

func (p *ParquetHandler) WriteData(any domain.Talent, businessId int) {
	var filePath string

	switch businessId {
	case mq.UnCleansingRepoId:
		repoMutex.Lock()
		defer repoMutex.Unlock()

		index := atomic.AddInt64(&repoIndex, 1)
		filePath = dao.RepoPrefix + common.UnderScore + dao.OutputFilepath + common.UnderScore + strconv.FormatInt(index, 10) + dao.ParquetSuffix

	case mq.UnCleansingUserId:
		userMutex.Lock()
		defer userMutex.Unlock()
		index := atomic.AddInt64(&userIndex, 1)
		filePath = dao.ContributorPrefix + common.UnderScore + dao.OutputFilepath + common.UnderScore + strconv.FormatInt(index, 10) + dao.ParquetSuffix
	}

	// 创建文件
	f, err2 := local.NewLocalFileWriter(filePath)
	if err2 != nil {
		log.GetTextLogger().Error("Error creating file %s: %v", filePath, err2)
		return
	}
	defer f.Close()

	switch businessId {
	case mq.UnCleansingUserId:
		// 创建 Parquet Writer
		pw, err := writer.NewParquetWriter(f, any.Contributors, dao.ParquetDefaultGoroutine)
		if err != nil {
			log.GetTextLogger().Warn("Error creating Parquet writer for %s: %v", filePath, err)
			return
		}
		defer pw.WriteStop()
		defer pw.Flush(true)
		pw.CompressionType = parquet.CompressionCodec_SNAPPY

		// 写入数据
		for _, record := range *any.Contributors {
			if err := pw.Write(record); err != nil {
				log.GetTextLogger().Error("Write error for file %s: %v", filePath, err)
			}
		}

	case mq.UnCleansingRepoId:
		// 创建 Parquet Writer
		pw, err := writer.NewParquetWriter(f, any.Repos, dao.ParquetDefaultGoroutine)
		if err != nil {
			log.GetTextLogger().Warn("Error creating Parquet writer for %s: %v", filePath, err)
			return
		}
		defer pw.WriteStop()
		defer pw.Flush(true)
		pw.CompressionType = parquet.CompressionCodec_SNAPPY

		// 写入数据
		for _, record := range *any.Repos {
			if err := pw.Write(record); err != nil {
				log.GetTextLogger().Error("Write error for file %s: %v", filePath, err)
			}
		}
	}

	log.GetTextLogger().Info("Finished writing to %s", filePath)
}

func (p *ParquetHandler) ReadContributorData(filePath string) []domain.Talent {
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

// TODO 方法路径由调用者决定，示例有
//
//	filePath := dao.OutputFilepath + common.UnderScore + strconv.Itoa(getParquetIndex()) + dao.ParquetSuffix
func (p *ParquetHandler) ReadData(filePath string, businessId int) *domain.Talent {
	var contributors []domain.Contributor
	var repos []domain.Repo

	switch businessId {
	case mq.UnCleansingUserId:
		fileReader, _ := local.NewLocalFileReader(filePath)
		defer fileReader.Close()
		parquetReader, err := reader.NewParquetReader(fileReader, new(domain.Contributor), 4)
		if err != nil {
			log.GetTextLogger().Error("Error creating Parquet reader for %s: %v", filePath, err)
		}
		defer parquetReader.ReadStop()

		num := int(parquetReader.GetNumRows())
		for j := 0; j < num; j++ {
			var talentBatch []domain.Contributor
			if err := parquetReader.Read(&talentBatch); err != nil {
				log.GetTextLogger().Error("Read error for file %s: %v", filePath, err)
				break
			}
			contributors = append(contributors, talentBatch...)
		}

	case mq.UnCleansingRepoId:
		fileReader, _ := local.NewLocalFileReader(filePath)
		defer fileReader.Close()
		parquetReader, err := reader.NewParquetReader(fileReader, new(domain.Repo), 4)
		if err != nil {
			log.GetTextLogger().Error("Error creating Parquet reader for %s: %v", filePath, err)
		}
		defer parquetReader.ReadStop()

		num := int(parquetReader.GetNumRows())
		for j := 0; j < num; j++ {
			var talentBatch []domain.Repo
			if err := parquetReader.Read(&talentBatch); err != nil {
				log.GetTextLogger().Error("Read error for file %s: %v", filePath, err)
				break
			}
			repos = append(repos, talentBatch...)
		}
	}

	return &domain.Talent{Contributors: &contributors, Repos: &repos}
}

func (p *ParquetHandler) GetIndex(businessId int) int {
	switch businessId {
	case mq.UnCleansingUserId:
		return int(userIndex)
	case mq.UnCleansingRepoId:
		return int(repoIndex)
	default:
		panic("error")
	}
}
