package types_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogoprototypes "github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/axelarnetwork/axelar-core/app"
	"github.com/axelarnetwork/axelar-core/testutils"
	"github.com/axelarnetwork/axelar-core/testutils/rand"
	"github.com/axelarnetwork/axelar-core/utils"
	"github.com/axelarnetwork/axelar-core/x/tss/tofnd"
	"github.com/axelarnetwork/axelar-core/x/vote/exported"
	voteMock "github.com/axelarnetwork/axelar-core/x/vote/exported/mock"
	"github.com/axelarnetwork/axelar-core/x/vote/types"
	"github.com/axelarnetwork/axelar-core/x/vote/types/mock"
	. "github.com/axelarnetwork/utils/test"
)

func TestNewTalliedVote(t *testing.T) {
	t.Run("panic on nil data", func(t *testing.T) {
		assert.Panics(t, func() {
			types.NewTalliedVote(rand.ValAddr(), rand.PosI64(), nil)
		})
	})

	t.Run("panic on nil voter", func(t *testing.T) {
		assert.Panics(t, func() {
			types.NewTalliedVote(nil, rand.PosI64(), &gogoprototypes.BoolValue{Value: true})
		})
	})
}

func TestTalliedVote_Marshaling(t *testing.T) {
	encCfg := app.MakeEncodingConfig()
	cdc := encCfg.Codec

	output := tofnd.KeygenOutput{PubKey: []byte("a public key"), GroupRecoverInfo: []byte{0}, PrivateRecoverInfo: []byte{0, 1, 2, 3}}
	data := tofnd.MessageOut_KeygenResult{KeygenResultData: &tofnd.MessageOut_KeygenResult_Data{Data: &output}}
	vote := types.NewTalliedVote(rand.ValAddr(), 23, &data)

	bz := cdc.MustMarshalLengthPrefixed(&vote)
	var actual types.TalliedVote
	cdc.MustUnmarshalLengthPrefixed(bz, &actual)

	assert.Equal(t, vote, actual)

	bz = cdc.MustMarshalJSON(&vote)
	var actual2 types.TalliedVote
	cdc.MustUnmarshalJSON(bz, &actual2)

	assert.Equal(t, vote.Tally, actual2.Tally)
	assert.Equal(t, vote.Data.GetCachedValue(), actual2.Data.GetCachedValue())
}

func TestPoll_Expiry(t *testing.T) {
	setup := func() exported.PollMetadata {
		key := exported.NewPollKey(randomNormalizedStr(5, 20), randomNormalizedStr(5, 20))
		return types.NewPollMetaData(key, types.DefaultParams().DefaultVotingThreshold, []exported.Voter{})
	}
	repeats := 20
	notExpiredStates := []exported.PollState{exported.NonExistent, exported.Pending, exported.Completed, exported.Failed}

	t.Run("poll is not expired", testutils.Func(func(t *testing.T) {
		metadata := setup()
		initialState := notExpiredStates[rand.I64Between(0, int64(len(notExpiredStates)))]
		metadata.State = initialState
		expiry := rand.PosI64()
		metadata.ExpiresAt = expiry

		poll := types.NewPoll(metadata, &mock.StoreMock{})

		assert.True(t, poll.Is(initialState))
	}).Repeat(repeats))

	t.Run("pending poll expires", testutils.Func(func(t *testing.T) {
		metadata := setup()
		metadata.State = exported.Pending
		expiry := rand.I64Between(0, 1000000)
		metadata.ExpiresAt = expiry

		poll := types.NewPoll(metadata, &mock.StoreMock{})

		assert.True(t, poll.Is(exported.Pending))
		assert.True(t, poll.Is(exported.Expired))
	}).Repeat(repeats))

	t.Run("not-pending poll does not expire", testutils.Func(func(t *testing.T) {
		initialState := notExpiredStates[rand.I64GenBetween(0, int64(len(notExpiredStates))).
			Where(func(i int64) bool { return i != int64(exported.Pending) }).
			Next()]
		metadata := setup()
		metadata.State = initialState
		expiry := rand.I64Between(0, 1000000)
		metadata.ExpiresAt = expiry
		poll := types.NewPoll(metadata, &mock.StoreMock{})

		assert.True(t, poll.Is(initialState))
		assert.False(t, poll.Is(exported.Expired))
	}).Repeat(repeats))
}

func TestPoll_Is(t *testing.T) {
	for _, state := range exported.PollState_value {
		poll := types.Poll{PollMetadata: exported.PollMetadata{State: exported.PollState(state)}}

		assert.True(t, poll.Is(exported.PollState(state)))

		for _, otherState := range exported.PollState_value {
			if otherState == state {
				continue
			}
			assert.False(t, poll.Is(exported.PollState(otherState)), "poll: %s, other: %s", poll.State, exported.PollState(otherState))
		}
	}
}

func TestPoll_Vote(t *testing.T) {
	var (
		poll      *types.Poll
		pollStore *mock.StoreMock
	)

	repeats := 20

	givenPoll := Given("poll", func(_ *testing.T) {
		voterCount := rand.I64Between(10, 20)
		voters := make([]exported.Voter, voterCount)
		for i := 0; i < int(voterCount); i++ {
			voters[i] = exported.Voter{
				Validator:   rand.ValAddr(),
				VotingPower: rand.I64Between(10, 100),
			}
		}

		pollKey := exported.NewPollKey(randomNormalizedStr(5, 20), randomNormalizedStr(5, 20))
		pollMetadata := types.NewPollMetaData(pollKey, types.DefaultParams().DefaultVotingThreshold, voters)

		pollStore = &mock.StoreMock{}
		poll = types.NewPoll(pollMetadata, pollStore)
	})

	withState := func(state exported.PollState) func(_ *testing.T) {
		return func(_ *testing.T) {
			poll.State = state
		}
	}

	givenPoll.
		When("poll does not exist", withState(exported.NonExistent)).
		Then("should return error", func(t *testing.T) {
			result, voted, err := poll.Vote(
				rand.Of(poll.Voters...).Validator,
				rand.PosI64(),
				&gogoprototypes.BoolValue{Value: true},
			)

			assert.Nil(t, result)
			assert.False(t, voted)
			assert.ErrorContains(t, err, "poll does not exist")
		}).
		Run(t, repeats)

	givenPoll.
		When("poll is failed", withState(exported.Failed)).
		Then("should do nothing", func(t *testing.T) {
			result, voted, err := poll.Vote(
				rand.Of(poll.Voters...).Validator,
				rand.PosI64(),
				&gogoprototypes.BoolValue{Value: true},
			)

			assert.Nil(t, result)
			assert.False(t, voted)
			assert.NoError(t, err)
		}).
		Run(t, repeats)

	givenPoll.
		When("poll is expired", withState(exported.Expired)).
		Then("should do nothing", func(t *testing.T) {
			result, voted, err := poll.Vote(
				rand.Of(poll.Voters...).Validator,
				rand.PosI64(),
				&gogoprototypes.BoolValue{Value: true},
			)

			assert.Nil(t, result)
			assert.False(t, voted)
			assert.NoError(t, err)
		}).
		Run(t, repeats)

	blockHeight := rand.I64Between(100, 1000)
	givenPoll.
		When("poll is completed", withState(exported.Completed)).
		And().
		When("poll is within its grace period", func(_ *testing.T) {
			poll.GracePeriod = uint64(rand.I64Between(1, blockHeight))
			poll.CompletedAt = rand.I64Between(blockHeight-int64(poll.GracePeriod), blockHeight+1)
		}).
		Then("should allow late vote", func(t *testing.T) {
			voter := rand.Of(poll.Voters...)
			pollStore.HasVotedFunc = func(v sdk.ValAddress) bool { return !v.Equals(voter.Validator) }
			pollStore.SetVoteFunc = func(voter sdk.ValAddress, data codec.ProtoMarshaler, votingPower int64, isLate bool) {}

			result, voted, err := poll.Vote(
				voter.Validator,
				blockHeight,
				&gogoprototypes.BoolValue{Value: true},
			)

			assert.Nil(t, result)
			assert.True(t, voted)
			assert.NoError(t, err)

			assert.Len(t, pollStore.SetVoteCalls(), 1)
			assert.Equal(t, voter.Validator, pollStore.SetVoteCalls()[0].Voter)
			assert.NotNil(t, pollStore.SetVoteCalls()[0].Data)
			assert.Equal(t, voter.VotingPower, pollStore.SetVoteCalls()[0].VotingPower)
			assert.True(t, pollStore.SetVoteCalls()[0].IsLate)
		}).
		Run(t, repeats)

	blockHeight = rand.I64Between(100, 1000)
	givenPoll.
		When("poll is completed", withState(exported.Completed)).
		And().
		When("poll is not within its grace period", func(_ *testing.T) {
			poll.GracePeriod = uint64(rand.I64Between(1, blockHeight))

			if rand.Bools(0.5).Next() {
				poll.State |= exported.Expired
			} else {
				poll.CompletedAt = rand.I64Between(0, blockHeight-int64(poll.GracePeriod))
			}
		}).
		Then("should allow late vote", func(t *testing.T) {
			voter := rand.Of(poll.Voters...)
			pollStore.HasVotedFunc = func(v sdk.ValAddress) bool { return !v.Equals(voter.Validator) }

			result, voted, err := poll.Vote(
				voter.Validator,
				blockHeight,
				&gogoprototypes.BoolValue{Value: true},
			)

			assert.Nil(t, result)
			assert.False(t, voted)
			assert.NoError(t, err)
		}).
		Run(t, repeats)

	var voter exported.Voter
	givenPoll.
		When("poll is pending", withState(exported.Pending)).
		And().
		When("voter is not eligible", func(t *testing.T) {
			if rand.Bools(0.5).Next() {
				pollStore.HasVotedFunc = func(v sdk.ValAddress) bool { return false }
				voter = exported.Voter{Validator: rand.ValAddr()}
			} else {
				pollStore.HasVotedFunc = func(v sdk.ValAddress) bool { return v.Equals(voter.Validator) }
			}
		}).
		Then("should return error", func(t *testing.T) {
			result, voted, err := poll.Vote(
				voter.Validator,
				rand.PosI64(),
				&gogoprototypes.BoolValue{Value: true},
			)

			assert.Nil(t, result)
			assert.False(t, voted)
			assert.ErrorContains(t, err, "is not eligible")
		}).
		Run(t, repeats)

}

func TestPoll_Initialize(t *testing.T) {
	var (
		previousPoll exported.Poll
	)

	store := &mock.StoreMock{
		GetPollFunc:     func(exported.PollKey) exported.Poll { return previousPoll },
		SetMetadataFunc: func(exported.PollMetadata) {},
		DeletePollFunc:  func() {},
	}

	repeats := 20

	testCases := []struct {
		label        string
		previousPoll exported.Poll
		expectError  bool
	}{
		{"poll can be overridden", &voteMock.PollMock{DeleteFunc: func() error { return nil }}, false},
		{"poll can not be overridden", &voteMock.PollMock{DeleteFunc: func() error { return fmt.Errorf("no delete") }}, true},
	}

	for _, testCase := range testCases {
		t.Run(testCase.label, testutils.Func(func(t *testing.T) {
			previousPoll = testCase.previousPoll
			poll := types.NewPoll(newRandomPollMetadata(), store).WithLogger(log.TestingLogger())

			if testCase.expectError {
				assert.Error(t, poll.Initialize(rand.PosI64()))
			} else {
				assert.NoError(t, poll.Initialize(rand.PosI64()))
			}
		}).Repeat(repeats))
	}
}

func TestPoll_Delete(t *testing.T) {
	var store *mock.StoreMock
	setup := func(pollState exported.PollState) exported.Poll {
		store = &mock.StoreMock{DeletePollFunc: func() {}}
		metadata := newRandomPollMetadata()
		metadata.State = pollState

		return types.NewPoll(metadata, store).WithLogger(log.TestingLogger())
	}

	t.Run("nonexistent", func(t *testing.T) {
		poll := setup(exported.NonExistent)
		assert.NoError(t, poll.Delete())
		assert.Len(t, store.DeletePollCalls(), 0)
	})

	testCases := []struct {
		label            string
		pollState        exported.PollState
		deleteSuccessful bool
	}{
		{"pending", exported.Pending, false},
		{"completed", exported.Completed, false},
		{"failed", exported.Failed, false},
		{"expired", exported.Expired, false},
		{"allow override", exported.AllowOverride, true},
	}

	for _, testCase := range testCases {
		t.Run(testCase.label, func(t *testing.T) {
			poll := setup(testCase.pollState)

			if testCase.deleteSuccessful {
				assert.NoError(t, poll.Delete())
				assert.Len(t, store.DeletePollCalls(), 1)
			} else {
				assert.Error(t, poll.Delete())
				assert.Len(t, store.DeletePollCalls(), 0)
			}
		})
	}
}

func randomEvenVotingPowers() map[string]int64 {
	votingPowers := make(map[string]int64)

	total := sdk.ZeroInt()
	for i := 0; i < int(rand.I64Between(1, 20)); i++ {
		addr := rand.ValAddr()
		votingPower := rand.I64Between(1, 100)
		votingPowers[addr.String()] = votingPower
		total = total.AddRaw(votingPower)
	}

	// redraw voting power if any one votingPower is greater than half of the total
	for _, votingPower := range votingPowers {
		if total.QuoRaw(2).LT(sdk.NewInt(votingPower)) {
			return randomEvenVotingPowers()
		}
	}

	return votingPowers
}

func getValues(m map[string]types.TalliedVote) []types.TalliedVote {
	votes := make([]types.TalliedVote, 0, len(m))
	for _, vote := range m {
		votes = append(votes, vote)
	}
	return votes
}

func newRandomPollMetadata() exported.PollMetadata {
	key := exported.NewPollKey(randomNormalizedStr(5, 20), randomNormalizedStr(5, 20))
	poll := types.NewPollMetaData(key, types.DefaultParams().DefaultVotingThreshold, []exported.Voter{})
	poll.ExpiresAt = rand.I64Between(1, 1000000)
	return poll
}

func randomNormalizedStr(min, max int) string {
	return strings.ReplaceAll(utils.NormalizeString(rand.StrBetween(min, max)), utils.DefaultDelimiter, "-")
}
