// Code generated by hertz generator.

package Api

import (
	"context"

	"github.com/CyanAsterisk/FreeCar/shared/middleware"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/limiter"
	"github.com/hertz-contrib/requestid"
	"go.opentelemetry.io/otel/trace"
)

func rootMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		// use gzip mw
		gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions([]string{".jpg", ".mp4", ".png"})),
		// use limiter mw
		limiter.AdaptiveLimit(limiter.WithCPUThreshold(900)),
		// use requestId mw & bind with traceId
		requestid.New(
			requestid.WithGenerator(func(ctx context.Context, c *app.RequestContext) string {
				traceID := trace.SpanFromContext(ctx).SpanContext().TraceID().String()
				return traceID
			}),
		),
	}
}

func _createcarMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}

func _getcarMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}

func _submitprofileMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}

func _clearprofileMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}

func _tripMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}

func _updatetripMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}

func _gettripMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}

func _gettripsMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}

func _authMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _profileMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}

func _createprofilephotoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _clearprofilephotoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _photoMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}

func _completeprofilephotoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getcarsMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}

func _upload_vatarMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}

func _getuserinfoMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}

func _updateuserinfoMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}
