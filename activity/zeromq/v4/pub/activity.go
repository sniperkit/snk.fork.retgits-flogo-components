package pub

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"

	zmq "github.com/alecthomas/gozmq"
	zmq "github.com/pebbe/zmq4"
)

const (
	ivService = "service"
	ivURI     = "uri"
	ivTopic   = "topic"
	ivMessage = "message"
	ovoutput  = "output"
)

// log is the default package logger
var log = logger.GetLogger("activity-tibco-zmqpub-v4")
var subsExpected = 1

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	syncServiceHost := context.GetInput(ivService).(string)
	URI := context.GetInput(ivURI).(string)
	Topic := context.GetInput(ivTopic).(string)
	Message := context.GetInput(ivMessage).(string)

	Publisher, _ := context.NewSocket(zmq.PUB)
	defer Publisher.close()
	Publisher.Bind(URI)

	// sync service should run at 14444 port to test the synchronization at client
	syncservice, _ := context.NewSocket(zmq.REP)
	defer syncservice.close()
	syncservice.Bind(syncServiceHost)

	for i := 0; i < subsExpected; i = i + 1 {
		syncservice.Recv(0)
		syncservice.Send([]byte(""), 0)
	}

	for true {
		out, status, err := Publisher.send([][]byte{[]byte(Topic), []byte(Message)}, 0)
	}

	// Set the output value in the context
	if err != nil {
		context.SetOutput(ovoutput, status.Error.Error())
		return false, err
	}

	log.Debugf("Timestamp of the publish response: [%v]", res.Timestamp)

	context.SetOutput(ovoutput, status.syncservice)
	return true, nil
}
