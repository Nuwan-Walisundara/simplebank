package db

import (
	"context"
	"database/sql"
	"fmt"
)

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	db *sql.DB
	*Queries
}
type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferParams) (TrnasferTxResult, error)
}

var txKey = struct{}{}

// NewStore creates a new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) executeTx(ctx context.Context, fn func(*Queries) error) error {
	tx, errr := store.db.BeginTx(ctx, nil)

	if errr != nil {
		return errr
	}
	q := New(tx)
	err := fn(q)

	if err != nil {
		if rberr := tx.Rollback(); rberr != nil {
			return fmt.Errorf("tx error: %v , rb error : %v", err, rberr)
		}
		return err
	}

	return tx.Commit()
}

//TransferParams contains the from account id , to account id and the transfering amount
type TransferParams struct {
	FromAccountId int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Balance       int64 `json:"ammount"`
}
type TrnasferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

/*
* TransferTx perfoms mony transfer from one account (fromAccount) to another account( toAccount)
* this create
*	1. transfer record
* 	2. add entries
*	3. Update from account balance and to account balance
* All the actions are handle in one database transaction
 */
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferParams) (TrnasferTxResult, error) {
	var tr TrnasferTxResult

	err := store.executeTx(ctx, func(q *Queries) error {
		var err error

		txName := ctx.Value(txKey)
		fmt.Println(txName, "create Transfer")

		tr.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountId,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Balance,
		})
		if err != nil {
			return err
		}

		fmt.Println(txName, "create Entry1")
		tr.FromEntry, err = q.CreateEntries(ctx, CreateEntriesParams{
			AccountID: arg.FromAccountId,
			Amount:    -arg.Balance,
		})
		if err != nil {
			return err
		}
		fmt.Println(txName, "create Entry2")
		tr.ToEntry, err = q.CreateEntries(ctx, CreateEntriesParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Balance,
		})
		if err != nil {
			return err
		}
		/*
		 * 1.get account 1
		 * 2. update account 1 balance
		 * 3.get account 2
		 * 4.update account2 balance
		 */
		if arg.FromAccountId > arg.ToAccountID {

			fmt.Println(txName, "get account 1 for Update   ")
			account1, err := q.GetAccountForUpdate(ctx, arg.FromAccountId)

			if err != nil {
				return err
			}
			fmt.Println(txName, "update account 1    ")
			bal := account1.Balance - arg.Balance
			tr.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
				ID:      account1.ID,
				Balance: bal,
			})
			fmt.Println(txName, "get account 2 for Update   ")
			account2, err := q.GetAccountForUpdate(ctx, arg.ToAccountID)
			if err != nil {
				return err
			}
			bal2 := account2.Balance + arg.Balance
			fmt.Println(txName, "update account 2    ")
			tr.ToAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
				ID:      account2.ID,
				Balance: bal2,
			})

		} else {

			fmt.Println(txName, "get account 2 for Update   ")
			account2, err := q.GetAccountForUpdate(ctx, arg.ToAccountID)
			if err != nil {
				return err
			}
			bal2 := account2.Balance + arg.Balance
			fmt.Println(txName, "update account 2    ")
			tr.ToAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
				ID:      account2.ID,
				Balance: bal2,
			})

			fmt.Println(txName, "get account 1 for Update   ")
			account1, err := q.GetAccountForUpdate(ctx, arg.FromAccountId)

			if err != nil {
				return err
			}
			fmt.Println(txName, "update account 1    ")
			bal := account1.Balance - arg.Balance
			tr.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
				ID:      account1.ID,
				Balance: bal,
			})

		}

		return nil
	})
	return tr, err
}

func oderdedTxTrnasfer(ctx context.Context, account1, account2 Account) {

}
