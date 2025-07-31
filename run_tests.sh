#!/bin/bash

# VocaGame E-Wallet Test Runner
echo "🚀 Running VocaGame E-Wallet Test Suite"
echo "======================================"

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    if [ $1 -eq 0 ]; then
        echo -e "${GREEN}✅ $2 PASSED${NC}"
    else
        echo -e "${RED}❌ $2 FAILED${NC}"
        return 1
    fi
}

# Test counter
total_tests=0
passed_tests=0

echo ""
echo "📋 Running Unit Tests..."
echo "----------------------"

# Run deposit tests
echo "🔸 Testing Deposit Logic..."
go test ./internal/usecase/wallet/deposit_test.go -v
deposit_result=$?
total_tests=$((total_tests + 1))
if print_status $deposit_result "Deposit Unit Tests"; then
    passed_tests=$((passed_tests + 1))
fi

echo ""

# Run transfer tests  
echo "🔸 Testing Transfer Logic..."
go test ./internal/usecase/wallet/transfer_test.go -v
transfer_result=$?
total_tests=$((total_tests + 1))
if print_status $transfer_result "Transfer Unit Tests"; then
    passed_tests=$((passed_tests + 1))
fi

echo ""

# Run all wallet usecase tests together
echo "🔸 Testing All Wallet Use Cases..."
go test ./internal/usecase/wallet/... -v
wallet_result=$?
total_tests=$((total_tests + 1))
if print_status $wallet_result "Wallet Use Case Tests"; then
    passed_tests=$((passed_tests + 1))
fi

echo ""
echo "🌐 Running Integration Tests..."
echo "----------------------------"

# Run integration tests
echo "🔸 Testing API Integration..."
go test ./test/integration/api_test.go -v
integration_result=$?
total_tests=$((total_tests + 1))
if print_status $integration_result "Integration Tests"; then
    passed_tests=$((passed_tests + 1))
fi

echo ""
echo "📊 Test Summary"
echo "==============="
echo "Total Test Suites: $total_tests"
echo "Passed: $passed_tests"
echo "Failed: $((total_tests - passed_tests))"

if [ $passed_tests -eq $total_tests ]; then
    echo -e "${GREEN}🎉 All tests passed! The e-wallet system is ready.${NC}"
    exit 0
else
    echo -e "${RED}⚠️  Some tests failed. Please check the output above.${NC}"
    exit 1
fi
