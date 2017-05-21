package bolt

import (
	"github.com/asdine/lobby"
	"github.com/asdine/lobby/bolt/internal"
	"github.com/asdine/storm"
	"github.com/pkg/errors"
)

var _ lobby.Bucket = new(Bucket)

// NewBucket returns a Bucket
func NewBucket(node storm.Node) *Bucket {
	return &Bucket{
		node: node,
	}
}

// Bucket is a BoltDB implementation of a bucket.
type Bucket struct {
	node storm.Node
}

// Put value to the bucket. Returns an Item.
func (b *Bucket) Put(key string, value []byte) (*lobby.Item, error) {
	var i internal.Item

	tx, err := b.node.Begin(true)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create transaction")
	}
	defer tx.Rollback()

	err = tx.One("Key", key, &i)
	if err != nil {
		if err != storm.ErrNotFound {
			return nil, errors.Wrap(err, "failed to fetch item")
		}

		i = internal.Item{
			Key:   key,
			Value: value,
		}
	} else {
		i.Value = value
	}

	err = tx.Save(&i)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, "failed to commit")
	}

	return &lobby.Item{
		Key:   i.Key,
		Value: i.Value,
	}, nil
}

// Get an item by key.
func (b *Bucket) Get(key string) (*lobby.Item, error) {
	var i internal.Item

	err := b.node.One("Key", key, &i)
	if err != nil {
		if err == storm.ErrNotFound {
			return nil, lobby.ErrKeyNotFound
		}

		return nil, errors.Wrap(err, "failed to fetch item")
	}

	return &lobby.Item{
		Key:   i.Key,
		Value: i.Value,
	}, nil
}

// Delete item from the bucket
func (b *Bucket) Delete(key string) error {
	var i internal.Item

	tx, err := b.node.Begin(true)
	if err != nil {
		return errors.Wrap(err, "failed to create transaction")
	}
	defer tx.Rollback()

	err = tx.One("Key", key, &i)
	if err != nil {
		if err == storm.ErrNotFound {
			return lobby.ErrKeyNotFound
		}
		return errors.Wrap(err, "failed to fetch item")
	}

	err = tx.DeleteStruct(&i)
	if err != nil {
		return errors.Wrap(err, "failed to delete item")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "failed to commit")
	}

	return nil
}

// Page returns a list of items
func (b *Bucket) Page(page int, perPage int) ([]lobby.Item, error) {
	var skip int
	var list []internal.Item

	if page <= 0 {
		return nil, nil
	}

	if perPage >= 0 {
		skip = (page - 1) * perPage
	}

	err := b.node.All(&list, storm.Skip(skip), storm.Limit(perPage))
	if err != nil {
		return nil, errors.Wrap(err, "boltdb.bucket.Page failed to fetch items")
	}

	items := make([]lobby.Item, len(list))
	for i := range list {
		items[i].Key = list[i].Key
		items[i].Value = list[i].Value
	}
	return items, nil
}

// Close the bucket session
func (b *Bucket) Close() error {
	return nil
}