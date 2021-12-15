package crypto

import (
	fmt "fmt"
	io "io"
	reflect "reflect"
	sync "sync"

	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

var (
	md_PublicKey           protoreflect.MessageDescriptor
	fd_PublicKey_ed25519   protoreflect.FieldDescriptor
	fd_PublicKey_secp256k1 protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_crypto_keys_proto_init()
	md_PublicKey = File_tendermint_crypto_keys_proto.Messages().ByName("PublicKey")
	fd_PublicKey_ed25519 = md_PublicKey.Fields().ByName("ed25519")
	fd_PublicKey_secp256k1 = md_PublicKey.Fields().ByName("secp256k1")
}

var _ protoreflect.Message = (*fastReflection_PublicKey)(nil)

type fastReflection_PublicKey PublicKey

func (x *PublicKey) ProtoReflect() protoreflect.Message {
	return (*fastReflection_PublicKey)(x)
}

func (x *PublicKey) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_crypto_keys_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_PublicKey_messageType fastReflection_PublicKey_messageType
var _ protoreflect.MessageType = fastReflection_PublicKey_messageType{}

type fastReflection_PublicKey_messageType struct{}

func (x fastReflection_PublicKey_messageType) Zero() protoreflect.Message {
	return (*fastReflection_PublicKey)(nil)
}
func (x fastReflection_PublicKey_messageType) New() protoreflect.Message {
	return new(fastReflection_PublicKey)
}
func (x fastReflection_PublicKey_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_PublicKey
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_PublicKey) Descriptor() protoreflect.MessageDescriptor {
	return md_PublicKey
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_PublicKey) Type() protoreflect.MessageType {
	return _fastReflection_PublicKey_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_PublicKey) New() protoreflect.Message {
	return new(fastReflection_PublicKey)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_PublicKey) Interface() protoreflect.ProtoMessage {
	return (*PublicKey)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_PublicKey) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Sum != nil {
		switch o := x.Sum.(type) {
		case *PublicKey_Ed25519:
			v := o.Ed25519
			value := protoreflect.ValueOfBytes(v)
			if !f(fd_PublicKey_ed25519, value) {
				return
			}
		case *PublicKey_Secp256K1:
			v := o.Secp256K1
			value := protoreflect.ValueOfBytes(v)
			if !f(fd_PublicKey_secp256k1, value) {
				return
			}
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
func (x *fastReflection_PublicKey) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.crypto.PublicKey.ed25519":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*PublicKey_Ed25519); ok {
			return true
		} else {
			return false
		}
	case "tendermint.crypto.PublicKey.secp256k1":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*PublicKey_Secp256K1); ok {
			return true
		} else {
			return false
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.crypto.PublicKey"))
		}
		panic(fmt.Errorf("message tendermint.crypto.PublicKey does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PublicKey) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.crypto.PublicKey.ed25519":
		x.Sum = nil
	case "tendermint.crypto.PublicKey.secp256k1":
		x.Sum = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.crypto.PublicKey"))
		}
		panic(fmt.Errorf("message tendermint.crypto.PublicKey does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_PublicKey) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.crypto.PublicKey.ed25519":
		if x.Sum == nil {
			return protoreflect.ValueOfBytes(nil)
		} else if v, ok := x.Sum.(*PublicKey_Ed25519); ok {
			return protoreflect.ValueOfBytes(v.Ed25519)
		} else {
			return protoreflect.ValueOfBytes(nil)
		}
	case "tendermint.crypto.PublicKey.secp256k1":
		if x.Sum == nil {
			return protoreflect.ValueOfBytes(nil)
		} else if v, ok := x.Sum.(*PublicKey_Secp256K1); ok {
			return protoreflect.ValueOfBytes(v.Secp256K1)
		} else {
			return protoreflect.ValueOfBytes(nil)
		}
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.crypto.PublicKey"))
		}
		panic(fmt.Errorf("message tendermint.crypto.PublicKey does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_PublicKey) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.crypto.PublicKey.ed25519":
		cv := value.Bytes()
		x.Sum = &PublicKey_Ed25519{Ed25519: cv}
	case "tendermint.crypto.PublicKey.secp256k1":
		cv := value.Bytes()
		x.Sum = &PublicKey_Secp256K1{Secp256K1: cv}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.crypto.PublicKey"))
		}
		panic(fmt.Errorf("message tendermint.crypto.PublicKey does not contain field %s", fd.FullName()))
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
func (x *fastReflection_PublicKey) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.crypto.PublicKey.ed25519":
		panic(fmt.Errorf("field ed25519 of message tendermint.crypto.PublicKey is not mutable"))
	case "tendermint.crypto.PublicKey.secp256k1":
		panic(fmt.Errorf("field secp256k1 of message tendermint.crypto.PublicKey is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.crypto.PublicKey"))
		}
		panic(fmt.Errorf("message tendermint.crypto.PublicKey does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_PublicKey) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.crypto.PublicKey.ed25519":
		return protoreflect.ValueOfBytes(nil)
	case "tendermint.crypto.PublicKey.secp256k1":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.crypto.PublicKey"))
		}
		panic(fmt.Errorf("message tendermint.crypto.PublicKey does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_PublicKey) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "tendermint.crypto.PublicKey.sum":
		if x.Sum == nil {
			return nil
		}
		switch x.Sum.(type) {
		case *PublicKey_Ed25519:
			return x.Descriptor().Fields().ByName("ed25519")
		case *PublicKey_Secp256K1:
			return x.Descriptor().Fields().ByName("secp256k1")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.crypto.PublicKey", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_PublicKey) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PublicKey) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_PublicKey) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_PublicKey) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*PublicKey)
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
		switch x := x.Sum.(type) {
		case *PublicKey_Ed25519:
			if x == nil {
				break
			}
			l = len(x.Ed25519)
			n += 1 + l + runtime.Sov(uint64(l))
		case *PublicKey_Secp256K1:
			if x == nil {
				break
			}
			l = len(x.Secp256K1)
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
		x := input.Message.Interface().(*PublicKey)
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
		switch x := x.Sum.(type) {
		case *PublicKey_Ed25519:
			i -= len(x.Ed25519)
			copy(dAtA[i:], x.Ed25519)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Ed25519)))
			i--
			dAtA[i] = 0xa
		case *PublicKey_Secp256K1:
			i -= len(x.Secp256K1)
			copy(dAtA[i:], x.Secp256K1)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Secp256K1)))
			i--
			dAtA[i] = 0x12
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
		x := input.Message.Interface().(*PublicKey)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PublicKey: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PublicKey: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Ed25519", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v := make([]byte, postIndex-iNdEx)
				copy(v, dAtA[iNdEx:postIndex])
				x.Sum = &PublicKey_Ed25519{v}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Secp256K1", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v := make([]byte, postIndex-iNdEx)
				copy(v, dAtA[iNdEx:postIndex])
				x.Sum = &PublicKey_Secp256K1{v}
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
// source: tendermint/crypto/keys.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PublicKey defines the keys available for use with Tendermint Validators
type PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Sum:
	//	*PublicKey_Ed25519
	//	*PublicKey_Secp256K1
	Sum isPublicKey_Sum `protobuf_oneof:"sum"`
}

func (x *PublicKey) Reset() {
	*x = PublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_crypto_keys_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicKey) ProtoMessage() {}

// Deprecated: Use PublicKey.ProtoReflect.Descriptor instead.
func (*PublicKey) Descriptor() ([]byte, []int) {
	return file_tendermint_crypto_keys_proto_rawDescGZIP(), []int{0}
}

func (x *PublicKey) GetSum() isPublicKey_Sum {
	if x != nil {
		return x.Sum
	}
	return nil
}

func (x *PublicKey) GetEd25519() []byte {
	if x, ok := x.GetSum().(*PublicKey_Ed25519); ok {
		return x.Ed25519
	}
	return nil
}

func (x *PublicKey) GetSecp256K1() []byte {
	if x, ok := x.GetSum().(*PublicKey_Secp256K1); ok {
		return x.Secp256K1
	}
	return nil
}

type isPublicKey_Sum interface {
	isPublicKey_Sum()
}

type PublicKey_Ed25519 struct {
	Ed25519 []byte `protobuf:"bytes,1,opt,name=ed25519,proto3,oneof"`
}

type PublicKey_Secp256K1 struct {
	Secp256K1 []byte `protobuf:"bytes,2,opt,name=secp256k1,proto3,oneof"`
}

func (*PublicKey_Ed25519) isPublicKey_Sum() {}

func (*PublicKey_Secp256K1) isPublicKey_Sum() {}

var File_tendermint_crypto_keys_proto protoreflect.FileDescriptor

var file_tendermint_crypto_keys_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x07, 0x65, 0x64, 0x32, 0x35, 0x35, 0x31, 0x39, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x07, 0x65, 0x64, 0x32, 0x35, 0x35, 0x31, 0x39,
	0x12, 0x1e, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x6b, 0x31, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x09, 0x73, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x6b, 0x31,
	0x3a, 0x08, 0xe8, 0xa0, 0x1f, 0x01, 0xe8, 0xa1, 0x1f, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x73, 0x75,
	0x6d, 0x42, 0xbb, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x42, 0x09, 0x4b, 0x65, 0x79,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x54,
	0x43, 0x58, 0xaa, 0x02, 0x11, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0xca, 0x02, 0x11, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x5c, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0xe2, 0x02, 0x1d, 0x54, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x5c, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x54, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x3a, 0x3a, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tendermint_crypto_keys_proto_rawDescOnce sync.Once
	file_tendermint_crypto_keys_proto_rawDescData = file_tendermint_crypto_keys_proto_rawDesc
)

func file_tendermint_crypto_keys_proto_rawDescGZIP() []byte {
	file_tendermint_crypto_keys_proto_rawDescOnce.Do(func() {
		file_tendermint_crypto_keys_proto_rawDescData = protoimpl.X.CompressGZIP(file_tendermint_crypto_keys_proto_rawDescData)
	})
	return file_tendermint_crypto_keys_proto_rawDescData
}

var file_tendermint_crypto_keys_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tendermint_crypto_keys_proto_goTypes = []interface{}{
	(*PublicKey)(nil), // 0: tendermint.crypto.PublicKey
}
var file_tendermint_crypto_keys_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tendermint_crypto_keys_proto_init() }
func file_tendermint_crypto_keys_proto_init() {
	if File_tendermint_crypto_keys_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tendermint_crypto_keys_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicKey); i {
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
	file_tendermint_crypto_keys_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PublicKey_Ed25519)(nil),
		(*PublicKey_Secp256K1)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tendermint_crypto_keys_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tendermint_crypto_keys_proto_goTypes,
		DependencyIndexes: file_tendermint_crypto_keys_proto_depIdxs,
		MessageInfos:      file_tendermint_crypto_keys_proto_msgTypes,
	}.Build()
	File_tendermint_crypto_keys_proto = out.File
	file_tendermint_crypto_keys_proto_rawDesc = nil
	file_tendermint_crypto_keys_proto_goTypes = nil
	file_tendermint_crypto_keys_proto_depIdxs = nil
}
