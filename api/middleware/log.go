package middleware

import (
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	slogmulti "github.com/samber/slog-multi"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"opencsg.com/csghub-server/api/httpbase"
	"opencsg.com/csghub-server/common/config"
)

func Log(config *config.Config) gin.HandlerFunc {
	handlers := []slog.Handler{
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: false,
			Level:     slog.LevelInfo,
		}),
	}
	if config.Instrumentation.OTLPEndpoint != "" && config.Instrumentation.OTLPLogging {
		handlers = append(handlers, otelslog.NewHandler("csghub-server"))
	}

	l := slog.New(slogmulti.Fanout(handlers...))
	return func(ctx *gin.Context) {
		startTime := time.Now()

		ctx.Next()

		latency := time.Since(startTime).Milliseconds()
		l.InfoContext(ctx, "http request", slog.String("ip", ctx.ClientIP()),
			slog.String("method", ctx.Request.Method),
			slog.Int("latency(ms)", int(latency)),
			slog.Int("status", ctx.Writer.Status()),
			slog.String("current_user", httpbase.GetCurrentUser(ctx)),
			slog.Any("auth_type", httpbase.GetAuthType(ctx)),
			slog.String("url", ctx.Request.URL.RequestURI()),
			slog.String("full_path", ctx.FullPath()),
		)
	}
}
