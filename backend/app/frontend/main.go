// Code generated by hertz generator.

package main

import (
	"context"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"github.com/hertz-contrib/pprof"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
	"github.com/lpernett/godotenv"
	"github.com/zjhM3l/go-e-commerce/backend/app/frontend/biz/router"
	"github.com/zjhM3l/go-e-commerce/backend/app/frontend/conf"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// load env
	godotenv.Load()	
	// init dal
	// dal.Init()
	address := conf.GetConf().Hertz.Address
	h := server.New(server.WithHostPorts(address))

	registerMiddleware(h)

	// add a ping route to test
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	router.GeneratedRegister(h)
	// load template
	h.LoadHTMLGlob("template/*")
	// load static file
	h.Static("/static", "./")

	// sign-in
	// tmpl里面{{define "sign-in"}}{{end}}, 这样可以不用写后缀
	h.GET("/sign-in", func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(consts.StatusOK, "sign-in", utils.H{"title": "Sign In"})
	})

	// about， 同理无后缀
	h.GET("/about", func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(consts.StatusOK, "about", utils.H{"title": "About"})
	})

	// sign-up，同理无后缀
	h.GET("/sign-up", func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(consts.StatusOK, "sign-up", utils.H{"title": "Sign Up"})
	})

	h.Spin()
}

func registerMiddleware(h *server.Hertz) {
	// session
	// 为了方便后续更改，使用conf.GetConf().Redis.Address，config.yaml中配置
	// 隐私信息不应该写在代码中，改成从环境变量中读取.env
	store, _ := redis.NewStore(10, "tcp", conf.GetConf().Redis.Address, "", []byte(os.Getenv("SESSION_SECRET")))
	h.Use(sessions.New("go-ecommerce", store))

	// log
	logger := hertzlogrus.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Hertz.LogFileName,
			MaxSize:    conf.GetConf().Hertz.LogMaxSize,
			MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
			MaxAge:     conf.GetConf().Hertz.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	hlog.SetOutput(asyncWriter)
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		asyncWriter.Sync()
	})

	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}

	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())
}
