package services

import (
	"context"
	"testing"

	"github.com/demo-talent/internal/entities"
	"github.com/demo-talent/internal/repository/mocks"
	"github.com/golang/mock/gomock"
)

func Test_BudgetService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBudgetRepositoryInterface(ctrl)

	type fields struct {
		repo *mocks.MockBudgetRepositoryInterface
	}
	type args struct {
		ctx context.Context
		b   *entities.Budget
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		setupMock func(*mocks.MockBudgetRepositoryInterface)
	}{
		{
			name: "CreateBudget_Success",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				ctx: context.TODO(),
				b:   &entities.Budget{Description: "Test budget", Amount: 1000, StartDate: "2021-01-01", EndDate: "2021-12-31"},
			},
			wantErr: false,
			setupMock: func(m *mocks.MockBudgetRepositoryInterface) {
				m.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil) // Expect the Create method to be called once with any arguments and to return nil
			},
		},
		// Add test to simulate an insert
		
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock(tt.fields.repo)

			s := &BudgetService{
				repo: tt.fields.repo,
			}

			if err := s.Create(tt.args.ctx, tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("BudgetService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}