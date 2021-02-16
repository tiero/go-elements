package mobile

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil"
	"github.com/vulpemventures/go-elements/confidential"
	"github.com/vulpemventures/go-elements/network"
	"github.com/vulpemventures/go-elements/payment"
	"github.com/vulpemventures/go-elements/transaction"
)

// PrivateKey implements an Identity interface and holds a single EC private key
type PrivateKey struct {
	network     *network.Network
	signingKey  *btcec.PrivateKey
	blindingKey *btcec.PrivateKey
	script      []byte

	address AddressExtended
	payment *payment.Payment
}

// NewPrivateKeyFromWIF returns a PrivateKey with given signing and blinding wifs
func NewPrivateKeyFromWIF(chain, signingWIF, blindingWIF string) (Identity, error) {
	net := &network.Liquid
	if chain == "regtest" {
		net = &network.Regtest
	}

	sign, err := btcutil.DecodeWIF(signingWIF)
	if err != nil {
		return nil, err
	}
	blind, err := btcutil.DecodeWIF(signingWIF)
	if err != nil {
		return nil, err
	}

	return newPrivateKey(net, sign.PrivKey, blind.PrivKey)
}

//NewPrivateKeyRandom returns a PrivateKey randomly generated
func NewPrivateKeyRandom(chain string) (Identity, error) {
	net := &network.Liquid
	if chain == "regtest" {
		net = &network.Regtest
	}

	signKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, err
	}

	blindKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, err
	}

	return newPrivateKey(net, signKey, blindKey)
}

func newPrivateKey(
	net *network.Network,
	signKey, blindKey *btcec.PrivateKey,
) (*PrivateKey, error) {

	pay := payment.FromPublicKey(signKey.PubKey(), net, blindKey.PubKey())

	addr, err := pay.ConfidentialWitnessPubKeyHash()
	if err != nil {
		return nil, err
	}

	return &PrivateKey{
		network:     net,
		signingKey:  signKey,
		blindingKey: blindKey,
		script:      pay.Script,
		payment:     pay,
		address: AddressExtended{
			ConfidentialAddress: addr,
			BlindingPrivateKey:  blindKey.Serialize(),
		},
	}, nil
}


// IsAbleToSign returns true if Identity has key signing material
func (pk *PrivateKey) IsAbleToSign() bool {
	return true
}


// AddressByPath ...
func (pk *PrivateKey) AddressByPath(account, change, index int64) AddressExtended {
	return pk.address
}

// NextAddress returns the confidential segwit address of the identity
func (pk *PrivateKey) NextAddress(account int64) AddressExtended {
	return pk.address
}

// NextChangeAddress returns the confidential segwit address of the identity
func (pk *PrivateKey) NextChangeAddress(account int64) AddressExtended {
	return pk.address
}

// UnblindOutput returns explicit value and asset along with value blinding factors revealed
func (pk *PrivateKey) UnblindOutput(prevout *transaction.TxOutput) (*confidential.UnblindOutputResult, error) {
	return nil, nil
}

// Blind uses the blinding key to produce a blinded pset
func (pk *PrivateKey) Blind(pset string) (string, error) {
	return pset, nil
}

// Sign uses the signing keys to produce a signed pset
func (pk *PrivateKey) Sign(pset string) string, error {
	return pset, nil
}
