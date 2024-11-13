package tests

import (
	"context"
	"reflect"
	"testing"

	"github.com/fleimkeipa/case/model"
	"github.com/fleimkeipa/case/pkg"
	"github.com/fleimkeipa/case/repositories"

	"github.com/go-pg/pg"
)

func TestProductDBRepository_Create(t *testing.T) {
	testDB, terminateContainer = pkg.GetTestInstance(context.Background())
	defer terminateContainer()

	type fields struct {
		db *pg.DB
	}
	type args struct {
		product *model.Product
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Product
		wantErr bool
	}{
		{
			name: "correct",
			fields: fields{
				db: testDB,
			},
			args: args{
				product: &model.Product{
					Barcode: "1234",
					BrandID: 1,
				},
			},
			want: &model.Product{
				Barcode: "1234",
				BrandID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := repositories.NewProductDBRepository(testDB)
			got, err := rc.Create(tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductDBRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductDBRepository.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
