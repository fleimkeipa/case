package tests

import (
	"context"
	"testing"

	"github.com/fleimkeipa/case/pkg"
	"github.com/fleimkeipa/case/repositories"
	"github.com/fleimkeipa/case/repositories/interfaces"
	"github.com/fleimkeipa/case/uc"
)

func TestProductCacheUC_IsExist(t *testing.T) {
	testCache, terminateContainer = pkg.GetCacheTestInstance(context.Background())
	defer terminateContainer()

	type fields struct {
		repo interfaces.CacheRepository
	}
	type tempData struct {
		brandID int
		barcode string
	}
	type args struct {
		ctx     context.Context
		brandID int
		barcode string
	}
	tests := []struct {
		name      string
		fields    fields
		tempDatas []tempData
		args      args
		want      bool
	}{
		{
			name: "correct - not exist",
			fields: fields{
				repo: repositories.NewCacheRepository(testCache),
			},
			args: args{
				ctx:     context.Background(),
				brandID: 123,
				barcode: "abc",
			},
			want: false,
		},
		{
			name: "correct - exist",
			fields: fields{
				repo: repositories.NewCacheRepository(testCache),
			},
			tempDatas: []tempData{
				{
					brandID: 123,
					barcode: "abc",
				},
			},
			args: args{
				ctx:     context.Background(),
				brandID: 123,
				barcode: "abc",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, tempData := range tt.tempDatas {
				addTestCacheData(tt.args.ctx, uc.ProductCacheID(tempData.brandID, tempData.barcode), tempData.barcode)
			}
			rc := uc.NewProductCacheUC(tt.fields.repo)
			if got := rc.IsExist(tt.args.ctx, tt.args.brandID, tt.args.barcode); got != tt.want {
				t.Errorf("ProductCacheUC.IsExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
