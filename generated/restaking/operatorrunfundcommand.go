// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// OperatorRunFundCommand is the `operator_run_fund_command` instruction.
type OperatorRunFundCommand struct {
	ForceResetCommand *OperationCommandEntry `bin:"optional"`

	// [0] = [WRITE, SIGNER] operator
	//
	// [1] = [] system_program
	//
	// [2] = [WRITE] receipt_token_mint
	//
	// [3] = [WRITE] fund_account
	//
	// [4] = [] event_authority
	//
	// [5] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewOperatorRunFundCommandInstructionBuilder creates a new `OperatorRunFundCommand` instruction builder.
func NewOperatorRunFundCommandInstructionBuilder() *OperatorRunFundCommand {
	nd := &OperatorRunFundCommand{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	nd.AccountMetaSlice[1] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	return nd
}

// SetForceResetCommand sets the "force_reset_command" parameter.
func (inst *OperatorRunFundCommand) SetForceResetCommand(force_reset_command OperationCommandEntry) *OperatorRunFundCommand {
	inst.ForceResetCommand = &force_reset_command
	return inst
}

// SetOperatorAccount sets the "operator" account.
func (inst *OperatorRunFundCommand) SetOperatorAccount(operator ag_solanago.PublicKey) *OperatorRunFundCommand {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(operator).WRITE().SIGNER()
	return inst
}

// GetOperatorAccount gets the "operator" account.
func (inst *OperatorRunFundCommand) GetOperatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *OperatorRunFundCommand) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *OperatorRunFundCommand {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *OperatorRunFundCommand) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *OperatorRunFundCommand) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *OperatorRunFundCommand {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(receiptTokenMint).WRITE()
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *OperatorRunFundCommand) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetFundAccountAccount sets the "fund_account" account.
func (inst *OperatorRunFundCommand) SetFundAccountAccount(fundAccount ag_solanago.PublicKey) *OperatorRunFundCommand {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(fundAccount).WRITE()
	return inst
}

func (inst *OperatorRunFundCommand) findFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *OperatorRunFundCommand) FindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *OperatorRunFundCommand) MustFindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundAccountAddress finds FundAccount account address with given seeds.
func (inst *OperatorRunFundCommand) FindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *OperatorRunFundCommand) MustFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundAccountAccount gets the "fund_account" account.
func (inst *OperatorRunFundCommand) GetFundAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *OperatorRunFundCommand) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *OperatorRunFundCommand {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(eventAuthority)
	return inst
}

func (inst *OperatorRunFundCommand) findFindEventAuthorityAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: __event_authority
	seeds = append(seeds, []byte{byte(0x5f), byte(0x5f), byte(0x65), byte(0x76), byte(0x65), byte(0x6e), byte(0x74), byte(0x5f), byte(0x61), byte(0x75), byte(0x74), byte(0x68), byte(0x6f), byte(0x72), byte(0x69), byte(0x74), byte(0x79)})

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindEventAuthorityAddressWithBumpSeed calculates EventAuthority account address with given seeds and a known bump seed.
func (inst *OperatorRunFundCommand) FindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindEventAuthorityAddress(bumpSeed)
	return
}

func (inst *OperatorRunFundCommand) MustFindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindEventAuthorityAddress finds EventAuthority account address with given seeds.
func (inst *OperatorRunFundCommand) FindEventAuthorityAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindEventAuthorityAddress(0)
	return
}

func (inst *OperatorRunFundCommand) MustFindEventAuthorityAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *OperatorRunFundCommand) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetProgramAccount sets the "program" account.
func (inst *OperatorRunFundCommand) SetProgramAccount(program ag_solanago.PublicKey) *OperatorRunFundCommand {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *OperatorRunFundCommand) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst OperatorRunFundCommand) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_OperatorRunFundCommand,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst OperatorRunFundCommand) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *OperatorRunFundCommand) Validate() error {
	// Check whether all (required) parameters are set:
	{
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Operator is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.FundAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *OperatorRunFundCommand) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("OperatorRunFundCommand")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  ForceResetCommand (OPT)", inst.ForceResetCommand))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("          operator", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    system_program", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("receipt_token_mint", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("             fund_", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("   event_authority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("           program", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj OperatorRunFundCommand) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ForceResetCommand` param (optional):
	{
		if obj.ForceResetCommand == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.ForceResetCommand)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (obj *OperatorRunFundCommand) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ForceResetCommand` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.ForceResetCommand)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// NewOperatorRunFundCommandInstruction declares a new OperatorRunFundCommand instruction with the provided parameters and accounts.
func NewOperatorRunFundCommandInstruction(
	// Parameters:
	force_reset_command OperationCommandEntry,
	// Accounts:
	operator ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	fundAccount ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *OperatorRunFundCommand {
	return NewOperatorRunFundCommandInstructionBuilder().
		SetForceResetCommand(force_reset_command).
		SetOperatorAccount(operator).
		SetSystemProgramAccount(systemProgram).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetFundAccountAccount(fundAccount).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
