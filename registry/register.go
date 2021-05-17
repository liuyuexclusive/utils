package registry

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	SERVICE_KEY_SPLIT = "_"
)

type Service struct {
	name          string           //like prefix
	id            string           //id is a uuid
	cli           *clientv3.Client //etcd client
	leaseID       clientv3.LeaseID //leanse
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	val           string //value
}

// key
func (s *Service) key() string {
	return fmt.Sprintf("%s%s%s", s.name, SERVICE_KEY_SPLIT, s.id)
}

// NewService
func NewService(endpoints []string, name, val string, lease int64) (*Service, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	ser := &Service{
		cli:  cli,
		name: name,
		val:  val,
		id:   uuid.New().String(),
	}

	//get lease
	if err := ser.putKeyWithLease(lease); err != nil {
		return nil, err
	}

	go ser.listenLeaseRespChan()

	return ser, nil
}

//set lease
func (s *Service) putKeyWithLease(lease int64) error {
	//set time
	resp, err := s.cli.Grant(context.Background(), lease)
	if err != nil {
		return err
	}
	//set key,value with lease
	_, err = s.cli.Put(context.Background(), s.key(), s.val, clientv3.WithLease(resp.ID))
	if err != nil {
		return err
	}
	//keep alive
	leaseRespChan, err := s.cli.KeepAlive(context.Background(), resp.ID)

	if err != nil {
		return err
	}
	s.leaseID = resp.ID
	log.Println(s.leaseID)
	s.keepAliveChan = leaseRespChan
	log.Printf("Put key:%s  val:%s  success!", s.key(), s.val)
	return nil
}

//listenLeaseRespChan
func (s *Service) listenLeaseRespChan() {
	for _ = range s.keepAliveChan {
		// log.Printf("lease %v successed\n", leaseKeepResp)
	}
	log.Printf("lease %d closed\n", s.leaseID)
}

// Close
func (s *Service) Close() error {
	if _, err := s.cli.Revoke(context.Background(), s.leaseID); err != nil {
		return err
	}
	log.Printf("revoke lease %d\n", s.leaseID)
	return s.cli.Close()
}
