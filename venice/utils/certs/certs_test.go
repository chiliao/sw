// {C} Copyright 2017 Pensando Systems Inc. All rights reserved.

package certs

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net/rpc"
	"os"
	"reflect"
	"strings"
	"testing"

	. "github.com/pensando/sw/venice/utils/testutils"
)

const (
	numRsaBits = 4096
	days       = 365
)

func testSaveAndReadPrivateKey(privateKey crypto.PrivateKey, t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "cert_test")
	if err != nil {
		log.Fatal(err)
	}
	tmpfileName := tmpfile.Name()
	defer os.Remove(tmpfileName)

	err = SavePrivateKey(tmpfileName, privateKey)
	AssertOk(t, err, "SavePrivateKey fail")

	readPrivateKey, err := ReadPrivateKey(tmpfileName)
	AssertOk(t, err, "ReadPrivateKey fail")
	Assert(t, reflect.DeepEqual(privateKey, readPrivateKey), ("ReadPrivateKey is not same as expected PrivateKey"))
}

func TestSaveAndReadRsaPrivateKey(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, numRsaBits)
	AssertOk(t, err, "GenerateKey fail")
	testSaveAndReadPrivateKey(privateKey, t)
}

func TestSaveAndReadEcPrivateKey(t *testing.T) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	AssertOk(t, err, "GenerateKey fail. err: %s")
	testSaveAndReadPrivateKey(privateKey, t)
}

func testSaveAndReadCertificate(privateKey crypto.PrivateKey, t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "cert_test")
	AssertOk(t, err, "Error generating temp file")
	tmpfileName := tmpfile.Name()
	defer os.Remove(tmpfileName)

	cert, err := SelfSign(days, "", privateKey)
	AssertOk(t, err, "Error creating self-signed certificate")
	err = SaveCertificate(tmpfileName, cert)
	AssertOk(t, err, "Error saving certificate")
	err = SaveCertificate("/tmp", cert)
	Assert(t, err != nil, "SaveCertificate succeeded writing to invalid filename.")

	readCerts, err := ReadCertificates(tmpfileName)
	AssertOk(t, err, "ReadCertificate fail")
	Assert(t, cert.Equal(readCerts[0]), "read cert is not same as expected cert")
	// Readcertificate should produce the same result as ReadCertificates when there is only 1 cert
	readOneCert, err := ReadCertificate(tmpfileName)
	AssertOk(t, err, "Error reading certificate using ReadCertificate")
	Assert(t, cert.Equal(readOneCert), "read cert is not same as expected cert")
}

func TestSaveAndReadRsaCertificate(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, numRsaBits)
	if err != nil {
		t.Fatalf("GenerateKey fail. err: %s", err.Error())
	}

	testSaveAndReadCertificate(privateKey, t)
}

func TestSaveAndReadEcCertificate(t *testing.T) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	AssertOk(t, err, "GenerateKey fail")
	testSaveAndReadCertificate(privateKey, t)
}

func testSaveAndReadCSR(privateKey crypto.PrivateKey, t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "cert_test")
	AssertOk(t, err, "Error creating temporary file")
	tmpfileName := tmpfile.Name()
	defer os.Remove(tmpfileName)

	certSignReq, err := CreateCSR(privateKey, nil, nil)
	AssertOk(t, err, "Error creating CSR")
	SaveCSR(tmpfileName, certSignReq)
	readcsr, err := ReadCSR(tmpfileName)
	AssertOk(t, err, "Error reading CSR")
	Assert(t, reflect.DeepEqual(readcsr, certSignReq), "read CSR is not same as expected cert")
}

func TestSaveAndReadRsaCSR(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, numRsaBits)
	AssertOk(t, err, "GenerateKey fail.")
	testSaveAndReadCSR(privateKey, t)
}

func TestSaveAndReadEcCSR(t *testing.T) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	AssertOk(t, err, "GenerateKey fail.")
	testSaveAndReadCSR(privateKey, t)
}

// Check that we can properly read a private key file generated by OpenSSL
// OpenSSL puts key parameters both in the key block itself and in a separate "EC PARAMETERS" block
func TestReadOpenSSLGeneratedEcKey(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "cert_test")
	AssertOk(t, err, "Error creating temporary file")
	tmpFileName := tmpFile.Name()
	defer os.Remove(tmpFileName)

	keyParamBlock := `
-----BEGIN EC PARAMETERS-----
BgUrgQQAIw==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MIHcAgEBBEIBw19j4zd8aEMsCqBsGfrLT93ywnovsOEmTGkHnNZxQ+9U3HZvYEZA
QMUobxlj891ioExvRwm7aY7r6Hjnb+lCkLqgBwYFK4EEACOhgYkDgYYABADqG0/0
cp2+HjmqafBSgYonsrGboMHkLfT2J7YdGKZCCyebJMoDf6JBZxwcOKJ9mFj6wUy/
x0bxRsNd/YdNH9uiQwBt7vHGUb1uyEniyoFPyoVQqn6mqdp2nY21OwkHcMQ6U6C1
Uqvhc8wvGrVwYLlrIcGNcnZxEglGXJXTFwxQWSMuQQ==
-----END EC PRIVATE KEY-----
`
	_, err = tmpFile.WriteString(keyParamBlock)
	AssertOk(t, err, "Error writing to temp file")
	tmpFile.Sync()

	privateKey, err := ReadPrivateKey(tmpFileName)
	AssertOk(t, err, "Error reading private key from file")

	// compare against the key that we can parse directly out of the
	// "EC PRIVATE KEY" block
	_, keyBlock := pem.Decode([]byte(keyParamBlock))
	key, _ := pem.Decode(keyBlock)
	privateKeyRef, err := x509.ParseECPrivateKey(key.Bytes)
	AssertOk(t, err, "Error parsing EC private key")
	Assert(t, reflect.DeepEqual(privateKey, privateKeyRef), "Private key does not match")

	// Try reverse as well, as it seems to be legal
	revKeyParamBlock := `
-----BEGIN EC PRIVATE KEY-----
MIHcAgEBBEIBw19j4zd8aEMsCqBsGfrLT93ywnovsOEmTGkHnNZxQ+9U3HZvYEZA
QMUobxlj891ioExvRwm7aY7r6Hjnb+lCkLqgBwYFK4EEACOhgYkDgYYABADqG0/0
cp2+HjmqafBSgYonsrGboMHkLfT2J7YdGKZCCyebJMoDf6JBZxwcOKJ9mFj6wUy/
x0bxRsNd/YdNH9uiQwBt7vHGUb1uyEniyoFPyoVQqn6mqdp2nY21OwkHcMQ6U6C1
Uqvhc8wvGrVwYLlrIcGNcnZxEglGXJXTFwxQWSMuQQ==
-----END EC PRIVATE KEY-----
-----BEGIN EC PARAMETERS-----
BgUrgQQAIw==
-----END EC PARAMETERS-----
`
	// rewind to beginning of file
	tmpFile.Seek(0, 0)
	_, err = tmpFile.WriteString(revKeyParamBlock)
	AssertOk(t, err, "Error writing to temp file")
	tmpFile.Sync()

	privateKey, err = ReadPrivateKey(tmpFileName)
	AssertOk(t, err, "Error reading EC private key from file")
	Assert(t, reflect.DeepEqual(privateKey, privateKeyRef), "Private key does not match")
}

func TestCertificateValidationRoutines(t *testing.T) {
	// RSA keys
	rsaPrivateKey, err := rsa.GenerateKey(rand.Reader, numRsaBits)
	AssertOk(t, err, "GenerateKey fail")
	rsaPrivateKey2, err := rsa.GenerateKey(rand.Reader, numRsaBits)
	AssertOk(t, err, "GenerateKey fail")

	// ECDSA keys
	ecdsaPrivateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	AssertOk(t, err, "GenerateKey fail")
	ecdsaPrivateKey2, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	AssertOk(t, err, "GenerateKey fail")

	keys := [][]crypto.PrivateKey{
		{rsaPrivateKey, rsaPrivateKey2},
		{ecdsaPrivateKey, ecdsaPrivateKey2}}

	for i := 0; i < len(keys); i++ {
		privateKey := keys[i][0]
		altPrivateKey := keys[i][1]

		cert, err := SelfSign(days, "", privateKey)
		AssertOk(t, err, "Error creating self-signed certificate")
		Assert(t, IsSelfSigned(cert), "Failed to detect self-signed certificate")
		valid, err := ValidateKeyCertificatePair(privateKey, cert)
		AssertOk(t, err, "Error validating certificate")
		Assert(t, valid, "Certificate failed key validation")

		// this is expected to fail
		valid, err = ValidateKeyCertificatePair(altPrivateKey, cert)
		Assert(t, (!valid || err != nil), "Certificate did not fail key validation as expected")
	}
}

type Arith int
type Args struct {
	A, B int
}

type Reply struct {
	C int
}

func (t *Arith) Add(args Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}

func startServer(t *testing.T, serverCertFile, serverPrivKeyFile, caCertFile string) string {
	rpc.Register(new(Arith))

	cert, err := tls.LoadX509KeyPair(serverCertFile, serverPrivKeyFile)
	AssertOk(t, err, "server: loadkeys")

	bytes, err := ioutil.ReadFile(caCertFile)
	AssertOk(t, err, fmt.Sprintf("Error reading CA cert file %v", caCertFile))
	block, _ := pem.Decode(bytes)
	ca, err := x509.ParseCertificate(block.Bytes)
	AssertOk(t, err, "Error parsing certificate")

	certPool := NewCertPool([]*x509.Certificate{ca})
	config := tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,

		// not needed for pure  servers. But if we are clients also, we will need this.
		RootCAs: certPool,
	}
	config.Rand = rand.Reader

	service, _ := os.Hostname()
	service = service + ":0" // any available address
	listener, err := tls.Listen("tcp", service, &config)
	AssertOk(t, err, "Error listening")
	serverAddr := listener.Addr().String()
	log.Println("Test RPC server listening on", serverAddr)
	go rpc.Accept(listener)
	return serverAddr
}

func generateKeysAndCerts(t *testing.T, caCertFile, serverCertFile, serverPrivKeyFile, clientCertFile, clientPrivKeyFile string) {
	days := 365
	numRsaBits := 2048
	caprivatekey, err := rsa.GenerateKey(rand.Reader, numRsaBits)
	AssertOk(t, err, "GenerateKey fail")
	cacert, err := SelfSign(days, "", caprivatekey)
	AssertOk(t, err, "Error generating self-signed certificate")
	SaveCertificate(caCertFile, cacert)

	srvprivatekey, err := rsa.GenerateKey(rand.Reader, numRsaBits)
	AssertOk(t, err, "GenerateKey fail")
	SavePrivateKey(serverPrivKeyFile, srvprivatekey)
	csr, err := CreateCSR(srvprivatekey, nil, nil)
	AssertOk(t, err, "Error generating CSR")
	srvcert, err := SignCSRwithCA(days, csr, cacert, caprivatekey)
	AssertOk(t, err, "Error signing CSR")
	SaveCertificate(serverCertFile, srvcert)

	clientprivatekey, err := rsa.GenerateKey(rand.Reader, numRsaBits)
	AssertOk(t, err, "GenerateKey fail")
	SavePrivateKey(clientPrivKeyFile, clientprivatekey)
	csr, err = CreateCSR(clientprivatekey, nil, nil)
	AssertOk(t, err, "Error generating CSR")
	clientcert, err := SignCSRwithCA(days, csr, cacert, caprivatekey)
	AssertOk(t, err, "Error signing CSR")
	SaveCertificate(clientCertFile, clientcert)
}

func TestRPC(t *testing.T) {
	F, err := ioutil.TempFile("", "serverCertFile")
	AssertOk(t, err, "client: loadkeys")
	serverCertFile := F.Name()
	defer os.Remove(serverCertFile)

	F, err = ioutil.TempFile("", "serverPrivKeyFile")
	AssertOk(t, err, "Error creating temporary file")
	serverPrivKeyFile := F.Name()
	defer os.Remove(serverPrivKeyFile)

	F, err = ioutil.TempFile("", "caCertFile")
	AssertOk(t, err, "Error creating temporary file")
	caCertFile := F.Name()
	defer os.Remove(caCertFile)

	F, err = ioutil.TempFile("", "clientCertFile")
	AssertOk(t, err, "Error creating temporary file")
	clientCertFile := F.Name()
	defer os.Remove(clientCertFile)

	F, err = ioutil.TempFile("", "clientPrivKeyFile")
	AssertOk(t, err, "Error creating temporary file")
	clientPrivKeyFile := F.Name()
	defer os.Remove(clientPrivKeyFile)

	generateKeysAndCerts(t, caCertFile, serverCertFile, serverPrivKeyFile, clientCertFile, clientPrivKeyFile)
	serverAddr := startServer(t, serverCertFile, serverPrivKeyFile, caCertFile)

	// start client with the IPaddress:port and connect to server with RPC request
	startClientAndDoRPC(t, serverAddr, clientCertFile, clientPrivKeyFile, caCertFile)

	serverPort := strings.Split(serverAddr, ":")[1]
	hostname, _ := os.Hostname()
	serverAddrWithName := hostname + ":" + serverPort
	// start client with the hostname:port and connect to server with RPC request
	startClientAndDoRPC(t, serverAddrWithName, clientCertFile, clientPrivKeyFile, caCertFile)
}

func startClientAndDoRPC(t *testing.T, addr string, clientCertFile, clientPrivKeyFile, caCertFile string) {
	cert, err := tls.LoadX509KeyPair(clientCertFile, clientPrivKeyFile)
	AssertOk(t, err, "client: loadkeys")
	bytes, err := ioutil.ReadFile(caCertFile)
	AssertOk(t, err, fmt.Sprintf("Error reading caCertFile: %v", caCertFile))
	block, _ := pem.Decode(bytes)
	ca, err := x509.ParseCertificate(block.Bytes)
	AssertOk(t, err, "Error parsing certificate")

	certPool := NewCertPool([]*x509.Certificate{ca})
	config := tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,

		// not needed for pure  clients. But if we are server also, then we will need this.
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  certPool,
	}

	conn, err := tls.Dial("tcp", addr, &config)
	AssertOk(t, err, fmt.Sprintf("client: dial error while connecting to server: %s", addr))
	defer conn.Close()
	log.Println("client: Dial: ", addr, " connected to: ", conn.RemoteAddr())
	rpcClient := rpc.NewClient(conn)

	args := &Args{7, 8}
	reply := new(Reply)
	err = rpcClient.Call("Arith.Add", args, reply)
	if err != nil {
		t.Errorf("Add: expected no error but got string %q", err.Error())
	}
	if reply.C != args.A+args.B {
		t.Errorf("Add: expected %d got %d", reply.C, args.A+args.B)
	}
}
