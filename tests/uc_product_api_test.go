package tests

import (
	"context"
	"os"
	"testing"

	"github.com/fleimkeipa/case/model"
	"github.com/fleimkeipa/case/pkg"
	"github.com/fleimkeipa/case/repositories"
	"github.com/fleimkeipa/case/repositories/interfaces"
	"github.com/fleimkeipa/case/uc"
)

func init() {
	os.Setenv("API_KEY", "")
	os.Setenv("API_SECRET", "")
	os.Setenv("SUPPLIER_ID", "")
}

func TestProductAPIUC_FindAll(t *testing.T) {
	type fields struct {
		apiRepo interfaces.ProductAPIRepository
		dbRepo  interfaces.ProductDBRepository
		cacheUC *uc.ProductCacheUC
	}
	type args struct {
		ctx  context.Context
		opts model.ProductListOpts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.ProductsResponse
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				apiRepo: repositories.NewProductAPIRepository(pkg.NewHTTPClient(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))),
				dbRepo:  repositories.NewProductDBRepository(pkg.NewPSQLClient()),
				cacheUC: uc.NewProductCacheUC(repositories.NewCacheRepository(pkg.NewRedisClient())),
			},
			args: args{
				ctx: context.TODO(),
				opts: model.ProductListOpts{
					PaginationOpts: model.PaginationOpts{
						Page: 0,
						Size: 30,
					},
					SuplierID: model.Filter{
						Value:    os.Getenv("SUPPLIER_ID"),
						IsSended: true,
					},
				},
			},
			want:    &model.ProductsResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cacheUC := uc.NewProductDBUC(tt.fields.dbRepo, tt.fields.cacheUC)
			rc := uc.NewProductAPIUC(tt.fields.apiRepo, *cacheUC)
			got, err := rc.FindAll(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductAPIUC.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Content == nil {
				t.Errorf("ProductAPIUC.FindAll() got.Content = %v, want %v", got.Content, tt.want.Content)
			}
		})
	}
}
