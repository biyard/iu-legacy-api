package main

import (
	"biyard.co/common/base"
	"biyard.co/common/logger"
	"biyard.co/iuniverse-api/config"
	m1 "biyard.co/iuniverse-api/controllers/m1"
	v1 "biyard.co/iuniverse-api/controllers/v1"
)

func main() {
	conf := config.DefaultAppConfig()
	base.Load(&conf)

	server := base.NewServer(conf.Config)
	ctx := server.Context()

	logger.Debugf(nil, "config: %+v", conf)
	server.Group("/v1/fee", v1.NewFeeController(ctx, conf))
	server.Group("/v1/auth", v1.NewAuthController(ctx, conf))
	server.Group("/v1/game", v1.NewRoulletUserController(ctx, conf))
	server.Group("/v1/image", v1.NewImageController(ctx, conf))
	server.Group("/v1/time", v1.NewTimeController(ctx, conf))
	server.Group("/v1/md", v1.NewMdController(ctx, conf))
	server.Group("/v1/leaderboard", v1.NewLeaderboardController(ctx, conf))
	server.Group("/v1/calendar", v1.NewCalendarController(ctx, conf))
	server.Group("/v1/heroes-song-contest", v1.NewHeroesSongContestController(ctx, conf))
	server.Group("/v1/mission", v1.NewMissionController(ctx, conf))
	server.Group("/m1/verification", m1.NewVerificationController(ctx, conf))
	server.Group("/m1/time", m1.NewTimeController(ctx, conf))
	server.Group("/m1/md", m1.NewMdController(ctx, conf))
	server.Group("/m1/leaderboard", m1.NewLeaderboardAdminController(ctx, conf))
	server.Group("/m1/calendar", m1.NewCalendarAdminController(ctx, conf))
	server.Group("/m1/shop", m1.NewShopAdminController(ctx, conf))
	server.Group("/m1/icp", m1.NewIcpController(ctx, conf))
	server.Group("/m1/heroes-song-contest", m1.NewHeroesSongContestAdminController(ctx, conf))
	server.Group("/m1/mission", m1.NewMissionAdminController(ctx, conf))

	server.Start()
}
