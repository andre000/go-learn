package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("shoud be able make deposits w/ Bitcoins and return the balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("shoud be able to withdraw Bitcoins", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("should throw an error when trying to withdraw more than the current balance", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assert(t *testing.T, received Bitcoin, expected Bitcoin) {
	t.Helper()
	if received != expected {
		t.Errorf("❌ received %s expected %s", received, expected)
	}
}

func assertBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()
	received := wallet.Balance()
	assert(t, received, expected)
}

func assertError(t *testing.T, err error, expected error) {
	t.Helper()
	if err == nil {
		t.Fatal("❌ expected error but didn't get any")
	}

	if err != expected {
		t.Errorf("❌ expected %q received %q", expected, err)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("❌ expected nil received error")
	}
}
