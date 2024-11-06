package handler

import (
	"gae-backend-storage/bootstrap"
	"gae-backend-storage/constant/sys"
)

type ParquetStorageEngine interface {
	Setup()
}

type storageEngine struct {
	es *bootstrap.SearchEngine
}

func (s *storageEngine) Setup() {
	handler := NewIntermediateHandler()
	readData := handler.ReadData("../../algorithm/data/contributor_output.parquet", sys.UnCleansingUserId)
	_ = readData.Contributors

	data := handler.ReadData("../../algorithm/data/repo_output.parquet", sys.UnCleansingRepoId)
	_ = data.Repos

}

func NewStorageEngine(b *bootstrap.SearchEngine) ParquetStorageEngine {
	return &storageEngine{es: b}
}
