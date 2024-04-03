package services

import (
	"context"
	"errors"
	"testing"

	"github.com/demo-talent/internal/entities"
	"github.com/demo-talent/internal/repository/mocks"
	"github.com/golang/mock/gomock"
)

func Test_expenseServiceImpl_CreateExpense(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockExpenseRepositoryInterface(ctrl)

	type fields struct {
		repo *mocks.MockExpenseRepositoryInterface
	}
	type args struct {
		ctx context.Context
		e   *entities.Expense
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		setupMock func(*mocks.MockExpenseRepositoryInterface)
	}{
		{
			name: "CreateExpense_Success",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				ctx: context.TODO(),
				e:   &entities.Expense{ID: "1", Description: "Test expense"},
			},
			wantErr: false,
			setupMock: func(m *mocks.MockExpenseRepositoryInterface) {
				m.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil) // Expect the Create method to be called once with any arguments and to return nil
			},
		},
		// Add test to simulate an insert
		{
			name: "CreateExpense_Error",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				ctx: context.TODO(),
				e:   &entities.Expense{ID: "1", Description: "Test expense"},
			},
			wantErr: true,
			setupMock: func(m *mocks.MockExpenseRepositoryInterface) {
				m.EXPECT().Create(gomock.Any(), gomock.Any()).Return(errors.New("error")) // Expect the Create method to be called once with any arguments and to return an error
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock(tt.fields.repo) // Setup the mock expectations
			s := &expenseServiceImpl{
				repo: tt.fields.repo,
			}
			if err := s.CreateExpense(tt.args.ctx, tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("expenseServiceImpl.CreateExpense() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
