package rmq

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type TestQueue struct {
	name           string
	LastDeliveries []string
}

func NewTestQueue(name string) *TestQueue {
	queue := &TestQueue{name: name}
	queue.Reset()
	return queue
}

func (queue *TestQueue) String() string {
	return queue.name
}

func (queue *TestQueue) Publish(payload string, checkAlreadyExist ...bool) error {
	queue.LastDeliveries = append(queue.LastDeliveries, payload)
	return nil
}

func (queue *TestQueue) PublishBytes(payload ...[]byte) error {
	for _, b := range payload {
		err := queue.Publish(string(b))
		if err != nil {
			return err
		}
	}
	return nil
}

func (*TestQueue) SetPushQueue(Queue)                                   { panic(errorNotSupported) }
func (*TestQueue) StartConsuming(int64, time.Duration) error            { panic(errorNotSupported) }
func (*TestQueue) StopConsuming() <-chan struct{}                       { panic(errorNotSupported) }
func (*TestQueue) AddConsumer(string, Consumer) (string, error)         { panic(errorNotSupported) }
func (*TestQueue) AddConsumerFunc(string, ConsumerFunc) (string, error) { panic(errorNotSupported) }
func (*TestQueue) AddBatchConsumer(string, int64, time.Duration, BatchConsumer) (string, error) {
	panic(errorNotSupported)
}
func (*TestQueue) ReturnUnacked(int64) (int64, error)  { panic(errorNotSupported) }
func (*TestQueue) ReturnRejected(int64) (int64, error) { panic(errorNotSupported) }
func (*TestQueue) PurgeReady() (int64, error)          { panic(errorNotSupported) }
func (*TestQueue) PurgeRejected() (int64, error)       { panic(errorNotSupported) }
func (*TestQueue) GetRejected(int64) ([]string, error) { panic(errorNotSupported) }
func (*TestQueue) Destroy() (int64, int64, error)      { panic(errorNotSupported) }
func (*TestQueue) closeInStaleConnection() error       { panic(errorNotSupported) }
func (*TestQueue) readyCount() (int64, error)          { panic(errorNotSupported) }
func (*TestQueue) unackedCount() (int64, error)        { panic(errorNotSupported) }
func (*TestQueue) rejectedCount() (int64, error)       { panic(errorNotSupported) }
func (*TestQueue) getConsumers() ([]string, error)     { panic(errorNotSupported) }
func (*TestQueue) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	panic(errorNotSupported)
}
func (*TestQueue) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	panic(errorNotSupported)
}
func (*TestQueue) ScriptExists(ctx context.Context, hashes ...string) *redis.BoolSliceCmd {
	panic(errorNotSupported)
}
func (*TestQueue) ScriptLoad(ctx context.Context, script string) *redis.StringCmd {
	panic(errorNotSupported)
}

// test helper

func (queue *TestQueue) Reset() {
	queue.LastDeliveries = []string{}
}
