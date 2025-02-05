// Code generated by command: go run fe_amd64_asm.go -out ../fe_amd64.s -stubs ../fe_amd64.go -pkg edwards25519. DO NOT EDIT.

// +build amd64,gc,!purego

package edwards25519

// feMul sets out = a * b. It works like feMulGeneric.
//go:noescape
func feMul(out *fieldElement, a *fieldElement, b *fieldElement)

// feSquare sets out = a * a. It works like feSquareGeneric.
//go:noescape
func feSquare(out *fieldElement, a *fieldElement)
