package handler

import (
	"context"
	"fmt"
	"time"

	"go-micro.dev/v4/cache"
	log "go-micro.dev/v4/logger"

	pb "github.com/go-micro/examples/cache/proto"
)

type Cache struct {
	cache cache.Cache
}

func NewCache(opts ...cache.Option) *Cache {
	c := cache.NewCache(opts...)
	return &Cache{c}
}

func (c *Cache) Get(ctx context.Context, req *pb.GetRequest, rsp *pb.GetResponse) error {
	log.Infof("Received Cache.Get request: %v", req)

	v, e, err := c.cache.Get(ctx, req.Key)
	if err != nil {
		return err
	}

	rsp.Value = fmt.Sprintf("%v", v)
	rsp.Expiration = e.String()

	return nil
}

func (c *Cache) Put(ctx context.Context, req *pb.PutRequest, rsp *pb.PutResponse) error {
	log.Infof("Received Cache.Put request: %v", req)

	d, err := time.ParseDuration(req.Duration)
	if err != nil {
		return err
	}

	if err := c.cache.Put(ctx, req.Key, req.Value, d); err != nil {
		return err
	}

	return nil
}

func (c *Cache) Delete(ctx context.Context, req *pb.DeleteRequest, rsp *pb.DeleteResponse) error {
	log.Infof("Received Cache.Delete request: %v", req)

	if err := c.cache.Delete(ctx, req.Key); err != nil {
		return err
	}

	return nil
}
