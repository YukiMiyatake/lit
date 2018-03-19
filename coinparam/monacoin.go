package coinparam

import (
	"time"

	"github.com/adiabat/btcd/chaincfg/chainhash"
	"github.com/adiabat/btcd/wire"
	"golang.org/x/crypto/scrypt"
	//	"golang.org/x/crypto/scrypt"
)

// MonaCoinTestNet4Params are the parameters for the monacoin test network 4.
var MonaCoinTestNet4Params = Params{
	Name:          "monatest4",
	NetMagicBytes: 0xf1c8d2fd,
	DefaultPort:   "19403",
	DNSSeeds: []string{
		"testnet-dnsseed.monacoin.org",
		"electrumx1.testnet.monacoin.nl",
		"electrumx1.testnet.monacoin.ninja",
	},

	// Chain parameters
	GenesisBlock: &monaCoinTestNet4GenesisBlock, // no it's not
	GenesisHash:  &monaCoinTestNet4GenesisHash,
	PoWFunction: func(b []byte, height int32) chainhash.Hash {
		scryptBytes, _ := scrypt.Key(b, b, 1024, 1, 1, 32)
		asChainHash, _ := chainhash.NewHash(scryptBytes)
		return *asChainHash
	},
	DiffCalcFunction: diffBitcoin,
	/*
		StartHeader: [80]byte{
				0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0xd9, 0xce, 0xd4, 0xed,
				0x11, 0x30, 0xf7, 0xb7, 0xfa, 0xad, 0x9b, 0xe2,
				0x53, 0x23, 0xff, 0xaf, 0xa3, 0x32, 0x32, 0xa1,
				0x7c, 0x3e, 0xdf, 0x6c, 0xfd, 0x97, 0xbe, 0xe6,
				0xba, 0xfb, 0xdd, 0x97, 0xf6, 0x0b, 0xa1, 0x58,
				0xf0, 0xff, 0x0f, 0x1e, 0xe1, 0x79, 0x04, 0x00,
		},
	*/
	StartHeight:              48384,
	AssumeDiffBefore:         50401,
	FeePerByte:               800,
	PowLimit:                 monaCoinTestNet4PowLimit,
	PowLimitBits:             0x1e0ffff0,
	CoinbaseMaturity:         100,
	SubsidyReductionInterval: 1051200,
	TargetTimespan:           time.Second * 95040, // 1.1 weeks
	TargetTimePerBlock:       time.Second * 90,    // 90 seconds
	RetargetAdjustmentFactor: 4,                   // 25% less, 400% more
	ReduceMinDifficulty:      true,
	MinDiffReductionTime:     0,
	GenerateSupported:        false,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []Checkpoint{
		{1500, newHashFromStr("fa09400a2b4b62e852383d7abd67023e0aaad0e484e190cde9b6d6c079cf2a19")},
		{4000, newHashFromStr("6f554b86ba4379e181f890fb7b7eb14a47a04661b15cc42b3ab8dba48184bd13")},
		{8000, newHashFromStr("32c1f92d0e4968e26577859c31643a60650464953d85dde92f4b434e1af35258")},
		{16000, newHashFromStr("df4b1eec11c3ae392923278c0ec70beb162c5c80e2563208b9ea2b40e72d9128")},
		{32000, newHashFromStr("0e96364a1a7813a3bb7015ea6ee99fcbd76fbc198b13d4eb33915a0d1340f357")},
		{58700, newHashFromStr("668f0cfaca30134792208fd77ffb9f7405aec52b25d75ce5df032342d7b5c216")},
		{80000, newHashFromStr("4666cafb5b29e23acc991a9850ad4f21969695e52794f782fcb18eddc6cc0033")},
		{115000, newHashFromStr("8bd40f1c82cd0572b9e5e5eab29262acbb237597c487903fc53a3d5c07edda1a")},
	},
	// Enforce current block version once majority of the network has
	// upgraded.
	// 51% (51 / 100)
	// Reject previous block versions once a majority of the network has
	// upgraded.
	// 75% (75 / 100)
	BlockEnforceNumRequired: 51,
	BlockRejectNumRequired:  75,
	BlockUpgradeNumToCheck:  100,

	// Mempool parameters
	RelayNonStdTxs: true,

	// Address encoding magics
	PubKeyHashAddrID: 0x6f, // starts with m or n
	ScriptHashAddrID: 0xc4, // starts with 2
	Bech32Prefix:     "tmona",
	PrivateKeyID:     0xef, // starts with 9 7(uncompressed) or c (compressed)

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94}, // starts with tprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xcf}, // starts with tpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 99, // TODO:
}

/*
// LiteCoinTestNet4Params are the parameters for the litecoin test network 4.
var MonaRegNetParams = Params{
	Name:          "litereg",
	NetMagicBytes: 0xdab5bffa,
	DefaultPort:   "19444",
	DNSSeeds:      []string{},

	// Chain parameters
	GenesisBlock: &liteCoinRegTestGenesisBlock, // no it's not
	GenesisHash:  &liteCoinRegTestGenesisHash,
	PoWFunction: func(b []byte, height int32) chainhash.Hash {
		scryptBytes, _ := scrypt.Key(b, b, 1024, 1, 1, 32)
		asChainHash, _ := chainhash.NewHash(scryptBytes)
		return *asChainHash
	},
	DiffCalcFunction:         diffBitcoin,
	FeePerByte:               800,
	PowLimit:                 regressionPowLimit,
	PowLimitBits:             0x207fffff,
	CoinbaseMaturity:         100,
	SubsidyReductionInterval: 150,
	TargetTimespan:           time.Hour * 84,    // 84 hours (3.5 days)
	TargetTimePerBlock:       time.Second * 150, // 150 seconds (2.5 min)
	RetargetAdjustmentFactor: 4,                 // 25% less, 400% more
	ReduceMinDifficulty:      true,
	MinDiffReductionTime:     time.Minute * 10, // ?? unknown
	GenerateSupported:        true,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []Checkpoint{},

	// Enforce current block version once majority of the network has
	// upgraded.
	// 51% (51 / 100)
	// Reject previous block versions once a majority of the network has
	// upgraded.
	// 75% (75 / 100)
	BlockEnforceNumRequired: 51,
	BlockRejectNumRequired:  75,
	BlockUpgradeNumToCheck:  100,

	// Mempool parameters
	RelayNonStdTxs: true,

	// Address encoding magics
	PubKeyHashAddrID: 0x6f, // starts with m or n
	ScriptHashAddrID: 0xc4, // starts with 2
	Bech32Prefix:     "rltc",
	PrivateKeyID:     0xef, // starts with 9 7(uncompressed) or c (compressed)

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94}, // starts with tprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xcf}, // starts with tpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 258, // i dunno
}
*/
// liteCoinTestNet4GenesisHash is the first hash in litecoin testnet4
var monaCoinTestNet4GenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x06, 0x22, 0x6e, 0x46, 0x11, 0x1a, 0x0b, 0x59,
	0xca, 0xaf, 0x12, 0x60, 0x43, 0xeb, 0x5b, 0xbf,
	0x28, 0xc3, 0x4f, 0x3a, 0x5e, 0x33, 0x2a, 0x1f,
	0xc7, 0xb2, 0xb7, 0x3c, 0xf1, 0x88, 0x91, 0x0f,
})

var monaCoinTestNet4MerkleRoot = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xa6, 0x4b, 0xac, 0x07, 0xfe, 0x31, 0x87, 0x7f,
	0x31, 0xd0, 0x32, 0x52, 0x95, 0x3b, 0x3c, 0x32,
	0x39, 0x89, 0x33, 0xaf, 0x7a, 0x72, 0x41, 0x19,
	0xbc, 0x4d, 0x6f, 0xa4, 0xa8, 0x05, 0xe4, 0x35,
})

// monaCoinTestNet4GenesisBlock has is like completely its own thing
var monaCoinTestNet4GenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{}, // empty
		MerkleRoot: monaCoinTestNet4MerkleRoot,
		Timestamp:  time.Unix(0x58bf2dec, 0),
		Bits:       0x1e0ffff0,
		Nonce:      0x0020646c,
	},
	//	Transactions: []*wire.MsgTx{&genesisCoinbaseTx}, // this is wrong... will it break?
}

/*
// ==================== LiteRegNet

// liteCoinRegTestGenesisHash is the first hash in litecoin regtest
var liteCoinRegTestGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x06, 0x22, 0x6e, 0x46, 0x11, 0x1a, 0x0b, 0x59,
	0xca, 0xaf, 0x12, 0x60, 0x43, 0xeb, 0x5b, 0xbf,
	0x28, 0xc3, 0x4f, 0x3a, 0x5e, 0x33, 0x2a, 0x1f,
	0xc7, 0xb2, 0xb7, 0x3c, 0xf1, 0x88, 0x91, 0x0f,
})

// is this the same...?
var liteCoinRegTestMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xd9, 0xce, 0xd4, 0xed, 0x11, 0x30, 0xf7, 0xb7, 0
	0x53, 0x23, 0xff, 0xaf, 0xa3, 0x32, 0x32, 0xa1, 0x7c, 0x3e, 0xdf, 0x6c,
	0xfd, 0x97, 0xbe, 0xe6, 0xba, 0xfb, 0xdd, 0x97,
})

// liteCoinTestNet4GenesisBlock has is like completely its own thing
var liteCoinRegTestGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{}, // empty
		MerkleRoot: liteCoinRegTestMerkleRoot,
		Timestamp:  time.Unix(1296688602, 0), // later
		Bits:       0x207fffff,
		Nonce:      2,
	},
	//	Transactions: []*wire.MsgTx{&genesisCoinbaseTx}, // this is wrong... will it break?
}
*/
