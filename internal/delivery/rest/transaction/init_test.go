package transaction

import (
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"github.com/soerjadi/brick/internal/delivery/rest"
	"github.com/soerjadi/brick/internal/usecase/transaction"
)

func TestNewHandler(t *testing.T) {
	type args struct {
		usecase transaction.Usecase
	}
	tests := []struct {
		name string
		args args
		want rest.API
	}{
		{
			name: "success",
			args: args{
				usecase: nil,
			},
			want: &Handler{
				usecase: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.usecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_RegisterRoutes(t *testing.T) {
	type fields struct {
		usecase transaction.Usecase
	}
	type args struct {
		r *mux.Router
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "pass",
			args: args{
				r: mux.NewRouter(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				usecase: tt.fields.usecase,
			}
			h.RegisterRoutes(tt.args.r)
		})
	}
}
