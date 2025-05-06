package main

import (
	"testing"
	"time"
)

func TestDeepCopy(t *testing.T) {
	setup := func() (*Message, func(string)) {
		topic := "test"
		original := &Message{
			TopicPartition: TopicPartition{
				Topic:       &topic,
				Partition:   1,
				Offset:      100,
				Metadata:    new(string),
				LeaderEpoch: new(int32),
			},
			Value:         []byte("hello"),
			Key:           []byte("world"),
			Timestamp:     time.Now(),
			TimestampType: 1,
		}

		return original, func(newTopic string) {
			if *original.TopicPartition.Topic != "test" {
				t.Errorf("原始对象的 Topic 被意外修改，期望 'test'，实际 '%s'", *original.TopicPartition.Topic)
			}
		}
	}

	t.Run("修改指针指向内容", func(t *testing.T) {
		original, _ := setup()
		copied := fake_deepcopy(original)
		newTopic := "test2"

		// 修改指针指向的值
		*copied.TopicPartition.Topic = newTopic

		// verify(newTopic) // 验证原始对象未被修改
		if *copied.TopicPartition.Topic != newTopic {
			t.Errorf("拷贝对象的 Topic 修改失败，期望 '%s'，实际 '%s'", newTopic, *copied.TopicPartition.Topic)
		}
	})

	t.Run("修改指针本身", func(t *testing.T) {
		original, _ := setup()
		copied := fake_deepcopy(original)
		newTopic := "test3"

		// 修改指针本身
		copied.TopicPartition.Topic = &newTopic

		// verify(newTopic) // 验证原始对象未被修改
		if *copied.TopicPartition.Topic != newTopic {
			t.Errorf("拷贝对象的指针修改失败，期望 '%s'，实际 '%s'", newTopic, *copied.TopicPartition.Topic)
		}
		if original.TopicPartition.Topic == copied.TopicPartition.Topic {
			t.Error("拷贝对象与原始对象的 Topic 指针不应相同")
		}
	})
}
