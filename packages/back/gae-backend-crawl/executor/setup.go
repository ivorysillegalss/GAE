package executor

import (
	"gae-backend-crawl/domain"
)

type CrawlExecutor struct {
	repoCrawl domain.RepoCrawl
}

func (e *CrawlExecutor) Setup() {
	e.repoCrawl.DoCrawl()
}

func NewCrawlExecutor(crawl domain.RepoCrawl) *CrawlExecutor {
	return &CrawlExecutor{repoCrawl: crawl}
}
