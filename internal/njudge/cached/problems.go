package cached

import (
	"context"
	"github.com/erni27/imcache"
	"github.com/mraron/njudge/internal/njudge"
	"time"
)

type problemAndUser struct {
	ProblemID, UserID int
}

type SolvedStatusQuery struct {
	solvedStatusQuery njudge.SolvedStatusQuery
	cache             *imcache.Cache[problemAndUser, njudge.SolvedStatus]
	userCache         *imcache.Cache[int, map[int]njudge.SolvedStatus]
}

func (ssq *SolvedStatusQuery) GetSolvedStatus(ctx context.Context, problemID, userID int) (njudge.SolvedStatus, error) {
	key := problemAndUser{problemID, userID}
	if val, ok := ssq.cache.Get(key); ok {
		return val, nil
	}

	val, err := ssq.solvedStatusQuery.GetSolvedStatus(ctx, problemID, userID)
	if err != nil {
		return njudge.Unknown, err
	}

	ssq.cache.Set(key, val, imcache.WithDefaultExpiration())
	return val, nil
}

// GetSolvedStatusesForUser caches the whole per-user map. The cached map is shared between callers and must not be modified.
func (ssq *SolvedStatusQuery) GetSolvedStatusesForUser(ctx context.Context, userID int) (map[int]njudge.SolvedStatus, error) {
	if val, ok := ssq.userCache.Get(userID); ok {
		return val, nil
	}

	val, err := ssq.solvedStatusQuery.GetSolvedStatusesForUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	ssq.userCache.Set(userID, val, imcache.WithDefaultExpiration())
	return val, nil
}

func NewSolvedStatusQuery(solvedStatusQuery njudge.SolvedStatusQuery, ttl time.Duration) *SolvedStatusQuery {
	return &SolvedStatusQuery{
		cache: imcache.New[problemAndUser, njudge.SolvedStatus](
			imcache.WithDefaultExpirationOption[problemAndUser, njudge.SolvedStatus](ttl),
		),
		userCache: imcache.New[int, map[int]njudge.SolvedStatus](
			imcache.WithDefaultExpirationOption[int, map[int]njudge.SolvedStatus](ttl),
		),
		solvedStatusQuery: solvedStatusQuery,
	}
}
