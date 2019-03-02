package main

import (
	"context"

	"cloud.google.com/go/datastore"
)

type DatastoreStore struct {
	ds *datastore.Client
}

func NewDatastoreStore(ds *datastore.Client) *DatastoreStore {
	return &DatastoreStore{
		ds: ds,
	}
}

func CreateClient(ctx context.Context, projectID string) (*datastore.Client, error) {
	return datastore.NewClient(ctx, projectID)
}

func (ds *DatastoreStore) QueryKeysOnly(ctx context.Context) error {
	var s struct{}

	q := datastore.NewQuery("SpannerQueryStats").KeysOnly().Limit(1)
	iter := ds.ds.Run(ctx, q)
	_, err := iter.Next(&s)
	if err != nil {
		if err == datastore.ErrNoSuchEntity {
			return nil
		}
		return err
	}
	return nil
}
