// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package dummy

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type AccountData struct {
	Data      VersionedData
	Owner     ag_solanago.PublicKey
	CreatedAt int64
}

func (obj AccountData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Data` param:
	{
		tmp := versionedDataContainer{}
		switch realvalue := obj.Data.(type) {
		case *VersionedDataV1Tuple:
			tmp.Enum = 0
			tmp.V1 = *realvalue
		case *VersionedDataV2Tuple:
			tmp.Enum = 1
			tmp.V2 = *realvalue
		}
		err := encoder.Encode(tmp)
		if err != nil {
			return err
		}
	}
	// Serialize `Owner` param:
	err = encoder.Encode(obj.Owner)
	if err != nil {
		return err
	}
	// Serialize `CreatedAt` param:
	err = encoder.Encode(obj.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (obj *AccountData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Data`:
	{
		tmp := new(versionedDataContainer)
		err := decoder.Decode(tmp)
		if err != nil {
			return err
		}
		switch tmp.Enum {
		case 0:
			obj.Data = &tmp.V1
		case 1:
			obj.Data = &tmp.V2
		default:
			return fmt.Errorf("unknown enum index: %v", tmp.Enum)
		}
	}
	// Deserialize `Owner`:
	err = decoder.Decode(&obj.Owner)
	if err != nil {
		return err
	}
	// Deserialize `CreatedAt`:
	err = decoder.Decode(&obj.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

type DataV1 struct {
	Field1 uint64
	Field2 string
}

func (obj DataV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Field1` param:
	err = encoder.Encode(obj.Field1)
	if err != nil {
		return err
	}
	// Serialize `Field2` param:
	err = encoder.Encode(obj.Field2)
	if err != nil {
		return err
	}
	return nil
}

func (obj *DataV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Field1`:
	err = decoder.Decode(&obj.Field1)
	if err != nil {
		return err
	}
	// Deserialize `Field2`:
	err = decoder.Decode(&obj.Field2)
	if err != nil {
		return err
	}
	return nil
}

type DataV2 struct {
	Field1 uint64
	Field2 uint32
	Field3 string
	Field4 bool
}

func (obj DataV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Field1` param:
	err = encoder.Encode(obj.Field1)
	if err != nil {
		return err
	}
	// Serialize `Field2` param:
	err = encoder.Encode(obj.Field2)
	if err != nil {
		return err
	}
	// Serialize `Field3` param:
	err = encoder.Encode(obj.Field3)
	if err != nil {
		return err
	}
	// Serialize `Field4` param:
	err = encoder.Encode(obj.Field4)
	if err != nil {
		return err
	}
	return nil
}

func (obj *DataV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Field1`:
	err = decoder.Decode(&obj.Field1)
	if err != nil {
		return err
	}
	// Deserialize `Field2`:
	err = decoder.Decode(&obj.Field2)
	if err != nil {
		return err
	}
	// Deserialize `Field3`:
	err = decoder.Decode(&obj.Field3)
	if err != nil {
		return err
	}
	// Deserialize `Field4`:
	err = decoder.Decode(&obj.Field4)
	if err != nil {
		return err
	}
	return nil
}

type Decremented struct {
	User   ag_solanago.PublicKey
	Token  string
	Amount uint64
}

func (obj Decremented) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `User` param:
	err = encoder.Encode(obj.User)
	if err != nil {
		return err
	}
	// Serialize `Token` param:
	err = encoder.Encode(obj.Token)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Decremented) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `User`:
	err = decoder.Decode(&obj.User)
	if err != nil {
		return err
	}
	// Deserialize `Token`:
	err = decoder.Decode(&obj.Token)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

type Incremented struct {
	User   ag_solanago.PublicKey
	Token  string
	Amount uint64
}

func (obj Incremented) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `User` param:
	err = encoder.Encode(obj.User)
	if err != nil {
		return err
	}
	// Serialize `Token` param:
	err = encoder.Encode(obj.Token)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Incremented) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `User`:
	err = decoder.Decode(&obj.User)
	if err != nil {
		return err
	}
	// Deserialize `Token`:
	err = decoder.Decode(&obj.Token)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

type InstructionRequest interface {
	isInstructionRequest()
}

type instructionRequestContainer struct {
	Enum ag_binary.BorshEnum `borsh_enum:"true"`
	V1   InstructionRequestV1Tuple
	V2   InstructionRequestV2Tuple
}

type InstructionRequestV1Tuple struct {
	Elem0 RequestV1
}

func (_ InstructionRequestV1Tuple) isInstructionRequest() {}

type InstructionRequestV2Tuple struct {
	Elem0 RequestV2
}

func (_ InstructionRequestV2Tuple) isInstructionRequest() {}

type RequestV1 struct {
	Field1 uint64
	Field2 string
}

func (obj RequestV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Field1` param:
	err = encoder.Encode(obj.Field1)
	if err != nil {
		return err
	}
	// Serialize `Field2` param:
	err = encoder.Encode(obj.Field2)
	if err != nil {
		return err
	}
	return nil
}

func (obj *RequestV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Field1`:
	err = decoder.Decode(&obj.Field1)
	if err != nil {
		return err
	}
	// Deserialize `Field2`:
	err = decoder.Decode(&obj.Field2)
	if err != nil {
		return err
	}
	return nil
}

type RequestV2 struct {
	Field1 uint64
	Field2 uint32
	Field3 string
	Field4 bool
}

func (obj RequestV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Field1` param:
	err = encoder.Encode(obj.Field1)
	if err != nil {
		return err
	}
	// Serialize `Field2` param:
	err = encoder.Encode(obj.Field2)
	if err != nil {
		return err
	}
	// Serialize `Field3` param:
	err = encoder.Encode(obj.Field3)
	if err != nil {
		return err
	}
	// Serialize `Field4` param:
	err = encoder.Encode(obj.Field4)
	if err != nil {
		return err
	}
	return nil
}

func (obj *RequestV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Field1`:
	err = decoder.Decode(&obj.Field1)
	if err != nil {
		return err
	}
	// Deserialize `Field2`:
	err = decoder.Decode(&obj.Field2)
	if err != nil {
		return err
	}
	// Deserialize `Field3`:
	err = decoder.Decode(&obj.Field3)
	if err != nil {
		return err
	}
	// Deserialize `Field4`:
	err = decoder.Decode(&obj.Field4)
	if err != nil {
		return err
	}
	return nil
}

type UserTokenAmount struct {
	User   ag_solanago.PublicKey
	Bump   uint8
	Token  string
	Amount uint64
}

func (obj UserTokenAmount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `User` param:
	err = encoder.Encode(obj.User)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `Token` param:
	err = encoder.Encode(obj.Token)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *UserTokenAmount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `User`:
	err = decoder.Decode(&obj.User)
	if err != nil {
		return err
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `Token`:
	err = decoder.Decode(&obj.Token)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

type VersionedData interface {
	isVersionedData()
}

type versionedDataContainer struct {
	Enum ag_binary.BorshEnum `borsh_enum:"true"`
	V1   VersionedDataV1Tuple
	V2   VersionedDataV2Tuple
}

type VersionedDataV1Tuple struct {
	Elem0 DataV1
}

func (_ VersionedDataV1Tuple) isVersionedData() {}

type VersionedDataV2Tuple struct {
	Elem0 DataV2
}

func (_ VersionedDataV2Tuple) isVersionedData() {}

type VersionedEventV1 struct {
	Field1 uint64
	Field2 string
}

func (obj VersionedEventV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Field1` param:
	err = encoder.Encode(obj.Field1)
	if err != nil {
		return err
	}
	// Serialize `Field2` param:
	err = encoder.Encode(obj.Field2)
	if err != nil {
		return err
	}
	return nil
}

func (obj *VersionedEventV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Field1`:
	err = decoder.Decode(&obj.Field1)
	if err != nil {
		return err
	}
	// Deserialize `Field2`:
	err = decoder.Decode(&obj.Field2)
	if err != nil {
		return err
	}
	return nil
}

type VersionedEventV2 struct {
	Field1 uint64
	Field2 uint32
	Field3 string
	Field4 bool
}

func (obj VersionedEventV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Field1` param:
	err = encoder.Encode(obj.Field1)
	if err != nil {
		return err
	}
	// Serialize `Field2` param:
	err = encoder.Encode(obj.Field2)
	if err != nil {
		return err
	}
	// Serialize `Field3` param:
	err = encoder.Encode(obj.Field3)
	if err != nil {
		return err
	}
	// Serialize `Field4` param:
	err = encoder.Encode(obj.Field4)
	if err != nil {
		return err
	}
	return nil
}

func (obj *VersionedEventV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Field1`:
	err = decoder.Decode(&obj.Field1)
	if err != nil {
		return err
	}
	// Deserialize `Field2`:
	err = decoder.Decode(&obj.Field2)
	if err != nil {
		return err
	}
	// Deserialize `Field3`:
	err = decoder.Decode(&obj.Field3)
	if err != nil {
		return err
	}
	// Deserialize `Field4`:
	err = decoder.Decode(&obj.Field4)
	if err != nil {
		return err
	}
	return nil
}

type VersionedState interface {
	isVersionedState()
}

type versionedStateContainer struct {
	Enum ag_binary.BorshEnum `borsh_enum:"true"`
	V1   VersionedStateV1Tuple
	V2   VersionedStateV2Tuple
}

type VersionedStateV1Tuple struct {
	Elem0 VersionedStateV1
}

func (_ VersionedStateV1Tuple) isVersionedState() {}

type VersionedStateV2Tuple struct {
	Elem0 VersionedStateV2
}

func (_ VersionedStateV2Tuple) isVersionedState() {}

type VersionedStateV1 struct {
	Field1 uint64
	Field2 string
}

func (obj VersionedStateV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Field1` param:
	err = encoder.Encode(obj.Field1)
	if err != nil {
		return err
	}
	// Serialize `Field2` param:
	err = encoder.Encode(obj.Field2)
	if err != nil {
		return err
	}
	return nil
}

func (obj *VersionedStateV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Field1`:
	err = decoder.Decode(&obj.Field1)
	if err != nil {
		return err
	}
	// Deserialize `Field2`:
	err = decoder.Decode(&obj.Field2)
	if err != nil {
		return err
	}
	return nil
}

type VersionedStateV2 struct {
	Field1 uint64
	Field2 uint32
	Field3 string
	Field4 bool
}

func (obj VersionedStateV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Field1` param:
	err = encoder.Encode(obj.Field1)
	if err != nil {
		return err
	}
	// Serialize `Field2` param:
	err = encoder.Encode(obj.Field2)
	if err != nil {
		return err
	}
	// Serialize `Field3` param:
	err = encoder.Encode(obj.Field3)
	if err != nil {
		return err
	}
	// Serialize `Field4` param:
	err = encoder.Encode(obj.Field4)
	if err != nil {
		return err
	}
	return nil
}

func (obj *VersionedStateV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Field1`:
	err = decoder.Decode(&obj.Field1)
	if err != nil {
		return err
	}
	// Deserialize `Field2`:
	err = decoder.Decode(&obj.Field2)
	if err != nil {
		return err
	}
	// Deserialize `Field3`:
	err = decoder.Decode(&obj.Field3)
	if err != nil {
		return err
	}
	// Deserialize `Field4`:
	err = decoder.Decode(&obj.Field4)
	if err != nil {
		return err
	}
	return nil
}
