// {C} Copyright 2018 Pensando Systems Inc. All rights reserved.

#include <unistd.h>
#include <ev++.h>

#include "gtest/gtest.h"

#include "messenger_client.hpp"
#include "messenger_server.hpp"
#include "nic/delphi/utils/utest.hpp"

namespace {
using namespace std;
using namespace delphi;
using namespace delphi::messenger;

// event loop thread
void * startEventLoop(void* arg) {
    ev::default_loop *loop = (ev::default_loop *)arg;
    loop->run(0);

    return NULL;
}

class srvHandler : public ServerHandler {
public:
    int mountReq = 0;
    int changeReq = 0;
    vector<int> socklist;
    error HandleChangeReq(int sockCtx, vector<ObjectData> req, vector<ObjectData *> *resp) {
        changeReq += req.size();
        return error::OK();
    };
    error HandleMountReq(int sockCtx, MountReqMsgPtr req, MountRespMsgPtr resp) {
        this->socklist.push_back(sockCtx);
        mountReq += req->mounts().size();
        resp->set_serviceid(mountReq);

        // send a response
        for (int i = 0; i < req->mounts().size(); i++) {
            string out_str;
            TestObject tobj;
            ObjectData *od = resp->add_objects();

            ObjectMeta *meta = tobj.mutable_meta();
            meta->set_kind(tobj.GetDescriptor()->name());
            tobj.set_testdata1("Test Data");
            tobj.SerializeToString(&out_str);

            ObjectMeta *ometa = od->mutable_meta();
            ometa->CopyFrom(*meta);
            od->set_op(SetOp);
            od->set_data(out_str);
        }

        return error::OK();
    };
    error HandleSocketClosed(int sockCtx) {
        return error::OK();
    };
};

class clientHandler : public ClientHandler {

public:
    int notify = 0;
    int mountResp = 0;
    int client_id;
    clientHandler(int id) {
        notify = 0;
        mountResp = 0;
        client_id = id;
    }
    error HandleNotify(vector<ObjectData *> objlist) {
        notify++;
        return error::OK();
    };
    error HandleMountResp(uint16_t svcID, string status, vector<ObjectData *> objlist) {
        LogDebug("Client {} Received mount resp. svcID {}", client_id, svcID);
        for(vector<ObjectData *>::iterator iter=objlist.begin(); iter!=objlist.end(); ++iter){
            LogDebug("Mount resp {}", (*iter)->DebugString());
        }
        mountResp += objlist.size();
        return error::OK();
    };
};

#define NUM_CLIENTS 5

class MessangerTest : public testing::Test {
protected:
    pthread_t ev_thread_id = 0;
    ev::default_loop loop;
    shared_ptr<srvHandler> serverHandler;
    shared_ptr<clientHandler> clientHandlers[NUM_CLIENTS];
    MessangerServerPtr server;
    MessangerClientPtr clients[NUM_CLIENTS];
public:
    virtual void SetUp() {
        this->serverHandler = make_shared<srvHandler>();
        this->server = make_shared<MessangerServer>(this->serverHandler);

        // start the server
        this->server->Start();
        usleep(1000);

        // start the clients
        for (int i = 0; i < NUM_CLIENTS; i++) {
            this->clientHandlers[i] = make_shared<clientHandler>(i);
            this->clients[i] = make_shared<MessangerClient>(this->clientHandlers[i]);
            // connect to the server
            this->clients[i]->Connect();
        }

        // start event loop
        pthread_create(&ev_thread_id, 0, &startEventLoop, (void*)&loop);
        usleep(1000);
    }
    virtual void TearDown() {
        // kill the event loop thread
        pthread_cancel(ev_thread_id);
        pthread_join(ev_thread_id, NULL);
        usleep(1000);

        // stop all clients
        for (int i = 0; i < NUM_CLIENTS; i++) {
            this->clients[i]->Close();
        }

        // stop the server
        this->server->Stop();
        usleep(1000);

        LogDebug("Stopping event loop");
        loop.break_loop(ev::ALL);
        usleep(1000);
    }
};

TEST_F(MessangerTest, BasicMountTest) {
    // send mount request
    for (int i = 0; i < NUM_CLIENTS; i++) {
        vector<MountDataPtr>mnts;
        MountDataPtr mnt(make_shared<MountData>());
        mnt->set_kind("Endpoint");
        mnt->set_mode(ReadWriteMode);
        mnts.push_back(mnt);
        string svcName = "TestService-" + std::to_string(i);
        this->clients[i]->SendMountReq(svcName, mnts);
    }
    usleep(1000);

    // verify server got all mount requests
    ASSERT_EQ_EVENTUALLY(serverHandler->mountReq, NUM_CLIENTS) << "Server did not receive all mount requests";
    LogDebug("Server got {} mount requests", serverHandler->mountReq);

    // verify clients got mount resp
    for (int i = 0; i < NUM_CLIENTS; i++) {
        ASSERT_EQ_EVENTUALLY(clientHandlers[i]->mountResp, 1) << "client did not receive mount resp";
        LogDebug("Client {} got {} mount resps\n", i, clientHandlers[i]->mountResp);
    }
}

TEST_F(MessangerTest, ChangeReqTest) {
    // send mount request
    for (int i = 0; i < NUM_CLIENTS; i++) {
        vector<MountDataPtr>mnts;
        MountDataPtr mnt(make_shared<MountData>());
        mnt->set_kind("Endpoint");
        mnt->set_mode(ReadWriteMode);
        mnts.push_back(mnt);
        string svcName = "TestService-" + std::to_string(i);
        this->clients[i]->SendMountReq(svcName, mnts);
    }
    usleep(1000);

    // verify server got all mount requests
    ASSERT_EQ_EVENTUALLY(serverHandler->mountReq, NUM_CLIENTS) << "Server did not receive all mount requests";
    LogDebug("Server got {} mount requests", serverHandler->mountReq);

    // send change req from each client
    for (int i = 0; i < NUM_CLIENTS; i++) {
        string out_str;
        TestObject tobj;
        ObjectData *od =  new ObjectData;
        vector<ObjectData *> objlist;

        ObjectMeta *meta = tobj.mutable_meta();
        meta->set_kind(tobj.GetDescriptor()->name());
        tobj.mutable_key()->set_idx((uint32_t)i);
        tobj.set_testdata1("Test Data");
        tobj.SerializeToString(&out_str);

        ObjectMeta *ometa = od->mutable_meta();
        ometa->CopyFrom(*meta);
        od->set_op(SetOp);
        od->set_data(out_str);
        objlist.push_back(od);

        this->clients[i]->SendChangeReq(objlist);
    }
    usleep(1000);

    ASSERT_EQ_EVENTUALLY(serverHandler->changeReq, NUM_CLIENTS) << "Server did not receive all change requests";
    LogDebug("Server got {} change requests\n", serverHandler->changeReq);

    // send a notify message to each client
    for(vector<int>::iterator iter=serverHandler->socklist.begin(); iter!=serverHandler->socklist.end(); ++iter){
        int sock = (*iter);
        string out_str;
        TestObject tobj;
        ObjectData *od = new ObjectData;
        vector<ObjectData *> objlist;

        ObjectMeta *meta = tobj.mutable_meta();
        meta->set_kind(tobj.GetDescriptor()->name());
        tobj.set_testdata1("Test Data");
        tobj.SerializeToString(&out_str);

        ObjectMeta *ometa = od->mutable_meta();
        ometa->CopyFrom(*meta);
        od->set_op(SetOp);
        od->set_data(out_str);
        objlist.push_back(od);

        this->server->SendNotify(sock, objlist);
    }
    usleep(1000 * 100);

    // verify clients got notify message
    for (int i = 0; i < NUM_CLIENTS; i++) {
        ASSERT_EQ_EVENTUALLY(clientHandlers[i]->notify, 1) << "client did not receive notify message " << i;
        LogDebug("Client {} got {} notify messages\n", i, clientHandlers[i]->notify);
    }

}


TEST_F(MessangerTest, BidirMsgBenchmark) {
    int numIter = 200;
    int batchSize = 5000;
    for (int i = 0; i < numIter; i++) {
        vector<MountDataPtr>mnts;
        for (int j = 0; j < batchSize; j++) {
            MountDataPtr mnt = make_shared<MountData>();
            mnt->set_kind("Endpoint");
            mnt->set_mode(ReadWriteMode);
            mnts.push_back(mnt);
        }
        this->clients[0]->SendMountReq("TestService", mnts);
    }

    for (int iter = 0; iter < 1000; iter++) {
        if ((serverHandler->mountReq == numIter * batchSize) && (clientHandlers[0]->mountResp == numIter * batchSize)) {
            break;
        }
        usleep(1000 * 10);
    }
    ASSERT_EQ_EVENTUALLY(serverHandler->mountReq, numIter * batchSize) << "server did not receive all mount requests";
    ASSERT_EQ_EVENTUALLY(clientHandlers[0]->mountResp, numIter * batchSize) << "client did not receive all mount response";
}


TEST_F(MessangerTest, Client2ServerBenchmark) {
    int numIter = 500;
    int batchSize = 1000;
    for (int i = 0; i < numIter; i++) {
        vector<ObjectData *> objlist;
        for (int j = 0; j < batchSize; j++) {
            string out_str;
            TestObject tobj;
            ObjectData *od = new ObjectData;
            assert(od != NULL);

            ObjectMeta *meta = tobj.mutable_meta();
            meta->set_kind(tobj.GetDescriptor()->name());
            tobj.set_testdata1("Test Data");
            tobj.SerializeToString(&out_str);

            ObjectMeta *ometa = od->mutable_meta();
            ometa->CopyFrom(*meta);
            od->set_op(SetOp);
            od->set_data(out_str);
            objlist.push_back(od);

        }
        this->clients[0]->SendChangeReq(objlist);
    }

    for (int iter = 0; iter < 1000; iter++) {
        if (serverHandler->changeReq == numIter * batchSize)  {
            break;
        }
        usleep(1000 * 10);
    }
    ASSERT_EQ(serverHandler->changeReq, numIter * batchSize) << "server did not receive all change requests";

}

} // namespace

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
