// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package networkCliUtilsBackend is a auto generated package.
Input file: svc_network.proto
*/
package cli

import (
	"context"
	"fmt"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/apiclient"
	"github.com/pensando/sw/api/generated/network"
	loginctx "github.com/pensando/sw/api/login/context"
	"github.com/pensando/sw/venice/cli/gen"
)

func restGetNetwork(hostname, tenant, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*network.Network); ok {
		nv, err := restcl.NetworkV1().Network().Get(loginCtx, &v.ObjectMeta)
		if err != nil {
			return err
		}
		*v = *nv
	}

	if v, ok := obj.(*network.NetworkList); ok {
		opts := api.ListWatchOptions{ObjectMeta: api.ObjectMeta{Tenant: tenant}}
		nv, err := restcl.NetworkV1().Network().List(loginCtx, &opts)
		if err != nil {
			return err
		}
		v.Items = nv
	}
	return nil

}

func restDeleteNetwork(hostname, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*network.Network); ok {
		nv, err := restcl.NetworkV1().Network().Delete(loginCtx, &v.ObjectMeta)
		if err != nil {
			return err
		}
		*v = *nv
	}
	return nil

}

func restPostNetwork(hostname, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*network.Network); ok {
		nv, err := restcl.NetworkV1().Network().Create(loginCtx, v)
		if err != nil {
			return err
		}
		*v = *nv
	}
	return nil

}

func restPutNetwork(hostname, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*network.Network); ok {
		nv, err := restcl.NetworkV1().Network().Update(loginCtx, v)
		if err != nil {
			return err
		}
		*v = *nv
	}
	return nil

}

func restGetService(hostname, tenant, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*network.Service); ok {
		nv, err := restcl.NetworkV1().Service().Get(loginCtx, &v.ObjectMeta)
		if err != nil {
			return err
		}
		*v = *nv
	}

	if v, ok := obj.(*network.ServiceList); ok {
		opts := api.ListWatchOptions{ObjectMeta: api.ObjectMeta{Tenant: tenant}}
		nv, err := restcl.NetworkV1().Service().List(loginCtx, &opts)
		if err != nil {
			return err
		}
		v.Items = nv
	}
	return nil

}

func restDeleteService(hostname, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*network.Service); ok {
		nv, err := restcl.NetworkV1().Service().Delete(loginCtx, &v.ObjectMeta)
		if err != nil {
			return err
		}
		*v = *nv
	}
	return nil

}

func restPostService(hostname, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*network.Service); ok {
		nv, err := restcl.NetworkV1().Service().Create(loginCtx, v)
		if err != nil {
			return err
		}
		*v = *nv
	}
	return nil

}

func restPutService(hostname, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*network.Service); ok {
		nv, err := restcl.NetworkV1().Service().Update(loginCtx, v)
		if err != nil {
			return err
		}
		*v = *nv
	}
	return nil

}

func restGetLbPolicy(hostname, tenant, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*network.LbPolicy); ok {
		nv, err := restcl.NetworkV1().LbPolicy().Get(loginCtx, &v.ObjectMeta)
		if err != nil {
			return err
		}
		*v = *nv
	}

	if v, ok := obj.(*network.LbPolicyList); ok {
		opts := api.ListWatchOptions{ObjectMeta: api.ObjectMeta{Tenant: tenant}}
		nv, err := restcl.NetworkV1().LbPolicy().List(loginCtx, &opts)
		if err != nil {
			return err
		}
		v.Items = nv
	}
	return nil

}

func restDeleteLbPolicy(hostname, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*network.LbPolicy); ok {
		nv, err := restcl.NetworkV1().LbPolicy().Delete(loginCtx, &v.ObjectMeta)
		if err != nil {
			return err
		}
		*v = *nv
	}
	return nil

}

func restPostLbPolicy(hostname, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*network.LbPolicy); ok {
		nv, err := restcl.NetworkV1().LbPolicy().Create(loginCtx, v)
		if err != nil {
			return err
		}
		*v = *nv
	}
	return nil

}

func restPutLbPolicy(hostname, token string, obj interface{}) error {

	restcl, err := apiclient.NewRestAPIClient(hostname)
	if err != nil {
		return fmt.Errorf("cannot create REST client")
	}
	defer restcl.Close()
	loginCtx := loginctx.NewContextWithAuthzHeader(context.Background(), "Bearer "+token)

	if v, ok := obj.(*network.LbPolicy); ok {
		nv, err := restcl.NetworkV1().LbPolicy().Update(loginCtx, v)
		if err != nil {
			return err
		}
		*v = *nv
	}
	return nil

}

func init() {
	cl := gen.GetInfo()
	if cl == nil {
		return
	}

	cl.AddRestPostFunc("network.Network", "v1", restPostNetwork)
	cl.AddRestDeleteFunc("network.Network", "v1", restDeleteNetwork)
	cl.AddRestPutFunc("network.Network", "v1", restPutNetwork)
	cl.AddRestGetFunc("network.Network", "v1", restGetNetwork)

	cl.AddRestPostFunc("network.Service", "v1", restPostService)
	cl.AddRestDeleteFunc("network.Service", "v1", restDeleteService)
	cl.AddRestPutFunc("network.Service", "v1", restPutService)
	cl.AddRestGetFunc("network.Service", "v1", restGetService)

	cl.AddRestPostFunc("network.LbPolicy", "v1", restPostLbPolicy)
	cl.AddRestDeleteFunc("network.LbPolicy", "v1", restDeleteLbPolicy)
	cl.AddRestPutFunc("network.LbPolicy", "v1", restPutLbPolicy)
	cl.AddRestGetFunc("network.LbPolicy", "v1", restGetLbPolicy)

}
