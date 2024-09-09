// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// FundManagerUpdateSolCapacityAmount is the `fund_manager_update_sol_capacity_amount` instruction.
type FundManagerUpdateSolCapacityAmount struct {
	CapacityAmount *uint64

	// [0] = [SIGNER] fund_manager
	//
	// [1] = [] receipt_token_mint
	//
	// [2] = [WRITE] fund_account
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewFundManagerUpdateSolCapacityAmountInstructionBuilder creates a new `FundManagerUpdateSolCapacityAmount` instruction builder.
func NewFundManagerUpdateSolCapacityAmountInstructionBuilder() *FundManagerUpdateSolCapacityAmount {
	nd := &FundManagerUpdateSolCapacityAmount{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	nd.AccountMetaSlice[0] = ag_solanago.Meta(Addresses["5UpLTLA7Wjqp7qdfjuTtPcUw3aVtbqFA5Mgm34mxPNg2"]).SIGNER()
	nd.AccountMetaSlice[1] = ag_solanago.Meta(Addresses["FRAGSEthVFL7fdqM8hxfxkfCZzUvmg21cqPJVvC1qdbo"])
	return nd
}

// SetCapacityAmount sets the "capacity_amount" parameter.
func (inst *FundManagerUpdateSolCapacityAmount) SetCapacityAmount(capacity_amount uint64) *FundManagerUpdateSolCapacityAmount {
	inst.CapacityAmount = &capacity_amount
	return inst
}

// SetFundManagerAccount sets the "fund_manager" account.
func (inst *FundManagerUpdateSolCapacityAmount) SetFundManagerAccount(fundManager ag_solanago.PublicKey) *FundManagerUpdateSolCapacityAmount {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(fundManager).SIGNER()
	return inst
}

// GetFundManagerAccount gets the "fund_manager" account.
func (inst *FundManagerUpdateSolCapacityAmount) GetFundManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *FundManagerUpdateSolCapacityAmount) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *FundManagerUpdateSolCapacityAmount {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(receiptTokenMint)
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *FundManagerUpdateSolCapacityAmount) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetFundAccountAccount sets the "fund_account" account.
func (inst *FundManagerUpdateSolCapacityAmount) SetFundAccountAccount(fundAccount ag_solanago.PublicKey) *FundManagerUpdateSolCapacityAmount {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(fundAccount).WRITE()
	return inst
}

func (inst *FundManagerUpdateSolCapacityAmount) findFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: fund
	seeds = append(seeds, []byte{byte(0x66), byte(0x75), byte(0x6e), byte(0x64)})
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindFundAccountAddressWithBumpSeed calculates FundAccount account address with given seeds and a known bump seed.
func (inst *FundManagerUpdateSolCapacityAmount) FindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *FundManagerUpdateSolCapacityAmount) MustFindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundAccountAddress finds FundAccount account address with given seeds.
func (inst *FundManagerUpdateSolCapacityAmount) FindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *FundManagerUpdateSolCapacityAmount) MustFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundAccountAccount gets the "fund_account" account.
func (inst *FundManagerUpdateSolCapacityAmount) GetFundAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst FundManagerUpdateSolCapacityAmount) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_FundManagerUpdateSolCapacityAmount,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FundManagerUpdateSolCapacityAmount) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FundManagerUpdateSolCapacityAmount) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.CapacityAmount == nil {
			return errors.New("CapacityAmount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.FundManager is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.FundAccount is not set")
		}
	}
	return nil
}

func (inst *FundManagerUpdateSolCapacityAmount) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FundManagerUpdateSolCapacityAmount")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" CapacityAmount", *inst.CapacityAmount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      fund_manager", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("receipt_token_mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("             fund_", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj FundManagerUpdateSolCapacityAmount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `CapacityAmount` param:
	err = encoder.Encode(obj.CapacityAmount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *FundManagerUpdateSolCapacityAmount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `CapacityAmount`:
	err = decoder.Decode(&obj.CapacityAmount)
	if err != nil {
		return err
	}
	return nil
}

// NewFundManagerUpdateSolCapacityAmountInstruction declares a new FundManagerUpdateSolCapacityAmount instruction with the provided parameters and accounts.
func NewFundManagerUpdateSolCapacityAmountInstruction(
	// Parameters:
	capacity_amount uint64,
	// Accounts:
	fundManager ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	fundAccount ag_solanago.PublicKey) *FundManagerUpdateSolCapacityAmount {
	return NewFundManagerUpdateSolCapacityAmountInstructionBuilder().
		SetCapacityAmount(capacity_amount).
		SetFundManagerAccount(fundManager).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetFundAccountAccount(fundAccount)
}
