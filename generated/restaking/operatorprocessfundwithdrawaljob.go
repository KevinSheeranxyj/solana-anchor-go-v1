// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// OperatorProcessFundWithdrawalJob is the `operator_process_fund_withdrawal_job` instruction.
type OperatorProcessFundWithdrawalJob struct {
	Forced *bool

	// [0] = [SIGNER] operator
	//
	// [1] = [WRITE] receipt_token_mint
	//
	// [2] = [] receipt_token_program
	//
	// [3] = [] receipt_token_lock_authority
	//
	// [4] = [WRITE] receipt_token_lock_account
	//
	// [5] = [WRITE] fund_account
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewOperatorProcessFundWithdrawalJobInstructionBuilder creates a new `OperatorProcessFundWithdrawalJob` instruction builder.
func NewOperatorProcessFundWithdrawalJobInstructionBuilder() *OperatorProcessFundWithdrawalJob {
	nd := &OperatorProcessFundWithdrawalJob{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	nd.AccountMetaSlice[1] = ag_solanago.Meta(Addresses["FRAGSEthVFL7fdqM8hxfxkfCZzUvmg21cqPJVvC1qdbo"]).WRITE()
	nd.AccountMetaSlice[2] = ag_solanago.Meta(Addresses["TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb"])
	return nd
}

// SetForced sets the "forced" parameter.
func (inst *OperatorProcessFundWithdrawalJob) SetForced(forced bool) *OperatorProcessFundWithdrawalJob {
	inst.Forced = &forced
	return inst
}

// SetOperatorAccount sets the "operator" account.
func (inst *OperatorProcessFundWithdrawalJob) SetOperatorAccount(operator ag_solanago.PublicKey) *OperatorProcessFundWithdrawalJob {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(operator).SIGNER()
	return inst
}

// GetOperatorAccount gets the "operator" account.
func (inst *OperatorProcessFundWithdrawalJob) GetOperatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *OperatorProcessFundWithdrawalJob) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *OperatorProcessFundWithdrawalJob {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(receiptTokenMint).WRITE()
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *OperatorProcessFundWithdrawalJob) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetReceiptTokenProgramAccount sets the "receipt_token_program" account.
func (inst *OperatorProcessFundWithdrawalJob) SetReceiptTokenProgramAccount(receiptTokenProgram ag_solanago.PublicKey) *OperatorProcessFundWithdrawalJob {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(receiptTokenProgram)
	return inst
}

// GetReceiptTokenProgramAccount gets the "receipt_token_program" account.
func (inst *OperatorProcessFundWithdrawalJob) GetReceiptTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetReceiptTokenLockAuthorityAccount sets the "receipt_token_lock_authority" account.
func (inst *OperatorProcessFundWithdrawalJob) SetReceiptTokenLockAuthorityAccount(receiptTokenLockAuthority ag_solanago.PublicKey) *OperatorProcessFundWithdrawalJob {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(receiptTokenLockAuthority)
	return inst
}

func (inst *OperatorProcessFundWithdrawalJob) findFindReceiptTokenLockAuthorityAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: receipt_token_lock_authority
	seeds = append(seeds, []byte{byte(0x72), byte(0x65), byte(0x63), byte(0x65), byte(0x69), byte(0x70), byte(0x74), byte(0x5f), byte(0x74), byte(0x6f), byte(0x6b), byte(0x65), byte(0x6e), byte(0x5f), byte(0x6c), byte(0x6f), byte(0x63), byte(0x6b), byte(0x5f), byte(0x61), byte(0x75), byte(0x74), byte(0x68), byte(0x6f), byte(0x72), byte(0x69), byte(0x74), byte(0x79)})
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

// FindReceiptTokenLockAuthorityAddressWithBumpSeed calculates ReceiptTokenLockAuthority account address with given seeds and a known bump seed.
func (inst *OperatorProcessFundWithdrawalJob) FindReceiptTokenLockAuthorityAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindReceiptTokenLockAuthorityAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *OperatorProcessFundWithdrawalJob) MustFindReceiptTokenLockAuthorityAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindReceiptTokenLockAuthorityAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindReceiptTokenLockAuthorityAddress finds ReceiptTokenLockAuthority account address with given seeds.
func (inst *OperatorProcessFundWithdrawalJob) FindReceiptTokenLockAuthorityAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindReceiptTokenLockAuthorityAddress(receiptTokenMint, 0)
	return
}

func (inst *OperatorProcessFundWithdrawalJob) MustFindReceiptTokenLockAuthorityAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindReceiptTokenLockAuthorityAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetReceiptTokenLockAuthorityAccount gets the "receipt_token_lock_authority" account.
func (inst *OperatorProcessFundWithdrawalJob) GetReceiptTokenLockAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetReceiptTokenLockAccountAccount sets the "receipt_token_lock_account" account.
func (inst *OperatorProcessFundWithdrawalJob) SetReceiptTokenLockAccountAccount(receiptTokenLockAccount ag_solanago.PublicKey) *OperatorProcessFundWithdrawalJob {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(receiptTokenLockAccount).WRITE()
	return inst
}

func (inst *OperatorProcessFundWithdrawalJob) findFindReceiptTokenLockAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: receipt_token_lock
	seeds = append(seeds, []byte{byte(0x72), byte(0x65), byte(0x63), byte(0x65), byte(0x69), byte(0x70), byte(0x74), byte(0x5f), byte(0x74), byte(0x6f), byte(0x6b), byte(0x65), byte(0x6e), byte(0x5f), byte(0x6c), byte(0x6f), byte(0x63), byte(0x6b)})
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

// FindReceiptTokenLockAccountAddressWithBumpSeed calculates ReceiptTokenLockAccount account address with given seeds and a known bump seed.
func (inst *OperatorProcessFundWithdrawalJob) FindReceiptTokenLockAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindReceiptTokenLockAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *OperatorProcessFundWithdrawalJob) MustFindReceiptTokenLockAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindReceiptTokenLockAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindReceiptTokenLockAccountAddress finds ReceiptTokenLockAccount account address with given seeds.
func (inst *OperatorProcessFundWithdrawalJob) FindReceiptTokenLockAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindReceiptTokenLockAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *OperatorProcessFundWithdrawalJob) MustFindReceiptTokenLockAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindReceiptTokenLockAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetReceiptTokenLockAccountAccount gets the "receipt_token_lock_account" account.
func (inst *OperatorProcessFundWithdrawalJob) GetReceiptTokenLockAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetFundAccountAccount sets the "fund_account" account.
func (inst *OperatorProcessFundWithdrawalJob) SetFundAccountAccount(fundAccount ag_solanago.PublicKey) *OperatorProcessFundWithdrawalJob {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(fundAccount).WRITE()
	return inst
}

func (inst *OperatorProcessFundWithdrawalJob) findFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *OperatorProcessFundWithdrawalJob) FindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *OperatorProcessFundWithdrawalJob) MustFindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundAccountAddress finds FundAccount account address with given seeds.
func (inst *OperatorProcessFundWithdrawalJob) FindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *OperatorProcessFundWithdrawalJob) MustFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundAccountAccount gets the "fund_account" account.
func (inst *OperatorProcessFundWithdrawalJob) GetFundAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst OperatorProcessFundWithdrawalJob) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_OperatorProcessFundWithdrawalJob,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst OperatorProcessFundWithdrawalJob) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *OperatorProcessFundWithdrawalJob) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Forced == nil {
			return errors.New("Forced parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Operator is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ReceiptTokenProgram is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.ReceiptTokenLockAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.ReceiptTokenLockAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.FundAccount is not set")
		}
	}
	return nil
}

func (inst *OperatorProcessFundWithdrawalJob) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("OperatorProcessFundWithdrawalJob")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Forced", *inst.Forced))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                    operator", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          receipt_token_mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("       receipt_token_program", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("receipt_token_lock_authority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("         receipt_token_lock_", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                       fund_", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj OperatorProcessFundWithdrawalJob) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Forced` param:
	err = encoder.Encode(obj.Forced)
	if err != nil {
		return err
	}
	return nil
}
func (obj *OperatorProcessFundWithdrawalJob) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Forced`:
	err = decoder.Decode(&obj.Forced)
	if err != nil {
		return err
	}
	return nil
}

// NewOperatorProcessFundWithdrawalJobInstruction declares a new OperatorProcessFundWithdrawalJob instruction with the provided parameters and accounts.
func NewOperatorProcessFundWithdrawalJobInstruction(
	// Parameters:
	forced bool,
	// Accounts:
	operator ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	receiptTokenProgram ag_solanago.PublicKey,
	receiptTokenLockAuthority ag_solanago.PublicKey,
	receiptTokenLockAccount ag_solanago.PublicKey,
	fundAccount ag_solanago.PublicKey) *OperatorProcessFundWithdrawalJob {
	return NewOperatorProcessFundWithdrawalJobInstructionBuilder().
		SetForced(forced).
		SetOperatorAccount(operator).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetReceiptTokenProgramAccount(receiptTokenProgram).
		SetReceiptTokenLockAuthorityAccount(receiptTokenLockAuthority).
		SetReceiptTokenLockAccountAccount(receiptTokenLockAccount).
		SetFundAccountAccount(fundAccount)
}
