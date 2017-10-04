// {C} Copyright 2017 Pensando Systems Inc. All rights reserved.

package keymgr

import (
	"crypto"
	"crypto/x509"
	"io"
)

// KeyType encodes the algorithm and key size of the key pair
type KeyType int

// Known types for asymmetric keys
const (
	Unknown KeyType = iota
	RSA1024
	RSA2048
	RSA4096
	ECDSA224
	ECDSA256
	ECDSA384
	ECDSA521
)

// ObjectType represents the type of the objects that can be handled by KeyMgr
type ObjectType int

// ObjectType values
const (
	ObjectTypeKeyPair ObjectType = iota
	ObjectTypeCertificate
	ObjectTypeCertificateBundle
)

// Object represents a cryptographic artifact that can be stored and retrieved
// ID space is shared across object types
type Object interface {
	// ID is a unique identifier that allows client to refer to the object
	// when performing operations (lookup, destroy, etc.)
	ID() string

	// ObjectType is the type of the object (key, certificate, etc.)
	Type() ObjectType
}

// the members and methods that are common to all objects handled by KeyMgr
type object struct {
	id      string
	objType ObjectType
}

func (o *object) ID() string {
	return o.id
}

func (o *object) Type() ObjectType {
	return o.objType
}

// KeyPair is a KeyMgr object representing a (public key, private key) pair.
// It implements the crypto.Signer interface, so it can be passed to any Go crypto
// library function that needs to sign data. The public key can be retrieved using
// the PublicKey() method, the private key is only available indirectly using
// the Sign() method.
type KeyPair struct {
	object
	KeyType KeyType
	Signer  crypto.Signer
}

// Public returns the public key in the pair.
// Implements Go crypto.Signer Public() function.
func (kp *KeyPair) Public() crypto.PublicKey {
	return kp.Signer.Public()
}

// Sign signs the message using the private key in the pair
// Implements Go crypto.Signer Sign() function.
// Since KeyPairs are generated by backend, the implementation details are backend-specific.
func (kp *KeyPair) Sign(rand io.Reader, msg []byte, opts crypto.SignerOpts) (signature []byte, err error) {
	return kp.Signer.Sign(rand, msg, opts)
}

// NewKeyPairObject is a convenience function that takes a crypto.Signer and
// an ID and returns a KeyMgr Object
func NewKeyPairObject(id string, signer crypto.Signer) *KeyPair {
	return &KeyPair{
		object: object{
			id:      id,
			objType: ObjectTypeKeyPair,
		},
		KeyType: getKeyType(signer.Public()),
		Signer:  signer,
	}
}

// Certificate is a KeyMgr object holding a x509 Certificate
type Certificate struct {
	object
	Certificate *x509.Certificate
}

// NewCertificateObject is a convenience function that takes a x509.Certificate and
// returns a KeyMgr Object
func NewCertificateObject(id string, cert *x509.Certificate) *Certificate {
	return &Certificate{
		object: object{
			id:      id,
			objType: ObjectTypeCertificate,
		},
		Certificate: cert,
	}
}

// CertificateBundle is a KeyMgr object holding an ordered collection of x509 Certificates.
// This is KeyMgr-level abstraction, backends are not expected to have any notion of it.
// Certificates are cached individually. There is no transaction support. If we fail storing
// one of the certificates, client is supposed to handle the failure and cleanup the truncated
// bundle.
type CertificateBundle struct {
	object
	Certificates []*x509.Certificate
}

// NewCertificateBundleObject is a convenience function that takes an array of x509.Certificates
// and returns a KeyMgr Object
func NewCertificateBundleObject(id string, certs []*x509.Certificate) *CertificateBundle {
	return &CertificateBundle{
		object: object{
			id:      id,
			objType: ObjectTypeCertificateBundle,
		},
		Certificates: certs,
	}
}

// Backend is the interface that needs to be implemented by KeyMgr backends
// Initialization functions are backend-specific and so are not part of the interface
type Backend interface {
	// Creates a key pair, with the private key confined inside the backend
	CreateKeyPair(id string, kt KeyType) (*KeyPair, error)

	// Store an object of type ObjectType, indexed with the supplied id.
	// If object already exists, an error is returned.
	StoreObject(obj Object) error

	// Retrieved the object with the given id. If object is not found, return nil.
	// Error is set only if there was an error in accessing the backend
	GetObject(id string, Type ObjectType) (Object, error)

	// Destroy the object with the given id and type. If object is not found, return nil
	// Error is set only if there was an error in accessing the backend
	DestroyObject(ID string, Type ObjectType) error

	// Frees up the resources held by the backend
	Close() error
}
