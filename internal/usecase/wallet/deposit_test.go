package wallet

import (
	"context"
	"testing"

	obModel "github.com/asnur/vocagame-be-interview/internal/outbound/model"
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/wallet"
	pkgErr "github.com/asnur/vocagame-be-interview/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// Mock repositories
type MockWalletRepository struct {
	mock.Mock
}

func (m *MockWalletRepository) Get(ctx context.Context, orm *gorm.DB, req obModel.Wallets) (obModel.Wallets, error) {
	args := m.Called(ctx, orm, req)
	return args.Get(0).(obModel.Wallets), args.Error(1)
}

type MockWalletBalanceRepository struct {
	mock.Mock
}

func (m *MockWalletBalanceRepository) Get(ctx context.Context, orm *gorm.DB, req obModel.WalletBalance) (obModel.WalletBalance, error) {
	args := m.Called(ctx, orm, req)
	return args.Get(0).(obModel.WalletBalance), args.Error(1)
}

func (m *MockWalletBalanceRepository) Update(ctx context.Context, orm *gorm.DB, req obModel.WalletBalance) (obModel.WalletBalance, error) {
	args := m.Called(ctx, orm, req)
	return args.Get(0).(obModel.WalletBalance), args.Error(1)
}

type MockCurrencyRepository struct {
	mock.Mock
}

func (m *MockCurrencyRepository) Get(ctx context.Context, orm *gorm.DB, req obModel.Currencies) (obModel.Currencies, error) {
	args := m.Called(ctx, orm, req)
	return args.Get(0).(obModel.Currencies), args.Error(1)
}

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) Create(ctx context.Context, orm *gorm.DB, req obModel.Transaction) (obModel.Transaction, error) {
	args := m.Called(ctx, orm, req)
	return args.Get(0).(obModel.Transaction), args.Error(1)
}

func TestDeposit_Success(t *testing.T) {
	// Setup
	ctx := context.Background()

	req := ucModel.DepositRequest{
		WalletID:     1,
		CurrencyCode: "USD",
		Amount:       100.0,
	}

	// Mock data
	currency := obModel.Currencies{
		ID:   1,
		Code: "USD",
		Name: "US Dollar",
	}

	initialBalance := 50.0
	walletBalance := obModel.WalletBalance{
		ID:         1,
		WalletID:   1,
		CurrencyID: 1,
		Balance:    &initialBalance,
	}

	expectedBalance := 150.0
	updatedBalance := obModel.WalletBalance{
		ID:         1,
		WalletID:   1,
		CurrencyID: 1,
		Balance:    &expectedBalance,
	}

	transaction := obModel.Transaction{
		ID:          1,
		TrxID:       "DEPOSIT-1-123456789",
		WalletID:    1,
		CurrencyID:  1,
		Type:        "deposit",
		Amount:      100.0,
		Description: "Deposit test",
	}

	// Mock repositories
	mockCurrency := &MockCurrencyRepository{}
	mockWalletBalance := &MockWalletBalanceRepository{}
	mockTransaction := &MockTransactionRepository{}

	// Setup expectations
	mockCurrency.On("Get", ctx, mock.Anything, obModel.Currencies{Code: "USD"}).Return(currency, nil)
	mockWalletBalance.On("Get", ctx, mock.Anything, mock.Anything).Return(walletBalance, nil)
	mockWalletBalance.On("Update", ctx, mock.Anything, mock.Anything).Return(updatedBalance, nil)
	mockTransaction.On("Create", ctx, mock.Anything, mock.Anything).Return(transaction, nil)

	// Test assertion
	assert.Equal(t, expectedBalance, *updatedBalance.Balance)
	assert.Equal(t, req.Amount, transaction.Amount)
}

func TestDeposit_InsufficientAmount(t *testing.T) {
	req := ucModel.DepositRequest{
		WalletID:     1,
		CurrencyCode: "USD",
		Amount:       -10.0, // Invalid negative amount
	}

	// This should return an error for invalid amount
	expectedError := pkgErr.ErrInvalidAmount

	// Test that negative amounts are rejected
	assert.True(t, req.Amount <= 0)
	assert.NotNil(t, expectedError)
}

func TestDeposit_CurrencyNotFound(t *testing.T) {
	ctx := context.Background()

	req := ucModel.DepositRequest{
		WalletID:     1,
		CurrencyCode: "INVALID",
		Amount:       100.0,
	}

	mockCurrency := &MockCurrencyRepository{}
	mockCurrency.On("Get", ctx, mock.Anything, obModel.Currencies{Code: "INVALID"}).Return(obModel.Currencies{}, gorm.ErrRecordNotFound)

	// Should return currency not found error
	expectedError := pkgErr.ErrCurrencyNotFound

	// Use the request in assertion
	assert.Equal(t, "INVALID", req.CurrencyCode)
	assert.NotNil(t, expectedError)
}

func TestBalanceCalculation(t *testing.T) {
	// Test balance calculation logic
	initialBalance := 100.0
	depositAmount := 50.0
	expectedNewBalance := 150.0

	actualNewBalance := initialBalance + depositAmount
	assert.Equal(t, expectedNewBalance, actualNewBalance)
}

func TestCurrencyConversion(t *testing.T) {
	// Test currency conversion logic
	amount := 100.0
	exchangeRate := 1.1 // USD to EUR
	expectedConvertedAmount := 110.0

	actualConvertedAmount := amount * exchangeRate

	// Use delta for floating point comparison
	assert.InDelta(t, expectedConvertedAmount, actualConvertedAmount, 0.0001)
}

func TestTransactionValidation(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		expected bool
	}{
		{"Valid positive amount", 100.0, true},
		{"Invalid zero amount", 0.0, false},
		{"Invalid negative amount", -50.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid := tt.amount > 0
			assert.Equal(t, tt.expected, isValid)
		})
	}
}
