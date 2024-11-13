package tests

import (
	"os"
	"testing"

	"github.com/fleimkeipa/case/model"
	"github.com/fleimkeipa/case/pkg"
	"github.com/fleimkeipa/case/repositories"
)

func init() {
	os.Setenv("API_KEY", "")
	os.Setenv("API_SECRET", "")
	os.Setenv("SUPPLIER_ID", "")
}

func TestProductAPIRepository_FindAll(t *testing.T) {
	type fields struct {
		Client pkg.Client
	}
	type args struct {
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
				Client: pkg.NewHTTPClient(os.Getenv("API_KEY"), os.Getenv("API_SECRET")),
			},
			args: args{
				opts: model.ProductListOpts{
					PaginationOpts: model.PaginationOpts{
						Size: 10,
						Page: 0,
					},
					SuplierID: model.Filter{
						Value: os.Getenv("SUPPLIER_ID"),
					},
				},
			},
			want:    &model.ProductsResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := &repositories.ProductAPIRepository{
				Client: tt.fields.Client,
			}
			got, err := rc.FindAll(tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductAPIRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Content == nil {
				t.Errorf("ProductAPIRepository.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
