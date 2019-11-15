package cluster

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/pensando/sw/api/generated/security"

	"github.com/pensando/sw/api/generated/network"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/errors"
	"github.com/pensando/sw/api/generated/apiclient"
	"github.com/pensando/sw/api/generated/cluster"
	"github.com/pensando/sw/venice/globals"
)

var _ = Describe("Config SnapShot and restore", func() {
	var grpcClient apiclient.Services
	var restClient apiclient.Services
	var lctx context.Context

	BeforeEach(func() {
		var err error
		validateCluster()
		lctx = ts.tu.MustGetLoggedInContext(context.Background())
		grpcClient = ts.tu.APIClient

		apigwAddr := ts.tu.ClusterVIP + ":" + globals.APIGwRESTPort
		restClient, err = apiclient.NewRestAPIClient(apigwAddr)
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		restClient.Close()
	})

	downloadSnapshot := func(ctx context.Context, filename string) (bytes.Buffer, int, error) {
		var ret bytes.Buffer
		fr, err := ts.tu.SnapshotVOSClient.GetObject(ctx, filename)
		if err != nil {
			return ret, 0, err
		}
		buf := make([]byte, 10*1024)
		totsize := 0
		for {
			n, err := fr.Read(buf)
			if err != nil && err != io.EOF {
				return ret, 0, err
			}
			if n == 0 {
				break
			}
			totsize += n
			if _, err = ret.Write(buf[:n]); err != nil {
				return ret, 0, err
			}
		}
		return ret, totsize, nil
	}

	Context("Snapshot and Restore", func() {
		Context("SnapShot", func() {
			It("Test Multiple snapshots and restore operations", func() {
				By("Save a snapshot")
				{
					cfg := cluster.ConfigurationSnapshot{
						ObjectMeta: api.ObjectMeta{
							Name: "GlobalSnapshot",
						},
						Spec: cluster.ConfigurationSnapshotSpec{
							Destination: cluster.SnapshotDestination{
								Type: cluster.SnapshotDestinationType_ObjectStore.String(),
							},
						},
					}

					if _, err := restClient.ClusterV1().ConfigurationSnapshot().Get(lctx, &api.ObjectMeta{}); err != nil {
						_, err = restClient.ClusterV1().ConfigurationSnapshot().Create(lctx, &cfg)
						Expect(err).To(BeNil())
					}

					req := cluster.ConfigurationSnapshotRequest{}
					var resp *cluster.ConfigurationSnapshot
					var err error
					// Sometimes it takes time for the Object store to be ready, hence the eventually
					Eventually(func() bool {
						resp, err = restClient.ClusterV1().ConfigurationSnapshot().Save(lctx, &req)
						if err != nil {
							return false
						}
						return true
					}, 120, 10).Should(BeTrue(), "Snapshot did not succeed in the given time")
					Expect(err).To(BeNil())
					Expect(resp.Status.LastSnapshot).ToNot(BeNil())
					name := resp.Status.LastSnapshot.URI[strings.LastIndex(resp.Status.LastSnapshot.URI, "/")+1:]
					wrbuf, len, err := downloadSnapshot(lctx, name)
					Expect(err).To(BeNil())

					metadata := map[string]string{
						"Description": "First Save",
					}
					By(fmt.Sprintf("Read file [%s] [%d]bytes", name, len))

					len, err = uploadFile(lctx, "snapshots", "firstsnapshot", metadata, wrbuf.Bytes())
					By(fmt.Sprintf("Wrote [firstsnapshot] [%d] (%s)", len, err))
				}

				By("Make configuration Changes, and save the snapshot")
				{
					{ // Create tenant
						tenant := cluster.Tenant{
							TypeMeta: api.TypeMeta{
								Kind:       "Tenant",
								APIVersion: "v1",
							},
							ObjectMeta: api.ObjectMeta{
								Name: globals.DefaultTenant,
							},
							Spec: cluster.TenantSpec{
								AdminUser: "admin",
							},
						}

						_, err := grpcClient.ClusterV1().Tenant().Get(lctx, &tenant.ObjectMeta)
						if err == nil {
							// Delete all networks
							netws, err := grpcClient.NetworkV1().Network().List(lctx, &api.ListWatchOptions{ObjectMeta: api.ObjectMeta{Tenant: globals.DefaultTenant}})
							By(fmt.Sprintf("got networks [%+v]", netws))
							Expect(err).Should(BeNil(), fmt.Sprintf("got error listing networks (%s)", err))
							for _, n := range netws {
								if strings.HasPrefix("SnapshotTest", n.Name) {
									_, err = grpcClient.NetworkV1().Network().Delete(lctx, &api.ObjectMeta{Tenant: n.Tenant, Name: n.Name})
									Expect(err).Should(BeNil(), fmt.Sprintf("got error deleting networks[%v] (%v)", n.Name, apierrors.FromError(err)))
								}
							}
						} else {
							ret, err := grpcClient.ClusterV1().Tenant().Create(lctx, &tenant)
							Expect(err).To(BeNil())
							Expect(reflect.DeepEqual(ret.Spec, tenant.Spec)).To(Equal(true))
						}
					}
					{ // Create networks
						netw := network.Network{
							TypeMeta: api.TypeMeta{
								Kind:       "Network",
								APIVersion: "v1",
							},
							ObjectMeta: api.ObjectMeta{
								Tenant:    globals.DefaultTenant,
								Name:      "SnapshotTest-Network1",
								Namespace: globals.DefaultNamespace,
							},
							Spec: network.NetworkSpec{
								Type:        "vlan",
								IPv4Subnet:  "255.255.255.0",
								IPv4Gateway: "10.1.1.1",
							},
						}
						ret, err := grpcClient.NetworkV1().Network().Create(lctx, &netw)
						Expect(err).To(BeNil())
						Expect(reflect.DeepEqual(ret.Spec, netw.Spec)).To(Equal(true))
						_, err = grpcClient.NetworkV1().Network().Get(lctx, &netw.ObjectMeta)
						Expect(err).To(BeNil())

						req := cluster.ConfigurationSnapshotRequest{}
						resp, err := restClient.ClusterV1().ConfigurationSnapshot().Save(lctx, &req)
						Expect(err).To(BeNil())
						Expect(resp.Status.LastSnapshot).ToNot(BeNil())
						name := resp.Status.LastSnapshot.URI[strings.LastIndex(resp.Status.LastSnapshot.URI, "/")+1:]
						wrbuf, len, err := downloadSnapshot(lctx, name)
						Expect(err).To(BeNil())

						metadata := map[string]string{
							"Description": "Save with configuration change",
						}
						By(fmt.Sprintf("Read file [%s] [%d]bytes", name, len))

						len, err = uploadFile(lctx, "snapshots", "secondsnapshot", metadata, wrbuf.Bytes())
						By(fmt.Sprintf("Wrote [secondsnapshot] [%d] (%s)", len, err))
					}
				}

				By("Restore the first snapshot")
				{
					req := cluster.SnapshotRestore{
						ObjectMeta: api.ObjectMeta{
							Name: "SnapshotRestore",
						},
						Spec: cluster.SnapshotRestoreSpec{
							SnapshotPath: "firstsnapshot",
						},
					}
					netwMeta := api.ObjectMeta{
						Tenant:    globals.DefaultTenant,
						Name:      "SnapshotTest-Network1",
						Namespace: globals.DefaultNamespace,
					}
					_, err := restClient.NetworkV1().Network().Get(lctx, &netwMeta)
					Expect(err).To(BeNil())
					resp, err := restClient.ClusterV1().SnapshotRestore().Restore(lctx, &req)
					Expect(err).To(BeNil())

					Eventually(func() bool {
						lctx = ts.tu.MustGetLoggedInContext(context.Background())
						resp, err = restClient.ClusterV1().SnapshotRestore().Get(lctx, &resp.ObjectMeta)
						Expect(err).To(BeNil())
						if resp.Status.Status == "completed" {
							return true
						}
						return false
					}, 60, 1).Should(BeTrue(), "Restore did not reach success in the alloted time")

					_, err = restClient.NetworkV1().Network().Get(lctx, &netwMeta)
					Expect(err).To(Not(BeNil()))
				}

				By("Restore the second snapshot")
				{
					req := cluster.SnapshotRestore{
						ObjectMeta: api.ObjectMeta{
							Name: "SnapshotRestore",
						},
						Spec: cluster.SnapshotRestoreSpec{
							SnapshotPath: "secondsnapshot",
						},
					}
					netwMeta := api.ObjectMeta{
						Tenant:    globals.DefaultTenant,
						Name:      "SnapshotTest-Network1",
						Namespace: globals.DefaultNamespace,
					}
					_, err := restClient.NetworkV1().Network().Get(lctx, &netwMeta)
					Expect(err).To(Not(BeNil()))

					resp, err := restClient.ClusterV1().SnapshotRestore().Restore(lctx, &req)
					Expect(err).To(BeNil(), "retore operation returned error")

					Eventually(func() bool {
						lctx = ts.tu.MustGetLoggedInContext(context.Background())
						resp, err = restClient.ClusterV1().SnapshotRestore().Get(lctx, &resp.ObjectMeta)
						Expect(err).To(BeNil())
						if resp.Status.Status == "completed" {
							return true
						}
						return false
					}, 60, 1).Should(BeTrue(), "Restore did not reach success in the alloted time")

					_, err = restClient.NetworkV1().Network().Get(lctx, &netwMeta)
					Expect(err).To(BeNil(), "expecting network object to be present")
				}
			})
		})

		It("Scale comfiguration operations", func() {
			// Create 2000 Apps and save configuration
			cfg := cluster.ConfigurationSnapshot{
				ObjectMeta: api.ObjectMeta{
					Name: "GlobalSnapshot",
				},
				Spec: cluster.ConfigurationSnapshotSpec{
					Destination: cluster.SnapshotDestination{
						Type: cluster.SnapshotDestinationType_ObjectStore.String(),
					},
				},
			}

			if _, err := restClient.ClusterV1().ConfigurationSnapshot().Get(lctx, &api.ObjectMeta{}); err != nil {
				_, err = restClient.ClusterV1().ConfigurationSnapshot().Create(lctx, &cfg)
				Expect(err).To(BeNil())
			}

			req := cluster.ConfigurationSnapshotRequest{}
			var resp *cluster.ConfigurationSnapshot
			var err error
			// Sometimes it takes time for the Object store to be ready, hence the eventually
			Eventually(func() bool {
				resp, err = restClient.ClusterV1().ConfigurationSnapshot().Save(lctx, &req)
				if err != nil {
					return false
				}
				return true
			}, 120, 10).Should(BeTrue(), "Snapshot did not succeed in the given time")
			Expect(err).To(BeNil())
			Expect(resp.Status.LastSnapshot).ToNot(BeNil())
			origName := resp.Status.LastSnapshot.URI[strings.LastIndex(resp.Status.LastSnapshot.URI, "/")+1:]

			app := security.App{
				ObjectMeta: api.ObjectMeta{
					Tenant:    globals.DefaultTenant,
					Namespace: globals.DefaultNamespace,
				},
				Spec: security.AppSpec{
					ProtoPorts: []security.ProtoPort{
						{Protocol: "tcp", Ports: "80"},
					},
				},
			}
			ctx := context.TODO()
			al, err := grpcClient.SecurityV1().App().List(ctx, &api.ListWatchOptions{})
			Expect(err).To(BeNil())
			origLen := len(al)

			for i := 0; i < 2048; i++ {
				app.Name = fmt.Sprintf("scaleSnapshotApp-%d", i)
				_, err = grpcClient.SecurityV1().App().Create(ctx, &app)
				Expect(err).To(BeNil(), "App creation failed [%d[(%s)", i, err)
			}

			// Create second snapshot
			resp, err = restClient.ClusterV1().ConfigurationSnapshot().Save(lctx, &req)
			Expect(err).To(BeNil())
			Expect(resp.Status.LastSnapshot).ToNot(BeNil())
			scaleName := resp.Status.LastSnapshot.URI[strings.LastIndex(resp.Status.LastSnapshot.URI, "/")+1:]

			al, err = grpcClient.SecurityV1().App().List(ctx, &api.ListWatchOptions{})
			Expect(len(al) == 2048+origLen, "number of apps did not match [%d] origi [%d]", len(al), origLen)

			// restore original
			By("restoring original cofiguration")
			restoreReq := cluster.SnapshotRestore{
				ObjectMeta: api.ObjectMeta{
					Name: "SnapshotRestore",
				},
				Spec: cluster.SnapshotRestoreSpec{
					SnapshotPath: origName,
				},
			}

			restoreResp, err := restClient.ClusterV1().SnapshotRestore().Restore(lctx, &restoreReq)
			Expect(err).To(BeNil(), "retore operation returned error")

			Eventually(func() bool {
				lctx = ts.tu.MustGetLoggedInContext(context.Background())
				restoreResp, err = restClient.ClusterV1().SnapshotRestore().Get(lctx, &resp.ObjectMeta)
				Expect(err).To(BeNil())
				if restoreResp.Status.Status == "completed" {
					return true
				}
				return false
			}, 60, 1).Should(BeTrue(), "Restore did not reach success in the alloted time")

			al, err = grpcClient.SecurityV1().App().List(ctx, &api.ListWatchOptions{})
			Expect(len(al) == origLen, "number of apps did not match [%d] origi [%d]", len(al), origLen)

			By("restoring Scale cofiguration")
			restoreReq.Spec.SnapshotPath = scaleName
			restoreResp, err = restClient.ClusterV1().SnapshotRestore().Restore(lctx, &restoreReq)
			Expect(err).To(BeNil(), "retore operation returned error")

			Eventually(func() bool {

				lctx = ts.tu.MustGetLoggedInContext(context.Background())
				restoreResp, err = restClient.ClusterV1().SnapshotRestore().Get(lctx, &resp.ObjectMeta)
				Expect(err).To(BeNil())
				if restoreResp.Status.Status == "completed" {
					return true
				}
				return false
			}, 60, 1).Should(BeTrue(), "Restore did not reach success in the alloted time")

			al, err = grpcClient.SecurityV1().App().List(ctx, &api.ListWatchOptions{})
			Expect(len(al) == 2048+origLen, "number of apps did not match [%d] origi [%d]", len(al), origLen)

			By("restore back to original state before exiting")
			restoreReq.Spec.SnapshotPath = origName
			restoreResp, err = restClient.ClusterV1().SnapshotRestore().Restore(lctx, &restoreReq)
			Expect(err).To(BeNil(), "retore operation returned error")
		})
	})
})