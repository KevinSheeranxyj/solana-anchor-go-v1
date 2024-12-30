// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
	ag_v5 "github.com/vmihailenco/msgpack/v5"
)

// FundManagerClearUserSolWithdrawalRequests is the `fund_manager_clear_user_sol_withdrawal_requests` instruction.
type FundManagerClearUserSolWithdrawalRequests struct {
	User                    *ag_solanago.PublicKey
	NumExpectedRequestsLeft *uint8

	// [0] = [SIGNER] fund_manager
	//
	// [1] = [] fund_account
	//
	// [2] = [WRITE] user_fund_account
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewFundManagerClearUserSolWithdrawalRequestsInstructionBuilder creates a new `FundManagerClearUserSolWithdrawalRequests` instruction builder.
func NewFundManagerClearUserSolWithdrawalRequestsInstructionBuilder() *FundManagerClearUserSolWithdrawalRequests {
	nd := &FundManagerClearUserSolWithdrawalRequests{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	nd.AccountMetaSlice[0] = ag_solanago.Meta(Addresses["5UpLTLA7Wjqp7qdfjuTtPcUw3aVtbqFA5Mgm34mxPNg2"]).SIGNER()
	return nd
}

// SetUser sets the "_user" parameter.
func (inst *FundManagerClearUserSolWithdrawalRequests) SetUser(_user ag_solanago.PublicKey) *FundManagerClearUserSolWithdrawalRequests {
	inst.User = &_user
	return inst
}

// SetNumExpectedRequestsLeft sets the "num_expected_requests_left" parameter.
func (inst *FundManagerClearUserSolWithdrawalRequests) SetNumExpectedRequestsLeft(num_expected_requests_left uint8) *FundManagerClearUserSolWithdrawalRequests {
	inst.NumExpectedRequestsLeft = &num_expected_requests_left
	return inst
}

// SetFundManagerAccount sets the "fund_manager" account.
func (inst *FundManagerClearUserSolWithdrawalRequests) SetFundManagerAccount(fundManager ag_solanago.PublicKey) *FundManagerClearUserSolWithdrawalRequests {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(fundManager).SIGNER()
	return inst
}

// GetFundManagerAccount gets the "fund_manager" account.
func (inst *FundManagerClearUserSolWithdrawalRequests) GetFundManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetFundAccountAccount sets the "fund_account" account.
func (inst *FundManagerClearUserSolWithdrawalRequests) SetFundAccountAccount(fundAccount ag_solanago.PublicKey) *FundManagerClearUserSolWithdrawalRequests {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(fundAccount)
	return inst
}

func (inst *FundManagerClearUserSolWithdrawalRequests) findFindFundAccountAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: fund
	seeds = append(seeds, []byte{byte(0x66), byte(0x75), byte(0x6e), byte(0x64)})
	// const (raw): [214 52 8 155 182 149 115 57 20 131 125 232 82 251 210 76 255 40 78 39 34 166 52 128 105 118 67 202 117 247 108 146]
	seeds = append(seeds, []byte{byte(0xd6), byte(0x34), byte(0x8), byte(0x9b), byte(0xb6), byte(0x95), byte(0x73), byte(0x39), byte(0x14), byte(0x83), byte(0x7d), byte(0xe8), byte(0x52), byte(0xfb), byte(0xd2), byte(0x4c), byte(0xff), byte(0x28), byte(0x4e), byte(0x27), byte(0x22), byte(0xa6), byte(0x34), byte(0x80), byte(0x69), byte(0x76), byte(0x43), byte(0xca), byte(0x75), byte(0xf7), byte(0x6c), byte(0x92)})

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindFundAccountAddressWithBumpSeed calculates FundAccount account address with given seeds and a known bump seed.
func (inst *FundManagerClearUserSolWithdrawalRequests) FindFundAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundAccountAddress(bumpSeed)
	return
}

func (inst *FundManagerClearUserSolWithdrawalRequests) MustFindFundAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundAccountAddress finds FundAccount account address with given seeds.
func (inst *FundManagerClearUserSolWithdrawalRequests) FindFundAccountAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundAccountAddress(0)
	return
}

func (inst *FundManagerClearUserSolWithdrawalRequests) MustFindFundAccountAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundAccountAccount gets the "fund_account" account.
func (inst *FundManagerClearUserSolWithdrawalRequests) GetFundAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetUserFundAccountAccount sets the "user_fund_account" account.
func (inst *FundManagerClearUserSolWithdrawalRequests) SetUserFundAccountAccount(userFundAccount ag_solanago.PublicKey) *FundManagerClearUserSolWithdrawalRequests {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(userFundAccount).WRITE()
	return inst
}

func (inst *FundManagerClearUserSolWithdrawalRequests) findFindUserFundAccountAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: user_fund
	seeds = append(seeds, []byte{byte(0x75), byte(0x73), byte(0x65), byte(0x72), byte(0x5f), byte(0x66), byte(0x75), byte(0x6e), byte(0x64)})
	// const (raw): [214 52 8 155 182 149 115 57 20 131 125 232 82 251 210 76 255 40 78 39 34 166 52 128 105 118 67 202 117 247 108 146]
	seeds = append(seeds, []byte{byte(0xd6), byte(0x34), byte(0x8), byte(0x9b), byte(0xb6), byte(0x95), byte(0x73), byte(0x39), byte(0x14), byte(0x83), byte(0x7d), byte(0xe8), byte(0x52), byte(0xfb), byte(0xd2), byte(0x4c), byte(0xff), byte(0x28), byte(0x4e), byte(0x27), byte(0x22), byte(0xa6), byte(0x34), byte(0x80), byte(0x69), byte(0x76), byte(0x43), byte(0xca), byte(0x75), byte(0xf7), byte(0x6c), byte(0x92)})
	// arg: User
	userSeed, err := ag_v5.Marshal(inst.User)
	if err != nil {
		return
	}
	seeds = append(seeds, userSeed)

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindUserFundAccountAddressWithBumpSeed calculates UserFundAccount account address with given seeds and a known bump seed.
func (inst *FundManagerClearUserSolWithdrawalRequests) FindUserFundAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindUserFundAccountAddress(bumpSeed)
	return
}

func (inst *FundManagerClearUserSolWithdrawalRequests) MustFindUserFundAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserFundAccountAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindUserFundAccountAddress finds UserFundAccount account address with given seeds.
func (inst *FundManagerClearUserSolWithdrawalRequests) FindUserFundAccountAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindUserFundAccountAddress(0)
	return
}

func (inst *FundManagerClearUserSolWithdrawalRequests) MustFindUserFundAccountAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserFundAccountAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetUserFundAccountAccount gets the "user_fund_account" account.
func (inst *FundManagerClearUserSolWithdrawalRequests) GetUserFundAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst FundManagerClearUserSolWithdrawalRequests) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_FundManagerClearUserSolWithdrawalRequests,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FundManagerClearUserSolWithdrawalRequests) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FundManagerClearUserSolWithdrawalRequests) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.User == nil {
			return errors.New("User parameter is not set")
		}
		if inst.NumExpectedRequestsLeft == nil {
			return errors.New("NumExpectedRequestsLeft parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.FundManager is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.FundAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.UserFundAccount is not set")
		}
	}
	return nil
}

func (inst *FundManagerClearUserSolWithdrawalRequests) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FundManagerClearUserSolWithdrawalRequests")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("                      User", *inst.User))
						paramsBranch.Child(ag_format.Param("   NumExpectedRequestsLeft", *inst.NumExpectedRequestsLeft))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("fund_manager", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       fund_", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("  user_fund_", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj FundManagerClearUserSolWithdrawalRequests) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `User` param:
	err = encoder.Encode(obj.User)
	if err != nil {
		return err
	}
	// Serialize `NumExpectedRequestsLeft` param:
	err = encoder.Encode(obj.NumExpectedRequestsLeft)
	if err != nil {
		return err
	}
	return nil
}
func (obj *FundManagerClearUserSolWithdrawalRequests) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `User`:
	err = decoder.Decode(&obj.User)
	if err != nil {
		return err
	}
	// Deserialize `NumExpectedRequestsLeft`:
	err = decoder.Decode(&obj.NumExpectedRequestsLeft)
	if err != nil {
		return err
	}
	return nil
}

// NewFundManagerClearUserSolWithdrawalRequestsInstruction declares a new FundManagerClearUserSolWithdrawalRequests instruction with the provided parameters and accounts.
func NewFundManagerClearUserSolWithdrawalRequestsInstruction(
	// Parameters:
	_user ag_solanago.PublicKey,
	num_expected_requests_left uint8,
	// Accounts:
	fundManager ag_solanago.PublicKey,
	fundAccount ag_solanago.PublicKey,
	userFundAccount ag_solanago.PublicKey) *FundManagerClearUserSolWithdrawalRequests {
	return NewFundManagerClearUserSolWithdrawalRequestsInstructionBuilder().
		SetUser(_user).
		SetNumExpectedRequestsLeft(num_expected_requests_left).
		SetFundManagerAccount(fundManager).
		SetFundAccountAccount(fundAccount).
		SetUserFundAccountAccount(userFundAccount)
}
