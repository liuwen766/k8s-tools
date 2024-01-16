package main

import (
	"os"
	"time"
)

const (
	BroadCasting MessageModel = iota
	Clustering

	ConsumeFromLastOffset ConsumeFromWhere = iota
	ConsumeFromFirstOffset
	ConsumeFromTimestamp
)

type Consumer struct {
	isStart    bool
	group      string
	nameServer *NamesrvAddr
}

func (c *Consumer) newRocketMqConsumer() (PushConsumer, error) {
	opts := []Option{
		WithGroupName(c.group),
		WithNameServer(c.nameServer),
		WithConsumerModel(Clustering),
		WithVIPChannel(true),
	}

	if os.Getenv("ConsumeFrom") == "FirstOffset" { //项目第一次启动，需激活消费组
		opts = append(opts, WithConsumeFromWhere(ConsumeFromFirstOffset))
	} else if os.Getenv("ConsumeFrom") == "Timestamp" {
		opts = append(opts, WithConsumeFromWhere(ConsumeFromTimestamp))
	}

	mc, err := NewPushConsumer(opts...)
	if err != nil {
		return nil, err
	}

	return mc, nil
}

type PushConsumer interface {
	// Start the PullConsumer for consuming message
	Start() error

	// Shutdown the PullConsumer, all offset of MessageQueue will be sync to broker before process exit
	Shutdown() error
	// Subscribe a topic for consuming
	Subscribe(topic string) error
	error

	// Unsubscribe a topic
	Unsubscribe(topic string) error
}

type Option func(*consumerOptions)

type consumerOptions struct {
	ClientOptions
	ConsumeTimestamp          string
	ConsumerPullTimeout       time.Duration
	ConsumerModel             MessageModel
	PullThresholdForTopic     int
	PullThresholdSizeForTopic int
	FromWhere                 ConsumeFromWhere
	// Message pull Interval
	PullInterval time.Duration
	// Batch consumption size
	ConsumeMessageBatchMaxSize int
	// Batch pull size
	PullBatchSize int32
	// Whether update subscription relationship when every pull
	PostSubscriptionWhenPull bool
	MaxReconsumeTimes        int32
	// Suspending pulling time for cases requiring slow pulling like flow-control scenario.
	SuspendCurrentQueueTimeMillis time.Duration
	// Maximum amount of time a message may block the consuming thread.
	ConsumeTimeout             time.Duration
	MaxTimeConsumeContinuously time.Duration
	//
	AutoCommit            bool
	RebalanceLockInterval time.Duration
}

func defaultPushConsumerOptions() consumerOptions {
	opts := consumerOptions{
		ClientOptions:              DefaultClientOptions(),
		MaxTimeConsumeContinuously: time.Duration(60 * time.Second),
		RebalanceLockInterval:      20 * time.Second,
		MaxReconsumeTimes:          -1,
		ConsumerModel:              Clustering,
		AutoCommit:                 true,
	}
	opts.ClientOptions.GroupName = "DEFAULT_CONSUMER"
	return opts
}

func DefaultClientOptions() ClientOptions {
	opts := ClientOptions{
		InstanceName: "DEFAULT",
		RetryTimes:   3,
	}
	return opts
}

type ClientOptions struct {
	GroupName         string
	NameServerAddrs   string
	Namesrv           *string
	ClientIP          string
	InstanceName      string
	UnitMode          bool
	UnitName          string
	VIPChannelEnabled bool
	RetryTimes        int
	Namespace         string
}

// WithGroupName set group name address
func WithGroupName(group string) Option {
	return func(opts *consumerOptions) {
		if group == "" {
			return
		}
		opts.GroupName = group
	}
}

// WithNameServer set NameServer address, only support one NameServer cluster in alpha2
func WithNameServer(nameServers *NamesrvAddr) Option {
	return func(options *consumerOptions) {
		//options.Resolver = primitive.NewPassthroughResolver(nameServers)
	}
}

type NamesrvAddr []string

type MessageModel int

func WithConsumerModel(m MessageModel) Option {
	return func(options *consumerOptions) {
		options.ConsumerModel = m
	}
}

func WithVIPChannel(enable bool) Option {
	return func(opts *consumerOptions) {
		opts.VIPChannelEnabled = enable
	}
}

type ConsumeFromWhere int

func WithConsumeFromWhere(w ConsumeFromWhere) Option {
	return func(options *consumerOptions) {
		options.FromWhere = w
	}
}
func NewPushConsumer(opts ...Option) (PushConsumer, error) {
	return NewPushConsumer(opts...)
}
