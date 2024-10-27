package application

import (
	"errors"
	"github.com/iqbalbachmid/billing-engine/domain/entity"
	"github.com/iqbalbachmid/billing-engine/domain/repository"
	mocks "github.com/iqbalbachmid/billing-engine/mocks/domain/repository"
	"testing"
	"time"
)

func TestLoanService_GetOutstanding(t *testing.T) {
	mockLoanScheduleRepo := mocks.NewLoanScheduleRepository(t)

	type fields struct {
		loanRepo         repository.LoanRepository
		loanScheduleRepo repository.LoanScheduleRepository
		paymentRepo      repository.PaymentRepository
	}
	type args struct {
		loanID int
	}
	var tests = []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
		mock    func()
	}{
		{
			name: "should return error if loan not found",
			fields: fields{
				loanScheduleRepo: mockLoanScheduleRepo,
			},
			args: args{
				loanID: 1,
			},
			want:    0,
			wantErr: true,
			mock: func() {
				mockLoanScheduleRepo.EXPECT().GetByLoanID(1).Return(nil, errors.New("loan not found")).Once()
			},
		},
		{
			name: "should return outstanding successfully",
			fields: fields{
				loanScheduleRepo: mockLoanScheduleRepo,
			},
			args: args{
				loanID: 1,
			},
			want:    220000,
			wantErr: false,
			mock: func() {
				mockLoanScheduleRepo.EXPECT().GetByLoanID(1).Return([]entity.LoanSchedule{
					{
						ScheduleID:      1,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusPaid,
					},
					{
						ScheduleID:      2,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusDue,
					},
					{
						ScheduleID:      3,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusUnspecified,
					},
				}, nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			s := &LoanService{
				loanRepo:         tt.fields.loanRepo,
				loanScheduleRepo: tt.fields.loanScheduleRepo,
				paymentRepo:      tt.fields.paymentRepo,
			}
			got, err := s.GetOutstanding(tt.args.loanID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOutstanding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetOutstanding() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoanService_IsDelinquent(t *testing.T) {
	mockLoanScheduleRepo := mocks.NewLoanScheduleRepository(t)

	type fields struct {
		loanRepo         repository.LoanRepository
		loanScheduleRepo repository.LoanScheduleRepository
		paymentRepo      repository.PaymentRepository
	}
	type args struct {
		loanID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
		mock    func()
	}{
		{
			name: "should return error if loan not found",
			fields: fields{
				loanScheduleRepo: mockLoanScheduleRepo,
			},
			args: args{
				loanID: 1,
			},
			want:    false,
			wantErr: true,
			mock: func() {
				mockLoanScheduleRepo.EXPECT().GetByLoanID(1).Return(nil, errors.New("loan not found")).Once()
			},
		},
		{
			name: "should return true if overdue = 2",
			fields: fields{
				loanScheduleRepo: mockLoanScheduleRepo,
			},
			args: args{
				loanID: 1,
			},
			want:    true,
			wantErr: false,
			mock: func() {
				mockLoanScheduleRepo.EXPECT().GetByLoanID(1).Return([]entity.LoanSchedule{
					{
						ScheduleID:      1,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusPaid,
					},
					{
						ScheduleID:      2,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusOverdue,
					},
					{
						ScheduleID:      3,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusOverdue,
					},
					{
						ScheduleID:      4,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusUnspecified,
					},
				}, nil).Once()
			},
		},
		{
			name: "should return false if overdue < 2",
			fields: fields{
				loanScheduleRepo: mockLoanScheduleRepo,
			},
			args: args{
				loanID: 1,
			},
			want:    false,
			wantErr: false,
			mock: func() {
				mockLoanScheduleRepo.EXPECT().GetByLoanID(1).Return([]entity.LoanSchedule{
					{
						ScheduleID:      1,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusPaid,
					},
					{
						ScheduleID:      2,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusOverdue,
					},
					{
						ScheduleID:      3,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusUnspecified,
					},
					{
						ScheduleID:      4,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusUnspecified,
					},
				}, nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			s := &LoanService{
				loanRepo:         tt.fields.loanRepo,
				loanScheduleRepo: tt.fields.loanScheduleRepo,
				paymentRepo:      tt.fields.paymentRepo,
			}
			got, err := s.IsDelinquent(tt.args.loanID)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsDelinquent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsDelinquent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoanService_MakePayment(t *testing.T) {
	var (
		mockLoanScheduleRepository = mocks.NewLoanScheduleRepository(t)
		mockPaymentRepository      = mocks.NewPaymentRepository(t)
	)

	type fields struct {
		loanRepo         repository.LoanRepository
		loanScheduleRepo repository.LoanScheduleRepository
		paymentRepo      repository.PaymentRepository
	}
	type args struct {
		loanID        int
		paymentAmount float64
		paymentMethod string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "should return error if loan not found",
			fields: fields{
				loanScheduleRepo: mockLoanScheduleRepository,
			},
			args: args{
				loanID: 1,
			},
			wantErr: true,
			mock: func() {
				mockLoanScheduleRepository.EXPECT().GetByLoanID(1).Return(nil, errors.New("loan not found")).Once()
			},
		},
		{
			name: "should return error if payment amount is more than outstanding",
			fields: fields{
				loanScheduleRepo: mockLoanScheduleRepository,
			},
			args: args{
				loanID:        1,
				paymentAmount: 230000,
			},
			wantErr: true,
			mock: func() {
				mockLoanScheduleRepository.EXPECT().GetByLoanID(1).Return([]entity.LoanSchedule{
					{
						ScheduleID:      1,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusPaid,
					},
					{
						ScheduleID:      2,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusPaid,
					},
					{
						ScheduleID:      3,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusDue,
					},
					{
						ScheduleID:      4,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusUnspecified,
					},
				}, nil).Twice()
			},
		},
		{
			name: "should return error if payment amount is not multiple of schedule amount",
			fields: fields{
				loanScheduleRepo: mockLoanScheduleRepository,
			},
			args: args{
				loanID:        1,
				paymentAmount: 230000,
			},
			wantErr: true,
			mock: func() {
				mockLoanScheduleRepository.EXPECT().GetByLoanID(1).Return([]entity.LoanSchedule{
					{
						ScheduleID:      1,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusPaid,
					},
					{
						ScheduleID:      2,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusPaid,
					},
					{
						ScheduleID:      3,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusDue,
					},
					{
						ScheduleID:      4,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusUnspecified,
					},
					{
						ScheduleID:      5,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusUnspecified,
					},
				}, nil).Twice()
			},
		},
		{
			name: "should return error if payment repo fail",
			fields: fields{
				loanScheduleRepo: mockLoanScheduleRepository,
				paymentRepo:      mockPaymentRepository,
			},
			args: args{
				loanID:        1,
				paymentAmount: 220000,
				paymentMethod: "bank transfer",
			},
			wantErr: true,
			mock: func() {
				mockLoanScheduleRepository.EXPECT().GetByLoanID(1).Return([]entity.LoanSchedule{
					{
						ScheduleID:      1,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusPaid,
					},
					{
						ScheduleID:      2,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusPaid,
					},
					{
						ScheduleID:      3,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusDue,
					},
					{
						ScheduleID:      4,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusUnspecified,
					},
					{
						ScheduleID:      5,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusUnspecified,
					},
				}, nil).Twice()

				mockPaymentRepository.EXPECT().CreatePaymentAndUpdateLoanSchedules(entity.Payment{
					LoanID:        1,
					PaymentDate:   time.Date(2024, time.October, 28, 0, 0, 0, 0, time.UTC),
					AmountPaid:    220000,
					PaymentMethod: "bank transfer",
					Status:        entity.Status,
				}, []entity.LoanSchedule{
					{
						ScheduleID:      3,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusPaid,
					},
					{
						ScheduleID:      4,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusPaid,
					},
				}).Return(0, errors.New("failed to create payment")).Once()
			},
		},
		{
			name: "should not return error and make payment successfully",
			fields: fields{
				loanScheduleRepo: mockLoanScheduleRepository,
				paymentRepo:      mockPaymentRepository,
			},
			args: args{
				loanID:        1,
				paymentAmount: 220000,
				paymentMethod: "bank transfer",
			},
			wantErr: false,
			mock: func() {
				mockLoanScheduleRepository.EXPECT().GetByLoanID(1).Return([]entity.LoanSchedule{
					{
						ScheduleID:      1,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusPaid,
					},
					{
						ScheduleID:      2,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusPaid,
					},
					{
						ScheduleID:      3,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusDue,
					},
					{
						ScheduleID:      4,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusUnspecified,
					},
					{
						ScheduleID:      5,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusUnspecified,
					},
				}, nil).Twice()

				mockPaymentRepository.EXPECT().CreatePaymentAndUpdateLoanSchedules(entity.Payment{
					LoanID:        1,
					PaymentDate:   time.Date(2024, time.October, 28, 0, 0, 0, 0, time.UTC),
					AmountPaid:    220000,
					PaymentMethod: "bank transfer",
					Status:        entity.Status,
				}, []entity.LoanSchedule{
					{
						ScheduleID:      3,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusPaid,
					},
					{
						ScheduleID:      4,
						LoanID:          1,
						PrincipalAmount: 100000,
						InterestAmount:  10000,
						TotalDue:        110000,
						PaymentStatus:   entity.PaymentStatusPaid,
					},
				}).Return(200, nil).Once()
			},
		},
	}
	for _, tt := range tests {
		tt.mock()

		t.Run(tt.name, func(t *testing.T) {
			s := &LoanService{
				loanRepo:         tt.fields.loanRepo,
				loanScheduleRepo: tt.fields.loanScheduleRepo,
				paymentRepo:      tt.fields.paymentRepo,
				timeNow: func() time.Time {
					return time.Date(2024, time.October, 28, 0, 0, 0, 0, time.UTC)
				},
			}
			if err := s.MakePayment(tt.args.loanID, tt.args.paymentAmount, tt.args.paymentMethod); (err != nil) != tt.wantErr {
				t.Errorf("MakePayment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
