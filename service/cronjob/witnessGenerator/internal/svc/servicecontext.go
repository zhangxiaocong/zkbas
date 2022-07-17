package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/zecrey-labs/zecrey-legend/common/model/account"
	"github.com/zecrey-labs/zecrey-legend/common/model/block"
	"github.com/zecrey-labs/zecrey-legend/common/model/blockForProof"
	"github.com/zecrey-labs/zecrey-legend/common/model/liquidity"
	"github.com/zecrey-labs/zecrey-legend/common/model/nft"
	"github.com/zecrey-labs/zecrey-legend/common/model/proofSender"
	"github.com/zecrey-labs/zecrey-legend/service/cronjob/witnessGenerator/internal/config"
)

type ServiceContext struct {
	Config config.Config

	RedisConn *redis.Redis

	BlockModel            block.BlockModel
	AccountModel          account.AccountModel
	AccountHistoryModel   account.AccountHistoryModel
	LiquidityHistoryModel liquidity.LiquidityHistoryModel
	NftHistoryModel       nft.L2NftHistoryModel
	ProofSenderModel      proofSender.ProofSenderModel
	BlockForProofModel    blockForProof.BlockForProofModel
}

func WithRedis(redisType string, redisPass string) redis.Option {
	return func(p *redis.Redis) {
		p.Type = redisType
		p.Pass = redisPass
	}
}

func NewServiceContext(c config.Config) *ServiceContext {
	gormPointer, err := gorm.Open(postgres.Open(c.Postgres.DataSource))
	if err != nil {
		logx.Errorf("gorm connect db error, err = %s", err.Error())
	}
	conn := sqlx.NewSqlConn("postgres", c.Postgres.DataSource)
	redisConn := redis.New(c.CacheRedis[0].Host, WithRedis(c.CacheRedis[0].Type, c.CacheRedis[0].Pass))
	return &ServiceContext{
		Config:                c,
		RedisConn:             redisConn,
		BlockModel:            block.NewBlockModel(conn, c.CacheRedis, gormPointer, redisConn),
		BlockForProofModel:    blockForProof.NewBlockForProofModel(conn, c.CacheRedis, gormPointer),
		AccountModel:          account.NewAccountModel(conn, c.CacheRedis, gormPointer),
		AccountHistoryModel:   account.NewAccountHistoryModel(conn, c.CacheRedis, gormPointer),
		LiquidityHistoryModel: liquidity.NewLiquidityHistoryModel(conn, c.CacheRedis, gormPointer),
		NftHistoryModel:       nft.NewL2NftHistoryModel(conn, c.CacheRedis, gormPointer),
		ProofSenderModel:      proofSender.NewProofSenderModel(gormPointer),
	}
}