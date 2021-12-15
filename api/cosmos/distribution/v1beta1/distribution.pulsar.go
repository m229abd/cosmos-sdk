package distributionv1beta1

import (
	fmt "fmt"
	io "io"
	reflect "reflect"
	sync "sync"

	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
)

var (
	md_Params                       protoreflect.MessageDescriptor
	fd_Params_community_tax         protoreflect.FieldDescriptor
	fd_Params_base_proposer_reward  protoreflect.FieldDescriptor
	fd_Params_bonus_proposer_reward protoreflect.FieldDescriptor
	fd_Params_withdraw_addr_enabled protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_distribution_proto_init()
	md_Params = File_cosmos_distribution_v1beta1_distribution_proto.Messages().ByName("Params")
	fd_Params_community_tax = md_Params.Fields().ByName("community_tax")
	fd_Params_base_proposer_reward = md_Params.Fields().ByName("base_proposer_reward")
	fd_Params_bonus_proposer_reward = md_Params.Fields().ByName("bonus_proposer_reward")
	fd_Params_withdraw_addr_enabled = md_Params.Fields().ByName("withdraw_addr_enabled")
}

var _ protoreflect.Message = (*fastReflection_Params)(nil)

type fastReflection_Params Params

func (x *Params) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Params)(x)
}

func (x *Params) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Params_messageType fastReflection_Params_messageType
var _ protoreflect.MessageType = fastReflection_Params_messageType{}

type fastReflection_Params_messageType struct{}

func (x fastReflection_Params_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Params)(nil)
}
func (x fastReflection_Params_messageType) New() protoreflect.Message {
	return new(fastReflection_Params)
}
func (x fastReflection_Params_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Params) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Params) Type() protoreflect.MessageType {
	return _fastReflection_Params_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Params) New() protoreflect.Message {
	return new(fastReflection_Params)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Params) Interface() protoreflect.ProtoMessage {
	return (*Params)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Params) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.CommunityTax != "" {
		value := protoreflect.ValueOfString(x.CommunityTax)
		if !f(fd_Params_community_tax, value) {
			return
		}
	}
	if x.BaseProposerReward != "" {
		value := protoreflect.ValueOfString(x.BaseProposerReward)
		if !f(fd_Params_base_proposer_reward, value) {
			return
		}
	}
	if x.BonusProposerReward != "" {
		value := protoreflect.ValueOfString(x.BonusProposerReward)
		if !f(fd_Params_bonus_proposer_reward, value) {
			return
		}
	}
	if x.WithdrawAddrEnabled != false {
		value := protoreflect.ValueOfBool(x.WithdrawAddrEnabled)
		if !f(fd_Params_withdraw_addr_enabled, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Params) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.Params.community_tax":
		return x.CommunityTax != ""
	case "cosmos.distribution.v1beta1.Params.base_proposer_reward":
		return x.BaseProposerReward != ""
	case "cosmos.distribution.v1beta1.Params.bonus_proposer_reward":
		return x.BonusProposerReward != ""
	case "cosmos.distribution.v1beta1.Params.withdraw_addr_enabled":
		return x.WithdrawAddrEnabled != false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.Params.community_tax":
		x.CommunityTax = ""
	case "cosmos.distribution.v1beta1.Params.base_proposer_reward":
		x.BaseProposerReward = ""
	case "cosmos.distribution.v1beta1.Params.bonus_proposer_reward":
		x.BonusProposerReward = ""
	case "cosmos.distribution.v1beta1.Params.withdraw_addr_enabled":
		x.WithdrawAddrEnabled = false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Params) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.Params.community_tax":
		value := x.CommunityTax
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.Params.base_proposer_reward":
		value := x.BaseProposerReward
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.Params.bonus_proposer_reward":
		value := x.BonusProposerReward
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.Params.withdraw_addr_enabled":
		value := x.WithdrawAddrEnabled
		return protoreflect.ValueOfBool(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.Params does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.Params.community_tax":
		x.CommunityTax = value.Interface().(string)
	case "cosmos.distribution.v1beta1.Params.base_proposer_reward":
		x.BaseProposerReward = value.Interface().(string)
	case "cosmos.distribution.v1beta1.Params.bonus_proposer_reward":
		x.BonusProposerReward = value.Interface().(string)
	case "cosmos.distribution.v1beta1.Params.withdraw_addr_enabled":
		x.WithdrawAddrEnabled = value.Bool()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.Params.community_tax":
		panic(fmt.Errorf("field community_tax of message cosmos.distribution.v1beta1.Params is not mutable"))
	case "cosmos.distribution.v1beta1.Params.base_proposer_reward":
		panic(fmt.Errorf("field base_proposer_reward of message cosmos.distribution.v1beta1.Params is not mutable"))
	case "cosmos.distribution.v1beta1.Params.bonus_proposer_reward":
		panic(fmt.Errorf("field bonus_proposer_reward of message cosmos.distribution.v1beta1.Params is not mutable"))
	case "cosmos.distribution.v1beta1.Params.withdraw_addr_enabled":
		panic(fmt.Errorf("field withdraw_addr_enabled of message cosmos.distribution.v1beta1.Params is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Params) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.Params.community_tax":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.Params.base_proposer_reward":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.Params.bonus_proposer_reward":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.Params.withdraw_addr_enabled":
		return protoreflect.ValueOfBool(false)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Params) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.Params", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Params) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Params) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Params) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.CommunityTax)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.BaseProposerReward)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.BonusProposerReward)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.WithdrawAddrEnabled {
			n += 2
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.WithdrawAddrEnabled {
			i--
			if x.WithdrawAddrEnabled {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x20
		}
		if len(x.BonusProposerReward) > 0 {
			i -= len(x.BonusProposerReward)
			copy(dAtA[i:], x.BonusProposerReward)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BonusProposerReward)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.BaseProposerReward) > 0 {
			i -= len(x.BaseProposerReward)
			copy(dAtA[i:], x.BaseProposerReward)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BaseProposerReward)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.CommunityTax) > 0 {
			i -= len(x.CommunityTax)
			copy(dAtA[i:], x.CommunityTax)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.CommunityTax)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CommunityTax", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.CommunityTax = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BaseProposerReward", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BaseProposerReward = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BonusProposerReward", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BonusProposerReward = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field WithdrawAddrEnabled", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.WithdrawAddrEnabled = bool(v != 0)
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_ValidatorHistoricalRewards_1_list)(nil)

type _ValidatorHistoricalRewards_1_list struct {
	list *[]*v1beta1.DecCoin
}

func (x *_ValidatorHistoricalRewards_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_ValidatorHistoricalRewards_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_ValidatorHistoricalRewards_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.DecCoin)
	(*x.list)[i] = concreteValue
}

func (x *_ValidatorHistoricalRewards_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.DecCoin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_ValidatorHistoricalRewards_1_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.DecCoin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_ValidatorHistoricalRewards_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_ValidatorHistoricalRewards_1_list) NewElement() protoreflect.Value {
	v := new(v1beta1.DecCoin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_ValidatorHistoricalRewards_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_ValidatorHistoricalRewards                         protoreflect.MessageDescriptor
	fd_ValidatorHistoricalRewards_cumulative_reward_ratio protoreflect.FieldDescriptor
	fd_ValidatorHistoricalRewards_reference_count         protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_distribution_proto_init()
	md_ValidatorHistoricalRewards = File_cosmos_distribution_v1beta1_distribution_proto.Messages().ByName("ValidatorHistoricalRewards")
	fd_ValidatorHistoricalRewards_cumulative_reward_ratio = md_ValidatorHistoricalRewards.Fields().ByName("cumulative_reward_ratio")
	fd_ValidatorHistoricalRewards_reference_count = md_ValidatorHistoricalRewards.Fields().ByName("reference_count")
}

var _ protoreflect.Message = (*fastReflection_ValidatorHistoricalRewards)(nil)

type fastReflection_ValidatorHistoricalRewards ValidatorHistoricalRewards

func (x *ValidatorHistoricalRewards) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ValidatorHistoricalRewards)(x)
}

func (x *ValidatorHistoricalRewards) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ValidatorHistoricalRewards_messageType fastReflection_ValidatorHistoricalRewards_messageType
var _ protoreflect.MessageType = fastReflection_ValidatorHistoricalRewards_messageType{}

type fastReflection_ValidatorHistoricalRewards_messageType struct{}

func (x fastReflection_ValidatorHistoricalRewards_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ValidatorHistoricalRewards)(nil)
}
func (x fastReflection_ValidatorHistoricalRewards_messageType) New() protoreflect.Message {
	return new(fastReflection_ValidatorHistoricalRewards)
}
func (x fastReflection_ValidatorHistoricalRewards_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorHistoricalRewards
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ValidatorHistoricalRewards) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorHistoricalRewards
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ValidatorHistoricalRewards) Type() protoreflect.MessageType {
	return _fastReflection_ValidatorHistoricalRewards_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ValidatorHistoricalRewards) New() protoreflect.Message {
	return new(fastReflection_ValidatorHistoricalRewards)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ValidatorHistoricalRewards) Interface() protoreflect.ProtoMessage {
	return (*ValidatorHistoricalRewards)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ValidatorHistoricalRewards) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.CumulativeRewardRatio) != 0 {
		value := protoreflect.ValueOfList(&_ValidatorHistoricalRewards_1_list{list: &x.CumulativeRewardRatio})
		if !f(fd_ValidatorHistoricalRewards_cumulative_reward_ratio, value) {
			return
		}
	}
	if x.ReferenceCount != uint32(0) {
		value := protoreflect.ValueOfUint32(x.ReferenceCount)
		if !f(fd_ValidatorHistoricalRewards_reference_count, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ValidatorHistoricalRewards) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewards.cumulative_reward_ratio":
		return len(x.CumulativeRewardRatio) != 0
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewards.reference_count":
		return x.ReferenceCount != uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorHistoricalRewards"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorHistoricalRewards does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorHistoricalRewards) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewards.cumulative_reward_ratio":
		x.CumulativeRewardRatio = nil
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewards.reference_count":
		x.ReferenceCount = uint32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorHistoricalRewards"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorHistoricalRewards does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ValidatorHistoricalRewards) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewards.cumulative_reward_ratio":
		if len(x.CumulativeRewardRatio) == 0 {
			return protoreflect.ValueOfList(&_ValidatorHistoricalRewards_1_list{})
		}
		listValue := &_ValidatorHistoricalRewards_1_list{list: &x.CumulativeRewardRatio}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewards.reference_count":
		value := x.ReferenceCount
		return protoreflect.ValueOfUint32(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorHistoricalRewards"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorHistoricalRewards does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorHistoricalRewards) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewards.cumulative_reward_ratio":
		lv := value.List()
		clv := lv.(*_ValidatorHistoricalRewards_1_list)
		x.CumulativeRewardRatio = *clv.list
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewards.reference_count":
		x.ReferenceCount = uint32(value.Uint())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorHistoricalRewards"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorHistoricalRewards does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorHistoricalRewards) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewards.cumulative_reward_ratio":
		if x.CumulativeRewardRatio == nil {
			x.CumulativeRewardRatio = []*v1beta1.DecCoin{}
		}
		value := &_ValidatorHistoricalRewards_1_list{list: &x.CumulativeRewardRatio}
		return protoreflect.ValueOfList(value)
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewards.reference_count":
		panic(fmt.Errorf("field reference_count of message cosmos.distribution.v1beta1.ValidatorHistoricalRewards is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorHistoricalRewards"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorHistoricalRewards does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ValidatorHistoricalRewards) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewards.cumulative_reward_ratio":
		list := []*v1beta1.DecCoin{}
		return protoreflect.ValueOfList(&_ValidatorHistoricalRewards_1_list{list: &list})
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewards.reference_count":
		return protoreflect.ValueOfUint32(uint32(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorHistoricalRewards"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorHistoricalRewards does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ValidatorHistoricalRewards) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.ValidatorHistoricalRewards", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ValidatorHistoricalRewards) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorHistoricalRewards) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ValidatorHistoricalRewards) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ValidatorHistoricalRewards) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ValidatorHistoricalRewards)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.CumulativeRewardRatio) > 0 {
			for _, e := range x.CumulativeRewardRatio {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.ReferenceCount != 0 {
			n += 1 + runtime.Sov(uint64(x.ReferenceCount))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ValidatorHistoricalRewards)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.ReferenceCount != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ReferenceCount))
			i--
			dAtA[i] = 0x10
		}
		if len(x.CumulativeRewardRatio) > 0 {
			for iNdEx := len(x.CumulativeRewardRatio) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.CumulativeRewardRatio[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ValidatorHistoricalRewards)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorHistoricalRewards: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorHistoricalRewards: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CumulativeRewardRatio", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.CumulativeRewardRatio = append(x.CumulativeRewardRatio, &v1beta1.DecCoin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.CumulativeRewardRatio[len(x.CumulativeRewardRatio)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ReferenceCount", wireType)
				}
				x.ReferenceCount = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ReferenceCount |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_ValidatorCurrentRewards_1_list)(nil)

type _ValidatorCurrentRewards_1_list struct {
	list *[]*v1beta1.DecCoin
}

func (x *_ValidatorCurrentRewards_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_ValidatorCurrentRewards_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_ValidatorCurrentRewards_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.DecCoin)
	(*x.list)[i] = concreteValue
}

func (x *_ValidatorCurrentRewards_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.DecCoin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_ValidatorCurrentRewards_1_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.DecCoin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_ValidatorCurrentRewards_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_ValidatorCurrentRewards_1_list) NewElement() protoreflect.Value {
	v := new(v1beta1.DecCoin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_ValidatorCurrentRewards_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_ValidatorCurrentRewards         protoreflect.MessageDescriptor
	fd_ValidatorCurrentRewards_rewards protoreflect.FieldDescriptor
	fd_ValidatorCurrentRewards_period  protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_distribution_proto_init()
	md_ValidatorCurrentRewards = File_cosmos_distribution_v1beta1_distribution_proto.Messages().ByName("ValidatorCurrentRewards")
	fd_ValidatorCurrentRewards_rewards = md_ValidatorCurrentRewards.Fields().ByName("rewards")
	fd_ValidatorCurrentRewards_period = md_ValidatorCurrentRewards.Fields().ByName("period")
}

var _ protoreflect.Message = (*fastReflection_ValidatorCurrentRewards)(nil)

type fastReflection_ValidatorCurrentRewards ValidatorCurrentRewards

func (x *ValidatorCurrentRewards) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ValidatorCurrentRewards)(x)
}

func (x *ValidatorCurrentRewards) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ValidatorCurrentRewards_messageType fastReflection_ValidatorCurrentRewards_messageType
var _ protoreflect.MessageType = fastReflection_ValidatorCurrentRewards_messageType{}

type fastReflection_ValidatorCurrentRewards_messageType struct{}

func (x fastReflection_ValidatorCurrentRewards_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ValidatorCurrentRewards)(nil)
}
func (x fastReflection_ValidatorCurrentRewards_messageType) New() protoreflect.Message {
	return new(fastReflection_ValidatorCurrentRewards)
}
func (x fastReflection_ValidatorCurrentRewards_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorCurrentRewards
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ValidatorCurrentRewards) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorCurrentRewards
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ValidatorCurrentRewards) Type() protoreflect.MessageType {
	return _fastReflection_ValidatorCurrentRewards_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ValidatorCurrentRewards) New() protoreflect.Message {
	return new(fastReflection_ValidatorCurrentRewards)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ValidatorCurrentRewards) Interface() protoreflect.ProtoMessage {
	return (*ValidatorCurrentRewards)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ValidatorCurrentRewards) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Rewards) != 0 {
		value := protoreflect.ValueOfList(&_ValidatorCurrentRewards_1_list{list: &x.Rewards})
		if !f(fd_ValidatorCurrentRewards_rewards, value) {
			return
		}
	}
	if x.Period != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Period)
		if !f(fd_ValidatorCurrentRewards_period, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ValidatorCurrentRewards) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewards.rewards":
		return len(x.Rewards) != 0
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewards.period":
		return x.Period != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorCurrentRewards"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorCurrentRewards does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorCurrentRewards) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewards.rewards":
		x.Rewards = nil
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewards.period":
		x.Period = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorCurrentRewards"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorCurrentRewards does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ValidatorCurrentRewards) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewards.rewards":
		if len(x.Rewards) == 0 {
			return protoreflect.ValueOfList(&_ValidatorCurrentRewards_1_list{})
		}
		listValue := &_ValidatorCurrentRewards_1_list{list: &x.Rewards}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewards.period":
		value := x.Period
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorCurrentRewards"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorCurrentRewards does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorCurrentRewards) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewards.rewards":
		lv := value.List()
		clv := lv.(*_ValidatorCurrentRewards_1_list)
		x.Rewards = *clv.list
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewards.period":
		x.Period = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorCurrentRewards"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorCurrentRewards does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorCurrentRewards) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewards.rewards":
		if x.Rewards == nil {
			x.Rewards = []*v1beta1.DecCoin{}
		}
		value := &_ValidatorCurrentRewards_1_list{list: &x.Rewards}
		return protoreflect.ValueOfList(value)
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewards.period":
		panic(fmt.Errorf("field period of message cosmos.distribution.v1beta1.ValidatorCurrentRewards is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorCurrentRewards"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorCurrentRewards does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ValidatorCurrentRewards) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewards.rewards":
		list := []*v1beta1.DecCoin{}
		return protoreflect.ValueOfList(&_ValidatorCurrentRewards_1_list{list: &list})
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewards.period":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorCurrentRewards"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorCurrentRewards does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ValidatorCurrentRewards) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.ValidatorCurrentRewards", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ValidatorCurrentRewards) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorCurrentRewards) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ValidatorCurrentRewards) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ValidatorCurrentRewards) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ValidatorCurrentRewards)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.Rewards) > 0 {
			for _, e := range x.Rewards {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Period != 0 {
			n += 1 + runtime.Sov(uint64(x.Period))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ValidatorCurrentRewards)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Period != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Period))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Rewards) > 0 {
			for iNdEx := len(x.Rewards) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Rewards[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ValidatorCurrentRewards)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorCurrentRewards: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorCurrentRewards: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Rewards = append(x.Rewards, &v1beta1.DecCoin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Rewards[len(x.Rewards)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
				}
				x.Period = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Period |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_ValidatorAccumulatedCommission_1_list)(nil)

type _ValidatorAccumulatedCommission_1_list struct {
	list *[]*v1beta1.DecCoin
}

func (x *_ValidatorAccumulatedCommission_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_ValidatorAccumulatedCommission_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_ValidatorAccumulatedCommission_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.DecCoin)
	(*x.list)[i] = concreteValue
}

func (x *_ValidatorAccumulatedCommission_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.DecCoin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_ValidatorAccumulatedCommission_1_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.DecCoin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_ValidatorAccumulatedCommission_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_ValidatorAccumulatedCommission_1_list) NewElement() protoreflect.Value {
	v := new(v1beta1.DecCoin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_ValidatorAccumulatedCommission_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_ValidatorAccumulatedCommission            protoreflect.MessageDescriptor
	fd_ValidatorAccumulatedCommission_commission protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_distribution_proto_init()
	md_ValidatorAccumulatedCommission = File_cosmos_distribution_v1beta1_distribution_proto.Messages().ByName("ValidatorAccumulatedCommission")
	fd_ValidatorAccumulatedCommission_commission = md_ValidatorAccumulatedCommission.Fields().ByName("commission")
}

var _ protoreflect.Message = (*fastReflection_ValidatorAccumulatedCommission)(nil)

type fastReflection_ValidatorAccumulatedCommission ValidatorAccumulatedCommission

func (x *ValidatorAccumulatedCommission) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ValidatorAccumulatedCommission)(x)
}

func (x *ValidatorAccumulatedCommission) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ValidatorAccumulatedCommission_messageType fastReflection_ValidatorAccumulatedCommission_messageType
var _ protoreflect.MessageType = fastReflection_ValidatorAccumulatedCommission_messageType{}

type fastReflection_ValidatorAccumulatedCommission_messageType struct{}

func (x fastReflection_ValidatorAccumulatedCommission_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ValidatorAccumulatedCommission)(nil)
}
func (x fastReflection_ValidatorAccumulatedCommission_messageType) New() protoreflect.Message {
	return new(fastReflection_ValidatorAccumulatedCommission)
}
func (x fastReflection_ValidatorAccumulatedCommission_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorAccumulatedCommission
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ValidatorAccumulatedCommission) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorAccumulatedCommission
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ValidatorAccumulatedCommission) Type() protoreflect.MessageType {
	return _fastReflection_ValidatorAccumulatedCommission_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ValidatorAccumulatedCommission) New() protoreflect.Message {
	return new(fastReflection_ValidatorAccumulatedCommission)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ValidatorAccumulatedCommission) Interface() protoreflect.ProtoMessage {
	return (*ValidatorAccumulatedCommission)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ValidatorAccumulatedCommission) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Commission) != 0 {
		value := protoreflect.ValueOfList(&_ValidatorAccumulatedCommission_1_list{list: &x.Commission})
		if !f(fd_ValidatorAccumulatedCommission_commission, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ValidatorAccumulatedCommission) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorAccumulatedCommission.commission":
		return len(x.Commission) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorAccumulatedCommission"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorAccumulatedCommission does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorAccumulatedCommission) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorAccumulatedCommission.commission":
		x.Commission = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorAccumulatedCommission"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorAccumulatedCommission does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ValidatorAccumulatedCommission) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorAccumulatedCommission.commission":
		if len(x.Commission) == 0 {
			return protoreflect.ValueOfList(&_ValidatorAccumulatedCommission_1_list{})
		}
		listValue := &_ValidatorAccumulatedCommission_1_list{list: &x.Commission}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorAccumulatedCommission"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorAccumulatedCommission does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorAccumulatedCommission) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorAccumulatedCommission.commission":
		lv := value.List()
		clv := lv.(*_ValidatorAccumulatedCommission_1_list)
		x.Commission = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorAccumulatedCommission"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorAccumulatedCommission does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorAccumulatedCommission) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorAccumulatedCommission.commission":
		if x.Commission == nil {
			x.Commission = []*v1beta1.DecCoin{}
		}
		value := &_ValidatorAccumulatedCommission_1_list{list: &x.Commission}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorAccumulatedCommission"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorAccumulatedCommission does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ValidatorAccumulatedCommission) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorAccumulatedCommission.commission":
		list := []*v1beta1.DecCoin{}
		return protoreflect.ValueOfList(&_ValidatorAccumulatedCommission_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorAccumulatedCommission"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorAccumulatedCommission does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ValidatorAccumulatedCommission) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.ValidatorAccumulatedCommission", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ValidatorAccumulatedCommission) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorAccumulatedCommission) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ValidatorAccumulatedCommission) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ValidatorAccumulatedCommission) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ValidatorAccumulatedCommission)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.Commission) > 0 {
			for _, e := range x.Commission {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ValidatorAccumulatedCommission)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Commission) > 0 {
			for iNdEx := len(x.Commission) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Commission[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ValidatorAccumulatedCommission)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorAccumulatedCommission: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorAccumulatedCommission: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Commission", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Commission = append(x.Commission, &v1beta1.DecCoin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Commission[len(x.Commission)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_ValidatorOutstandingRewards_1_list)(nil)

type _ValidatorOutstandingRewards_1_list struct {
	list *[]*v1beta1.DecCoin
}

func (x *_ValidatorOutstandingRewards_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_ValidatorOutstandingRewards_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_ValidatorOutstandingRewards_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.DecCoin)
	(*x.list)[i] = concreteValue
}

func (x *_ValidatorOutstandingRewards_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.DecCoin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_ValidatorOutstandingRewards_1_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.DecCoin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_ValidatorOutstandingRewards_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_ValidatorOutstandingRewards_1_list) NewElement() protoreflect.Value {
	v := new(v1beta1.DecCoin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_ValidatorOutstandingRewards_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_ValidatorOutstandingRewards         protoreflect.MessageDescriptor
	fd_ValidatorOutstandingRewards_rewards protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_distribution_proto_init()
	md_ValidatorOutstandingRewards = File_cosmos_distribution_v1beta1_distribution_proto.Messages().ByName("ValidatorOutstandingRewards")
	fd_ValidatorOutstandingRewards_rewards = md_ValidatorOutstandingRewards.Fields().ByName("rewards")
}

var _ protoreflect.Message = (*fastReflection_ValidatorOutstandingRewards)(nil)

type fastReflection_ValidatorOutstandingRewards ValidatorOutstandingRewards

func (x *ValidatorOutstandingRewards) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ValidatorOutstandingRewards)(x)
}

func (x *ValidatorOutstandingRewards) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ValidatorOutstandingRewards_messageType fastReflection_ValidatorOutstandingRewards_messageType
var _ protoreflect.MessageType = fastReflection_ValidatorOutstandingRewards_messageType{}

type fastReflection_ValidatorOutstandingRewards_messageType struct{}

func (x fastReflection_ValidatorOutstandingRewards_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ValidatorOutstandingRewards)(nil)
}
func (x fastReflection_ValidatorOutstandingRewards_messageType) New() protoreflect.Message {
	return new(fastReflection_ValidatorOutstandingRewards)
}
func (x fastReflection_ValidatorOutstandingRewards_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorOutstandingRewards
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ValidatorOutstandingRewards) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorOutstandingRewards
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ValidatorOutstandingRewards) Type() protoreflect.MessageType {
	return _fastReflection_ValidatorOutstandingRewards_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ValidatorOutstandingRewards) New() protoreflect.Message {
	return new(fastReflection_ValidatorOutstandingRewards)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ValidatorOutstandingRewards) Interface() protoreflect.ProtoMessage {
	return (*ValidatorOutstandingRewards)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ValidatorOutstandingRewards) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Rewards) != 0 {
		value := protoreflect.ValueOfList(&_ValidatorOutstandingRewards_1_list{list: &x.Rewards})
		if !f(fd_ValidatorOutstandingRewards_rewards, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ValidatorOutstandingRewards) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorOutstandingRewards.rewards":
		return len(x.Rewards) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorOutstandingRewards"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorOutstandingRewards does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorOutstandingRewards) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorOutstandingRewards.rewards":
		x.Rewards = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorOutstandingRewards"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorOutstandingRewards does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ValidatorOutstandingRewards) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorOutstandingRewards.rewards":
		if len(x.Rewards) == 0 {
			return protoreflect.ValueOfList(&_ValidatorOutstandingRewards_1_list{})
		}
		listValue := &_ValidatorOutstandingRewards_1_list{list: &x.Rewards}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorOutstandingRewards"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorOutstandingRewards does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorOutstandingRewards) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorOutstandingRewards.rewards":
		lv := value.List()
		clv := lv.(*_ValidatorOutstandingRewards_1_list)
		x.Rewards = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorOutstandingRewards"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorOutstandingRewards does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorOutstandingRewards) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorOutstandingRewards.rewards":
		if x.Rewards == nil {
			x.Rewards = []*v1beta1.DecCoin{}
		}
		value := &_ValidatorOutstandingRewards_1_list{list: &x.Rewards}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorOutstandingRewards"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorOutstandingRewards does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ValidatorOutstandingRewards) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorOutstandingRewards.rewards":
		list := []*v1beta1.DecCoin{}
		return protoreflect.ValueOfList(&_ValidatorOutstandingRewards_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorOutstandingRewards"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorOutstandingRewards does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ValidatorOutstandingRewards) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.ValidatorOutstandingRewards", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ValidatorOutstandingRewards) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorOutstandingRewards) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ValidatorOutstandingRewards) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ValidatorOutstandingRewards) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ValidatorOutstandingRewards)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.Rewards) > 0 {
			for _, e := range x.Rewards {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ValidatorOutstandingRewards)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Rewards) > 0 {
			for iNdEx := len(x.Rewards) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Rewards[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ValidatorOutstandingRewards)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorOutstandingRewards: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorOutstandingRewards: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Rewards = append(x.Rewards, &v1beta1.DecCoin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Rewards[len(x.Rewards)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_ValidatorSlashEvent                  protoreflect.MessageDescriptor
	fd_ValidatorSlashEvent_validator_period protoreflect.FieldDescriptor
	fd_ValidatorSlashEvent_fraction         protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_distribution_proto_init()
	md_ValidatorSlashEvent = File_cosmos_distribution_v1beta1_distribution_proto.Messages().ByName("ValidatorSlashEvent")
	fd_ValidatorSlashEvent_validator_period = md_ValidatorSlashEvent.Fields().ByName("validator_period")
	fd_ValidatorSlashEvent_fraction = md_ValidatorSlashEvent.Fields().ByName("fraction")
}

var _ protoreflect.Message = (*fastReflection_ValidatorSlashEvent)(nil)

type fastReflection_ValidatorSlashEvent ValidatorSlashEvent

func (x *ValidatorSlashEvent) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ValidatorSlashEvent)(x)
}

func (x *ValidatorSlashEvent) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ValidatorSlashEvent_messageType fastReflection_ValidatorSlashEvent_messageType
var _ protoreflect.MessageType = fastReflection_ValidatorSlashEvent_messageType{}

type fastReflection_ValidatorSlashEvent_messageType struct{}

func (x fastReflection_ValidatorSlashEvent_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ValidatorSlashEvent)(nil)
}
func (x fastReflection_ValidatorSlashEvent_messageType) New() protoreflect.Message {
	return new(fastReflection_ValidatorSlashEvent)
}
func (x fastReflection_ValidatorSlashEvent_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorSlashEvent
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ValidatorSlashEvent) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorSlashEvent
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ValidatorSlashEvent) Type() protoreflect.MessageType {
	return _fastReflection_ValidatorSlashEvent_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ValidatorSlashEvent) New() protoreflect.Message {
	return new(fastReflection_ValidatorSlashEvent)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ValidatorSlashEvent) Interface() protoreflect.ProtoMessage {
	return (*ValidatorSlashEvent)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ValidatorSlashEvent) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ValidatorPeriod != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ValidatorPeriod)
		if !f(fd_ValidatorSlashEvent_validator_period, value) {
			return
		}
	}
	if x.Fraction != "" {
		value := protoreflect.ValueOfString(x.Fraction)
		if !f(fd_ValidatorSlashEvent_fraction, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ValidatorSlashEvent) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorSlashEvent.validator_period":
		return x.ValidatorPeriod != uint64(0)
	case "cosmos.distribution.v1beta1.ValidatorSlashEvent.fraction":
		return x.Fraction != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorSlashEvent"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorSlashEvent does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorSlashEvent) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorSlashEvent.validator_period":
		x.ValidatorPeriod = uint64(0)
	case "cosmos.distribution.v1beta1.ValidatorSlashEvent.fraction":
		x.Fraction = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorSlashEvent"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorSlashEvent does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ValidatorSlashEvent) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorSlashEvent.validator_period":
		value := x.ValidatorPeriod
		return protoreflect.ValueOfUint64(value)
	case "cosmos.distribution.v1beta1.ValidatorSlashEvent.fraction":
		value := x.Fraction
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorSlashEvent"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorSlashEvent does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorSlashEvent) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorSlashEvent.validator_period":
		x.ValidatorPeriod = value.Uint()
	case "cosmos.distribution.v1beta1.ValidatorSlashEvent.fraction":
		x.Fraction = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorSlashEvent"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorSlashEvent does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorSlashEvent) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorSlashEvent.validator_period":
		panic(fmt.Errorf("field validator_period of message cosmos.distribution.v1beta1.ValidatorSlashEvent is not mutable"))
	case "cosmos.distribution.v1beta1.ValidatorSlashEvent.fraction":
		panic(fmt.Errorf("field fraction of message cosmos.distribution.v1beta1.ValidatorSlashEvent is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorSlashEvent"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorSlashEvent does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ValidatorSlashEvent) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorSlashEvent.validator_period":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.distribution.v1beta1.ValidatorSlashEvent.fraction":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorSlashEvent"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorSlashEvent does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ValidatorSlashEvent) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.ValidatorSlashEvent", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ValidatorSlashEvent) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorSlashEvent) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ValidatorSlashEvent) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ValidatorSlashEvent) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ValidatorSlashEvent)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.ValidatorPeriod != 0 {
			n += 1 + runtime.Sov(uint64(x.ValidatorPeriod))
		}
		l = len(x.Fraction)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ValidatorSlashEvent)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Fraction) > 0 {
			i -= len(x.Fraction)
			copy(dAtA[i:], x.Fraction)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Fraction)))
			i--
			dAtA[i] = 0x12
		}
		if x.ValidatorPeriod != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ValidatorPeriod))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ValidatorSlashEvent)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorSlashEvent: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorSlashEvent: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorPeriod", wireType)
				}
				x.ValidatorPeriod = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ValidatorPeriod |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Fraction", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Fraction = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_ValidatorSlashEvents_1_list)(nil)

type _ValidatorSlashEvents_1_list struct {
	list *[]*ValidatorSlashEvent
}

func (x *_ValidatorSlashEvents_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_ValidatorSlashEvents_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_ValidatorSlashEvents_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ValidatorSlashEvent)
	(*x.list)[i] = concreteValue
}

func (x *_ValidatorSlashEvents_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ValidatorSlashEvent)
	*x.list = append(*x.list, concreteValue)
}

func (x *_ValidatorSlashEvents_1_list) AppendMutable() protoreflect.Value {
	v := new(ValidatorSlashEvent)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_ValidatorSlashEvents_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_ValidatorSlashEvents_1_list) NewElement() protoreflect.Value {
	v := new(ValidatorSlashEvent)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_ValidatorSlashEvents_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_ValidatorSlashEvents                        protoreflect.MessageDescriptor
	fd_ValidatorSlashEvents_validator_slash_events protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_distribution_proto_init()
	md_ValidatorSlashEvents = File_cosmos_distribution_v1beta1_distribution_proto.Messages().ByName("ValidatorSlashEvents")
	fd_ValidatorSlashEvents_validator_slash_events = md_ValidatorSlashEvents.Fields().ByName("validator_slash_events")
}

var _ protoreflect.Message = (*fastReflection_ValidatorSlashEvents)(nil)

type fastReflection_ValidatorSlashEvents ValidatorSlashEvents

func (x *ValidatorSlashEvents) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ValidatorSlashEvents)(x)
}

func (x *ValidatorSlashEvents) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ValidatorSlashEvents_messageType fastReflection_ValidatorSlashEvents_messageType
var _ protoreflect.MessageType = fastReflection_ValidatorSlashEvents_messageType{}

type fastReflection_ValidatorSlashEvents_messageType struct{}

func (x fastReflection_ValidatorSlashEvents_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ValidatorSlashEvents)(nil)
}
func (x fastReflection_ValidatorSlashEvents_messageType) New() protoreflect.Message {
	return new(fastReflection_ValidatorSlashEvents)
}
func (x fastReflection_ValidatorSlashEvents_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorSlashEvents
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ValidatorSlashEvents) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorSlashEvents
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ValidatorSlashEvents) Type() protoreflect.MessageType {
	return _fastReflection_ValidatorSlashEvents_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ValidatorSlashEvents) New() protoreflect.Message {
	return new(fastReflection_ValidatorSlashEvents)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ValidatorSlashEvents) Interface() protoreflect.ProtoMessage {
	return (*ValidatorSlashEvents)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ValidatorSlashEvents) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.ValidatorSlashEvents) != 0 {
		value := protoreflect.ValueOfList(&_ValidatorSlashEvents_1_list{list: &x.ValidatorSlashEvents})
		if !f(fd_ValidatorSlashEvents_validator_slash_events, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ValidatorSlashEvents) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorSlashEvents.validator_slash_events":
		return len(x.ValidatorSlashEvents) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorSlashEvents"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorSlashEvents does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorSlashEvents) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorSlashEvents.validator_slash_events":
		x.ValidatorSlashEvents = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorSlashEvents"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorSlashEvents does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ValidatorSlashEvents) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorSlashEvents.validator_slash_events":
		if len(x.ValidatorSlashEvents) == 0 {
			return protoreflect.ValueOfList(&_ValidatorSlashEvents_1_list{})
		}
		listValue := &_ValidatorSlashEvents_1_list{list: &x.ValidatorSlashEvents}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorSlashEvents"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorSlashEvents does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorSlashEvents) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorSlashEvents.validator_slash_events":
		lv := value.List()
		clv := lv.(*_ValidatorSlashEvents_1_list)
		x.ValidatorSlashEvents = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorSlashEvents"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorSlashEvents does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorSlashEvents) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorSlashEvents.validator_slash_events":
		if x.ValidatorSlashEvents == nil {
			x.ValidatorSlashEvents = []*ValidatorSlashEvent{}
		}
		value := &_ValidatorSlashEvents_1_list{list: &x.ValidatorSlashEvents}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorSlashEvents"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorSlashEvents does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ValidatorSlashEvents) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorSlashEvents.validator_slash_events":
		list := []*ValidatorSlashEvent{}
		return protoreflect.ValueOfList(&_ValidatorSlashEvents_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorSlashEvents"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorSlashEvents does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ValidatorSlashEvents) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.ValidatorSlashEvents", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ValidatorSlashEvents) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorSlashEvents) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ValidatorSlashEvents) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ValidatorSlashEvents) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ValidatorSlashEvents)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.ValidatorSlashEvents) > 0 {
			for _, e := range x.ValidatorSlashEvents {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ValidatorSlashEvents)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.ValidatorSlashEvents) > 0 {
			for iNdEx := len(x.ValidatorSlashEvents) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.ValidatorSlashEvents[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ValidatorSlashEvents)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorSlashEvents: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorSlashEvents: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorSlashEvents", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ValidatorSlashEvents = append(x.ValidatorSlashEvents, &ValidatorSlashEvent{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ValidatorSlashEvents[len(x.ValidatorSlashEvents)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_FeePool_1_list)(nil)

type _FeePool_1_list struct {
	list *[]*v1beta1.DecCoin
}

func (x *_FeePool_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_FeePool_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_FeePool_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.DecCoin)
	(*x.list)[i] = concreteValue
}

func (x *_FeePool_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.DecCoin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_FeePool_1_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.DecCoin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_FeePool_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_FeePool_1_list) NewElement() protoreflect.Value {
	v := new(v1beta1.DecCoin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_FeePool_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_FeePool                protoreflect.MessageDescriptor
	fd_FeePool_community_pool protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_distribution_proto_init()
	md_FeePool = File_cosmos_distribution_v1beta1_distribution_proto.Messages().ByName("FeePool")
	fd_FeePool_community_pool = md_FeePool.Fields().ByName("community_pool")
}

var _ protoreflect.Message = (*fastReflection_FeePool)(nil)

type fastReflection_FeePool FeePool

func (x *FeePool) ProtoReflect() protoreflect.Message {
	return (*fastReflection_FeePool)(x)
}

func (x *FeePool) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_FeePool_messageType fastReflection_FeePool_messageType
var _ protoreflect.MessageType = fastReflection_FeePool_messageType{}

type fastReflection_FeePool_messageType struct{}

func (x fastReflection_FeePool_messageType) Zero() protoreflect.Message {
	return (*fastReflection_FeePool)(nil)
}
func (x fastReflection_FeePool_messageType) New() protoreflect.Message {
	return new(fastReflection_FeePool)
}
func (x fastReflection_FeePool_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_FeePool
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_FeePool) Descriptor() protoreflect.MessageDescriptor {
	return md_FeePool
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_FeePool) Type() protoreflect.MessageType {
	return _fastReflection_FeePool_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_FeePool) New() protoreflect.Message {
	return new(fastReflection_FeePool)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_FeePool) Interface() protoreflect.ProtoMessage {
	return (*FeePool)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_FeePool) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.CommunityPool) != 0 {
		value := protoreflect.ValueOfList(&_FeePool_1_list{list: &x.CommunityPool})
		if !f(fd_FeePool_community_pool, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_FeePool) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.FeePool.community_pool":
		return len(x.CommunityPool) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.FeePool"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.FeePool does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_FeePool) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.FeePool.community_pool":
		x.CommunityPool = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.FeePool"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.FeePool does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_FeePool) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.FeePool.community_pool":
		if len(x.CommunityPool) == 0 {
			return protoreflect.ValueOfList(&_FeePool_1_list{})
		}
		listValue := &_FeePool_1_list{list: &x.CommunityPool}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.FeePool"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.FeePool does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_FeePool) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.FeePool.community_pool":
		lv := value.List()
		clv := lv.(*_FeePool_1_list)
		x.CommunityPool = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.FeePool"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.FeePool does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_FeePool) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.FeePool.community_pool":
		if x.CommunityPool == nil {
			x.CommunityPool = []*v1beta1.DecCoin{}
		}
		value := &_FeePool_1_list{list: &x.CommunityPool}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.FeePool"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.FeePool does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_FeePool) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.FeePool.community_pool":
		list := []*v1beta1.DecCoin{}
		return protoreflect.ValueOfList(&_FeePool_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.FeePool"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.FeePool does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_FeePool) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.FeePool", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_FeePool) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_FeePool) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_FeePool) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_FeePool) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*FeePool)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.CommunityPool) > 0 {
			for _, e := range x.CommunityPool {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*FeePool)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.CommunityPool) > 0 {
			for iNdEx := len(x.CommunityPool) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.CommunityPool[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*FeePool)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: FeePool: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: FeePool: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CommunityPool", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.CommunityPool = append(x.CommunityPool, &v1beta1.DecCoin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.CommunityPool[len(x.CommunityPool)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_CommunityPoolSpendProposal_4_list)(nil)

type _CommunityPoolSpendProposal_4_list struct {
	list *[]*v1beta1.Coin
}

func (x *_CommunityPoolSpendProposal_4_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_CommunityPoolSpendProposal_4_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_CommunityPoolSpendProposal_4_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	(*x.list)[i] = concreteValue
}

func (x *_CommunityPoolSpendProposal_4_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_CommunityPoolSpendProposal_4_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.Coin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_CommunityPoolSpendProposal_4_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_CommunityPoolSpendProposal_4_list) NewElement() protoreflect.Value {
	v := new(v1beta1.Coin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_CommunityPoolSpendProposal_4_list) IsValid() bool {
	return x.list != nil
}

var (
	md_CommunityPoolSpendProposal             protoreflect.MessageDescriptor
	fd_CommunityPoolSpendProposal_title       protoreflect.FieldDescriptor
	fd_CommunityPoolSpendProposal_description protoreflect.FieldDescriptor
	fd_CommunityPoolSpendProposal_recipient   protoreflect.FieldDescriptor
	fd_CommunityPoolSpendProposal_amount      protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_distribution_proto_init()
	md_CommunityPoolSpendProposal = File_cosmos_distribution_v1beta1_distribution_proto.Messages().ByName("CommunityPoolSpendProposal")
	fd_CommunityPoolSpendProposal_title = md_CommunityPoolSpendProposal.Fields().ByName("title")
	fd_CommunityPoolSpendProposal_description = md_CommunityPoolSpendProposal.Fields().ByName("description")
	fd_CommunityPoolSpendProposal_recipient = md_CommunityPoolSpendProposal.Fields().ByName("recipient")
	fd_CommunityPoolSpendProposal_amount = md_CommunityPoolSpendProposal.Fields().ByName("amount")
}

var _ protoreflect.Message = (*fastReflection_CommunityPoolSpendProposal)(nil)

type fastReflection_CommunityPoolSpendProposal CommunityPoolSpendProposal

func (x *CommunityPoolSpendProposal) ProtoReflect() protoreflect.Message {
	return (*fastReflection_CommunityPoolSpendProposal)(x)
}

func (x *CommunityPoolSpendProposal) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_CommunityPoolSpendProposal_messageType fastReflection_CommunityPoolSpendProposal_messageType
var _ protoreflect.MessageType = fastReflection_CommunityPoolSpendProposal_messageType{}

type fastReflection_CommunityPoolSpendProposal_messageType struct{}

func (x fastReflection_CommunityPoolSpendProposal_messageType) Zero() protoreflect.Message {
	return (*fastReflection_CommunityPoolSpendProposal)(nil)
}
func (x fastReflection_CommunityPoolSpendProposal_messageType) New() protoreflect.Message {
	return new(fastReflection_CommunityPoolSpendProposal)
}
func (x fastReflection_CommunityPoolSpendProposal_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_CommunityPoolSpendProposal
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_CommunityPoolSpendProposal) Descriptor() protoreflect.MessageDescriptor {
	return md_CommunityPoolSpendProposal
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_CommunityPoolSpendProposal) Type() protoreflect.MessageType {
	return _fastReflection_CommunityPoolSpendProposal_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_CommunityPoolSpendProposal) New() protoreflect.Message {
	return new(fastReflection_CommunityPoolSpendProposal)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_CommunityPoolSpendProposal) Interface() protoreflect.ProtoMessage {
	return (*CommunityPoolSpendProposal)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_CommunityPoolSpendProposal) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Title != "" {
		value := protoreflect.ValueOfString(x.Title)
		if !f(fd_CommunityPoolSpendProposal_title, value) {
			return
		}
	}
	if x.Description != "" {
		value := protoreflect.ValueOfString(x.Description)
		if !f(fd_CommunityPoolSpendProposal_description, value) {
			return
		}
	}
	if x.Recipient != "" {
		value := protoreflect.ValueOfString(x.Recipient)
		if !f(fd_CommunityPoolSpendProposal_recipient, value) {
			return
		}
	}
	if len(x.Amount) != 0 {
		value := protoreflect.ValueOfList(&_CommunityPoolSpendProposal_4_list{list: &x.Amount})
		if !f(fd_CommunityPoolSpendProposal_amount, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_CommunityPoolSpendProposal) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.title":
		return x.Title != ""
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.description":
		return x.Description != ""
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.recipient":
		return x.Recipient != ""
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.amount":
		return len(x.Amount) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.CommunityPoolSpendProposal"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.CommunityPoolSpendProposal does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CommunityPoolSpendProposal) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.title":
		x.Title = ""
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.description":
		x.Description = ""
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.recipient":
		x.Recipient = ""
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.amount":
		x.Amount = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.CommunityPoolSpendProposal"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.CommunityPoolSpendProposal does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_CommunityPoolSpendProposal) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.title":
		value := x.Title
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.description":
		value := x.Description
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.recipient":
		value := x.Recipient
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.amount":
		if len(x.Amount) == 0 {
			return protoreflect.ValueOfList(&_CommunityPoolSpendProposal_4_list{})
		}
		listValue := &_CommunityPoolSpendProposal_4_list{list: &x.Amount}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.CommunityPoolSpendProposal"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.CommunityPoolSpendProposal does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CommunityPoolSpendProposal) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.title":
		x.Title = value.Interface().(string)
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.description":
		x.Description = value.Interface().(string)
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.recipient":
		x.Recipient = value.Interface().(string)
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.amount":
		lv := value.List()
		clv := lv.(*_CommunityPoolSpendProposal_4_list)
		x.Amount = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.CommunityPoolSpendProposal"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.CommunityPoolSpendProposal does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CommunityPoolSpendProposal) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.amount":
		if x.Amount == nil {
			x.Amount = []*v1beta1.Coin{}
		}
		value := &_CommunityPoolSpendProposal_4_list{list: &x.Amount}
		return protoreflect.ValueOfList(value)
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.title":
		panic(fmt.Errorf("field title of message cosmos.distribution.v1beta1.CommunityPoolSpendProposal is not mutable"))
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.description":
		panic(fmt.Errorf("field description of message cosmos.distribution.v1beta1.CommunityPoolSpendProposal is not mutable"))
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.recipient":
		panic(fmt.Errorf("field recipient of message cosmos.distribution.v1beta1.CommunityPoolSpendProposal is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.CommunityPoolSpendProposal"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.CommunityPoolSpendProposal does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_CommunityPoolSpendProposal) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.title":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.description":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.recipient":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposal.amount":
		list := []*v1beta1.Coin{}
		return protoreflect.ValueOfList(&_CommunityPoolSpendProposal_4_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.CommunityPoolSpendProposal"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.CommunityPoolSpendProposal does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_CommunityPoolSpendProposal) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.CommunityPoolSpendProposal", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_CommunityPoolSpendProposal) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CommunityPoolSpendProposal) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_CommunityPoolSpendProposal) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_CommunityPoolSpendProposal) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*CommunityPoolSpendProposal)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Title)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Description)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Recipient)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Amount) > 0 {
			for _, e := range x.Amount {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*CommunityPoolSpendProposal)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Amount) > 0 {
			for iNdEx := len(x.Amount) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Amount[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x22
			}
		}
		if len(x.Recipient) > 0 {
			i -= len(x.Recipient)
			copy(dAtA[i:], x.Recipient)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Recipient)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Description) > 0 {
			i -= len(x.Description)
			copy(dAtA[i:], x.Description)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Description)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Title) > 0 {
			i -= len(x.Title)
			copy(dAtA[i:], x.Title)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Title)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*CommunityPoolSpendProposal)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: CommunityPoolSpendProposal: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: CommunityPoolSpendProposal: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Title = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Description = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Recipient = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Amount = append(x.Amount, &v1beta1.Coin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Amount[len(x.Amount)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_DelegatorStartingInfo                 protoreflect.MessageDescriptor
	fd_DelegatorStartingInfo_previous_period protoreflect.FieldDescriptor
	fd_DelegatorStartingInfo_stake           protoreflect.FieldDescriptor
	fd_DelegatorStartingInfo_height          protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_distribution_proto_init()
	md_DelegatorStartingInfo = File_cosmos_distribution_v1beta1_distribution_proto.Messages().ByName("DelegatorStartingInfo")
	fd_DelegatorStartingInfo_previous_period = md_DelegatorStartingInfo.Fields().ByName("previous_period")
	fd_DelegatorStartingInfo_stake = md_DelegatorStartingInfo.Fields().ByName("stake")
	fd_DelegatorStartingInfo_height = md_DelegatorStartingInfo.Fields().ByName("height")
}

var _ protoreflect.Message = (*fastReflection_DelegatorStartingInfo)(nil)

type fastReflection_DelegatorStartingInfo DelegatorStartingInfo

func (x *DelegatorStartingInfo) ProtoReflect() protoreflect.Message {
	return (*fastReflection_DelegatorStartingInfo)(x)
}

func (x *DelegatorStartingInfo) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_DelegatorStartingInfo_messageType fastReflection_DelegatorStartingInfo_messageType
var _ protoreflect.MessageType = fastReflection_DelegatorStartingInfo_messageType{}

type fastReflection_DelegatorStartingInfo_messageType struct{}

func (x fastReflection_DelegatorStartingInfo_messageType) Zero() protoreflect.Message {
	return (*fastReflection_DelegatorStartingInfo)(nil)
}
func (x fastReflection_DelegatorStartingInfo_messageType) New() protoreflect.Message {
	return new(fastReflection_DelegatorStartingInfo)
}
func (x fastReflection_DelegatorStartingInfo_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_DelegatorStartingInfo
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_DelegatorStartingInfo) Descriptor() protoreflect.MessageDescriptor {
	return md_DelegatorStartingInfo
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_DelegatorStartingInfo) Type() protoreflect.MessageType {
	return _fastReflection_DelegatorStartingInfo_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_DelegatorStartingInfo) New() protoreflect.Message {
	return new(fastReflection_DelegatorStartingInfo)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_DelegatorStartingInfo) Interface() protoreflect.ProtoMessage {
	return (*DelegatorStartingInfo)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_DelegatorStartingInfo) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.PreviousPeriod != uint64(0) {
		value := protoreflect.ValueOfUint64(x.PreviousPeriod)
		if !f(fd_DelegatorStartingInfo_previous_period, value) {
			return
		}
	}
	if x.Stake != "" {
		value := protoreflect.ValueOfString(x.Stake)
		if !f(fd_DelegatorStartingInfo_stake, value) {
			return
		}
	}
	if x.Height != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Height)
		if !f(fd_DelegatorStartingInfo_height, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_DelegatorStartingInfo) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.DelegatorStartingInfo.previous_period":
		return x.PreviousPeriod != uint64(0)
	case "cosmos.distribution.v1beta1.DelegatorStartingInfo.stake":
		return x.Stake != ""
	case "cosmos.distribution.v1beta1.DelegatorStartingInfo.height":
		return x.Height != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegatorStartingInfo"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegatorStartingInfo does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DelegatorStartingInfo) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.DelegatorStartingInfo.previous_period":
		x.PreviousPeriod = uint64(0)
	case "cosmos.distribution.v1beta1.DelegatorStartingInfo.stake":
		x.Stake = ""
	case "cosmos.distribution.v1beta1.DelegatorStartingInfo.height":
		x.Height = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegatorStartingInfo"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegatorStartingInfo does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_DelegatorStartingInfo) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.DelegatorStartingInfo.previous_period":
		value := x.PreviousPeriod
		return protoreflect.ValueOfUint64(value)
	case "cosmos.distribution.v1beta1.DelegatorStartingInfo.stake":
		value := x.Stake
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.DelegatorStartingInfo.height":
		value := x.Height
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegatorStartingInfo"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegatorStartingInfo does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DelegatorStartingInfo) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.DelegatorStartingInfo.previous_period":
		x.PreviousPeriod = value.Uint()
	case "cosmos.distribution.v1beta1.DelegatorStartingInfo.stake":
		x.Stake = value.Interface().(string)
	case "cosmos.distribution.v1beta1.DelegatorStartingInfo.height":
		x.Height = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegatorStartingInfo"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegatorStartingInfo does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DelegatorStartingInfo) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.DelegatorStartingInfo.previous_period":
		panic(fmt.Errorf("field previous_period of message cosmos.distribution.v1beta1.DelegatorStartingInfo is not mutable"))
	case "cosmos.distribution.v1beta1.DelegatorStartingInfo.stake":
		panic(fmt.Errorf("field stake of message cosmos.distribution.v1beta1.DelegatorStartingInfo is not mutable"))
	case "cosmos.distribution.v1beta1.DelegatorStartingInfo.height":
		panic(fmt.Errorf("field height of message cosmos.distribution.v1beta1.DelegatorStartingInfo is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegatorStartingInfo"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegatorStartingInfo does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_DelegatorStartingInfo) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.DelegatorStartingInfo.previous_period":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.distribution.v1beta1.DelegatorStartingInfo.stake":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.DelegatorStartingInfo.height":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegatorStartingInfo"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegatorStartingInfo does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_DelegatorStartingInfo) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.DelegatorStartingInfo", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_DelegatorStartingInfo) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DelegatorStartingInfo) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_DelegatorStartingInfo) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_DelegatorStartingInfo) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*DelegatorStartingInfo)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.PreviousPeriod != 0 {
			n += 1 + runtime.Sov(uint64(x.PreviousPeriod))
		}
		l = len(x.Stake)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Height != 0 {
			n += 1 + runtime.Sov(uint64(x.Height))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*DelegatorStartingInfo)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Height != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Height))
			i--
			dAtA[i] = 0x18
		}
		if len(x.Stake) > 0 {
			i -= len(x.Stake)
			copy(dAtA[i:], x.Stake)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Stake)))
			i--
			dAtA[i] = 0x12
		}
		if x.PreviousPeriod != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.PreviousPeriod))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*DelegatorStartingInfo)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DelegatorStartingInfo: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DelegatorStartingInfo: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PreviousPeriod", wireType)
				}
				x.PreviousPeriod = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.PreviousPeriod |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Stake", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Stake = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
				}
				x.Height = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Height |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_DelegationDelegatorReward_2_list)(nil)

type _DelegationDelegatorReward_2_list struct {
	list *[]*v1beta1.DecCoin
}

func (x *_DelegationDelegatorReward_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_DelegationDelegatorReward_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_DelegationDelegatorReward_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.DecCoin)
	(*x.list)[i] = concreteValue
}

func (x *_DelegationDelegatorReward_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.DecCoin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_DelegationDelegatorReward_2_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.DecCoin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_DelegationDelegatorReward_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_DelegationDelegatorReward_2_list) NewElement() protoreflect.Value {
	v := new(v1beta1.DecCoin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_DelegationDelegatorReward_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_DelegationDelegatorReward                   protoreflect.MessageDescriptor
	fd_DelegationDelegatorReward_validator_address protoreflect.FieldDescriptor
	fd_DelegationDelegatorReward_reward            protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_distribution_proto_init()
	md_DelegationDelegatorReward = File_cosmos_distribution_v1beta1_distribution_proto.Messages().ByName("DelegationDelegatorReward")
	fd_DelegationDelegatorReward_validator_address = md_DelegationDelegatorReward.Fields().ByName("validator_address")
	fd_DelegationDelegatorReward_reward = md_DelegationDelegatorReward.Fields().ByName("reward")
}

var _ protoreflect.Message = (*fastReflection_DelegationDelegatorReward)(nil)

type fastReflection_DelegationDelegatorReward DelegationDelegatorReward

func (x *DelegationDelegatorReward) ProtoReflect() protoreflect.Message {
	return (*fastReflection_DelegationDelegatorReward)(x)
}

func (x *DelegationDelegatorReward) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_DelegationDelegatorReward_messageType fastReflection_DelegationDelegatorReward_messageType
var _ protoreflect.MessageType = fastReflection_DelegationDelegatorReward_messageType{}

type fastReflection_DelegationDelegatorReward_messageType struct{}

func (x fastReflection_DelegationDelegatorReward_messageType) Zero() protoreflect.Message {
	return (*fastReflection_DelegationDelegatorReward)(nil)
}
func (x fastReflection_DelegationDelegatorReward_messageType) New() protoreflect.Message {
	return new(fastReflection_DelegationDelegatorReward)
}
func (x fastReflection_DelegationDelegatorReward_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_DelegationDelegatorReward
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_DelegationDelegatorReward) Descriptor() protoreflect.MessageDescriptor {
	return md_DelegationDelegatorReward
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_DelegationDelegatorReward) Type() protoreflect.MessageType {
	return _fastReflection_DelegationDelegatorReward_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_DelegationDelegatorReward) New() protoreflect.Message {
	return new(fastReflection_DelegationDelegatorReward)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_DelegationDelegatorReward) Interface() protoreflect.ProtoMessage {
	return (*DelegationDelegatorReward)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_DelegationDelegatorReward) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ValidatorAddress != "" {
		value := protoreflect.ValueOfString(x.ValidatorAddress)
		if !f(fd_DelegationDelegatorReward_validator_address, value) {
			return
		}
	}
	if len(x.Reward) != 0 {
		value := protoreflect.ValueOfList(&_DelegationDelegatorReward_2_list{list: &x.Reward})
		if !f(fd_DelegationDelegatorReward_reward, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_DelegationDelegatorReward) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.DelegationDelegatorReward.validator_address":
		return x.ValidatorAddress != ""
	case "cosmos.distribution.v1beta1.DelegationDelegatorReward.reward":
		return len(x.Reward) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegationDelegatorReward"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegationDelegatorReward does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DelegationDelegatorReward) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.DelegationDelegatorReward.validator_address":
		x.ValidatorAddress = ""
	case "cosmos.distribution.v1beta1.DelegationDelegatorReward.reward":
		x.Reward = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegationDelegatorReward"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegationDelegatorReward does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_DelegationDelegatorReward) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.DelegationDelegatorReward.validator_address":
		value := x.ValidatorAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.DelegationDelegatorReward.reward":
		if len(x.Reward) == 0 {
			return protoreflect.ValueOfList(&_DelegationDelegatorReward_2_list{})
		}
		listValue := &_DelegationDelegatorReward_2_list{list: &x.Reward}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegationDelegatorReward"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegationDelegatorReward does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DelegationDelegatorReward) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.DelegationDelegatorReward.validator_address":
		x.ValidatorAddress = value.Interface().(string)
	case "cosmos.distribution.v1beta1.DelegationDelegatorReward.reward":
		lv := value.List()
		clv := lv.(*_DelegationDelegatorReward_2_list)
		x.Reward = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegationDelegatorReward"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegationDelegatorReward does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DelegationDelegatorReward) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.DelegationDelegatorReward.reward":
		if x.Reward == nil {
			x.Reward = []*v1beta1.DecCoin{}
		}
		value := &_DelegationDelegatorReward_2_list{list: &x.Reward}
		return protoreflect.ValueOfList(value)
	case "cosmos.distribution.v1beta1.DelegationDelegatorReward.validator_address":
		panic(fmt.Errorf("field validator_address of message cosmos.distribution.v1beta1.DelegationDelegatorReward is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegationDelegatorReward"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegationDelegatorReward does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_DelegationDelegatorReward) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.DelegationDelegatorReward.validator_address":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.DelegationDelegatorReward.reward":
		list := []*v1beta1.DecCoin{}
		return protoreflect.ValueOfList(&_DelegationDelegatorReward_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegationDelegatorReward"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegationDelegatorReward does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_DelegationDelegatorReward) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.DelegationDelegatorReward", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_DelegationDelegatorReward) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DelegationDelegatorReward) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_DelegationDelegatorReward) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_DelegationDelegatorReward) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*DelegationDelegatorReward)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.ValidatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Reward) > 0 {
			for _, e := range x.Reward {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*DelegationDelegatorReward)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Reward) > 0 {
			for iNdEx := len(x.Reward) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Reward[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x12
			}
		}
		if len(x.ValidatorAddress) > 0 {
			i -= len(x.ValidatorAddress)
			copy(dAtA[i:], x.ValidatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorAddress)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*DelegationDelegatorReward)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DelegationDelegatorReward: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DelegationDelegatorReward: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ValidatorAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Reward", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Reward = append(x.Reward, &v1beta1.DecCoin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Reward[len(x.Reward)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_CommunityPoolSpendProposalWithDeposit             protoreflect.MessageDescriptor
	fd_CommunityPoolSpendProposalWithDeposit_title       protoreflect.FieldDescriptor
	fd_CommunityPoolSpendProposalWithDeposit_description protoreflect.FieldDescriptor
	fd_CommunityPoolSpendProposalWithDeposit_recipient   protoreflect.FieldDescriptor
	fd_CommunityPoolSpendProposalWithDeposit_amount      protoreflect.FieldDescriptor
	fd_CommunityPoolSpendProposalWithDeposit_deposit     protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_distribution_proto_init()
	md_CommunityPoolSpendProposalWithDeposit = File_cosmos_distribution_v1beta1_distribution_proto.Messages().ByName("CommunityPoolSpendProposalWithDeposit")
	fd_CommunityPoolSpendProposalWithDeposit_title = md_CommunityPoolSpendProposalWithDeposit.Fields().ByName("title")
	fd_CommunityPoolSpendProposalWithDeposit_description = md_CommunityPoolSpendProposalWithDeposit.Fields().ByName("description")
	fd_CommunityPoolSpendProposalWithDeposit_recipient = md_CommunityPoolSpendProposalWithDeposit.Fields().ByName("recipient")
	fd_CommunityPoolSpendProposalWithDeposit_amount = md_CommunityPoolSpendProposalWithDeposit.Fields().ByName("amount")
	fd_CommunityPoolSpendProposalWithDeposit_deposit = md_CommunityPoolSpendProposalWithDeposit.Fields().ByName("deposit")
}

var _ protoreflect.Message = (*fastReflection_CommunityPoolSpendProposalWithDeposit)(nil)

type fastReflection_CommunityPoolSpendProposalWithDeposit CommunityPoolSpendProposalWithDeposit

func (x *CommunityPoolSpendProposalWithDeposit) ProtoReflect() protoreflect.Message {
	return (*fastReflection_CommunityPoolSpendProposalWithDeposit)(x)
}

func (x *CommunityPoolSpendProposalWithDeposit) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_CommunityPoolSpendProposalWithDeposit_messageType fastReflection_CommunityPoolSpendProposalWithDeposit_messageType
var _ protoreflect.MessageType = fastReflection_CommunityPoolSpendProposalWithDeposit_messageType{}

type fastReflection_CommunityPoolSpendProposalWithDeposit_messageType struct{}

func (x fastReflection_CommunityPoolSpendProposalWithDeposit_messageType) Zero() protoreflect.Message {
	return (*fastReflection_CommunityPoolSpendProposalWithDeposit)(nil)
}
func (x fastReflection_CommunityPoolSpendProposalWithDeposit_messageType) New() protoreflect.Message {
	return new(fastReflection_CommunityPoolSpendProposalWithDeposit)
}
func (x fastReflection_CommunityPoolSpendProposalWithDeposit_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_CommunityPoolSpendProposalWithDeposit
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_CommunityPoolSpendProposalWithDeposit) Descriptor() protoreflect.MessageDescriptor {
	return md_CommunityPoolSpendProposalWithDeposit
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_CommunityPoolSpendProposalWithDeposit) Type() protoreflect.MessageType {
	return _fastReflection_CommunityPoolSpendProposalWithDeposit_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_CommunityPoolSpendProposalWithDeposit) New() protoreflect.Message {
	return new(fastReflection_CommunityPoolSpendProposalWithDeposit)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_CommunityPoolSpendProposalWithDeposit) Interface() protoreflect.ProtoMessage {
	return (*CommunityPoolSpendProposalWithDeposit)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_CommunityPoolSpendProposalWithDeposit) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Title != "" {
		value := protoreflect.ValueOfString(x.Title)
		if !f(fd_CommunityPoolSpendProposalWithDeposit_title, value) {
			return
		}
	}
	if x.Description != "" {
		value := protoreflect.ValueOfString(x.Description)
		if !f(fd_CommunityPoolSpendProposalWithDeposit_description, value) {
			return
		}
	}
	if x.Recipient != "" {
		value := protoreflect.ValueOfString(x.Recipient)
		if !f(fd_CommunityPoolSpendProposalWithDeposit_recipient, value) {
			return
		}
	}
	if x.Amount != "" {
		value := protoreflect.ValueOfString(x.Amount)
		if !f(fd_CommunityPoolSpendProposalWithDeposit_amount, value) {
			return
		}
	}
	if x.Deposit != "" {
		value := protoreflect.ValueOfString(x.Deposit)
		if !f(fd_CommunityPoolSpendProposalWithDeposit_deposit, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_CommunityPoolSpendProposalWithDeposit) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.title":
		return x.Title != ""
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.description":
		return x.Description != ""
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.recipient":
		return x.Recipient != ""
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.amount":
		return x.Amount != ""
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.deposit":
		return x.Deposit != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CommunityPoolSpendProposalWithDeposit) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.title":
		x.Title = ""
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.description":
		x.Description = ""
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.recipient":
		x.Recipient = ""
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.amount":
		x.Amount = ""
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.deposit":
		x.Deposit = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_CommunityPoolSpendProposalWithDeposit) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.title":
		value := x.Title
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.description":
		value := x.Description
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.recipient":
		value := x.Recipient
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.amount":
		value := x.Amount
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.deposit":
		value := x.Deposit
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CommunityPoolSpendProposalWithDeposit) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.title":
		x.Title = value.Interface().(string)
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.description":
		x.Description = value.Interface().(string)
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.recipient":
		x.Recipient = value.Interface().(string)
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.amount":
		x.Amount = value.Interface().(string)
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.deposit":
		x.Deposit = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CommunityPoolSpendProposalWithDeposit) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.title":
		panic(fmt.Errorf("field title of message cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit is not mutable"))
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.description":
		panic(fmt.Errorf("field description of message cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit is not mutable"))
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.recipient":
		panic(fmt.Errorf("field recipient of message cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit is not mutable"))
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.amount":
		panic(fmt.Errorf("field amount of message cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit is not mutable"))
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.deposit":
		panic(fmt.Errorf("field deposit of message cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_CommunityPoolSpendProposalWithDeposit) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.title":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.description":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.recipient":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.amount":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.deposit":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_CommunityPoolSpendProposalWithDeposit) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_CommunityPoolSpendProposalWithDeposit) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CommunityPoolSpendProposalWithDeposit) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_CommunityPoolSpendProposalWithDeposit) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_CommunityPoolSpendProposalWithDeposit) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*CommunityPoolSpendProposalWithDeposit)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Title)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Description)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Recipient)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Amount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Deposit)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*CommunityPoolSpendProposalWithDeposit)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Deposit) > 0 {
			i -= len(x.Deposit)
			copy(dAtA[i:], x.Deposit)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Deposit)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.Amount) > 0 {
			i -= len(x.Amount)
			copy(dAtA[i:], x.Amount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Amount)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Recipient) > 0 {
			i -= len(x.Recipient)
			copy(dAtA[i:], x.Recipient)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Recipient)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Description) > 0 {
			i -= len(x.Description)
			copy(dAtA[i:], x.Description)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Description)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Title) > 0 {
			i -= len(x.Title)
			copy(dAtA[i:], x.Title)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Title)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*CommunityPoolSpendProposalWithDeposit)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: CommunityPoolSpendProposalWithDeposit: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: CommunityPoolSpendProposalWithDeposit: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Title = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Description = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Recipient = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Amount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Deposit = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.19.1
// source: cosmos/distribution/v1beta1/distribution.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Params defines the set of params for the distribution module.
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommunityTax        string `protobuf:"bytes,1,opt,name=community_tax,json=communityTax,proto3" json:"community_tax,omitempty"`
	BaseProposerReward  string `protobuf:"bytes,2,opt,name=base_proposer_reward,json=baseProposerReward,proto3" json:"base_proposer_reward,omitempty"`
	BonusProposerReward string `protobuf:"bytes,3,opt,name=bonus_proposer_reward,json=bonusProposerReward,proto3" json:"bonus_proposer_reward,omitempty"`
	WithdrawAddrEnabled bool   `protobuf:"varint,4,opt,name=withdraw_addr_enabled,json=withdrawAddrEnabled,proto3" json:"withdraw_addr_enabled,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_distribution_proto_rawDescGZIP(), []int{0}
}

func (x *Params) GetCommunityTax() string {
	if x != nil {
		return x.CommunityTax
	}
	return ""
}

func (x *Params) GetBaseProposerReward() string {
	if x != nil {
		return x.BaseProposerReward
	}
	return ""
}

func (x *Params) GetBonusProposerReward() string {
	if x != nil {
		return x.BonusProposerReward
	}
	return ""
}

func (x *Params) GetWithdrawAddrEnabled() bool {
	if x != nil {
		return x.WithdrawAddrEnabled
	}
	return false
}

// ValidatorHistoricalRewards represents historical rewards for a validator.
// Height is implicit within the store key.
// Cumulative reward ratio is the sum from the zeroeth period
// until this period of rewards / tokens, per the spec.
// The reference count indicates the number of objects
// which might need to reference this historical entry at any point.
// ReferenceCount =
//    number of outstanding delegations which ended the associated period (and
//    might need to read that record)
//  + number of slashes which ended the associated period (and might need to
//  read that record)
//  + one per validator for the zeroeth period, set on initialization
type ValidatorHistoricalRewards struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CumulativeRewardRatio []*v1beta1.DecCoin `protobuf:"bytes,1,rep,name=cumulative_reward_ratio,json=cumulativeRewardRatio,proto3" json:"cumulative_reward_ratio,omitempty"`
	ReferenceCount        uint32             `protobuf:"varint,2,opt,name=reference_count,json=referenceCount,proto3" json:"reference_count,omitempty"`
}

func (x *ValidatorHistoricalRewards) Reset() {
	*x = ValidatorHistoricalRewards{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorHistoricalRewards) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorHistoricalRewards) ProtoMessage() {}

// Deprecated: Use ValidatorHistoricalRewards.ProtoReflect.Descriptor instead.
func (*ValidatorHistoricalRewards) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_distribution_proto_rawDescGZIP(), []int{1}
}

func (x *ValidatorHistoricalRewards) GetCumulativeRewardRatio() []*v1beta1.DecCoin {
	if x != nil {
		return x.CumulativeRewardRatio
	}
	return nil
}

func (x *ValidatorHistoricalRewards) GetReferenceCount() uint32 {
	if x != nil {
		return x.ReferenceCount
	}
	return 0
}

// ValidatorCurrentRewards represents current rewards and current
// period for a validator kept as a running counter and incremented
// each block as long as the validator's tokens remain constant.
type ValidatorCurrentRewards struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rewards []*v1beta1.DecCoin `protobuf:"bytes,1,rep,name=rewards,proto3" json:"rewards,omitempty"`
	Period  uint64             `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
}

func (x *ValidatorCurrentRewards) Reset() {
	*x = ValidatorCurrentRewards{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorCurrentRewards) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorCurrentRewards) ProtoMessage() {}

// Deprecated: Use ValidatorCurrentRewards.ProtoReflect.Descriptor instead.
func (*ValidatorCurrentRewards) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_distribution_proto_rawDescGZIP(), []int{2}
}

func (x *ValidatorCurrentRewards) GetRewards() []*v1beta1.DecCoin {
	if x != nil {
		return x.Rewards
	}
	return nil
}

func (x *ValidatorCurrentRewards) GetPeriod() uint64 {
	if x != nil {
		return x.Period
	}
	return 0
}

// ValidatorAccumulatedCommission represents accumulated commission
// for a validator kept as a running counter, can be withdrawn at any time.
type ValidatorAccumulatedCommission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commission []*v1beta1.DecCoin `protobuf:"bytes,1,rep,name=commission,proto3" json:"commission,omitempty"`
}

func (x *ValidatorAccumulatedCommission) Reset() {
	*x = ValidatorAccumulatedCommission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorAccumulatedCommission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorAccumulatedCommission) ProtoMessage() {}

// Deprecated: Use ValidatorAccumulatedCommission.ProtoReflect.Descriptor instead.
func (*ValidatorAccumulatedCommission) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_distribution_proto_rawDescGZIP(), []int{3}
}

func (x *ValidatorAccumulatedCommission) GetCommission() []*v1beta1.DecCoin {
	if x != nil {
		return x.Commission
	}
	return nil
}

// ValidatorOutstandingRewards represents outstanding (un-withdrawn) rewards
// for a validator inexpensive to track, allows simple sanity checks.
type ValidatorOutstandingRewards struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rewards []*v1beta1.DecCoin `protobuf:"bytes,1,rep,name=rewards,proto3" json:"rewards,omitempty"`
}

func (x *ValidatorOutstandingRewards) Reset() {
	*x = ValidatorOutstandingRewards{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorOutstandingRewards) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorOutstandingRewards) ProtoMessage() {}

// Deprecated: Use ValidatorOutstandingRewards.ProtoReflect.Descriptor instead.
func (*ValidatorOutstandingRewards) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_distribution_proto_rawDescGZIP(), []int{4}
}

func (x *ValidatorOutstandingRewards) GetRewards() []*v1beta1.DecCoin {
	if x != nil {
		return x.Rewards
	}
	return nil
}

// ValidatorSlashEvent represents a validator slash event.
// Height is implicit within the store key.
// This is needed to calculate appropriate amount of staking tokens
// for delegations which are withdrawn after a slash has occurred.
type ValidatorSlashEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValidatorPeriod uint64 `protobuf:"varint,1,opt,name=validator_period,json=validatorPeriod,proto3" json:"validator_period,omitempty"`
	Fraction        string `protobuf:"bytes,2,opt,name=fraction,proto3" json:"fraction,omitempty"`
}

func (x *ValidatorSlashEvent) Reset() {
	*x = ValidatorSlashEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorSlashEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorSlashEvent) ProtoMessage() {}

// Deprecated: Use ValidatorSlashEvent.ProtoReflect.Descriptor instead.
func (*ValidatorSlashEvent) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_distribution_proto_rawDescGZIP(), []int{5}
}

func (x *ValidatorSlashEvent) GetValidatorPeriod() uint64 {
	if x != nil {
		return x.ValidatorPeriod
	}
	return 0
}

func (x *ValidatorSlashEvent) GetFraction() string {
	if x != nil {
		return x.Fraction
	}
	return ""
}

// ValidatorSlashEvents is a collection of ValidatorSlashEvent messages.
type ValidatorSlashEvents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValidatorSlashEvents []*ValidatorSlashEvent `protobuf:"bytes,1,rep,name=validator_slash_events,json=validatorSlashEvents,proto3" json:"validator_slash_events,omitempty"`
}

func (x *ValidatorSlashEvents) Reset() {
	*x = ValidatorSlashEvents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorSlashEvents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorSlashEvents) ProtoMessage() {}

// Deprecated: Use ValidatorSlashEvents.ProtoReflect.Descriptor instead.
func (*ValidatorSlashEvents) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_distribution_proto_rawDescGZIP(), []int{6}
}

func (x *ValidatorSlashEvents) GetValidatorSlashEvents() []*ValidatorSlashEvent {
	if x != nil {
		return x.ValidatorSlashEvents
	}
	return nil
}

// FeePool is the global fee pool for distribution.
type FeePool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommunityPool []*v1beta1.DecCoin `protobuf:"bytes,1,rep,name=community_pool,json=communityPool,proto3" json:"community_pool,omitempty"`
}

func (x *FeePool) Reset() {
	*x = FeePool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeePool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeePool) ProtoMessage() {}

// Deprecated: Use FeePool.ProtoReflect.Descriptor instead.
func (*FeePool) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_distribution_proto_rawDescGZIP(), []int{7}
}

func (x *FeePool) GetCommunityPool() []*v1beta1.DecCoin {
	if x != nil {
		return x.CommunityPool
	}
	return nil
}

// CommunityPoolSpendProposal details a proposal for use of community funds,
// together with how many coins are proposed to be spent, and to which
// recipient account.
type CommunityPoolSpendProposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string          `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string          `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Recipient   string          `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Amount      []*v1beta1.Coin `protobuf:"bytes,4,rep,name=amount,proto3" json:"amount,omitempty"`
}

func (x *CommunityPoolSpendProposal) Reset() {
	*x = CommunityPoolSpendProposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommunityPoolSpendProposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommunityPoolSpendProposal) ProtoMessage() {}

// Deprecated: Use CommunityPoolSpendProposal.ProtoReflect.Descriptor instead.
func (*CommunityPoolSpendProposal) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_distribution_proto_rawDescGZIP(), []int{8}
}

func (x *CommunityPoolSpendProposal) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CommunityPoolSpendProposal) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CommunityPoolSpendProposal) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *CommunityPoolSpendProposal) GetAmount() []*v1beta1.Coin {
	if x != nil {
		return x.Amount
	}
	return nil
}

// DelegatorStartingInfo represents the starting info for a delegator reward
// period. It tracks the previous validator period, the delegation's amount of
// staking token, and the creation height (to check later on if any slashes have
// occurred). NOTE: Even though validators are slashed to whole staking tokens,
// the delegators within the validator may be left with less than a full token,
// thus sdk.Dec is used.
type DelegatorStartingInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreviousPeriod uint64 `protobuf:"varint,1,opt,name=previous_period,json=previousPeriod,proto3" json:"previous_period,omitempty"`
	Stake          string `protobuf:"bytes,2,opt,name=stake,proto3" json:"stake,omitempty"`
	Height         uint64 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *DelegatorStartingInfo) Reset() {
	*x = DelegatorStartingInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelegatorStartingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelegatorStartingInfo) ProtoMessage() {}

// Deprecated: Use DelegatorStartingInfo.ProtoReflect.Descriptor instead.
func (*DelegatorStartingInfo) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_distribution_proto_rawDescGZIP(), []int{9}
}

func (x *DelegatorStartingInfo) GetPreviousPeriod() uint64 {
	if x != nil {
		return x.PreviousPeriod
	}
	return 0
}

func (x *DelegatorStartingInfo) GetStake() string {
	if x != nil {
		return x.Stake
	}
	return ""
}

func (x *DelegatorStartingInfo) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

// DelegationDelegatorReward represents the properties
// of a delegator's delegation reward.
type DelegationDelegatorReward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValidatorAddress string             `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	Reward           []*v1beta1.DecCoin `protobuf:"bytes,2,rep,name=reward,proto3" json:"reward,omitempty"`
}

func (x *DelegationDelegatorReward) Reset() {
	*x = DelegationDelegatorReward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelegationDelegatorReward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelegationDelegatorReward) ProtoMessage() {}

// Deprecated: Use DelegationDelegatorReward.ProtoReflect.Descriptor instead.
func (*DelegationDelegatorReward) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_distribution_proto_rawDescGZIP(), []int{10}
}

func (x *DelegationDelegatorReward) GetValidatorAddress() string {
	if x != nil {
		return x.ValidatorAddress
	}
	return ""
}

func (x *DelegationDelegatorReward) GetReward() []*v1beta1.DecCoin {
	if x != nil {
		return x.Reward
	}
	return nil
}

// CommunityPoolSpendProposalWithDeposit defines a CommunityPoolSpendProposal
// with a deposit
type CommunityPoolSpendProposalWithDeposit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Recipient   string `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Amount      string `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Deposit     string `protobuf:"bytes,5,opt,name=deposit,proto3" json:"deposit,omitempty"`
}

func (x *CommunityPoolSpendProposalWithDeposit) Reset() {
	*x = CommunityPoolSpendProposalWithDeposit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommunityPoolSpendProposalWithDeposit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommunityPoolSpendProposalWithDeposit) ProtoMessage() {}

// Deprecated: Use CommunityPoolSpendProposalWithDeposit.ProtoReflect.Descriptor instead.
func (*CommunityPoolSpendProposalWithDeposit) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_distribution_proto_rawDescGZIP(), []int{11}
}

func (x *CommunityPoolSpendProposalWithDeposit) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CommunityPoolSpendProposalWithDeposit) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CommunityPoolSpendProposalWithDeposit) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *CommunityPoolSpendProposalWithDeposit) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *CommunityPoolSpendProposalWithDeposit) GetDeposit() string {
	if x != nil {
		return x.Deposit
	}
	return ""
}

var File_cosmos_distribution_v1beta1_distribution_proto protoreflect.FileDescriptor

var file_cosmos_distribution_v1beta1_distribution_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14, 0x67,
	0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87,
	0x03, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x61, 0x0a, 0x0d, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x61, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x3c, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63,
	0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x0c,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x54, 0x61, 0x78, 0x12, 0x6e, 0x0a, 0x14,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3c, 0xc8, 0xde, 0x1f, 0x00,
	0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x12, 0x62, 0x61, 0x73, 0x65, 0x50, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x70, 0x0a, 0x15,
	0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x5f, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3c, 0xc8, 0xde, 0x1f,
	0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0xd2, 0xb4, 0x2d, 0x0a, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x13, 0x62, 0x6f, 0x6e, 0x75, 0x73,
	0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x32,
	0x0a, 0x15, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x5f,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x77,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x41, 0x64, 0x64, 0x72, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x3a, 0x04, 0x98, 0xa0, 0x1f, 0x00, 0x22, 0xd1, 0x01, 0x0a, 0x1a, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x89, 0x01, 0x0a, 0x17, 0x63, 0x75, 0x6d, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x44, 0x65, 0x63, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x33, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f,
	0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x15, 0x63, 0x75,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x61,
	0x74, 0x69, 0x6f, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x9e, 0x01, 0x0a,
	0x17, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x6b, 0x0a, 0x07, 0x72, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x44, 0x65, 0x63, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x33, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f,
	0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x07, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0x93, 0x01,
	0x0a, 0x1e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x63, 0x63, 0x75, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x71, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x43, 0x6f,
	0x69, 0x6e, 0x42, 0x33, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44,
	0x65, 0x63, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x8a, 0x01, 0x0a, 0x1b, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x4f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x73, 0x12, 0x6b, 0x0a, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x43, 0x6f,
	0x69, 0x6e, 0x42, 0x33, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44,
	0x65, 0x63, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73,
	0x22, 0x9a, 0x01, 0x0a, 0x13, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x6c,
	0x61, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x12, 0x58, 0x0a, 0x08, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3c, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x44, 0x65, 0x63, 0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x44, 0x65, 0x63, 0x52, 0x08, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8a, 0x01,
	0x0a, 0x14, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x6c, 0x61, 0x73, 0x68,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x6c, 0x0a, 0x16, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x5f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x6c,
	0x61, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x14,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x3a, 0x04, 0x98, 0xa0, 0x1f, 0x00, 0x22, 0x83, 0x01, 0x0a, 0x07, 0x46,
	0x65, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x78, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x33, 0xc8, 0xde,
	0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x43, 0x6f, 0x69, 0x6e,
	0x73, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x6f, 0x6c,
	0x22, 0xe5, 0x01, 0x0a, 0x1a, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x50, 0x6f,
	0x6f, 0x6c, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x63, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e,
	0x42, 0x30, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x69,
	0x6e, 0x73, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x0c, 0x88, 0xa0, 0x1f, 0x00,
	0x98, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0xc1, 0x01, 0x0a, 0x15, 0x44, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x52, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x6b, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3c, 0xc8, 0xde, 0x1f, 0x00,
	0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x05, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x12,
	0x2b, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x13, 0xea, 0xde, 0x1f, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0xd7, 0x01, 0x0a,
	0x19, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x45, 0x0a, 0x11, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x69, 0x0a, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x43, 0x6f, 0x69, 0x6e, 0x42,
	0x33, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x43,
	0x6f, 0x69, 0x6e, 0x73, 0x52, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x3a, 0x08, 0x88, 0xa0,
	0x1f, 0x00, 0x98, 0xa0, 0x1f, 0x01, 0x22, 0xb9, 0x01, 0x0a, 0x25, 0x43, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00, 0x98, 0xa0,
	0x1f, 0x01, 0x42, 0x98, 0x02, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x11, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03,
	0x43, 0x44, 0x58, 0xaa, 0x02, 0x1b, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x44, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0xca, 0x02, 0x1b, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2,
	0x02, 0x27, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1d, 0x43, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x3a, 0x3a, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa8, 0xe2, 0x1e, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_distribution_v1beta1_distribution_proto_rawDescOnce sync.Once
	file_cosmos_distribution_v1beta1_distribution_proto_rawDescData = file_cosmos_distribution_v1beta1_distribution_proto_rawDesc
)

func file_cosmos_distribution_v1beta1_distribution_proto_rawDescGZIP() []byte {
	file_cosmos_distribution_v1beta1_distribution_proto_rawDescOnce.Do(func() {
		file_cosmos_distribution_v1beta1_distribution_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_distribution_v1beta1_distribution_proto_rawDescData)
	})
	return file_cosmos_distribution_v1beta1_distribution_proto_rawDescData
}

var file_cosmos_distribution_v1beta1_distribution_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_cosmos_distribution_v1beta1_distribution_proto_goTypes = []interface{}{
	(*Params)(nil),                                // 0: cosmos.distribution.v1beta1.Params
	(*ValidatorHistoricalRewards)(nil),            // 1: cosmos.distribution.v1beta1.ValidatorHistoricalRewards
	(*ValidatorCurrentRewards)(nil),               // 2: cosmos.distribution.v1beta1.ValidatorCurrentRewards
	(*ValidatorAccumulatedCommission)(nil),        // 3: cosmos.distribution.v1beta1.ValidatorAccumulatedCommission
	(*ValidatorOutstandingRewards)(nil),           // 4: cosmos.distribution.v1beta1.ValidatorOutstandingRewards
	(*ValidatorSlashEvent)(nil),                   // 5: cosmos.distribution.v1beta1.ValidatorSlashEvent
	(*ValidatorSlashEvents)(nil),                  // 6: cosmos.distribution.v1beta1.ValidatorSlashEvents
	(*FeePool)(nil),                               // 7: cosmos.distribution.v1beta1.FeePool
	(*CommunityPoolSpendProposal)(nil),            // 8: cosmos.distribution.v1beta1.CommunityPoolSpendProposal
	(*DelegatorStartingInfo)(nil),                 // 9: cosmos.distribution.v1beta1.DelegatorStartingInfo
	(*DelegationDelegatorReward)(nil),             // 10: cosmos.distribution.v1beta1.DelegationDelegatorReward
	(*CommunityPoolSpendProposalWithDeposit)(nil), // 11: cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit
	(*v1beta1.DecCoin)(nil),                       // 12: cosmos.base.v1beta1.DecCoin
	(*v1beta1.Coin)(nil),                          // 13: cosmos.base.v1beta1.Coin
}
var file_cosmos_distribution_v1beta1_distribution_proto_depIdxs = []int32{
	12, // 0: cosmos.distribution.v1beta1.ValidatorHistoricalRewards.cumulative_reward_ratio:type_name -> cosmos.base.v1beta1.DecCoin
	12, // 1: cosmos.distribution.v1beta1.ValidatorCurrentRewards.rewards:type_name -> cosmos.base.v1beta1.DecCoin
	12, // 2: cosmos.distribution.v1beta1.ValidatorAccumulatedCommission.commission:type_name -> cosmos.base.v1beta1.DecCoin
	12, // 3: cosmos.distribution.v1beta1.ValidatorOutstandingRewards.rewards:type_name -> cosmos.base.v1beta1.DecCoin
	5,  // 4: cosmos.distribution.v1beta1.ValidatorSlashEvents.validator_slash_events:type_name -> cosmos.distribution.v1beta1.ValidatorSlashEvent
	12, // 5: cosmos.distribution.v1beta1.FeePool.community_pool:type_name -> cosmos.base.v1beta1.DecCoin
	13, // 6: cosmos.distribution.v1beta1.CommunityPoolSpendProposal.amount:type_name -> cosmos.base.v1beta1.Coin
	12, // 7: cosmos.distribution.v1beta1.DelegationDelegatorReward.reward:type_name -> cosmos.base.v1beta1.DecCoin
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_cosmos_distribution_v1beta1_distribution_proto_init() }
func file_cosmos_distribution_v1beta1_distribution_proto_init() {
	if File_cosmos_distribution_v1beta1_distribution_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorHistoricalRewards); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorCurrentRewards); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorAccumulatedCommission); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorOutstandingRewards); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorSlashEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorSlashEvents); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeePool); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommunityPoolSpendProposal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelegatorStartingInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelegationDelegatorReward); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_distribution_v1beta1_distribution_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommunityPoolSpendProposalWithDeposit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_distribution_v1beta1_distribution_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_distribution_v1beta1_distribution_proto_goTypes,
		DependencyIndexes: file_cosmos_distribution_v1beta1_distribution_proto_depIdxs,
		MessageInfos:      file_cosmos_distribution_v1beta1_distribution_proto_msgTypes,
	}.Build()
	File_cosmos_distribution_v1beta1_distribution_proto = out.File
	file_cosmos_distribution_v1beta1_distribution_proto_rawDesc = nil
	file_cosmos_distribution_v1beta1_distribution_proto_goTypes = nil
	file_cosmos_distribution_v1beta1_distribution_proto_depIdxs = nil
}
