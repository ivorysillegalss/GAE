package crawl

import (
	"context"
	"github.com/google/go-github/github"
	"log"
	"sync"
	"time"

	"golang.org/x/oauth2"
)

// TokenManager 负责管理多个 GitHub tokens
type TokenManager struct {
	tokens      []string        // 存储所有 token
	tokenStatus map[string]bool // 记录每个 token 是否可用
	mu          sync.Mutex      // 保护 token 的并发访问
	current     int             // 当前使用的 token 索引
}

// NewTokenManager 初始化 TokenManager
func NewTokenManager(tokens []string) *TokenManager {
	tokenStatus := make(map[string]bool)
	for _, token := range tokens {
		tokenStatus[token] = true // 初始状态为可用
	}
	return &TokenManager{
		tokens:      tokens,
		tokenStatus: tokenStatus,
		current:     0,
	}
}

// GetClient 返回一个 OAuth2 客户端，使用当前的可用 token
func (tm *TokenManager) GetClient(ctx context.Context) *github.Client {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	token := tm.getCurrentToken()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

// getCurrentToken 获取当前可用的 token，如果达到限额则切换
func (tm *TokenManager) getCurrentToken() string {
	for {
		token := tm.tokens[tm.current]
		if tm.tokenStatus[token] {
			return token
		}
		tm.current = (tm.current + 1) % len(tm.tokens)
		time.Sleep(2 * time.Second)
	}
}

// MarkTokenAsLimited 将当前 token 标记为达到限额
func (tm *TokenManager) MarkTokenAsLimited() {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	currentToken := tm.tokens[tm.current]
	tm.tokenStatus[currentToken] = false

	// 设置定时器，在限额重置时间后恢复 token 可用状态
	resetTime := time.Now().Add(time.Hour) // GitHub API 重置时间一般为 1 小时
	go func(token string) {
		time.Sleep(resetTime.Sub(time.Now()))
		tm.mu.Lock()
		tm.tokenStatus[token] = true
		tm.mu.Unlock()
	}(currentToken)

	// 切换到下一个 token
	tm.current = (tm.current + 1) % len(tm.tokens)
}

// CheckRateLimit 检查当前 token 的限额
func (tm *TokenManager) CheckRateLimit(client *github.Client) bool {
	rate, _, err := client.RateLimits(context.Background())
	if err != nil {
		log.Printf("Error fetching rate limit: %v", err)
		return false
	}

	coreLimit := rate.Core
	if coreLimit.Remaining == 0 {
		tm.MarkTokenAsLimited()
		return true
	}
	return false
}
