package chengsheng

import (
	"golang.org/x/sync/errgroup"
	"time"
)


func (s *Service) SpiderConsumer(session *kafka.ConsumerSession, msgs <-chan *sarama.ConsumerMessage) error {
	s.wg.Add(1)
	defer s.wg.Done()

	fun := s.BatchKafkaMsg(session, msgs, 100)
	for {
		batchMsg, ok := fun()
		if !ok {
			xlog.Warnf("[SpiderConsumer] get empty kafka data")
			return nil
		}

		eg := errgroup.Group{}
		eg.GOMAXPROCS(20)
		for _, msg := range messages {
			tmp := msg
			eg.Go(func(ctx context.Context) error {
				// todo 业务
				return nil
			})
		}
		err := eg.Wait()
		if err != nil {
			xlog.Errorf("[SpiderConsumer] eg.wait err:%+v", err)
		}
	}
}

func (s *Service) BatchKafkaMsg(session *kafka.ConsumerSession, msgs <-chan *sarama.ConsumerMessage, num int) func() ([]*sarama.ConsumerMessage, bool) {
	var lastMsgSpider *sarama.ConsumerMessage

	return func() ([]*sarama.ConsumerMessage, bool) {
		msg := make([]*sarama.ConsumerMessage, 0)
		ticker := time.NewTicker(5 * time.Second)
		defer func() {
			ticker.Stop()
		}()
		if lastMsgSpider != nil {
			session.Commit(lastMsgSpider)
			lastMsgSpider = nil
		}
		for {
			select {
			case tempMsg, ok := <-msgs:
				if !ok {
					return msg, ok
				}
				lastMsgSpider = tempMsg
				msg = append(msg, tempMsg)
				if len(msg) >= num {
					return msg, true
				}
			case <-ticker.C:
				if len(msg) > 0 {
					return msg, true
				}
			}
		}
	}
}
