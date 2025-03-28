package main

import (
	"encoding/json"
	"fmt"
	"os"

	. "github.com/dave/jennifer/jen"
	"github.com/davecgh/go-spew/spew"
	. "github.com/gagliardetto/utilz"
)

func main() {

	filenames := []string{
		// "idl_files/zero_copy.json",
		// "idl_files/typescript.json",
		// "idl_files/sysvars.json",
		"idl_files/swap.json",
		// "idl_files/pyth.json",
		// "idl_files/multisig.json",
		// "idl_files/misc.json",
		// "idl_files/lockup.json",
		// "idl_files/ido_pool.json",
		// "idl_files/events.json",
		// "idl_files/escrow.json",
		// "idl_files/errors.json",
		// "idl_files/composite.json",
		// "idl_files/chat.json",
		// "idl_files/cashiers_check.json",
		// "idl_files/counter_auth.json",
		// "idl_files/counter.json",
	}
	for _, idlFilepath := range filenames {
		Ln(LimeBG(idlFilepath))
		// idlFilepath := "/home/withparty/go/src/github.com/project-serum/anchor/examples/escrow/target/idl/escrow.json"
		idlFile, err := os.Open(idlFilepath)
		if err != nil {
			panic(err)
		}

		dec := json.NewDecoder(idlFile)

		var idl IDL

		err = dec.Decode(&idl)
		if err != nil {
			panic(err)
		}

		spew.Dump(idl)

		err = GenerateClient(idl)
		if err != nil {
			panic(err)
		}
	}
}

func typeStringToType(ts IdlTypeAsString) *Statement {
	stat := newStatement()
	switch ts {
	case IdlTypeBool:
		stat.Bool()
	case IdlTypeU8:
		stat.Uint8()
	case IdlTypeI8:
		stat.Int8()
	case IdlTypeU16:
		// TODO: some types have their implementation in github.com/dfuse-io/binary
		stat.Uint16()
	case IdlTypeI16:
		stat.Int16()
	case IdlTypeU32:
		stat.Uint32()
	case IdlTypeI32:
		stat.Int32()
	case IdlTypeU64:
		stat.Uint64()
	case IdlTypeI64:
		stat.Int64()
	case IdlTypeU128:
		stat.Qual("github.com/dfuse-io/binary", "Uint128")
	case IdlTypeI128:
		stat.Qual("github.com/dfuse-io/binary", "Int128")
	case IdlTypeBytes:
		// TODO:
		stat.Qual("github.com/dfuse-io/binary", "HexBytes")
	case IdlTypeString:
		stat.String()
	case IdlTypePublicKey:
		stat.Qual("github.com/gagliardetto/solana-go", "PublicKey")
	default:
		panic(Sf("unknown type string: %s", ts))
	}

	return stat
}

func idlTypeToType(envel IdlTypeEnvelope) *Statement {
	switch {
	case envel.IsString():
		return typeStringToType(envel.GetString())
	case envel.IsIdlTypeDefined():
		return Id(envel.GetIdlTypeDefined().Defined)
	default:
		panic(spew.Sdump(envel))
	}
}

func GenerateClient(idl IDL) error {
	// TODO:
	// - validate IDL (???)
	// - create new go file
	// - add instructions, etc.

	file := NewGoFile(idl.Name, true)

	// Instructions:
	for _, instruction := range idl.Instructions {
		insExportedName := ToCamel(instruction.Name)

		fmt.Println(RedBG(instruction.Name))

		{
			code := Empty()
			code.Commentf(
				"%s is the `%s` instruction.",
				insExportedName,
				instruction.Name,
			).Line()
			code.Type().Id(insExportedName).StructFunc(func(fieldsGroup *Group) {
				for _, arg := range instruction.Args {
					fieldsGroup.Id(ToCamel(arg.Name)).Add(
						DoGroup(func(fieldTypeGroup *Group) {
							if arg.Type.IsString() {
								fieldTypeGroup.Add(typeStringToType(arg.Type.GetString()))
							}

							if arg.Type.IsArray() {
								arr := arg.Type.GetArray()
								_ = arr

								if arr.Thing.IsString() {
									fieldTypeGroup.Index()
									fieldTypeGroup.Add(typeStringToType(arr.Thing.GetString()))
								}
							}
						}),
					)
				}

				fieldsGroup.Line()

				fieldsGroup.Qual("github.com/gagliardetto/solana-go", "AccountMetaSlice").Tag(map[string]string{
					"bin": "-",
				})
			})

			file.Add(code.Line())
		}

		if len(instruction.Accounts) > 0 {
			builderFuncName := "New" + insExportedName + "Builder"
			code := Empty()
			code.Commentf(
				"%s initializes a new %s builder.",
				builderFuncName,
				insExportedName,
			).Line()
			//
			code.Func().Id(builderFuncName).Params().Op("*").Id(insExportedName).
				BlockFunc(func(gr *Group) {
					gr.Return().Op("&").Id(insExportedName).Block(
						Id("AccountMetaSlice").Op(":").Make(Qual("github.com/gagliardetto/solana-go", "AccountMetaSlice"), Lit(len(instruction.Accounts))).Op(","),
					)
				})
			file.Add(code.Line())
		}

		{
			// Create parameters setters:
			code := Empty()
			for _, arg := range instruction.Args {
				exportedArgName := ToCamel(arg.Name)
				code.Line().Line().Func().Params(Id("ins").Op("*").Id(insExportedName)).Id("Set" + exportedArgName).
					Params(
						ListFunc(func(st *Group) {
							// Parameters:
							st.Id(arg.Name).Add(idlTypeToType(arg.Type))
							// TODO: determine the right type for the arg.

						}),
					).
					Params(
						ListFunc(func(st *Group) {
							// Results:
							st.Op("*").Id(insExportedName)
						}),
					).
					BlockFunc(func(gr *Group) {
						// Body:
						gr.Id("ins").Dot(exportedArgName).Op("=").Id(arg.Name)

						gr.Return().Id("ins")
					})
			}

			file.Add(code.Line())
		}
		{
			// Account setters/getters:
			code := Empty()
			for accountIndex, account := range instruction.Accounts {
				spew.Dump(account)
				// single account (???)
				// TODO: is this a parameter, or a hardcoded value?
				if account.IdlAccount != nil {
					exportedAccountName := ToCamel(account.IdlAccount.Name)
					lowerAccountName := ToLowerCamel(account.IdlAccount.Name)

					// Create account setters:
					code.Line().Line().Func().Params(Id("ins").Op("*").Id(insExportedName)).Id("Set" + exportedAccountName + "Account").
						Params(
							ListFunc(func(st *Group) {
								// Parameters:
								st.Id(lowerAccountName).Qual("github.com/gagliardetto/solana-go", "PublicKey")
							}),
						).
						Params(
							ListFunc(func(st *Group) {
								// Results:
								st.Op("*").Id(insExportedName)
							}),
						).
						BlockFunc(func(gr *Group) {
							// Body:
							def := Id("ins").Dot("AccountMetaSlice").Index(Lit(accountIndex)).
								Op("=").Qual("github.com/gagliardetto/solana-go", "NewMeta").Call(Id(lowerAccountName))
							if account.IdlAccount.IsMut {
								def.Dot("WRITE").Call()
							}
							if account.IdlAccount.IsSigner {
								def.Dot("SIGNER").Call()
							}

							gr.Add(def)

							gr.Return().Id("ins")
						})

					// Create account getters:
					code.Line().Line().Func().Params(Id("ins").Op("*").Id(insExportedName)).Id("Get" + exportedAccountName + "Account").
						Params(
							ListFunc(func(st *Group) {
								// Parameters:
							}),
						).
						Params(
							ListFunc(func(st *Group) {
								// Results:
								st.Op("*").Qual("github.com/gagliardetto/solana-go", "PublicKey")
							}),
						).
						BlockFunc(func(gr *Group) {
							// Body:
							gr.Id("ac").Op(":=").Id("ins").Dot("AccountMetaSlice").Index(Lit(accountIndex))

							gr.If(Id("ac").Op("==").Nil()).Block(
								Return().Nil(),
							)

							gr.Return(Op("&").Id("ac").Dot("PublicKey"))
						})
				}

				// manu account (???)
				// TODO: are these all the wanted parameter accounts, or a list of valid accounts?
				if account.IdlAccounts != nil {
					for accountIndex, account := range account.IdlAccounts.Accounts {
						exportedAccountName := ToCamel(account.IdlAccount.Name)
						lowerAccountName := ToLowerCamel(account.IdlAccount.Name)

						// Create account setters:
						code.Line().Line().Func().Params(Id("ins").Op("*").Id(insExportedName)).Id("Set" + exportedAccountName + "Account").
							Params(
								ListFunc(func(st *Group) {
									// Parameters:
									st.Id(lowerAccountName).Qual("github.com/gagliardetto/solana-go", "PublicKey")
								}),
							).
							Params(
								ListFunc(func(st *Group) {
									// Results:
									st.Op("*").Id(insExportedName)
								}),
							).
							BlockFunc(func(gr *Group) {
								// Body:
								def := Id("ins").Dot("AccountMetaSlice").Index(Lit(accountIndex)).
									Op("=").Qual("github.com/gagliardetto/solana-go", "NewMeta").Call(Id(lowerAccountName))
								if account.IdlAccount.IsMut {
									def.Dot("WRITE").Call()
								}
								if account.IdlAccount.IsSigner {
									def.Dot("SIGNER").Call()
								}

								gr.Add(def)

								gr.Return().Id("ins")
							})

						// Create account getters:
						code.Line().Line().Func().Params(Id("ins").Op("*").Id(insExportedName)).Id("Get" + exportedAccountName + "Account").
							Params(
								ListFunc(func(st *Group) {
									// Parameters:
								}),
							).
							Params(
								ListFunc(func(st *Group) {
									// Results:
									st.Op("*").Qual("github.com/gagliardetto/solana-go", "PublicKey")
								}),
							).
							BlockFunc(func(gr *Group) {
								// Body:
								gr.Id("ac").Op(":=").Id("ins").Dot("AccountMetaSlice").Index(Lit(accountIndex))

								gr.If(Id("ac").Op("==").Nil()).Block(
									Return().Nil(),
								)

								gr.Return(Op("&").Id("ac").Dot("PublicKey"))
							})
					}
				}

			}

			file.Add(code.Line())
		}
	}

	{
		// Types:
		for _, typ := range idl.Types {
			switch typ.Type.Kind {
			case IdlTypeDefTyKindStruct:
				code := Empty()
				code.Type().Id(typ.Name).StructFunc(func(fieldsGroup *Group) {
					for _, field := range *typ.Type.Fields {
						fieldsGroup.Id(ToCamel(field.Name)).Add(
							DoGroup(func(fieldTypeGroup *Group) {
								if field.Type.IsString() {
									fieldTypeGroup.Add(typeStringToType(field.Type.GetString()))
								}

								if field.Type.IsArray() {
									arr := field.Type.GetArray()
									_ = arr

									if arr.Thing.IsString() {
										fieldTypeGroup.Index()
										fieldTypeGroup.Add(typeStringToType(arr.Thing.GetString()))
									}
								}
							}),
						)
					}
				})

				file.Add(code.Line())
			case IdlTypeDefTyKindEnum:
				code := Empty()
				code.Type().Id(typ.Name).String()

				code.Line().Const().Parens(DoGroup(func(gr *Group) {
					for _, variant := range typ.Type.Variants {
						gr.Id(variant.Name).Id(typ.Name).Op("=").Lit(variant.Name).Line()
					}
					// TODO: check for fields, etc.
				}))
				file.Add(code.Line())

				// panic(Sf("not implemented: %s", spew.Sdump(typ)))
			default:
				panic(Sf("not implemented: %s", spew.Sdump(typ.Type.Kind)))
			}
		}
	}

	{
		// Accounts:
		for _, acc := range idl.Accounts {
			switch acc.Type.Kind {
			case IdlTypeDefTyKindStruct:
				code := Empty()
				code.Type().Id(acc.Name).StructFunc(func(fieldsGroup *Group) {
					for _, field := range *acc.Type.Fields {
						fieldsGroup.Id(ToCamel(field.Name)).Add(
							DoGroup(func(fieldTypeGroup *Group) {
								if field.Type.IsString() {
									fieldTypeGroup.Add(typeStringToType(field.Type.GetString()))
								}

								if field.Type.IsArray() {
									arr := field.Type.GetArray()
									_ = arr

									if arr.Thing.IsString() {
										fieldTypeGroup.Index()
										fieldTypeGroup.Add(typeStringToType(arr.Thing.GetString()))
									} else if arr.Thing.IsIdlTypeDefined() {
										fieldTypeGroup.Index()
										fieldTypeGroup.Add(Id(arr.Thing.GetIdlTypeDefined().Defined))
									} else {
										panic(spew.Sdump(arr))
									}
								}
							}),
						)
					}
				})

				file.Add(code.Line())
			case IdlTypeDefTyKindEnum:
				panic("not implemented")
			default:
				panic("not implemented")
			}
		}
	}

	{
		err := file.Render(os.Stdout)
		if err != nil {
			panic(err)
		}
	}
	return nil
}
