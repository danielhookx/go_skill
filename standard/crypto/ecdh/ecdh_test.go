package ecdh

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"testing"
)

func TestECDH(t *testing.T) {
	privateA, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	privateB, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	publicA := privateA.PublicKey
	publicB := privateB.PublicKey

	t.Logf("Alice private key is %x\n", privateA.D)
	t.Logf("Bob private key is %x\n", privateB.D)

	t.Logf("Alice publuic key is (%x,%x)\n", publicA.X, publicA.Y)
	t.Logf("Bob private key is (%x,%x)\n", publicB.X, publicB.Y)

	a, _ := privateA.Curve.ScalarMult(publicB.X, publicB.Y, privateA.D.Bytes())
	a, _ := publicA.Curve.ScalarMult(publicA.X, publicA.Y, privb.D.Bytes())
	b, _ := pubb.Curve.ScalarMult(pubb.X, pubb.Y, priva.D.Bytes())

	shared1 := sha256.Sum256(a.Bytes())
	shared2 := sha256.Sum256(b.Bytes())

}
