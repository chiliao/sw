package apisrvpkg

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/pensando/sw/api"
	apisrv "github.com/pensando/sw/venice/apiserver"
	compliance "github.com/pensando/sw/venice/utils/kvstore/compliance"
	"github.com/pensando/sw/venice/utils/kvstore/store"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/runtime"
)

type TestType1 struct {
	api.TypeMeta
	api.ObjectMeta
	A string
}

func (t *TestType1) Clone(into interface{}) (interface{}, error) {
	if into != nil {
		in := into.(*TestType1)
		*in = *t
		return in, nil
	}
	ret := *t
	return &ret, nil
}

type TestType2 struct {
	api.TypeMeta
	api.ObjectMeta
	A int
}

func (t *TestType2) Clone(into interface{}) (interface{}, error) {
	if into != nil {
		in := into.(*TestType2)
		*in = *t
		return in, nil
	}
	ret := *t
	return &ret, nil
}

func TestMain(m *testing.M) {
	// Setup the Apiserver Singleton.
	fmt.Printf("Setting up test main")
	_ = MustGetAPIServer()
	buf := &bytes.Buffer{}

	logConfig := log.GetDefaultConfig("TestApiServer")
	l := log.GetNewLogger(logConfig).SetOutput(buf)
	singletonAPISrv.version = "v1"
	singletonAPISrv.Logger = l
	s := runtime.NewScheme()
	s.AddKnownTypes(&TestType1{}, &TestType2{}, &compliance.TestObj{})
	// Add a few KV connections in the pool
	config := apisrv.Config{
		GrpcServerPort: ":0",
		DebugMode:      true,
		Logger:         l,
		Version:        "v1",
		Scheme:         s,
		Kvstore: store.Config{
			Type:  store.KVStoreTypeMemkv,
			Codec: runtime.NewJSONCodec(s),
		},
		KVPoolSize:       1,
		AllowMultiTenant: true,
	}
	singletonAPISrv.config = config

	for i := 0; i < 5; i++ {
		singletonAPISrv.addKvConnToPool()
	}
	rcode := m.Run()
	//fmt.Printf("Test Logs == \n %v\n", buf.String())
	os.Exit(rcode)
}
