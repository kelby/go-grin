package libtx

import "keychain"

type Aggsig interface {
	Create_secnonce(secp *Secp256k1) (SecretKey, error)

	Calculate_partial_sig(secp *Secp256k1, sec_key *SecretKey, sec_nonce *SecretKey, nonce_sum *PublicKey, fee uint64, lock_height uint64) (Signature, error)

	Verify_partial_sig(secp *Secp256k1, sig *Signature, pub_nonce_sum *PublicKey, pubkey *PublicKey, fee uint64, lock_height uint64) error

	Sign_from_key_id(secp *Secp256k1, k *keychain.Keychain, msg *Message, key_id *keychain.Identifier) (Signature, error)

	Verify_single_from_commit(secp *Secp256k1, sig *Signature, msg *Message, commit *Commitment) error

	Verify_sig_build_msg(secp *Secp256k1, sig *Signature, pubkey *PublicKey, fee u64, lock_height u64) error

	Verify_single(secp *Secp256k1, sig *Signature, msg *Message, pubnonce *PublicKey, pubkey *PublicKey, is_partial bool) bool

	Add_signatures(secp *Secp256k1, part_sigs []*Signature, nonce_sum *PublicKey) (Signature, error)

	Sign_with_blinding(secp *Secp256k1, msg *Message, blinding *keychain.BlindingFactor) (Signature, error)
}
