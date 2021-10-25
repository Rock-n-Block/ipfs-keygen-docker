package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"encoding/base64"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
)


func main() {
	size := flag.Int("bitsize", 2048, "select the bitsize of the key to generate")
	typ := flag.String("type", "RSA", "select type of key to generate (RSA or Ed25519)")

	flag.Parse()

	var atyp int
	switch strings.ToLower(*typ) {
	case "rsa":
		atyp = crypto.RSA
	case "ed25519":
		atyp = crypto.Ed25519
	default:
		fmt.Fprintln(os.Stderr, "unrecognized key type: ", *typ)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Generating a %d bit %s key...\n", *size, *typ)
	priv, pub, err := crypto.GenerateKeyPair(atyp, *size)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr, "Success!")

	pid, err := peer.IDFromPublicKey(pub)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "Key ID: %s\n", pid.Pretty())
        
    privbytes, err := crypto.MarshalPrivateKey(priv)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
                os.Exit(1)
	}
 	fmt.Fprintf(os.Stderr, "Private key: %s\n", base64.StdEncoding.EncodeToString(privBytes))
}
