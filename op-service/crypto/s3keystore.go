package crypto

import (
	"context"
	"errors"
	"math/big"
	"os"

	s3store "github.com/zircuit-labs/zkr-go-common/stores/s3"
	"github.com/zircuit-labs/zkr-go-common/xerrors/stacktrace"

	"github.com/zircuit-labs/l2-geth-public/accounts"
	"github.com/zircuit-labs/l2-geth-public/accounts/keystore"
	"github.com/zircuit-labs/l2-geth-public/core/types"
)

var (
	ErrNoFile       = errors.New("no keystore file supplied")
	ErrNoPassphrase = errors.New("no passphrase supplied")
)

// Keystore directory
var defaultKeystoreDir = ".tmpkeystore"

type S3KeystoreConfig struct {
	Passphrase  string `koanf:"passphrase"`
	S3File      string `koanf:"file"`
	KeystoreDir string `koanf:"keystoredir"`

	BlobStoreConfig s3store.BlobStoreConfig `koanf:"blobstore"`
}

type S3KeystoreService struct {
	config    S3KeystoreConfig
	blobStore *s3store.BlobStore
	ks        *keystore.KeyStore
	account   accounts.Account
}

func NewS3Keystore(config S3KeystoreConfig) (*S3KeystoreService, error) {
	if config.S3File == "" {
		return nil, ErrNoFile
	}
	if config.Passphrase == "" {
		return nil, ErrNoPassphrase
	}

	var keystoreDir string
	if config.KeystoreDir != "" {
		keystoreDir = config.KeystoreDir
	} else {
		tmpDir, err := os.MkdirTemp("", defaultKeystoreDir)
		if err != nil {
			return nil, stacktrace.Wrap(err)
		}
		keystoreDir = tmpDir
	}

	blobStore, err := s3store.NewBlobStoreFromConfig(config.BlobStoreConfig)
	if err != nil {
		return nil, err
	}

	keystoreFile, err := blobStore.Get(context.Background(), config.S3File)
	if err != nil {
		return nil, err
	}

	// Empty out directory
	err = os.RemoveAll(keystoreDir)
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}

	ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.Import(keystoreFile, config.Passphrase, config.Passphrase)
	if err != nil {
		return nil, stacktrace.Wrap(err)
	}

	return &S3KeystoreService{
		config:    config,
		blobStore: blobStore,
		ks:        ks,
		account:   account,
	}, nil
}

func (s *S3KeystoreService) SignTx(tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	return s.ks.SignTxWithPassphrase(s.account, s.config.Passphrase, tx, chainID)
}

// expose account
func (s *S3KeystoreService) Account() accounts.Account {
	return s.account
}
