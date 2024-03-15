package transaction

import (
	"context"
	"reflect"
	"testing"

	"github.com/soerjadi/brick/internal/mocks"
	"github.com/soerjadi/brick/internal/model"
	"github.com/soerjadi/brick/internal/repository/transaction"
	"go.uber.org/mock/gomock"
)

func Test_trxUsecase_CheckAccount(t *testing.T) {
	type fields struct {
		repository transaction.Repository
	}
	type args struct {
		ctx context.Context
		req model.CheckAccountRequest
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Account
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				req: model.CheckAccountRequest{},
			},
			fields: fields{
				repository: func() transaction.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockTransactionRepository(ctrl)

					mock.EXPECT().
						CheckAccount(gomock.Any(), gomock.Any()).
						Return(&model.Account{
							ID:         "1",
							Name:       "satu",
							Number:     "satunumber123",
							OriginBank: "Mandiri",
						}, nil)

					return mock
				}(),
			},
			want: &model.Account{
				ID:         "1",
				Name:       "satu",
				Number:     "satunumber123",
				OriginBank: "Mandiri",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := &trxUsecase{
				repository: tt.fields.repository,
			}

			got, err := usecase.CheckAccount(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.CheckAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.CheckAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
