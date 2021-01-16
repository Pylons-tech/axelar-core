// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"crypto/ecdsa"
	"github.com/axelarnetwork/axelar-core/x/balance/exported"
	"github.com/axelarnetwork/axelar-core/x/bitcoin/types"
	snapshot "github.com/axelarnetwork/axelar-core/x/snapshot/exported"
	tss "github.com/axelarnetwork/axelar-core/x/tss/exported"
	voting "github.com/axelarnetwork/axelar-core/x/vote/exported"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sync"
)

// Ensure, that VoterMock does implement types.Voter.
// If this is not the case, regenerate this file with moq.
var _ types.Voter = &VoterMock{}

// VoterMock is a mock implementation of types.Voter.
//
//     func TestSomethingThatUsesVoter(t *testing.T) {
//
//         // make and configure a mocked types.Voter
//         mockedVoter := &VoterMock{
//             DeletePollFunc: func(ctx sdk.Context, poll voting.PollMeta)  {
// 	               panic("mock out the DeletePoll method")
//             },
//             InitPollFunc: func(ctx sdk.Context, poll voting.PollMeta) error {
// 	               panic("mock out the InitPoll method")
//             },
//             RecordVoteFunc: func(ctx sdk.Context, vote voting.MsgVote) error {
// 	               panic("mock out the RecordVote method")
//             },
//             ResultFunc: func(ctx sdk.Context, poll voting.PollMeta) voting.VotingData {
// 	               panic("mock out the Result method")
//             },
//             TallyVoteFunc: func(ctx sdk.Context, vote voting.MsgVote) error {
// 	               panic("mock out the TallyVote method")
//             },
//         }
//
//         // use mockedVoter in code that requires types.Voter
//         // and then make assertions.
//
//     }
type VoterMock struct {
	// DeletePollFunc mocks the DeletePoll method.
	DeletePollFunc func(ctx sdk.Context, poll voting.PollMeta)

	// InitPollFunc mocks the InitPoll method.
	InitPollFunc func(ctx sdk.Context, poll voting.PollMeta) error

	// RecordVoteFunc mocks the RecordVote method.
	RecordVoteFunc func(ctx sdk.Context, vote voting.MsgVote) error

	// ResultFunc mocks the Result method.
	ResultFunc func(ctx sdk.Context, poll voting.PollMeta) voting.VotingData

	// TallyVoteFunc mocks the TallyVote method.
	TallyVoteFunc func(ctx sdk.Context, vote voting.MsgVote) error

	// calls tracks calls to the methods.
	calls struct {
		// DeletePoll holds details about calls to the DeletePoll method.
		DeletePoll []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Poll is the poll argument value.
			Poll voting.PollMeta
		}
		// InitPoll holds details about calls to the InitPoll method.
		InitPoll []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Poll is the poll argument value.
			Poll voting.PollMeta
		}
		// RecordVote holds details about calls to the RecordVote method.
		RecordVote []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Vote is the vote argument value.
			Vote voting.MsgVote
		}
		// Result holds details about calls to the Result method.
		Result []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Poll is the poll argument value.
			Poll voting.PollMeta
		}
		// TallyVote holds details about calls to the TallyVote method.
		TallyVote []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Vote is the vote argument value.
			Vote voting.MsgVote
		}
	}
	lockDeletePoll sync.RWMutex
	lockInitPoll   sync.RWMutex
	lockRecordVote sync.RWMutex
	lockResult     sync.RWMutex
	lockTallyVote  sync.RWMutex
}

// DeletePoll calls DeletePollFunc.
func (mock *VoterMock) DeletePoll(ctx sdk.Context, poll voting.PollMeta) {
	if mock.DeletePollFunc == nil {
		panic("VoterMock.DeletePollFunc: method is nil but Voter.DeletePoll was just called")
	}
	callInfo := struct {
		Ctx  sdk.Context
		Poll voting.PollMeta
	}{
		Ctx:  ctx,
		Poll: poll,
	}
	mock.lockDeletePoll.Lock()
	mock.calls.DeletePoll = append(mock.calls.DeletePoll, callInfo)
	mock.lockDeletePoll.Unlock()
	mock.DeletePollFunc(ctx, poll)
}

// DeletePollCalls gets all the calls that were made to DeletePoll.
// Check the length with:
//     len(mockedVoter.DeletePollCalls())
func (mock *VoterMock) DeletePollCalls() []struct {
	Ctx  sdk.Context
	Poll voting.PollMeta
} {
	var calls []struct {
		Ctx  sdk.Context
		Poll voting.PollMeta
	}
	mock.lockDeletePoll.RLock()
	calls = mock.calls.DeletePoll
	mock.lockDeletePoll.RUnlock()
	return calls
}

// InitPoll calls InitPollFunc.
func (mock *VoterMock) InitPoll(ctx sdk.Context, poll voting.PollMeta) error {
	if mock.InitPollFunc == nil {
		panic("VoterMock.InitPollFunc: method is nil but Voter.InitPoll was just called")
	}
	callInfo := struct {
		Ctx  sdk.Context
		Poll voting.PollMeta
	}{
		Ctx:  ctx,
		Poll: poll,
	}
	mock.lockInitPoll.Lock()
	mock.calls.InitPoll = append(mock.calls.InitPoll, callInfo)
	mock.lockInitPoll.Unlock()
	return mock.InitPollFunc(ctx, poll)
}

// InitPollCalls gets all the calls that were made to InitPoll.
// Check the length with:
//     len(mockedVoter.InitPollCalls())
func (mock *VoterMock) InitPollCalls() []struct {
	Ctx  sdk.Context
	Poll voting.PollMeta
} {
	var calls []struct {
		Ctx  sdk.Context
		Poll voting.PollMeta
	}
	mock.lockInitPoll.RLock()
	calls = mock.calls.InitPoll
	mock.lockInitPoll.RUnlock()
	return calls
}

// RecordVote calls RecordVoteFunc.
func (mock *VoterMock) RecordVote(ctx sdk.Context, vote voting.MsgVote) error {
	if mock.RecordVoteFunc == nil {
		panic("VoterMock.RecordVoteFunc: method is nil but Voter.RecordVote was just called")
	}
	callInfo := struct {
		Ctx  sdk.Context
		Vote voting.MsgVote
	}{
		Ctx:  ctx,
		Vote: vote,
	}
	mock.lockRecordVote.Lock()
	mock.calls.RecordVote = append(mock.calls.RecordVote, callInfo)
	mock.lockRecordVote.Unlock()
	return mock.RecordVoteFunc(ctx, vote)
}

// RecordVoteCalls gets all the calls that were made to RecordVote.
// Check the length with:
//     len(mockedVoter.RecordVoteCalls())
func (mock *VoterMock) RecordVoteCalls() []struct {
	Ctx  sdk.Context
	Vote voting.MsgVote
} {
	var calls []struct {
		Ctx  sdk.Context
		Vote voting.MsgVote
	}
	mock.lockRecordVote.RLock()
	calls = mock.calls.RecordVote
	mock.lockRecordVote.RUnlock()
	return calls
}

// Result calls ResultFunc.
func (mock *VoterMock) Result(ctx sdk.Context, poll voting.PollMeta) voting.VotingData {
	if mock.ResultFunc == nil {
		panic("VoterMock.ResultFunc: method is nil but Voter.Result was just called")
	}
	callInfo := struct {
		Ctx  sdk.Context
		Poll voting.PollMeta
	}{
		Ctx:  ctx,
		Poll: poll,
	}
	mock.lockResult.Lock()
	mock.calls.Result = append(mock.calls.Result, callInfo)
	mock.lockResult.Unlock()
	return mock.ResultFunc(ctx, poll)
}

// ResultCalls gets all the calls that were made to Result.
// Check the length with:
//     len(mockedVoter.ResultCalls())
func (mock *VoterMock) ResultCalls() []struct {
	Ctx  sdk.Context
	Poll voting.PollMeta
} {
	var calls []struct {
		Ctx  sdk.Context
		Poll voting.PollMeta
	}
	mock.lockResult.RLock()
	calls = mock.calls.Result
	mock.lockResult.RUnlock()
	return calls
}

// TallyVote calls TallyVoteFunc.
func (mock *VoterMock) TallyVote(ctx sdk.Context, vote voting.MsgVote) error {
	if mock.TallyVoteFunc == nil {
		panic("VoterMock.TallyVoteFunc: method is nil but Voter.TallyVote was just called")
	}
	callInfo := struct {
		Ctx  sdk.Context
		Vote voting.MsgVote
	}{
		Ctx:  ctx,
		Vote: vote,
	}
	mock.lockTallyVote.Lock()
	mock.calls.TallyVote = append(mock.calls.TallyVote, callInfo)
	mock.lockTallyVote.Unlock()
	return mock.TallyVoteFunc(ctx, vote)
}

// TallyVoteCalls gets all the calls that were made to TallyVote.
// Check the length with:
//     len(mockedVoter.TallyVoteCalls())
func (mock *VoterMock) TallyVoteCalls() []struct {
	Ctx  sdk.Context
	Vote voting.MsgVote
} {
	var calls []struct {
		Ctx  sdk.Context
		Vote voting.MsgVote
	}
	mock.lockTallyVote.RLock()
	calls = mock.calls.TallyVote
	mock.lockTallyVote.RUnlock()
	return calls
}

// Ensure, that SignerMock does implement types.Signer.
// If this is not the case, regenerate this file with moq.
var _ types.Signer = &SignerMock{}

// SignerMock is a mock implementation of types.Signer.
//
//     func TestSomethingThatUsesSigner(t *testing.T) {
//
//         // make and configure a mocked types.Signer
//         mockedSigner := &SignerMock{
//             GetCurrentMasterKeyFunc: func(ctx sdk.Context, chain exported.Chain) (ecdsa.PublicKey, bool) {
// 	               panic("mock out the GetCurrentMasterKey method")
//             },
//             GetCurrentMasterKeyIDFunc: func(ctx sdk.Context, chain exported.Chain) (string, bool) {
// 	               panic("mock out the GetCurrentMasterKeyID method")
//             },
//             GetKeyFunc: func(ctx sdk.Context, keyID string) (ecdsa.PublicKey, bool) {
// 	               panic("mock out the GetKey method")
//             },
//             GetKeyForSigIDFunc: func(ctx sdk.Context, sigID string) (ecdsa.PublicKey, bool) {
// 	               panic("mock out the GetKeyForSigID method")
//             },
//             GetNextMasterKeyFunc: func(ctx sdk.Context, chain exported.Chain) (ecdsa.PublicKey, bool) {
// 	               panic("mock out the GetNextMasterKey method")
//             },
//             GetSigFunc: func(ctx sdk.Context, sigID string) (tss.Signature, bool) {
// 	               panic("mock out the GetSig method")
//             },
//             GetSnapshotRoundForKeyIDFunc: func(ctx sdk.Context, keyID string) (int64, bool) {
// 	               panic("mock out the GetSnapshotRoundForKeyID method")
//             },
//             StartSignFunc: func(ctx sdk.Context, keyID string, sigID string, msg []byte, validators []snapshot.Validator) error {
// 	               panic("mock out the StartSign method")
//             },
//         }
//
//         // use mockedSigner in code that requires types.Signer
//         // and then make assertions.
//
//     }
type SignerMock struct {
	// GetCurrentMasterKeyFunc mocks the GetCurrentMasterKey method.
	GetCurrentMasterKeyFunc func(ctx sdk.Context, chain exported.Chain) (ecdsa.PublicKey, bool)

	// GetCurrentMasterKeyIDFunc mocks the GetCurrentMasterKeyID method.
	GetCurrentMasterKeyIDFunc func(ctx sdk.Context, chain exported.Chain) (string, bool)

	// GetKeyFunc mocks the GetKey method.
	GetKeyFunc func(ctx sdk.Context, keyID string) (ecdsa.PublicKey, bool)

	// GetKeyForSigIDFunc mocks the GetKeyForSigID method.
	GetKeyForSigIDFunc func(ctx sdk.Context, sigID string) (ecdsa.PublicKey, bool)

	// GetNextMasterKeyFunc mocks the GetNextMasterKey method.
	GetNextMasterKeyFunc func(ctx sdk.Context, chain exported.Chain) (ecdsa.PublicKey, bool)

	// GetSigFunc mocks the GetSig method.
	GetSigFunc func(ctx sdk.Context, sigID string) (tss.Signature, bool)

	// GetSnapshotRoundForKeyIDFunc mocks the GetSnapshotRoundForKeyID method.
	GetSnapshotRoundForKeyIDFunc func(ctx sdk.Context, keyID string) (int64, bool)

	// StartSignFunc mocks the StartSign method.
	StartSignFunc func(ctx sdk.Context, keyID string, sigID string, msg []byte, validators []snapshot.Validator) error

	// calls tracks calls to the methods.
	calls struct {
		// GetCurrentMasterKey holds details about calls to the GetCurrentMasterKey method.
		GetCurrentMasterKey []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Chain is the chain argument value.
			Chain exported.Chain
		}
		// GetCurrentMasterKeyID holds details about calls to the GetCurrentMasterKeyID method.
		GetCurrentMasterKeyID []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Chain is the chain argument value.
			Chain exported.Chain
		}
		// GetKey holds details about calls to the GetKey method.
		GetKey []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// KeyID is the keyID argument value.
			KeyID string
		}
		// GetKeyForSigID holds details about calls to the GetKeyForSigID method.
		GetKeyForSigID []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// SigID is the sigID argument value.
			SigID string
		}
		// GetNextMasterKey holds details about calls to the GetNextMasterKey method.
		GetNextMasterKey []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Chain is the chain argument value.
			Chain exported.Chain
		}
		// GetSig holds details about calls to the GetSig method.
		GetSig []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// SigID is the sigID argument value.
			SigID string
		}
		// GetSnapshotRoundForKeyID holds details about calls to the GetSnapshotRoundForKeyID method.
		GetSnapshotRoundForKeyID []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// KeyID is the keyID argument value.
			KeyID string
		}
		// StartSign holds details about calls to the StartSign method.
		StartSign []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// KeyID is the keyID argument value.
			KeyID string
			// SigID is the sigID argument value.
			SigID string
			// Msg is the msg argument value.
			Msg []byte
			// Validators is the validators argument value.
			Validators []snapshot.Validator
		}
	}
	lockGetCurrentMasterKey      sync.RWMutex
	lockGetCurrentMasterKeyID    sync.RWMutex
	lockGetKey                   sync.RWMutex
	lockGetKeyForSigID           sync.RWMutex
	lockGetNextMasterKey         sync.RWMutex
	lockGetSig                   sync.RWMutex
	lockGetSnapshotRoundForKeyID sync.RWMutex
	lockStartSign                sync.RWMutex
}

// GetCurrentMasterKey calls GetCurrentMasterKeyFunc.
func (mock *SignerMock) GetCurrentMasterKey(ctx sdk.Context, chain exported.Chain) (ecdsa.PublicKey, bool) {
	if mock.GetCurrentMasterKeyFunc == nil {
		panic("SignerMock.GetCurrentMasterKeyFunc: method is nil but Signer.GetCurrentMasterKey was just called")
	}
	callInfo := struct {
		Ctx   sdk.Context
		Chain exported.Chain
	}{
		Ctx:   ctx,
		Chain: chain,
	}
	mock.lockGetCurrentMasterKey.Lock()
	mock.calls.GetCurrentMasterKey = append(mock.calls.GetCurrentMasterKey, callInfo)
	mock.lockGetCurrentMasterKey.Unlock()
	return mock.GetCurrentMasterKeyFunc(ctx, chain)
}

// GetCurrentMasterKeyCalls gets all the calls that were made to GetCurrentMasterKey.
// Check the length with:
//     len(mockedSigner.GetCurrentMasterKeyCalls())
func (mock *SignerMock) GetCurrentMasterKeyCalls() []struct {
	Ctx   sdk.Context
	Chain exported.Chain
} {
	var calls []struct {
		Ctx   sdk.Context
		Chain exported.Chain
	}
	mock.lockGetCurrentMasterKey.RLock()
	calls = mock.calls.GetCurrentMasterKey
	mock.lockGetCurrentMasterKey.RUnlock()
	return calls
}

// GetCurrentMasterKeyID calls GetCurrentMasterKeyIDFunc.
func (mock *SignerMock) GetCurrentMasterKeyID(ctx sdk.Context, chain exported.Chain) (string, bool) {
	if mock.GetCurrentMasterKeyIDFunc == nil {
		panic("SignerMock.GetCurrentMasterKeyIDFunc: method is nil but Signer.GetCurrentMasterKeyID was just called")
	}
	callInfo := struct {
		Ctx   sdk.Context
		Chain exported.Chain
	}{
		Ctx:   ctx,
		Chain: chain,
	}
	mock.lockGetCurrentMasterKeyID.Lock()
	mock.calls.GetCurrentMasterKeyID = append(mock.calls.GetCurrentMasterKeyID, callInfo)
	mock.lockGetCurrentMasterKeyID.Unlock()
	return mock.GetCurrentMasterKeyIDFunc(ctx, chain)
}

// GetCurrentMasterKeyIDCalls gets all the calls that were made to GetCurrentMasterKeyID.
// Check the length with:
//     len(mockedSigner.GetCurrentMasterKeyIDCalls())
func (mock *SignerMock) GetCurrentMasterKeyIDCalls() []struct {
	Ctx   sdk.Context
	Chain exported.Chain
} {
	var calls []struct {
		Ctx   sdk.Context
		Chain exported.Chain
	}
	mock.lockGetCurrentMasterKeyID.RLock()
	calls = mock.calls.GetCurrentMasterKeyID
	mock.lockGetCurrentMasterKeyID.RUnlock()
	return calls
}

// GetKey calls GetKeyFunc.
func (mock *SignerMock) GetKey(ctx sdk.Context, keyID string) (ecdsa.PublicKey, bool) {
	if mock.GetKeyFunc == nil {
		panic("SignerMock.GetKeyFunc: method is nil but Signer.GetKey was just called")
	}
	callInfo := struct {
		Ctx   sdk.Context
		KeyID string
	}{
		Ctx:   ctx,
		KeyID: keyID,
	}
	mock.lockGetKey.Lock()
	mock.calls.GetKey = append(mock.calls.GetKey, callInfo)
	mock.lockGetKey.Unlock()
	return mock.GetKeyFunc(ctx, keyID)
}

// GetKeyCalls gets all the calls that were made to GetKey.
// Check the length with:
//     len(mockedSigner.GetKeyCalls())
func (mock *SignerMock) GetKeyCalls() []struct {
	Ctx   sdk.Context
	KeyID string
} {
	var calls []struct {
		Ctx   sdk.Context
		KeyID string
	}
	mock.lockGetKey.RLock()
	calls = mock.calls.GetKey
	mock.lockGetKey.RUnlock()
	return calls
}

// GetKeyForSigID calls GetKeyForSigIDFunc.
func (mock *SignerMock) GetKeyForSigID(ctx sdk.Context, sigID string) (ecdsa.PublicKey, bool) {
	if mock.GetKeyForSigIDFunc == nil {
		panic("SignerMock.GetKeyForSigIDFunc: method is nil but Signer.GetKeyForSigID was just called")
	}
	callInfo := struct {
		Ctx   sdk.Context
		SigID string
	}{
		Ctx:   ctx,
		SigID: sigID,
	}
	mock.lockGetKeyForSigID.Lock()
	mock.calls.GetKeyForSigID = append(mock.calls.GetKeyForSigID, callInfo)
	mock.lockGetKeyForSigID.Unlock()
	return mock.GetKeyForSigIDFunc(ctx, sigID)
}

// GetKeyForSigIDCalls gets all the calls that were made to GetKeyForSigID.
// Check the length with:
//     len(mockedSigner.GetKeyForSigIDCalls())
func (mock *SignerMock) GetKeyForSigIDCalls() []struct {
	Ctx   sdk.Context
	SigID string
} {
	var calls []struct {
		Ctx   sdk.Context
		SigID string
	}
	mock.lockGetKeyForSigID.RLock()
	calls = mock.calls.GetKeyForSigID
	mock.lockGetKeyForSigID.RUnlock()
	return calls
}

// GetNextMasterKey calls GetNextMasterKeyFunc.
func (mock *SignerMock) GetNextMasterKey(ctx sdk.Context, chain exported.Chain) (ecdsa.PublicKey, bool) {
	if mock.GetNextMasterKeyFunc == nil {
		panic("SignerMock.GetNextMasterKeyFunc: method is nil but Signer.GetNextMasterKey was just called")
	}
	callInfo := struct {
		Ctx   sdk.Context
		Chain exported.Chain
	}{
		Ctx:   ctx,
		Chain: chain,
	}
	mock.lockGetNextMasterKey.Lock()
	mock.calls.GetNextMasterKey = append(mock.calls.GetNextMasterKey, callInfo)
	mock.lockGetNextMasterKey.Unlock()
	return mock.GetNextMasterKeyFunc(ctx, chain)
}

// GetNextMasterKeyCalls gets all the calls that were made to GetNextMasterKey.
// Check the length with:
//     len(mockedSigner.GetNextMasterKeyCalls())
func (mock *SignerMock) GetNextMasterKeyCalls() []struct {
	Ctx   sdk.Context
	Chain exported.Chain
} {
	var calls []struct {
		Ctx   sdk.Context
		Chain exported.Chain
	}
	mock.lockGetNextMasterKey.RLock()
	calls = mock.calls.GetNextMasterKey
	mock.lockGetNextMasterKey.RUnlock()
	return calls
}

// GetSig calls GetSigFunc.
func (mock *SignerMock) GetSig(ctx sdk.Context, sigID string) (tss.Signature, bool) {
	if mock.GetSigFunc == nil {
		panic("SignerMock.GetSigFunc: method is nil but Signer.GetSig was just called")
	}
	callInfo := struct {
		Ctx   sdk.Context
		SigID string
	}{
		Ctx:   ctx,
		SigID: sigID,
	}
	mock.lockGetSig.Lock()
	mock.calls.GetSig = append(mock.calls.GetSig, callInfo)
	mock.lockGetSig.Unlock()
	return mock.GetSigFunc(ctx, sigID)
}

// GetSigCalls gets all the calls that were made to GetSig.
// Check the length with:
//     len(mockedSigner.GetSigCalls())
func (mock *SignerMock) GetSigCalls() []struct {
	Ctx   sdk.Context
	SigID string
} {
	var calls []struct {
		Ctx   sdk.Context
		SigID string
	}
	mock.lockGetSig.RLock()
	calls = mock.calls.GetSig
	mock.lockGetSig.RUnlock()
	return calls
}

// GetSnapshotRoundForKeyID calls GetSnapshotRoundForKeyIDFunc.
func (mock *SignerMock) GetSnapshotRoundForKeyID(ctx sdk.Context, keyID string) (int64, bool) {
	if mock.GetSnapshotRoundForKeyIDFunc == nil {
		panic("SignerMock.GetSnapshotRoundForKeyIDFunc: method is nil but Signer.GetSnapshotRoundForKeyID was just called")
	}
	callInfo := struct {
		Ctx   sdk.Context
		KeyID string
	}{
		Ctx:   ctx,
		KeyID: keyID,
	}
	mock.lockGetSnapshotRoundForKeyID.Lock()
	mock.calls.GetSnapshotRoundForKeyID = append(mock.calls.GetSnapshotRoundForKeyID, callInfo)
	mock.lockGetSnapshotRoundForKeyID.Unlock()
	return mock.GetSnapshotRoundForKeyIDFunc(ctx, keyID)
}

// GetSnapshotRoundForKeyIDCalls gets all the calls that were made to GetSnapshotRoundForKeyID.
// Check the length with:
//     len(mockedSigner.GetSnapshotRoundForKeyIDCalls())
func (mock *SignerMock) GetSnapshotRoundForKeyIDCalls() []struct {
	Ctx   sdk.Context
	KeyID string
} {
	var calls []struct {
		Ctx   sdk.Context
		KeyID string
	}
	mock.lockGetSnapshotRoundForKeyID.RLock()
	calls = mock.calls.GetSnapshotRoundForKeyID
	mock.lockGetSnapshotRoundForKeyID.RUnlock()
	return calls
}

// StartSign calls StartSignFunc.
func (mock *SignerMock) StartSign(ctx sdk.Context, keyID string, sigID string, msg []byte, validators []snapshot.Validator) error {
	if mock.StartSignFunc == nil {
		panic("SignerMock.StartSignFunc: method is nil but Signer.StartSign was just called")
	}
	callInfo := struct {
		Ctx        sdk.Context
		KeyID      string
		SigID      string
		Msg        []byte
		Validators []snapshot.Validator
	}{
		Ctx:        ctx,
		KeyID:      keyID,
		SigID:      sigID,
		Msg:        msg,
		Validators: validators,
	}
	mock.lockStartSign.Lock()
	mock.calls.StartSign = append(mock.calls.StartSign, callInfo)
	mock.lockStartSign.Unlock()
	return mock.StartSignFunc(ctx, keyID, sigID, msg, validators)
}

// StartSignCalls gets all the calls that were made to StartSign.
// Check the length with:
//     len(mockedSigner.StartSignCalls())
func (mock *SignerMock) StartSignCalls() []struct {
	Ctx        sdk.Context
	KeyID      string
	SigID      string
	Msg        []byte
	Validators []snapshot.Validator
} {
	var calls []struct {
		Ctx        sdk.Context
		KeyID      string
		SigID      string
		Msg        []byte
		Validators []snapshot.Validator
	}
	mock.lockStartSign.RLock()
	calls = mock.calls.StartSign
	mock.lockStartSign.RUnlock()
	return calls
}

// Ensure, that BalancerMock does implement types.Balancer.
// If this is not the case, regenerate this file with moq.
var _ types.Balancer = &BalancerMock{}

// BalancerMock is a mock implementation of types.Balancer.
//
//     func TestSomethingThatUsesBalancer(t *testing.T) {
//
//         // make and configure a mocked types.Balancer
//         mockedBalancer := &BalancerMock{
//             GetRecipientFunc: func(ctx sdk.Context, sender exported.CrossChainAddress) (exported.CrossChainAddress, bool) {
// 	               panic("mock out the GetRecipient method")
//             },
//             LinkAddressesFunc: func(ctx sdk.Context, sender exported.CrossChainAddress, recipient exported.CrossChainAddress)  {
// 	               panic("mock out the LinkAddresses method")
//             },
//             PrepareForTransferFunc: func(ctx sdk.Context, sender exported.CrossChainAddress, amount sdk.Coin) error {
// 	               panic("mock out the PrepareForTransfer method")
//             },
//         }
//
//         // use mockedBalancer in code that requires types.Balancer
//         // and then make assertions.
//
//     }
type BalancerMock struct {
	// GetRecipientFunc mocks the GetRecipient method.
	GetRecipientFunc func(ctx sdk.Context, sender exported.CrossChainAddress) (exported.CrossChainAddress, bool)

	// LinkAddressesFunc mocks the LinkAddresses method.
	LinkAddressesFunc func(ctx sdk.Context, sender exported.CrossChainAddress, recipient exported.CrossChainAddress)

	// PrepareForTransferFunc mocks the PrepareForTransfer method.
	PrepareForTransferFunc func(ctx sdk.Context, sender exported.CrossChainAddress, amount sdk.Coin) error

	// calls tracks calls to the methods.
	calls struct {
		// GetRecipient holds details about calls to the GetRecipient method.
		GetRecipient []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Sender is the sender argument value.
			Sender exported.CrossChainAddress
		}
		// LinkAddresses holds details about calls to the LinkAddresses method.
		LinkAddresses []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Sender is the sender argument value.
			Sender exported.CrossChainAddress
			// Recipient is the recipient argument value.
			Recipient exported.CrossChainAddress
		}
		// PrepareForTransfer holds details about calls to the PrepareForTransfer method.
		PrepareForTransfer []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Sender is the sender argument value.
			Sender exported.CrossChainAddress
			// Amount is the amount argument value.
			Amount sdk.Coin
		}
	}
	lockGetRecipient       sync.RWMutex
	lockLinkAddresses      sync.RWMutex
	lockPrepareForTransfer sync.RWMutex
}

// GetRecipient calls GetRecipientFunc.
func (mock *BalancerMock) GetRecipient(ctx sdk.Context, sender exported.CrossChainAddress) (exported.CrossChainAddress, bool) {
	if mock.GetRecipientFunc == nil {
		panic("BalancerMock.GetRecipientFunc: method is nil but Balancer.GetRecipient was just called")
	}
	callInfo := struct {
		Ctx    sdk.Context
		Sender exported.CrossChainAddress
	}{
		Ctx:    ctx,
		Sender: sender,
	}
	mock.lockGetRecipient.Lock()
	mock.calls.GetRecipient = append(mock.calls.GetRecipient, callInfo)
	mock.lockGetRecipient.Unlock()
	return mock.GetRecipientFunc(ctx, sender)
}

// GetRecipientCalls gets all the calls that were made to GetRecipient.
// Check the length with:
//     len(mockedBalancer.GetRecipientCalls())
func (mock *BalancerMock) GetRecipientCalls() []struct {
	Ctx    sdk.Context
	Sender exported.CrossChainAddress
} {
	var calls []struct {
		Ctx    sdk.Context
		Sender exported.CrossChainAddress
	}
	mock.lockGetRecipient.RLock()
	calls = mock.calls.GetRecipient
	mock.lockGetRecipient.RUnlock()
	return calls
}

// LinkAddresses calls LinkAddressesFunc.
func (mock *BalancerMock) LinkAddresses(ctx sdk.Context, sender exported.CrossChainAddress, recipient exported.CrossChainAddress) {
	if mock.LinkAddressesFunc == nil {
		panic("BalancerMock.LinkAddressesFunc: method is nil but Balancer.LinkAddresses was just called")
	}
	callInfo := struct {
		Ctx       sdk.Context
		Sender    exported.CrossChainAddress
		Recipient exported.CrossChainAddress
	}{
		Ctx:       ctx,
		Sender:    sender,
		Recipient: recipient,
	}
	mock.lockLinkAddresses.Lock()
	mock.calls.LinkAddresses = append(mock.calls.LinkAddresses, callInfo)
	mock.lockLinkAddresses.Unlock()
	mock.LinkAddressesFunc(ctx, sender, recipient)
}

// LinkAddressesCalls gets all the calls that were made to LinkAddresses.
// Check the length with:
//     len(mockedBalancer.LinkAddressesCalls())
func (mock *BalancerMock) LinkAddressesCalls() []struct {
	Ctx       sdk.Context
	Sender    exported.CrossChainAddress
	Recipient exported.CrossChainAddress
} {
	var calls []struct {
		Ctx       sdk.Context
		Sender    exported.CrossChainAddress
		Recipient exported.CrossChainAddress
	}
	mock.lockLinkAddresses.RLock()
	calls = mock.calls.LinkAddresses
	mock.lockLinkAddresses.RUnlock()
	return calls
}

// PrepareForTransfer calls PrepareForTransferFunc.
func (mock *BalancerMock) PrepareForTransfer(ctx sdk.Context, sender exported.CrossChainAddress, amount sdk.Coin) error {
	if mock.PrepareForTransferFunc == nil {
		panic("BalancerMock.PrepareForTransferFunc: method is nil but Balancer.PrepareForTransfer was just called")
	}
	callInfo := struct {
		Ctx    sdk.Context
		Sender exported.CrossChainAddress
		Amount sdk.Coin
	}{
		Ctx:    ctx,
		Sender: sender,
		Amount: amount,
	}
	mock.lockPrepareForTransfer.Lock()
	mock.calls.PrepareForTransfer = append(mock.calls.PrepareForTransfer, callInfo)
	mock.lockPrepareForTransfer.Unlock()
	return mock.PrepareForTransferFunc(ctx, sender, amount)
}

// PrepareForTransferCalls gets all the calls that were made to PrepareForTransfer.
// Check the length with:
//     len(mockedBalancer.PrepareForTransferCalls())
func (mock *BalancerMock) PrepareForTransferCalls() []struct {
	Ctx    sdk.Context
	Sender exported.CrossChainAddress
	Amount sdk.Coin
} {
	var calls []struct {
		Ctx    sdk.Context
		Sender exported.CrossChainAddress
		Amount sdk.Coin
	}
	mock.lockPrepareForTransfer.RLock()
	calls = mock.calls.PrepareForTransfer
	mock.lockPrepareForTransfer.RUnlock()
	return calls
}

// Ensure, that SnapshotterMock does implement types.Snapshotter.
// If this is not the case, regenerate this file with moq.
var _ types.Snapshotter = &SnapshotterMock{}

// SnapshotterMock is a mock implementation of types.Snapshotter.
//
//     func TestSomethingThatUsesSnapshotter(t *testing.T) {
//
//         // make and configure a mocked types.Snapshotter
//         mockedSnapshotter := &SnapshotterMock{
//             GetSnapshotFunc: func(ctx sdk.Context, round int64) (snapshot.Snapshot, bool) {
// 	               panic("mock out the GetSnapshot method")
//             },
//         }
//
//         // use mockedSnapshotter in code that requires types.Snapshotter
//         // and then make assertions.
//
//     }
type SnapshotterMock struct {
	// GetSnapshotFunc mocks the GetSnapshot method.
	GetSnapshotFunc func(ctx sdk.Context, round int64) (snapshot.Snapshot, bool)

	// calls tracks calls to the methods.
	calls struct {
		// GetSnapshot holds details about calls to the GetSnapshot method.
		GetSnapshot []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Round is the round argument value.
			Round int64
		}
	}
	lockGetSnapshot sync.RWMutex
}

// GetSnapshot calls GetSnapshotFunc.
func (mock *SnapshotterMock) GetSnapshot(ctx sdk.Context, round int64) (snapshot.Snapshot, bool) {
	if mock.GetSnapshotFunc == nil {
		panic("SnapshotterMock.GetSnapshotFunc: method is nil but Snapshotter.GetSnapshot was just called")
	}
	callInfo := struct {
		Ctx   sdk.Context
		Round int64
	}{
		Ctx:   ctx,
		Round: round,
	}
	mock.lockGetSnapshot.Lock()
	mock.calls.GetSnapshot = append(mock.calls.GetSnapshot, callInfo)
	mock.lockGetSnapshot.Unlock()
	return mock.GetSnapshotFunc(ctx, round)
}

// GetSnapshotCalls gets all the calls that were made to GetSnapshot.
// Check the length with:
//     len(mockedSnapshotter.GetSnapshotCalls())
func (mock *SnapshotterMock) GetSnapshotCalls() []struct {
	Ctx   sdk.Context
	Round int64
} {
	var calls []struct {
		Ctx   sdk.Context
		Round int64
	}
	mock.lockGetSnapshot.RLock()
	calls = mock.calls.GetSnapshot
	mock.lockGetSnapshot.RUnlock()
	return calls
}
