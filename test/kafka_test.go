/**
 *Created by He Haitao at 2019/12/19 10:27 上午
 */
package test

import (
	"testing"

	logger "github.com/tal-tech/loggerX"
	"github.com/tal-tech/xtools/flagutil"
	"github.com/tal-tech/xtools/kafkautil"
)

func BenchmarkKafkaFor10Byte(b *testing.B) {
	b.StopTimer()
	str := []byte("kafkakafka")
	b.Logf("Message len is %v", len(str))
	flagutil.SetConfig("/home/hehaitao/project/pan/conf/conf_for_kafka_test.ini")
	logger.InitLogger("/home/hehaitao/project/pan/conf/log.xml")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err := kafkautil.Send2Proxy("hehaitao_test", str)
		if err != nil {
			b.Log(err)
		}
	}
}

func BenchmarkKafkaFor100Byte(b *testing.B) {
	b.StopTimer()
	str := []byte("kafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafka")
	b.Logf("Message len is %v", len(str))
	flagutil.SetConfig("/home/hehaitao/project/pan/conf/conf_for_kafka_test.ini")
	logger.InitLogger("/home/hehaitao/project/pan/conf/log.xml")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err := kafkautil.Send2Proxy("hehaitao_test", str)
		if err != nil {
			b.Log(err)
		}
	}
}

func BenchmarkKafkaFor1000Byte(b *testing.B) {
	b.StopTimer()
	str := []byte("kafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafka")
	b.Logf("Message len is %v", len(str))
	flagutil.SetConfig("/home/hehaitao/project/pan/conf/conf_for_kafka_test.ini")
	logger.InitLogger("/home/hehaitao/project/pan/conf/log.xml")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err := kafkautil.Send2Proxy("hehaitao_test", str)
		if err != nil {
			b.Log(err)
		}
	}
}

func BenchmarkKafkaFor5000Byte(b *testing.B) {
	b.StopTimer()
	str := []byte("kafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafkakafka")
	b.Logf("Message len is %v", len(str))
	flagutil.SetConfig("/home/hehaitao/project/pan/conf/conf_for_kafka_test.ini")
	logger.InitLogger("/home/hehaitao/project/pan/conf/log.xml")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err := kafkautil.Send2Proxy("hehaitao_test", str)
		if err != nil {
			b.Log(err)
		}
	}
}
