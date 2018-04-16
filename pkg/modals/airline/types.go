package airline

import (
	"code.ysitd.cloud/component/aviation/runway"
	"context"
)

type Store interface {
	GetRevision(ctx context.Context, revision string) (runway.Airline, error)
}
