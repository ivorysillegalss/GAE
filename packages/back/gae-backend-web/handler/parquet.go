package handler

import (
	"gae-backend-web/constant/mq"
	"gae-backend-web/domain"
	"gae-backend-web/infrastructure/log"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
	"sync"
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
	ReadData(filePath string, businessId int) *domain.Talent
	GetIndex(businessId int) int
}

type ParquetHandler struct {
}

func (p *ParquetHandler) ReadData(filePath string, businessId int) *domain.Talent {
	var contributors []*domain.Contributor
	var repos []*domain.Repo

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
			var talentBatch []*domain.Contributor
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
			var talentBatch []*domain.Repo
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
