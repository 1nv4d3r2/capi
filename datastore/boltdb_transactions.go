package datastore

import (
	// Standard Library Imports
	"context"
	"fmt"

	// External Imports
	"github.com/asdine/storm"
	log "github.com/sirupsen/logrus"
	bolt "go.etcd.io/bbolt"

	// Internal Imports
	"github.com/aciddude/capi"
)

// Transactions provides a boltdb backed datastore implementation for storing
// blockchain transactions.
type Transactions struct {
	// dbpath is the path to where the boltdb datafiles are stored.
	DBPath string

	// DB contains the boltdb handler to access the underlying buckets.
	DB *bolt.DB
	// storm is a wrapper around BoltDB that provides higher level ORM based
	// operations used internally to easily perform queries.
	storm *storm.DB
	// Coins contains the list of coins supported by the underlying DB.
	Coins []string
	// Options are specific boltdb configurations.
	Options *bolt.Options
}

// CreateSchema implements datastore.Creater
func (b *Transactions) CreateSchema(ctx context.Context, coins []string) (err error) {
	b.DB, err = bolt.Open(b.DBPath, 0666, b.Options)
	if err != nil {
		return err
	}

	b.storm, err = storm.Open(b.DBPath, storm.UseDB(b.DB))
	if err != nil {
		return err
	}

	for _, coin := range coins {
		err := b.DB.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(coin))
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
			return nil
		})

		if err != nil {
			return err
		}
	}

	return nil
}

// Configure implements Configurer.Configure
func (b *Transactions) Configure(ctx context.Context, coins []string) (err error) {
	if b.DB == nil {
		b.DB, err = bolt.Open(b.DBPath, 0666, b.Options)
		if err != nil {
			return err
		}
	}

	if b.storm == nil {
		b.storm, err = storm.Open(b.DBPath, storm.UseDB(b.DB))
		if err != nil {
			return err
		}
	}

	// Update database buckets on first up to ensure any added coins are
	// bucketed correctly.
	err = b.DB.Update(func(tx *bolt.Tx) error {
		for _, coin := range coins {
			_, err := tx.CreateBucketIfNotExists([]byte(coin))
			if err != nil {
				return fmt.Errorf("error creating bucket: %s", err)
			}
		}
		return nil
	})

	return err
}

// IsCreated returns a bool to trigger crea
func (b *Transactions) IsCreated(ctx context.Context, coins []string) bool {
	logger := log.WithFields(log.Fields{
		"package": "datastore",
		"manager": "Transactions",
		"method":  "IsCreated",
	})

	dbpath := fmt.Sprintf("%s/%s", b.DBPath, boltdbPathTransactions)
	ok, err := fileExists(dbpath)
	if err != nil {
		logger.WithError(err).Debug("error checking for database")
	}

	b.DBPath = dbpath
	return ok
}

// Close implements stater.Close.
func (t *Transactions) Close() error {
	return t.DB.Close()
}

// List enables filtering for transactions in the datastore.
func (b *Transactions) List(ctx context.Context, coin string, query ListTransactionsRequest) (transactions []*capi.Transaction, err error) {
	// TODO: Implement logic

	return nil, nil
}

// Create creates a transaction in the datastore.
func (b *Transactions) Create(ctx context.Context, coin string, query CreateTransactionRequest) (transaction *capi.Transaction, err error) {
	// TODO: Implement logic

	return nil, nil
}

// CreateBulk enables bulk creation of transactions in the datastore.
func (b *Transactions) CreateBulk(ctx context.Context, coin string, query []CreateTransactionRequest) (transactions []*capi.Transaction, err error) {
	// TODO: Implement logic

	return nil, nil
}

// Update enables updating a transaction resource in the datastore.
func (b *Transactions) Update(ctx context.Context, coin string, query UpdateTransactionRequest) (transaction *capi.Transaction, err error) {
	// TODO: Implement logic

	return nil, nil
}

// Delete removes a transaction from the datastore.
func (b *Transactions) Delete(ctx context.Context, coin string, query DeleteTransactionRequest) (err error) {
	// TODO: Implement logic

	return nil
}
