package wallet

import (
	"errors"
	"fmt"
	"testing"
)

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})

}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

//The language permits us to write w.balance,
// without an explicit dereference.
// These pointers to structs even have their own name:
// struct pointers and they are automatically dereferenced.

func (w *Wallet) Deposit(val Bitcoin) {
	// fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += val
}
func (w *Wallet) Withdraw(val Bitcoin) error {
	// fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	if w.balance-val < 0 {
		return ErrInsufficientFunds
	}
	w.balance -= val
	return nil
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

/*
Pointers

-Go copies values when you pass them to functions/methods,
	so if you're writing a function that needs to mutate state
	you'll need it to take a pointer to the thing you want to change.
-The fact that Go takes a copy of values is useful a lot of the time
	but sometimes you won't want your system to make a copy of something, in which case you need to pass a reference.
	Examples include referencing very large data structures or things where only one instance is necessary (like database connection pools).

nil

-Pointers can be nil
-When a function returns a pointer to something,
	you need to make sure you check if it's nil
	or you might raise a runtime exception - the compiler won't help you here.
-Useful for when you want to describe a value that could be missing

Errors
-Errors are the way to signify failure when calling a function/method.
-By listening to our tests we concluded that checking for a string
	in an error would result in a flaky test. So we refactored our implementation
	to use a meaningful value instead and this resulted in easier to test code
 	and concluded this would be easier for users of our API too.
*/
