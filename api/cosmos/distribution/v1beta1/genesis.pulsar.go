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
	md_DelegatorWithdrawInfo                   protoreflect.MessageDescriptor
	fd_DelegatorWithdrawInfo_delegator_address protoreflect.FieldDescriptor
	fd_DelegatorWithdrawInfo_withdraw_address  protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_genesis_proto_init()
	md_DelegatorWithdrawInfo = File_cosmos_distribution_v1beta1_genesis_proto.Messages().ByName("DelegatorWithdrawInfo")
	fd_DelegatorWithdrawInfo_delegator_address = md_DelegatorWithdrawInfo.Fields().ByName("delegator_address")
	fd_DelegatorWithdrawInfo_withdraw_address = md_DelegatorWithdrawInfo.Fields().ByName("withdraw_address")
}

var _ protoreflect.Message = (*fastReflection_DelegatorWithdrawInfo)(nil)

type fastReflection_DelegatorWithdrawInfo DelegatorWithdrawInfo

func (x *DelegatorWithdrawInfo) ProtoReflect() protoreflect.Message {
	return (*fastReflection_DelegatorWithdrawInfo)(x)
}

func (x *DelegatorWithdrawInfo) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_DelegatorWithdrawInfo_messageType fastReflection_DelegatorWithdrawInfo_messageType
var _ protoreflect.MessageType = fastReflection_DelegatorWithdrawInfo_messageType{}

type fastReflection_DelegatorWithdrawInfo_messageType struct{}

func (x fastReflection_DelegatorWithdrawInfo_messageType) Zero() protoreflect.Message {
	return (*fastReflection_DelegatorWithdrawInfo)(nil)
}
func (x fastReflection_DelegatorWithdrawInfo_messageType) New() protoreflect.Message {
	return new(fastReflection_DelegatorWithdrawInfo)
}
func (x fastReflection_DelegatorWithdrawInfo_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_DelegatorWithdrawInfo
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_DelegatorWithdrawInfo) Descriptor() protoreflect.MessageDescriptor {
	return md_DelegatorWithdrawInfo
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_DelegatorWithdrawInfo) Type() protoreflect.MessageType {
	return _fastReflection_DelegatorWithdrawInfo_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_DelegatorWithdrawInfo) New() protoreflect.Message {
	return new(fastReflection_DelegatorWithdrawInfo)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_DelegatorWithdrawInfo) Interface() protoreflect.ProtoMessage {
	return (*DelegatorWithdrawInfo)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_DelegatorWithdrawInfo) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.DelegatorAddress != "" {
		value := protoreflect.ValueOfString(x.DelegatorAddress)
		if !f(fd_DelegatorWithdrawInfo_delegator_address, value) {
			return
		}
	}
	if x.WithdrawAddress != "" {
		value := protoreflect.ValueOfString(x.WithdrawAddress)
		if !f(fd_DelegatorWithdrawInfo_withdraw_address, value) {
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
func (x *fastReflection_DelegatorWithdrawInfo) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.DelegatorWithdrawInfo.delegator_address":
		return x.DelegatorAddress != ""
	case "cosmos.distribution.v1beta1.DelegatorWithdrawInfo.withdraw_address":
		return x.WithdrawAddress != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegatorWithdrawInfo"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegatorWithdrawInfo does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DelegatorWithdrawInfo) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.DelegatorWithdrawInfo.delegator_address":
		x.DelegatorAddress = ""
	case "cosmos.distribution.v1beta1.DelegatorWithdrawInfo.withdraw_address":
		x.WithdrawAddress = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegatorWithdrawInfo"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegatorWithdrawInfo does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_DelegatorWithdrawInfo) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.DelegatorWithdrawInfo.delegator_address":
		value := x.DelegatorAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.DelegatorWithdrawInfo.withdraw_address":
		value := x.WithdrawAddress
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegatorWithdrawInfo"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegatorWithdrawInfo does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_DelegatorWithdrawInfo) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.DelegatorWithdrawInfo.delegator_address":
		x.DelegatorAddress = value.Interface().(string)
	case "cosmos.distribution.v1beta1.DelegatorWithdrawInfo.withdraw_address":
		x.WithdrawAddress = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegatorWithdrawInfo"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegatorWithdrawInfo does not contain field %s", fd.FullName()))
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
func (x *fastReflection_DelegatorWithdrawInfo) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.DelegatorWithdrawInfo.delegator_address":
		panic(fmt.Errorf("field delegator_address of message cosmos.distribution.v1beta1.DelegatorWithdrawInfo is not mutable"))
	case "cosmos.distribution.v1beta1.DelegatorWithdrawInfo.withdraw_address":
		panic(fmt.Errorf("field withdraw_address of message cosmos.distribution.v1beta1.DelegatorWithdrawInfo is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegatorWithdrawInfo"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegatorWithdrawInfo does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_DelegatorWithdrawInfo) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.DelegatorWithdrawInfo.delegator_address":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.DelegatorWithdrawInfo.withdraw_address":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegatorWithdrawInfo"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegatorWithdrawInfo does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_DelegatorWithdrawInfo) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.DelegatorWithdrawInfo", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_DelegatorWithdrawInfo) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DelegatorWithdrawInfo) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_DelegatorWithdrawInfo) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_DelegatorWithdrawInfo) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*DelegatorWithdrawInfo)
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
		l = len(x.DelegatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.WithdrawAddress)
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
		x := input.Message.Interface().(*DelegatorWithdrawInfo)
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
		if len(x.WithdrawAddress) > 0 {
			i -= len(x.WithdrawAddress)
			copy(dAtA[i:], x.WithdrawAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.WithdrawAddress)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.DelegatorAddress) > 0 {
			i -= len(x.DelegatorAddress)
			copy(dAtA[i:], x.DelegatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DelegatorAddress)))
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
		x := input.Message.Interface().(*DelegatorWithdrawInfo)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DelegatorWithdrawInfo: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DelegatorWithdrawInfo: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
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
				x.DelegatorAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field WithdrawAddress", wireType)
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
				x.WithdrawAddress = string(dAtA[iNdEx:postIndex])
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

var _ protoreflect.List = (*_ValidatorOutstandingRewardsRecord_2_list)(nil)

type _ValidatorOutstandingRewardsRecord_2_list struct {
	list *[]*v1beta1.DecCoin
}

func (x *_ValidatorOutstandingRewardsRecord_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_ValidatorOutstandingRewardsRecord_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_ValidatorOutstandingRewardsRecord_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.DecCoin)
	(*x.list)[i] = concreteValue
}

func (x *_ValidatorOutstandingRewardsRecord_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.DecCoin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_ValidatorOutstandingRewardsRecord_2_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.DecCoin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_ValidatorOutstandingRewardsRecord_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_ValidatorOutstandingRewardsRecord_2_list) NewElement() protoreflect.Value {
	v := new(v1beta1.DecCoin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_ValidatorOutstandingRewardsRecord_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_ValidatorOutstandingRewardsRecord                     protoreflect.MessageDescriptor
	fd_ValidatorOutstandingRewardsRecord_validator_address   protoreflect.FieldDescriptor
	fd_ValidatorOutstandingRewardsRecord_outstanding_rewards protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_genesis_proto_init()
	md_ValidatorOutstandingRewardsRecord = File_cosmos_distribution_v1beta1_genesis_proto.Messages().ByName("ValidatorOutstandingRewardsRecord")
	fd_ValidatorOutstandingRewardsRecord_validator_address = md_ValidatorOutstandingRewardsRecord.Fields().ByName("validator_address")
	fd_ValidatorOutstandingRewardsRecord_outstanding_rewards = md_ValidatorOutstandingRewardsRecord.Fields().ByName("outstanding_rewards")
}

var _ protoreflect.Message = (*fastReflection_ValidatorOutstandingRewardsRecord)(nil)

type fastReflection_ValidatorOutstandingRewardsRecord ValidatorOutstandingRewardsRecord

func (x *ValidatorOutstandingRewardsRecord) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ValidatorOutstandingRewardsRecord)(x)
}

func (x *ValidatorOutstandingRewardsRecord) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ValidatorOutstandingRewardsRecord_messageType fastReflection_ValidatorOutstandingRewardsRecord_messageType
var _ protoreflect.MessageType = fastReflection_ValidatorOutstandingRewardsRecord_messageType{}

type fastReflection_ValidatorOutstandingRewardsRecord_messageType struct{}

func (x fastReflection_ValidatorOutstandingRewardsRecord_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ValidatorOutstandingRewardsRecord)(nil)
}
func (x fastReflection_ValidatorOutstandingRewardsRecord_messageType) New() protoreflect.Message {
	return new(fastReflection_ValidatorOutstandingRewardsRecord)
}
func (x fastReflection_ValidatorOutstandingRewardsRecord_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorOutstandingRewardsRecord
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ValidatorOutstandingRewardsRecord) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorOutstandingRewardsRecord
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ValidatorOutstandingRewardsRecord) Type() protoreflect.MessageType {
	return _fastReflection_ValidatorOutstandingRewardsRecord_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ValidatorOutstandingRewardsRecord) New() protoreflect.Message {
	return new(fastReflection_ValidatorOutstandingRewardsRecord)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ValidatorOutstandingRewardsRecord) Interface() protoreflect.ProtoMessage {
	return (*ValidatorOutstandingRewardsRecord)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ValidatorOutstandingRewardsRecord) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ValidatorAddress != "" {
		value := protoreflect.ValueOfString(x.ValidatorAddress)
		if !f(fd_ValidatorOutstandingRewardsRecord_validator_address, value) {
			return
		}
	}
	if len(x.OutstandingRewards) != 0 {
		value := protoreflect.ValueOfList(&_ValidatorOutstandingRewardsRecord_2_list{list: &x.OutstandingRewards})
		if !f(fd_ValidatorOutstandingRewardsRecord_outstanding_rewards, value) {
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
func (x *fastReflection_ValidatorOutstandingRewardsRecord) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord.validator_address":
		return x.ValidatorAddress != ""
	case "cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord.outstanding_rewards":
		return len(x.OutstandingRewards) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorOutstandingRewardsRecord) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord.validator_address":
		x.ValidatorAddress = ""
	case "cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord.outstanding_rewards":
		x.OutstandingRewards = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ValidatorOutstandingRewardsRecord) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord.validator_address":
		value := x.ValidatorAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord.outstanding_rewards":
		if len(x.OutstandingRewards) == 0 {
			return protoreflect.ValueOfList(&_ValidatorOutstandingRewardsRecord_2_list{})
		}
		listValue := &_ValidatorOutstandingRewardsRecord_2_list{list: &x.OutstandingRewards}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ValidatorOutstandingRewardsRecord) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord.validator_address":
		x.ValidatorAddress = value.Interface().(string)
	case "cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord.outstanding_rewards":
		lv := value.List()
		clv := lv.(*_ValidatorOutstandingRewardsRecord_2_list)
		x.OutstandingRewards = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ValidatorOutstandingRewardsRecord) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord.outstanding_rewards":
		if x.OutstandingRewards == nil {
			x.OutstandingRewards = []*v1beta1.DecCoin{}
		}
		value := &_ValidatorOutstandingRewardsRecord_2_list{list: &x.OutstandingRewards}
		return protoreflect.ValueOfList(value)
	case "cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord.validator_address":
		panic(fmt.Errorf("field validator_address of message cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ValidatorOutstandingRewardsRecord) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord.validator_address":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord.outstanding_rewards":
		list := []*v1beta1.DecCoin{}
		return protoreflect.ValueOfList(&_ValidatorOutstandingRewardsRecord_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ValidatorOutstandingRewardsRecord) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ValidatorOutstandingRewardsRecord) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorOutstandingRewardsRecord) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ValidatorOutstandingRewardsRecord) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ValidatorOutstandingRewardsRecord) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ValidatorOutstandingRewardsRecord)
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
		if len(x.OutstandingRewards) > 0 {
			for _, e := range x.OutstandingRewards {
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
		x := input.Message.Interface().(*ValidatorOutstandingRewardsRecord)
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
		if len(x.OutstandingRewards) > 0 {
			for iNdEx := len(x.OutstandingRewards) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.OutstandingRewards[iNdEx])
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
		x := input.Message.Interface().(*ValidatorOutstandingRewardsRecord)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorOutstandingRewardsRecord: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorOutstandingRewardsRecord: illegal tag %d (wire type %d)", fieldNum, wire)
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field OutstandingRewards", wireType)
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
				x.OutstandingRewards = append(x.OutstandingRewards, &v1beta1.DecCoin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.OutstandingRewards[len(x.OutstandingRewards)-1]); err != nil {
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
	md_ValidatorAccumulatedCommissionRecord                   protoreflect.MessageDescriptor
	fd_ValidatorAccumulatedCommissionRecord_validator_address protoreflect.FieldDescriptor
	fd_ValidatorAccumulatedCommissionRecord_accumulated       protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_genesis_proto_init()
	md_ValidatorAccumulatedCommissionRecord = File_cosmos_distribution_v1beta1_genesis_proto.Messages().ByName("ValidatorAccumulatedCommissionRecord")
	fd_ValidatorAccumulatedCommissionRecord_validator_address = md_ValidatorAccumulatedCommissionRecord.Fields().ByName("validator_address")
	fd_ValidatorAccumulatedCommissionRecord_accumulated = md_ValidatorAccumulatedCommissionRecord.Fields().ByName("accumulated")
}

var _ protoreflect.Message = (*fastReflection_ValidatorAccumulatedCommissionRecord)(nil)

type fastReflection_ValidatorAccumulatedCommissionRecord ValidatorAccumulatedCommissionRecord

func (x *ValidatorAccumulatedCommissionRecord) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ValidatorAccumulatedCommissionRecord)(x)
}

func (x *ValidatorAccumulatedCommissionRecord) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ValidatorAccumulatedCommissionRecord_messageType fastReflection_ValidatorAccumulatedCommissionRecord_messageType
var _ protoreflect.MessageType = fastReflection_ValidatorAccumulatedCommissionRecord_messageType{}

type fastReflection_ValidatorAccumulatedCommissionRecord_messageType struct{}

func (x fastReflection_ValidatorAccumulatedCommissionRecord_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ValidatorAccumulatedCommissionRecord)(nil)
}
func (x fastReflection_ValidatorAccumulatedCommissionRecord_messageType) New() protoreflect.Message {
	return new(fastReflection_ValidatorAccumulatedCommissionRecord)
}
func (x fastReflection_ValidatorAccumulatedCommissionRecord_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorAccumulatedCommissionRecord
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ValidatorAccumulatedCommissionRecord) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorAccumulatedCommissionRecord
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ValidatorAccumulatedCommissionRecord) Type() protoreflect.MessageType {
	return _fastReflection_ValidatorAccumulatedCommissionRecord_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ValidatorAccumulatedCommissionRecord) New() protoreflect.Message {
	return new(fastReflection_ValidatorAccumulatedCommissionRecord)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ValidatorAccumulatedCommissionRecord) Interface() protoreflect.ProtoMessage {
	return (*ValidatorAccumulatedCommissionRecord)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ValidatorAccumulatedCommissionRecord) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ValidatorAddress != "" {
		value := protoreflect.ValueOfString(x.ValidatorAddress)
		if !f(fd_ValidatorAccumulatedCommissionRecord_validator_address, value) {
			return
		}
	}
	if x.Accumulated != nil {
		value := protoreflect.ValueOfMessage(x.Accumulated.ProtoReflect())
		if !f(fd_ValidatorAccumulatedCommissionRecord_accumulated, value) {
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
func (x *fastReflection_ValidatorAccumulatedCommissionRecord) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord.validator_address":
		return x.ValidatorAddress != ""
	case "cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord.accumulated":
		return x.Accumulated != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorAccumulatedCommissionRecord) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord.validator_address":
		x.ValidatorAddress = ""
	case "cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord.accumulated":
		x.Accumulated = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ValidatorAccumulatedCommissionRecord) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord.validator_address":
		value := x.ValidatorAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord.accumulated":
		value := x.Accumulated
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ValidatorAccumulatedCommissionRecord) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord.validator_address":
		x.ValidatorAddress = value.Interface().(string)
	case "cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord.accumulated":
		x.Accumulated = value.Message().Interface().(*ValidatorAccumulatedCommission)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ValidatorAccumulatedCommissionRecord) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord.accumulated":
		if x.Accumulated == nil {
			x.Accumulated = new(ValidatorAccumulatedCommission)
		}
		return protoreflect.ValueOfMessage(x.Accumulated.ProtoReflect())
	case "cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord.validator_address":
		panic(fmt.Errorf("field validator_address of message cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ValidatorAccumulatedCommissionRecord) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord.validator_address":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord.accumulated":
		m := new(ValidatorAccumulatedCommission)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ValidatorAccumulatedCommissionRecord) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ValidatorAccumulatedCommissionRecord) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorAccumulatedCommissionRecord) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ValidatorAccumulatedCommissionRecord) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ValidatorAccumulatedCommissionRecord) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ValidatorAccumulatedCommissionRecord)
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
		if x.Accumulated != nil {
			l = options.Size(x.Accumulated)
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
		x := input.Message.Interface().(*ValidatorAccumulatedCommissionRecord)
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
		if x.Accumulated != nil {
			encoded, err := options.Marshal(x.Accumulated)
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
		x := input.Message.Interface().(*ValidatorAccumulatedCommissionRecord)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorAccumulatedCommissionRecord: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorAccumulatedCommissionRecord: illegal tag %d (wire type %d)", fieldNum, wire)
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Accumulated", wireType)
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
				if x.Accumulated == nil {
					x.Accumulated = &ValidatorAccumulatedCommission{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Accumulated); err != nil {
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
	md_ValidatorHistoricalRewardsRecord                   protoreflect.MessageDescriptor
	fd_ValidatorHistoricalRewardsRecord_validator_address protoreflect.FieldDescriptor
	fd_ValidatorHistoricalRewardsRecord_period            protoreflect.FieldDescriptor
	fd_ValidatorHistoricalRewardsRecord_rewards           protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_genesis_proto_init()
	md_ValidatorHistoricalRewardsRecord = File_cosmos_distribution_v1beta1_genesis_proto.Messages().ByName("ValidatorHistoricalRewardsRecord")
	fd_ValidatorHistoricalRewardsRecord_validator_address = md_ValidatorHistoricalRewardsRecord.Fields().ByName("validator_address")
	fd_ValidatorHistoricalRewardsRecord_period = md_ValidatorHistoricalRewardsRecord.Fields().ByName("period")
	fd_ValidatorHistoricalRewardsRecord_rewards = md_ValidatorHistoricalRewardsRecord.Fields().ByName("rewards")
}

var _ protoreflect.Message = (*fastReflection_ValidatorHistoricalRewardsRecord)(nil)

type fastReflection_ValidatorHistoricalRewardsRecord ValidatorHistoricalRewardsRecord

func (x *ValidatorHistoricalRewardsRecord) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ValidatorHistoricalRewardsRecord)(x)
}

func (x *ValidatorHistoricalRewardsRecord) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ValidatorHistoricalRewardsRecord_messageType fastReflection_ValidatorHistoricalRewardsRecord_messageType
var _ protoreflect.MessageType = fastReflection_ValidatorHistoricalRewardsRecord_messageType{}

type fastReflection_ValidatorHistoricalRewardsRecord_messageType struct{}

func (x fastReflection_ValidatorHistoricalRewardsRecord_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ValidatorHistoricalRewardsRecord)(nil)
}
func (x fastReflection_ValidatorHistoricalRewardsRecord_messageType) New() protoreflect.Message {
	return new(fastReflection_ValidatorHistoricalRewardsRecord)
}
func (x fastReflection_ValidatorHistoricalRewardsRecord_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorHistoricalRewardsRecord
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ValidatorHistoricalRewardsRecord) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorHistoricalRewardsRecord
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ValidatorHistoricalRewardsRecord) Type() protoreflect.MessageType {
	return _fastReflection_ValidatorHistoricalRewardsRecord_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ValidatorHistoricalRewardsRecord) New() protoreflect.Message {
	return new(fastReflection_ValidatorHistoricalRewardsRecord)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ValidatorHistoricalRewardsRecord) Interface() protoreflect.ProtoMessage {
	return (*ValidatorHistoricalRewardsRecord)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ValidatorHistoricalRewardsRecord) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ValidatorAddress != "" {
		value := protoreflect.ValueOfString(x.ValidatorAddress)
		if !f(fd_ValidatorHistoricalRewardsRecord_validator_address, value) {
			return
		}
	}
	if x.Period != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Period)
		if !f(fd_ValidatorHistoricalRewardsRecord_period, value) {
			return
		}
	}
	if x.Rewards != nil {
		value := protoreflect.ValueOfMessage(x.Rewards.ProtoReflect())
		if !f(fd_ValidatorHistoricalRewardsRecord_rewards, value) {
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
func (x *fastReflection_ValidatorHistoricalRewardsRecord) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord.validator_address":
		return x.ValidatorAddress != ""
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord.period":
		return x.Period != uint64(0)
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord.rewards":
		return x.Rewards != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorHistoricalRewardsRecord) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord.validator_address":
		x.ValidatorAddress = ""
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord.period":
		x.Period = uint64(0)
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord.rewards":
		x.Rewards = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ValidatorHistoricalRewardsRecord) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord.validator_address":
		value := x.ValidatorAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord.period":
		value := x.Period
		return protoreflect.ValueOfUint64(value)
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord.rewards":
		value := x.Rewards
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ValidatorHistoricalRewardsRecord) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord.validator_address":
		x.ValidatorAddress = value.Interface().(string)
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord.period":
		x.Period = value.Uint()
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord.rewards":
		x.Rewards = value.Message().Interface().(*ValidatorHistoricalRewards)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ValidatorHistoricalRewardsRecord) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord.rewards":
		if x.Rewards == nil {
			x.Rewards = new(ValidatorHistoricalRewards)
		}
		return protoreflect.ValueOfMessage(x.Rewards.ProtoReflect())
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord.validator_address":
		panic(fmt.Errorf("field validator_address of message cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord is not mutable"))
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord.period":
		panic(fmt.Errorf("field period of message cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ValidatorHistoricalRewardsRecord) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord.validator_address":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord.period":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord.rewards":
		m := new(ValidatorHistoricalRewards)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ValidatorHistoricalRewardsRecord) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ValidatorHistoricalRewardsRecord) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorHistoricalRewardsRecord) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ValidatorHistoricalRewardsRecord) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ValidatorHistoricalRewardsRecord) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ValidatorHistoricalRewardsRecord)
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
		if x.Period != 0 {
			n += 1 + runtime.Sov(uint64(x.Period))
		}
		if x.Rewards != nil {
			l = options.Size(x.Rewards)
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
		x := input.Message.Interface().(*ValidatorHistoricalRewardsRecord)
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
		if x.Rewards != nil {
			encoded, err := options.Marshal(x.Rewards)
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
			dAtA[i] = 0x1a
		}
		if x.Period != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Period))
			i--
			dAtA[i] = 0x10
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
		x := input.Message.Interface().(*ValidatorHistoricalRewardsRecord)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorHistoricalRewardsRecord: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorHistoricalRewardsRecord: illegal tag %d (wire type %d)", fieldNum, wire)
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
			case 3:
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
				if x.Rewards == nil {
					x.Rewards = &ValidatorHistoricalRewards{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Rewards); err != nil {
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
	md_ValidatorCurrentRewardsRecord                   protoreflect.MessageDescriptor
	fd_ValidatorCurrentRewardsRecord_validator_address protoreflect.FieldDescriptor
	fd_ValidatorCurrentRewardsRecord_rewards           protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_genesis_proto_init()
	md_ValidatorCurrentRewardsRecord = File_cosmos_distribution_v1beta1_genesis_proto.Messages().ByName("ValidatorCurrentRewardsRecord")
	fd_ValidatorCurrentRewardsRecord_validator_address = md_ValidatorCurrentRewardsRecord.Fields().ByName("validator_address")
	fd_ValidatorCurrentRewardsRecord_rewards = md_ValidatorCurrentRewardsRecord.Fields().ByName("rewards")
}

var _ protoreflect.Message = (*fastReflection_ValidatorCurrentRewardsRecord)(nil)

type fastReflection_ValidatorCurrentRewardsRecord ValidatorCurrentRewardsRecord

func (x *ValidatorCurrentRewardsRecord) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ValidatorCurrentRewardsRecord)(x)
}

func (x *ValidatorCurrentRewardsRecord) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ValidatorCurrentRewardsRecord_messageType fastReflection_ValidatorCurrentRewardsRecord_messageType
var _ protoreflect.MessageType = fastReflection_ValidatorCurrentRewardsRecord_messageType{}

type fastReflection_ValidatorCurrentRewardsRecord_messageType struct{}

func (x fastReflection_ValidatorCurrentRewardsRecord_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ValidatorCurrentRewardsRecord)(nil)
}
func (x fastReflection_ValidatorCurrentRewardsRecord_messageType) New() protoreflect.Message {
	return new(fastReflection_ValidatorCurrentRewardsRecord)
}
func (x fastReflection_ValidatorCurrentRewardsRecord_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorCurrentRewardsRecord
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ValidatorCurrentRewardsRecord) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorCurrentRewardsRecord
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ValidatorCurrentRewardsRecord) Type() protoreflect.MessageType {
	return _fastReflection_ValidatorCurrentRewardsRecord_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ValidatorCurrentRewardsRecord) New() protoreflect.Message {
	return new(fastReflection_ValidatorCurrentRewardsRecord)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ValidatorCurrentRewardsRecord) Interface() protoreflect.ProtoMessage {
	return (*ValidatorCurrentRewardsRecord)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ValidatorCurrentRewardsRecord) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ValidatorAddress != "" {
		value := protoreflect.ValueOfString(x.ValidatorAddress)
		if !f(fd_ValidatorCurrentRewardsRecord_validator_address, value) {
			return
		}
	}
	if x.Rewards != nil {
		value := protoreflect.ValueOfMessage(x.Rewards.ProtoReflect())
		if !f(fd_ValidatorCurrentRewardsRecord_rewards, value) {
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
func (x *fastReflection_ValidatorCurrentRewardsRecord) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord.validator_address":
		return x.ValidatorAddress != ""
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord.rewards":
		return x.Rewards != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorCurrentRewardsRecord) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord.validator_address":
		x.ValidatorAddress = ""
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord.rewards":
		x.Rewards = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ValidatorCurrentRewardsRecord) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord.validator_address":
		value := x.ValidatorAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord.rewards":
		value := x.Rewards
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ValidatorCurrentRewardsRecord) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord.validator_address":
		x.ValidatorAddress = value.Interface().(string)
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord.rewards":
		x.Rewards = value.Message().Interface().(*ValidatorCurrentRewards)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ValidatorCurrentRewardsRecord) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord.rewards":
		if x.Rewards == nil {
			x.Rewards = new(ValidatorCurrentRewards)
		}
		return protoreflect.ValueOfMessage(x.Rewards.ProtoReflect())
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord.validator_address":
		panic(fmt.Errorf("field validator_address of message cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ValidatorCurrentRewardsRecord) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord.validator_address":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord.rewards":
		m := new(ValidatorCurrentRewards)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ValidatorCurrentRewardsRecord) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ValidatorCurrentRewardsRecord) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorCurrentRewardsRecord) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ValidatorCurrentRewardsRecord) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ValidatorCurrentRewardsRecord) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ValidatorCurrentRewardsRecord)
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
		if x.Rewards != nil {
			l = options.Size(x.Rewards)
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
		x := input.Message.Interface().(*ValidatorCurrentRewardsRecord)
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
		if x.Rewards != nil {
			encoded, err := options.Marshal(x.Rewards)
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
		x := input.Message.Interface().(*ValidatorCurrentRewardsRecord)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorCurrentRewardsRecord: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorCurrentRewardsRecord: illegal tag %d (wire type %d)", fieldNum, wire)
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
				if x.Rewards == nil {
					x.Rewards = &ValidatorCurrentRewards{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Rewards); err != nil {
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
	md_DelegatorStartingInfoRecord                   protoreflect.MessageDescriptor
	fd_DelegatorStartingInfoRecord_delegator_address protoreflect.FieldDescriptor
	fd_DelegatorStartingInfoRecord_validator_address protoreflect.FieldDescriptor
	fd_DelegatorStartingInfoRecord_starting_info     protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_genesis_proto_init()
	md_DelegatorStartingInfoRecord = File_cosmos_distribution_v1beta1_genesis_proto.Messages().ByName("DelegatorStartingInfoRecord")
	fd_DelegatorStartingInfoRecord_delegator_address = md_DelegatorStartingInfoRecord.Fields().ByName("delegator_address")
	fd_DelegatorStartingInfoRecord_validator_address = md_DelegatorStartingInfoRecord.Fields().ByName("validator_address")
	fd_DelegatorStartingInfoRecord_starting_info = md_DelegatorStartingInfoRecord.Fields().ByName("starting_info")
}

var _ protoreflect.Message = (*fastReflection_DelegatorStartingInfoRecord)(nil)

type fastReflection_DelegatorStartingInfoRecord DelegatorStartingInfoRecord

func (x *DelegatorStartingInfoRecord) ProtoReflect() protoreflect.Message {
	return (*fastReflection_DelegatorStartingInfoRecord)(x)
}

func (x *DelegatorStartingInfoRecord) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_DelegatorStartingInfoRecord_messageType fastReflection_DelegatorStartingInfoRecord_messageType
var _ protoreflect.MessageType = fastReflection_DelegatorStartingInfoRecord_messageType{}

type fastReflection_DelegatorStartingInfoRecord_messageType struct{}

func (x fastReflection_DelegatorStartingInfoRecord_messageType) Zero() protoreflect.Message {
	return (*fastReflection_DelegatorStartingInfoRecord)(nil)
}
func (x fastReflection_DelegatorStartingInfoRecord_messageType) New() protoreflect.Message {
	return new(fastReflection_DelegatorStartingInfoRecord)
}
func (x fastReflection_DelegatorStartingInfoRecord_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_DelegatorStartingInfoRecord
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_DelegatorStartingInfoRecord) Descriptor() protoreflect.MessageDescriptor {
	return md_DelegatorStartingInfoRecord
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_DelegatorStartingInfoRecord) Type() protoreflect.MessageType {
	return _fastReflection_DelegatorStartingInfoRecord_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_DelegatorStartingInfoRecord) New() protoreflect.Message {
	return new(fastReflection_DelegatorStartingInfoRecord)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_DelegatorStartingInfoRecord) Interface() protoreflect.ProtoMessage {
	return (*DelegatorStartingInfoRecord)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_DelegatorStartingInfoRecord) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.DelegatorAddress != "" {
		value := protoreflect.ValueOfString(x.DelegatorAddress)
		if !f(fd_DelegatorStartingInfoRecord_delegator_address, value) {
			return
		}
	}
	if x.ValidatorAddress != "" {
		value := protoreflect.ValueOfString(x.ValidatorAddress)
		if !f(fd_DelegatorStartingInfoRecord_validator_address, value) {
			return
		}
	}
	if x.StartingInfo != nil {
		value := protoreflect.ValueOfMessage(x.StartingInfo.ProtoReflect())
		if !f(fd_DelegatorStartingInfoRecord_starting_info, value) {
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
func (x *fastReflection_DelegatorStartingInfoRecord) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.DelegatorStartingInfoRecord.delegator_address":
		return x.DelegatorAddress != ""
	case "cosmos.distribution.v1beta1.DelegatorStartingInfoRecord.validator_address":
		return x.ValidatorAddress != ""
	case "cosmos.distribution.v1beta1.DelegatorStartingInfoRecord.starting_info":
		return x.StartingInfo != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegatorStartingInfoRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegatorStartingInfoRecord does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DelegatorStartingInfoRecord) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.DelegatorStartingInfoRecord.delegator_address":
		x.DelegatorAddress = ""
	case "cosmos.distribution.v1beta1.DelegatorStartingInfoRecord.validator_address":
		x.ValidatorAddress = ""
	case "cosmos.distribution.v1beta1.DelegatorStartingInfoRecord.starting_info":
		x.StartingInfo = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegatorStartingInfoRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegatorStartingInfoRecord does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_DelegatorStartingInfoRecord) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.DelegatorStartingInfoRecord.delegator_address":
		value := x.DelegatorAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.DelegatorStartingInfoRecord.validator_address":
		value := x.ValidatorAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.DelegatorStartingInfoRecord.starting_info":
		value := x.StartingInfo
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegatorStartingInfoRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegatorStartingInfoRecord does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_DelegatorStartingInfoRecord) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.DelegatorStartingInfoRecord.delegator_address":
		x.DelegatorAddress = value.Interface().(string)
	case "cosmos.distribution.v1beta1.DelegatorStartingInfoRecord.validator_address":
		x.ValidatorAddress = value.Interface().(string)
	case "cosmos.distribution.v1beta1.DelegatorStartingInfoRecord.starting_info":
		x.StartingInfo = value.Message().Interface().(*DelegatorStartingInfo)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegatorStartingInfoRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegatorStartingInfoRecord does not contain field %s", fd.FullName()))
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
func (x *fastReflection_DelegatorStartingInfoRecord) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.DelegatorStartingInfoRecord.starting_info":
		if x.StartingInfo == nil {
			x.StartingInfo = new(DelegatorStartingInfo)
		}
		return protoreflect.ValueOfMessage(x.StartingInfo.ProtoReflect())
	case "cosmos.distribution.v1beta1.DelegatorStartingInfoRecord.delegator_address":
		panic(fmt.Errorf("field delegator_address of message cosmos.distribution.v1beta1.DelegatorStartingInfoRecord is not mutable"))
	case "cosmos.distribution.v1beta1.DelegatorStartingInfoRecord.validator_address":
		panic(fmt.Errorf("field validator_address of message cosmos.distribution.v1beta1.DelegatorStartingInfoRecord is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegatorStartingInfoRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegatorStartingInfoRecord does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_DelegatorStartingInfoRecord) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.DelegatorStartingInfoRecord.delegator_address":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.DelegatorStartingInfoRecord.validator_address":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.DelegatorStartingInfoRecord.starting_info":
		m := new(DelegatorStartingInfo)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.DelegatorStartingInfoRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.DelegatorStartingInfoRecord does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_DelegatorStartingInfoRecord) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.DelegatorStartingInfoRecord", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_DelegatorStartingInfoRecord) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DelegatorStartingInfoRecord) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_DelegatorStartingInfoRecord) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_DelegatorStartingInfoRecord) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*DelegatorStartingInfoRecord)
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
		l = len(x.DelegatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ValidatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.StartingInfo != nil {
			l = options.Size(x.StartingInfo)
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
		x := input.Message.Interface().(*DelegatorStartingInfoRecord)
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
		if x.StartingInfo != nil {
			encoded, err := options.Marshal(x.StartingInfo)
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
			dAtA[i] = 0x1a
		}
		if len(x.ValidatorAddress) > 0 {
			i -= len(x.ValidatorAddress)
			copy(dAtA[i:], x.ValidatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorAddress)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.DelegatorAddress) > 0 {
			i -= len(x.DelegatorAddress)
			copy(dAtA[i:], x.DelegatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DelegatorAddress)))
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
		x := input.Message.Interface().(*DelegatorStartingInfoRecord)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DelegatorStartingInfoRecord: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DelegatorStartingInfoRecord: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
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
				x.DelegatorAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
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
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field StartingInfo", wireType)
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
				if x.StartingInfo == nil {
					x.StartingInfo = &DelegatorStartingInfo{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.StartingInfo); err != nil {
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
	md_ValidatorSlashEventRecord                       protoreflect.MessageDescriptor
	fd_ValidatorSlashEventRecord_validator_address     protoreflect.FieldDescriptor
	fd_ValidatorSlashEventRecord_height                protoreflect.FieldDescriptor
	fd_ValidatorSlashEventRecord_period                protoreflect.FieldDescriptor
	fd_ValidatorSlashEventRecord_validator_slash_event protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_genesis_proto_init()
	md_ValidatorSlashEventRecord = File_cosmos_distribution_v1beta1_genesis_proto.Messages().ByName("ValidatorSlashEventRecord")
	fd_ValidatorSlashEventRecord_validator_address = md_ValidatorSlashEventRecord.Fields().ByName("validator_address")
	fd_ValidatorSlashEventRecord_height = md_ValidatorSlashEventRecord.Fields().ByName("height")
	fd_ValidatorSlashEventRecord_period = md_ValidatorSlashEventRecord.Fields().ByName("period")
	fd_ValidatorSlashEventRecord_validator_slash_event = md_ValidatorSlashEventRecord.Fields().ByName("validator_slash_event")
}

var _ protoreflect.Message = (*fastReflection_ValidatorSlashEventRecord)(nil)

type fastReflection_ValidatorSlashEventRecord ValidatorSlashEventRecord

func (x *ValidatorSlashEventRecord) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ValidatorSlashEventRecord)(x)
}

func (x *ValidatorSlashEventRecord) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ValidatorSlashEventRecord_messageType fastReflection_ValidatorSlashEventRecord_messageType
var _ protoreflect.MessageType = fastReflection_ValidatorSlashEventRecord_messageType{}

type fastReflection_ValidatorSlashEventRecord_messageType struct{}

func (x fastReflection_ValidatorSlashEventRecord_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ValidatorSlashEventRecord)(nil)
}
func (x fastReflection_ValidatorSlashEventRecord_messageType) New() protoreflect.Message {
	return new(fastReflection_ValidatorSlashEventRecord)
}
func (x fastReflection_ValidatorSlashEventRecord_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorSlashEventRecord
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ValidatorSlashEventRecord) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorSlashEventRecord
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ValidatorSlashEventRecord) Type() protoreflect.MessageType {
	return _fastReflection_ValidatorSlashEventRecord_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ValidatorSlashEventRecord) New() protoreflect.Message {
	return new(fastReflection_ValidatorSlashEventRecord)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ValidatorSlashEventRecord) Interface() protoreflect.ProtoMessage {
	return (*ValidatorSlashEventRecord)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ValidatorSlashEventRecord) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ValidatorAddress != "" {
		value := protoreflect.ValueOfString(x.ValidatorAddress)
		if !f(fd_ValidatorSlashEventRecord_validator_address, value) {
			return
		}
	}
	if x.Height != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Height)
		if !f(fd_ValidatorSlashEventRecord_height, value) {
			return
		}
	}
	if x.Period != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Period)
		if !f(fd_ValidatorSlashEventRecord_period, value) {
			return
		}
	}
	if x.ValidatorSlashEvent != nil {
		value := protoreflect.ValueOfMessage(x.ValidatorSlashEvent.ProtoReflect())
		if !f(fd_ValidatorSlashEventRecord_validator_slash_event, value) {
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
func (x *fastReflection_ValidatorSlashEventRecord) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.validator_address":
		return x.ValidatorAddress != ""
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.height":
		return x.Height != uint64(0)
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.period":
		return x.Period != uint64(0)
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.validator_slash_event":
		return x.ValidatorSlashEvent != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorSlashEventRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorSlashEventRecord does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorSlashEventRecord) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.validator_address":
		x.ValidatorAddress = ""
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.height":
		x.Height = uint64(0)
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.period":
		x.Period = uint64(0)
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.validator_slash_event":
		x.ValidatorSlashEvent = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorSlashEventRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorSlashEventRecord does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ValidatorSlashEventRecord) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.validator_address":
		value := x.ValidatorAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.height":
		value := x.Height
		return protoreflect.ValueOfUint64(value)
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.period":
		value := x.Period
		return protoreflect.ValueOfUint64(value)
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.validator_slash_event":
		value := x.ValidatorSlashEvent
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorSlashEventRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorSlashEventRecord does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ValidatorSlashEventRecord) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.validator_address":
		x.ValidatorAddress = value.Interface().(string)
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.height":
		x.Height = value.Uint()
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.period":
		x.Period = value.Uint()
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.validator_slash_event":
		x.ValidatorSlashEvent = value.Message().Interface().(*ValidatorSlashEvent)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorSlashEventRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorSlashEventRecord does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ValidatorSlashEventRecord) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.validator_slash_event":
		if x.ValidatorSlashEvent == nil {
			x.ValidatorSlashEvent = new(ValidatorSlashEvent)
		}
		return protoreflect.ValueOfMessage(x.ValidatorSlashEvent.ProtoReflect())
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.validator_address":
		panic(fmt.Errorf("field validator_address of message cosmos.distribution.v1beta1.ValidatorSlashEventRecord is not mutable"))
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.height":
		panic(fmt.Errorf("field height of message cosmos.distribution.v1beta1.ValidatorSlashEventRecord is not mutable"))
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.period":
		panic(fmt.Errorf("field period of message cosmos.distribution.v1beta1.ValidatorSlashEventRecord is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorSlashEventRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorSlashEventRecord does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ValidatorSlashEventRecord) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.validator_address":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.height":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.period":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.distribution.v1beta1.ValidatorSlashEventRecord.validator_slash_event":
		m := new(ValidatorSlashEvent)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.ValidatorSlashEventRecord"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.ValidatorSlashEventRecord does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ValidatorSlashEventRecord) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.ValidatorSlashEventRecord", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ValidatorSlashEventRecord) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorSlashEventRecord) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ValidatorSlashEventRecord) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ValidatorSlashEventRecord) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ValidatorSlashEventRecord)
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
		if x.Height != 0 {
			n += 1 + runtime.Sov(uint64(x.Height))
		}
		if x.Period != 0 {
			n += 1 + runtime.Sov(uint64(x.Period))
		}
		if x.ValidatorSlashEvent != nil {
			l = options.Size(x.ValidatorSlashEvent)
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
		x := input.Message.Interface().(*ValidatorSlashEventRecord)
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
		if x.ValidatorSlashEvent != nil {
			encoded, err := options.Marshal(x.ValidatorSlashEvent)
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
		if x.Period != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Period))
			i--
			dAtA[i] = 0x18
		}
		if x.Height != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Height))
			i--
			dAtA[i] = 0x10
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
		x := input.Message.Interface().(*ValidatorSlashEventRecord)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorSlashEventRecord: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorSlashEventRecord: illegal tag %d (wire type %d)", fieldNum, wire)
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
			case 3:
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
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorSlashEvent", wireType)
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
				if x.ValidatorSlashEvent == nil {
					x.ValidatorSlashEvent = &ValidatorSlashEvent{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ValidatorSlashEvent); err != nil {
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

var _ protoreflect.List = (*_GenesisState_3_list)(nil)

type _GenesisState_3_list struct {
	list *[]*DelegatorWithdrawInfo
}

func (x *_GenesisState_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*DelegatorWithdrawInfo)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*DelegatorWithdrawInfo)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_3_list) AppendMutable() protoreflect.Value {
	v := new(DelegatorWithdrawInfo)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_3_list) NewElement() protoreflect.Value {
	v := new(DelegatorWithdrawInfo)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_3_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_GenesisState_5_list)(nil)

type _GenesisState_5_list struct {
	list *[]*ValidatorOutstandingRewardsRecord
}

func (x *_GenesisState_5_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_5_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_5_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ValidatorOutstandingRewardsRecord)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_5_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ValidatorOutstandingRewardsRecord)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_5_list) AppendMutable() protoreflect.Value {
	v := new(ValidatorOutstandingRewardsRecord)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_5_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_5_list) NewElement() protoreflect.Value {
	v := new(ValidatorOutstandingRewardsRecord)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_5_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_GenesisState_6_list)(nil)

type _GenesisState_6_list struct {
	list *[]*ValidatorAccumulatedCommissionRecord
}

func (x *_GenesisState_6_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_6_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_6_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ValidatorAccumulatedCommissionRecord)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_6_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ValidatorAccumulatedCommissionRecord)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_6_list) AppendMutable() protoreflect.Value {
	v := new(ValidatorAccumulatedCommissionRecord)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_6_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_6_list) NewElement() protoreflect.Value {
	v := new(ValidatorAccumulatedCommissionRecord)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_6_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_GenesisState_7_list)(nil)

type _GenesisState_7_list struct {
	list *[]*ValidatorHistoricalRewardsRecord
}

func (x *_GenesisState_7_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_7_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_7_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ValidatorHistoricalRewardsRecord)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_7_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ValidatorHistoricalRewardsRecord)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_7_list) AppendMutable() protoreflect.Value {
	v := new(ValidatorHistoricalRewardsRecord)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_7_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_7_list) NewElement() protoreflect.Value {
	v := new(ValidatorHistoricalRewardsRecord)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_7_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_GenesisState_8_list)(nil)

type _GenesisState_8_list struct {
	list *[]*ValidatorCurrentRewardsRecord
}

func (x *_GenesisState_8_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_8_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_8_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ValidatorCurrentRewardsRecord)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_8_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ValidatorCurrentRewardsRecord)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_8_list) AppendMutable() protoreflect.Value {
	v := new(ValidatorCurrentRewardsRecord)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_8_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_8_list) NewElement() protoreflect.Value {
	v := new(ValidatorCurrentRewardsRecord)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_8_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_GenesisState_9_list)(nil)

type _GenesisState_9_list struct {
	list *[]*DelegatorStartingInfoRecord
}

func (x *_GenesisState_9_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_9_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_9_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*DelegatorStartingInfoRecord)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_9_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*DelegatorStartingInfoRecord)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_9_list) AppendMutable() protoreflect.Value {
	v := new(DelegatorStartingInfoRecord)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_9_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_9_list) NewElement() protoreflect.Value {
	v := new(DelegatorStartingInfoRecord)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_9_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_GenesisState_10_list)(nil)

type _GenesisState_10_list struct {
	list *[]*ValidatorSlashEventRecord
}

func (x *_GenesisState_10_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_10_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_10_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ValidatorSlashEventRecord)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_10_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ValidatorSlashEventRecord)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_10_list) AppendMutable() protoreflect.Value {
	v := new(ValidatorSlashEventRecord)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_10_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_10_list) NewElement() protoreflect.Value {
	v := new(ValidatorSlashEventRecord)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_10_list) IsValid() bool {
	return x.list != nil
}

var (
	md_GenesisState                                   protoreflect.MessageDescriptor
	fd_GenesisState_params                            protoreflect.FieldDescriptor
	fd_GenesisState_fee_pool                          protoreflect.FieldDescriptor
	fd_GenesisState_delegator_withdraw_infos          protoreflect.FieldDescriptor
	fd_GenesisState_previous_proposer                 protoreflect.FieldDescriptor
	fd_GenesisState_outstanding_rewards               protoreflect.FieldDescriptor
	fd_GenesisState_validator_accumulated_commissions protoreflect.FieldDescriptor
	fd_GenesisState_validator_historical_rewards      protoreflect.FieldDescriptor
	fd_GenesisState_validator_current_rewards         protoreflect.FieldDescriptor
	fd_GenesisState_delegator_starting_infos          protoreflect.FieldDescriptor
	fd_GenesisState_validator_slash_events            protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_distribution_v1beta1_genesis_proto_init()
	md_GenesisState = File_cosmos_distribution_v1beta1_genesis_proto.Messages().ByName("GenesisState")
	fd_GenesisState_params = md_GenesisState.Fields().ByName("params")
	fd_GenesisState_fee_pool = md_GenesisState.Fields().ByName("fee_pool")
	fd_GenesisState_delegator_withdraw_infos = md_GenesisState.Fields().ByName("delegator_withdraw_infos")
	fd_GenesisState_previous_proposer = md_GenesisState.Fields().ByName("previous_proposer")
	fd_GenesisState_outstanding_rewards = md_GenesisState.Fields().ByName("outstanding_rewards")
	fd_GenesisState_validator_accumulated_commissions = md_GenesisState.Fields().ByName("validator_accumulated_commissions")
	fd_GenesisState_validator_historical_rewards = md_GenesisState.Fields().ByName("validator_historical_rewards")
	fd_GenesisState_validator_current_rewards = md_GenesisState.Fields().ByName("validator_current_rewards")
	fd_GenesisState_delegator_starting_infos = md_GenesisState.Fields().ByName("delegator_starting_infos")
	fd_GenesisState_validator_slash_events = md_GenesisState.Fields().ByName("validator_slash_events")
}

var _ protoreflect.Message = (*fastReflection_GenesisState)(nil)

type fastReflection_GenesisState GenesisState

func (x *GenesisState) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GenesisState)(x)
}

func (x *GenesisState) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GenesisState_messageType fastReflection_GenesisState_messageType
var _ protoreflect.MessageType = fastReflection_GenesisState_messageType{}

type fastReflection_GenesisState_messageType struct{}

func (x fastReflection_GenesisState_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GenesisState)(nil)
}
func (x fastReflection_GenesisState_messageType) New() protoreflect.Message {
	return new(fastReflection_GenesisState)
}
func (x fastReflection_GenesisState_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GenesisState
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GenesisState) Descriptor() protoreflect.MessageDescriptor {
	return md_GenesisState
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GenesisState) Type() protoreflect.MessageType {
	return _fastReflection_GenesisState_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GenesisState) New() protoreflect.Message {
	return new(fastReflection_GenesisState)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GenesisState) Interface() protoreflect.ProtoMessage {
	return (*GenesisState)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GenesisState) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Params != nil {
		value := protoreflect.ValueOfMessage(x.Params.ProtoReflect())
		if !f(fd_GenesisState_params, value) {
			return
		}
	}
	if x.FeePool != nil {
		value := protoreflect.ValueOfMessage(x.FeePool.ProtoReflect())
		if !f(fd_GenesisState_fee_pool, value) {
			return
		}
	}
	if len(x.DelegatorWithdrawInfos) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_3_list{list: &x.DelegatorWithdrawInfos})
		if !f(fd_GenesisState_delegator_withdraw_infos, value) {
			return
		}
	}
	if x.PreviousProposer != "" {
		value := protoreflect.ValueOfString(x.PreviousProposer)
		if !f(fd_GenesisState_previous_proposer, value) {
			return
		}
	}
	if len(x.OutstandingRewards) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_5_list{list: &x.OutstandingRewards})
		if !f(fd_GenesisState_outstanding_rewards, value) {
			return
		}
	}
	if len(x.ValidatorAccumulatedCommissions) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_6_list{list: &x.ValidatorAccumulatedCommissions})
		if !f(fd_GenesisState_validator_accumulated_commissions, value) {
			return
		}
	}
	if len(x.ValidatorHistoricalRewards) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_7_list{list: &x.ValidatorHistoricalRewards})
		if !f(fd_GenesisState_validator_historical_rewards, value) {
			return
		}
	}
	if len(x.ValidatorCurrentRewards) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_8_list{list: &x.ValidatorCurrentRewards})
		if !f(fd_GenesisState_validator_current_rewards, value) {
			return
		}
	}
	if len(x.DelegatorStartingInfos) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_9_list{list: &x.DelegatorStartingInfos})
		if !f(fd_GenesisState_delegator_starting_infos, value) {
			return
		}
	}
	if len(x.ValidatorSlashEvents) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_10_list{list: &x.ValidatorSlashEvents})
		if !f(fd_GenesisState_validator_slash_events, value) {
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
func (x *fastReflection_GenesisState) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.GenesisState.params":
		return x.Params != nil
	case "cosmos.distribution.v1beta1.GenesisState.fee_pool":
		return x.FeePool != nil
	case "cosmos.distribution.v1beta1.GenesisState.delegator_withdraw_infos":
		return len(x.DelegatorWithdrawInfos) != 0
	case "cosmos.distribution.v1beta1.GenesisState.previous_proposer":
		return x.PreviousProposer != ""
	case "cosmos.distribution.v1beta1.GenesisState.outstanding_rewards":
		return len(x.OutstandingRewards) != 0
	case "cosmos.distribution.v1beta1.GenesisState.validator_accumulated_commissions":
		return len(x.ValidatorAccumulatedCommissions) != 0
	case "cosmos.distribution.v1beta1.GenesisState.validator_historical_rewards":
		return len(x.ValidatorHistoricalRewards) != 0
	case "cosmos.distribution.v1beta1.GenesisState.validator_current_rewards":
		return len(x.ValidatorCurrentRewards) != 0
	case "cosmos.distribution.v1beta1.GenesisState.delegator_starting_infos":
		return len(x.DelegatorStartingInfos) != 0
	case "cosmos.distribution.v1beta1.GenesisState.validator_slash_events":
		return len(x.ValidatorSlashEvents) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.GenesisState"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.GenesisState does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.GenesisState.params":
		x.Params = nil
	case "cosmos.distribution.v1beta1.GenesisState.fee_pool":
		x.FeePool = nil
	case "cosmos.distribution.v1beta1.GenesisState.delegator_withdraw_infos":
		x.DelegatorWithdrawInfos = nil
	case "cosmos.distribution.v1beta1.GenesisState.previous_proposer":
		x.PreviousProposer = ""
	case "cosmos.distribution.v1beta1.GenesisState.outstanding_rewards":
		x.OutstandingRewards = nil
	case "cosmos.distribution.v1beta1.GenesisState.validator_accumulated_commissions":
		x.ValidatorAccumulatedCommissions = nil
	case "cosmos.distribution.v1beta1.GenesisState.validator_historical_rewards":
		x.ValidatorHistoricalRewards = nil
	case "cosmos.distribution.v1beta1.GenesisState.validator_current_rewards":
		x.ValidatorCurrentRewards = nil
	case "cosmos.distribution.v1beta1.GenesisState.delegator_starting_infos":
		x.DelegatorStartingInfos = nil
	case "cosmos.distribution.v1beta1.GenesisState.validator_slash_events":
		x.ValidatorSlashEvents = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.GenesisState"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.GenesisState does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GenesisState) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.distribution.v1beta1.GenesisState.params":
		value := x.Params
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.distribution.v1beta1.GenesisState.fee_pool":
		value := x.FeePool
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.distribution.v1beta1.GenesisState.delegator_withdraw_infos":
		if len(x.DelegatorWithdrawInfos) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_3_list{})
		}
		listValue := &_GenesisState_3_list{list: &x.DelegatorWithdrawInfos}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.distribution.v1beta1.GenesisState.previous_proposer":
		value := x.PreviousProposer
		return protoreflect.ValueOfString(value)
	case "cosmos.distribution.v1beta1.GenesisState.outstanding_rewards":
		if len(x.OutstandingRewards) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_5_list{})
		}
		listValue := &_GenesisState_5_list{list: &x.OutstandingRewards}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.distribution.v1beta1.GenesisState.validator_accumulated_commissions":
		if len(x.ValidatorAccumulatedCommissions) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_6_list{})
		}
		listValue := &_GenesisState_6_list{list: &x.ValidatorAccumulatedCommissions}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.distribution.v1beta1.GenesisState.validator_historical_rewards":
		if len(x.ValidatorHistoricalRewards) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_7_list{})
		}
		listValue := &_GenesisState_7_list{list: &x.ValidatorHistoricalRewards}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.distribution.v1beta1.GenesisState.validator_current_rewards":
		if len(x.ValidatorCurrentRewards) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_8_list{})
		}
		listValue := &_GenesisState_8_list{list: &x.ValidatorCurrentRewards}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.distribution.v1beta1.GenesisState.delegator_starting_infos":
		if len(x.DelegatorStartingInfos) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_9_list{})
		}
		listValue := &_GenesisState_9_list{list: &x.DelegatorStartingInfos}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.distribution.v1beta1.GenesisState.validator_slash_events":
		if len(x.ValidatorSlashEvents) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_10_list{})
		}
		listValue := &_GenesisState_10_list{list: &x.ValidatorSlashEvents}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.GenesisState"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.GenesisState does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_GenesisState) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.GenesisState.params":
		x.Params = value.Message().Interface().(*Params)
	case "cosmos.distribution.v1beta1.GenesisState.fee_pool":
		x.FeePool = value.Message().Interface().(*FeePool)
	case "cosmos.distribution.v1beta1.GenesisState.delegator_withdraw_infos":
		lv := value.List()
		clv := lv.(*_GenesisState_3_list)
		x.DelegatorWithdrawInfos = *clv.list
	case "cosmos.distribution.v1beta1.GenesisState.previous_proposer":
		x.PreviousProposer = value.Interface().(string)
	case "cosmos.distribution.v1beta1.GenesisState.outstanding_rewards":
		lv := value.List()
		clv := lv.(*_GenesisState_5_list)
		x.OutstandingRewards = *clv.list
	case "cosmos.distribution.v1beta1.GenesisState.validator_accumulated_commissions":
		lv := value.List()
		clv := lv.(*_GenesisState_6_list)
		x.ValidatorAccumulatedCommissions = *clv.list
	case "cosmos.distribution.v1beta1.GenesisState.validator_historical_rewards":
		lv := value.List()
		clv := lv.(*_GenesisState_7_list)
		x.ValidatorHistoricalRewards = *clv.list
	case "cosmos.distribution.v1beta1.GenesisState.validator_current_rewards":
		lv := value.List()
		clv := lv.(*_GenesisState_8_list)
		x.ValidatorCurrentRewards = *clv.list
	case "cosmos.distribution.v1beta1.GenesisState.delegator_starting_infos":
		lv := value.List()
		clv := lv.(*_GenesisState_9_list)
		x.DelegatorStartingInfos = *clv.list
	case "cosmos.distribution.v1beta1.GenesisState.validator_slash_events":
		lv := value.List()
		clv := lv.(*_GenesisState_10_list)
		x.ValidatorSlashEvents = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.GenesisState"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.GenesisState does not contain field %s", fd.FullName()))
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
func (x *fastReflection_GenesisState) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.GenesisState.params":
		if x.Params == nil {
			x.Params = new(Params)
		}
		return protoreflect.ValueOfMessage(x.Params.ProtoReflect())
	case "cosmos.distribution.v1beta1.GenesisState.fee_pool":
		if x.FeePool == nil {
			x.FeePool = new(FeePool)
		}
		return protoreflect.ValueOfMessage(x.FeePool.ProtoReflect())
	case "cosmos.distribution.v1beta1.GenesisState.delegator_withdraw_infos":
		if x.DelegatorWithdrawInfos == nil {
			x.DelegatorWithdrawInfos = []*DelegatorWithdrawInfo{}
		}
		value := &_GenesisState_3_list{list: &x.DelegatorWithdrawInfos}
		return protoreflect.ValueOfList(value)
	case "cosmos.distribution.v1beta1.GenesisState.outstanding_rewards":
		if x.OutstandingRewards == nil {
			x.OutstandingRewards = []*ValidatorOutstandingRewardsRecord{}
		}
		value := &_GenesisState_5_list{list: &x.OutstandingRewards}
		return protoreflect.ValueOfList(value)
	case "cosmos.distribution.v1beta1.GenesisState.validator_accumulated_commissions":
		if x.ValidatorAccumulatedCommissions == nil {
			x.ValidatorAccumulatedCommissions = []*ValidatorAccumulatedCommissionRecord{}
		}
		value := &_GenesisState_6_list{list: &x.ValidatorAccumulatedCommissions}
		return protoreflect.ValueOfList(value)
	case "cosmos.distribution.v1beta1.GenesisState.validator_historical_rewards":
		if x.ValidatorHistoricalRewards == nil {
			x.ValidatorHistoricalRewards = []*ValidatorHistoricalRewardsRecord{}
		}
		value := &_GenesisState_7_list{list: &x.ValidatorHistoricalRewards}
		return protoreflect.ValueOfList(value)
	case "cosmos.distribution.v1beta1.GenesisState.validator_current_rewards":
		if x.ValidatorCurrentRewards == nil {
			x.ValidatorCurrentRewards = []*ValidatorCurrentRewardsRecord{}
		}
		value := &_GenesisState_8_list{list: &x.ValidatorCurrentRewards}
		return protoreflect.ValueOfList(value)
	case "cosmos.distribution.v1beta1.GenesisState.delegator_starting_infos":
		if x.DelegatorStartingInfos == nil {
			x.DelegatorStartingInfos = []*DelegatorStartingInfoRecord{}
		}
		value := &_GenesisState_9_list{list: &x.DelegatorStartingInfos}
		return protoreflect.ValueOfList(value)
	case "cosmos.distribution.v1beta1.GenesisState.validator_slash_events":
		if x.ValidatorSlashEvents == nil {
			x.ValidatorSlashEvents = []*ValidatorSlashEventRecord{}
		}
		value := &_GenesisState_10_list{list: &x.ValidatorSlashEvents}
		return protoreflect.ValueOfList(value)
	case "cosmos.distribution.v1beta1.GenesisState.previous_proposer":
		panic(fmt.Errorf("field previous_proposer of message cosmos.distribution.v1beta1.GenesisState is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.GenesisState"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.GenesisState does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GenesisState) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.distribution.v1beta1.GenesisState.params":
		m := new(Params)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.distribution.v1beta1.GenesisState.fee_pool":
		m := new(FeePool)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.distribution.v1beta1.GenesisState.delegator_withdraw_infos":
		list := []*DelegatorWithdrawInfo{}
		return protoreflect.ValueOfList(&_GenesisState_3_list{list: &list})
	case "cosmos.distribution.v1beta1.GenesisState.previous_proposer":
		return protoreflect.ValueOfString("")
	case "cosmos.distribution.v1beta1.GenesisState.outstanding_rewards":
		list := []*ValidatorOutstandingRewardsRecord{}
		return protoreflect.ValueOfList(&_GenesisState_5_list{list: &list})
	case "cosmos.distribution.v1beta1.GenesisState.validator_accumulated_commissions":
		list := []*ValidatorAccumulatedCommissionRecord{}
		return protoreflect.ValueOfList(&_GenesisState_6_list{list: &list})
	case "cosmos.distribution.v1beta1.GenesisState.validator_historical_rewards":
		list := []*ValidatorHistoricalRewardsRecord{}
		return protoreflect.ValueOfList(&_GenesisState_7_list{list: &list})
	case "cosmos.distribution.v1beta1.GenesisState.validator_current_rewards":
		list := []*ValidatorCurrentRewardsRecord{}
		return protoreflect.ValueOfList(&_GenesisState_8_list{list: &list})
	case "cosmos.distribution.v1beta1.GenesisState.delegator_starting_infos":
		list := []*DelegatorStartingInfoRecord{}
		return protoreflect.ValueOfList(&_GenesisState_9_list{list: &list})
	case "cosmos.distribution.v1beta1.GenesisState.validator_slash_events":
		list := []*ValidatorSlashEventRecord{}
		return protoreflect.ValueOfList(&_GenesisState_10_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.distribution.v1beta1.GenesisState"))
		}
		panic(fmt.Errorf("message cosmos.distribution.v1beta1.GenesisState does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GenesisState) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.distribution.v1beta1.GenesisState", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GenesisState) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_GenesisState) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GenesisState) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GenesisState)
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
		if x.Params != nil {
			l = options.Size(x.Params)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.FeePool != nil {
			l = options.Size(x.FeePool)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.DelegatorWithdrawInfos) > 0 {
			for _, e := range x.DelegatorWithdrawInfos {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		l = len(x.PreviousProposer)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.OutstandingRewards) > 0 {
			for _, e := range x.OutstandingRewards {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.ValidatorAccumulatedCommissions) > 0 {
			for _, e := range x.ValidatorAccumulatedCommissions {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.ValidatorHistoricalRewards) > 0 {
			for _, e := range x.ValidatorHistoricalRewards {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.ValidatorCurrentRewards) > 0 {
			for _, e := range x.ValidatorCurrentRewards {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.DelegatorStartingInfos) > 0 {
			for _, e := range x.DelegatorStartingInfos {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
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
		x := input.Message.Interface().(*GenesisState)
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
				dAtA[i] = 0x52
			}
		}
		if len(x.DelegatorStartingInfos) > 0 {
			for iNdEx := len(x.DelegatorStartingInfos) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.DelegatorStartingInfos[iNdEx])
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
				dAtA[i] = 0x4a
			}
		}
		if len(x.ValidatorCurrentRewards) > 0 {
			for iNdEx := len(x.ValidatorCurrentRewards) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.ValidatorCurrentRewards[iNdEx])
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
				dAtA[i] = 0x42
			}
		}
		if len(x.ValidatorHistoricalRewards) > 0 {
			for iNdEx := len(x.ValidatorHistoricalRewards) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.ValidatorHistoricalRewards[iNdEx])
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
				dAtA[i] = 0x3a
			}
		}
		if len(x.ValidatorAccumulatedCommissions) > 0 {
			for iNdEx := len(x.ValidatorAccumulatedCommissions) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.ValidatorAccumulatedCommissions[iNdEx])
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
				dAtA[i] = 0x32
			}
		}
		if len(x.OutstandingRewards) > 0 {
			for iNdEx := len(x.OutstandingRewards) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.OutstandingRewards[iNdEx])
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
				dAtA[i] = 0x2a
			}
		}
		if len(x.PreviousProposer) > 0 {
			i -= len(x.PreviousProposer)
			copy(dAtA[i:], x.PreviousProposer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.PreviousProposer)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.DelegatorWithdrawInfos) > 0 {
			for iNdEx := len(x.DelegatorWithdrawInfos) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.DelegatorWithdrawInfos[iNdEx])
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
				dAtA[i] = 0x1a
			}
		}
		if x.FeePool != nil {
			encoded, err := options.Marshal(x.FeePool)
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
		if x.Params != nil {
			encoded, err := options.Marshal(x.Params)
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
		x := input.Message.Interface().(*GenesisState)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
				if x.Params == nil {
					x.Params = &Params{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Params); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field FeePool", wireType)
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
				if x.FeePool == nil {
					x.FeePool = &FeePool{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.FeePool); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatorWithdrawInfos", wireType)
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
				x.DelegatorWithdrawInfos = append(x.DelegatorWithdrawInfos, &DelegatorWithdrawInfo{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.DelegatorWithdrawInfos[len(x.DelegatorWithdrawInfos)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PreviousProposer", wireType)
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
				x.PreviousProposer = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field OutstandingRewards", wireType)
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
				x.OutstandingRewards = append(x.OutstandingRewards, &ValidatorOutstandingRewardsRecord{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.OutstandingRewards[len(x.OutstandingRewards)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorAccumulatedCommissions", wireType)
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
				x.ValidatorAccumulatedCommissions = append(x.ValidatorAccumulatedCommissions, &ValidatorAccumulatedCommissionRecord{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ValidatorAccumulatedCommissions[len(x.ValidatorAccumulatedCommissions)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorHistoricalRewards", wireType)
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
				x.ValidatorHistoricalRewards = append(x.ValidatorHistoricalRewards, &ValidatorHistoricalRewardsRecord{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ValidatorHistoricalRewards[len(x.ValidatorHistoricalRewards)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 8:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorCurrentRewards", wireType)
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
				x.ValidatorCurrentRewards = append(x.ValidatorCurrentRewards, &ValidatorCurrentRewardsRecord{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ValidatorCurrentRewards[len(x.ValidatorCurrentRewards)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 9:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatorStartingInfos", wireType)
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
				x.DelegatorStartingInfos = append(x.DelegatorStartingInfos, &DelegatorStartingInfoRecord{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.DelegatorStartingInfos[len(x.DelegatorStartingInfos)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 10:
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
				x.ValidatorSlashEvents = append(x.ValidatorSlashEvents, &ValidatorSlashEventRecord{})
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.19.1
// source: cosmos/distribution/v1beta1/genesis.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// DelegatorWithdrawInfo is the address for where distributions rewards are
// withdrawn to by default this struct is only used at genesis to feed in
// default withdraw addresses.
type DelegatorWithdrawInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delegator_address is the address of the delegator.
	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
	// withdraw_address is the address to withdraw the delegation rewards to.
	WithdrawAddress string `protobuf:"bytes,2,opt,name=withdraw_address,json=withdrawAddress,proto3" json:"withdraw_address,omitempty"`
}

func (x *DelegatorWithdrawInfo) Reset() {
	*x = DelegatorWithdrawInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelegatorWithdrawInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelegatorWithdrawInfo) ProtoMessage() {}

// Deprecated: Use DelegatorWithdrawInfo.ProtoReflect.Descriptor instead.
func (*DelegatorWithdrawInfo) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_genesis_proto_rawDescGZIP(), []int{0}
}

func (x *DelegatorWithdrawInfo) GetDelegatorAddress() string {
	if x != nil {
		return x.DelegatorAddress
	}
	return ""
}

func (x *DelegatorWithdrawInfo) GetWithdrawAddress() string {
	if x != nil {
		return x.WithdrawAddress
	}
	return ""
}

// ValidatorOutstandingRewardsRecord is used for import/export via genesis json.
type ValidatorOutstandingRewardsRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// validator_address is the address of the validator.
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// outstanding_rewards represents the outstanding rewards of a validator.
	OutstandingRewards []*v1beta1.DecCoin `protobuf:"bytes,2,rep,name=outstanding_rewards,json=outstandingRewards,proto3" json:"outstanding_rewards,omitempty"`
}

func (x *ValidatorOutstandingRewardsRecord) Reset() {
	*x = ValidatorOutstandingRewardsRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorOutstandingRewardsRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorOutstandingRewardsRecord) ProtoMessage() {}

// Deprecated: Use ValidatorOutstandingRewardsRecord.ProtoReflect.Descriptor instead.
func (*ValidatorOutstandingRewardsRecord) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_genesis_proto_rawDescGZIP(), []int{1}
}

func (x *ValidatorOutstandingRewardsRecord) GetValidatorAddress() string {
	if x != nil {
		return x.ValidatorAddress
	}
	return ""
}

func (x *ValidatorOutstandingRewardsRecord) GetOutstandingRewards() []*v1beta1.DecCoin {
	if x != nil {
		return x.OutstandingRewards
	}
	return nil
}

// ValidatorAccumulatedCommissionRecord is used for import / export via genesis
// json.
type ValidatorAccumulatedCommissionRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// validator_address is the address of the validator.
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// accumulated is the accumulated commission of a validator.
	Accumulated *ValidatorAccumulatedCommission `protobuf:"bytes,2,opt,name=accumulated,proto3" json:"accumulated,omitempty"`
}

func (x *ValidatorAccumulatedCommissionRecord) Reset() {
	*x = ValidatorAccumulatedCommissionRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorAccumulatedCommissionRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorAccumulatedCommissionRecord) ProtoMessage() {}

// Deprecated: Use ValidatorAccumulatedCommissionRecord.ProtoReflect.Descriptor instead.
func (*ValidatorAccumulatedCommissionRecord) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_genesis_proto_rawDescGZIP(), []int{2}
}

func (x *ValidatorAccumulatedCommissionRecord) GetValidatorAddress() string {
	if x != nil {
		return x.ValidatorAddress
	}
	return ""
}

func (x *ValidatorAccumulatedCommissionRecord) GetAccumulated() *ValidatorAccumulatedCommission {
	if x != nil {
		return x.Accumulated
	}
	return nil
}

// ValidatorHistoricalRewardsRecord is used for import / export via genesis
// json.
type ValidatorHistoricalRewardsRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// validator_address is the address of the validator.
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// period defines the period the historical rewards apply to.
	Period uint64 `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
	// rewards defines the historical rewards of a validator.
	Rewards *ValidatorHistoricalRewards `protobuf:"bytes,3,opt,name=rewards,proto3" json:"rewards,omitempty"`
}

func (x *ValidatorHistoricalRewardsRecord) Reset() {
	*x = ValidatorHistoricalRewardsRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorHistoricalRewardsRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorHistoricalRewardsRecord) ProtoMessage() {}

// Deprecated: Use ValidatorHistoricalRewardsRecord.ProtoReflect.Descriptor instead.
func (*ValidatorHistoricalRewardsRecord) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_genesis_proto_rawDescGZIP(), []int{3}
}

func (x *ValidatorHistoricalRewardsRecord) GetValidatorAddress() string {
	if x != nil {
		return x.ValidatorAddress
	}
	return ""
}

func (x *ValidatorHistoricalRewardsRecord) GetPeriod() uint64 {
	if x != nil {
		return x.Period
	}
	return 0
}

func (x *ValidatorHistoricalRewardsRecord) GetRewards() *ValidatorHistoricalRewards {
	if x != nil {
		return x.Rewards
	}
	return nil
}

// ValidatorCurrentRewardsRecord is used for import / export via genesis json.
type ValidatorCurrentRewardsRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// validator_address is the address of the validator.
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// rewards defines the current rewards of a validator.
	Rewards *ValidatorCurrentRewards `protobuf:"bytes,2,opt,name=rewards,proto3" json:"rewards,omitempty"`
}

func (x *ValidatorCurrentRewardsRecord) Reset() {
	*x = ValidatorCurrentRewardsRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorCurrentRewardsRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorCurrentRewardsRecord) ProtoMessage() {}

// Deprecated: Use ValidatorCurrentRewardsRecord.ProtoReflect.Descriptor instead.
func (*ValidatorCurrentRewardsRecord) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_genesis_proto_rawDescGZIP(), []int{4}
}

func (x *ValidatorCurrentRewardsRecord) GetValidatorAddress() string {
	if x != nil {
		return x.ValidatorAddress
	}
	return ""
}

func (x *ValidatorCurrentRewardsRecord) GetRewards() *ValidatorCurrentRewards {
	if x != nil {
		return x.Rewards
	}
	return nil
}

// DelegatorStartingInfoRecord used for import / export via genesis json.
type DelegatorStartingInfoRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delegator_address is the address of the delegator.
	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
	// validator_address is the address of the validator.
	ValidatorAddress string `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// starting_info defines the starting info of a delegator.
	StartingInfo *DelegatorStartingInfo `protobuf:"bytes,3,opt,name=starting_info,json=startingInfo,proto3" json:"starting_info,omitempty"`
}

func (x *DelegatorStartingInfoRecord) Reset() {
	*x = DelegatorStartingInfoRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelegatorStartingInfoRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelegatorStartingInfoRecord) ProtoMessage() {}

// Deprecated: Use DelegatorStartingInfoRecord.ProtoReflect.Descriptor instead.
func (*DelegatorStartingInfoRecord) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_genesis_proto_rawDescGZIP(), []int{5}
}

func (x *DelegatorStartingInfoRecord) GetDelegatorAddress() string {
	if x != nil {
		return x.DelegatorAddress
	}
	return ""
}

func (x *DelegatorStartingInfoRecord) GetValidatorAddress() string {
	if x != nil {
		return x.ValidatorAddress
	}
	return ""
}

func (x *DelegatorStartingInfoRecord) GetStartingInfo() *DelegatorStartingInfo {
	if x != nil {
		return x.StartingInfo
	}
	return nil
}

// ValidatorSlashEventRecord is used for import / export via genesis json.
type ValidatorSlashEventRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// validator_address is the address of the validator.
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// height defines the block height at which the slash event occurred.
	Height uint64 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	// period is the period of the slash event.
	Period uint64 `protobuf:"varint,3,opt,name=period,proto3" json:"period,omitempty"`
	// validator_slash_event describes the slash event.
	ValidatorSlashEvent *ValidatorSlashEvent `protobuf:"bytes,4,opt,name=validator_slash_event,json=validatorSlashEvent,proto3" json:"validator_slash_event,omitempty"`
}

func (x *ValidatorSlashEventRecord) Reset() {
	*x = ValidatorSlashEventRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorSlashEventRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorSlashEventRecord) ProtoMessage() {}

// Deprecated: Use ValidatorSlashEventRecord.ProtoReflect.Descriptor instead.
func (*ValidatorSlashEventRecord) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_genesis_proto_rawDescGZIP(), []int{6}
}

func (x *ValidatorSlashEventRecord) GetValidatorAddress() string {
	if x != nil {
		return x.ValidatorAddress
	}
	return ""
}

func (x *ValidatorSlashEventRecord) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *ValidatorSlashEventRecord) GetPeriod() uint64 {
	if x != nil {
		return x.Period
	}
	return 0
}

func (x *ValidatorSlashEventRecord) GetValidatorSlashEvent() *ValidatorSlashEvent {
	if x != nil {
		return x.ValidatorSlashEvent
	}
	return nil
}

// GenesisState defines the distribution module's genesis state.
type GenesisState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// params defines all the parameters of the module.
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	// fee_pool defines the fee pool at genesis.
	FeePool *FeePool `protobuf:"bytes,2,opt,name=fee_pool,json=feePool,proto3" json:"fee_pool,omitempty"`
	// fee_pool defines the delegator withdraw infos at genesis.
	DelegatorWithdrawInfos []*DelegatorWithdrawInfo `protobuf:"bytes,3,rep,name=delegator_withdraw_infos,json=delegatorWithdrawInfos,proto3" json:"delegator_withdraw_infos,omitempty"`
	// fee_pool defines the previous proposer at genesis.
	PreviousProposer string `protobuf:"bytes,4,opt,name=previous_proposer,json=previousProposer,proto3" json:"previous_proposer,omitempty"`
	// fee_pool defines the outstanding rewards of all validators at genesis.
	OutstandingRewards []*ValidatorOutstandingRewardsRecord `protobuf:"bytes,5,rep,name=outstanding_rewards,json=outstandingRewards,proto3" json:"outstanding_rewards,omitempty"`
	// fee_pool defines the accumulated commisions of all validators at genesis.
	ValidatorAccumulatedCommissions []*ValidatorAccumulatedCommissionRecord `protobuf:"bytes,6,rep,name=validator_accumulated_commissions,json=validatorAccumulatedCommissions,proto3" json:"validator_accumulated_commissions,omitempty"`
	// fee_pool defines the historical rewards of all validators at genesis.
	ValidatorHistoricalRewards []*ValidatorHistoricalRewardsRecord `protobuf:"bytes,7,rep,name=validator_historical_rewards,json=validatorHistoricalRewards,proto3" json:"validator_historical_rewards,omitempty"`
	// fee_pool defines the current rewards of all validators at genesis.
	ValidatorCurrentRewards []*ValidatorCurrentRewardsRecord `protobuf:"bytes,8,rep,name=validator_current_rewards,json=validatorCurrentRewards,proto3" json:"validator_current_rewards,omitempty"`
	// fee_pool defines the delegator starting infos at genesis.
	DelegatorStartingInfos []*DelegatorStartingInfoRecord `protobuf:"bytes,9,rep,name=delegator_starting_infos,json=delegatorStartingInfos,proto3" json:"delegator_starting_infos,omitempty"`
	// fee_pool defines the validator slash events at genesis.
	ValidatorSlashEvents []*ValidatorSlashEventRecord `protobuf:"bytes,10,rep,name=validator_slash_events,json=validatorSlashEvents,proto3" json:"validator_slash_events,omitempty"`
}

func (x *GenesisState) Reset() {
	*x = GenesisState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState) ProtoMessage() {}

// Deprecated: Use GenesisState.ProtoReflect.Descriptor instead.
func (*GenesisState) Descriptor() ([]byte, []int) {
	return file_cosmos_distribution_v1beta1_genesis_proto_rawDescGZIP(), []int{7}
}

func (x *GenesisState) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *GenesisState) GetFeePool() *FeePool {
	if x != nil {
		return x.FeePool
	}
	return nil
}

func (x *GenesisState) GetDelegatorWithdrawInfos() []*DelegatorWithdrawInfo {
	if x != nil {
		return x.DelegatorWithdrawInfos
	}
	return nil
}

func (x *GenesisState) GetPreviousProposer() string {
	if x != nil {
		return x.PreviousProposer
	}
	return ""
}

func (x *GenesisState) GetOutstandingRewards() []*ValidatorOutstandingRewardsRecord {
	if x != nil {
		return x.OutstandingRewards
	}
	return nil
}

func (x *GenesisState) GetValidatorAccumulatedCommissions() []*ValidatorAccumulatedCommissionRecord {
	if x != nil {
		return x.ValidatorAccumulatedCommissions
	}
	return nil
}

func (x *GenesisState) GetValidatorHistoricalRewards() []*ValidatorHistoricalRewardsRecord {
	if x != nil {
		return x.ValidatorHistoricalRewards
	}
	return nil
}

func (x *GenesisState) GetValidatorCurrentRewards() []*ValidatorCurrentRewardsRecord {
	if x != nil {
		return x.ValidatorCurrentRewards
	}
	return nil
}

func (x *GenesisState) GetDelegatorStartingInfos() []*DelegatorStartingInfoRecord {
	if x != nil {
		return x.DelegatorStartingInfos
	}
	return nil
}

func (x *GenesisState) GetValidatorSlashEvents() []*ValidatorSlashEventRecord {
	if x != nil {
		return x.ValidatorSlashEvents
	}
	return nil
}

var File_cosmos_distribution_v1beta1_genesis_proto protoreflect.FileDescriptor

var file_cosmos_distribution_v1beta1_genesis_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x45, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18,
	0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x43, 0x0a, 0x10, 0x77, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0f,
	0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3a,
	0x08, 0x88, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0xf9, 0x01, 0x0a, 0x21, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12,
	0x45, 0x0a, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x82, 0x01, 0x0a, 0x13, 0x6f, 0x75, 0x74, 0x73, 0x74,
	0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x43, 0x6f,
	0x69, 0x6e, 0x42, 0x33, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44,
	0x65, 0x63, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x12, 0x6f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x3a, 0x08, 0x88, 0xa0, 0x1f,
	0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0xdc, 0x01, 0x0a, 0x24, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x45,
	0x0a, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x63, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00,
	0xe8, 0xa0, 0x1f, 0x00, 0x22, 0xe4, 0x01, 0x0a, 0x20, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x45, 0x0a, 0x11, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x10,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x57, 0x0a, 0x07, 0x72, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x73, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x73, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0xc6, 0x01, 0x0a, 0x1d,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x45, 0x0a,
	0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x54, 0x0a, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x42, 0x04, 0xc8, 0xde, 0x1f,
	0x00, 0x52, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00,
	0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x94, 0x02, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x6f, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x45, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x45, 0x0a, 0x11, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x5d, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f,
	0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x04, 0xc8,
	0xde, 0x1f, 0x00, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x88, 0x02, 0x0a, 0x19,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x45, 0x0a, 0x11, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x10,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x12, 0x6a, 0x0a, 0x15, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x6c,
	0x61, 0x73, 0x68, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x30, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x13, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x3a, 0x08, 0x88, 0xa0,
	0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0xde, 0x08, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x73,
	0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x04, 0xc8, 0xde,
	0x1f, 0x00, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x45, 0x0a, 0x08, 0x66, 0x65,
	0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x65, 0x65, 0x50, 0x6f,
	0x6f, 0x6c, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x07, 0x66, 0x65, 0x65, 0x50, 0x6f, 0x6f,
	0x6c, 0x12, 0x72, 0x0a, 0x18, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x77,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x16, 0x64,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x45, 0x0a, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75,
	0x73, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x6f, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x12, 0x75, 0x0a, 0x13,
	0x6f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x4f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52,
	0x12, 0x6f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x73, 0x12, 0x93, 0x01, 0x0a, 0x21, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x61, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x41, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x1f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x85, 0x01, 0x0a, 0x1c, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63,
	0x61, 0x6c, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63,
	0x61, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42,
	0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x1a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x73, 0x12, 0x7c, 0x0a, 0x19, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12,
	0x78, 0x0a, 0x18, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x04, 0xc8, 0xde, 0x1f,
	0x00, 0x52, 0x16, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x72, 0x0a, 0x16, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x14, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x08, 0x88,
	0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x42, 0x93, 0x02, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0c, 0x47, 0x65, 0x6e,
	0x65, 0x73, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74,
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
	file_cosmos_distribution_v1beta1_genesis_proto_rawDescOnce sync.Once
	file_cosmos_distribution_v1beta1_genesis_proto_rawDescData = file_cosmos_distribution_v1beta1_genesis_proto_rawDesc
)

func file_cosmos_distribution_v1beta1_genesis_proto_rawDescGZIP() []byte {
	file_cosmos_distribution_v1beta1_genesis_proto_rawDescOnce.Do(func() {
		file_cosmos_distribution_v1beta1_genesis_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_distribution_v1beta1_genesis_proto_rawDescData)
	})
	return file_cosmos_distribution_v1beta1_genesis_proto_rawDescData
}

var file_cosmos_distribution_v1beta1_genesis_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cosmos_distribution_v1beta1_genesis_proto_goTypes = []interface{}{
	(*DelegatorWithdrawInfo)(nil),                // 0: cosmos.distribution.v1beta1.DelegatorWithdrawInfo
	(*ValidatorOutstandingRewardsRecord)(nil),    // 1: cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord
	(*ValidatorAccumulatedCommissionRecord)(nil), // 2: cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord
	(*ValidatorHistoricalRewardsRecord)(nil),     // 3: cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord
	(*ValidatorCurrentRewardsRecord)(nil),        // 4: cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord
	(*DelegatorStartingInfoRecord)(nil),          // 5: cosmos.distribution.v1beta1.DelegatorStartingInfoRecord
	(*ValidatorSlashEventRecord)(nil),            // 6: cosmos.distribution.v1beta1.ValidatorSlashEventRecord
	(*GenesisState)(nil),                         // 7: cosmos.distribution.v1beta1.GenesisState
	(*v1beta1.DecCoin)(nil),                      // 8: cosmos.base.v1beta1.DecCoin
	(*ValidatorAccumulatedCommission)(nil),       // 9: cosmos.distribution.v1beta1.ValidatorAccumulatedCommission
	(*ValidatorHistoricalRewards)(nil),           // 10: cosmos.distribution.v1beta1.ValidatorHistoricalRewards
	(*ValidatorCurrentRewards)(nil),              // 11: cosmos.distribution.v1beta1.ValidatorCurrentRewards
	(*DelegatorStartingInfo)(nil),                // 12: cosmos.distribution.v1beta1.DelegatorStartingInfo
	(*ValidatorSlashEvent)(nil),                  // 13: cosmos.distribution.v1beta1.ValidatorSlashEvent
	(*Params)(nil),                               // 14: cosmos.distribution.v1beta1.Params
	(*FeePool)(nil),                              // 15: cosmos.distribution.v1beta1.FeePool
}
var file_cosmos_distribution_v1beta1_genesis_proto_depIdxs = []int32{
	8,  // 0: cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord.outstanding_rewards:type_name -> cosmos.base.v1beta1.DecCoin
	9,  // 1: cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord.accumulated:type_name -> cosmos.distribution.v1beta1.ValidatorAccumulatedCommission
	10, // 2: cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord.rewards:type_name -> cosmos.distribution.v1beta1.ValidatorHistoricalRewards
	11, // 3: cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord.rewards:type_name -> cosmos.distribution.v1beta1.ValidatorCurrentRewards
	12, // 4: cosmos.distribution.v1beta1.DelegatorStartingInfoRecord.starting_info:type_name -> cosmos.distribution.v1beta1.DelegatorStartingInfo
	13, // 5: cosmos.distribution.v1beta1.ValidatorSlashEventRecord.validator_slash_event:type_name -> cosmos.distribution.v1beta1.ValidatorSlashEvent
	14, // 6: cosmos.distribution.v1beta1.GenesisState.params:type_name -> cosmos.distribution.v1beta1.Params
	15, // 7: cosmos.distribution.v1beta1.GenesisState.fee_pool:type_name -> cosmos.distribution.v1beta1.FeePool
	0,  // 8: cosmos.distribution.v1beta1.GenesisState.delegator_withdraw_infos:type_name -> cosmos.distribution.v1beta1.DelegatorWithdrawInfo
	1,  // 9: cosmos.distribution.v1beta1.GenesisState.outstanding_rewards:type_name -> cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord
	2,  // 10: cosmos.distribution.v1beta1.GenesisState.validator_accumulated_commissions:type_name -> cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord
	3,  // 11: cosmos.distribution.v1beta1.GenesisState.validator_historical_rewards:type_name -> cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord
	4,  // 12: cosmos.distribution.v1beta1.GenesisState.validator_current_rewards:type_name -> cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord
	5,  // 13: cosmos.distribution.v1beta1.GenesisState.delegator_starting_infos:type_name -> cosmos.distribution.v1beta1.DelegatorStartingInfoRecord
	6,  // 14: cosmos.distribution.v1beta1.GenesisState.validator_slash_events:type_name -> cosmos.distribution.v1beta1.ValidatorSlashEventRecord
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_cosmos_distribution_v1beta1_genesis_proto_init() }
func file_cosmos_distribution_v1beta1_genesis_proto_init() {
	if File_cosmos_distribution_v1beta1_genesis_proto != nil {
		return
	}
	file_cosmos_distribution_v1beta1_distribution_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelegatorWithdrawInfo); i {
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
		file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorOutstandingRewardsRecord); i {
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
		file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorAccumulatedCommissionRecord); i {
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
		file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorHistoricalRewardsRecord); i {
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
		file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorCurrentRewardsRecord); i {
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
		file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelegatorStartingInfoRecord); i {
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
		file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorSlashEventRecord); i {
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
		file_cosmos_distribution_v1beta1_genesis_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState); i {
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
			RawDescriptor: file_cosmos_distribution_v1beta1_genesis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_distribution_v1beta1_genesis_proto_goTypes,
		DependencyIndexes: file_cosmos_distribution_v1beta1_genesis_proto_depIdxs,
		MessageInfos:      file_cosmos_distribution_v1beta1_genesis_proto_msgTypes,
	}.Build()
	File_cosmos_distribution_v1beta1_genesis_proto = out.File
	file_cosmos_distribution_v1beta1_genesis_proto_rawDesc = nil
	file_cosmos_distribution_v1beta1_genesis_proto_goTypes = nil
	file_cosmos_distribution_v1beta1_genesis_proto_depIdxs = nil
}
