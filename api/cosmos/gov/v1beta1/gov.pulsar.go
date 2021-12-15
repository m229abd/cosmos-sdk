package govv1beta1

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

var (
	md_WeightedVoteOption        protoreflect.MessageDescriptor
	fd_WeightedVoteOption_option protoreflect.FieldDescriptor
	fd_WeightedVoteOption_weight protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_gov_v1beta1_gov_proto_init()
	md_WeightedVoteOption = File_cosmos_gov_v1beta1_gov_proto.Messages().ByName("WeightedVoteOption")
	fd_WeightedVoteOption_option = md_WeightedVoteOption.Fields().ByName("option")
	fd_WeightedVoteOption_weight = md_WeightedVoteOption.Fields().ByName("weight")
}

var _ protoreflect.Message = (*fastReflection_WeightedVoteOption)(nil)

type fastReflection_WeightedVoteOption WeightedVoteOption

func (x *WeightedVoteOption) ProtoReflect() protoreflect.Message {
	return (*fastReflection_WeightedVoteOption)(x)
}

func (x *WeightedVoteOption) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_gov_v1beta1_gov_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_WeightedVoteOption_messageType fastReflection_WeightedVoteOption_messageType
var _ protoreflect.MessageType = fastReflection_WeightedVoteOption_messageType{}

type fastReflection_WeightedVoteOption_messageType struct{}

func (x fastReflection_WeightedVoteOption_messageType) Zero() protoreflect.Message {
	return (*fastReflection_WeightedVoteOption)(nil)
}
func (x fastReflection_WeightedVoteOption_messageType) New() protoreflect.Message {
	return new(fastReflection_WeightedVoteOption)
}
func (x fastReflection_WeightedVoteOption_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_WeightedVoteOption
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_WeightedVoteOption) Descriptor() protoreflect.MessageDescriptor {
	return md_WeightedVoteOption
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_WeightedVoteOption) Type() protoreflect.MessageType {
	return _fastReflection_WeightedVoteOption_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_WeightedVoteOption) New() protoreflect.Message {
	return new(fastReflection_WeightedVoteOption)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_WeightedVoteOption) Interface() protoreflect.ProtoMessage {
	return (*WeightedVoteOption)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_WeightedVoteOption) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Option != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Option))
		if !f(fd_WeightedVoteOption_option, value) {
			return
		}
	}
	if x.Weight != "" {
		value := protoreflect.ValueOfString(x.Weight)
		if !f(fd_WeightedVoteOption_weight, value) {
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
func (x *fastReflection_WeightedVoteOption) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.WeightedVoteOption.option":
		return x.Option != 0
	case "cosmos.gov.v1beta1.WeightedVoteOption.weight":
		return x.Weight != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.WeightedVoteOption"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.WeightedVoteOption does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_WeightedVoteOption) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.WeightedVoteOption.option":
		x.Option = 0
	case "cosmos.gov.v1beta1.WeightedVoteOption.weight":
		x.Weight = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.WeightedVoteOption"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.WeightedVoteOption does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_WeightedVoteOption) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.gov.v1beta1.WeightedVoteOption.option":
		value := x.Option
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "cosmos.gov.v1beta1.WeightedVoteOption.weight":
		value := x.Weight
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.WeightedVoteOption"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.WeightedVoteOption does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_WeightedVoteOption) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.WeightedVoteOption.option":
		x.Option = (VoteOption)(value.Enum())
	case "cosmos.gov.v1beta1.WeightedVoteOption.weight":
		x.Weight = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.WeightedVoteOption"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.WeightedVoteOption does not contain field %s", fd.FullName()))
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
func (x *fastReflection_WeightedVoteOption) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.WeightedVoteOption.option":
		panic(fmt.Errorf("field option of message cosmos.gov.v1beta1.WeightedVoteOption is not mutable"))
	case "cosmos.gov.v1beta1.WeightedVoteOption.weight":
		panic(fmt.Errorf("field weight of message cosmos.gov.v1beta1.WeightedVoteOption is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.WeightedVoteOption"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.WeightedVoteOption does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_WeightedVoteOption) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.WeightedVoteOption.option":
		return protoreflect.ValueOfEnum(0)
	case "cosmos.gov.v1beta1.WeightedVoteOption.weight":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.WeightedVoteOption"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.WeightedVoteOption does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_WeightedVoteOption) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.gov.v1beta1.WeightedVoteOption", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_WeightedVoteOption) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_WeightedVoteOption) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_WeightedVoteOption) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_WeightedVoteOption) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*WeightedVoteOption)
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
		if x.Option != 0 {
			n += 1 + runtime.Sov(uint64(x.Option))
		}
		l = len(x.Weight)
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
		x := input.Message.Interface().(*WeightedVoteOption)
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
		if len(x.Weight) > 0 {
			i -= len(x.Weight)
			copy(dAtA[i:], x.Weight)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Weight)))
			i--
			dAtA[i] = 0x12
		}
		if x.Option != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Option))
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
		x := input.Message.Interface().(*WeightedVoteOption)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: WeightedVoteOption: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: WeightedVoteOption: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Option", wireType)
				}
				x.Option = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Option |= VoteOption(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
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
				x.Weight = string(dAtA[iNdEx:postIndex])
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
	md_TextProposal             protoreflect.MessageDescriptor
	fd_TextProposal_title       protoreflect.FieldDescriptor
	fd_TextProposal_description protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_gov_v1beta1_gov_proto_init()
	md_TextProposal = File_cosmos_gov_v1beta1_gov_proto.Messages().ByName("TextProposal")
	fd_TextProposal_title = md_TextProposal.Fields().ByName("title")
	fd_TextProposal_description = md_TextProposal.Fields().ByName("description")
}

var _ protoreflect.Message = (*fastReflection_TextProposal)(nil)

type fastReflection_TextProposal TextProposal

func (x *TextProposal) ProtoReflect() protoreflect.Message {
	return (*fastReflection_TextProposal)(x)
}

func (x *TextProposal) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_gov_v1beta1_gov_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_TextProposal_messageType fastReflection_TextProposal_messageType
var _ protoreflect.MessageType = fastReflection_TextProposal_messageType{}

type fastReflection_TextProposal_messageType struct{}

func (x fastReflection_TextProposal_messageType) Zero() protoreflect.Message {
	return (*fastReflection_TextProposal)(nil)
}
func (x fastReflection_TextProposal_messageType) New() protoreflect.Message {
	return new(fastReflection_TextProposal)
}
func (x fastReflection_TextProposal_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_TextProposal
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_TextProposal) Descriptor() protoreflect.MessageDescriptor {
	return md_TextProposal
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_TextProposal) Type() protoreflect.MessageType {
	return _fastReflection_TextProposal_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_TextProposal) New() protoreflect.Message {
	return new(fastReflection_TextProposal)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_TextProposal) Interface() protoreflect.ProtoMessage {
	return (*TextProposal)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_TextProposal) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Title != "" {
		value := protoreflect.ValueOfString(x.Title)
		if !f(fd_TextProposal_title, value) {
			return
		}
	}
	if x.Description != "" {
		value := protoreflect.ValueOfString(x.Description)
		if !f(fd_TextProposal_description, value) {
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
func (x *fastReflection_TextProposal) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.TextProposal.title":
		return x.Title != ""
	case "cosmos.gov.v1beta1.TextProposal.description":
		return x.Description != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.TextProposal"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.TextProposal does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TextProposal) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.TextProposal.title":
		x.Title = ""
	case "cosmos.gov.v1beta1.TextProposal.description":
		x.Description = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.TextProposal"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.TextProposal does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_TextProposal) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.gov.v1beta1.TextProposal.title":
		value := x.Title
		return protoreflect.ValueOfString(value)
	case "cosmos.gov.v1beta1.TextProposal.description":
		value := x.Description
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.TextProposal"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.TextProposal does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_TextProposal) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.TextProposal.title":
		x.Title = value.Interface().(string)
	case "cosmos.gov.v1beta1.TextProposal.description":
		x.Description = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.TextProposal"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.TextProposal does not contain field %s", fd.FullName()))
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
func (x *fastReflection_TextProposal) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.TextProposal.title":
		panic(fmt.Errorf("field title of message cosmos.gov.v1beta1.TextProposal is not mutable"))
	case "cosmos.gov.v1beta1.TextProposal.description":
		panic(fmt.Errorf("field description of message cosmos.gov.v1beta1.TextProposal is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.TextProposal"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.TextProposal does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_TextProposal) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.TextProposal.title":
		return protoreflect.ValueOfString("")
	case "cosmos.gov.v1beta1.TextProposal.description":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.TextProposal"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.TextProposal does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_TextProposal) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.gov.v1beta1.TextProposal", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_TextProposal) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TextProposal) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_TextProposal) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_TextProposal) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*TextProposal)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*TextProposal)
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
		x := input.Message.Interface().(*TextProposal)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TextProposal: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TextProposal: illegal tag %d (wire type %d)", fieldNum, wire)
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

var _ protoreflect.List = (*_Deposit_3_list)(nil)

type _Deposit_3_list struct {
	list *[]*v1beta1.Coin
}

func (x *_Deposit_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Deposit_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_Deposit_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	(*x.list)[i] = concreteValue
}

func (x *_Deposit_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_Deposit_3_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.Coin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Deposit_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_Deposit_3_list) NewElement() protoreflect.Value {
	v := new(v1beta1.Coin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Deposit_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Deposit             protoreflect.MessageDescriptor
	fd_Deposit_proposal_id protoreflect.FieldDescriptor
	fd_Deposit_depositor   protoreflect.FieldDescriptor
	fd_Deposit_amount      protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_gov_v1beta1_gov_proto_init()
	md_Deposit = File_cosmos_gov_v1beta1_gov_proto.Messages().ByName("Deposit")
	fd_Deposit_proposal_id = md_Deposit.Fields().ByName("proposal_id")
	fd_Deposit_depositor = md_Deposit.Fields().ByName("depositor")
	fd_Deposit_amount = md_Deposit.Fields().ByName("amount")
}

var _ protoreflect.Message = (*fastReflection_Deposit)(nil)

type fastReflection_Deposit Deposit

func (x *Deposit) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Deposit)(x)
}

func (x *Deposit) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_gov_v1beta1_gov_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Deposit_messageType fastReflection_Deposit_messageType
var _ protoreflect.MessageType = fastReflection_Deposit_messageType{}

type fastReflection_Deposit_messageType struct{}

func (x fastReflection_Deposit_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Deposit)(nil)
}
func (x fastReflection_Deposit_messageType) New() protoreflect.Message {
	return new(fastReflection_Deposit)
}
func (x fastReflection_Deposit_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Deposit
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Deposit) Descriptor() protoreflect.MessageDescriptor {
	return md_Deposit
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Deposit) Type() protoreflect.MessageType {
	return _fastReflection_Deposit_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Deposit) New() protoreflect.Message {
	return new(fastReflection_Deposit)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Deposit) Interface() protoreflect.ProtoMessage {
	return (*Deposit)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Deposit) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ProposalId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ProposalId)
		if !f(fd_Deposit_proposal_id, value) {
			return
		}
	}
	if x.Depositor != "" {
		value := protoreflect.ValueOfString(x.Depositor)
		if !f(fd_Deposit_depositor, value) {
			return
		}
	}
	if len(x.Amount) != 0 {
		value := protoreflect.ValueOfList(&_Deposit_3_list{list: &x.Amount})
		if !f(fd_Deposit_amount, value) {
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
func (x *fastReflection_Deposit) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.Deposit.proposal_id":
		return x.ProposalId != uint64(0)
	case "cosmos.gov.v1beta1.Deposit.depositor":
		return x.Depositor != ""
	case "cosmos.gov.v1beta1.Deposit.amount":
		return len(x.Amount) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.Deposit"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.Deposit does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Deposit) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.Deposit.proposal_id":
		x.ProposalId = uint64(0)
	case "cosmos.gov.v1beta1.Deposit.depositor":
		x.Depositor = ""
	case "cosmos.gov.v1beta1.Deposit.amount":
		x.Amount = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.Deposit"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.Deposit does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Deposit) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.gov.v1beta1.Deposit.proposal_id":
		value := x.ProposalId
		return protoreflect.ValueOfUint64(value)
	case "cosmos.gov.v1beta1.Deposit.depositor":
		value := x.Depositor
		return protoreflect.ValueOfString(value)
	case "cosmos.gov.v1beta1.Deposit.amount":
		if len(x.Amount) == 0 {
			return protoreflect.ValueOfList(&_Deposit_3_list{})
		}
		listValue := &_Deposit_3_list{list: &x.Amount}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.Deposit"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.Deposit does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Deposit) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.Deposit.proposal_id":
		x.ProposalId = value.Uint()
	case "cosmos.gov.v1beta1.Deposit.depositor":
		x.Depositor = value.Interface().(string)
	case "cosmos.gov.v1beta1.Deposit.amount":
		lv := value.List()
		clv := lv.(*_Deposit_3_list)
		x.Amount = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.Deposit"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.Deposit does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Deposit) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.Deposit.amount":
		if x.Amount == nil {
			x.Amount = []*v1beta1.Coin{}
		}
		value := &_Deposit_3_list{list: &x.Amount}
		return protoreflect.ValueOfList(value)
	case "cosmos.gov.v1beta1.Deposit.proposal_id":
		panic(fmt.Errorf("field proposal_id of message cosmos.gov.v1beta1.Deposit is not mutable"))
	case "cosmos.gov.v1beta1.Deposit.depositor":
		panic(fmt.Errorf("field depositor of message cosmos.gov.v1beta1.Deposit is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.Deposit"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.Deposit does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Deposit) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.Deposit.proposal_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.gov.v1beta1.Deposit.depositor":
		return protoreflect.ValueOfString("")
	case "cosmos.gov.v1beta1.Deposit.amount":
		list := []*v1beta1.Coin{}
		return protoreflect.ValueOfList(&_Deposit_3_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.Deposit"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.Deposit does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Deposit) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.gov.v1beta1.Deposit", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Deposit) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Deposit) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Deposit) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Deposit) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Deposit)
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
		if x.ProposalId != 0 {
			n += 1 + runtime.Sov(uint64(x.ProposalId))
		}
		l = len(x.Depositor)
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
		x := input.Message.Interface().(*Deposit)
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
				dAtA[i] = 0x1a
			}
		}
		if len(x.Depositor) > 0 {
			i -= len(x.Depositor)
			copy(dAtA[i:], x.Depositor)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Depositor)))
			i--
			dAtA[i] = 0x12
		}
		if x.ProposalId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ProposalId))
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
		x := input.Message.Interface().(*Deposit)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Deposit: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Deposit: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
				}
				x.ProposalId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ProposalId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
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
				x.Depositor = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
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

var _ protoreflect.List = (*_Proposal_7_list)(nil)

type _Proposal_7_list struct {
	list *[]*v1beta1.Coin
}

func (x *_Proposal_7_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Proposal_7_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_Proposal_7_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	(*x.list)[i] = concreteValue
}

func (x *_Proposal_7_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_Proposal_7_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.Coin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Proposal_7_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_Proposal_7_list) NewElement() protoreflect.Value {
	v := new(v1beta1.Coin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Proposal_7_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Proposal                    protoreflect.MessageDescriptor
	fd_Proposal_proposal_id        protoreflect.FieldDescriptor
	fd_Proposal_content            protoreflect.FieldDescriptor
	fd_Proposal_status             protoreflect.FieldDescriptor
	fd_Proposal_final_tally_result protoreflect.FieldDescriptor
	fd_Proposal_submit_time        protoreflect.FieldDescriptor
	fd_Proposal_deposit_end_time   protoreflect.FieldDescriptor
	fd_Proposal_total_deposit      protoreflect.FieldDescriptor
	fd_Proposal_voting_start_time  protoreflect.FieldDescriptor
	fd_Proposal_voting_end_time    protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_gov_v1beta1_gov_proto_init()
	md_Proposal = File_cosmos_gov_v1beta1_gov_proto.Messages().ByName("Proposal")
	fd_Proposal_proposal_id = md_Proposal.Fields().ByName("proposal_id")
	fd_Proposal_content = md_Proposal.Fields().ByName("content")
	fd_Proposal_status = md_Proposal.Fields().ByName("status")
	fd_Proposal_final_tally_result = md_Proposal.Fields().ByName("final_tally_result")
	fd_Proposal_submit_time = md_Proposal.Fields().ByName("submit_time")
	fd_Proposal_deposit_end_time = md_Proposal.Fields().ByName("deposit_end_time")
	fd_Proposal_total_deposit = md_Proposal.Fields().ByName("total_deposit")
	fd_Proposal_voting_start_time = md_Proposal.Fields().ByName("voting_start_time")
	fd_Proposal_voting_end_time = md_Proposal.Fields().ByName("voting_end_time")
}

var _ protoreflect.Message = (*fastReflection_Proposal)(nil)

type fastReflection_Proposal Proposal

func (x *Proposal) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Proposal)(x)
}

func (x *Proposal) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_gov_v1beta1_gov_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Proposal_messageType fastReflection_Proposal_messageType
var _ protoreflect.MessageType = fastReflection_Proposal_messageType{}

type fastReflection_Proposal_messageType struct{}

func (x fastReflection_Proposal_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Proposal)(nil)
}
func (x fastReflection_Proposal_messageType) New() protoreflect.Message {
	return new(fastReflection_Proposal)
}
func (x fastReflection_Proposal_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Proposal
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Proposal) Descriptor() protoreflect.MessageDescriptor {
	return md_Proposal
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Proposal) Type() protoreflect.MessageType {
	return _fastReflection_Proposal_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Proposal) New() protoreflect.Message {
	return new(fastReflection_Proposal)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Proposal) Interface() protoreflect.ProtoMessage {
	return (*Proposal)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Proposal) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ProposalId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ProposalId)
		if !f(fd_Proposal_proposal_id, value) {
			return
		}
	}
	if x.Content != nil {
		value := protoreflect.ValueOfMessage(x.Content.ProtoReflect())
		if !f(fd_Proposal_content, value) {
			return
		}
	}
	if x.Status != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Status))
		if !f(fd_Proposal_status, value) {
			return
		}
	}
	if x.FinalTallyResult != nil {
		value := protoreflect.ValueOfMessage(x.FinalTallyResult.ProtoReflect())
		if !f(fd_Proposal_final_tally_result, value) {
			return
		}
	}
	if x.SubmitTime != nil {
		value := protoreflect.ValueOfMessage(x.SubmitTime.ProtoReflect())
		if !f(fd_Proposal_submit_time, value) {
			return
		}
	}
	if x.DepositEndTime != nil {
		value := protoreflect.ValueOfMessage(x.DepositEndTime.ProtoReflect())
		if !f(fd_Proposal_deposit_end_time, value) {
			return
		}
	}
	if len(x.TotalDeposit) != 0 {
		value := protoreflect.ValueOfList(&_Proposal_7_list{list: &x.TotalDeposit})
		if !f(fd_Proposal_total_deposit, value) {
			return
		}
	}
	if x.VotingStartTime != nil {
		value := protoreflect.ValueOfMessage(x.VotingStartTime.ProtoReflect())
		if !f(fd_Proposal_voting_start_time, value) {
			return
		}
	}
	if x.VotingEndTime != nil {
		value := protoreflect.ValueOfMessage(x.VotingEndTime.ProtoReflect())
		if !f(fd_Proposal_voting_end_time, value) {
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
func (x *fastReflection_Proposal) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.Proposal.proposal_id":
		return x.ProposalId != uint64(0)
	case "cosmos.gov.v1beta1.Proposal.content":
		return x.Content != nil
	case "cosmos.gov.v1beta1.Proposal.status":
		return x.Status != 0
	case "cosmos.gov.v1beta1.Proposal.final_tally_result":
		return x.FinalTallyResult != nil
	case "cosmos.gov.v1beta1.Proposal.submit_time":
		return x.SubmitTime != nil
	case "cosmos.gov.v1beta1.Proposal.deposit_end_time":
		return x.DepositEndTime != nil
	case "cosmos.gov.v1beta1.Proposal.total_deposit":
		return len(x.TotalDeposit) != 0
	case "cosmos.gov.v1beta1.Proposal.voting_start_time":
		return x.VotingStartTime != nil
	case "cosmos.gov.v1beta1.Proposal.voting_end_time":
		return x.VotingEndTime != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.Proposal"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.Proposal does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Proposal) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.Proposal.proposal_id":
		x.ProposalId = uint64(0)
	case "cosmos.gov.v1beta1.Proposal.content":
		x.Content = nil
	case "cosmos.gov.v1beta1.Proposal.status":
		x.Status = 0
	case "cosmos.gov.v1beta1.Proposal.final_tally_result":
		x.FinalTallyResult = nil
	case "cosmos.gov.v1beta1.Proposal.submit_time":
		x.SubmitTime = nil
	case "cosmos.gov.v1beta1.Proposal.deposit_end_time":
		x.DepositEndTime = nil
	case "cosmos.gov.v1beta1.Proposal.total_deposit":
		x.TotalDeposit = nil
	case "cosmos.gov.v1beta1.Proposal.voting_start_time":
		x.VotingStartTime = nil
	case "cosmos.gov.v1beta1.Proposal.voting_end_time":
		x.VotingEndTime = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.Proposal"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.Proposal does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Proposal) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.gov.v1beta1.Proposal.proposal_id":
		value := x.ProposalId
		return protoreflect.ValueOfUint64(value)
	case "cosmos.gov.v1beta1.Proposal.content":
		value := x.Content
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.gov.v1beta1.Proposal.status":
		value := x.Status
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "cosmos.gov.v1beta1.Proposal.final_tally_result":
		value := x.FinalTallyResult
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.gov.v1beta1.Proposal.submit_time":
		value := x.SubmitTime
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.gov.v1beta1.Proposal.deposit_end_time":
		value := x.DepositEndTime
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.gov.v1beta1.Proposal.total_deposit":
		if len(x.TotalDeposit) == 0 {
			return protoreflect.ValueOfList(&_Proposal_7_list{})
		}
		listValue := &_Proposal_7_list{list: &x.TotalDeposit}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.gov.v1beta1.Proposal.voting_start_time":
		value := x.VotingStartTime
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.gov.v1beta1.Proposal.voting_end_time":
		value := x.VotingEndTime
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.Proposal"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.Proposal does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Proposal) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.Proposal.proposal_id":
		x.ProposalId = value.Uint()
	case "cosmos.gov.v1beta1.Proposal.content":
		x.Content = value.Message().Interface().(*anypb.Any)
	case "cosmos.gov.v1beta1.Proposal.status":
		x.Status = (ProposalStatus)(value.Enum())
	case "cosmos.gov.v1beta1.Proposal.final_tally_result":
		x.FinalTallyResult = value.Message().Interface().(*TallyResult)
	case "cosmos.gov.v1beta1.Proposal.submit_time":
		x.SubmitTime = value.Message().Interface().(*timestamppb.Timestamp)
	case "cosmos.gov.v1beta1.Proposal.deposit_end_time":
		x.DepositEndTime = value.Message().Interface().(*timestamppb.Timestamp)
	case "cosmos.gov.v1beta1.Proposal.total_deposit":
		lv := value.List()
		clv := lv.(*_Proposal_7_list)
		x.TotalDeposit = *clv.list
	case "cosmos.gov.v1beta1.Proposal.voting_start_time":
		x.VotingStartTime = value.Message().Interface().(*timestamppb.Timestamp)
	case "cosmos.gov.v1beta1.Proposal.voting_end_time":
		x.VotingEndTime = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.Proposal"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.Proposal does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Proposal) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.Proposal.content":
		if x.Content == nil {
			x.Content = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.Content.ProtoReflect())
	case "cosmos.gov.v1beta1.Proposal.final_tally_result":
		if x.FinalTallyResult == nil {
			x.FinalTallyResult = new(TallyResult)
		}
		return protoreflect.ValueOfMessage(x.FinalTallyResult.ProtoReflect())
	case "cosmos.gov.v1beta1.Proposal.submit_time":
		if x.SubmitTime == nil {
			x.SubmitTime = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.SubmitTime.ProtoReflect())
	case "cosmos.gov.v1beta1.Proposal.deposit_end_time":
		if x.DepositEndTime == nil {
			x.DepositEndTime = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.DepositEndTime.ProtoReflect())
	case "cosmos.gov.v1beta1.Proposal.total_deposit":
		if x.TotalDeposit == nil {
			x.TotalDeposit = []*v1beta1.Coin{}
		}
		value := &_Proposal_7_list{list: &x.TotalDeposit}
		return protoreflect.ValueOfList(value)
	case "cosmos.gov.v1beta1.Proposal.voting_start_time":
		if x.VotingStartTime == nil {
			x.VotingStartTime = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.VotingStartTime.ProtoReflect())
	case "cosmos.gov.v1beta1.Proposal.voting_end_time":
		if x.VotingEndTime == nil {
			x.VotingEndTime = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.VotingEndTime.ProtoReflect())
	case "cosmos.gov.v1beta1.Proposal.proposal_id":
		panic(fmt.Errorf("field proposal_id of message cosmos.gov.v1beta1.Proposal is not mutable"))
	case "cosmos.gov.v1beta1.Proposal.status":
		panic(fmt.Errorf("field status of message cosmos.gov.v1beta1.Proposal is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.Proposal"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.Proposal does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Proposal) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.Proposal.proposal_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.gov.v1beta1.Proposal.content":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.gov.v1beta1.Proposal.status":
		return protoreflect.ValueOfEnum(0)
	case "cosmos.gov.v1beta1.Proposal.final_tally_result":
		m := new(TallyResult)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.gov.v1beta1.Proposal.submit_time":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.gov.v1beta1.Proposal.deposit_end_time":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.gov.v1beta1.Proposal.total_deposit":
		list := []*v1beta1.Coin{}
		return protoreflect.ValueOfList(&_Proposal_7_list{list: &list})
	case "cosmos.gov.v1beta1.Proposal.voting_start_time":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.gov.v1beta1.Proposal.voting_end_time":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.Proposal"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.Proposal does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Proposal) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.gov.v1beta1.Proposal", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Proposal) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Proposal) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Proposal) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Proposal) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Proposal)
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
		if x.ProposalId != 0 {
			n += 1 + runtime.Sov(uint64(x.ProposalId))
		}
		if x.Content != nil {
			l = options.Size(x.Content)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Status != 0 {
			n += 1 + runtime.Sov(uint64(x.Status))
		}
		if x.FinalTallyResult != nil {
			l = options.Size(x.FinalTallyResult)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.SubmitTime != nil {
			l = options.Size(x.SubmitTime)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.DepositEndTime != nil {
			l = options.Size(x.DepositEndTime)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.TotalDeposit) > 0 {
			for _, e := range x.TotalDeposit {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.VotingStartTime != nil {
			l = options.Size(x.VotingStartTime)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.VotingEndTime != nil {
			l = options.Size(x.VotingEndTime)
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
		x := input.Message.Interface().(*Proposal)
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
		if x.VotingEndTime != nil {
			encoded, err := options.Marshal(x.VotingEndTime)
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
		if x.VotingStartTime != nil {
			encoded, err := options.Marshal(x.VotingStartTime)
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
		if len(x.TotalDeposit) > 0 {
			for iNdEx := len(x.TotalDeposit) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.TotalDeposit[iNdEx])
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
		if x.DepositEndTime != nil {
			encoded, err := options.Marshal(x.DepositEndTime)
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
		if x.SubmitTime != nil {
			encoded, err := options.Marshal(x.SubmitTime)
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
		if x.FinalTallyResult != nil {
			encoded, err := options.Marshal(x.FinalTallyResult)
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
		if x.Status != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Status))
			i--
			dAtA[i] = 0x18
		}
		if x.Content != nil {
			encoded, err := options.Marshal(x.Content)
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
		if x.ProposalId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ProposalId))
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
		x := input.Message.Interface().(*Proposal)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Proposal: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Proposal: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
				}
				x.ProposalId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ProposalId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
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
				if x.Content == nil {
					x.Content = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Content); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
				}
				x.Status = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Status |= ProposalStatus(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field FinalTallyResult", wireType)
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
				if x.FinalTallyResult == nil {
					x.FinalTallyResult = &TallyResult{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.FinalTallyResult); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SubmitTime", wireType)
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
				if x.SubmitTime == nil {
					x.SubmitTime = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.SubmitTime); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DepositEndTime", wireType)
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
				if x.DepositEndTime == nil {
					x.DepositEndTime = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.DepositEndTime); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TotalDeposit", wireType)
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
				x.TotalDeposit = append(x.TotalDeposit, &v1beta1.Coin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.TotalDeposit[len(x.TotalDeposit)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 8:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field VotingStartTime", wireType)
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
				if x.VotingStartTime == nil {
					x.VotingStartTime = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.VotingStartTime); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 9:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field VotingEndTime", wireType)
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
				if x.VotingEndTime == nil {
					x.VotingEndTime = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.VotingEndTime); err != nil {
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
	md_TallyResult              protoreflect.MessageDescriptor
	fd_TallyResult_yes          protoreflect.FieldDescriptor
	fd_TallyResult_abstain      protoreflect.FieldDescriptor
	fd_TallyResult_no           protoreflect.FieldDescriptor
	fd_TallyResult_no_with_veto protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_gov_v1beta1_gov_proto_init()
	md_TallyResult = File_cosmos_gov_v1beta1_gov_proto.Messages().ByName("TallyResult")
	fd_TallyResult_yes = md_TallyResult.Fields().ByName("yes")
	fd_TallyResult_abstain = md_TallyResult.Fields().ByName("abstain")
	fd_TallyResult_no = md_TallyResult.Fields().ByName("no")
	fd_TallyResult_no_with_veto = md_TallyResult.Fields().ByName("no_with_veto")
}

var _ protoreflect.Message = (*fastReflection_TallyResult)(nil)

type fastReflection_TallyResult TallyResult

func (x *TallyResult) ProtoReflect() protoreflect.Message {
	return (*fastReflection_TallyResult)(x)
}

func (x *TallyResult) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_gov_v1beta1_gov_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_TallyResult_messageType fastReflection_TallyResult_messageType
var _ protoreflect.MessageType = fastReflection_TallyResult_messageType{}

type fastReflection_TallyResult_messageType struct{}

func (x fastReflection_TallyResult_messageType) Zero() protoreflect.Message {
	return (*fastReflection_TallyResult)(nil)
}
func (x fastReflection_TallyResult_messageType) New() protoreflect.Message {
	return new(fastReflection_TallyResult)
}
func (x fastReflection_TallyResult_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_TallyResult
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_TallyResult) Descriptor() protoreflect.MessageDescriptor {
	return md_TallyResult
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_TallyResult) Type() protoreflect.MessageType {
	return _fastReflection_TallyResult_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_TallyResult) New() protoreflect.Message {
	return new(fastReflection_TallyResult)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_TallyResult) Interface() protoreflect.ProtoMessage {
	return (*TallyResult)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_TallyResult) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Yes != "" {
		value := protoreflect.ValueOfString(x.Yes)
		if !f(fd_TallyResult_yes, value) {
			return
		}
	}
	if x.Abstain != "" {
		value := protoreflect.ValueOfString(x.Abstain)
		if !f(fd_TallyResult_abstain, value) {
			return
		}
	}
	if x.No != "" {
		value := protoreflect.ValueOfString(x.No)
		if !f(fd_TallyResult_no, value) {
			return
		}
	}
	if x.NoWithVeto != "" {
		value := protoreflect.ValueOfString(x.NoWithVeto)
		if !f(fd_TallyResult_no_with_veto, value) {
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
func (x *fastReflection_TallyResult) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.TallyResult.yes":
		return x.Yes != ""
	case "cosmos.gov.v1beta1.TallyResult.abstain":
		return x.Abstain != ""
	case "cosmos.gov.v1beta1.TallyResult.no":
		return x.No != ""
	case "cosmos.gov.v1beta1.TallyResult.no_with_veto":
		return x.NoWithVeto != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.TallyResult"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.TallyResult does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TallyResult) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.TallyResult.yes":
		x.Yes = ""
	case "cosmos.gov.v1beta1.TallyResult.abstain":
		x.Abstain = ""
	case "cosmos.gov.v1beta1.TallyResult.no":
		x.No = ""
	case "cosmos.gov.v1beta1.TallyResult.no_with_veto":
		x.NoWithVeto = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.TallyResult"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.TallyResult does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_TallyResult) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.gov.v1beta1.TallyResult.yes":
		value := x.Yes
		return protoreflect.ValueOfString(value)
	case "cosmos.gov.v1beta1.TallyResult.abstain":
		value := x.Abstain
		return protoreflect.ValueOfString(value)
	case "cosmos.gov.v1beta1.TallyResult.no":
		value := x.No
		return protoreflect.ValueOfString(value)
	case "cosmos.gov.v1beta1.TallyResult.no_with_veto":
		value := x.NoWithVeto
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.TallyResult"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.TallyResult does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_TallyResult) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.TallyResult.yes":
		x.Yes = value.Interface().(string)
	case "cosmos.gov.v1beta1.TallyResult.abstain":
		x.Abstain = value.Interface().(string)
	case "cosmos.gov.v1beta1.TallyResult.no":
		x.No = value.Interface().(string)
	case "cosmos.gov.v1beta1.TallyResult.no_with_veto":
		x.NoWithVeto = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.TallyResult"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.TallyResult does not contain field %s", fd.FullName()))
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
func (x *fastReflection_TallyResult) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.TallyResult.yes":
		panic(fmt.Errorf("field yes of message cosmos.gov.v1beta1.TallyResult is not mutable"))
	case "cosmos.gov.v1beta1.TallyResult.abstain":
		panic(fmt.Errorf("field abstain of message cosmos.gov.v1beta1.TallyResult is not mutable"))
	case "cosmos.gov.v1beta1.TallyResult.no":
		panic(fmt.Errorf("field no of message cosmos.gov.v1beta1.TallyResult is not mutable"))
	case "cosmos.gov.v1beta1.TallyResult.no_with_veto":
		panic(fmt.Errorf("field no_with_veto of message cosmos.gov.v1beta1.TallyResult is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.TallyResult"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.TallyResult does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_TallyResult) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.TallyResult.yes":
		return protoreflect.ValueOfString("")
	case "cosmos.gov.v1beta1.TallyResult.abstain":
		return protoreflect.ValueOfString("")
	case "cosmos.gov.v1beta1.TallyResult.no":
		return protoreflect.ValueOfString("")
	case "cosmos.gov.v1beta1.TallyResult.no_with_veto":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.TallyResult"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.TallyResult does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_TallyResult) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.gov.v1beta1.TallyResult", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_TallyResult) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TallyResult) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_TallyResult) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_TallyResult) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*TallyResult)
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
		l = len(x.Yes)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Abstain)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.No)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.NoWithVeto)
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
		x := input.Message.Interface().(*TallyResult)
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
		if len(x.NoWithVeto) > 0 {
			i -= len(x.NoWithVeto)
			copy(dAtA[i:], x.NoWithVeto)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.NoWithVeto)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.No) > 0 {
			i -= len(x.No)
			copy(dAtA[i:], x.No)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.No)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Abstain) > 0 {
			i -= len(x.Abstain)
			copy(dAtA[i:], x.Abstain)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Abstain)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Yes) > 0 {
			i -= len(x.Yes)
			copy(dAtA[i:], x.Yes)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Yes)))
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
		x := input.Message.Interface().(*TallyResult)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TallyResult: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TallyResult: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Yes", wireType)
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
				x.Yes = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Abstain", wireType)
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
				x.Abstain = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field No", wireType)
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
				x.No = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NoWithVeto", wireType)
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
				x.NoWithVeto = string(dAtA[iNdEx:postIndex])
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

var _ protoreflect.List = (*_Vote_4_list)(nil)

type _Vote_4_list struct {
	list *[]*WeightedVoteOption
}

func (x *_Vote_4_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Vote_4_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_Vote_4_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*WeightedVoteOption)
	(*x.list)[i] = concreteValue
}

func (x *_Vote_4_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*WeightedVoteOption)
	*x.list = append(*x.list, concreteValue)
}

func (x *_Vote_4_list) AppendMutable() protoreflect.Value {
	v := new(WeightedVoteOption)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Vote_4_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_Vote_4_list) NewElement() protoreflect.Value {
	v := new(WeightedVoteOption)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Vote_4_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Vote             protoreflect.MessageDescriptor
	fd_Vote_proposal_id protoreflect.FieldDescriptor
	fd_Vote_voter       protoreflect.FieldDescriptor
	fd_Vote_option      protoreflect.FieldDescriptor
	fd_Vote_options     protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_gov_v1beta1_gov_proto_init()
	md_Vote = File_cosmos_gov_v1beta1_gov_proto.Messages().ByName("Vote")
	fd_Vote_proposal_id = md_Vote.Fields().ByName("proposal_id")
	fd_Vote_voter = md_Vote.Fields().ByName("voter")
	fd_Vote_option = md_Vote.Fields().ByName("option")
	fd_Vote_options = md_Vote.Fields().ByName("options")
}

var _ protoreflect.Message = (*fastReflection_Vote)(nil)

type fastReflection_Vote Vote

func (x *Vote) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Vote)(x)
}

func (x *Vote) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_gov_v1beta1_gov_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Vote_messageType fastReflection_Vote_messageType
var _ protoreflect.MessageType = fastReflection_Vote_messageType{}

type fastReflection_Vote_messageType struct{}

func (x fastReflection_Vote_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Vote)(nil)
}
func (x fastReflection_Vote_messageType) New() protoreflect.Message {
	return new(fastReflection_Vote)
}
func (x fastReflection_Vote_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Vote
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Vote) Descriptor() protoreflect.MessageDescriptor {
	return md_Vote
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Vote) Type() protoreflect.MessageType {
	return _fastReflection_Vote_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Vote) New() protoreflect.Message {
	return new(fastReflection_Vote)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Vote) Interface() protoreflect.ProtoMessage {
	return (*Vote)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Vote) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ProposalId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ProposalId)
		if !f(fd_Vote_proposal_id, value) {
			return
		}
	}
	if x.Voter != "" {
		value := protoreflect.ValueOfString(x.Voter)
		if !f(fd_Vote_voter, value) {
			return
		}
	}
	if x.Option != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Option))
		if !f(fd_Vote_option, value) {
			return
		}
	}
	if len(x.Options) != 0 {
		value := protoreflect.ValueOfList(&_Vote_4_list{list: &x.Options})
		if !f(fd_Vote_options, value) {
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
func (x *fastReflection_Vote) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.Vote.proposal_id":
		return x.ProposalId != uint64(0)
	case "cosmos.gov.v1beta1.Vote.voter":
		return x.Voter != ""
	case "cosmos.gov.v1beta1.Vote.option":
		return x.Option != 0
	case "cosmos.gov.v1beta1.Vote.options":
		return len(x.Options) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.Vote"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.Vote does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Vote) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.Vote.proposal_id":
		x.ProposalId = uint64(0)
	case "cosmos.gov.v1beta1.Vote.voter":
		x.Voter = ""
	case "cosmos.gov.v1beta1.Vote.option":
		x.Option = 0
	case "cosmos.gov.v1beta1.Vote.options":
		x.Options = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.Vote"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.Vote does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Vote) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.gov.v1beta1.Vote.proposal_id":
		value := x.ProposalId
		return protoreflect.ValueOfUint64(value)
	case "cosmos.gov.v1beta1.Vote.voter":
		value := x.Voter
		return protoreflect.ValueOfString(value)
	case "cosmos.gov.v1beta1.Vote.option":
		value := x.Option
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "cosmos.gov.v1beta1.Vote.options":
		if len(x.Options) == 0 {
			return protoreflect.ValueOfList(&_Vote_4_list{})
		}
		listValue := &_Vote_4_list{list: &x.Options}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.Vote"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.Vote does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Vote) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.Vote.proposal_id":
		x.ProposalId = value.Uint()
	case "cosmos.gov.v1beta1.Vote.voter":
		x.Voter = value.Interface().(string)
	case "cosmos.gov.v1beta1.Vote.option":
		x.Option = (VoteOption)(value.Enum())
	case "cosmos.gov.v1beta1.Vote.options":
		lv := value.List()
		clv := lv.(*_Vote_4_list)
		x.Options = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.Vote"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.Vote does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Vote) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.Vote.options":
		if x.Options == nil {
			x.Options = []*WeightedVoteOption{}
		}
		value := &_Vote_4_list{list: &x.Options}
		return protoreflect.ValueOfList(value)
	case "cosmos.gov.v1beta1.Vote.proposal_id":
		panic(fmt.Errorf("field proposal_id of message cosmos.gov.v1beta1.Vote is not mutable"))
	case "cosmos.gov.v1beta1.Vote.voter":
		panic(fmt.Errorf("field voter of message cosmos.gov.v1beta1.Vote is not mutable"))
	case "cosmos.gov.v1beta1.Vote.option":
		panic(fmt.Errorf("field option of message cosmos.gov.v1beta1.Vote is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.Vote"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.Vote does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Vote) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.Vote.proposal_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.gov.v1beta1.Vote.voter":
		return protoreflect.ValueOfString("")
	case "cosmos.gov.v1beta1.Vote.option":
		return protoreflect.ValueOfEnum(0)
	case "cosmos.gov.v1beta1.Vote.options":
		list := []*WeightedVoteOption{}
		return protoreflect.ValueOfList(&_Vote_4_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.Vote"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.Vote does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Vote) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.gov.v1beta1.Vote", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Vote) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Vote) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Vote) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Vote) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Vote)
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
		if x.ProposalId != 0 {
			n += 1 + runtime.Sov(uint64(x.ProposalId))
		}
		l = len(x.Voter)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Option != 0 {
			n += 1 + runtime.Sov(uint64(x.Option))
		}
		if len(x.Options) > 0 {
			for _, e := range x.Options {
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
		x := input.Message.Interface().(*Vote)
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
		if len(x.Options) > 0 {
			for iNdEx := len(x.Options) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Options[iNdEx])
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
		if x.Option != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Option))
			i--
			dAtA[i] = 0x18
		}
		if len(x.Voter) > 0 {
			i -= len(x.Voter)
			copy(dAtA[i:], x.Voter)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Voter)))
			i--
			dAtA[i] = 0x12
		}
		if x.ProposalId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ProposalId))
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
		x := input.Message.Interface().(*Vote)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Vote: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Vote: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
				}
				x.ProposalId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ProposalId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
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
				x.Voter = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Option", wireType)
				}
				x.Option = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Option |= VoteOption(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
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
				x.Options = append(x.Options, &WeightedVoteOption{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Options[len(x.Options)-1]); err != nil {
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

var _ protoreflect.List = (*_DepositParams_1_list)(nil)

type _DepositParams_1_list struct {
	list *[]*v1beta1.Coin
}

func (x *_DepositParams_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_DepositParams_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_DepositParams_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	(*x.list)[i] = concreteValue
}

func (x *_DepositParams_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_DepositParams_1_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.Coin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_DepositParams_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_DepositParams_1_list) NewElement() protoreflect.Value {
	v := new(v1beta1.Coin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_DepositParams_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_DepositParams                    protoreflect.MessageDescriptor
	fd_DepositParams_min_deposit        protoreflect.FieldDescriptor
	fd_DepositParams_max_deposit_period protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_gov_v1beta1_gov_proto_init()
	md_DepositParams = File_cosmos_gov_v1beta1_gov_proto.Messages().ByName("DepositParams")
	fd_DepositParams_min_deposit = md_DepositParams.Fields().ByName("min_deposit")
	fd_DepositParams_max_deposit_period = md_DepositParams.Fields().ByName("max_deposit_period")
}

var _ protoreflect.Message = (*fastReflection_DepositParams)(nil)

type fastReflection_DepositParams DepositParams

func (x *DepositParams) ProtoReflect() protoreflect.Message {
	return (*fastReflection_DepositParams)(x)
}

func (x *DepositParams) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_gov_v1beta1_gov_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_DepositParams_messageType fastReflection_DepositParams_messageType
var _ protoreflect.MessageType = fastReflection_DepositParams_messageType{}

type fastReflection_DepositParams_messageType struct{}

func (x fastReflection_DepositParams_messageType) Zero() protoreflect.Message {
	return (*fastReflection_DepositParams)(nil)
}
func (x fastReflection_DepositParams_messageType) New() protoreflect.Message {
	return new(fastReflection_DepositParams)
}
func (x fastReflection_DepositParams_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_DepositParams
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_DepositParams) Descriptor() protoreflect.MessageDescriptor {
	return md_DepositParams
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_DepositParams) Type() protoreflect.MessageType {
	return _fastReflection_DepositParams_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_DepositParams) New() protoreflect.Message {
	return new(fastReflection_DepositParams)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_DepositParams) Interface() protoreflect.ProtoMessage {
	return (*DepositParams)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_DepositParams) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.MinDeposit) != 0 {
		value := protoreflect.ValueOfList(&_DepositParams_1_list{list: &x.MinDeposit})
		if !f(fd_DepositParams_min_deposit, value) {
			return
		}
	}
	if x.MaxDepositPeriod != nil {
		value := protoreflect.ValueOfMessage(x.MaxDepositPeriod.ProtoReflect())
		if !f(fd_DepositParams_max_deposit_period, value) {
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
func (x *fastReflection_DepositParams) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.DepositParams.min_deposit":
		return len(x.MinDeposit) != 0
	case "cosmos.gov.v1beta1.DepositParams.max_deposit_period":
		return x.MaxDepositPeriod != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.DepositParams"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.DepositParams does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DepositParams) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.DepositParams.min_deposit":
		x.MinDeposit = nil
	case "cosmos.gov.v1beta1.DepositParams.max_deposit_period":
		x.MaxDepositPeriod = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.DepositParams"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.DepositParams does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_DepositParams) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.gov.v1beta1.DepositParams.min_deposit":
		if len(x.MinDeposit) == 0 {
			return protoreflect.ValueOfList(&_DepositParams_1_list{})
		}
		listValue := &_DepositParams_1_list{list: &x.MinDeposit}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.gov.v1beta1.DepositParams.max_deposit_period":
		value := x.MaxDepositPeriod
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.DepositParams"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.DepositParams does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_DepositParams) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.DepositParams.min_deposit":
		lv := value.List()
		clv := lv.(*_DepositParams_1_list)
		x.MinDeposit = *clv.list
	case "cosmos.gov.v1beta1.DepositParams.max_deposit_period":
		x.MaxDepositPeriod = value.Message().Interface().(*durationpb.Duration)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.DepositParams"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.DepositParams does not contain field %s", fd.FullName()))
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
func (x *fastReflection_DepositParams) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.DepositParams.min_deposit":
		if x.MinDeposit == nil {
			x.MinDeposit = []*v1beta1.Coin{}
		}
		value := &_DepositParams_1_list{list: &x.MinDeposit}
		return protoreflect.ValueOfList(value)
	case "cosmos.gov.v1beta1.DepositParams.max_deposit_period":
		if x.MaxDepositPeriod == nil {
			x.MaxDepositPeriod = new(durationpb.Duration)
		}
		return protoreflect.ValueOfMessage(x.MaxDepositPeriod.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.DepositParams"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.DepositParams does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_DepositParams) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.DepositParams.min_deposit":
		list := []*v1beta1.Coin{}
		return protoreflect.ValueOfList(&_DepositParams_1_list{list: &list})
	case "cosmos.gov.v1beta1.DepositParams.max_deposit_period":
		m := new(durationpb.Duration)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.DepositParams"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.DepositParams does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_DepositParams) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.gov.v1beta1.DepositParams", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_DepositParams) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DepositParams) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_DepositParams) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_DepositParams) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*DepositParams)
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
		if len(x.MinDeposit) > 0 {
			for _, e := range x.MinDeposit {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.MaxDepositPeriod != nil {
			l = options.Size(x.MaxDepositPeriod)
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
		x := input.Message.Interface().(*DepositParams)
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
		if x.MaxDepositPeriod != nil {
			encoded, err := options.Marshal(x.MaxDepositPeriod)
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
		if len(x.MinDeposit) > 0 {
			for iNdEx := len(x.MinDeposit) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.MinDeposit[iNdEx])
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
		x := input.Message.Interface().(*DepositParams)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DepositParams: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DepositParams: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MinDeposit", wireType)
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
				x.MinDeposit = append(x.MinDeposit, &v1beta1.Coin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.MinDeposit[len(x.MinDeposit)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MaxDepositPeriod", wireType)
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
				if x.MaxDepositPeriod == nil {
					x.MaxDepositPeriod = &durationpb.Duration{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.MaxDepositPeriod); err != nil {
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
	md_VotingParams               protoreflect.MessageDescriptor
	fd_VotingParams_voting_period protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_gov_v1beta1_gov_proto_init()
	md_VotingParams = File_cosmos_gov_v1beta1_gov_proto.Messages().ByName("VotingParams")
	fd_VotingParams_voting_period = md_VotingParams.Fields().ByName("voting_period")
}

var _ protoreflect.Message = (*fastReflection_VotingParams)(nil)

type fastReflection_VotingParams VotingParams

func (x *VotingParams) ProtoReflect() protoreflect.Message {
	return (*fastReflection_VotingParams)(x)
}

func (x *VotingParams) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_gov_v1beta1_gov_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_VotingParams_messageType fastReflection_VotingParams_messageType
var _ protoreflect.MessageType = fastReflection_VotingParams_messageType{}

type fastReflection_VotingParams_messageType struct{}

func (x fastReflection_VotingParams_messageType) Zero() protoreflect.Message {
	return (*fastReflection_VotingParams)(nil)
}
func (x fastReflection_VotingParams_messageType) New() protoreflect.Message {
	return new(fastReflection_VotingParams)
}
func (x fastReflection_VotingParams_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_VotingParams
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_VotingParams) Descriptor() protoreflect.MessageDescriptor {
	return md_VotingParams
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_VotingParams) Type() protoreflect.MessageType {
	return _fastReflection_VotingParams_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_VotingParams) New() protoreflect.Message {
	return new(fastReflection_VotingParams)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_VotingParams) Interface() protoreflect.ProtoMessage {
	return (*VotingParams)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_VotingParams) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.VotingPeriod != nil {
		value := protoreflect.ValueOfMessage(x.VotingPeriod.ProtoReflect())
		if !f(fd_VotingParams_voting_period, value) {
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
func (x *fastReflection_VotingParams) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.VotingParams.voting_period":
		return x.VotingPeriod != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.VotingParams"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.VotingParams does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_VotingParams) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.VotingParams.voting_period":
		x.VotingPeriod = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.VotingParams"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.VotingParams does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_VotingParams) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.gov.v1beta1.VotingParams.voting_period":
		value := x.VotingPeriod
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.VotingParams"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.VotingParams does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_VotingParams) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.VotingParams.voting_period":
		x.VotingPeriod = value.Message().Interface().(*durationpb.Duration)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.VotingParams"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.VotingParams does not contain field %s", fd.FullName()))
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
func (x *fastReflection_VotingParams) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.VotingParams.voting_period":
		if x.VotingPeriod == nil {
			x.VotingPeriod = new(durationpb.Duration)
		}
		return protoreflect.ValueOfMessage(x.VotingPeriod.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.VotingParams"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.VotingParams does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_VotingParams) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.VotingParams.voting_period":
		m := new(durationpb.Duration)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.VotingParams"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.VotingParams does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_VotingParams) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.gov.v1beta1.VotingParams", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_VotingParams) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_VotingParams) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_VotingParams) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_VotingParams) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*VotingParams)
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
		if x.VotingPeriod != nil {
			l = options.Size(x.VotingPeriod)
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
		x := input.Message.Interface().(*VotingParams)
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
		if x.VotingPeriod != nil {
			encoded, err := options.Marshal(x.VotingPeriod)
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
		x := input.Message.Interface().(*VotingParams)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: VotingParams: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: VotingParams: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field VotingPeriod", wireType)
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
				if x.VotingPeriod == nil {
					x.VotingPeriod = &durationpb.Duration{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.VotingPeriod); err != nil {
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
	md_TallyParams                protoreflect.MessageDescriptor
	fd_TallyParams_quorum         protoreflect.FieldDescriptor
	fd_TallyParams_threshold      protoreflect.FieldDescriptor
	fd_TallyParams_veto_threshold protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_gov_v1beta1_gov_proto_init()
	md_TallyParams = File_cosmos_gov_v1beta1_gov_proto.Messages().ByName("TallyParams")
	fd_TallyParams_quorum = md_TallyParams.Fields().ByName("quorum")
	fd_TallyParams_threshold = md_TallyParams.Fields().ByName("threshold")
	fd_TallyParams_veto_threshold = md_TallyParams.Fields().ByName("veto_threshold")
}

var _ protoreflect.Message = (*fastReflection_TallyParams)(nil)

type fastReflection_TallyParams TallyParams

func (x *TallyParams) ProtoReflect() protoreflect.Message {
	return (*fastReflection_TallyParams)(x)
}

func (x *TallyParams) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_gov_v1beta1_gov_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_TallyParams_messageType fastReflection_TallyParams_messageType
var _ protoreflect.MessageType = fastReflection_TallyParams_messageType{}

type fastReflection_TallyParams_messageType struct{}

func (x fastReflection_TallyParams_messageType) Zero() protoreflect.Message {
	return (*fastReflection_TallyParams)(nil)
}
func (x fastReflection_TallyParams_messageType) New() protoreflect.Message {
	return new(fastReflection_TallyParams)
}
func (x fastReflection_TallyParams_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_TallyParams
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_TallyParams) Descriptor() protoreflect.MessageDescriptor {
	return md_TallyParams
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_TallyParams) Type() protoreflect.MessageType {
	return _fastReflection_TallyParams_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_TallyParams) New() protoreflect.Message {
	return new(fastReflection_TallyParams)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_TallyParams) Interface() protoreflect.ProtoMessage {
	return (*TallyParams)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_TallyParams) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Quorum) != 0 {
		value := protoreflect.ValueOfBytes(x.Quorum)
		if !f(fd_TallyParams_quorum, value) {
			return
		}
	}
	if len(x.Threshold) != 0 {
		value := protoreflect.ValueOfBytes(x.Threshold)
		if !f(fd_TallyParams_threshold, value) {
			return
		}
	}
	if len(x.VetoThreshold) != 0 {
		value := protoreflect.ValueOfBytes(x.VetoThreshold)
		if !f(fd_TallyParams_veto_threshold, value) {
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
func (x *fastReflection_TallyParams) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.TallyParams.quorum":
		return len(x.Quorum) != 0
	case "cosmos.gov.v1beta1.TallyParams.threshold":
		return len(x.Threshold) != 0
	case "cosmos.gov.v1beta1.TallyParams.veto_threshold":
		return len(x.VetoThreshold) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.TallyParams"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.TallyParams does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TallyParams) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.TallyParams.quorum":
		x.Quorum = nil
	case "cosmos.gov.v1beta1.TallyParams.threshold":
		x.Threshold = nil
	case "cosmos.gov.v1beta1.TallyParams.veto_threshold":
		x.VetoThreshold = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.TallyParams"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.TallyParams does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_TallyParams) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.gov.v1beta1.TallyParams.quorum":
		value := x.Quorum
		return protoreflect.ValueOfBytes(value)
	case "cosmos.gov.v1beta1.TallyParams.threshold":
		value := x.Threshold
		return protoreflect.ValueOfBytes(value)
	case "cosmos.gov.v1beta1.TallyParams.veto_threshold":
		value := x.VetoThreshold
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.TallyParams"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.TallyParams does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_TallyParams) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.TallyParams.quorum":
		x.Quorum = value.Bytes()
	case "cosmos.gov.v1beta1.TallyParams.threshold":
		x.Threshold = value.Bytes()
	case "cosmos.gov.v1beta1.TallyParams.veto_threshold":
		x.VetoThreshold = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.TallyParams"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.TallyParams does not contain field %s", fd.FullName()))
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
func (x *fastReflection_TallyParams) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.TallyParams.quorum":
		panic(fmt.Errorf("field quorum of message cosmos.gov.v1beta1.TallyParams is not mutable"))
	case "cosmos.gov.v1beta1.TallyParams.threshold":
		panic(fmt.Errorf("field threshold of message cosmos.gov.v1beta1.TallyParams is not mutable"))
	case "cosmos.gov.v1beta1.TallyParams.veto_threshold":
		panic(fmt.Errorf("field veto_threshold of message cosmos.gov.v1beta1.TallyParams is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.TallyParams"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.TallyParams does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_TallyParams) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.gov.v1beta1.TallyParams.quorum":
		return protoreflect.ValueOfBytes(nil)
	case "cosmos.gov.v1beta1.TallyParams.threshold":
		return protoreflect.ValueOfBytes(nil)
	case "cosmos.gov.v1beta1.TallyParams.veto_threshold":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.gov.v1beta1.TallyParams"))
		}
		panic(fmt.Errorf("message cosmos.gov.v1beta1.TallyParams does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_TallyParams) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.gov.v1beta1.TallyParams", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_TallyParams) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TallyParams) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_TallyParams) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_TallyParams) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*TallyParams)
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
		l = len(x.Quorum)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Threshold)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.VetoThreshold)
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
		x := input.Message.Interface().(*TallyParams)
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
		if len(x.VetoThreshold) > 0 {
			i -= len(x.VetoThreshold)
			copy(dAtA[i:], x.VetoThreshold)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.VetoThreshold)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Threshold) > 0 {
			i -= len(x.Threshold)
			copy(dAtA[i:], x.Threshold)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Threshold)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Quorum) > 0 {
			i -= len(x.Quorum)
			copy(dAtA[i:], x.Quorum)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Quorum)))
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
		x := input.Message.Interface().(*TallyParams)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TallyParams: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TallyParams: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Quorum", wireType)
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
				x.Quorum = append(x.Quorum[:0], dAtA[iNdEx:postIndex]...)
				if x.Quorum == nil {
					x.Quorum = []byte{}
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
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
				x.Threshold = append(x.Threshold[:0], dAtA[iNdEx:postIndex]...)
				if x.Threshold == nil {
					x.Threshold = []byte{}
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field VetoThreshold", wireType)
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
				x.VetoThreshold = append(x.VetoThreshold[:0], dAtA[iNdEx:postIndex]...)
				if x.VetoThreshold == nil {
					x.VetoThreshold = []byte{}
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
// source: cosmos/gov/v1beta1/gov.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// VoteOption enumerates the valid vote options for a given governance proposal.
type VoteOption int32

const (
	// VOTE_OPTION_UNSPECIFIED defines a no-op vote option.
	VoteOption_VOTE_OPTION_UNSPECIFIED VoteOption = 0
	// VOTE_OPTION_YES defines a yes vote option.
	VoteOption_VOTE_OPTION_YES VoteOption = 1
	// VOTE_OPTION_ABSTAIN defines an abstain vote option.
	VoteOption_VOTE_OPTION_ABSTAIN VoteOption = 2
	// VOTE_OPTION_NO defines a no vote option.
	VoteOption_VOTE_OPTION_NO VoteOption = 3
	// VOTE_OPTION_NO_WITH_VETO defines a no with veto vote option.
	VoteOption_VOTE_OPTION_NO_WITH_VETO VoteOption = 4
)

// Enum value maps for VoteOption.
var (
	VoteOption_name = map[int32]string{
		0: "VOTE_OPTION_UNSPECIFIED",
		1: "VOTE_OPTION_YES",
		2: "VOTE_OPTION_ABSTAIN",
		3: "VOTE_OPTION_NO",
		4: "VOTE_OPTION_NO_WITH_VETO",
	}
	VoteOption_value = map[string]int32{
		"VOTE_OPTION_UNSPECIFIED":  0,
		"VOTE_OPTION_YES":          1,
		"VOTE_OPTION_ABSTAIN":      2,
		"VOTE_OPTION_NO":           3,
		"VOTE_OPTION_NO_WITH_VETO": 4,
	}
)

func (x VoteOption) Enum() *VoteOption {
	p := new(VoteOption)
	*p = x
	return p
}

func (x VoteOption) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VoteOption) Descriptor() protoreflect.EnumDescriptor {
	return file_cosmos_gov_v1beta1_gov_proto_enumTypes[0].Descriptor()
}

func (VoteOption) Type() protoreflect.EnumType {
	return &file_cosmos_gov_v1beta1_gov_proto_enumTypes[0]
}

func (x VoteOption) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VoteOption.Descriptor instead.
func (VoteOption) EnumDescriptor() ([]byte, []int) {
	return file_cosmos_gov_v1beta1_gov_proto_rawDescGZIP(), []int{0}
}

// ProposalStatus enumerates the valid statuses of a proposal.
type ProposalStatus int32

const (
	// PROPOSAL_STATUS_UNSPECIFIED defines the default propopsal status.
	ProposalStatus_PROPOSAL_STATUS_UNSPECIFIED ProposalStatus = 0
	// PROPOSAL_STATUS_DEPOSIT_PERIOD defines a proposal status during the deposit
	// period.
	ProposalStatus_PROPOSAL_STATUS_DEPOSIT_PERIOD ProposalStatus = 1
	// PROPOSAL_STATUS_VOTING_PERIOD defines a proposal status during the voting
	// period.
	ProposalStatus_PROPOSAL_STATUS_VOTING_PERIOD ProposalStatus = 2
	// PROPOSAL_STATUS_PASSED defines a proposal status of a proposal that has
	// passed.
	ProposalStatus_PROPOSAL_STATUS_PASSED ProposalStatus = 3
	// PROPOSAL_STATUS_REJECTED defines a proposal status of a proposal that has
	// been rejected.
	ProposalStatus_PROPOSAL_STATUS_REJECTED ProposalStatus = 4
	// PROPOSAL_STATUS_FAILED defines a proposal status of a proposal that has
	// failed.
	ProposalStatus_PROPOSAL_STATUS_FAILED ProposalStatus = 5
)

// Enum value maps for ProposalStatus.
var (
	ProposalStatus_name = map[int32]string{
		0: "PROPOSAL_STATUS_UNSPECIFIED",
		1: "PROPOSAL_STATUS_DEPOSIT_PERIOD",
		2: "PROPOSAL_STATUS_VOTING_PERIOD",
		3: "PROPOSAL_STATUS_PASSED",
		4: "PROPOSAL_STATUS_REJECTED",
		5: "PROPOSAL_STATUS_FAILED",
	}
	ProposalStatus_value = map[string]int32{
		"PROPOSAL_STATUS_UNSPECIFIED":    0,
		"PROPOSAL_STATUS_DEPOSIT_PERIOD": 1,
		"PROPOSAL_STATUS_VOTING_PERIOD":  2,
		"PROPOSAL_STATUS_PASSED":         3,
		"PROPOSAL_STATUS_REJECTED":       4,
		"PROPOSAL_STATUS_FAILED":         5,
	}
)

func (x ProposalStatus) Enum() *ProposalStatus {
	p := new(ProposalStatus)
	*p = x
	return p
}

func (x ProposalStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProposalStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_cosmos_gov_v1beta1_gov_proto_enumTypes[1].Descriptor()
}

func (ProposalStatus) Type() protoreflect.EnumType {
	return &file_cosmos_gov_v1beta1_gov_proto_enumTypes[1]
}

func (x ProposalStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProposalStatus.Descriptor instead.
func (ProposalStatus) EnumDescriptor() ([]byte, []int) {
	return file_cosmos_gov_v1beta1_gov_proto_rawDescGZIP(), []int{1}
}

// WeightedVoteOption defines a unit of vote for vote split.
//
// Since: cosmos-sdk 0.43
type WeightedVoteOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Option VoteOption `protobuf:"varint,1,opt,name=option,proto3,enum=cosmos.gov.v1beta1.VoteOption" json:"option,omitempty"`
	Weight string     `protobuf:"bytes,2,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *WeightedVoteOption) Reset() {
	*x = WeightedVoteOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_gov_v1beta1_gov_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeightedVoteOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeightedVoteOption) ProtoMessage() {}

// Deprecated: Use WeightedVoteOption.ProtoReflect.Descriptor instead.
func (*WeightedVoteOption) Descriptor() ([]byte, []int) {
	return file_cosmos_gov_v1beta1_gov_proto_rawDescGZIP(), []int{0}
}

func (x *WeightedVoteOption) GetOption() VoteOption {
	if x != nil {
		return x.Option
	}
	return VoteOption_VOTE_OPTION_UNSPECIFIED
}

func (x *WeightedVoteOption) GetWeight() string {
	if x != nil {
		return x.Weight
	}
	return ""
}

// TextProposal defines a standard text proposal whose changes need to be
// manually updated in case of approval.
type TextProposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *TextProposal) Reset() {
	*x = TextProposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_gov_v1beta1_gov_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextProposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextProposal) ProtoMessage() {}

// Deprecated: Use TextProposal.ProtoReflect.Descriptor instead.
func (*TextProposal) Descriptor() ([]byte, []int) {
	return file_cosmos_gov_v1beta1_gov_proto_rawDescGZIP(), []int{1}
}

func (x *TextProposal) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TextProposal) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Deposit defines an amount deposited by an account address to an active
// proposal.
type Deposit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProposalId uint64          `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	Depositor  string          `protobuf:"bytes,2,opt,name=depositor,proto3" json:"depositor,omitempty"`
	Amount     []*v1beta1.Coin `protobuf:"bytes,3,rep,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Deposit) Reset() {
	*x = Deposit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_gov_v1beta1_gov_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deposit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deposit) ProtoMessage() {}

// Deprecated: Use Deposit.ProtoReflect.Descriptor instead.
func (*Deposit) Descriptor() ([]byte, []int) {
	return file_cosmos_gov_v1beta1_gov_proto_rawDescGZIP(), []int{2}
}

func (x *Deposit) GetProposalId() uint64 {
	if x != nil {
		return x.ProposalId
	}
	return 0
}

func (x *Deposit) GetDepositor() string {
	if x != nil {
		return x.Depositor
	}
	return ""
}

func (x *Deposit) GetAmount() []*v1beta1.Coin {
	if x != nil {
		return x.Amount
	}
	return nil
}

// Proposal defines the core field members of a governance proposal.
type Proposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProposalId       uint64                 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	Content          *anypb.Any             `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Status           ProposalStatus         `protobuf:"varint,3,opt,name=status,proto3,enum=cosmos.gov.v1beta1.ProposalStatus" json:"status,omitempty"`
	FinalTallyResult *TallyResult           `protobuf:"bytes,4,opt,name=final_tally_result,json=finalTallyResult,proto3" json:"final_tally_result,omitempty"`
	SubmitTime       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=submit_time,json=submitTime,proto3" json:"submit_time,omitempty"`
	DepositEndTime   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=deposit_end_time,json=depositEndTime,proto3" json:"deposit_end_time,omitempty"`
	TotalDeposit     []*v1beta1.Coin        `protobuf:"bytes,7,rep,name=total_deposit,json=totalDeposit,proto3" json:"total_deposit,omitempty"`
	VotingStartTime  *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=voting_start_time,json=votingStartTime,proto3" json:"voting_start_time,omitempty"`
	VotingEndTime    *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=voting_end_time,json=votingEndTime,proto3" json:"voting_end_time,omitempty"`
}

func (x *Proposal) Reset() {
	*x = Proposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_gov_v1beta1_gov_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proposal) ProtoMessage() {}

// Deprecated: Use Proposal.ProtoReflect.Descriptor instead.
func (*Proposal) Descriptor() ([]byte, []int) {
	return file_cosmos_gov_v1beta1_gov_proto_rawDescGZIP(), []int{3}
}

func (x *Proposal) GetProposalId() uint64 {
	if x != nil {
		return x.ProposalId
	}
	return 0
}

func (x *Proposal) GetContent() *anypb.Any {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Proposal) GetStatus() ProposalStatus {
	if x != nil {
		return x.Status
	}
	return ProposalStatus_PROPOSAL_STATUS_UNSPECIFIED
}

func (x *Proposal) GetFinalTallyResult() *TallyResult {
	if x != nil {
		return x.FinalTallyResult
	}
	return nil
}

func (x *Proposal) GetSubmitTime() *timestamppb.Timestamp {
	if x != nil {
		return x.SubmitTime
	}
	return nil
}

func (x *Proposal) GetDepositEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DepositEndTime
	}
	return nil
}

func (x *Proposal) GetTotalDeposit() []*v1beta1.Coin {
	if x != nil {
		return x.TotalDeposit
	}
	return nil
}

func (x *Proposal) GetVotingStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.VotingStartTime
	}
	return nil
}

func (x *Proposal) GetVotingEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.VotingEndTime
	}
	return nil
}

// TallyResult defines a standard tally for a governance proposal.
type TallyResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Yes        string `protobuf:"bytes,1,opt,name=yes,proto3" json:"yes,omitempty"`
	Abstain    string `protobuf:"bytes,2,opt,name=abstain,proto3" json:"abstain,omitempty"`
	No         string `protobuf:"bytes,3,opt,name=no,proto3" json:"no,omitempty"`
	NoWithVeto string `protobuf:"bytes,4,opt,name=no_with_veto,json=noWithVeto,proto3" json:"no_with_veto,omitempty"`
}

func (x *TallyResult) Reset() {
	*x = TallyResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_gov_v1beta1_gov_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TallyResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TallyResult) ProtoMessage() {}

// Deprecated: Use TallyResult.ProtoReflect.Descriptor instead.
func (*TallyResult) Descriptor() ([]byte, []int) {
	return file_cosmos_gov_v1beta1_gov_proto_rawDescGZIP(), []int{4}
}

func (x *TallyResult) GetYes() string {
	if x != nil {
		return x.Yes
	}
	return ""
}

func (x *TallyResult) GetAbstain() string {
	if x != nil {
		return x.Abstain
	}
	return ""
}

func (x *TallyResult) GetNo() string {
	if x != nil {
		return x.No
	}
	return ""
}

func (x *TallyResult) GetNoWithVeto() string {
	if x != nil {
		return x.NoWithVeto
	}
	return ""
}

// Vote defines a vote on a governance proposal.
// A Vote consists of a proposal ID, the voter, and the vote option.
type Vote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	Voter      string `protobuf:"bytes,2,opt,name=voter,proto3" json:"voter,omitempty"`
	// Deprecated: Prefer to use `options` instead. This field is set in queries
	// if and only if `len(options) == 1` and that option has weight 1. In all
	// other cases, this field will default to VOTE_OPTION_UNSPECIFIED.
	//
	// Deprecated: Do not use.
	Option VoteOption `protobuf:"varint,3,opt,name=option,proto3,enum=cosmos.gov.v1beta1.VoteOption" json:"option,omitempty"`
	// Since: cosmos-sdk 0.43
	Options []*WeightedVoteOption `protobuf:"bytes,4,rep,name=options,proto3" json:"options,omitempty"`
}

func (x *Vote) Reset() {
	*x = Vote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_gov_v1beta1_gov_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vote) ProtoMessage() {}

// Deprecated: Use Vote.ProtoReflect.Descriptor instead.
func (*Vote) Descriptor() ([]byte, []int) {
	return file_cosmos_gov_v1beta1_gov_proto_rawDescGZIP(), []int{5}
}

func (x *Vote) GetProposalId() uint64 {
	if x != nil {
		return x.ProposalId
	}
	return 0
}

func (x *Vote) GetVoter() string {
	if x != nil {
		return x.Voter
	}
	return ""
}

// Deprecated: Do not use.
func (x *Vote) GetOption() VoteOption {
	if x != nil {
		return x.Option
	}
	return VoteOption_VOTE_OPTION_UNSPECIFIED
}

func (x *Vote) GetOptions() []*WeightedVoteOption {
	if x != nil {
		return x.Options
	}
	return nil
}

// DepositParams defines the params for deposits on governance proposals.
type DepositParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  Minimum deposit for a proposal to enter voting period.
	MinDeposit []*v1beta1.Coin `protobuf:"bytes,1,rep,name=min_deposit,json=minDeposit,proto3" json:"min_deposit,omitempty"`
	//  Maximum period for Atom holders to deposit on a proposal. Initial value: 2
	//  months.
	MaxDepositPeriod *durationpb.Duration `protobuf:"bytes,2,opt,name=max_deposit_period,json=maxDepositPeriod,proto3" json:"max_deposit_period,omitempty"`
}

func (x *DepositParams) Reset() {
	*x = DepositParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_gov_v1beta1_gov_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositParams) ProtoMessage() {}

// Deprecated: Use DepositParams.ProtoReflect.Descriptor instead.
func (*DepositParams) Descriptor() ([]byte, []int) {
	return file_cosmos_gov_v1beta1_gov_proto_rawDescGZIP(), []int{6}
}

func (x *DepositParams) GetMinDeposit() []*v1beta1.Coin {
	if x != nil {
		return x.MinDeposit
	}
	return nil
}

func (x *DepositParams) GetMaxDepositPeriod() *durationpb.Duration {
	if x != nil {
		return x.MaxDepositPeriod
	}
	return nil
}

// VotingParams defines the params for voting on governance proposals.
type VotingParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  Length of the voting period.
	VotingPeriod *durationpb.Duration `protobuf:"bytes,1,opt,name=voting_period,json=votingPeriod,proto3" json:"voting_period,omitempty"`
}

func (x *VotingParams) Reset() {
	*x = VotingParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_gov_v1beta1_gov_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VotingParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VotingParams) ProtoMessage() {}

// Deprecated: Use VotingParams.ProtoReflect.Descriptor instead.
func (*VotingParams) Descriptor() ([]byte, []int) {
	return file_cosmos_gov_v1beta1_gov_proto_rawDescGZIP(), []int{7}
}

func (x *VotingParams) GetVotingPeriod() *durationpb.Duration {
	if x != nil {
		return x.VotingPeriod
	}
	return nil
}

// TallyParams defines the params for tallying votes on governance proposals.
type TallyParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  Minimum percentage of total stake needed to vote for a result to be
	//  considered valid.
	Quorum []byte `protobuf:"bytes,1,opt,name=quorum,proto3" json:"quorum,omitempty"`
	//  Minimum proportion of Yes votes for proposal to pass. Default value: 0.5.
	Threshold []byte `protobuf:"bytes,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	//  Minimum value of Veto votes to Total votes ratio for proposal to be
	//  vetoed. Default value: 1/3.
	VetoThreshold []byte `protobuf:"bytes,3,opt,name=veto_threshold,json=vetoThreshold,proto3" json:"veto_threshold,omitempty"`
}

func (x *TallyParams) Reset() {
	*x = TallyParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_gov_v1beta1_gov_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TallyParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TallyParams) ProtoMessage() {}

// Deprecated: Use TallyParams.ProtoReflect.Descriptor instead.
func (*TallyParams) Descriptor() ([]byte, []int) {
	return file_cosmos_gov_v1beta1_gov_proto_rawDescGZIP(), []int{8}
}

func (x *TallyParams) GetQuorum() []byte {
	if x != nil {
		return x.Quorum
	}
	return nil
}

func (x *TallyParams) GetThreshold() []byte {
	if x != nil {
		return x.Threshold
	}
	return nil
}

func (x *TallyParams) GetVetoThreshold() []byte {
	if x != nil {
		return x.VetoThreshold
	}
	return nil
}

var File_cosmos_gov_v1beta1_gov_proto protoreflect.FileDescriptor

var file_cosmos_gov_v1beta1_gov_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x76, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x6f, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x6f, 0x76, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa2, 0x01, 0x0a, 0x12, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x56, 0x6f, 0x74, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x67, 0x6f, 0x76, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x6f, 0x74, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x54,
	0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3c,
	0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0xd2, 0xb4,
	0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x22, 0x57, 0x0a, 0x0c, 0x54, 0x65, 0x78, 0x74, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x0f, 0xe8, 0xa0,
	0x1f, 0x01, 0xca, 0xb4, 0x2d, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xd1, 0x01,
	0x0a, 0x07, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x09, 0x64, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2,
	0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x12, 0x63, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x30, 0xc8,
	0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f,
	0x00, 0x22, 0xb0, 0x05, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x27,
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x06, 0xea, 0xde, 0x1f, 0x02, 0x69, 0x64, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x0b,
	0xca, 0xb4, 0x2d, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x6f,
	0x76, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x53, 0x0a, 0x12, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x61, 0x6c, 0x6c, 0x79, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x6f, 0x76, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x54, 0x61, 0x6c, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x04, 0xc8,
	0xde, 0x1f, 0x00, 0x52, 0x10, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x61, 0x6c, 0x6c, 0x79, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x45, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0x90, 0xdf, 0x1f, 0x01,
	0x52, 0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x10,
	0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x0e, 0x64, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x70, 0x0a, 0x0d,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x30,
	0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x73,
	0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x12, 0x50,
	0x0a, 0x11, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0x90, 0xdf, 0x1f, 0x01, 0x52,
	0x0f, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x4c, 0x0a, 0x0f, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0x90, 0xdf, 0x1f, 0x01, 0x52,
	0x0d, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x3a, 0x04,
	0xe8, 0xa0, 0x1f, 0x01, 0x22, 0xe9, 0x02, 0x0a, 0x0b, 0x54, 0x61, 0x6c, 0x6c, 0x79, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x4e, 0x0a, 0x03, 0x79, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x3c, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e,
	0x74, 0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x52,
	0x03, 0x79, 0x65, 0x73, 0x12, 0x56, 0x0a, 0x07, 0x61, 0x62, 0x73, 0x74, 0x61, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3c, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x49, 0x6e, 0x74, 0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x49, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x62, 0x73, 0x74, 0x61, 0x69, 0x6e, 0x12, 0x4c, 0x0a, 0x02,
	0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3c, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde,
	0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x02, 0x6e, 0x6f, 0x12, 0x5e, 0x0a, 0x0c, 0x6e, 0x6f,
	0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x76, 0x65, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x3c, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74,
	0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x0a,
	0x6e, 0x6f, 0x57, 0x69, 0x74, 0x68, 0x56, 0x65, 0x74, 0x6f, 0x3a, 0x04, 0xe8, 0xa0, 0x1f, 0x01,
	0x22, 0xe5, 0x01, 0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x6f,
	0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x06, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x6f, 0x76, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x56, 0x6f, 0x74, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x02, 0x18, 0x01, 0x52, 0x06,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x67, 0x6f, 0x76, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x65, 0x64, 0x56, 0x6f, 0x74, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x08,
	0x98, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x8a, 0x02, 0x0a, 0x0d, 0x44, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x85, 0x01, 0x0a, 0x0b, 0x6d,
	0x69, 0x6e, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x49, 0xc8, 0xde, 0x1f,
	0x00, 0xea, 0xde, 0x1f, 0x15, 0x6d, 0x69, 0x6e, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0xaa, 0xdf, 0x1f, 0x28, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x0a, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x12, 0x71, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x28, 0xc8, 0xde, 0x1f, 0x00, 0xea,
	0xde, 0x1f, 0x1c, 0x6d, 0x61, 0x78, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x5f, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x98,
	0xdf, 0x1f, 0x01, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0x73, 0x0a, 0x0c, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x63, 0x0a, 0x0d, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x23, 0xc8, 0xde, 0x1f, 0x00, 0xea, 0xde, 0x1f,
	0x17, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x2c, 0x6f,
	0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x98, 0xdf, 0x1f, 0x01, 0x52, 0x0c, 0x76, 0x6f,
	0x74, 0x69, 0x6e, 0x67, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0xc1, 0x02, 0x0a, 0x0b, 0x54,
	0x61, 0x6c, 0x6c, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x5a, 0x0a, 0x06, 0x71, 0x75,
	0x6f, 0x72, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x42, 0xc8, 0xde, 0x1f, 0x00,
	0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0xea, 0xde, 0x1f, 0x10, 0x71, 0x75,
	0x6f, 0x72, 0x75, 0x6d, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x06,
	0x71, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x12, 0x63, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x45, 0xc8, 0xde, 0x1f, 0x00, 0xda,
	0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0xea, 0xde, 0x1f, 0x13, 0x74, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x71, 0x0a, 0x0e, 0x76,
	0x65, 0x74, 0x6f, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x42, 0x4a, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x44, 0x65, 0x63, 0xea, 0xde, 0x1f, 0x18, 0x76, 0x65, 0x74, 0x6f, 0x5f, 0x74, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x0d, 0x76, 0x65, 0x74, 0x6f, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x2a, 0xe6,
	0x01, 0x0a, 0x0a, 0x56, 0x6f, 0x74, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a,
	0x17, 0x56, 0x4f, 0x54, 0x45, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x0f, 0x8a, 0x9d, 0x20, 0x0b,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0f, 0x56,
	0x4f, 0x54, 0x45, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x59, 0x45, 0x53, 0x10, 0x01,
	0x1a, 0x0d, 0x8a, 0x9d, 0x20, 0x09, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x65, 0x73, 0x12,
	0x2a, 0x0a, 0x13, 0x56, 0x4f, 0x54, 0x45, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41,
	0x42, 0x53, 0x54, 0x41, 0x49, 0x4e, 0x10, 0x02, 0x1a, 0x11, 0x8a, 0x9d, 0x20, 0x0d, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x62, 0x73, 0x74, 0x61, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0e, 0x56,
	0x4f, 0x54, 0x45, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x10, 0x03, 0x1a,
	0x0c, 0x8a, 0x9d, 0x20, 0x08, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x12, 0x32, 0x0a,
	0x18, 0x56, 0x4f, 0x54, 0x45, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x5f,
	0x57, 0x49, 0x54, 0x48, 0x5f, 0x56, 0x45, 0x54, 0x4f, 0x10, 0x04, 0x1a, 0x14, 0x8a, 0x9d, 0x20,
	0x10, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x57, 0x69, 0x74, 0x68, 0x56, 0x65, 0x74,
	0x6f, 0x1a, 0x04, 0x88, 0xa3, 0x1e, 0x00, 0x2a, 0xcc, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x0a, 0x1b, 0x50, 0x52,
	0x4f, 0x50, 0x4f, 0x53, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x0d, 0x8a, 0x9d, 0x20,
	0x09, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x69, 0x6c, 0x12, 0x3b, 0x0a, 0x1e, 0x50, 0x52,
	0x4f, 0x50, 0x4f, 0x53, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45,
	0x50, 0x4f, 0x53, 0x49, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x49, 0x4f, 0x44, 0x10, 0x01, 0x1a, 0x17,
	0x8a, 0x9d, 0x20, 0x13, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x39, 0x0a, 0x1d, 0x50, 0x52, 0x4f, 0x50, 0x4f,
	0x53, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x56, 0x4f, 0x54, 0x49, 0x4e,
	0x47, 0x5f, 0x50, 0x45, 0x52, 0x49, 0x4f, 0x44, 0x10, 0x02, 0x1a, 0x16, 0x8a, 0x9d, 0x20, 0x12,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x12, 0x2c, 0x0a, 0x16, 0x50, 0x52, 0x4f, 0x50, 0x4f, 0x53, 0x41, 0x4c, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x45, 0x44, 0x10, 0x03, 0x1a, 0x10,
	0x8a, 0x9d, 0x20, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x61, 0x73, 0x73, 0x65, 0x64,
	0x12, 0x30, 0x0a, 0x18, 0x50, 0x52, 0x4f, 0x50, 0x4f, 0x53, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x04, 0x1a, 0x12,
	0x8a, 0x9d, 0x20, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x12, 0x2c, 0x0a, 0x16, 0x50, 0x52, 0x4f, 0x50, 0x4f, 0x53, 0x41, 0x4c, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x1a, 0x10,
	0x8a, 0x9d, 0x20, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x1a, 0x04, 0x88, 0xa3, 0x1e, 0x00, 0x42, 0xd8, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x67, 0x6f, 0x76, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x42, 0x08, 0x47, 0x6f, 0x76, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x76, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x3b, 0x67, 0x6f, 0x76, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03,
	0x43, 0x47, 0x58, 0xaa, 0x02, 0x12, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x47, 0x6f, 0x76,
	0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x12, 0x43, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x5c, 0x47, 0x6f, 0x76, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x1e,
	0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x47, 0x6f, 0x76, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x14, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x76, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0xc8, 0xe1, 0x1e, 0x00, 0xd8, 0xe1, 0x1e, 0x00, 0x80, 0xe2, 0x1e,
	0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_gov_v1beta1_gov_proto_rawDescOnce sync.Once
	file_cosmos_gov_v1beta1_gov_proto_rawDescData = file_cosmos_gov_v1beta1_gov_proto_rawDesc
)

func file_cosmos_gov_v1beta1_gov_proto_rawDescGZIP() []byte {
	file_cosmos_gov_v1beta1_gov_proto_rawDescOnce.Do(func() {
		file_cosmos_gov_v1beta1_gov_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_gov_v1beta1_gov_proto_rawDescData)
	})
	return file_cosmos_gov_v1beta1_gov_proto_rawDescData
}

var file_cosmos_gov_v1beta1_gov_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cosmos_gov_v1beta1_gov_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_cosmos_gov_v1beta1_gov_proto_goTypes = []interface{}{
	(VoteOption)(0),               // 0: cosmos.gov.v1beta1.VoteOption
	(ProposalStatus)(0),           // 1: cosmos.gov.v1beta1.ProposalStatus
	(*WeightedVoteOption)(nil),    // 2: cosmos.gov.v1beta1.WeightedVoteOption
	(*TextProposal)(nil),          // 3: cosmos.gov.v1beta1.TextProposal
	(*Deposit)(nil),               // 4: cosmos.gov.v1beta1.Deposit
	(*Proposal)(nil),              // 5: cosmos.gov.v1beta1.Proposal
	(*TallyResult)(nil),           // 6: cosmos.gov.v1beta1.TallyResult
	(*Vote)(nil),                  // 7: cosmos.gov.v1beta1.Vote
	(*DepositParams)(nil),         // 8: cosmos.gov.v1beta1.DepositParams
	(*VotingParams)(nil),          // 9: cosmos.gov.v1beta1.VotingParams
	(*TallyParams)(nil),           // 10: cosmos.gov.v1beta1.TallyParams
	(*v1beta1.Coin)(nil),          // 11: cosmos.base.v1beta1.Coin
	(*anypb.Any)(nil),             // 12: google.protobuf.Any
	(*timestamppb.Timestamp)(nil), // 13: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 14: google.protobuf.Duration
}
var file_cosmos_gov_v1beta1_gov_proto_depIdxs = []int32{
	0,  // 0: cosmos.gov.v1beta1.WeightedVoteOption.option:type_name -> cosmos.gov.v1beta1.VoteOption
	11, // 1: cosmos.gov.v1beta1.Deposit.amount:type_name -> cosmos.base.v1beta1.Coin
	12, // 2: cosmos.gov.v1beta1.Proposal.content:type_name -> google.protobuf.Any
	1,  // 3: cosmos.gov.v1beta1.Proposal.status:type_name -> cosmos.gov.v1beta1.ProposalStatus
	6,  // 4: cosmos.gov.v1beta1.Proposal.final_tally_result:type_name -> cosmos.gov.v1beta1.TallyResult
	13, // 5: cosmos.gov.v1beta1.Proposal.submit_time:type_name -> google.protobuf.Timestamp
	13, // 6: cosmos.gov.v1beta1.Proposal.deposit_end_time:type_name -> google.protobuf.Timestamp
	11, // 7: cosmos.gov.v1beta1.Proposal.total_deposit:type_name -> cosmos.base.v1beta1.Coin
	13, // 8: cosmos.gov.v1beta1.Proposal.voting_start_time:type_name -> google.protobuf.Timestamp
	13, // 9: cosmos.gov.v1beta1.Proposal.voting_end_time:type_name -> google.protobuf.Timestamp
	0,  // 10: cosmos.gov.v1beta1.Vote.option:type_name -> cosmos.gov.v1beta1.VoteOption
	2,  // 11: cosmos.gov.v1beta1.Vote.options:type_name -> cosmos.gov.v1beta1.WeightedVoteOption
	11, // 12: cosmos.gov.v1beta1.DepositParams.min_deposit:type_name -> cosmos.base.v1beta1.Coin
	14, // 13: cosmos.gov.v1beta1.DepositParams.max_deposit_period:type_name -> google.protobuf.Duration
	14, // 14: cosmos.gov.v1beta1.VotingParams.voting_period:type_name -> google.protobuf.Duration
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_cosmos_gov_v1beta1_gov_proto_init() }
func file_cosmos_gov_v1beta1_gov_proto_init() {
	if File_cosmos_gov_v1beta1_gov_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_gov_v1beta1_gov_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeightedVoteOption); i {
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
		file_cosmos_gov_v1beta1_gov_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextProposal); i {
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
		file_cosmos_gov_v1beta1_gov_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deposit); i {
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
		file_cosmos_gov_v1beta1_gov_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proposal); i {
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
		file_cosmos_gov_v1beta1_gov_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TallyResult); i {
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
		file_cosmos_gov_v1beta1_gov_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vote); i {
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
		file_cosmos_gov_v1beta1_gov_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositParams); i {
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
		file_cosmos_gov_v1beta1_gov_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VotingParams); i {
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
		file_cosmos_gov_v1beta1_gov_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TallyParams); i {
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
			RawDescriptor: file_cosmos_gov_v1beta1_gov_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_gov_v1beta1_gov_proto_goTypes,
		DependencyIndexes: file_cosmos_gov_v1beta1_gov_proto_depIdxs,
		EnumInfos:         file_cosmos_gov_v1beta1_gov_proto_enumTypes,
		MessageInfos:      file_cosmos_gov_v1beta1_gov_proto_msgTypes,
	}.Build()
	File_cosmos_gov_v1beta1_gov_proto = out.File
	file_cosmos_gov_v1beta1_gov_proto_rawDesc = nil
	file_cosmos_gov_v1beta1_gov_proto_goTypes = nil
	file_cosmos_gov_v1beta1_gov_proto_depIdxs = nil
}
