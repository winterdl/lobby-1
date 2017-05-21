package rpc

import (
	"io"
	"log"
	"os"

	"github.com/asdine/lobby"
	"github.com/asdine/lobby/json"
	"github.com/asdine/lobby/rpc/proto"
	"github.com/asdine/lobby/validation"
	"golang.org/x/net/context"
)

func newBucketService(r lobby.Registry) *bucketService {
	return &bucketService{
		registry: r,
		logger:   log.New(os.Stderr, "", log.LstdFlags),
	}
}

type bucketService struct {
	registry lobby.Registry
	logger   *log.Logger
}

// Put an item in the bucket.
func (s *bucketService) Put(stream proto.BucketService_PutServer) error {
	var itemCount int32

	for {
		newItem, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.PutSummary{
				ItemCount: itemCount,
			})
		}
		if err != nil {
			return Error(err, s.logger)
		}

		itemCount++
		err = validation.Validate(newItem)
		if err != nil {
			return Error(err, s.logger)
		}

		b, err := s.registry.Bucket(newItem.Bucket)
		if err != nil {
			return Error(err, s.logger)
		}

		data := json.ToValidJSONFromBytes(newItem.Item.Value)
		_, err = b.Put(newItem.Item.Key, data)
		if err != nil {
			return Error(err, s.logger)
		}
	}
}

func (s *bucketService) Get(ctx context.Context, key *proto.Key) (*proto.Item, error) {
	return nil, nil
}

func (s *bucketService) Delete(ctx context.Context, key *proto.Key) (*proto.Empty, error) {
	return nil, nil
}

func (s *bucketService) List(page *proto.Page, stream proto.BucketService_ListServer) error {
	return nil
}
