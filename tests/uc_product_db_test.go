package tests

import (
	"context"
	"reflect"
	"testing"

	"github.com/fleimkeipa/case/model"
	"github.com/fleimkeipa/case/pkg"
	"github.com/fleimkeipa/case/repositories"
	"github.com/fleimkeipa/case/repositories/interfaces"
	"github.com/fleimkeipa/case/uc"
)

func TestProductDBUC_Create(t *testing.T) {
	testCache, _ = pkg.GetCacheTestInstance(context.Background())
	testDB, terminateContainer = pkg.GetTestInstance(context.Background())
	defer terminateContainer()

	type fields struct {
		repo  interfaces.ProductDBRepository
		cache *uc.ProductCacheUC
	}
	type tempData struct {
		brandID int
		barcode string
	}
	type args struct {
		ctx     context.Context
		product *model.Product
	}
	tests := []struct {
		name      string
		fields    fields
		tempDatas []tempData
		args      args
		want      *model.Product
		wantErr   bool
	}{
		{
			name: "correct - not exist",
			fields: fields{
				repo:  repositories.NewProductDBRepository(testDB),
				cache: uc.NewProductCacheUC(repositories.NewCacheRepository(testCache)),
			},
			args: args{
				ctx: context.TODO(),
				product: &model.Product{
					Barcode: "1234",
					BrandID: 12,
				},
			},
			want: &model.Product{
				Barcode: "1234",
				BrandID: 12,
			},
			wantErr: false,
		},
		{
			name: "correct - exist",
			fields: fields{
				repo:  repositories.NewProductDBRepository(testDB),
				cache: uc.NewProductCacheUC(repositories.NewCacheRepository(testCache)),
			},
			tempDatas: []tempData{
				{
					brandID: 12,
					barcode: "1234",
				},
			},
			args: args{
				ctx: context.TODO(),
				product: &model.Product{
					Barcode: "1234",
					BrandID: 12,
				},
			},
			want: &model.Product{
				Barcode: "1234",
				BrandID: 12,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := uc.NewProductDBUC(tt.fields.repo, tt.fields.cache)
			got, err := rc.Create(tt.args.ctx, tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductDBUC.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductDBUC.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
