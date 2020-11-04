package pc

import (
	"../util"
	"fmt"
	"log"
	"reflect"
	"time"
)

type Test struct{}

var retrySleepTime int64 = 2

func (t *Test) HD0001(p util.Protocol, body []byte) {
	log.Println("handshake")
	// 断连重试，重试间隔上次一倍
	if p.StatusCode != 0 {
		time.Sleep(time.Duration(retrySleepTime) * time.Second)
		retrySleepTime *= 2
		handShake()
	} else {
		retrySleepTime = 2
	}
}

func (t *Test) HD0002(p util.Protocol, body []byte) {

}

func callReflect(any interface{}, name string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	if v := reflect.ValueOf(any).MethodByName(name); v.String() == "<invalid Value>" {
		return nil
	} else {
		return v.Call(inputs)
	}
}

const HandlerPrefix = "HD"

func handle(p util.Protocol, body []byte) {
	methodName := HandlerPrefix + fmt.Sprintf("%04x", p.Command)
	callReflect(&Test{}, methodName, p, body)

}
