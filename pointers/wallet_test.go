package pointers

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(10)

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Whthdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(10)
		assertNotError(t, err)
		assertBalance(t, wallet, want)

	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Error("wanted an error but didn't get one")

	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func assertNotError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
