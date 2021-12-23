package fintech

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, w *Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if got.Error() != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, &wallet, Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		wallet.Withdraw(10)
		assertBalance(t, &wallet, Bitcoin(10))
	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(100)
		assertError(t, err, "insufficient funds")
		assertBalance(t, &wallet, startingBalance)
	})
}
