package integration

import (
	"crypto/ecdsa"
	"crypto/rand"
	"github.com/KYVENetwork/chain/app"
	"github.com/KYVENetwork/chain/x/pool"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	stakerstypes "github.com/KYVENetwork/chain/x/stakers/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	tmcrypto "github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmversion "github.com/tendermint/tendermint/proto/tendermint/version"
	"github.com/tendermint/tendermint/version"
	"testing"
	"time"
)

func NewCleanChain() KeeperTestSuite {
	s := KeeperTestSuite{}
	s.SetupTest()
	return s
}

func (suite *KeeperTestSuite) RunTxWithResult(msg sdk.Msg) (*sdk.Result, error) {
	cachedCtx, commit := suite.ctx.CacheContext()
	// TODO generalize for all types of supported messages (modules)
	resp, err := pool.NewHandler(suite.app.PoolKeeper)(cachedCtx, msg)
	if err == nil {
		commit()
		return resp, nil
	}
	return nil, err
}

func (suite *KeeperTestSuite) RunTx(msg sdk.Msg) (success bool) {
	cachedCtx, commit := suite.ctx.CacheContext()
	// TODO generalize for all types of supported messages (modules)
	_, err := pool.NewHandler(suite.app.PoolKeeper)(cachedCtx, msg)
	if err == nil {
		commit()
		return true
	}
	return false
}

func (suite *KeeperTestSuite) RunTxSuccess(t *testing.T, msg sdk.Msg) {
	success := suite.RunTx(msg)
	require.True(t, success)
}

type QueryClients struct {
	poolClient    pooltypes.QueryClient
	stakersClient stakerstypes.QueryClient
}

type KeeperTestSuite struct {
	suite.Suite

	ctx sdk.Context

	app         *app.App
	queries     QueryClients
	address     common.Address
	signer      keyring.Signer
	consAddress sdk.ConsAddress
	validator   stakingtypes.Validator
	denom       string
}

func (suite *KeeperTestSuite) App() *app.App {
	return suite.app
}
func (suite *KeeperTestSuite) Ctx() sdk.Context {
	return suite.ctx
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.app = app.Setup()
	suite.SetupApp()
}

func (suite *KeeperTestSuite) SetupApp() {
	//t := suite.T()

	suite.app = app.Setup()

	suite.denom = "tkyve"

	//fmt.Printf("%s\n", sdk.GetConfig().GetBech32AccountAddrPrefix())
	//fmt.Printf("%s\n", sdk.GetConfig().GetBech32AccountPubPrefix())
	//fmt.Printf("%s\n", sdk.GetConfig().GetBech32ValidatorAddrPrefix())
	//fmt.Printf("%s\n", sdk.GetConfig().GetBech32ValidatorPubPrefix())
	//fmt.Printf("%s\n", sdk.GetConfig().GetBech32ConsensusAddrPrefix())
	//fmt.Printf("%s\n", sdk.GetConfig().GetBech32ConsensusPubPrefix())

	//sdk.GetConfig().SetBech32PrefixForAccount("kyve", "kyvepub")
	//sdk.GetConfig().SetBech32PrefixForValidator("kyvevaloper", "kyvevaloperpub")
	//sdk.GetConfig().SetBech32PrefixForValidator("kyvevalcons", "kyvevalconspub")

	// consensus key
	privKey, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	_ = err
	//require.NoError(t, err)

	addressBytes := tmcrypto.Address(crypto.PubkeyToAddress(privKey.PublicKey).Bytes())
	suite.address = common.BytesToAddress(addressBytes)

	// consensus key
	privKey, err = ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	//require.NoError(t, err)

	ePriv := ed25519.GenPrivKeyFromSecret([]byte{1})
	suite.consAddress = sdk.ConsAddress(ePriv.PubKey().Address())

	suite.ctx = suite.app.BaseApp.NewContext(false, tmproto.Header{
		Height:          1,
		ChainID:         "kyve-test",
		Time:            time.Now().UTC(),
		ProposerAddress: suite.consAddress.Bytes(),

		Version: tmversion.Consensus{
			Block: version.BlockProtocol,
		},
		LastBlockId: tmproto.BlockID{
			Hash: tmhash.Sum([]byte("block_id")),
			PartSetHeader: tmproto.PartSetHeader{
				Total: 11,
				Hash:  tmhash.Sum([]byte("partset_header")),
			},
		},
		AppHash:            tmhash.Sum([]byte("app")),
		DataHash:           tmhash.Sum([]byte("data")),
		EvidenceHash:       tmhash.Sum([]byte("evidence")),
		ValidatorsHash:     tmhash.Sum([]byte("validators")),
		NextValidatorsHash: tmhash.Sum([]byte("next_validators")),
		ConsensusHash:      tmhash.Sum([]byte("consensus")),
		LastResultsHash:    tmhash.Sum([]byte("last_result")),
	})
	suite.RegisterQueryClients()

	mintParams := suite.app.MintKeeper.GetParams(suite.ctx)
	mintParams.MintDenom = suite.denom
	suite.app.MintKeeper.SetParams(suite.ctx, mintParams)

	stakingParams := suite.app.StakingKeeper.GetParams(suite.ctx)
	stakingParams.BondDenom = suite.denom
	suite.app.StakingKeeper.SetParams(suite.ctx, stakingParams)

	// Set Validator
	valAddr := sdk.ValAddress(suite.address.Bytes())
	validator, err := stakingtypes.NewValidator(valAddr, ePriv.PubKey(), stakingtypes.Description{})
	//require.NoError(t, err)
	validator = stakingkeeper.TestingUpdateValidator(suite.app.StakingKeeper, suite.ctx, validator, true)
	suite.app.StakingKeeper.AfterValidatorCreated(suite.ctx, validator.GetOperator())
	err = suite.app.StakingKeeper.SetValidatorByConsAddr(suite.ctx, validator)
	//require.NoError(t, err)
	validators := suite.app.StakingKeeper.GetValidators(suite.ctx, 1)
	suite.validator = validators[0]
}

func (suite *KeeperTestSuite) Commit() {
	suite.CommitAfter(time.Second * 0)
}

func (suite *KeeperTestSuite) CommitAfterSeconds(seconds uint64) {
	suite.CommitAfter(time.Second * time.Duration(seconds))
}

func (suite *KeeperTestSuite) CommitAfter(t time.Duration) {
	header := suite.ctx.BlockHeader()
	suite.app.EndBlock(abci.RequestEndBlock{Height: header.Height})
	_ = suite.app.Commit()

	header.Height += 1
	header.Time = header.Time.Add(t)
	suite.app.BeginBlock(abci.RequestBeginBlock{Header: header})

	suite.ctx = suite.app.BaseApp.NewContext(false, header)

	suite.RegisterQueryClients()
}

func (suite *KeeperTestSuite) RegisterQueryClients() {
	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, suite.app.InterfaceRegistry())

	pooltypes.RegisterQueryServer(queryHelper, suite.app.PoolKeeper)
	suite.queries.poolClient = pooltypes.NewQueryClient(queryHelper)

	stakerstypes.RegisterQueryServer(queryHelper, suite.app.StakersKeeper)
	suite.queries.stakersClient = stakerstypes.NewQueryClient(queryHelper)
}
