// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AdminInitializeFundAccount is the `admin_initialize_fund_account` instruction.
type AdminInitializeFundAccount struct {

	// [0] = [WRITE, SIGNER] payer
	//
	// [1] = [SIGNER] admin
	//
	// [2] = [] system_program
	//
	// [3] = [WRITE] receipt_token_mint
	//
	// [4] = [] receipt_token_program
	//
	// [5] = [WRITE] fund_account
	//
	// [6] = [] fund_receipt_token_lock_account
	//
	// [7] = [] fund_reserve_account
	//
	// [8] = [] event_authority
	//
	// [9] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAdminInitializeFundAccountInstructionBuilder creates a new `AdminInitializeFundAccount` instruction builder.
func NewAdminInitializeFundAccountInstructionBuilder() *AdminInitializeFundAccount {
	nd := &AdminInitializeFundAccount{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 10),
	}
	nd.AccountMetaSlice[1] = ag_solanago.Meta(Addresses["fragkamrANLvuZYQPcmPsCATQAabkqNGH6gxqqPG3aP"]).SIGNER()
	nd.AccountMetaSlice[2] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	nd.AccountMetaSlice[4] = ag_solanago.Meta(Addresses["TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb"])
	return nd
}

// SetPayerAccount sets the "payer" account.
func (inst *AdminInitializeFundAccount) SetPayerAccount(payer ag_solanago.PublicKey) *AdminInitializeFundAccount {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *AdminInitializeFundAccount) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAccount sets the "admin" account.
func (inst *AdminInitializeFundAccount) SetAdminAccount(admin ag_solanago.PublicKey) *AdminInitializeFundAccount {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *AdminInitializeFundAccount) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *AdminInitializeFundAccount) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *AdminInitializeFundAccount {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *AdminInitializeFundAccount) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *AdminInitializeFundAccount) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *AdminInitializeFundAccount {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(receiptTokenMint).WRITE()
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *AdminInitializeFundAccount) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetReceiptTokenProgramAccount sets the "receipt_token_program" account.
func (inst *AdminInitializeFundAccount) SetReceiptTokenProgramAccount(receiptTokenProgram ag_solanago.PublicKey) *AdminInitializeFundAccount {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(receiptTokenProgram)
	return inst
}

// GetReceiptTokenProgramAccount gets the "receipt_token_program" account.
func (inst *AdminInitializeFundAccount) GetReceiptTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetFundAccountAccount sets the "fund_account" account.
func (inst *AdminInitializeFundAccount) SetFundAccountAccount(fundAccount ag_solanago.PublicKey) *AdminInitializeFundAccount {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(fundAccount).WRITE()
	return inst
}

func (inst *AdminInitializeFundAccount) findFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *AdminInitializeFundAccount) FindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *AdminInitializeFundAccount) MustFindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundAccountAddress finds FundAccount account address with given seeds.
func (inst *AdminInitializeFundAccount) FindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *AdminInitializeFundAccount) MustFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundAccountAccount gets the "fund_account" account.
func (inst *AdminInitializeFundAccount) GetFundAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetFundReceiptTokenLockAccountAccount sets the "fund_receipt_token_lock_account" account.
func (inst *AdminInitializeFundAccount) SetFundReceiptTokenLockAccountAccount(fundReceiptTokenLockAccount ag_solanago.PublicKey) *AdminInitializeFundAccount {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(fundReceiptTokenLockAccount)
	return inst
}

func (inst *AdminInitializeFundAccount) findFindFundReceiptTokenLockAccountAddress(fundAccount ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// path: fundAccount
	seeds = append(seeds, fundAccount.Bytes())
	// path: receiptTokenProgram
	seeds = append(seeds, receiptTokenProgram.Bytes())
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())

	programID := Addresses["ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"]

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, programID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, programID)
	}
	return
}

// FindFundReceiptTokenLockAccountAddressWithBumpSeed calculates FundReceiptTokenLockAccount account address with given seeds and a known bump seed.
func (inst *AdminInitializeFundAccount) FindFundReceiptTokenLockAccountAddressWithBumpSeed(fundAccount ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundReceiptTokenLockAccountAddress(fundAccount, receiptTokenProgram, receiptTokenMint, bumpSeed)
	return
}

func (inst *AdminInitializeFundAccount) MustFindFundReceiptTokenLockAccountAddressWithBumpSeed(fundAccount ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundReceiptTokenLockAccountAddress(fundAccount, receiptTokenProgram, receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundReceiptTokenLockAccountAddress finds FundReceiptTokenLockAccount account address with given seeds.
func (inst *AdminInitializeFundAccount) FindFundReceiptTokenLockAccountAddress(fundAccount ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundReceiptTokenLockAccountAddress(fundAccount, receiptTokenProgram, receiptTokenMint, 0)
	return
}

func (inst *AdminInitializeFundAccount) MustFindFundReceiptTokenLockAccountAddress(fundAccount ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundReceiptTokenLockAccountAddress(fundAccount, receiptTokenProgram, receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundReceiptTokenLockAccountAccount gets the "fund_receipt_token_lock_account" account.
func (inst *AdminInitializeFundAccount) GetFundReceiptTokenLockAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetFundReserveAccountAccount sets the "fund_reserve_account" account.
func (inst *AdminInitializeFundAccount) SetFundReserveAccountAccount(fundReserveAccount ag_solanago.PublicKey) *AdminInitializeFundAccount {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(fundReserveAccount)
	return inst
}

func (inst *AdminInitializeFundAccount) findFindFundReserveAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: fund_reserve
	seeds = append(seeds, []byte{byte(0x66), byte(0x75), byte(0x6e), byte(0x64), byte(0x5f), byte(0x72), byte(0x65), byte(0x73), byte(0x65), byte(0x72), byte(0x76), byte(0x65)})
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

// FindFundReserveAccountAddressWithBumpSeed calculates FundReserveAccount account address with given seeds and a known bump seed.
func (inst *AdminInitializeFundAccount) FindFundReserveAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundReserveAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *AdminInitializeFundAccount) MustFindFundReserveAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundReserveAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundReserveAccountAddress finds FundReserveAccount account address with given seeds.
func (inst *AdminInitializeFundAccount) FindFundReserveAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundReserveAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *AdminInitializeFundAccount) MustFindFundReserveAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundReserveAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundReserveAccountAccount gets the "fund_reserve_account" account.
func (inst *AdminInitializeFundAccount) GetFundReserveAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *AdminInitializeFundAccount) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *AdminInitializeFundAccount {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(eventAuthority)
	return inst
}

func (inst *AdminInitializeFundAccount) findFindEventAuthorityAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *AdminInitializeFundAccount) FindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindEventAuthorityAddress(bumpSeed)
	return
}

func (inst *AdminInitializeFundAccount) MustFindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindEventAuthorityAddress finds EventAuthority account address with given seeds.
func (inst *AdminInitializeFundAccount) FindEventAuthorityAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindEventAuthorityAddress(0)
	return
}

func (inst *AdminInitializeFundAccount) MustFindEventAuthorityAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *AdminInitializeFundAccount) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetProgramAccount sets the "program" account.
func (inst *AdminInitializeFundAccount) SetProgramAccount(program ag_solanago.PublicKey) *AdminInitializeFundAccount {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *AdminInitializeFundAccount) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

func (inst AdminInitializeFundAccount) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AdminInitializeFundAccount,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AdminInitializeFundAccount) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AdminInitializeFundAccount) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.ReceiptTokenProgram is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.FundAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.FundReceiptTokenLockAccount is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.FundReserveAccount is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *AdminInitializeFundAccount) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AdminInitializeFundAccount")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=10]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                   payer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                   admin", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("          system_program", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("      receipt_token_mint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("   receipt_token_program", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                   fund_", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("fund_receipt_token_lock_", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("           fund_reserve_", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("         event_authority", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("                 program", inst.AccountMetaSlice.Get(9)))
					})
				})
		})
}

func (obj AdminInitializeFundAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *AdminInitializeFundAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewAdminInitializeFundAccountInstruction declares a new AdminInitializeFundAccount instruction with the provided parameters and accounts.
func NewAdminInitializeFundAccountInstruction(
	// Accounts:
	payer ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	receiptTokenProgram ag_solanago.PublicKey,
	fundAccount ag_solanago.PublicKey,
	fundReceiptTokenLockAccount ag_solanago.PublicKey,
	fundReserveAccount ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *AdminInitializeFundAccount {
	return NewAdminInitializeFundAccountInstructionBuilder().
		SetPayerAccount(payer).
		SetAdminAccount(admin).
		SetSystemProgramAccount(systemProgram).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetReceiptTokenProgramAccount(receiptTokenProgram).
		SetFundAccountAccount(fundAccount).
		SetFundReceiptTokenLockAccountAccount(fundReceiptTokenLockAccount).
		SetFundReserveAccountAccount(fundReserveAccount).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
