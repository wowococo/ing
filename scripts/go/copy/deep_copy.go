package main

import (
	"time"
)

type Offset int64
type TimestampType int

type Header struct {
	Key   string
	Value []byte
}

type TopicPartition struct {
	Topic       *string
	Partition   int32
	Offset      Offset
	Metadata    *string
	Error       error
	LeaderEpoch *int32
}

type Message struct {
	TopicPartition TopicPartition
	Value          []byte
	Key            []byte
	Timestamp      time.Time
	TimestampType  TimestampType
	Opuque         interface{}
	Headers        []Header
}

func fake_deepcopy(msg *Message) *Message {
	// 深拷贝 Message 结构体
	copiedMsg := &Message{
		TopicPartition: TopicPartition{
			Topic:       msg.TopicPartition.Topic,
			Partition:   msg.TopicPartition.Partition,
			Offset:      msg.TopicPartition.Offset,
			Metadata:    msg.TopicPartition.Metadata,
			Error:       msg.TopicPartition.Error,
			LeaderEpoch: msg.TopicPartition.LeaderEpoch,
		},
		Value:         msg.Value,
		Key:           msg.Key,
		Timestamp:     msg.Timestamp,
		TimestampType: msg.TimestampType,
		Opuque:        msg.Opuque,
		Headers:       msg.Headers,
	}

	return copiedMsg
}

// ... existing code ...

func DeepCopyMessage(src *Message) *Message {
	if src == nil {
		return nil
	}

	dst := &Message{
		TopicPartition: DeepCopyTopicPartition(src.TopicPartition),
		Value:          make([]byte, len(src.Value)),
		Key:            make([]byte, len(src.Key)),
		Timestamp:      src.Timestamp,
		TimestampType:  src.TimestampType,
		Opuque:         src.Opuque,
		Headers:        make([]Header, len(src.Headers)),
	}

	copy(dst.Value, src.Value)
	copy(dst.Key, src.Key)

	// Deep copy Headers
	for i := range src.Headers {
		dst.Headers[i].Key = src.Headers[i].Key
		dst.Headers[i].Value = make([]byte, len(src.Headers[i].Value))
		copy(dst.Headers[i].Value, src.Headers[i].Value)
	}

	return dst
}

// ... existing code ...

// 首先要新建对象
// 然后深拷贝引用类型，比如指针、切片、map、管道
func DeepCopyTopicPartition(src TopicPartition) TopicPartition {
	dst := TopicPartition{
		Partition: src.Partition,
		Offset:    src.Offset,
		Error:     src.Error,
	}

	// 深拷贝字符串指针
	if src.Topic != nil {
		// 值拷贝
		copiedTopic := *src.Topic
		dst.Topic = &copiedTopic
	}

	// 深拷贝 Metadata 指针
	if src.Metadata != nil {
		copiedMetadata := *src.Metadata
		dst.Metadata = &copiedMetadata
	}

	// 深拷贝 LeaderEpoch 指针
	if src.LeaderEpoch != nil {
		copiedLeaderEpoch := *src.LeaderEpoch
		dst.LeaderEpoch = &copiedLeaderEpoch
	}

	return dst
}

func main() {
	topicName := "original-topic"
	original := &Message{
		TopicPartition: TopicPartition{
			Topic:     &topicName,
			Partition: 0,
		},
		// 其他字段初始化...
	}

	copied := DeepCopyMessage(original)
	*copied.TopicPartition.Topic = "modified-topic" // 不会影响original的值
}
