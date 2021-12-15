package feegrantv1beta1

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
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
)

var _ protoreflect.List = (*_BasicAllowance_1_list)(nil)

type _BasicAllowance_1_list struct {
	list *[]*v1beta1.Coin
}

func (x *_BasicAllowance_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_BasicAllowance_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_BasicAllowance_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	(*x.list)[i] = concreteValue
}

func (x *_BasicAllowance_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_BasicAllowance_1_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.Coin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_BasicAllowance_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_BasicAllowance_1_list) NewElement() protoreflect.Value {
	v := new(v1beta1.Coin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_BasicAllowance_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_BasicAllowance             protoreflect.MessageDescriptor
	fd_BasicAllowance_spend_limit protoreflect.FieldDescriptor
	fd_BasicAllowance_expiration  protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_feegrant_v1beta1_feegrant_proto_init()
	md_BasicAllowance = File_cosmos_feegrant_v1beta1_feegrant_proto.Messages().ByName("BasicAllowance")
	fd_BasicAllowance_spend_limit = md_BasicAllowance.Fields().ByName("spend_limit")
	fd_BasicAllowance_expiration = md_BasicAllowance.Fields().ByName("expiration")
}

var _ protoreflect.Message = (*fastReflection_BasicAllowance)(nil)

type fastReflection_BasicAllowance BasicAllowance

func (x *BasicAllowance) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BasicAllowance)(x)
}

func (x *BasicAllowance) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_feegrant_v1beta1_feegrant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BasicAllowance_messageType fastReflection_BasicAllowance_messageType
var _ protoreflect.MessageType = fastReflection_BasicAllowance_messageType{}

type fastReflection_BasicAllowance_messageType struct{}

func (x fastReflection_BasicAllowance_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BasicAllowance)(nil)
}
func (x fastReflection_BasicAllowance_messageType) New() protoreflect.Message {
	return new(fastReflection_BasicAllowance)
}
func (x fastReflection_BasicAllowance_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BasicAllowance
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BasicAllowance) Descriptor() protoreflect.MessageDescriptor {
	return md_BasicAllowance
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BasicAllowance) Type() protoreflect.MessageType {
	return _fastReflection_BasicAllowance_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BasicAllowance) New() protoreflect.Message {
	return new(fastReflection_BasicAllowance)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BasicAllowance) Interface() protoreflect.ProtoMessage {
	return (*BasicAllowance)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BasicAllowance) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.SpendLimit) != 0 {
		value := protoreflect.ValueOfList(&_BasicAllowance_1_list{list: &x.SpendLimit})
		if !f(fd_BasicAllowance_spend_limit, value) {
			return
		}
	}
	if x.Expiration != nil {
		value := protoreflect.ValueOfMessage(x.Expiration.ProtoReflect())
		if !f(fd_BasicAllowance_expiration, value) {
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
func (x *fastReflection_BasicAllowance) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.BasicAllowance.spend_limit":
		return len(x.SpendLimit) != 0
	case "cosmos.feegrant.v1beta1.BasicAllowance.expiration":
		return x.Expiration != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.BasicAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.BasicAllowance does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BasicAllowance) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.BasicAllowance.spend_limit":
		x.SpendLimit = nil
	case "cosmos.feegrant.v1beta1.BasicAllowance.expiration":
		x.Expiration = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.BasicAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.BasicAllowance does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BasicAllowance) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.feegrant.v1beta1.BasicAllowance.spend_limit":
		if len(x.SpendLimit) == 0 {
			return protoreflect.ValueOfList(&_BasicAllowance_1_list{})
		}
		listValue := &_BasicAllowance_1_list{list: &x.SpendLimit}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.feegrant.v1beta1.BasicAllowance.expiration":
		value := x.Expiration
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.BasicAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.BasicAllowance does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_BasicAllowance) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.BasicAllowance.spend_limit":
		lv := value.List()
		clv := lv.(*_BasicAllowance_1_list)
		x.SpendLimit = *clv.list
	case "cosmos.feegrant.v1beta1.BasicAllowance.expiration":
		x.Expiration = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.BasicAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.BasicAllowance does not contain field %s", fd.FullName()))
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
func (x *fastReflection_BasicAllowance) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.BasicAllowance.spend_limit":
		if x.SpendLimit == nil {
			x.SpendLimit = []*v1beta1.Coin{}
		}
		value := &_BasicAllowance_1_list{list: &x.SpendLimit}
		return protoreflect.ValueOfList(value)
	case "cosmos.feegrant.v1beta1.BasicAllowance.expiration":
		if x.Expiration == nil {
			x.Expiration = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.Expiration.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.BasicAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.BasicAllowance does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BasicAllowance) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.BasicAllowance.spend_limit":
		list := []*v1beta1.Coin{}
		return protoreflect.ValueOfList(&_BasicAllowance_1_list{list: &list})
	case "cosmos.feegrant.v1beta1.BasicAllowance.expiration":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.BasicAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.BasicAllowance does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BasicAllowance) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.feegrant.v1beta1.BasicAllowance", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BasicAllowance) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BasicAllowance) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_BasicAllowance) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BasicAllowance) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BasicAllowance)
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
		if len(x.SpendLimit) > 0 {
			for _, e := range x.SpendLimit {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Expiration != nil {
			l = options.Size(x.Expiration)
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
		x := input.Message.Interface().(*BasicAllowance)
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
		if x.Expiration != nil {
			encoded, err := options.Marshal(x.Expiration)
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
		if len(x.SpendLimit) > 0 {
			for iNdEx := len(x.SpendLimit) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.SpendLimit[iNdEx])
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
		x := input.Message.Interface().(*BasicAllowance)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BasicAllowance: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BasicAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SpendLimit", wireType)
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
				x.SpendLimit = append(x.SpendLimit, &v1beta1.Coin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.SpendLimit[len(x.SpendLimit)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
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
				if x.Expiration == nil {
					x.Expiration = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Expiration); err != nil {
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

var _ protoreflect.List = (*_PeriodicAllowance_3_list)(nil)

type _PeriodicAllowance_3_list struct {
	list *[]*v1beta1.Coin
}

func (x *_PeriodicAllowance_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_PeriodicAllowance_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_PeriodicAllowance_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	(*x.list)[i] = concreteValue
}

func (x *_PeriodicAllowance_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_PeriodicAllowance_3_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.Coin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_PeriodicAllowance_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_PeriodicAllowance_3_list) NewElement() protoreflect.Value {
	v := new(v1beta1.Coin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_PeriodicAllowance_3_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_PeriodicAllowance_4_list)(nil)

type _PeriodicAllowance_4_list struct {
	list *[]*v1beta1.Coin
}

func (x *_PeriodicAllowance_4_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_PeriodicAllowance_4_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_PeriodicAllowance_4_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	(*x.list)[i] = concreteValue
}

func (x *_PeriodicAllowance_4_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_PeriodicAllowance_4_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.Coin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_PeriodicAllowance_4_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_PeriodicAllowance_4_list) NewElement() protoreflect.Value {
	v := new(v1beta1.Coin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_PeriodicAllowance_4_list) IsValid() bool {
	return x.list != nil
}

var (
	md_PeriodicAllowance                    protoreflect.MessageDescriptor
	fd_PeriodicAllowance_basic              protoreflect.FieldDescriptor
	fd_PeriodicAllowance_period             protoreflect.FieldDescriptor
	fd_PeriodicAllowance_period_spend_limit protoreflect.FieldDescriptor
	fd_PeriodicAllowance_period_can_spend   protoreflect.FieldDescriptor
	fd_PeriodicAllowance_period_reset       protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_feegrant_v1beta1_feegrant_proto_init()
	md_PeriodicAllowance = File_cosmos_feegrant_v1beta1_feegrant_proto.Messages().ByName("PeriodicAllowance")
	fd_PeriodicAllowance_basic = md_PeriodicAllowance.Fields().ByName("basic")
	fd_PeriodicAllowance_period = md_PeriodicAllowance.Fields().ByName("period")
	fd_PeriodicAllowance_period_spend_limit = md_PeriodicAllowance.Fields().ByName("period_spend_limit")
	fd_PeriodicAllowance_period_can_spend = md_PeriodicAllowance.Fields().ByName("period_can_spend")
	fd_PeriodicAllowance_period_reset = md_PeriodicAllowance.Fields().ByName("period_reset")
}

var _ protoreflect.Message = (*fastReflection_PeriodicAllowance)(nil)

type fastReflection_PeriodicAllowance PeriodicAllowance

func (x *PeriodicAllowance) ProtoReflect() protoreflect.Message {
	return (*fastReflection_PeriodicAllowance)(x)
}

func (x *PeriodicAllowance) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_feegrant_v1beta1_feegrant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_PeriodicAllowance_messageType fastReflection_PeriodicAllowance_messageType
var _ protoreflect.MessageType = fastReflection_PeriodicAllowance_messageType{}

type fastReflection_PeriodicAllowance_messageType struct{}

func (x fastReflection_PeriodicAllowance_messageType) Zero() protoreflect.Message {
	return (*fastReflection_PeriodicAllowance)(nil)
}
func (x fastReflection_PeriodicAllowance_messageType) New() protoreflect.Message {
	return new(fastReflection_PeriodicAllowance)
}
func (x fastReflection_PeriodicAllowance_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_PeriodicAllowance
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_PeriodicAllowance) Descriptor() protoreflect.MessageDescriptor {
	return md_PeriodicAllowance
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_PeriodicAllowance) Type() protoreflect.MessageType {
	return _fastReflection_PeriodicAllowance_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_PeriodicAllowance) New() protoreflect.Message {
	return new(fastReflection_PeriodicAllowance)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_PeriodicAllowance) Interface() protoreflect.ProtoMessage {
	return (*PeriodicAllowance)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_PeriodicAllowance) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Basic != nil {
		value := protoreflect.ValueOfMessage(x.Basic.ProtoReflect())
		if !f(fd_PeriodicAllowance_basic, value) {
			return
		}
	}
	if x.Period != nil {
		value := protoreflect.ValueOfMessage(x.Period.ProtoReflect())
		if !f(fd_PeriodicAllowance_period, value) {
			return
		}
	}
	if len(x.PeriodSpendLimit) != 0 {
		value := protoreflect.ValueOfList(&_PeriodicAllowance_3_list{list: &x.PeriodSpendLimit})
		if !f(fd_PeriodicAllowance_period_spend_limit, value) {
			return
		}
	}
	if len(x.PeriodCanSpend) != 0 {
		value := protoreflect.ValueOfList(&_PeriodicAllowance_4_list{list: &x.PeriodCanSpend})
		if !f(fd_PeriodicAllowance_period_can_spend, value) {
			return
		}
	}
	if x.PeriodReset != nil {
		value := protoreflect.ValueOfMessage(x.PeriodReset.ProtoReflect())
		if !f(fd_PeriodicAllowance_period_reset, value) {
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
func (x *fastReflection_PeriodicAllowance) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.basic":
		return x.Basic != nil
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period":
		return x.Period != nil
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period_spend_limit":
		return len(x.PeriodSpendLimit) != 0
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period_can_spend":
		return len(x.PeriodCanSpend) != 0
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period_reset":
		return x.PeriodReset != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.PeriodicAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.PeriodicAllowance does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PeriodicAllowance) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.basic":
		x.Basic = nil
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period":
		x.Period = nil
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period_spend_limit":
		x.PeriodSpendLimit = nil
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period_can_spend":
		x.PeriodCanSpend = nil
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period_reset":
		x.PeriodReset = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.PeriodicAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.PeriodicAllowance does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_PeriodicAllowance) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.basic":
		value := x.Basic
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period":
		value := x.Period
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period_spend_limit":
		if len(x.PeriodSpendLimit) == 0 {
			return protoreflect.ValueOfList(&_PeriodicAllowance_3_list{})
		}
		listValue := &_PeriodicAllowance_3_list{list: &x.PeriodSpendLimit}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period_can_spend":
		if len(x.PeriodCanSpend) == 0 {
			return protoreflect.ValueOfList(&_PeriodicAllowance_4_list{})
		}
		listValue := &_PeriodicAllowance_4_list{list: &x.PeriodCanSpend}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period_reset":
		value := x.PeriodReset
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.PeriodicAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.PeriodicAllowance does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_PeriodicAllowance) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.basic":
		x.Basic = value.Message().Interface().(*BasicAllowance)
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period":
		x.Period = value.Message().Interface().(*durationpb.Duration)
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period_spend_limit":
		lv := value.List()
		clv := lv.(*_PeriodicAllowance_3_list)
		x.PeriodSpendLimit = *clv.list
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period_can_spend":
		lv := value.List()
		clv := lv.(*_PeriodicAllowance_4_list)
		x.PeriodCanSpend = *clv.list
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period_reset":
		x.PeriodReset = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.PeriodicAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.PeriodicAllowance does not contain field %s", fd.FullName()))
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
func (x *fastReflection_PeriodicAllowance) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.basic":
		if x.Basic == nil {
			x.Basic = new(BasicAllowance)
		}
		return protoreflect.ValueOfMessage(x.Basic.ProtoReflect())
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period":
		if x.Period == nil {
			x.Period = new(durationpb.Duration)
		}
		return protoreflect.ValueOfMessage(x.Period.ProtoReflect())
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period_spend_limit":
		if x.PeriodSpendLimit == nil {
			x.PeriodSpendLimit = []*v1beta1.Coin{}
		}
		value := &_PeriodicAllowance_3_list{list: &x.PeriodSpendLimit}
		return protoreflect.ValueOfList(value)
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period_can_spend":
		if x.PeriodCanSpend == nil {
			x.PeriodCanSpend = []*v1beta1.Coin{}
		}
		value := &_PeriodicAllowance_4_list{list: &x.PeriodCanSpend}
		return protoreflect.ValueOfList(value)
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period_reset":
		if x.PeriodReset == nil {
			x.PeriodReset = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.PeriodReset.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.PeriodicAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.PeriodicAllowance does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_PeriodicAllowance) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.basic":
		m := new(BasicAllowance)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period":
		m := new(durationpb.Duration)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period_spend_limit":
		list := []*v1beta1.Coin{}
		return protoreflect.ValueOfList(&_PeriodicAllowance_3_list{list: &list})
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period_can_spend":
		list := []*v1beta1.Coin{}
		return protoreflect.ValueOfList(&_PeriodicAllowance_4_list{list: &list})
	case "cosmos.feegrant.v1beta1.PeriodicAllowance.period_reset":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.PeriodicAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.PeriodicAllowance does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_PeriodicAllowance) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.feegrant.v1beta1.PeriodicAllowance", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_PeriodicAllowance) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PeriodicAllowance) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_PeriodicAllowance) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_PeriodicAllowance) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*PeriodicAllowance)
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
		if x.Basic != nil {
			l = options.Size(x.Basic)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Period != nil {
			l = options.Size(x.Period)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.PeriodSpendLimit) > 0 {
			for _, e := range x.PeriodSpendLimit {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.PeriodCanSpend) > 0 {
			for _, e := range x.PeriodCanSpend {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.PeriodReset != nil {
			l = options.Size(x.PeriodReset)
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
		x := input.Message.Interface().(*PeriodicAllowance)
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
		if x.PeriodReset != nil {
			encoded, err := options.Marshal(x.PeriodReset)
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
		if len(x.PeriodCanSpend) > 0 {
			for iNdEx := len(x.PeriodCanSpend) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.PeriodCanSpend[iNdEx])
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
		if len(x.PeriodSpendLimit) > 0 {
			for iNdEx := len(x.PeriodSpendLimit) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.PeriodSpendLimit[iNdEx])
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
		if x.Period != nil {
			encoded, err := options.Marshal(x.Period)
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
		if x.Basic != nil {
			encoded, err := options.Marshal(x.Basic)
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
		x := input.Message.Interface().(*PeriodicAllowance)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PeriodicAllowance: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PeriodicAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Basic", wireType)
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
				if x.Basic == nil {
					x.Basic = &BasicAllowance{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Basic); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
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
				if x.Period == nil {
					x.Period = &durationpb.Duration{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Period); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PeriodSpendLimit", wireType)
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
				x.PeriodSpendLimit = append(x.PeriodSpendLimit, &v1beta1.Coin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.PeriodSpendLimit[len(x.PeriodSpendLimit)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PeriodCanSpend", wireType)
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
				x.PeriodCanSpend = append(x.PeriodCanSpend, &v1beta1.Coin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.PeriodCanSpend[len(x.PeriodCanSpend)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PeriodReset", wireType)
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
				if x.PeriodReset == nil {
					x.PeriodReset = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.PeriodReset); err != nil {
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

var _ protoreflect.List = (*_AllowedMsgAllowance_2_list)(nil)

type _AllowedMsgAllowance_2_list struct {
	list *[]string
}

func (x *_AllowedMsgAllowance_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_AllowedMsgAllowance_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_AllowedMsgAllowance_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_AllowedMsgAllowance_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_AllowedMsgAllowance_2_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message AllowedMsgAllowance at list field AllowedMessages as it is not of Message kind"))
}

func (x *_AllowedMsgAllowance_2_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_AllowedMsgAllowance_2_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_AllowedMsgAllowance_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_AllowedMsgAllowance                  protoreflect.MessageDescriptor
	fd_AllowedMsgAllowance_allowance        protoreflect.FieldDescriptor
	fd_AllowedMsgAllowance_allowed_messages protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_feegrant_v1beta1_feegrant_proto_init()
	md_AllowedMsgAllowance = File_cosmos_feegrant_v1beta1_feegrant_proto.Messages().ByName("AllowedMsgAllowance")
	fd_AllowedMsgAllowance_allowance = md_AllowedMsgAllowance.Fields().ByName("allowance")
	fd_AllowedMsgAllowance_allowed_messages = md_AllowedMsgAllowance.Fields().ByName("allowed_messages")
}

var _ protoreflect.Message = (*fastReflection_AllowedMsgAllowance)(nil)

type fastReflection_AllowedMsgAllowance AllowedMsgAllowance

func (x *AllowedMsgAllowance) ProtoReflect() protoreflect.Message {
	return (*fastReflection_AllowedMsgAllowance)(x)
}

func (x *AllowedMsgAllowance) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_feegrant_v1beta1_feegrant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_AllowedMsgAllowance_messageType fastReflection_AllowedMsgAllowance_messageType
var _ protoreflect.MessageType = fastReflection_AllowedMsgAllowance_messageType{}

type fastReflection_AllowedMsgAllowance_messageType struct{}

func (x fastReflection_AllowedMsgAllowance_messageType) Zero() protoreflect.Message {
	return (*fastReflection_AllowedMsgAllowance)(nil)
}
func (x fastReflection_AllowedMsgAllowance_messageType) New() protoreflect.Message {
	return new(fastReflection_AllowedMsgAllowance)
}
func (x fastReflection_AllowedMsgAllowance_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_AllowedMsgAllowance
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_AllowedMsgAllowance) Descriptor() protoreflect.MessageDescriptor {
	return md_AllowedMsgAllowance
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_AllowedMsgAllowance) Type() protoreflect.MessageType {
	return _fastReflection_AllowedMsgAllowance_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_AllowedMsgAllowance) New() protoreflect.Message {
	return new(fastReflection_AllowedMsgAllowance)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_AllowedMsgAllowance) Interface() protoreflect.ProtoMessage {
	return (*AllowedMsgAllowance)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_AllowedMsgAllowance) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Allowance != nil {
		value := protoreflect.ValueOfMessage(x.Allowance.ProtoReflect())
		if !f(fd_AllowedMsgAllowance_allowance, value) {
			return
		}
	}
	if len(x.AllowedMessages) != 0 {
		value := protoreflect.ValueOfList(&_AllowedMsgAllowance_2_list{list: &x.AllowedMessages})
		if !f(fd_AllowedMsgAllowance_allowed_messages, value) {
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
func (x *fastReflection_AllowedMsgAllowance) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.AllowedMsgAllowance.allowance":
		return x.Allowance != nil
	case "cosmos.feegrant.v1beta1.AllowedMsgAllowance.allowed_messages":
		return len(x.AllowedMessages) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.AllowedMsgAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.AllowedMsgAllowance does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AllowedMsgAllowance) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.AllowedMsgAllowance.allowance":
		x.Allowance = nil
	case "cosmos.feegrant.v1beta1.AllowedMsgAllowance.allowed_messages":
		x.AllowedMessages = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.AllowedMsgAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.AllowedMsgAllowance does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_AllowedMsgAllowance) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.feegrant.v1beta1.AllowedMsgAllowance.allowance":
		value := x.Allowance
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.feegrant.v1beta1.AllowedMsgAllowance.allowed_messages":
		if len(x.AllowedMessages) == 0 {
			return protoreflect.ValueOfList(&_AllowedMsgAllowance_2_list{})
		}
		listValue := &_AllowedMsgAllowance_2_list{list: &x.AllowedMessages}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.AllowedMsgAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.AllowedMsgAllowance does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_AllowedMsgAllowance) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.AllowedMsgAllowance.allowance":
		x.Allowance = value.Message().Interface().(*anypb.Any)
	case "cosmos.feegrant.v1beta1.AllowedMsgAllowance.allowed_messages":
		lv := value.List()
		clv := lv.(*_AllowedMsgAllowance_2_list)
		x.AllowedMessages = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.AllowedMsgAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.AllowedMsgAllowance does not contain field %s", fd.FullName()))
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
func (x *fastReflection_AllowedMsgAllowance) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.AllowedMsgAllowance.allowance":
		if x.Allowance == nil {
			x.Allowance = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.Allowance.ProtoReflect())
	case "cosmos.feegrant.v1beta1.AllowedMsgAllowance.allowed_messages":
		if x.AllowedMessages == nil {
			x.AllowedMessages = []string{}
		}
		value := &_AllowedMsgAllowance_2_list{list: &x.AllowedMessages}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.AllowedMsgAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.AllowedMsgAllowance does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_AllowedMsgAllowance) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.AllowedMsgAllowance.allowance":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.feegrant.v1beta1.AllowedMsgAllowance.allowed_messages":
		list := []string{}
		return protoreflect.ValueOfList(&_AllowedMsgAllowance_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.AllowedMsgAllowance"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.AllowedMsgAllowance does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_AllowedMsgAllowance) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.feegrant.v1beta1.AllowedMsgAllowance", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_AllowedMsgAllowance) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AllowedMsgAllowance) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_AllowedMsgAllowance) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_AllowedMsgAllowance) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*AllowedMsgAllowance)
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
		if x.Allowance != nil {
			l = options.Size(x.Allowance)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.AllowedMessages) > 0 {
			for _, s := range x.AllowedMessages {
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
		x := input.Message.Interface().(*AllowedMsgAllowance)
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
		if len(x.AllowedMessages) > 0 {
			for iNdEx := len(x.AllowedMessages) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.AllowedMessages[iNdEx])
				copy(dAtA[i:], x.AllowedMessages[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.AllowedMessages[iNdEx])))
				i--
				dAtA[i] = 0x12
			}
		}
		if x.Allowance != nil {
			encoded, err := options.Marshal(x.Allowance)
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
		x := input.Message.Interface().(*AllowedMsgAllowance)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AllowedMsgAllowance: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AllowedMsgAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Allowance", wireType)
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
				if x.Allowance == nil {
					x.Allowance = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Allowance); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AllowedMessages", wireType)
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
				x.AllowedMessages = append(x.AllowedMessages, string(dAtA[iNdEx:postIndex]))
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
	md_Grant           protoreflect.MessageDescriptor
	fd_Grant_granter   protoreflect.FieldDescriptor
	fd_Grant_grantee   protoreflect.FieldDescriptor
	fd_Grant_allowance protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_feegrant_v1beta1_feegrant_proto_init()
	md_Grant = File_cosmos_feegrant_v1beta1_feegrant_proto.Messages().ByName("Grant")
	fd_Grant_granter = md_Grant.Fields().ByName("granter")
	fd_Grant_grantee = md_Grant.Fields().ByName("grantee")
	fd_Grant_allowance = md_Grant.Fields().ByName("allowance")
}

var _ protoreflect.Message = (*fastReflection_Grant)(nil)

type fastReflection_Grant Grant

func (x *Grant) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Grant)(x)
}

func (x *Grant) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_feegrant_v1beta1_feegrant_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Grant_messageType fastReflection_Grant_messageType
var _ protoreflect.MessageType = fastReflection_Grant_messageType{}

type fastReflection_Grant_messageType struct{}

func (x fastReflection_Grant_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Grant)(nil)
}
func (x fastReflection_Grant_messageType) New() protoreflect.Message {
	return new(fastReflection_Grant)
}
func (x fastReflection_Grant_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Grant
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Grant) Descriptor() protoreflect.MessageDescriptor {
	return md_Grant
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Grant) Type() protoreflect.MessageType {
	return _fastReflection_Grant_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Grant) New() protoreflect.Message {
	return new(fastReflection_Grant)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Grant) Interface() protoreflect.ProtoMessage {
	return (*Grant)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Grant) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Granter != "" {
		value := protoreflect.ValueOfString(x.Granter)
		if !f(fd_Grant_granter, value) {
			return
		}
	}
	if x.Grantee != "" {
		value := protoreflect.ValueOfString(x.Grantee)
		if !f(fd_Grant_grantee, value) {
			return
		}
	}
	if x.Allowance != nil {
		value := protoreflect.ValueOfMessage(x.Allowance.ProtoReflect())
		if !f(fd_Grant_allowance, value) {
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
func (x *fastReflection_Grant) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.Grant.granter":
		return x.Granter != ""
	case "cosmos.feegrant.v1beta1.Grant.grantee":
		return x.Grantee != ""
	case "cosmos.feegrant.v1beta1.Grant.allowance":
		return x.Allowance != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.Grant"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.Grant does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Grant) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.Grant.granter":
		x.Granter = ""
	case "cosmos.feegrant.v1beta1.Grant.grantee":
		x.Grantee = ""
	case "cosmos.feegrant.v1beta1.Grant.allowance":
		x.Allowance = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.Grant"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.Grant does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Grant) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.feegrant.v1beta1.Grant.granter":
		value := x.Granter
		return protoreflect.ValueOfString(value)
	case "cosmos.feegrant.v1beta1.Grant.grantee":
		value := x.Grantee
		return protoreflect.ValueOfString(value)
	case "cosmos.feegrant.v1beta1.Grant.allowance":
		value := x.Allowance
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.Grant"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.Grant does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Grant) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.Grant.granter":
		x.Granter = value.Interface().(string)
	case "cosmos.feegrant.v1beta1.Grant.grantee":
		x.Grantee = value.Interface().(string)
	case "cosmos.feegrant.v1beta1.Grant.allowance":
		x.Allowance = value.Message().Interface().(*anypb.Any)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.Grant"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.Grant does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Grant) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.Grant.allowance":
		if x.Allowance == nil {
			x.Allowance = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.Allowance.ProtoReflect())
	case "cosmos.feegrant.v1beta1.Grant.granter":
		panic(fmt.Errorf("field granter of message cosmos.feegrant.v1beta1.Grant is not mutable"))
	case "cosmos.feegrant.v1beta1.Grant.grantee":
		panic(fmt.Errorf("field grantee of message cosmos.feegrant.v1beta1.Grant is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.Grant"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.Grant does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Grant) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.feegrant.v1beta1.Grant.granter":
		return protoreflect.ValueOfString("")
	case "cosmos.feegrant.v1beta1.Grant.grantee":
		return protoreflect.ValueOfString("")
	case "cosmos.feegrant.v1beta1.Grant.allowance":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.feegrant.v1beta1.Grant"))
		}
		panic(fmt.Errorf("message cosmos.feegrant.v1beta1.Grant does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Grant) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.feegrant.v1beta1.Grant", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Grant) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Grant) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Grant) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Grant) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Grant)
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
		l = len(x.Granter)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Grantee)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Allowance != nil {
			l = options.Size(x.Allowance)
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
		x := input.Message.Interface().(*Grant)
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
		if x.Allowance != nil {
			encoded, err := options.Marshal(x.Allowance)
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
		if len(x.Grantee) > 0 {
			i -= len(x.Grantee)
			copy(dAtA[i:], x.Grantee)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Grantee)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Granter) > 0 {
			i -= len(x.Granter)
			copy(dAtA[i:], x.Granter)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Granter)))
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
		x := input.Message.Interface().(*Grant)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Grant: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Grant: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
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
				x.Granter = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
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
				x.Grantee = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Allowance", wireType)
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
				if x.Allowance == nil {
					x.Allowance = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Allowance); err != nil {
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

// Since: cosmos-sdk 0.43

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.19.1
// source: cosmos/feegrant/v1beta1/feegrant.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// BasicAllowance implements Allowance with a one-time grant of tokens
// that optionally expires. The grantee can use up to SpendLimit to cover fees.
type BasicAllowance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// spend_limit specifies the maximum amount of tokens that can be spent
	// by this allowance and will be updated as tokens are spent. If it is
	// empty, there is no spend limit and any amount of coins can be spent.
	SpendLimit []*v1beta1.Coin `protobuf:"bytes,1,rep,name=spend_limit,json=spendLimit,proto3" json:"spend_limit,omitempty"`
	// expiration specifies an optional time when this allowance expires
	Expiration *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (x *BasicAllowance) Reset() {
	*x = BasicAllowance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_feegrant_v1beta1_feegrant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicAllowance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicAllowance) ProtoMessage() {}

// Deprecated: Use BasicAllowance.ProtoReflect.Descriptor instead.
func (*BasicAllowance) Descriptor() ([]byte, []int) {
	return file_cosmos_feegrant_v1beta1_feegrant_proto_rawDescGZIP(), []int{0}
}

func (x *BasicAllowance) GetSpendLimit() []*v1beta1.Coin {
	if x != nil {
		return x.SpendLimit
	}
	return nil
}

func (x *BasicAllowance) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

// PeriodicAllowance extends Allowance to allow for both a maximum cap,
// as well as a limit per time period.
type PeriodicAllowance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// basic specifies a struct of `BasicAllowance`
	Basic *BasicAllowance `protobuf:"bytes,1,opt,name=basic,proto3" json:"basic,omitempty"`
	// period specifies the time duration in which period_spend_limit coins can
	// be spent before that allowance is reset
	Period *durationpb.Duration `protobuf:"bytes,2,opt,name=period,proto3" json:"period,omitempty"`
	// period_spend_limit specifies the maximum number of coins that can be spent
	// in the period
	PeriodSpendLimit []*v1beta1.Coin `protobuf:"bytes,3,rep,name=period_spend_limit,json=periodSpendLimit,proto3" json:"period_spend_limit,omitempty"`
	// period_can_spend is the number of coins left to be spent before the period_reset time
	PeriodCanSpend []*v1beta1.Coin `protobuf:"bytes,4,rep,name=period_can_spend,json=periodCanSpend,proto3" json:"period_can_spend,omitempty"`
	// period_reset is the time at which this period resets and a new one begins,
	// it is calculated from the start time of the first transaction after the
	// last period ended
	PeriodReset *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=period_reset,json=periodReset,proto3" json:"period_reset,omitempty"`
}

func (x *PeriodicAllowance) Reset() {
	*x = PeriodicAllowance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_feegrant_v1beta1_feegrant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeriodicAllowance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeriodicAllowance) ProtoMessage() {}

// Deprecated: Use PeriodicAllowance.ProtoReflect.Descriptor instead.
func (*PeriodicAllowance) Descriptor() ([]byte, []int) {
	return file_cosmos_feegrant_v1beta1_feegrant_proto_rawDescGZIP(), []int{1}
}

func (x *PeriodicAllowance) GetBasic() *BasicAllowance {
	if x != nil {
		return x.Basic
	}
	return nil
}

func (x *PeriodicAllowance) GetPeriod() *durationpb.Duration {
	if x != nil {
		return x.Period
	}
	return nil
}

func (x *PeriodicAllowance) GetPeriodSpendLimit() []*v1beta1.Coin {
	if x != nil {
		return x.PeriodSpendLimit
	}
	return nil
}

func (x *PeriodicAllowance) GetPeriodCanSpend() []*v1beta1.Coin {
	if x != nil {
		return x.PeriodCanSpend
	}
	return nil
}

func (x *PeriodicAllowance) GetPeriodReset() *timestamppb.Timestamp {
	if x != nil {
		return x.PeriodReset
	}
	return nil
}

// AllowedMsgAllowance creates allowance only for specified message types.
type AllowedMsgAllowance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// allowance can be any of basic and filtered fee allowance.
	Allowance *anypb.Any `protobuf:"bytes,1,opt,name=allowance,proto3" json:"allowance,omitempty"`
	// allowed_messages are the messages for which the grantee has the access.
	AllowedMessages []string `protobuf:"bytes,2,rep,name=allowed_messages,json=allowedMessages,proto3" json:"allowed_messages,omitempty"`
}

func (x *AllowedMsgAllowance) Reset() {
	*x = AllowedMsgAllowance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_feegrant_v1beta1_feegrant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowedMsgAllowance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowedMsgAllowance) ProtoMessage() {}

// Deprecated: Use AllowedMsgAllowance.ProtoReflect.Descriptor instead.
func (*AllowedMsgAllowance) Descriptor() ([]byte, []int) {
	return file_cosmos_feegrant_v1beta1_feegrant_proto_rawDescGZIP(), []int{2}
}

func (x *AllowedMsgAllowance) GetAllowance() *anypb.Any {
	if x != nil {
		return x.Allowance
	}
	return nil
}

func (x *AllowedMsgAllowance) GetAllowedMessages() []string {
	if x != nil {
		return x.AllowedMessages
	}
	return nil
}

// Grant is stored in the KVStore to record a grant with full context
type Grant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// granter is the address of the user granting an allowance of their funds.
	Granter string `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	// grantee is the address of the user being granted an allowance of another user's funds.
	Grantee string `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// allowance can be any of basic and filtered fee allowance.
	Allowance *anypb.Any `protobuf:"bytes,3,opt,name=allowance,proto3" json:"allowance,omitempty"`
}

func (x *Grant) Reset() {
	*x = Grant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_feegrant_v1beta1_feegrant_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grant) ProtoMessage() {}

// Deprecated: Use Grant.ProtoReflect.Descriptor instead.
func (*Grant) Descriptor() ([]byte, []int) {
	return file_cosmos_feegrant_v1beta1_feegrant_proto_rawDescGZIP(), []int{3}
}

func (x *Grant) GetGranter() string {
	if x != nil {
		return x.Granter
	}
	return ""
}

func (x *Grant) GetGrantee() string {
	if x != nil {
		return x.Grantee
	}
	return ""
}

func (x *Grant) GetAllowance() *anypb.Any {
	if x != nil {
		return x.Allowance
	}
	return nil
}

var File_cosmos_feegrant_v1beta1_feegrant_proto protoreflect.FileDescriptor

var file_cosmos_feegrant_v1beta1_feegrant_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3,
	0x01, 0x0a, 0x0e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x6c, 0x0a, 0x0b, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69,
	0x6e, 0x42, 0x30, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x69, 0x6e, 0x73, 0x52, 0x0a, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x40, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x04, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x3a, 0x11, 0xca, 0xb4, 0x2d, 0x0d, 0x46, 0x65, 0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61,
	0x6e, 0x63, 0x65, 0x49, 0x22, 0xe3, 0x03, 0x0a, 0x11, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69,
	0x63, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x05, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e,
	0x63, 0x65, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x12,
	0x3b, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00,
	0x98, 0xdf, 0x1f, 0x01, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x79, 0x0a, 0x12,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43,
	0x6f, 0x69, 0x6e, 0x42, 0x30, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x10, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x70, 0x65,
	0x6e, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x75, 0x0a, 0x10, 0x70, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x5f, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x30, 0xc8, 0xde,
	0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x0e,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x43, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x47,
	0x0a, 0x0c, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x3a, 0x11, 0xca, 0xb4, 0x2d, 0x0d, 0x46, 0x65, 0x65,
	0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x22, 0x9e, 0x01, 0x0a, 0x13, 0x41,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x4d, 0x73, 0x67, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x45, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x11, 0xca, 0xb4, 0x2d,
	0x0d, 0x46, 0x65, 0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x52, 0x09,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x3a, 0x15, 0x88, 0xa0, 0x1f, 0x00, 0xca, 0xb4, 0x2d, 0x0d, 0x46, 0x65,
	0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x22, 0xb6, 0x01, 0x0a, 0x05,
	0x47, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x07, 0x67, 0x72, 0x61,
	0x6e, 0x74, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x12, 0x45, 0x0a,
	0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x11, 0xca, 0xb4, 0x2d, 0x0d, 0x46, 0x65, 0x65, 0x41,
	0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x52, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x61, 0x6e, 0x63, 0x65, 0x42, 0xf4, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x42, 0x0d, 0x46, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x66,
	0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b,
	0x66, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2,
	0x02, 0x03, 0x43, 0x46, 0x58, 0xaa, 0x02, 0x17, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x46,
	0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca,
	0x02, 0x17, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x46, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x23, 0x43, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x5c, 0x46, 0x65, 0x65, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x19, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x46, 0x65, 0x65, 0x67, 0x72, 0x61,
	0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cosmos_feegrant_v1beta1_feegrant_proto_rawDescOnce sync.Once
	file_cosmos_feegrant_v1beta1_feegrant_proto_rawDescData = file_cosmos_feegrant_v1beta1_feegrant_proto_rawDesc
)

func file_cosmos_feegrant_v1beta1_feegrant_proto_rawDescGZIP() []byte {
	file_cosmos_feegrant_v1beta1_feegrant_proto_rawDescOnce.Do(func() {
		file_cosmos_feegrant_v1beta1_feegrant_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_feegrant_v1beta1_feegrant_proto_rawDescData)
	})
	return file_cosmos_feegrant_v1beta1_feegrant_proto_rawDescData
}

var file_cosmos_feegrant_v1beta1_feegrant_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cosmos_feegrant_v1beta1_feegrant_proto_goTypes = []interface{}{
	(*BasicAllowance)(nil),        // 0: cosmos.feegrant.v1beta1.BasicAllowance
	(*PeriodicAllowance)(nil),     // 1: cosmos.feegrant.v1beta1.PeriodicAllowance
	(*AllowedMsgAllowance)(nil),   // 2: cosmos.feegrant.v1beta1.AllowedMsgAllowance
	(*Grant)(nil),                 // 3: cosmos.feegrant.v1beta1.Grant
	(*v1beta1.Coin)(nil),          // 4: cosmos.base.v1beta1.Coin
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 6: google.protobuf.Duration
	(*anypb.Any)(nil),             // 7: google.protobuf.Any
}
var file_cosmos_feegrant_v1beta1_feegrant_proto_depIdxs = []int32{
	4, // 0: cosmos.feegrant.v1beta1.BasicAllowance.spend_limit:type_name -> cosmos.base.v1beta1.Coin
	5, // 1: cosmos.feegrant.v1beta1.BasicAllowance.expiration:type_name -> google.protobuf.Timestamp
	0, // 2: cosmos.feegrant.v1beta1.PeriodicAllowance.basic:type_name -> cosmos.feegrant.v1beta1.BasicAllowance
	6, // 3: cosmos.feegrant.v1beta1.PeriodicAllowance.period:type_name -> google.protobuf.Duration
	4, // 4: cosmos.feegrant.v1beta1.PeriodicAllowance.period_spend_limit:type_name -> cosmos.base.v1beta1.Coin
	4, // 5: cosmos.feegrant.v1beta1.PeriodicAllowance.period_can_spend:type_name -> cosmos.base.v1beta1.Coin
	5, // 6: cosmos.feegrant.v1beta1.PeriodicAllowance.period_reset:type_name -> google.protobuf.Timestamp
	7, // 7: cosmos.feegrant.v1beta1.AllowedMsgAllowance.allowance:type_name -> google.protobuf.Any
	7, // 8: cosmos.feegrant.v1beta1.Grant.allowance:type_name -> google.protobuf.Any
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_cosmos_feegrant_v1beta1_feegrant_proto_init() }
func file_cosmos_feegrant_v1beta1_feegrant_proto_init() {
	if File_cosmos_feegrant_v1beta1_feegrant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_feegrant_v1beta1_feegrant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicAllowance); i {
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
		file_cosmos_feegrant_v1beta1_feegrant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeriodicAllowance); i {
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
		file_cosmos_feegrant_v1beta1_feegrant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowedMsgAllowance); i {
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
		file_cosmos_feegrant_v1beta1_feegrant_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Grant); i {
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
			RawDescriptor: file_cosmos_feegrant_v1beta1_feegrant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_feegrant_v1beta1_feegrant_proto_goTypes,
		DependencyIndexes: file_cosmos_feegrant_v1beta1_feegrant_proto_depIdxs,
		MessageInfos:      file_cosmos_feegrant_v1beta1_feegrant_proto_msgTypes,
	}.Build()
	File_cosmos_feegrant_v1beta1_feegrant_proto = out.File
	file_cosmos_feegrant_v1beta1_feegrant_proto_rawDesc = nil
	file_cosmos_feegrant_v1beta1_feegrant_proto_goTypes = nil
	file_cosmos_feegrant_v1beta1_feegrant_proto_depIdxs = nil
}
