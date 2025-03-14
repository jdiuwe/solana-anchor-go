// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AdminUpdateFundAccountIfNeeded is the `admin_update_fund_account_if_needed` instruction.
type AdminUpdateFundAccountIfNeeded struct {
	DesiredAccountSize *uint32 `bin:"optional"`

	// [0] = [WRITE, SIGNER] payer
	//
	// [1] = [SIGNER] admin
	//
	// [2] = [] system_program
	//
	// [3] = [] receipt_token_mint
	//
	// [4] = [WRITE] fund_account
	//
	// [5] = [] fund_reserve_account
	//
	// [6] = [] event_authority
	//
	// [7] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAdminUpdateFundAccountIfNeededInstructionBuilder creates a new `AdminUpdateFundAccountIfNeeded` instruction builder.
func NewAdminUpdateFundAccountIfNeededInstructionBuilder() *AdminUpdateFundAccountIfNeeded {
	nd := &AdminUpdateFundAccountIfNeeded{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	nd.AccountMetaSlice[1] = ag_solanago.Meta(Addresses["fragkamrANLvuZYQPcmPsCATQAabkqNGH6gxqqPG3aP"]).SIGNER()
	nd.AccountMetaSlice[2] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	return nd
}

// SetDesiredAccountSize sets the "desired_account_size" parameter.
func (inst *AdminUpdateFundAccountIfNeeded) SetDesiredAccountSize(desired_account_size uint32) *AdminUpdateFundAccountIfNeeded {
	inst.DesiredAccountSize = &desired_account_size
	return inst
}

// SetPayerAccount sets the "payer" account.
func (inst *AdminUpdateFundAccountIfNeeded) SetPayerAccount(payer ag_solanago.PublicKey) *AdminUpdateFundAccountIfNeeded {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *AdminUpdateFundAccountIfNeeded) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAccount sets the "admin" account.
func (inst *AdminUpdateFundAccountIfNeeded) SetAdminAccount(admin ag_solanago.PublicKey) *AdminUpdateFundAccountIfNeeded {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *AdminUpdateFundAccountIfNeeded) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *AdminUpdateFundAccountIfNeeded) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *AdminUpdateFundAccountIfNeeded {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *AdminUpdateFundAccountIfNeeded) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *AdminUpdateFundAccountIfNeeded) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *AdminUpdateFundAccountIfNeeded {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(receiptTokenMint)
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *AdminUpdateFundAccountIfNeeded) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetFundAccountAccount sets the "fund_account" account.
func (inst *AdminUpdateFundAccountIfNeeded) SetFundAccountAccount(fundAccount ag_solanago.PublicKey) *AdminUpdateFundAccountIfNeeded {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(fundAccount).WRITE()
	return inst
}

func (inst *AdminUpdateFundAccountIfNeeded) findFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *AdminUpdateFundAccountIfNeeded) FindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *AdminUpdateFundAccountIfNeeded) MustFindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundAccountAddress finds FundAccount account address with given seeds.
func (inst *AdminUpdateFundAccountIfNeeded) FindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *AdminUpdateFundAccountIfNeeded) MustFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundAccountAccount gets the "fund_account" account.
func (inst *AdminUpdateFundAccountIfNeeded) GetFundAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetFundReserveAccountAccount sets the "fund_reserve_account" account.
func (inst *AdminUpdateFundAccountIfNeeded) SetFundReserveAccountAccount(fundReserveAccount ag_solanago.PublicKey) *AdminUpdateFundAccountIfNeeded {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(fundReserveAccount)
	return inst
}

func (inst *AdminUpdateFundAccountIfNeeded) findFindFundReserveAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *AdminUpdateFundAccountIfNeeded) FindFundReserveAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundReserveAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *AdminUpdateFundAccountIfNeeded) MustFindFundReserveAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundReserveAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundReserveAccountAddress finds FundReserveAccount account address with given seeds.
func (inst *AdminUpdateFundAccountIfNeeded) FindFundReserveAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundReserveAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *AdminUpdateFundAccountIfNeeded) MustFindFundReserveAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundReserveAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundReserveAccountAccount gets the "fund_reserve_account" account.
func (inst *AdminUpdateFundAccountIfNeeded) GetFundReserveAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *AdminUpdateFundAccountIfNeeded) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *AdminUpdateFundAccountIfNeeded {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(eventAuthority)
	return inst
}

func (inst *AdminUpdateFundAccountIfNeeded) findFindEventAuthorityAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *AdminUpdateFundAccountIfNeeded) FindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindEventAuthorityAddress(bumpSeed)
	return
}

func (inst *AdminUpdateFundAccountIfNeeded) MustFindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindEventAuthorityAddress finds EventAuthority account address with given seeds.
func (inst *AdminUpdateFundAccountIfNeeded) FindEventAuthorityAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindEventAuthorityAddress(0)
	return
}

func (inst *AdminUpdateFundAccountIfNeeded) MustFindEventAuthorityAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *AdminUpdateFundAccountIfNeeded) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetProgramAccount sets the "program" account.
func (inst *AdminUpdateFundAccountIfNeeded) SetProgramAccount(program ag_solanago.PublicKey) *AdminUpdateFundAccountIfNeeded {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *AdminUpdateFundAccountIfNeeded) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

func (inst AdminUpdateFundAccountIfNeeded) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AdminUpdateFundAccountIfNeeded,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AdminUpdateFundAccountIfNeeded) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AdminUpdateFundAccountIfNeeded) Validate() error {
	// Check whether all (required) parameters are set:
	{
	}

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
			return errors.New("accounts.FundAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.FundReserveAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *AdminUpdateFundAccountIfNeeded) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AdminUpdateFundAccountIfNeeded")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  DesiredAccountSize (OPT)", inst.DesiredAccountSize))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("             payer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("             admin", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    system_program", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("receipt_token_mint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("             fund_", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("     fund_reserve_", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("   event_authority", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("           program", inst.AccountMetaSlice.Get(7)))
					})
				})
		})
}

func (obj AdminUpdateFundAccountIfNeeded) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DesiredAccountSize` param (optional):
	{
		if obj.DesiredAccountSize == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.DesiredAccountSize)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (obj *AdminUpdateFundAccountIfNeeded) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DesiredAccountSize` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.DesiredAccountSize)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// NewAdminUpdateFundAccountIfNeededInstruction declares a new AdminUpdateFundAccountIfNeeded instruction with the provided parameters and accounts.
func NewAdminUpdateFundAccountIfNeededInstruction(
	// Parameters:
	desired_account_size uint32,
	// Accounts:
	payer ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	fundAccount ag_solanago.PublicKey,
	fundReserveAccount ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *AdminUpdateFundAccountIfNeeded {
	return NewAdminUpdateFundAccountIfNeededInstructionBuilder().
		SetDesiredAccountSize(desired_account_size).
		SetPayerAccount(payer).
		SetAdminAccount(admin).
		SetSystemProgramAccount(systemProgram).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetFundAccountAccount(fundAccount).
		SetFundReserveAccountAccount(fundReserveAccount).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
