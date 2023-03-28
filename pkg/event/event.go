package event

import (
	"github.com/decred/dcrd/dcrec/secp256k1/v4"

	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"
)

type EventOptions struct {
	Content   string
	EventData []byte
	Kind      uint32
	Tags      [][]byte
}

func New(opt *EventOptions) (*Event, error) {
	newAEAD := func(key []byte) (cipher.AEAD, error) {
		block, err := aes.NewCipher(key)

		if err != nil {
			return nil, err
		}

		return cipher.NewGCM(block)
	}

	// Decode the hex-encoded pubkey of the recipient.
	pubKeyBytes, err := hex.DecodeString("04115c42e757b2efb7671c578530ec191a1" +
		"359381e6a71127a9d37c486fd30dae57e76dc58f693bd7e7010358ce6b165e483a29" +
		"21010db67ac11b1b51b651953d2") // uncompressed pubkey

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	pubKey, err := secp256k1.ParsePubKey(pubKeyBytes)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Derive an ephemeral public/private keypair for performing ECDHE with
	// the recipient.
	ephemeralPrivKey, err := secp256k1.GeneratePrivateKey()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	ephemeralPubKey := ephemeralPrivKey.PubKey().SerializeCompressed()

	// Using ECDHE, derive a shared symmetric key for encryption of the plaintext.
	cipherKey := sha256.Sum256(secp256k1.GenerateSharedSecret(ephemeralPrivKey, pubKey))

	// Seal the message using an AEAD.  Here we use AES-256-GCM.
	// The ephemeral public key must be included in this message, and becomes
	// the authenticated data for the AEAD.
	//
	// Note that unless a unique nonce can be guaranteed, the ephemeral
	// and/or shared keys must not be reused to encrypt different messages.
	// Doing so destroys the security of the scheme.  Random nonces may be
	// used if XChaCha20-Poly1305 is used instead, but the message must then
	// also encode the nonce (which we don't do here).
	//
	// Since a new ephemeral key is generated for every message ensuring there
	// is no key reuse and AES-GCM permits the nonce to be used as a counter,
	// the nonce is intentionally initialized to all zeros so it acts like the
	// first (and only) use of a counter.
	plaintext := []byte("test message")
	aead, err := newAEAD(cipherKey[:])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	nonce := make([]byte, aead.NonceSize())
	ciphertext := make([]byte, 4+len(ephemeralPubKey))
	binary.LittleEndian.PutUint32(ciphertext, uint32(len(ephemeralPubKey)))
	copy(ciphertext[4:], ephemeralPubKey)
	ciphertext = aead.Seal(ciphertext, nonce, plaintext, ephemeralPubKey)

	id := make([]byte, 32)
	publicKey := make([]byte, 32)
	createdAt := time.Now()

	base64.StdEncoding.Encode(id, opt.EventData)

	return &Event{
		Content:   opt.Content,
		CreatedAt: createdAt,
		ID:        id,
		Kind:      opt.Kind,
		PublicKey: publicKey,
		Tags:      opt.Tags,
	}, nil
}

func Validate(event *Event) error {
	return nil
}

type Event struct {
	ID        []byte
	PublicKey []byte
	CreatedAt time.Time
	Kind      uint32
	Tags      [][]byte
	Content   string
	Sig       string
}
