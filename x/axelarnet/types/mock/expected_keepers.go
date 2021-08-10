// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/axelarnetwork/axelar-core/x/axelarnet/types"
	nexus "github.com/axelarnetwork/axelar-core/x/nexus/exported"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sync"
)

// Ensure, that NexusMock does implement types.Nexus.
// If this is not the case, regenerate this file with moq.
var _ types.Nexus = &NexusMock{}

// NexusMock is a mock implementation of types.Nexus.
//
// 	func TestSomethingThatUsesNexus(t *testing.T) {
//
// 		// make and configure a mocked types.Nexus
// 		mockedNexus := &NexusMock{
// 			ArchivePendingTransferFunc: func(ctx sdk.Context, transfer nexus.CrossChainTransfer)  {
// 				panic("mock out the ArchivePendingTransfer method")
// 			},
// 			EnqueueForTransferFunc: func(ctx sdk.Context, sender nexus.CrossChainAddress, amount sdk.Coin) error {
// 				panic("mock out the EnqueueForTransfer method")
// 			},
// 			GetChainFunc: func(ctx sdk.Context, chain string) (nexus.Chain, bool) {
// 				panic("mock out the GetChain method")
// 			},
// 			GetTransfersForChainFunc: func(ctx sdk.Context, chain nexus.Chain, state nexus.TransferState) []nexus.CrossChainTransfer {
// 				panic("mock out the GetTransfersForChain method")
// 			},
// 			IsAssetRegisteredFunc: func(ctx sdk.Context, chainName string, denom string) bool {
// 				panic("mock out the IsAssetRegistered method")
// 			},
// 			LinkAddressesFunc: func(ctx sdk.Context, sender nexus.CrossChainAddress, recipient nexus.CrossChainAddress)  {
// 				panic("mock out the LinkAddresses method")
// 			},
// 			RegisterAssetFunc: func(ctx sdk.Context, chainName string, denom string)  {
// 				panic("mock out the RegisterAsset method")
// 			},
// 		}
//
// 		// use mockedNexus in code that requires types.Nexus
// 		// and then make assertions.
//
// 	}
type NexusMock struct {
	// ArchivePendingTransferFunc mocks the ArchivePendingTransfer method.
	ArchivePendingTransferFunc func(ctx sdk.Context, transfer nexus.CrossChainTransfer)

	// EnqueueForTransferFunc mocks the EnqueueForTransfer method.
	EnqueueForTransferFunc func(ctx sdk.Context, sender nexus.CrossChainAddress, amount sdk.Coin) error

	// GetChainFunc mocks the GetChain method.
	GetChainFunc func(ctx sdk.Context, chain string) (nexus.Chain, bool)

	// GetTransfersForChainFunc mocks the GetTransfersForChain method.
	GetTransfersForChainFunc func(ctx sdk.Context, chain nexus.Chain, state nexus.TransferState) []nexus.CrossChainTransfer

	// IsAssetRegisteredFunc mocks the IsAssetRegistered method.
	IsAssetRegisteredFunc func(ctx sdk.Context, chainName string, denom string) bool

	// LinkAddressesFunc mocks the LinkAddresses method.
	LinkAddressesFunc func(ctx sdk.Context, sender nexus.CrossChainAddress, recipient nexus.CrossChainAddress)

	// RegisterAssetFunc mocks the RegisterAsset method.
	RegisterAssetFunc func(ctx sdk.Context, chainName string, denom string)

	// calls tracks calls to the methods.
	calls struct {
		// ArchivePendingTransfer holds details about calls to the ArchivePendingTransfer method.
		ArchivePendingTransfer []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Transfer is the transfer argument value.
			Transfer nexus.CrossChainTransfer
		}
		// EnqueueForTransfer holds details about calls to the EnqueueForTransfer method.
		EnqueueForTransfer []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Sender is the sender argument value.
			Sender nexus.CrossChainAddress
			// Amount is the amount argument value.
			Amount sdk.Coin
		}
		// GetChain holds details about calls to the GetChain method.
		GetChain []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Chain is the chain argument value.
			Chain string
		}
		// GetTransfersForChain holds details about calls to the GetTransfersForChain method.
		GetTransfersForChain []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Chain is the chain argument value.
			Chain nexus.Chain
			// State is the state argument value.
			State nexus.TransferState
		}
		// IsAssetRegistered holds details about calls to the IsAssetRegistered method.
		IsAssetRegistered []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// ChainName is the chainName argument value.
			ChainName string
			// Denom is the denom argument value.
			Denom string
		}
		// LinkAddresses holds details about calls to the LinkAddresses method.
		LinkAddresses []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Sender is the sender argument value.
			Sender nexus.CrossChainAddress
			// Recipient is the recipient argument value.
			Recipient nexus.CrossChainAddress
		}
		// RegisterAsset holds details about calls to the RegisterAsset method.
		RegisterAsset []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// ChainName is the chainName argument value.
			ChainName string
			// Denom is the denom argument value.
			Denom string
		}
	}
	lockArchivePendingTransfer sync.RWMutex
	lockEnqueueForTransfer     sync.RWMutex
	lockGetChain               sync.RWMutex
	lockGetTransfersForChain   sync.RWMutex
	lockIsAssetRegistered      sync.RWMutex
	lockLinkAddresses          sync.RWMutex
	lockRegisterAsset          sync.RWMutex
}

// ArchivePendingTransfer calls ArchivePendingTransferFunc.
func (mock *NexusMock) ArchivePendingTransfer(ctx sdk.Context, transfer nexus.CrossChainTransfer) {
	if mock.ArchivePendingTransferFunc == nil {
		panic("NexusMock.ArchivePendingTransferFunc: method is nil but Nexus.ArchivePendingTransfer was just called")
	}
	callInfo := struct {
		Ctx      sdk.Context
		Transfer nexus.CrossChainTransfer
	}{
		Ctx:      ctx,
		Transfer: transfer,
	}
	mock.lockArchivePendingTransfer.Lock()
	mock.calls.ArchivePendingTransfer = append(mock.calls.ArchivePendingTransfer, callInfo)
	mock.lockArchivePendingTransfer.Unlock()
	mock.ArchivePendingTransferFunc(ctx, transfer)
}

// ArchivePendingTransferCalls gets all the calls that were made to ArchivePendingTransfer.
// Check the length with:
//     len(mockedNexus.ArchivePendingTransferCalls())
func (mock *NexusMock) ArchivePendingTransferCalls() []struct {
	Ctx      sdk.Context
	Transfer nexus.CrossChainTransfer
} {
	var calls []struct {
		Ctx      sdk.Context
		Transfer nexus.CrossChainTransfer
	}
	mock.lockArchivePendingTransfer.RLock()
	calls = mock.calls.ArchivePendingTransfer
	mock.lockArchivePendingTransfer.RUnlock()
	return calls
}

// EnqueueForTransfer calls EnqueueForTransferFunc.
func (mock *NexusMock) EnqueueForTransfer(ctx sdk.Context, sender nexus.CrossChainAddress, amount sdk.Coin) error {
	if mock.EnqueueForTransferFunc == nil {
		panic("NexusMock.EnqueueForTransferFunc: method is nil but Nexus.EnqueueForTransfer was just called")
	}
	callInfo := struct {
		Ctx    sdk.Context
		Sender nexus.CrossChainAddress
		Amount sdk.Coin
	}{
		Ctx:    ctx,
		Sender: sender,
		Amount: amount,
	}
	mock.lockEnqueueForTransfer.Lock()
	mock.calls.EnqueueForTransfer = append(mock.calls.EnqueueForTransfer, callInfo)
	mock.lockEnqueueForTransfer.Unlock()
	return mock.EnqueueForTransferFunc(ctx, sender, amount)
}

// EnqueueForTransferCalls gets all the calls that were made to EnqueueForTransfer.
// Check the length with:
//     len(mockedNexus.EnqueueForTransferCalls())
func (mock *NexusMock) EnqueueForTransferCalls() []struct {
	Ctx    sdk.Context
	Sender nexus.CrossChainAddress
	Amount sdk.Coin
} {
	var calls []struct {
		Ctx    sdk.Context
		Sender nexus.CrossChainAddress
		Amount sdk.Coin
	}
	mock.lockEnqueueForTransfer.RLock()
	calls = mock.calls.EnqueueForTransfer
	mock.lockEnqueueForTransfer.RUnlock()
	return calls
}

// GetChain calls GetChainFunc.
func (mock *NexusMock) GetChain(ctx sdk.Context, chain string) (nexus.Chain, bool) {
	if mock.GetChainFunc == nil {
		panic("NexusMock.GetChainFunc: method is nil but Nexus.GetChain was just called")
	}
	callInfo := struct {
		Ctx   sdk.Context
		Chain string
	}{
		Ctx:   ctx,
		Chain: chain,
	}
	mock.lockGetChain.Lock()
	mock.calls.GetChain = append(mock.calls.GetChain, callInfo)
	mock.lockGetChain.Unlock()
	return mock.GetChainFunc(ctx, chain)
}

// GetChainCalls gets all the calls that were made to GetChain.
// Check the length with:
//     len(mockedNexus.GetChainCalls())
func (mock *NexusMock) GetChainCalls() []struct {
	Ctx   sdk.Context
	Chain string
} {
	var calls []struct {
		Ctx   sdk.Context
		Chain string
	}
	mock.lockGetChain.RLock()
	calls = mock.calls.GetChain
	mock.lockGetChain.RUnlock()
	return calls
}

// GetTransfersForChain calls GetTransfersForChainFunc.
func (mock *NexusMock) GetTransfersForChain(ctx sdk.Context, chain nexus.Chain, state nexus.TransferState) []nexus.CrossChainTransfer {
	if mock.GetTransfersForChainFunc == nil {
		panic("NexusMock.GetTransfersForChainFunc: method is nil but Nexus.GetTransfersForChain was just called")
	}
	callInfo := struct {
		Ctx   sdk.Context
		Chain nexus.Chain
		State nexus.TransferState
	}{
		Ctx:   ctx,
		Chain: chain,
		State: state,
	}
	mock.lockGetTransfersForChain.Lock()
	mock.calls.GetTransfersForChain = append(mock.calls.GetTransfersForChain, callInfo)
	mock.lockGetTransfersForChain.Unlock()
	return mock.GetTransfersForChainFunc(ctx, chain, state)
}

// GetTransfersForChainCalls gets all the calls that were made to GetTransfersForChain.
// Check the length with:
//     len(mockedNexus.GetTransfersForChainCalls())
func (mock *NexusMock) GetTransfersForChainCalls() []struct {
	Ctx   sdk.Context
	Chain nexus.Chain
	State nexus.TransferState
} {
	var calls []struct {
		Ctx   sdk.Context
		Chain nexus.Chain
		State nexus.TransferState
	}
	mock.lockGetTransfersForChain.RLock()
	calls = mock.calls.GetTransfersForChain
	mock.lockGetTransfersForChain.RUnlock()
	return calls
}

// IsAssetRegistered calls IsAssetRegisteredFunc.
func (mock *NexusMock) IsAssetRegistered(ctx sdk.Context, chainName string, denom string) bool {
	if mock.IsAssetRegisteredFunc == nil {
		panic("NexusMock.IsAssetRegisteredFunc: method is nil but Nexus.IsAssetRegistered was just called")
	}
	callInfo := struct {
		Ctx       sdk.Context
		ChainName string
		Denom     string
	}{
		Ctx:       ctx,
		ChainName: chainName,
		Denom:     denom,
	}
	mock.lockIsAssetRegistered.Lock()
	mock.calls.IsAssetRegistered = append(mock.calls.IsAssetRegistered, callInfo)
	mock.lockIsAssetRegistered.Unlock()
	return mock.IsAssetRegisteredFunc(ctx, chainName, denom)
}

// IsAssetRegisteredCalls gets all the calls that were made to IsAssetRegistered.
// Check the length with:
//     len(mockedNexus.IsAssetRegisteredCalls())
func (mock *NexusMock) IsAssetRegisteredCalls() []struct {
	Ctx       sdk.Context
	ChainName string
	Denom     string
} {
	var calls []struct {
		Ctx       sdk.Context
		ChainName string
		Denom     string
	}
	mock.lockIsAssetRegistered.RLock()
	calls = mock.calls.IsAssetRegistered
	mock.lockIsAssetRegistered.RUnlock()
	return calls
}

// LinkAddresses calls LinkAddressesFunc.
func (mock *NexusMock) LinkAddresses(ctx sdk.Context, sender nexus.CrossChainAddress, recipient nexus.CrossChainAddress) {
	if mock.LinkAddressesFunc == nil {
		panic("NexusMock.LinkAddressesFunc: method is nil but Nexus.LinkAddresses was just called")
	}
	callInfo := struct {
		Ctx       sdk.Context
		Sender    nexus.CrossChainAddress
		Recipient nexus.CrossChainAddress
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
//     len(mockedNexus.LinkAddressesCalls())
func (mock *NexusMock) LinkAddressesCalls() []struct {
	Ctx       sdk.Context
	Sender    nexus.CrossChainAddress
	Recipient nexus.CrossChainAddress
} {
	var calls []struct {
		Ctx       sdk.Context
		Sender    nexus.CrossChainAddress
		Recipient nexus.CrossChainAddress
	}
	mock.lockLinkAddresses.RLock()
	calls = mock.calls.LinkAddresses
	mock.lockLinkAddresses.RUnlock()
	return calls
}

// RegisterAsset calls RegisterAssetFunc.
func (mock *NexusMock) RegisterAsset(ctx sdk.Context, chainName string, denom string) {
	if mock.RegisterAssetFunc == nil {
		panic("NexusMock.RegisterAssetFunc: method is nil but Nexus.RegisterAsset was just called")
	}
	callInfo := struct {
		Ctx       sdk.Context
		ChainName string
		Denom     string
	}{
		Ctx:       ctx,
		ChainName: chainName,
		Denom:     denom,
	}
	mock.lockRegisterAsset.Lock()
	mock.calls.RegisterAsset = append(mock.calls.RegisterAsset, callInfo)
	mock.lockRegisterAsset.Unlock()
	mock.RegisterAssetFunc(ctx, chainName, denom)
}

// RegisterAssetCalls gets all the calls that were made to RegisterAsset.
// Check the length with:
//     len(mockedNexus.RegisterAssetCalls())
func (mock *NexusMock) RegisterAssetCalls() []struct {
	Ctx       sdk.Context
	ChainName string
	Denom     string
} {
	var calls []struct {
		Ctx       sdk.Context
		ChainName string
		Denom     string
	}
	mock.lockRegisterAsset.RLock()
	calls = mock.calls.RegisterAsset
	mock.lockRegisterAsset.RUnlock()
	return calls
}

// Ensure, that BankKeeperMock does implement types.BankKeeper.
// If this is not the case, regenerate this file with moq.
var _ types.BankKeeper = &BankKeeperMock{}

// BankKeeperMock is a mock implementation of types.BankKeeper.
//
// 	func TestSomethingThatUsesBankKeeper(t *testing.T) {
//
// 		// make and configure a mocked types.BankKeeper
// 		mockedBankKeeper := &BankKeeperMock{
// 			BurnCoinsFunc: func(ctx sdk.Context, moduleName string, amt sdk.Coins) error {
// 				panic("mock out the BurnCoins method")
// 			},
// 			MintCoinsFunc: func(ctx sdk.Context, moduleName string, amt sdk.Coins) error {
// 				panic("mock out the MintCoins method")
// 			},
// 			SendCoinsFromAccountToModuleFunc: func(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
// 				panic("mock out the SendCoinsFromAccountToModule method")
// 			},
// 			SendCoinsFromModuleToAccountFunc: func(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
// 				panic("mock out the SendCoinsFromModuleToAccount method")
// 			},
// 		}
//
// 		// use mockedBankKeeper in code that requires types.BankKeeper
// 		// and then make assertions.
//
// 	}
type BankKeeperMock struct {
	// BurnCoinsFunc mocks the BurnCoins method.
	BurnCoinsFunc func(ctx sdk.Context, moduleName string, amt sdk.Coins) error

	// MintCoinsFunc mocks the MintCoins method.
	MintCoinsFunc func(ctx sdk.Context, moduleName string, amt sdk.Coins) error

	// SendCoinsFromAccountToModuleFunc mocks the SendCoinsFromAccountToModule method.
	SendCoinsFromAccountToModuleFunc func(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error

	// SendCoinsFromModuleToAccountFunc mocks the SendCoinsFromModuleToAccount method.
	SendCoinsFromModuleToAccountFunc func(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error

	// calls tracks calls to the methods.
	calls struct {
		// BurnCoins holds details about calls to the BurnCoins method.
		BurnCoins []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// ModuleName is the moduleName argument value.
			ModuleName string
			// Amt is the amt argument value.
			Amt sdk.Coins
		}
		// MintCoins holds details about calls to the MintCoins method.
		MintCoins []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// ModuleName is the moduleName argument value.
			ModuleName string
			// Amt is the amt argument value.
			Amt sdk.Coins
		}
		// SendCoinsFromAccountToModule holds details about calls to the SendCoinsFromAccountToModule method.
		SendCoinsFromAccountToModule []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// SenderAddr is the senderAddr argument value.
			SenderAddr sdk.AccAddress
			// RecipientModule is the recipientModule argument value.
			RecipientModule string
			// Amt is the amt argument value.
			Amt sdk.Coins
		}
		// SendCoinsFromModuleToAccount holds details about calls to the SendCoinsFromModuleToAccount method.
		SendCoinsFromModuleToAccount []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// SenderModule is the senderModule argument value.
			SenderModule string
			// RecipientAddr is the recipientAddr argument value.
			RecipientAddr sdk.AccAddress
			// Amt is the amt argument value.
			Amt sdk.Coins
		}
	}
	lockBurnCoins                    sync.RWMutex
	lockMintCoins                    sync.RWMutex
	lockSendCoinsFromAccountToModule sync.RWMutex
	lockSendCoinsFromModuleToAccount sync.RWMutex
}

// BurnCoins calls BurnCoinsFunc.
func (mock *BankKeeperMock) BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error {
	if mock.BurnCoinsFunc == nil {
		panic("BankKeeperMock.BurnCoinsFunc: method is nil but BankKeeper.BurnCoins was just called")
	}
	callInfo := struct {
		Ctx        sdk.Context
		ModuleName string
		Amt        sdk.Coins
	}{
		Ctx:        ctx,
		ModuleName: moduleName,
		Amt:        amt,
	}
	mock.lockBurnCoins.Lock()
	mock.calls.BurnCoins = append(mock.calls.BurnCoins, callInfo)
	mock.lockBurnCoins.Unlock()
	return mock.BurnCoinsFunc(ctx, moduleName, amt)
}

// BurnCoinsCalls gets all the calls that were made to BurnCoins.
// Check the length with:
//     len(mockedBankKeeper.BurnCoinsCalls())
func (mock *BankKeeperMock) BurnCoinsCalls() []struct {
	Ctx        sdk.Context
	ModuleName string
	Amt        sdk.Coins
} {
	var calls []struct {
		Ctx        sdk.Context
		ModuleName string
		Amt        sdk.Coins
	}
	mock.lockBurnCoins.RLock()
	calls = mock.calls.BurnCoins
	mock.lockBurnCoins.RUnlock()
	return calls
}

// MintCoins calls MintCoinsFunc.
func (mock *BankKeeperMock) MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error {
	if mock.MintCoinsFunc == nil {
		panic("BankKeeperMock.MintCoinsFunc: method is nil but BankKeeper.MintCoins was just called")
	}
	callInfo := struct {
		Ctx        sdk.Context
		ModuleName string
		Amt        sdk.Coins
	}{
		Ctx:        ctx,
		ModuleName: moduleName,
		Amt:        amt,
	}
	mock.lockMintCoins.Lock()
	mock.calls.MintCoins = append(mock.calls.MintCoins, callInfo)
	mock.lockMintCoins.Unlock()
	return mock.MintCoinsFunc(ctx, moduleName, amt)
}

// MintCoinsCalls gets all the calls that were made to MintCoins.
// Check the length with:
//     len(mockedBankKeeper.MintCoinsCalls())
func (mock *BankKeeperMock) MintCoinsCalls() []struct {
	Ctx        sdk.Context
	ModuleName string
	Amt        sdk.Coins
} {
	var calls []struct {
		Ctx        sdk.Context
		ModuleName string
		Amt        sdk.Coins
	}
	mock.lockMintCoins.RLock()
	calls = mock.calls.MintCoins
	mock.lockMintCoins.RUnlock()
	return calls
}

// SendCoinsFromAccountToModule calls SendCoinsFromAccountToModuleFunc.
func (mock *BankKeeperMock) SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	if mock.SendCoinsFromAccountToModuleFunc == nil {
		panic("BankKeeperMock.SendCoinsFromAccountToModuleFunc: method is nil but BankKeeper.SendCoinsFromAccountToModule was just called")
	}
	callInfo := struct {
		Ctx             sdk.Context
		SenderAddr      sdk.AccAddress
		RecipientModule string
		Amt             sdk.Coins
	}{
		Ctx:             ctx,
		SenderAddr:      senderAddr,
		RecipientModule: recipientModule,
		Amt:             amt,
	}
	mock.lockSendCoinsFromAccountToModule.Lock()
	mock.calls.SendCoinsFromAccountToModule = append(mock.calls.SendCoinsFromAccountToModule, callInfo)
	mock.lockSendCoinsFromAccountToModule.Unlock()
	return mock.SendCoinsFromAccountToModuleFunc(ctx, senderAddr, recipientModule, amt)
}

// SendCoinsFromAccountToModuleCalls gets all the calls that were made to SendCoinsFromAccountToModule.
// Check the length with:
//     len(mockedBankKeeper.SendCoinsFromAccountToModuleCalls())
func (mock *BankKeeperMock) SendCoinsFromAccountToModuleCalls() []struct {
	Ctx             sdk.Context
	SenderAddr      sdk.AccAddress
	RecipientModule string
	Amt             sdk.Coins
} {
	var calls []struct {
		Ctx             sdk.Context
		SenderAddr      sdk.AccAddress
		RecipientModule string
		Amt             sdk.Coins
	}
	mock.lockSendCoinsFromAccountToModule.RLock()
	calls = mock.calls.SendCoinsFromAccountToModule
	mock.lockSendCoinsFromAccountToModule.RUnlock()
	return calls
}

// SendCoinsFromModuleToAccount calls SendCoinsFromModuleToAccountFunc.
func (mock *BankKeeperMock) SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	if mock.SendCoinsFromModuleToAccountFunc == nil {
		panic("BankKeeperMock.SendCoinsFromModuleToAccountFunc: method is nil but BankKeeper.SendCoinsFromModuleToAccount was just called")
	}
	callInfo := struct {
		Ctx           sdk.Context
		SenderModule  string
		RecipientAddr sdk.AccAddress
		Amt           sdk.Coins
	}{
		Ctx:           ctx,
		SenderModule:  senderModule,
		RecipientAddr: recipientAddr,
		Amt:           amt,
	}
	mock.lockSendCoinsFromModuleToAccount.Lock()
	mock.calls.SendCoinsFromModuleToAccount = append(mock.calls.SendCoinsFromModuleToAccount, callInfo)
	mock.lockSendCoinsFromModuleToAccount.Unlock()
	return mock.SendCoinsFromModuleToAccountFunc(ctx, senderModule, recipientAddr, amt)
}

// SendCoinsFromModuleToAccountCalls gets all the calls that were made to SendCoinsFromModuleToAccount.
// Check the length with:
//     len(mockedBankKeeper.SendCoinsFromModuleToAccountCalls())
func (mock *BankKeeperMock) SendCoinsFromModuleToAccountCalls() []struct {
	Ctx           sdk.Context
	SenderModule  string
	RecipientAddr sdk.AccAddress
	Amt           sdk.Coins
} {
	var calls []struct {
		Ctx           sdk.Context
		SenderModule  string
		RecipientAddr sdk.AccAddress
		Amt           sdk.Coins
	}
	mock.lockSendCoinsFromModuleToAccount.RLock()
	calls = mock.calls.SendCoinsFromModuleToAccount
	mock.lockSendCoinsFromModuleToAccount.RUnlock()
	return calls
}