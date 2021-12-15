package authv1beta1

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
	anypb "google.golang.org/protobuf/types/known/anypb"
)

var (
	md_BaseAccount                protoreflect.MessageDescriptor
	fd_BaseAccount_address        protoreflect.FieldDescriptor
	fd_BaseAccount_pub_key        protoreflect.FieldDescriptor
	fd_BaseAccount_account_number protoreflect.FieldDescriptor
	fd_BaseAccount_sequence       protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_auth_v1beta1_auth_proto_init()
	md_BaseAccount = File_cosmos_auth_v1beta1_auth_proto.Messages().ByName("BaseAccount")
	fd_BaseAccount_address = md_BaseAccount.Fields().ByName("address")
	fd_BaseAccount_pub_key = md_BaseAccount.Fields().ByName("pub_key")
	fd_BaseAccount_account_number = md_BaseAccount.Fields().ByName("account_number")
	fd_BaseAccount_sequence = md_BaseAccount.Fields().ByName("sequence")
}

var _ protoreflect.Message = (*fastReflection_BaseAccount)(nil)

type fastReflection_BaseAccount BaseAccount

func (x *BaseAccount) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BaseAccount)(x)
}

func (x *BaseAccount) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_auth_v1beta1_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BaseAccount_messageType fastReflection_BaseAccount_messageType
var _ protoreflect.MessageType = fastReflection_BaseAccount_messageType{}

type fastReflection_BaseAccount_messageType struct{}

func (x fastReflection_BaseAccount_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BaseAccount)(nil)
}
func (x fastReflection_BaseAccount_messageType) New() protoreflect.Message {
	return new(fastReflection_BaseAccount)
}
func (x fastReflection_BaseAccount_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BaseAccount
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BaseAccount) Descriptor() protoreflect.MessageDescriptor {
	return md_BaseAccount
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BaseAccount) Type() protoreflect.MessageType {
	return _fastReflection_BaseAccount_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BaseAccount) New() protoreflect.Message {
	return new(fastReflection_BaseAccount)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BaseAccount) Interface() protoreflect.ProtoMessage {
	return (*BaseAccount)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BaseAccount) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_BaseAccount_address, value) {
			return
		}
	}
	if x.PubKey != nil {
		value := protoreflect.ValueOfMessage(x.PubKey.ProtoReflect())
		if !f(fd_BaseAccount_pub_key, value) {
			return
		}
	}
	if x.AccountNumber != uint64(0) {
		value := protoreflect.ValueOfUint64(x.AccountNumber)
		if !f(fd_BaseAccount_account_number, value) {
			return
		}
	}
	if x.Sequence != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Sequence)
		if !f(fd_BaseAccount_sequence, value) {
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
func (x *fastReflection_BaseAccount) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.auth.v1beta1.BaseAccount.address":
		return x.Address != ""
	case "cosmos.auth.v1beta1.BaseAccount.pub_key":
		return x.PubKey != nil
	case "cosmos.auth.v1beta1.BaseAccount.account_number":
		return x.AccountNumber != uint64(0)
	case "cosmos.auth.v1beta1.BaseAccount.sequence":
		return x.Sequence != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.auth.v1beta1.BaseAccount"))
		}
		panic(fmt.Errorf("message cosmos.auth.v1beta1.BaseAccount does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BaseAccount) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.auth.v1beta1.BaseAccount.address":
		x.Address = ""
	case "cosmos.auth.v1beta1.BaseAccount.pub_key":
		x.PubKey = nil
	case "cosmos.auth.v1beta1.BaseAccount.account_number":
		x.AccountNumber = uint64(0)
	case "cosmos.auth.v1beta1.BaseAccount.sequence":
		x.Sequence = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.auth.v1beta1.BaseAccount"))
		}
		panic(fmt.Errorf("message cosmos.auth.v1beta1.BaseAccount does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BaseAccount) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.auth.v1beta1.BaseAccount.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "cosmos.auth.v1beta1.BaseAccount.pub_key":
		value := x.PubKey
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.auth.v1beta1.BaseAccount.account_number":
		value := x.AccountNumber
		return protoreflect.ValueOfUint64(value)
	case "cosmos.auth.v1beta1.BaseAccount.sequence":
		value := x.Sequence
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.auth.v1beta1.BaseAccount"))
		}
		panic(fmt.Errorf("message cosmos.auth.v1beta1.BaseAccount does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_BaseAccount) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.auth.v1beta1.BaseAccount.address":
		x.Address = value.Interface().(string)
	case "cosmos.auth.v1beta1.BaseAccount.pub_key":
		x.PubKey = value.Message().Interface().(*anypb.Any)
	case "cosmos.auth.v1beta1.BaseAccount.account_number":
		x.AccountNumber = value.Uint()
	case "cosmos.auth.v1beta1.BaseAccount.sequence":
		x.Sequence = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.auth.v1beta1.BaseAccount"))
		}
		panic(fmt.Errorf("message cosmos.auth.v1beta1.BaseAccount does not contain field %s", fd.FullName()))
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
func (x *fastReflection_BaseAccount) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.auth.v1beta1.BaseAccount.pub_key":
		if x.PubKey == nil {
			x.PubKey = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.PubKey.ProtoReflect())
	case "cosmos.auth.v1beta1.BaseAccount.address":
		panic(fmt.Errorf("field address of message cosmos.auth.v1beta1.BaseAccount is not mutable"))
	case "cosmos.auth.v1beta1.BaseAccount.account_number":
		panic(fmt.Errorf("field account_number of message cosmos.auth.v1beta1.BaseAccount is not mutable"))
	case "cosmos.auth.v1beta1.BaseAccount.sequence":
		panic(fmt.Errorf("field sequence of message cosmos.auth.v1beta1.BaseAccount is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.auth.v1beta1.BaseAccount"))
		}
		panic(fmt.Errorf("message cosmos.auth.v1beta1.BaseAccount does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BaseAccount) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.auth.v1beta1.BaseAccount.address":
		return protoreflect.ValueOfString("")
	case "cosmos.auth.v1beta1.BaseAccount.pub_key":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.auth.v1beta1.BaseAccount.account_number":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.auth.v1beta1.BaseAccount.sequence":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.auth.v1beta1.BaseAccount"))
		}
		panic(fmt.Errorf("message cosmos.auth.v1beta1.BaseAccount does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BaseAccount) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.auth.v1beta1.BaseAccount", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BaseAccount) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BaseAccount) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_BaseAccount) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BaseAccount) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BaseAccount)
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
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.PubKey != nil {
			l = options.Size(x.PubKey)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.AccountNumber != 0 {
			n += 1 + runtime.Sov(uint64(x.AccountNumber))
		}
		if x.Sequence != 0 {
			n += 1 + runtime.Sov(uint64(x.Sequence))
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
		x := input.Message.Interface().(*BaseAccount)
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
		if x.Sequence != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Sequence))
			i--
			dAtA[i] = 0x20
		}
		if x.AccountNumber != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.AccountNumber))
			i--
			dAtA[i] = 0x18
		}
		if x.PubKey != nil {
			encoded, err := options.Marshal(x.PubKey)
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
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
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
		x := input.Message.Interface().(*BaseAccount)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BaseAccount: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BaseAccount: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
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
				if x.PubKey == nil {
					x.PubKey = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.PubKey); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AccountNumber", wireType)
				}
				x.AccountNumber = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.AccountNumber |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
				}
				x.Sequence = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Sequence |= uint64(b&0x7F) << shift
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

var _ protoreflect.List = (*_ModuleAccount_3_list)(nil)

type _ModuleAccount_3_list struct {
	list *[]string
}

func (x *_ModuleAccount_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_ModuleAccount_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_ModuleAccount_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_ModuleAccount_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_ModuleAccount_3_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message ModuleAccount at list field Permissions as it is not of Message kind"))
}

func (x *_ModuleAccount_3_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_ModuleAccount_3_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_ModuleAccount_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_ModuleAccount              protoreflect.MessageDescriptor
	fd_ModuleAccount_base_account protoreflect.FieldDescriptor
	fd_ModuleAccount_name         protoreflect.FieldDescriptor
	fd_ModuleAccount_permissions  protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_auth_v1beta1_auth_proto_init()
	md_ModuleAccount = File_cosmos_auth_v1beta1_auth_proto.Messages().ByName("ModuleAccount")
	fd_ModuleAccount_base_account = md_ModuleAccount.Fields().ByName("base_account")
	fd_ModuleAccount_name = md_ModuleAccount.Fields().ByName("name")
	fd_ModuleAccount_permissions = md_ModuleAccount.Fields().ByName("permissions")
}

var _ protoreflect.Message = (*fastReflection_ModuleAccount)(nil)

type fastReflection_ModuleAccount ModuleAccount

func (x *ModuleAccount) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ModuleAccount)(x)
}

func (x *ModuleAccount) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_auth_v1beta1_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ModuleAccount_messageType fastReflection_ModuleAccount_messageType
var _ protoreflect.MessageType = fastReflection_ModuleAccount_messageType{}

type fastReflection_ModuleAccount_messageType struct{}

func (x fastReflection_ModuleAccount_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ModuleAccount)(nil)
}
func (x fastReflection_ModuleAccount_messageType) New() protoreflect.Message {
	return new(fastReflection_ModuleAccount)
}
func (x fastReflection_ModuleAccount_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ModuleAccount
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ModuleAccount) Descriptor() protoreflect.MessageDescriptor {
	return md_ModuleAccount
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ModuleAccount) Type() protoreflect.MessageType {
	return _fastReflection_ModuleAccount_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ModuleAccount) New() protoreflect.Message {
	return new(fastReflection_ModuleAccount)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ModuleAccount) Interface() protoreflect.ProtoMessage {
	return (*ModuleAccount)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ModuleAccount) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BaseAccount != nil {
		value := protoreflect.ValueOfMessage(x.BaseAccount.ProtoReflect())
		if !f(fd_ModuleAccount_base_account, value) {
			return
		}
	}
	if x.Name != "" {
		value := protoreflect.ValueOfString(x.Name)
		if !f(fd_ModuleAccount_name, value) {
			return
		}
	}
	if len(x.Permissions) != 0 {
		value := protoreflect.ValueOfList(&_ModuleAccount_3_list{list: &x.Permissions})
		if !f(fd_ModuleAccount_permissions, value) {
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
func (x *fastReflection_ModuleAccount) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.auth.v1beta1.ModuleAccount.base_account":
		return x.BaseAccount != nil
	case "cosmos.auth.v1beta1.ModuleAccount.name":
		return x.Name != ""
	case "cosmos.auth.v1beta1.ModuleAccount.permissions":
		return len(x.Permissions) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.auth.v1beta1.ModuleAccount"))
		}
		panic(fmt.Errorf("message cosmos.auth.v1beta1.ModuleAccount does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ModuleAccount) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.auth.v1beta1.ModuleAccount.base_account":
		x.BaseAccount = nil
	case "cosmos.auth.v1beta1.ModuleAccount.name":
		x.Name = ""
	case "cosmos.auth.v1beta1.ModuleAccount.permissions":
		x.Permissions = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.auth.v1beta1.ModuleAccount"))
		}
		panic(fmt.Errorf("message cosmos.auth.v1beta1.ModuleAccount does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ModuleAccount) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.auth.v1beta1.ModuleAccount.base_account":
		value := x.BaseAccount
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.auth.v1beta1.ModuleAccount.name":
		value := x.Name
		return protoreflect.ValueOfString(value)
	case "cosmos.auth.v1beta1.ModuleAccount.permissions":
		if len(x.Permissions) == 0 {
			return protoreflect.ValueOfList(&_ModuleAccount_3_list{})
		}
		listValue := &_ModuleAccount_3_list{list: &x.Permissions}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.auth.v1beta1.ModuleAccount"))
		}
		panic(fmt.Errorf("message cosmos.auth.v1beta1.ModuleAccount does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ModuleAccount) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.auth.v1beta1.ModuleAccount.base_account":
		x.BaseAccount = value.Message().Interface().(*BaseAccount)
	case "cosmos.auth.v1beta1.ModuleAccount.name":
		x.Name = value.Interface().(string)
	case "cosmos.auth.v1beta1.ModuleAccount.permissions":
		lv := value.List()
		clv := lv.(*_ModuleAccount_3_list)
		x.Permissions = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.auth.v1beta1.ModuleAccount"))
		}
		panic(fmt.Errorf("message cosmos.auth.v1beta1.ModuleAccount does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ModuleAccount) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.auth.v1beta1.ModuleAccount.base_account":
		if x.BaseAccount == nil {
			x.BaseAccount = new(BaseAccount)
		}
		return protoreflect.ValueOfMessage(x.BaseAccount.ProtoReflect())
	case "cosmos.auth.v1beta1.ModuleAccount.permissions":
		if x.Permissions == nil {
			x.Permissions = []string{}
		}
		value := &_ModuleAccount_3_list{list: &x.Permissions}
		return protoreflect.ValueOfList(value)
	case "cosmos.auth.v1beta1.ModuleAccount.name":
		panic(fmt.Errorf("field name of message cosmos.auth.v1beta1.ModuleAccount is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.auth.v1beta1.ModuleAccount"))
		}
		panic(fmt.Errorf("message cosmos.auth.v1beta1.ModuleAccount does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ModuleAccount) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.auth.v1beta1.ModuleAccount.base_account":
		m := new(BaseAccount)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.auth.v1beta1.ModuleAccount.name":
		return protoreflect.ValueOfString("")
	case "cosmos.auth.v1beta1.ModuleAccount.permissions":
		list := []string{}
		return protoreflect.ValueOfList(&_ModuleAccount_3_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.auth.v1beta1.ModuleAccount"))
		}
		panic(fmt.Errorf("message cosmos.auth.v1beta1.ModuleAccount does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ModuleAccount) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.auth.v1beta1.ModuleAccount", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ModuleAccount) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ModuleAccount) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ModuleAccount) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ModuleAccount) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ModuleAccount)
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
		if x.BaseAccount != nil {
			l = options.Size(x.BaseAccount)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Name)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Permissions) > 0 {
			for _, s := range x.Permissions {
				l = len(s)
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
		x := input.Message.Interface().(*ModuleAccount)
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
		if len(x.Permissions) > 0 {
			for iNdEx := len(x.Permissions) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.Permissions[iNdEx])
				copy(dAtA[i:], x.Permissions[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Permissions[iNdEx])))
				i--
				dAtA[i] = 0x1a
			}
		}
		if len(x.Name) > 0 {
			i -= len(x.Name)
			copy(dAtA[i:], x.Name)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Name)))
			i--
			dAtA[i] = 0x12
		}
		if x.BaseAccount != nil {
			encoded, err := options.Marshal(x.BaseAccount)
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
		x := input.Message.Interface().(*ModuleAccount)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ModuleAccount: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ModuleAccount: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
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
				if x.BaseAccount == nil {
					x.BaseAccount = &BaseAccount{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.BaseAccount); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
				x.Name = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
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
				x.Permissions = append(x.Permissions, string(dAtA[iNdEx:postIndex]))
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
	md_Params                           protoreflect.MessageDescriptor
	fd_Params_max_memo_characters       protoreflect.FieldDescriptor
	fd_Params_tx_sig_limit              protoreflect.FieldDescriptor
	fd_Params_tx_size_cost_per_byte     protoreflect.FieldDescriptor
	fd_Params_sig_verify_cost_ed25519   protoreflect.FieldDescriptor
	fd_Params_sig_verify_cost_secp256k1 protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_auth_v1beta1_auth_proto_init()
	md_Params = File_cosmos_auth_v1beta1_auth_proto.Messages().ByName("Params")
	fd_Params_max_memo_characters = md_Params.Fields().ByName("max_memo_characters")
	fd_Params_tx_sig_limit = md_Params.Fields().ByName("tx_sig_limit")
	fd_Params_tx_size_cost_per_byte = md_Params.Fields().ByName("tx_size_cost_per_byte")
	fd_Params_sig_verify_cost_ed25519 = md_Params.Fields().ByName("sig_verify_cost_ed25519")
	fd_Params_sig_verify_cost_secp256k1 = md_Params.Fields().ByName("sig_verify_cost_secp256k1")
}

var _ protoreflect.Message = (*fastReflection_Params)(nil)

type fastReflection_Params Params

func (x *Params) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Params)(x)
}

func (x *Params) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_auth_v1beta1_auth_proto_msgTypes[2]
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
	if x.MaxMemoCharacters != uint64(0) {
		value := protoreflect.ValueOfUint64(x.MaxMemoCharacters)
		if !f(fd_Params_max_memo_characters, value) {
			return
		}
	}
	if x.TxSigLimit != uint64(0) {
		value := protoreflect.ValueOfUint64(x.TxSigLimit)
		if !f(fd_Params_tx_sig_limit, value) {
			return
		}
	}
	if x.TxSizeCostPerByte != uint64(0) {
		value := protoreflect.ValueOfUint64(x.TxSizeCostPerByte)
		if !f(fd_Params_tx_size_cost_per_byte, value) {
			return
		}
	}
	if x.SigVerifyCostEd25519 != uint64(0) {
		value := protoreflect.ValueOfUint64(x.SigVerifyCostEd25519)
		if !f(fd_Params_sig_verify_cost_ed25519, value) {
			return
		}
	}
	if x.SigVerifyCostSecp256K1 != uint64(0) {
		value := protoreflect.ValueOfUint64(x.SigVerifyCostSecp256K1)
		if !f(fd_Params_sig_verify_cost_secp256k1, value) {
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
	case "cosmos.auth.v1beta1.Params.max_memo_characters":
		return x.MaxMemoCharacters != uint64(0)
	case "cosmos.auth.v1beta1.Params.tx_sig_limit":
		return x.TxSigLimit != uint64(0)
	case "cosmos.auth.v1beta1.Params.tx_size_cost_per_byte":
		return x.TxSizeCostPerByte != uint64(0)
	case "cosmos.auth.v1beta1.Params.sig_verify_cost_ed25519":
		return x.SigVerifyCostEd25519 != uint64(0)
	case "cosmos.auth.v1beta1.Params.sig_verify_cost_secp256k1":
		return x.SigVerifyCostSecp256K1 != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.auth.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.auth.v1beta1.Params does not contain field %s", fd.FullName()))
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
	case "cosmos.auth.v1beta1.Params.max_memo_characters":
		x.MaxMemoCharacters = uint64(0)
	case "cosmos.auth.v1beta1.Params.tx_sig_limit":
		x.TxSigLimit = uint64(0)
	case "cosmos.auth.v1beta1.Params.tx_size_cost_per_byte":
		x.TxSizeCostPerByte = uint64(0)
	case "cosmos.auth.v1beta1.Params.sig_verify_cost_ed25519":
		x.SigVerifyCostEd25519 = uint64(0)
	case "cosmos.auth.v1beta1.Params.sig_verify_cost_secp256k1":
		x.SigVerifyCostSecp256K1 = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.auth.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.auth.v1beta1.Params does not contain field %s", fd.FullName()))
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
	case "cosmos.auth.v1beta1.Params.max_memo_characters":
		value := x.MaxMemoCharacters
		return protoreflect.ValueOfUint64(value)
	case "cosmos.auth.v1beta1.Params.tx_sig_limit":
		value := x.TxSigLimit
		return protoreflect.ValueOfUint64(value)
	case "cosmos.auth.v1beta1.Params.tx_size_cost_per_byte":
		value := x.TxSizeCostPerByte
		return protoreflect.ValueOfUint64(value)
	case "cosmos.auth.v1beta1.Params.sig_verify_cost_ed25519":
		value := x.SigVerifyCostEd25519
		return protoreflect.ValueOfUint64(value)
	case "cosmos.auth.v1beta1.Params.sig_verify_cost_secp256k1":
		value := x.SigVerifyCostSecp256K1
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.auth.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.auth.v1beta1.Params does not contain field %s", descriptor.FullName()))
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
	case "cosmos.auth.v1beta1.Params.max_memo_characters":
		x.MaxMemoCharacters = value.Uint()
	case "cosmos.auth.v1beta1.Params.tx_sig_limit":
		x.TxSigLimit = value.Uint()
	case "cosmos.auth.v1beta1.Params.tx_size_cost_per_byte":
		x.TxSizeCostPerByte = value.Uint()
	case "cosmos.auth.v1beta1.Params.sig_verify_cost_ed25519":
		x.SigVerifyCostEd25519 = value.Uint()
	case "cosmos.auth.v1beta1.Params.sig_verify_cost_secp256k1":
		x.SigVerifyCostSecp256K1 = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.auth.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.auth.v1beta1.Params does not contain field %s", fd.FullName()))
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
	case "cosmos.auth.v1beta1.Params.max_memo_characters":
		panic(fmt.Errorf("field max_memo_characters of message cosmos.auth.v1beta1.Params is not mutable"))
	case "cosmos.auth.v1beta1.Params.tx_sig_limit":
		panic(fmt.Errorf("field tx_sig_limit of message cosmos.auth.v1beta1.Params is not mutable"))
	case "cosmos.auth.v1beta1.Params.tx_size_cost_per_byte":
		panic(fmt.Errorf("field tx_size_cost_per_byte of message cosmos.auth.v1beta1.Params is not mutable"))
	case "cosmos.auth.v1beta1.Params.sig_verify_cost_ed25519":
		panic(fmt.Errorf("field sig_verify_cost_ed25519 of message cosmos.auth.v1beta1.Params is not mutable"))
	case "cosmos.auth.v1beta1.Params.sig_verify_cost_secp256k1":
		panic(fmt.Errorf("field sig_verify_cost_secp256k1 of message cosmos.auth.v1beta1.Params is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.auth.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.auth.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Params) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.auth.v1beta1.Params.max_memo_characters":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.auth.v1beta1.Params.tx_sig_limit":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.auth.v1beta1.Params.tx_size_cost_per_byte":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.auth.v1beta1.Params.sig_verify_cost_ed25519":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.auth.v1beta1.Params.sig_verify_cost_secp256k1":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.auth.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.auth.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Params) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.auth.v1beta1.Params", d.FullName()))
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
		if x.MaxMemoCharacters != 0 {
			n += 1 + runtime.Sov(uint64(x.MaxMemoCharacters))
		}
		if x.TxSigLimit != 0 {
			n += 1 + runtime.Sov(uint64(x.TxSigLimit))
		}
		if x.TxSizeCostPerByte != 0 {
			n += 1 + runtime.Sov(uint64(x.TxSizeCostPerByte))
		}
		if x.SigVerifyCostEd25519 != 0 {
			n += 1 + runtime.Sov(uint64(x.SigVerifyCostEd25519))
		}
		if x.SigVerifyCostSecp256K1 != 0 {
			n += 1 + runtime.Sov(uint64(x.SigVerifyCostSecp256K1))
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
		if x.SigVerifyCostSecp256K1 != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.SigVerifyCostSecp256K1))
			i--
			dAtA[i] = 0x28
		}
		if x.SigVerifyCostEd25519 != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.SigVerifyCostEd25519))
			i--
			dAtA[i] = 0x20
		}
		if x.TxSizeCostPerByte != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.TxSizeCostPerByte))
			i--
			dAtA[i] = 0x18
		}
		if x.TxSigLimit != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.TxSigLimit))
			i--
			dAtA[i] = 0x10
		}
		if x.MaxMemoCharacters != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.MaxMemoCharacters))
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
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MaxMemoCharacters", wireType)
				}
				x.MaxMemoCharacters = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.MaxMemoCharacters |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TxSigLimit", wireType)
				}
				x.TxSigLimit = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.TxSigLimit |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TxSizeCostPerByte", wireType)
				}
				x.TxSizeCostPerByte = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.TxSizeCostPerByte |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SigVerifyCostEd25519", wireType)
				}
				x.SigVerifyCostEd25519 = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.SigVerifyCostEd25519 |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SigVerifyCostSecp256K1", wireType)
				}
				x.SigVerifyCostSecp256K1 = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.SigVerifyCostSecp256K1 |= uint64(b&0x7F) << shift
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.19.1
// source: cosmos/auth/v1beta1/auth.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// BaseAccount defines a base account type. It contains all the necessary fields
// for basic account functionality. Any custom account type should extend this
// type for additional functionality (e.g. vesting).
type BaseAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address       string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PubKey        *anypb.Any `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	AccountNumber uint64     `protobuf:"varint,3,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	Sequence      uint64     `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *BaseAccount) Reset() {
	*x = BaseAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_auth_v1beta1_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseAccount) ProtoMessage() {}

// Deprecated: Use BaseAccount.ProtoReflect.Descriptor instead.
func (*BaseAccount) Descriptor() ([]byte, []int) {
	return file_cosmos_auth_v1beta1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *BaseAccount) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *BaseAccount) GetPubKey() *anypb.Any {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *BaseAccount) GetAccountNumber() uint64 {
	if x != nil {
		return x.AccountNumber
	}
	return 0
}

func (x *BaseAccount) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

// ModuleAccount defines an account for modules that holds coins on a pool.
type ModuleAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseAccount *BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3" json:"base_account,omitempty"`
	Name        string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Permissions []string     `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *ModuleAccount) Reset() {
	*x = ModuleAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_auth_v1beta1_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModuleAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuleAccount) ProtoMessage() {}

// Deprecated: Use ModuleAccount.ProtoReflect.Descriptor instead.
func (*ModuleAccount) Descriptor() ([]byte, []int) {
	return file_cosmos_auth_v1beta1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *ModuleAccount) GetBaseAccount() *BaseAccount {
	if x != nil {
		return x.BaseAccount
	}
	return nil
}

func (x *ModuleAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModuleAccount) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

// Params defines the parameters for the auth module.
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxMemoCharacters      uint64 `protobuf:"varint,1,opt,name=max_memo_characters,json=maxMemoCharacters,proto3" json:"max_memo_characters,omitempty"`
	TxSigLimit             uint64 `protobuf:"varint,2,opt,name=tx_sig_limit,json=txSigLimit,proto3" json:"tx_sig_limit,omitempty"`
	TxSizeCostPerByte      uint64 `protobuf:"varint,3,opt,name=tx_size_cost_per_byte,json=txSizeCostPerByte,proto3" json:"tx_size_cost_per_byte,omitempty"`
	SigVerifyCostEd25519   uint64 `protobuf:"varint,4,opt,name=sig_verify_cost_ed25519,json=sigVerifyCostEd25519,proto3" json:"sig_verify_cost_ed25519,omitempty"`
	SigVerifyCostSecp256K1 uint64 `protobuf:"varint,5,opt,name=sig_verify_cost_secp256k1,json=sigVerifyCostSecp256k1,proto3" json:"sig_verify_cost_secp256k1,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_auth_v1beta1_auth_proto_msgTypes[2]
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
	return file_cosmos_auth_v1beta1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *Params) GetMaxMemoCharacters() uint64 {
	if x != nil {
		return x.MaxMemoCharacters
	}
	return 0
}

func (x *Params) GetTxSigLimit() uint64 {
	if x != nil {
		return x.TxSigLimit
	}
	return 0
}

func (x *Params) GetTxSizeCostPerByte() uint64 {
	if x != nil {
		return x.TxSizeCostPerByte
	}
	return 0
}

func (x *Params) GetSigVerifyCostEd25519() uint64 {
	if x != nil {
		return x.SigVerifyCostEd25519
	}
	return 0
}

func (x *Params) GetSigVerifyCostSecp256K1() uint64 {
	if x != nil {
		return x.SigVerifyCostSecp256K1
	}
	return 0
}

var File_cosmos_auth_v1beta1_auth_proto protoreflect.FileDescriptor

var file_cosmos_auth_v1beta1_auth_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe7, 0x01, 0x0a, 0x0b, 0x42, 0x61, 0x73, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x32, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x47, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x18, 0xea, 0xde,
	0x1f, 0x14, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x2c, 0x6f, 0x6d, 0x69,
	0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x25,
	0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x3a, 0x18, 0x88, 0xa0, 0x1f, 0x00, 0x98, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0xca,
	0xb4, 0x2d, 0x08, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x22, 0xac, 0x01, 0x0a, 0x0d,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x49, 0x0a,
	0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x04, 0xd0, 0xde, 0x1f, 0x01, 0x52, 0x0b, 0x62, 0x61, 0x73,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x1a,
	0x88, 0xa0, 0x1f, 0x00, 0x98, 0xa0, 0x1f, 0x00, 0xca, 0xb4, 0x2d, 0x0e, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x22, 0xbe, 0x02, 0x0a, 0x06, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x6d, 0x65, 0x6d,
	0x6f, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x4d, 0x65, 0x6d, 0x6f, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x74, 0x78, 0x5f, 0x73, 0x69, 0x67, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x78, 0x53,
	0x69, 0x67, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x30, 0x0a, 0x15, 0x74, 0x78, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x74, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x43, 0x6f,
	0x73, 0x74, 0x50, 0x65, 0x72, 0x42, 0x79, 0x74, 0x65, 0x12, 0x4f, 0x0a, 0x17, 0x73, 0x69, 0x67,
	0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x65, 0x64, 0x32,
	0x35, 0x35, 0x31, 0x39, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x42, 0x18, 0xe2, 0xde, 0x1f, 0x14,
	0x53, 0x69, 0x67, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x73, 0x74, 0x45, 0x44, 0x32,
	0x35, 0x35, 0x31, 0x39, 0x52, 0x14, 0x73, 0x69, 0x67, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43,
	0x6f, 0x73, 0x74, 0x45, 0x64, 0x32, 0x35, 0x35, 0x31, 0x39, 0x12, 0x55, 0x0a, 0x19, 0x73, 0x69,
	0x67, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65,
	0x63, 0x70, 0x32, 0x35, 0x36, 0x6b, 0x31, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x42, 0x1a, 0xe2,
	0xde, 0x1f, 0x16, 0x53, 0x69, 0x67, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x73, 0x74,
	0x53, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x6b, 0x31, 0x52, 0x16, 0x73, 0x69, 0x67, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x6b,
	0x31, 0x3a, 0x08, 0x98, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x01, 0x42, 0xd4, 0x01, 0x0a, 0x17,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x09, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x41, 0x58, 0xaa, 0x02, 0x13, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xca, 0x02, 0x13, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x41, 0x75, 0x74, 0x68,
	0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x1f, 0x43, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x43, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_auth_v1beta1_auth_proto_rawDescOnce sync.Once
	file_cosmos_auth_v1beta1_auth_proto_rawDescData = file_cosmos_auth_v1beta1_auth_proto_rawDesc
)

func file_cosmos_auth_v1beta1_auth_proto_rawDescGZIP() []byte {
	file_cosmos_auth_v1beta1_auth_proto_rawDescOnce.Do(func() {
		file_cosmos_auth_v1beta1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_auth_v1beta1_auth_proto_rawDescData)
	})
	return file_cosmos_auth_v1beta1_auth_proto_rawDescData
}

var file_cosmos_auth_v1beta1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cosmos_auth_v1beta1_auth_proto_goTypes = []interface{}{
	(*BaseAccount)(nil),   // 0: cosmos.auth.v1beta1.BaseAccount
	(*ModuleAccount)(nil), // 1: cosmos.auth.v1beta1.ModuleAccount
	(*Params)(nil),        // 2: cosmos.auth.v1beta1.Params
	(*anypb.Any)(nil),     // 3: google.protobuf.Any
}
var file_cosmos_auth_v1beta1_auth_proto_depIdxs = []int32{
	3, // 0: cosmos.auth.v1beta1.BaseAccount.pub_key:type_name -> google.protobuf.Any
	0, // 1: cosmos.auth.v1beta1.ModuleAccount.base_account:type_name -> cosmos.auth.v1beta1.BaseAccount
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cosmos_auth_v1beta1_auth_proto_init() }
func file_cosmos_auth_v1beta1_auth_proto_init() {
	if File_cosmos_auth_v1beta1_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_auth_v1beta1_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseAccount); i {
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
		file_cosmos_auth_v1beta1_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModuleAccount); i {
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
		file_cosmos_auth_v1beta1_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_auth_v1beta1_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_auth_v1beta1_auth_proto_goTypes,
		DependencyIndexes: file_cosmos_auth_v1beta1_auth_proto_depIdxs,
		MessageInfos:      file_cosmos_auth_v1beta1_auth_proto_msgTypes,
	}.Build()
	File_cosmos_auth_v1beta1_auth_proto = out.File
	file_cosmos_auth_v1beta1_auth_proto_rawDesc = nil
	file_cosmos_auth_v1beta1_auth_proto_goTypes = nil
	file_cosmos_auth_v1beta1_auth_proto_depIdxs = nil
}
