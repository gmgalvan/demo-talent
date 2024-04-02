package services

import (
    "context"
    "testing"
    "fmt"
    "github.com/demo-talent/entities"
    "github.com/demo-talent/repository"
    "github.com/demo-talent/repository/mocks"
    "github.com/golang/mock/gomock"
)

func Test_budgetService_CreateBudget(t *testing.T) {
    type fields struct {
        repo repository.BudgetRepositoryInterface
    }
    type args struct {
        ctx context.Context
        b   *entities.Budget
    }
    tests := []struct {
        name      string
        fields    fields
        args      args
        setupMock func(*mocks.MockBudgetRepositoryInterface)
        wantErr   bool
    }{
        {
            name: "CreateBudget_Success",
            fields: fields{
                repo: mocks.NewMockBudgetRepositoryInterface(gomock.NewController(t)),
            },
            setupMock: func(m *mocks.MockBudgetRepositoryInterface) {
                m.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
            },
            args: args{
                ctx: context.TODO(),
                b:   &entities.Budget{ID: "1", Amount: 123},
            },
            wantErr: false,
        },
        // failed test case
        {
            name: "CreateBudget_Failure",
            fields: fields{
                repo: mocks.NewMockBudgetRepositoryInterface(gomock.NewController(t)),
            },
            args: args{
                ctx: context.TODO(),
                b:   &entities.Budget{ID: "1", Amount: 1000},
            },
            setupMock: func(m *mocks.MockBudgetRepositoryInterface) {
                m.EXPECT().Create(gomock.Any(), gomock.Any()).Return(fmt.Errorf("error creating budget"))
            },
            wantErr: true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            s := &budgetService{
                repo: tt.fields.repo,
            }
            tt.setupMock(tt.fields.repo.(*mocks.MockBudgetRepositoryInterface))
            if err := s.CreateBudget(tt.args.ctx, tt.args.b); (err != nil) != tt.wantErr {
                t.Errorf("budgetService.CreateBudget() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}