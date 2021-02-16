package mobile

import (
	"github.com/vulpemventures/go-elements/confidential"
	"github.com/vulpemventures/go-elements/transaction"
)

const (
	// PrivateKeyType ...
	PrivateKeyType = iota
	//MnemonicType ...
	MnemonicType
	// MasterPublicKeyType ...
	MasterPublicKeyType
)

// AddressExtended ...
type AddressExtended struct {
	ConfidentialAddress string
	BlindingPrivateKey  []byte
}

// Identity inteface defines the method for implementing an identity
type Identity interface {
	AddressByPath(account, change, index int64) AddressExtended
	NextAddress(account int64) AddressExtended
	NextChangeAddress(account int64) AddressExtended
	UnblindOutput(prevout *transaction.TxOutput) (*confidential.UnblindOutputResult, error)
	Blind(pset string) (string, error)
	Sign(pset string) (string, error)
	IsAbleToSign() bool
}
