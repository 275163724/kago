package kago

import (
	"github.com/Shopify/sarama"
	"log"
)

type AsyncProducer struct {
	producer        sarama.AsyncProducer
	Id              int
	ProducerGroupId string
}

func InitManualRetryAsyncProducer(addr []string, conf *Config) (*AsyncProducer, error) {
	conf.Config.Producer.Retry.Max = 0
	aSyncProducer := &AsyncProducer{
		Id:              0,
		ProducerGroupId: "",
	}
	var err error
	aSyncProducer.producer,err=sarama.NewAsyncProducer(addr,&conf.Config.Config)
	if err!=nil{
		log.Print(err)
		return nil, err
	}
	return aSyncProducer, nil
}

func InitManualRetryAsyncProducerGroup(addr []string, conf *Config, groupId string) ([]*AsyncProducer, error) {
	conf.Config.Producer.Retry.Max = 0
	producerAmount := conf.AsyncProducerAmount
	if producerAmount < 1 {
		producerAmount = 1
	}
	var producerSli []*AsyncProducer
	for i := 0; i < producerAmount; i++ {
		aSyncProducer := &AsyncProducer{
			Id:              i,
			ProducerGroupId: groupId,
		}
		var err error
		aSyncProducer.producer, err = sarama.NewAsyncProducer(addr, &conf.Config.Config)
		if err != nil {
			log.Print(err)
		} else {
			producerSli = append(producerSli, aSyncProducer)
		}
	}
	return producerSli, nil
}

func (asp *AsyncProducer) Send() (chan<- *ProducerMessage)  {

}

func (asp *AsyncProducer) Successes() (<-chan *ProducerMessage)  {

}

func (asp *AsyncProducer) Errors() (<-chan *ProducerError) {

}

func (asp *AsyncProducer) Close() (err error) {

}

