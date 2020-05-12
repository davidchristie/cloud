//go:generate go run github.com/vektah/dataloaden ProductLoader github.com/google/uuid.UUID *github.com/davidchristie/cloud/pkg/gateway/graph/model.Product

package dataloader

import (
	"context"
	"net/http"
	"time"

	"github.com/davidchristie/cloud/pkg/gateway/graph/convert"
	"github.com/davidchristie/cloud/pkg/gateway/graph/model"
	productReadAPI "github.com/davidchristie/cloud/pkg/product/read/api"
	"github.com/google/uuid"
)

const loadersKey = "dataloaders"

type Loaders struct {
	Product ProductLoader
}

func Middleware(productReadAPI productReadAPI.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {
			ctx := context.WithValue(req.Context(), loadersKey, &Loaders{
				Product: ProductLoader{
					maxBatch: 100,
					wait:     1 * time.Millisecond,
					fetch: func(ids []uuid.UUID) ([]*model.Product, []error) {
						if len(ids) == 1 {
							product, err := productReadAPI.Product(ids[0])
							if err != nil {
								return nil, []error{err}
							}
							return []*model.Product{convert.Product(product)}, nil
						} else {
							productsByID, err := productReadAPI.Products(ids)
							if err != nil {
								return nil, []error{err}
							}
							products := make([]*model.Product, len(ids))
							for i, id := range ids {
								products[i] = convert.Product(productsByID[id])
							}
							return products, nil
						}
					},
				},
			})
			req = req.WithContext(ctx)
			next.ServeHTTP(wrt, req)
		})
	}
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
