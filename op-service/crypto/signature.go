package crypto

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum/accounts"
	"github.com/zircuit-labs/l2-geth-public/accounts/abi/bind"
	"github.com/zircuit-labs/l2-geth-public/common"
	"github.com/zircuit-labs/l2-geth-public/core/types"
	"github.com/zircuit-labs/l2-geth-public/crypto"
	"github.com/zircuit-labs/l2-geth-public/log"

	hdwallet "github.com/ethereum-optimism/go-ethereum-hdwallet"
)

type SignerType string

const (
	// Mnemonic is a signer that uses a mnemonic to derive the private key
	Mnemonic SignerType = "mnemonic"

	// S3Keystore is a signer that uses a keystore file to sign transactions
	S3Keystore SignerType = "s3keystore"
)

func ParseSignerType(signerType string) (SignerType, error) {
	switch SignerType(signerType) {
	case Mnemonic:
		return Mnemonic, nil
	case S3Keystore:
		return S3Keystore, nil
	default:
		return "", fmt.Errorf("invalid signer type: %s", signerType)
	}
}

func PrivateKeySignerFn(key *ecdsa.PrivateKey, chainID *big.Int) bind.SignerFn {
	from := crypto.PubkeyToAddress(key.PublicKey)
	signer := types.LatestSignerForChainID(chainID)
	return func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if address != from {
			return nil, bind.ErrNotAuthorized
		}
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), key)
		if err != nil {
			return nil, err
		}
		return tx.WithSignature(signer, signature)
	}
}

// SignerFn is a generic transaction signing function. It may be a remote signer so it takes a context.
// It also takes the address that should be used to sign the transaction with.
type SignerFn func(context.Context, common.Address, *types.Transaction) (*types.Transaction, error)

// SignerFactory creates a SignerFn that is bound to a specific ChainID
type SignerFactory func(chainID *big.Int) SignerFn

// SignerFactoryFromConfig considers three ways that signers are created & then creates single factory from those config options.
// It can either take a remote signer (via opsigner.CLIConfig) or it can be provided either a mnemonic + derivation path or a private key.
// It prefers the remote signer, then the mnemonic or private key (only one of which can be provided).
func SignerFactoryFromConfig(
	l log.Logger,
	privateKey, mnemonic, hdPath string,
	signerType SignerType,
	keystoreConfig S3KeystoreConfig,
) (SignerFactory, common.Address, error) {
	var signer SignerFactory
	var fromAddress common.Address
	switch signerType {
	case S3Keystore:
		ks, err := NewS3Keystore(keystoreConfig)
		if err != nil {
			return nil, common.Address{}, fmt.Errorf("failed to create the s3keystore: %w", err)
		}

		fromAddress = ks.Account().Address
		signer = func(chainID *big.Int) SignerFn {
			return func(ctx context.Context, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
				return ks.SignTx(tx, chainID)
			}
		}
	default:
		var privKey *ecdsa.PrivateKey
		var err error

		if privateKey != "" && mnemonic != "" {
			return nil, common.Address{}, errors.New("cannot specify both a private key and a mnemonic")
		}
		if privateKey == "" {
			// Parse l2output wallet private key and L2OO contract address.
			wallet, err := hdwallet.NewFromMnemonic(mnemonic)
			if err != nil {
				return nil, common.Address{}, fmt.Errorf("failed to parse mnemonic: %w", err)
			}

			privKey, err = wallet.PrivateKey(ethereum.Account{
				URL: ethereum.URL{
					Path: hdPath,
				},
			})
			if err != nil {
				return nil, common.Address{}, fmt.Errorf("failed to create a wallet: %w", err)
			}
		} else {
			privKey, err = crypto.HexToECDSA(strings.TrimPrefix(privateKey, "0x"))
			if err != nil {
				return nil, common.Address{}, fmt.Errorf("failed to parse the private key: %w", err)
			}
		}
		fromAddress = crypto.PubkeyToAddress(privKey.PublicKey)
		signer = func(chainID *big.Int) SignerFn {
			s := PrivateKeySignerFn(privKey, chainID)
			return func(_ context.Context, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
				return s(addr, tx)
			}
		}
	}

	return signer, fromAddress, nil
}
