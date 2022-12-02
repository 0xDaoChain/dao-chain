package server

import (
	"github.com/0xDaoChain/dao-chain/consensus"
	consensusDev "github.com/0xDaoChain/dao-chain/consensus/dev"
	consensusDummy "github.com/0xDaoChain/dao-chain/consensus/dummy"
	consensusIBFT "github.com/0xDaoChain/dao-chain/consensus/ibft"
	"github.com/0xDaoChain/dao-chain/secrets"
	"github.com/0xDaoChain/dao-chain/secrets/awsssm"
	"github.com/0xDaoChain/dao-chain/secrets/gcpssm"
	"github.com/0xDaoChain/dao-chain/secrets/hashicorpvault"
	"github.com/0xDaoChain/dao-chain/secrets/local"
)

type ConsensusType string

const (
	DevConsensus   ConsensusType = "dev"
	IBFTConsensus  ConsensusType = "ibft"
	DummyConsensus ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:   consensusDev.Factory,
	IBFTConsensus:  consensusIBFT.Factory,
	DummyConsensus: consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
