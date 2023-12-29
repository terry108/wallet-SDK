package starknet

import (
	"github.com/NethermindEth/juno/core/felt"
	"github.com/coming-chat/wallet-SDK/core/base"
	"github.com/dontpanicdao/caigo/types"
	"github.com/xiang-xx/starknet.go/rpc"
)

type Transaction struct {
	calls   []types.FunctionCall
	details types.ExecuteDetails

	txnV1   rpc.InvokeTxnV1
	txnHash *felt.Felt
}

type SignedTransaction struct {
	// depoly Txn
	depolyTxn *rpc.DeployAccountTxn

	// Do you need to automatically deploy the contract address first when you send the transaction for the first time? default NO
	NeedAutoDeploy bool
	// invoke Txn
	invokeTxn *rpc.InvokeTxnV1
	Account   *Account
}

func (t *Transaction) SignWithAccount(account base.Account) (signedTxn *base.OptionalString, err error) {
	return nil, base.ErrUnsupportedFunction
}

func (t *Transaction) SignedTransactionWithAccount(account base.Account) (signedTx base.SignedTransaction, err error) {
	starknetAccount := AsStarknetAccount(account)
	if starknetAccount == nil {
		return nil, base.ErrInvalidAccountType
	}
	err = t.Sign(starknetAccount)
	if err != nil {
		return
	}
	return &SignedTransaction{
		Account:   starknetAccount,
		invokeTxn: &t.txnV1,
	}, nil
}

func (t *Transaction) Sign(acc *Account) error {
	if t.txnHash == nil {
		return base.ErrInvalidTransactionType
	}
	s1, s2, err := acc.SignHash(t.txnHash)
	if err != nil {
		return err
	}
	t.txnV1.Signature = []*felt.Felt{s1, s2}
	return nil
}

func (txn *SignedTransaction) HexString() (res *base.OptionalString, err error) {
	return nil, base.ErrUnsupportedFunction
}

func AsSignedTransaction(txn base.SignedTransaction) *SignedTransaction {
	if res, ok := txn.(*SignedTransaction); ok {
		return res
	}
	return nil
}
