package crawl

import (
	"gae-backend-crawl/bootstrap"
	"gae-backend-crawl/infrastructure/bloom"
)

func Setup(db *bootstrap.Databases, pools *bootstrap.PoolsFactory, env *bootstrap.Env, bloom bloom.Client) {
	crawl := NewRepoCrawl(bloom)
	crawl.DoCrawl()
}
