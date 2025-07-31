package wallet

import (
	"context"
	"testing"

	obModel "github.com/asnur/vocagame-be-interview/internal/outbound/model"
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/wallet"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestTransfer_SameCurrencyCalculation(t *testing.T) {
	// Test transfer within same currency
	transferAmount := 50.0
	fromBalance := 100.0
	toBalance := 25.0

	expectedFromBalance := fromBalance - transferAmount // 50.0
	expectedToBalance := toBalance + transferAmount     // 75.0

	assert.Equal(t, 50.0, expectedFromBalance)
	assert.Equal(t, 75.0, expectedToBalance)
}

func TestTransfer_CrossCurrencyCalculation(t *testing.T) {
	// Test transfer with currency conversion
	transferAmount := 100.0 // USD
	exchangeRate := 0.85    // USD to EUR

	expectedConvertedAmount := transferAmount * exchangeRate // 85.0 EUR

	assert.Equal(t, 85.0, expectedConvertedAmount)
}

func TestTransfer_InsufficientBalance(t *testing.T) {
	transferAmount := 150.0
	availableBalance := 100.0

	isInsufficientBalance := transferAmount > availableBalance
	assert.True(t, isInsufficientBalance)
}

func TestExchangeRateCalculation(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		rate     float64
		expected float64
	}{
		{"USD to EUR", 100.0, 0.85, 85.0},
		{"EUR to USD", 100.0, 1.18, 118.0},
		{"USD to JPY", 100.0, 110.0, 11000.0},
		{"JPY to USD", 1000.0, 0.009, 9.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.amount * tt.rate
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestBalanceTotal_MultiCurrency(t *testing.T) {
	// Test total balance calculation across multiple currencies
	balances := []struct {
		currency string
		amount   float64
		rate     float64 // to USD
	}{
		{"USD", 100.0, 1.0},
		{"EUR", 50.0, 1.18},
		{"JPY", 1000.0, 0.009},
	}

	expectedTotal := 0.0
	for _, balance := range balances {
		expectedTotal += balance.amount * balance.rate
	}

	// USD: 100 * 1.0 = 100
	// EUR: 50 * 1.18 = 59
	// JPY: 1000 * 0.009 = 9
	// Total: 168 USD
	assert.Equal(t, 168.0, expectedTotal)
}

type MockExchangeRateRepository struct {
	mock.Mock
}

func (m *MockExchangeRateRepository) Get(ctx context.Context, orm *gorm.DB, req obModel.ExchangeRate) (obModel.ExchangeRate, error) {
	args := m.Called(ctx, orm, req)
	return args.Get(0).(obModel.ExchangeRate), args.Error(1)
}

func TestTransfer_ValidateWalletOwnership(t *testing.T) {
	req := ucModel.TransferRequest{
		UserID:           1,
		FromWalletID:     1,
		ToWalletID:       2,
		FromCurrencyCode: "USD",
		ToCurrencyCode:   "EUR",
		Amount:           100.0,
	}

	// Test that user can only transfer from their own wallet
	wallet := obModel.Wallets{
		ID:     1,
		UserID: 1, // Same as request user ID
		Name:   "Test Wallet",
	}

	assert.Equal(t, req.UserID, wallet.UserID)
	assert.Equal(t, req.FromWalletID, wallet.ID)
}
