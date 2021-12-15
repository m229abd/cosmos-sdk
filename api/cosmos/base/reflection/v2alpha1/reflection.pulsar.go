package reflectionv2alpha1

import (
	context "context"
	fmt "fmt"
	io "io"
	reflect "reflect"
	sync "sync"

	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

var (
	md_AppDescriptor                protoreflect.MessageDescriptor
	fd_AppDescriptor_authn          protoreflect.FieldDescriptor
	fd_AppDescriptor_chain          protoreflect.FieldDescriptor
	fd_AppDescriptor_codec          protoreflect.FieldDescriptor
	fd_AppDescriptor_configuration  protoreflect.FieldDescriptor
	fd_AppDescriptor_query_services protoreflect.FieldDescriptor
	fd_AppDescriptor_tx             protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_AppDescriptor = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("AppDescriptor")
	fd_AppDescriptor_authn = md_AppDescriptor.Fields().ByName("authn")
	fd_AppDescriptor_chain = md_AppDescriptor.Fields().ByName("chain")
	fd_AppDescriptor_codec = md_AppDescriptor.Fields().ByName("codec")
	fd_AppDescriptor_configuration = md_AppDescriptor.Fields().ByName("configuration")
	fd_AppDescriptor_query_services = md_AppDescriptor.Fields().ByName("query_services")
	fd_AppDescriptor_tx = md_AppDescriptor.Fields().ByName("tx")
}

var _ protoreflect.Message = (*fastReflection_AppDescriptor)(nil)

type fastReflection_AppDescriptor AppDescriptor

func (x *AppDescriptor) ProtoReflect() protoreflect.Message {
	return (*fastReflection_AppDescriptor)(x)
}

func (x *AppDescriptor) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_AppDescriptor_messageType fastReflection_AppDescriptor_messageType
var _ protoreflect.MessageType = fastReflection_AppDescriptor_messageType{}

type fastReflection_AppDescriptor_messageType struct{}

func (x fastReflection_AppDescriptor_messageType) Zero() protoreflect.Message {
	return (*fastReflection_AppDescriptor)(nil)
}
func (x fastReflection_AppDescriptor_messageType) New() protoreflect.Message {
	return new(fastReflection_AppDescriptor)
}
func (x fastReflection_AppDescriptor_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_AppDescriptor
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_AppDescriptor) Descriptor() protoreflect.MessageDescriptor {
	return md_AppDescriptor
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_AppDescriptor) Type() protoreflect.MessageType {
	return _fastReflection_AppDescriptor_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_AppDescriptor) New() protoreflect.Message {
	return new(fastReflection_AppDescriptor)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_AppDescriptor) Interface() protoreflect.ProtoMessage {
	return (*AppDescriptor)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_AppDescriptor) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Authn != nil {
		value := protoreflect.ValueOfMessage(x.Authn.ProtoReflect())
		if !f(fd_AppDescriptor_authn, value) {
			return
		}
	}
	if x.Chain != nil {
		value := protoreflect.ValueOfMessage(x.Chain.ProtoReflect())
		if !f(fd_AppDescriptor_chain, value) {
			return
		}
	}
	if x.Codec != nil {
		value := protoreflect.ValueOfMessage(x.Codec.ProtoReflect())
		if !f(fd_AppDescriptor_codec, value) {
			return
		}
	}
	if x.Configuration != nil {
		value := protoreflect.ValueOfMessage(x.Configuration.ProtoReflect())
		if !f(fd_AppDescriptor_configuration, value) {
			return
		}
	}
	if x.QueryServices != nil {
		value := protoreflect.ValueOfMessage(x.QueryServices.ProtoReflect())
		if !f(fd_AppDescriptor_query_services, value) {
			return
		}
	}
	if x.Tx != nil {
		value := protoreflect.ValueOfMessage(x.Tx.ProtoReflect())
		if !f(fd_AppDescriptor_tx, value) {
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
func (x *fastReflection_AppDescriptor) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.authn":
		return x.Authn != nil
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.chain":
		return x.Chain != nil
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.codec":
		return x.Codec != nil
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.configuration":
		return x.Configuration != nil
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.query_services":
		return x.QueryServices != nil
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.tx":
		return x.Tx != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.AppDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.AppDescriptor does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AppDescriptor) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.authn":
		x.Authn = nil
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.chain":
		x.Chain = nil
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.codec":
		x.Codec = nil
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.configuration":
		x.Configuration = nil
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.query_services":
		x.QueryServices = nil
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.tx":
		x.Tx = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.AppDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.AppDescriptor does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_AppDescriptor) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.authn":
		value := x.Authn
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.chain":
		value := x.Chain
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.codec":
		value := x.Codec
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.configuration":
		value := x.Configuration
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.query_services":
		value := x.QueryServices
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.tx":
		value := x.Tx
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.AppDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.AppDescriptor does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_AppDescriptor) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.authn":
		x.Authn = value.Message().Interface().(*AuthnDescriptor)
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.chain":
		x.Chain = value.Message().Interface().(*ChainDescriptor)
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.codec":
		x.Codec = value.Message().Interface().(*CodecDescriptor)
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.configuration":
		x.Configuration = value.Message().Interface().(*ConfigurationDescriptor)
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.query_services":
		x.QueryServices = value.Message().Interface().(*QueryServicesDescriptor)
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.tx":
		x.Tx = value.Message().Interface().(*TxDescriptor)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.AppDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.AppDescriptor does not contain field %s", fd.FullName()))
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
func (x *fastReflection_AppDescriptor) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.authn":
		if x.Authn == nil {
			x.Authn = new(AuthnDescriptor)
		}
		return protoreflect.ValueOfMessage(x.Authn.ProtoReflect())
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.chain":
		if x.Chain == nil {
			x.Chain = new(ChainDescriptor)
		}
		return protoreflect.ValueOfMessage(x.Chain.ProtoReflect())
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.codec":
		if x.Codec == nil {
			x.Codec = new(CodecDescriptor)
		}
		return protoreflect.ValueOfMessage(x.Codec.ProtoReflect())
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.configuration":
		if x.Configuration == nil {
			x.Configuration = new(ConfigurationDescriptor)
		}
		return protoreflect.ValueOfMessage(x.Configuration.ProtoReflect())
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.query_services":
		if x.QueryServices == nil {
			x.QueryServices = new(QueryServicesDescriptor)
		}
		return protoreflect.ValueOfMessage(x.QueryServices.ProtoReflect())
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.tx":
		if x.Tx == nil {
			x.Tx = new(TxDescriptor)
		}
		return protoreflect.ValueOfMessage(x.Tx.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.AppDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.AppDescriptor does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_AppDescriptor) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.authn":
		m := new(AuthnDescriptor)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.chain":
		m := new(ChainDescriptor)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.codec":
		m := new(CodecDescriptor)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.configuration":
		m := new(ConfigurationDescriptor)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.query_services":
		m := new(QueryServicesDescriptor)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.base.reflection.v2alpha1.AppDescriptor.tx":
		m := new(TxDescriptor)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.AppDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.AppDescriptor does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_AppDescriptor) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.AppDescriptor", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_AppDescriptor) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AppDescriptor) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_AppDescriptor) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_AppDescriptor) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*AppDescriptor)
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
		if x.Authn != nil {
			l = options.Size(x.Authn)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Chain != nil {
			l = options.Size(x.Chain)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Codec != nil {
			l = options.Size(x.Codec)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Configuration != nil {
			l = options.Size(x.Configuration)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.QueryServices != nil {
			l = options.Size(x.QueryServices)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Tx != nil {
			l = options.Size(x.Tx)
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
		x := input.Message.Interface().(*AppDescriptor)
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
		if x.Tx != nil {
			encoded, err := options.Marshal(x.Tx)
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
		if x.QueryServices != nil {
			encoded, err := options.Marshal(x.QueryServices)
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
		if x.Configuration != nil {
			encoded, err := options.Marshal(x.Configuration)
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
		if x.Codec != nil {
			encoded, err := options.Marshal(x.Codec)
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
		if x.Chain != nil {
			encoded, err := options.Marshal(x.Chain)
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
		if x.Authn != nil {
			encoded, err := options.Marshal(x.Authn)
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
		x := input.Message.Interface().(*AppDescriptor)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AppDescriptor: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AppDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Authn", wireType)
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
				if x.Authn == nil {
					x.Authn = &AuthnDescriptor{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Authn); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
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
				if x.Chain == nil {
					x.Chain = &ChainDescriptor{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Chain); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Codec", wireType)
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
				if x.Codec == nil {
					x.Codec = &CodecDescriptor{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Codec); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Configuration", wireType)
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
				if x.Configuration == nil {
					x.Configuration = &ConfigurationDescriptor{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Configuration); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field QueryServices", wireType)
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
				if x.QueryServices == nil {
					x.QueryServices = &QueryServicesDescriptor{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.QueryServices); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
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
				if x.Tx == nil {
					x.Tx = &TxDescriptor{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Tx); err != nil {
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

var _ protoreflect.List = (*_TxDescriptor_2_list)(nil)

type _TxDescriptor_2_list struct {
	list *[]*MsgDescriptor
}

func (x *_TxDescriptor_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TxDescriptor_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_TxDescriptor_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgDescriptor)
	(*x.list)[i] = concreteValue
}

func (x *_TxDescriptor_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgDescriptor)
	*x.list = append(*x.list, concreteValue)
}

func (x *_TxDescriptor_2_list) AppendMutable() protoreflect.Value {
	v := new(MsgDescriptor)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TxDescriptor_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_TxDescriptor_2_list) NewElement() protoreflect.Value {
	v := new(MsgDescriptor)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TxDescriptor_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_TxDescriptor          protoreflect.MessageDescriptor
	fd_TxDescriptor_fullname protoreflect.FieldDescriptor
	fd_TxDescriptor_msgs     protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_TxDescriptor = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("TxDescriptor")
	fd_TxDescriptor_fullname = md_TxDescriptor.Fields().ByName("fullname")
	fd_TxDescriptor_msgs = md_TxDescriptor.Fields().ByName("msgs")
}

var _ protoreflect.Message = (*fastReflection_TxDescriptor)(nil)

type fastReflection_TxDescriptor TxDescriptor

func (x *TxDescriptor) ProtoReflect() protoreflect.Message {
	return (*fastReflection_TxDescriptor)(x)
}

func (x *TxDescriptor) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_TxDescriptor_messageType fastReflection_TxDescriptor_messageType
var _ protoreflect.MessageType = fastReflection_TxDescriptor_messageType{}

type fastReflection_TxDescriptor_messageType struct{}

func (x fastReflection_TxDescriptor_messageType) Zero() protoreflect.Message {
	return (*fastReflection_TxDescriptor)(nil)
}
func (x fastReflection_TxDescriptor_messageType) New() protoreflect.Message {
	return new(fastReflection_TxDescriptor)
}
func (x fastReflection_TxDescriptor_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_TxDescriptor
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_TxDescriptor) Descriptor() protoreflect.MessageDescriptor {
	return md_TxDescriptor
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_TxDescriptor) Type() protoreflect.MessageType {
	return _fastReflection_TxDescriptor_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_TxDescriptor) New() protoreflect.Message {
	return new(fastReflection_TxDescriptor)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_TxDescriptor) Interface() protoreflect.ProtoMessage {
	return (*TxDescriptor)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_TxDescriptor) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Fullname != "" {
		value := protoreflect.ValueOfString(x.Fullname)
		if !f(fd_TxDescriptor_fullname, value) {
			return
		}
	}
	if len(x.Msgs) != 0 {
		value := protoreflect.ValueOfList(&_TxDescriptor_2_list{list: &x.Msgs})
		if !f(fd_TxDescriptor_msgs, value) {
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
func (x *fastReflection_TxDescriptor) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.TxDescriptor.fullname":
		return x.Fullname != ""
	case "cosmos.base.reflection.v2alpha1.TxDescriptor.msgs":
		return len(x.Msgs) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.TxDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.TxDescriptor does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TxDescriptor) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.TxDescriptor.fullname":
		x.Fullname = ""
	case "cosmos.base.reflection.v2alpha1.TxDescriptor.msgs":
		x.Msgs = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.TxDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.TxDescriptor does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_TxDescriptor) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v2alpha1.TxDescriptor.fullname":
		value := x.Fullname
		return protoreflect.ValueOfString(value)
	case "cosmos.base.reflection.v2alpha1.TxDescriptor.msgs":
		if len(x.Msgs) == 0 {
			return protoreflect.ValueOfList(&_TxDescriptor_2_list{})
		}
		listValue := &_TxDescriptor_2_list{list: &x.Msgs}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.TxDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.TxDescriptor does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_TxDescriptor) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.TxDescriptor.fullname":
		x.Fullname = value.Interface().(string)
	case "cosmos.base.reflection.v2alpha1.TxDescriptor.msgs":
		lv := value.List()
		clv := lv.(*_TxDescriptor_2_list)
		x.Msgs = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.TxDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.TxDescriptor does not contain field %s", fd.FullName()))
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
func (x *fastReflection_TxDescriptor) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.TxDescriptor.msgs":
		if x.Msgs == nil {
			x.Msgs = []*MsgDescriptor{}
		}
		value := &_TxDescriptor_2_list{list: &x.Msgs}
		return protoreflect.ValueOfList(value)
	case "cosmos.base.reflection.v2alpha1.TxDescriptor.fullname":
		panic(fmt.Errorf("field fullname of message cosmos.base.reflection.v2alpha1.TxDescriptor is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.TxDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.TxDescriptor does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_TxDescriptor) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.TxDescriptor.fullname":
		return protoreflect.ValueOfString("")
	case "cosmos.base.reflection.v2alpha1.TxDescriptor.msgs":
		list := []*MsgDescriptor{}
		return protoreflect.ValueOfList(&_TxDescriptor_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.TxDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.TxDescriptor does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_TxDescriptor) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.TxDescriptor", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_TxDescriptor) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TxDescriptor) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_TxDescriptor) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_TxDescriptor) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*TxDescriptor)
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
		l = len(x.Fullname)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Msgs) > 0 {
			for _, e := range x.Msgs {
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
		x := input.Message.Interface().(*TxDescriptor)
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
		if len(x.Msgs) > 0 {
			for iNdEx := len(x.Msgs) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Msgs[iNdEx])
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
		if len(x.Fullname) > 0 {
			i -= len(x.Fullname)
			copy(dAtA[i:], x.Fullname)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Fullname)))
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
		x := input.Message.Interface().(*TxDescriptor)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TxDescriptor: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TxDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Fullname", wireType)
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
				x.Fullname = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Msgs", wireType)
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
				x.Msgs = append(x.Msgs, &MsgDescriptor{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Msgs[len(x.Msgs)-1]); err != nil {
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

var _ protoreflect.List = (*_AuthnDescriptor_1_list)(nil)

type _AuthnDescriptor_1_list struct {
	list *[]*SigningModeDescriptor
}

func (x *_AuthnDescriptor_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_AuthnDescriptor_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_AuthnDescriptor_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*SigningModeDescriptor)
	(*x.list)[i] = concreteValue
}

func (x *_AuthnDescriptor_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*SigningModeDescriptor)
	*x.list = append(*x.list, concreteValue)
}

func (x *_AuthnDescriptor_1_list) AppendMutable() protoreflect.Value {
	v := new(SigningModeDescriptor)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_AuthnDescriptor_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_AuthnDescriptor_1_list) NewElement() protoreflect.Value {
	v := new(SigningModeDescriptor)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_AuthnDescriptor_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_AuthnDescriptor            protoreflect.MessageDescriptor
	fd_AuthnDescriptor_sign_modes protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_AuthnDescriptor = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("AuthnDescriptor")
	fd_AuthnDescriptor_sign_modes = md_AuthnDescriptor.Fields().ByName("sign_modes")
}

var _ protoreflect.Message = (*fastReflection_AuthnDescriptor)(nil)

type fastReflection_AuthnDescriptor AuthnDescriptor

func (x *AuthnDescriptor) ProtoReflect() protoreflect.Message {
	return (*fastReflection_AuthnDescriptor)(x)
}

func (x *AuthnDescriptor) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_AuthnDescriptor_messageType fastReflection_AuthnDescriptor_messageType
var _ protoreflect.MessageType = fastReflection_AuthnDescriptor_messageType{}

type fastReflection_AuthnDescriptor_messageType struct{}

func (x fastReflection_AuthnDescriptor_messageType) Zero() protoreflect.Message {
	return (*fastReflection_AuthnDescriptor)(nil)
}
func (x fastReflection_AuthnDescriptor_messageType) New() protoreflect.Message {
	return new(fastReflection_AuthnDescriptor)
}
func (x fastReflection_AuthnDescriptor_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_AuthnDescriptor
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_AuthnDescriptor) Descriptor() protoreflect.MessageDescriptor {
	return md_AuthnDescriptor
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_AuthnDescriptor) Type() protoreflect.MessageType {
	return _fastReflection_AuthnDescriptor_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_AuthnDescriptor) New() protoreflect.Message {
	return new(fastReflection_AuthnDescriptor)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_AuthnDescriptor) Interface() protoreflect.ProtoMessage {
	return (*AuthnDescriptor)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_AuthnDescriptor) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.SignModes) != 0 {
		value := protoreflect.ValueOfList(&_AuthnDescriptor_1_list{list: &x.SignModes})
		if !f(fd_AuthnDescriptor_sign_modes, value) {
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
func (x *fastReflection_AuthnDescriptor) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.AuthnDescriptor.sign_modes":
		return len(x.SignModes) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.AuthnDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.AuthnDescriptor does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AuthnDescriptor) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.AuthnDescriptor.sign_modes":
		x.SignModes = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.AuthnDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.AuthnDescriptor does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_AuthnDescriptor) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v2alpha1.AuthnDescriptor.sign_modes":
		if len(x.SignModes) == 0 {
			return protoreflect.ValueOfList(&_AuthnDescriptor_1_list{})
		}
		listValue := &_AuthnDescriptor_1_list{list: &x.SignModes}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.AuthnDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.AuthnDescriptor does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_AuthnDescriptor) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.AuthnDescriptor.sign_modes":
		lv := value.List()
		clv := lv.(*_AuthnDescriptor_1_list)
		x.SignModes = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.AuthnDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.AuthnDescriptor does not contain field %s", fd.FullName()))
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
func (x *fastReflection_AuthnDescriptor) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.AuthnDescriptor.sign_modes":
		if x.SignModes == nil {
			x.SignModes = []*SigningModeDescriptor{}
		}
		value := &_AuthnDescriptor_1_list{list: &x.SignModes}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.AuthnDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.AuthnDescriptor does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_AuthnDescriptor) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.AuthnDescriptor.sign_modes":
		list := []*SigningModeDescriptor{}
		return protoreflect.ValueOfList(&_AuthnDescriptor_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.AuthnDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.AuthnDescriptor does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_AuthnDescriptor) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.AuthnDescriptor", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_AuthnDescriptor) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AuthnDescriptor) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_AuthnDescriptor) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_AuthnDescriptor) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*AuthnDescriptor)
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
		if len(x.SignModes) > 0 {
			for _, e := range x.SignModes {
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
		x := input.Message.Interface().(*AuthnDescriptor)
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
		if len(x.SignModes) > 0 {
			for iNdEx := len(x.SignModes) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.SignModes[iNdEx])
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
		x := input.Message.Interface().(*AuthnDescriptor)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AuthnDescriptor: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AuthnDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SignModes", wireType)
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
				x.SignModes = append(x.SignModes, &SigningModeDescriptor{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.SignModes[len(x.SignModes)-1]); err != nil {
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
	md_SigningModeDescriptor                                     protoreflect.MessageDescriptor
	fd_SigningModeDescriptor_name                                protoreflect.FieldDescriptor
	fd_SigningModeDescriptor_number                              protoreflect.FieldDescriptor
	fd_SigningModeDescriptor_authn_info_provider_method_fullname protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_SigningModeDescriptor = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("SigningModeDescriptor")
	fd_SigningModeDescriptor_name = md_SigningModeDescriptor.Fields().ByName("name")
	fd_SigningModeDescriptor_number = md_SigningModeDescriptor.Fields().ByName("number")
	fd_SigningModeDescriptor_authn_info_provider_method_fullname = md_SigningModeDescriptor.Fields().ByName("authn_info_provider_method_fullname")
}

var _ protoreflect.Message = (*fastReflection_SigningModeDescriptor)(nil)

type fastReflection_SigningModeDescriptor SigningModeDescriptor

func (x *SigningModeDescriptor) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SigningModeDescriptor)(x)
}

func (x *SigningModeDescriptor) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SigningModeDescriptor_messageType fastReflection_SigningModeDescriptor_messageType
var _ protoreflect.MessageType = fastReflection_SigningModeDescriptor_messageType{}

type fastReflection_SigningModeDescriptor_messageType struct{}

func (x fastReflection_SigningModeDescriptor_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SigningModeDescriptor)(nil)
}
func (x fastReflection_SigningModeDescriptor_messageType) New() protoreflect.Message {
	return new(fastReflection_SigningModeDescriptor)
}
func (x fastReflection_SigningModeDescriptor_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SigningModeDescriptor
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SigningModeDescriptor) Descriptor() protoreflect.MessageDescriptor {
	return md_SigningModeDescriptor
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SigningModeDescriptor) Type() protoreflect.MessageType {
	return _fastReflection_SigningModeDescriptor_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SigningModeDescriptor) New() protoreflect.Message {
	return new(fastReflection_SigningModeDescriptor)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SigningModeDescriptor) Interface() protoreflect.ProtoMessage {
	return (*SigningModeDescriptor)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SigningModeDescriptor) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Name != "" {
		value := protoreflect.ValueOfString(x.Name)
		if !f(fd_SigningModeDescriptor_name, value) {
			return
		}
	}
	if x.Number != int32(0) {
		value := protoreflect.ValueOfInt32(x.Number)
		if !f(fd_SigningModeDescriptor_number, value) {
			return
		}
	}
	if x.AuthnInfoProviderMethodFullname != "" {
		value := protoreflect.ValueOfString(x.AuthnInfoProviderMethodFullname)
		if !f(fd_SigningModeDescriptor_authn_info_provider_method_fullname, value) {
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
func (x *fastReflection_SigningModeDescriptor) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.SigningModeDescriptor.name":
		return x.Name != ""
	case "cosmos.base.reflection.v2alpha1.SigningModeDescriptor.number":
		return x.Number != int32(0)
	case "cosmos.base.reflection.v2alpha1.SigningModeDescriptor.authn_info_provider_method_fullname":
		return x.AuthnInfoProviderMethodFullname != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.SigningModeDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.SigningModeDescriptor does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SigningModeDescriptor) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.SigningModeDescriptor.name":
		x.Name = ""
	case "cosmos.base.reflection.v2alpha1.SigningModeDescriptor.number":
		x.Number = int32(0)
	case "cosmos.base.reflection.v2alpha1.SigningModeDescriptor.authn_info_provider_method_fullname":
		x.AuthnInfoProviderMethodFullname = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.SigningModeDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.SigningModeDescriptor does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SigningModeDescriptor) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v2alpha1.SigningModeDescriptor.name":
		value := x.Name
		return protoreflect.ValueOfString(value)
	case "cosmos.base.reflection.v2alpha1.SigningModeDescriptor.number":
		value := x.Number
		return protoreflect.ValueOfInt32(value)
	case "cosmos.base.reflection.v2alpha1.SigningModeDescriptor.authn_info_provider_method_fullname":
		value := x.AuthnInfoProviderMethodFullname
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.SigningModeDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.SigningModeDescriptor does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SigningModeDescriptor) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.SigningModeDescriptor.name":
		x.Name = value.Interface().(string)
	case "cosmos.base.reflection.v2alpha1.SigningModeDescriptor.number":
		x.Number = int32(value.Int())
	case "cosmos.base.reflection.v2alpha1.SigningModeDescriptor.authn_info_provider_method_fullname":
		x.AuthnInfoProviderMethodFullname = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.SigningModeDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.SigningModeDescriptor does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SigningModeDescriptor) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.SigningModeDescriptor.name":
		panic(fmt.Errorf("field name of message cosmos.base.reflection.v2alpha1.SigningModeDescriptor is not mutable"))
	case "cosmos.base.reflection.v2alpha1.SigningModeDescriptor.number":
		panic(fmt.Errorf("field number of message cosmos.base.reflection.v2alpha1.SigningModeDescriptor is not mutable"))
	case "cosmos.base.reflection.v2alpha1.SigningModeDescriptor.authn_info_provider_method_fullname":
		panic(fmt.Errorf("field authn_info_provider_method_fullname of message cosmos.base.reflection.v2alpha1.SigningModeDescriptor is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.SigningModeDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.SigningModeDescriptor does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SigningModeDescriptor) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.SigningModeDescriptor.name":
		return protoreflect.ValueOfString("")
	case "cosmos.base.reflection.v2alpha1.SigningModeDescriptor.number":
		return protoreflect.ValueOfInt32(int32(0))
	case "cosmos.base.reflection.v2alpha1.SigningModeDescriptor.authn_info_provider_method_fullname":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.SigningModeDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.SigningModeDescriptor does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SigningModeDescriptor) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.SigningModeDescriptor", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SigningModeDescriptor) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SigningModeDescriptor) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SigningModeDescriptor) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SigningModeDescriptor) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SigningModeDescriptor)
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
		l = len(x.Name)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Number != 0 {
			n += 1 + runtime.Sov(uint64(x.Number))
		}
		l = len(x.AuthnInfoProviderMethodFullname)
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
		x := input.Message.Interface().(*SigningModeDescriptor)
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
		if len(x.AuthnInfoProviderMethodFullname) > 0 {
			i -= len(x.AuthnInfoProviderMethodFullname)
			copy(dAtA[i:], x.AuthnInfoProviderMethodFullname)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.AuthnInfoProviderMethodFullname)))
			i--
			dAtA[i] = 0x1a
		}
		if x.Number != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Number))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Name) > 0 {
			i -= len(x.Name)
			copy(dAtA[i:], x.Name)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Name)))
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
		x := input.Message.Interface().(*SigningModeDescriptor)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SigningModeDescriptor: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SigningModeDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
				}
				x.Number = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Number |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AuthnInfoProviderMethodFullname", wireType)
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
				x.AuthnInfoProviderMethodFullname = string(dAtA[iNdEx:postIndex])
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
	md_ChainDescriptor    protoreflect.MessageDescriptor
	fd_ChainDescriptor_id protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_ChainDescriptor = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("ChainDescriptor")
	fd_ChainDescriptor_id = md_ChainDescriptor.Fields().ByName("id")
}

var _ protoreflect.Message = (*fastReflection_ChainDescriptor)(nil)

type fastReflection_ChainDescriptor ChainDescriptor

func (x *ChainDescriptor) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ChainDescriptor)(x)
}

func (x *ChainDescriptor) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ChainDescriptor_messageType fastReflection_ChainDescriptor_messageType
var _ protoreflect.MessageType = fastReflection_ChainDescriptor_messageType{}

type fastReflection_ChainDescriptor_messageType struct{}

func (x fastReflection_ChainDescriptor_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ChainDescriptor)(nil)
}
func (x fastReflection_ChainDescriptor_messageType) New() protoreflect.Message {
	return new(fastReflection_ChainDescriptor)
}
func (x fastReflection_ChainDescriptor_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ChainDescriptor
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ChainDescriptor) Descriptor() protoreflect.MessageDescriptor {
	return md_ChainDescriptor
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ChainDescriptor) Type() protoreflect.MessageType {
	return _fastReflection_ChainDescriptor_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ChainDescriptor) New() protoreflect.Message {
	return new(fastReflection_ChainDescriptor)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ChainDescriptor) Interface() protoreflect.ProtoMessage {
	return (*ChainDescriptor)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ChainDescriptor) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != "" {
		value := protoreflect.ValueOfString(x.Id)
		if !f(fd_ChainDescriptor_id, value) {
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
func (x *fastReflection_ChainDescriptor) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.ChainDescriptor.id":
		return x.Id != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.ChainDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.ChainDescriptor does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ChainDescriptor) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.ChainDescriptor.id":
		x.Id = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.ChainDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.ChainDescriptor does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ChainDescriptor) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v2alpha1.ChainDescriptor.id":
		value := x.Id
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.ChainDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.ChainDescriptor does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ChainDescriptor) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.ChainDescriptor.id":
		x.Id = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.ChainDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.ChainDescriptor does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ChainDescriptor) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.ChainDescriptor.id":
		panic(fmt.Errorf("field id of message cosmos.base.reflection.v2alpha1.ChainDescriptor is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.ChainDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.ChainDescriptor does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ChainDescriptor) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.ChainDescriptor.id":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.ChainDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.ChainDescriptor does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ChainDescriptor) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.ChainDescriptor", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ChainDescriptor) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ChainDescriptor) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ChainDescriptor) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ChainDescriptor) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ChainDescriptor)
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
		l = len(x.Id)
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
		x := input.Message.Interface().(*ChainDescriptor)
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
		if len(x.Id) > 0 {
			i -= len(x.Id)
			copy(dAtA[i:], x.Id)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Id)))
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
		x := input.Message.Interface().(*ChainDescriptor)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ChainDescriptor: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ChainDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
				x.Id = string(dAtA[iNdEx:postIndex])
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

var _ protoreflect.List = (*_CodecDescriptor_1_list)(nil)

type _CodecDescriptor_1_list struct {
	list *[]*InterfaceDescriptor
}

func (x *_CodecDescriptor_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_CodecDescriptor_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_CodecDescriptor_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*InterfaceDescriptor)
	(*x.list)[i] = concreteValue
}

func (x *_CodecDescriptor_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*InterfaceDescriptor)
	*x.list = append(*x.list, concreteValue)
}

func (x *_CodecDescriptor_1_list) AppendMutable() protoreflect.Value {
	v := new(InterfaceDescriptor)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_CodecDescriptor_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_CodecDescriptor_1_list) NewElement() protoreflect.Value {
	v := new(InterfaceDescriptor)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_CodecDescriptor_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_CodecDescriptor            protoreflect.MessageDescriptor
	fd_CodecDescriptor_interfaces protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_CodecDescriptor = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("CodecDescriptor")
	fd_CodecDescriptor_interfaces = md_CodecDescriptor.Fields().ByName("interfaces")
}

var _ protoreflect.Message = (*fastReflection_CodecDescriptor)(nil)

type fastReflection_CodecDescriptor CodecDescriptor

func (x *CodecDescriptor) ProtoReflect() protoreflect.Message {
	return (*fastReflection_CodecDescriptor)(x)
}

func (x *CodecDescriptor) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_CodecDescriptor_messageType fastReflection_CodecDescriptor_messageType
var _ protoreflect.MessageType = fastReflection_CodecDescriptor_messageType{}

type fastReflection_CodecDescriptor_messageType struct{}

func (x fastReflection_CodecDescriptor_messageType) Zero() protoreflect.Message {
	return (*fastReflection_CodecDescriptor)(nil)
}
func (x fastReflection_CodecDescriptor_messageType) New() protoreflect.Message {
	return new(fastReflection_CodecDescriptor)
}
func (x fastReflection_CodecDescriptor_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_CodecDescriptor
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_CodecDescriptor) Descriptor() protoreflect.MessageDescriptor {
	return md_CodecDescriptor
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_CodecDescriptor) Type() protoreflect.MessageType {
	return _fastReflection_CodecDescriptor_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_CodecDescriptor) New() protoreflect.Message {
	return new(fastReflection_CodecDescriptor)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_CodecDescriptor) Interface() protoreflect.ProtoMessage {
	return (*CodecDescriptor)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_CodecDescriptor) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Interfaces) != 0 {
		value := protoreflect.ValueOfList(&_CodecDescriptor_1_list{list: &x.Interfaces})
		if !f(fd_CodecDescriptor_interfaces, value) {
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
func (x *fastReflection_CodecDescriptor) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.CodecDescriptor.interfaces":
		return len(x.Interfaces) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.CodecDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.CodecDescriptor does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CodecDescriptor) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.CodecDescriptor.interfaces":
		x.Interfaces = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.CodecDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.CodecDescriptor does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_CodecDescriptor) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v2alpha1.CodecDescriptor.interfaces":
		if len(x.Interfaces) == 0 {
			return protoreflect.ValueOfList(&_CodecDescriptor_1_list{})
		}
		listValue := &_CodecDescriptor_1_list{list: &x.Interfaces}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.CodecDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.CodecDescriptor does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_CodecDescriptor) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.CodecDescriptor.interfaces":
		lv := value.List()
		clv := lv.(*_CodecDescriptor_1_list)
		x.Interfaces = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.CodecDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.CodecDescriptor does not contain field %s", fd.FullName()))
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
func (x *fastReflection_CodecDescriptor) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.CodecDescriptor.interfaces":
		if x.Interfaces == nil {
			x.Interfaces = []*InterfaceDescriptor{}
		}
		value := &_CodecDescriptor_1_list{list: &x.Interfaces}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.CodecDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.CodecDescriptor does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_CodecDescriptor) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.CodecDescriptor.interfaces":
		list := []*InterfaceDescriptor{}
		return protoreflect.ValueOfList(&_CodecDescriptor_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.CodecDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.CodecDescriptor does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_CodecDescriptor) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.CodecDescriptor", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_CodecDescriptor) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CodecDescriptor) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_CodecDescriptor) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_CodecDescriptor) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*CodecDescriptor)
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
		if len(x.Interfaces) > 0 {
			for _, e := range x.Interfaces {
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
		x := input.Message.Interface().(*CodecDescriptor)
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
		if len(x.Interfaces) > 0 {
			for iNdEx := len(x.Interfaces) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Interfaces[iNdEx])
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
		x := input.Message.Interface().(*CodecDescriptor)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: CodecDescriptor: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: CodecDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Interfaces", wireType)
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
				x.Interfaces = append(x.Interfaces, &InterfaceDescriptor{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Interfaces[len(x.Interfaces)-1]); err != nil {
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

var _ protoreflect.List = (*_InterfaceDescriptor_2_list)(nil)

type _InterfaceDescriptor_2_list struct {
	list *[]*InterfaceAcceptingMessageDescriptor
}

func (x *_InterfaceDescriptor_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_InterfaceDescriptor_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_InterfaceDescriptor_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*InterfaceAcceptingMessageDescriptor)
	(*x.list)[i] = concreteValue
}

func (x *_InterfaceDescriptor_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*InterfaceAcceptingMessageDescriptor)
	*x.list = append(*x.list, concreteValue)
}

func (x *_InterfaceDescriptor_2_list) AppendMutable() protoreflect.Value {
	v := new(InterfaceAcceptingMessageDescriptor)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_InterfaceDescriptor_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_InterfaceDescriptor_2_list) NewElement() protoreflect.Value {
	v := new(InterfaceAcceptingMessageDescriptor)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_InterfaceDescriptor_2_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_InterfaceDescriptor_3_list)(nil)

type _InterfaceDescriptor_3_list struct {
	list *[]*InterfaceImplementerDescriptor
}

func (x *_InterfaceDescriptor_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_InterfaceDescriptor_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_InterfaceDescriptor_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*InterfaceImplementerDescriptor)
	(*x.list)[i] = concreteValue
}

func (x *_InterfaceDescriptor_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*InterfaceImplementerDescriptor)
	*x.list = append(*x.list, concreteValue)
}

func (x *_InterfaceDescriptor_3_list) AppendMutable() protoreflect.Value {
	v := new(InterfaceImplementerDescriptor)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_InterfaceDescriptor_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_InterfaceDescriptor_3_list) NewElement() protoreflect.Value {
	v := new(InterfaceImplementerDescriptor)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_InterfaceDescriptor_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_InterfaceDescriptor                              protoreflect.MessageDescriptor
	fd_InterfaceDescriptor_fullname                     protoreflect.FieldDescriptor
	fd_InterfaceDescriptor_interface_accepting_messages protoreflect.FieldDescriptor
	fd_InterfaceDescriptor_interface_implementers       protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_InterfaceDescriptor = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("InterfaceDescriptor")
	fd_InterfaceDescriptor_fullname = md_InterfaceDescriptor.Fields().ByName("fullname")
	fd_InterfaceDescriptor_interface_accepting_messages = md_InterfaceDescriptor.Fields().ByName("interface_accepting_messages")
	fd_InterfaceDescriptor_interface_implementers = md_InterfaceDescriptor.Fields().ByName("interface_implementers")
}

var _ protoreflect.Message = (*fastReflection_InterfaceDescriptor)(nil)

type fastReflection_InterfaceDescriptor InterfaceDescriptor

func (x *InterfaceDescriptor) ProtoReflect() protoreflect.Message {
	return (*fastReflection_InterfaceDescriptor)(x)
}

func (x *InterfaceDescriptor) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_InterfaceDescriptor_messageType fastReflection_InterfaceDescriptor_messageType
var _ protoreflect.MessageType = fastReflection_InterfaceDescriptor_messageType{}

type fastReflection_InterfaceDescriptor_messageType struct{}

func (x fastReflection_InterfaceDescriptor_messageType) Zero() protoreflect.Message {
	return (*fastReflection_InterfaceDescriptor)(nil)
}
func (x fastReflection_InterfaceDescriptor_messageType) New() protoreflect.Message {
	return new(fastReflection_InterfaceDescriptor)
}
func (x fastReflection_InterfaceDescriptor_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_InterfaceDescriptor
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_InterfaceDescriptor) Descriptor() protoreflect.MessageDescriptor {
	return md_InterfaceDescriptor
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_InterfaceDescriptor) Type() protoreflect.MessageType {
	return _fastReflection_InterfaceDescriptor_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_InterfaceDescriptor) New() protoreflect.Message {
	return new(fastReflection_InterfaceDescriptor)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_InterfaceDescriptor) Interface() protoreflect.ProtoMessage {
	return (*InterfaceDescriptor)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_InterfaceDescriptor) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Fullname != "" {
		value := protoreflect.ValueOfString(x.Fullname)
		if !f(fd_InterfaceDescriptor_fullname, value) {
			return
		}
	}
	if len(x.InterfaceAcceptingMessages) != 0 {
		value := protoreflect.ValueOfList(&_InterfaceDescriptor_2_list{list: &x.InterfaceAcceptingMessages})
		if !f(fd_InterfaceDescriptor_interface_accepting_messages, value) {
			return
		}
	}
	if len(x.InterfaceImplementers) != 0 {
		value := protoreflect.ValueOfList(&_InterfaceDescriptor_3_list{list: &x.InterfaceImplementers})
		if !f(fd_InterfaceDescriptor_interface_implementers, value) {
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
func (x *fastReflection_InterfaceDescriptor) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.InterfaceDescriptor.fullname":
		return x.Fullname != ""
	case "cosmos.base.reflection.v2alpha1.InterfaceDescriptor.interface_accepting_messages":
		return len(x.InterfaceAcceptingMessages) != 0
	case "cosmos.base.reflection.v2alpha1.InterfaceDescriptor.interface_implementers":
		return len(x.InterfaceImplementers) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.InterfaceDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.InterfaceDescriptor does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_InterfaceDescriptor) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.InterfaceDescriptor.fullname":
		x.Fullname = ""
	case "cosmos.base.reflection.v2alpha1.InterfaceDescriptor.interface_accepting_messages":
		x.InterfaceAcceptingMessages = nil
	case "cosmos.base.reflection.v2alpha1.InterfaceDescriptor.interface_implementers":
		x.InterfaceImplementers = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.InterfaceDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.InterfaceDescriptor does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_InterfaceDescriptor) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v2alpha1.InterfaceDescriptor.fullname":
		value := x.Fullname
		return protoreflect.ValueOfString(value)
	case "cosmos.base.reflection.v2alpha1.InterfaceDescriptor.interface_accepting_messages":
		if len(x.InterfaceAcceptingMessages) == 0 {
			return protoreflect.ValueOfList(&_InterfaceDescriptor_2_list{})
		}
		listValue := &_InterfaceDescriptor_2_list{list: &x.InterfaceAcceptingMessages}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.base.reflection.v2alpha1.InterfaceDescriptor.interface_implementers":
		if len(x.InterfaceImplementers) == 0 {
			return protoreflect.ValueOfList(&_InterfaceDescriptor_3_list{})
		}
		listValue := &_InterfaceDescriptor_3_list{list: &x.InterfaceImplementers}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.InterfaceDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.InterfaceDescriptor does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_InterfaceDescriptor) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.InterfaceDescriptor.fullname":
		x.Fullname = value.Interface().(string)
	case "cosmos.base.reflection.v2alpha1.InterfaceDescriptor.interface_accepting_messages":
		lv := value.List()
		clv := lv.(*_InterfaceDescriptor_2_list)
		x.InterfaceAcceptingMessages = *clv.list
	case "cosmos.base.reflection.v2alpha1.InterfaceDescriptor.interface_implementers":
		lv := value.List()
		clv := lv.(*_InterfaceDescriptor_3_list)
		x.InterfaceImplementers = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.InterfaceDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.InterfaceDescriptor does not contain field %s", fd.FullName()))
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
func (x *fastReflection_InterfaceDescriptor) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.InterfaceDescriptor.interface_accepting_messages":
		if x.InterfaceAcceptingMessages == nil {
			x.InterfaceAcceptingMessages = []*InterfaceAcceptingMessageDescriptor{}
		}
		value := &_InterfaceDescriptor_2_list{list: &x.InterfaceAcceptingMessages}
		return protoreflect.ValueOfList(value)
	case "cosmos.base.reflection.v2alpha1.InterfaceDescriptor.interface_implementers":
		if x.InterfaceImplementers == nil {
			x.InterfaceImplementers = []*InterfaceImplementerDescriptor{}
		}
		value := &_InterfaceDescriptor_3_list{list: &x.InterfaceImplementers}
		return protoreflect.ValueOfList(value)
	case "cosmos.base.reflection.v2alpha1.InterfaceDescriptor.fullname":
		panic(fmt.Errorf("field fullname of message cosmos.base.reflection.v2alpha1.InterfaceDescriptor is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.InterfaceDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.InterfaceDescriptor does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_InterfaceDescriptor) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.InterfaceDescriptor.fullname":
		return protoreflect.ValueOfString("")
	case "cosmos.base.reflection.v2alpha1.InterfaceDescriptor.interface_accepting_messages":
		list := []*InterfaceAcceptingMessageDescriptor{}
		return protoreflect.ValueOfList(&_InterfaceDescriptor_2_list{list: &list})
	case "cosmos.base.reflection.v2alpha1.InterfaceDescriptor.interface_implementers":
		list := []*InterfaceImplementerDescriptor{}
		return protoreflect.ValueOfList(&_InterfaceDescriptor_3_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.InterfaceDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.InterfaceDescriptor does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_InterfaceDescriptor) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.InterfaceDescriptor", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_InterfaceDescriptor) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_InterfaceDescriptor) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_InterfaceDescriptor) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_InterfaceDescriptor) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*InterfaceDescriptor)
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
		l = len(x.Fullname)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.InterfaceAcceptingMessages) > 0 {
			for _, e := range x.InterfaceAcceptingMessages {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.InterfaceImplementers) > 0 {
			for _, e := range x.InterfaceImplementers {
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
		x := input.Message.Interface().(*InterfaceDescriptor)
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
		if len(x.InterfaceImplementers) > 0 {
			for iNdEx := len(x.InterfaceImplementers) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.InterfaceImplementers[iNdEx])
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
		if len(x.InterfaceAcceptingMessages) > 0 {
			for iNdEx := len(x.InterfaceAcceptingMessages) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.InterfaceAcceptingMessages[iNdEx])
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
		if len(x.Fullname) > 0 {
			i -= len(x.Fullname)
			copy(dAtA[i:], x.Fullname)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Fullname)))
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
		x := input.Message.Interface().(*InterfaceDescriptor)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: InterfaceDescriptor: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: InterfaceDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Fullname", wireType)
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
				x.Fullname = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field InterfaceAcceptingMessages", wireType)
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
				x.InterfaceAcceptingMessages = append(x.InterfaceAcceptingMessages, &InterfaceAcceptingMessageDescriptor{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.InterfaceAcceptingMessages[len(x.InterfaceAcceptingMessages)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field InterfaceImplementers", wireType)
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
				x.InterfaceImplementers = append(x.InterfaceImplementers, &InterfaceImplementerDescriptor{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.InterfaceImplementers[len(x.InterfaceImplementers)-1]); err != nil {
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
	md_InterfaceImplementerDescriptor          protoreflect.MessageDescriptor
	fd_InterfaceImplementerDescriptor_fullname protoreflect.FieldDescriptor
	fd_InterfaceImplementerDescriptor_type_url protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_InterfaceImplementerDescriptor = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("InterfaceImplementerDescriptor")
	fd_InterfaceImplementerDescriptor_fullname = md_InterfaceImplementerDescriptor.Fields().ByName("fullname")
	fd_InterfaceImplementerDescriptor_type_url = md_InterfaceImplementerDescriptor.Fields().ByName("type_url")
}

var _ protoreflect.Message = (*fastReflection_InterfaceImplementerDescriptor)(nil)

type fastReflection_InterfaceImplementerDescriptor InterfaceImplementerDescriptor

func (x *InterfaceImplementerDescriptor) ProtoReflect() protoreflect.Message {
	return (*fastReflection_InterfaceImplementerDescriptor)(x)
}

func (x *InterfaceImplementerDescriptor) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_InterfaceImplementerDescriptor_messageType fastReflection_InterfaceImplementerDescriptor_messageType
var _ protoreflect.MessageType = fastReflection_InterfaceImplementerDescriptor_messageType{}

type fastReflection_InterfaceImplementerDescriptor_messageType struct{}

func (x fastReflection_InterfaceImplementerDescriptor_messageType) Zero() protoreflect.Message {
	return (*fastReflection_InterfaceImplementerDescriptor)(nil)
}
func (x fastReflection_InterfaceImplementerDescriptor_messageType) New() protoreflect.Message {
	return new(fastReflection_InterfaceImplementerDescriptor)
}
func (x fastReflection_InterfaceImplementerDescriptor_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_InterfaceImplementerDescriptor
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_InterfaceImplementerDescriptor) Descriptor() protoreflect.MessageDescriptor {
	return md_InterfaceImplementerDescriptor
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_InterfaceImplementerDescriptor) Type() protoreflect.MessageType {
	return _fastReflection_InterfaceImplementerDescriptor_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_InterfaceImplementerDescriptor) New() protoreflect.Message {
	return new(fastReflection_InterfaceImplementerDescriptor)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_InterfaceImplementerDescriptor) Interface() protoreflect.ProtoMessage {
	return (*InterfaceImplementerDescriptor)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_InterfaceImplementerDescriptor) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Fullname != "" {
		value := protoreflect.ValueOfString(x.Fullname)
		if !f(fd_InterfaceImplementerDescriptor_fullname, value) {
			return
		}
	}
	if x.TypeUrl != "" {
		value := protoreflect.ValueOfString(x.TypeUrl)
		if !f(fd_InterfaceImplementerDescriptor_type_url, value) {
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
func (x *fastReflection_InterfaceImplementerDescriptor) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor.fullname":
		return x.Fullname != ""
	case "cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor.type_url":
		return x.TypeUrl != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_InterfaceImplementerDescriptor) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor.fullname":
		x.Fullname = ""
	case "cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor.type_url":
		x.TypeUrl = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_InterfaceImplementerDescriptor) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor.fullname":
		value := x.Fullname
		return protoreflect.ValueOfString(value)
	case "cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor.type_url":
		value := x.TypeUrl
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_InterfaceImplementerDescriptor) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor.fullname":
		x.Fullname = value.Interface().(string)
	case "cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor.type_url":
		x.TypeUrl = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor does not contain field %s", fd.FullName()))
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
func (x *fastReflection_InterfaceImplementerDescriptor) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor.fullname":
		panic(fmt.Errorf("field fullname of message cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor is not mutable"))
	case "cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor.type_url":
		panic(fmt.Errorf("field type_url of message cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_InterfaceImplementerDescriptor) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor.fullname":
		return protoreflect.ValueOfString("")
	case "cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor.type_url":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_InterfaceImplementerDescriptor) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_InterfaceImplementerDescriptor) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_InterfaceImplementerDescriptor) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_InterfaceImplementerDescriptor) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_InterfaceImplementerDescriptor) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*InterfaceImplementerDescriptor)
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
		l = len(x.Fullname)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.TypeUrl)
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
		x := input.Message.Interface().(*InterfaceImplementerDescriptor)
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
		if len(x.TypeUrl) > 0 {
			i -= len(x.TypeUrl)
			copy(dAtA[i:], x.TypeUrl)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TypeUrl)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Fullname) > 0 {
			i -= len(x.Fullname)
			copy(dAtA[i:], x.Fullname)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Fullname)))
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
		x := input.Message.Interface().(*InterfaceImplementerDescriptor)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: InterfaceImplementerDescriptor: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: InterfaceImplementerDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Fullname", wireType)
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
				x.Fullname = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TypeUrl", wireType)
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
				x.TypeUrl = string(dAtA[iNdEx:postIndex])
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

var _ protoreflect.List = (*_InterfaceAcceptingMessageDescriptor_2_list)(nil)

type _InterfaceAcceptingMessageDescriptor_2_list struct {
	list *[]string
}

func (x *_InterfaceAcceptingMessageDescriptor_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_InterfaceAcceptingMessageDescriptor_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_InterfaceAcceptingMessageDescriptor_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_InterfaceAcceptingMessageDescriptor_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_InterfaceAcceptingMessageDescriptor_2_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message InterfaceAcceptingMessageDescriptor at list field FieldDescriptorNames as it is not of Message kind"))
}

func (x *_InterfaceAcceptingMessageDescriptor_2_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_InterfaceAcceptingMessageDescriptor_2_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_InterfaceAcceptingMessageDescriptor_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_InterfaceAcceptingMessageDescriptor                        protoreflect.MessageDescriptor
	fd_InterfaceAcceptingMessageDescriptor_fullname               protoreflect.FieldDescriptor
	fd_InterfaceAcceptingMessageDescriptor_field_descriptor_names protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_InterfaceAcceptingMessageDescriptor = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("InterfaceAcceptingMessageDescriptor")
	fd_InterfaceAcceptingMessageDescriptor_fullname = md_InterfaceAcceptingMessageDescriptor.Fields().ByName("fullname")
	fd_InterfaceAcceptingMessageDescriptor_field_descriptor_names = md_InterfaceAcceptingMessageDescriptor.Fields().ByName("field_descriptor_names")
}

var _ protoreflect.Message = (*fastReflection_InterfaceAcceptingMessageDescriptor)(nil)

type fastReflection_InterfaceAcceptingMessageDescriptor InterfaceAcceptingMessageDescriptor

func (x *InterfaceAcceptingMessageDescriptor) ProtoReflect() protoreflect.Message {
	return (*fastReflection_InterfaceAcceptingMessageDescriptor)(x)
}

func (x *InterfaceAcceptingMessageDescriptor) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_InterfaceAcceptingMessageDescriptor_messageType fastReflection_InterfaceAcceptingMessageDescriptor_messageType
var _ protoreflect.MessageType = fastReflection_InterfaceAcceptingMessageDescriptor_messageType{}

type fastReflection_InterfaceAcceptingMessageDescriptor_messageType struct{}

func (x fastReflection_InterfaceAcceptingMessageDescriptor_messageType) Zero() protoreflect.Message {
	return (*fastReflection_InterfaceAcceptingMessageDescriptor)(nil)
}
func (x fastReflection_InterfaceAcceptingMessageDescriptor_messageType) New() protoreflect.Message {
	return new(fastReflection_InterfaceAcceptingMessageDescriptor)
}
func (x fastReflection_InterfaceAcceptingMessageDescriptor_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_InterfaceAcceptingMessageDescriptor
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_InterfaceAcceptingMessageDescriptor) Descriptor() protoreflect.MessageDescriptor {
	return md_InterfaceAcceptingMessageDescriptor
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_InterfaceAcceptingMessageDescriptor) Type() protoreflect.MessageType {
	return _fastReflection_InterfaceAcceptingMessageDescriptor_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_InterfaceAcceptingMessageDescriptor) New() protoreflect.Message {
	return new(fastReflection_InterfaceAcceptingMessageDescriptor)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_InterfaceAcceptingMessageDescriptor) Interface() protoreflect.ProtoMessage {
	return (*InterfaceAcceptingMessageDescriptor)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_InterfaceAcceptingMessageDescriptor) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Fullname != "" {
		value := protoreflect.ValueOfString(x.Fullname)
		if !f(fd_InterfaceAcceptingMessageDescriptor_fullname, value) {
			return
		}
	}
	if len(x.FieldDescriptorNames) != 0 {
		value := protoreflect.ValueOfList(&_InterfaceAcceptingMessageDescriptor_2_list{list: &x.FieldDescriptorNames})
		if !f(fd_InterfaceAcceptingMessageDescriptor_field_descriptor_names, value) {
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
func (x *fastReflection_InterfaceAcceptingMessageDescriptor) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor.fullname":
		return x.Fullname != ""
	case "cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor.field_descriptor_names":
		return len(x.FieldDescriptorNames) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_InterfaceAcceptingMessageDescriptor) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor.fullname":
		x.Fullname = ""
	case "cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor.field_descriptor_names":
		x.FieldDescriptorNames = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_InterfaceAcceptingMessageDescriptor) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor.fullname":
		value := x.Fullname
		return protoreflect.ValueOfString(value)
	case "cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor.field_descriptor_names":
		if len(x.FieldDescriptorNames) == 0 {
			return protoreflect.ValueOfList(&_InterfaceAcceptingMessageDescriptor_2_list{})
		}
		listValue := &_InterfaceAcceptingMessageDescriptor_2_list{list: &x.FieldDescriptorNames}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_InterfaceAcceptingMessageDescriptor) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor.fullname":
		x.Fullname = value.Interface().(string)
	case "cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor.field_descriptor_names":
		lv := value.List()
		clv := lv.(*_InterfaceAcceptingMessageDescriptor_2_list)
		x.FieldDescriptorNames = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor does not contain field %s", fd.FullName()))
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
func (x *fastReflection_InterfaceAcceptingMessageDescriptor) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor.field_descriptor_names":
		if x.FieldDescriptorNames == nil {
			x.FieldDescriptorNames = []string{}
		}
		value := &_InterfaceAcceptingMessageDescriptor_2_list{list: &x.FieldDescriptorNames}
		return protoreflect.ValueOfList(value)
	case "cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor.fullname":
		panic(fmt.Errorf("field fullname of message cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_InterfaceAcceptingMessageDescriptor) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor.fullname":
		return protoreflect.ValueOfString("")
	case "cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor.field_descriptor_names":
		list := []string{}
		return protoreflect.ValueOfList(&_InterfaceAcceptingMessageDescriptor_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_InterfaceAcceptingMessageDescriptor) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_InterfaceAcceptingMessageDescriptor) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_InterfaceAcceptingMessageDescriptor) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_InterfaceAcceptingMessageDescriptor) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_InterfaceAcceptingMessageDescriptor) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*InterfaceAcceptingMessageDescriptor)
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
		l = len(x.Fullname)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.FieldDescriptorNames) > 0 {
			for _, s := range x.FieldDescriptorNames {
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
		x := input.Message.Interface().(*InterfaceAcceptingMessageDescriptor)
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
		if len(x.FieldDescriptorNames) > 0 {
			for iNdEx := len(x.FieldDescriptorNames) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.FieldDescriptorNames[iNdEx])
				copy(dAtA[i:], x.FieldDescriptorNames[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.FieldDescriptorNames[iNdEx])))
				i--
				dAtA[i] = 0x12
			}
		}
		if len(x.Fullname) > 0 {
			i -= len(x.Fullname)
			copy(dAtA[i:], x.Fullname)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Fullname)))
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
		x := input.Message.Interface().(*InterfaceAcceptingMessageDescriptor)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: InterfaceAcceptingMessageDescriptor: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: InterfaceAcceptingMessageDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Fullname", wireType)
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
				x.Fullname = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field FieldDescriptorNames", wireType)
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
				x.FieldDescriptorNames = append(x.FieldDescriptorNames, string(dAtA[iNdEx:postIndex]))
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
	md_ConfigurationDescriptor                               protoreflect.MessageDescriptor
	fd_ConfigurationDescriptor_bech32_account_address_prefix protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_ConfigurationDescriptor = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("ConfigurationDescriptor")
	fd_ConfigurationDescriptor_bech32_account_address_prefix = md_ConfigurationDescriptor.Fields().ByName("bech32_account_address_prefix")
}

var _ protoreflect.Message = (*fastReflection_ConfigurationDescriptor)(nil)

type fastReflection_ConfigurationDescriptor ConfigurationDescriptor

func (x *ConfigurationDescriptor) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ConfigurationDescriptor)(x)
}

func (x *ConfigurationDescriptor) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ConfigurationDescriptor_messageType fastReflection_ConfigurationDescriptor_messageType
var _ protoreflect.MessageType = fastReflection_ConfigurationDescriptor_messageType{}

type fastReflection_ConfigurationDescriptor_messageType struct{}

func (x fastReflection_ConfigurationDescriptor_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ConfigurationDescriptor)(nil)
}
func (x fastReflection_ConfigurationDescriptor_messageType) New() protoreflect.Message {
	return new(fastReflection_ConfigurationDescriptor)
}
func (x fastReflection_ConfigurationDescriptor_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ConfigurationDescriptor
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ConfigurationDescriptor) Descriptor() protoreflect.MessageDescriptor {
	return md_ConfigurationDescriptor
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ConfigurationDescriptor) Type() protoreflect.MessageType {
	return _fastReflection_ConfigurationDescriptor_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ConfigurationDescriptor) New() protoreflect.Message {
	return new(fastReflection_ConfigurationDescriptor)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ConfigurationDescriptor) Interface() protoreflect.ProtoMessage {
	return (*ConfigurationDescriptor)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ConfigurationDescriptor) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Bech32AccountAddressPrefix != "" {
		value := protoreflect.ValueOfString(x.Bech32AccountAddressPrefix)
		if !f(fd_ConfigurationDescriptor_bech32_account_address_prefix, value) {
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
func (x *fastReflection_ConfigurationDescriptor) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.ConfigurationDescriptor.bech32_account_address_prefix":
		return x.Bech32AccountAddressPrefix != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.ConfigurationDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.ConfigurationDescriptor does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ConfigurationDescriptor) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.ConfigurationDescriptor.bech32_account_address_prefix":
		x.Bech32AccountAddressPrefix = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.ConfigurationDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.ConfigurationDescriptor does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ConfigurationDescriptor) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v2alpha1.ConfigurationDescriptor.bech32_account_address_prefix":
		value := x.Bech32AccountAddressPrefix
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.ConfigurationDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.ConfigurationDescriptor does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ConfigurationDescriptor) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.ConfigurationDescriptor.bech32_account_address_prefix":
		x.Bech32AccountAddressPrefix = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.ConfigurationDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.ConfigurationDescriptor does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ConfigurationDescriptor) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.ConfigurationDescriptor.bech32_account_address_prefix":
		panic(fmt.Errorf("field bech32_account_address_prefix of message cosmos.base.reflection.v2alpha1.ConfigurationDescriptor is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.ConfigurationDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.ConfigurationDescriptor does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ConfigurationDescriptor) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.ConfigurationDescriptor.bech32_account_address_prefix":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.ConfigurationDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.ConfigurationDescriptor does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ConfigurationDescriptor) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.ConfigurationDescriptor", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ConfigurationDescriptor) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ConfigurationDescriptor) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ConfigurationDescriptor) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ConfigurationDescriptor) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ConfigurationDescriptor)
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
		l = len(x.Bech32AccountAddressPrefix)
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
		x := input.Message.Interface().(*ConfigurationDescriptor)
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
		if len(x.Bech32AccountAddressPrefix) > 0 {
			i -= len(x.Bech32AccountAddressPrefix)
			copy(dAtA[i:], x.Bech32AccountAddressPrefix)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Bech32AccountAddressPrefix)))
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
		x := input.Message.Interface().(*ConfigurationDescriptor)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ConfigurationDescriptor: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ConfigurationDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Bech32AccountAddressPrefix", wireType)
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
				x.Bech32AccountAddressPrefix = string(dAtA[iNdEx:postIndex])
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
	md_MsgDescriptor              protoreflect.MessageDescriptor
	fd_MsgDescriptor_msg_type_url protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_MsgDescriptor = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("MsgDescriptor")
	fd_MsgDescriptor_msg_type_url = md_MsgDescriptor.Fields().ByName("msg_type_url")
}

var _ protoreflect.Message = (*fastReflection_MsgDescriptor)(nil)

type fastReflection_MsgDescriptor MsgDescriptor

func (x *MsgDescriptor) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgDescriptor)(x)
}

func (x *MsgDescriptor) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgDescriptor_messageType fastReflection_MsgDescriptor_messageType
var _ protoreflect.MessageType = fastReflection_MsgDescriptor_messageType{}

type fastReflection_MsgDescriptor_messageType struct{}

func (x fastReflection_MsgDescriptor_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgDescriptor)(nil)
}
func (x fastReflection_MsgDescriptor_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgDescriptor)
}
func (x fastReflection_MsgDescriptor_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgDescriptor
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgDescriptor) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgDescriptor
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgDescriptor) Type() protoreflect.MessageType {
	return _fastReflection_MsgDescriptor_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgDescriptor) New() protoreflect.Message {
	return new(fastReflection_MsgDescriptor)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgDescriptor) Interface() protoreflect.ProtoMessage {
	return (*MsgDescriptor)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgDescriptor) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.MsgTypeUrl != "" {
		value := protoreflect.ValueOfString(x.MsgTypeUrl)
		if !f(fd_MsgDescriptor_msg_type_url, value) {
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
func (x *fastReflection_MsgDescriptor) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.MsgDescriptor.msg_type_url":
		return x.MsgTypeUrl != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.MsgDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.MsgDescriptor does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgDescriptor) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.MsgDescriptor.msg_type_url":
		x.MsgTypeUrl = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.MsgDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.MsgDescriptor does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgDescriptor) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v2alpha1.MsgDescriptor.msg_type_url":
		value := x.MsgTypeUrl
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.MsgDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.MsgDescriptor does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgDescriptor) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.MsgDescriptor.msg_type_url":
		x.MsgTypeUrl = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.MsgDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.MsgDescriptor does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgDescriptor) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.MsgDescriptor.msg_type_url":
		panic(fmt.Errorf("field msg_type_url of message cosmos.base.reflection.v2alpha1.MsgDescriptor is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.MsgDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.MsgDescriptor does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgDescriptor) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.MsgDescriptor.msg_type_url":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.MsgDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.MsgDescriptor does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgDescriptor) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.MsgDescriptor", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgDescriptor) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgDescriptor) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgDescriptor) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgDescriptor) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgDescriptor)
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
		l = len(x.MsgTypeUrl)
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
		x := input.Message.Interface().(*MsgDescriptor)
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
		if len(x.MsgTypeUrl) > 0 {
			i -= len(x.MsgTypeUrl)
			copy(dAtA[i:], x.MsgTypeUrl)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.MsgTypeUrl)))
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
		x := input.Message.Interface().(*MsgDescriptor)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgDescriptor: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrl", wireType)
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
				x.MsgTypeUrl = string(dAtA[iNdEx:postIndex])
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
	md_GetAuthnDescriptorRequest protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_GetAuthnDescriptorRequest = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("GetAuthnDescriptorRequest")
}

var _ protoreflect.Message = (*fastReflection_GetAuthnDescriptorRequest)(nil)

type fastReflection_GetAuthnDescriptorRequest GetAuthnDescriptorRequest

func (x *GetAuthnDescriptorRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetAuthnDescriptorRequest)(x)
}

func (x *GetAuthnDescriptorRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetAuthnDescriptorRequest_messageType fastReflection_GetAuthnDescriptorRequest_messageType
var _ protoreflect.MessageType = fastReflection_GetAuthnDescriptorRequest_messageType{}

type fastReflection_GetAuthnDescriptorRequest_messageType struct{}

func (x fastReflection_GetAuthnDescriptorRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetAuthnDescriptorRequest)(nil)
}
func (x fastReflection_GetAuthnDescriptorRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_GetAuthnDescriptorRequest)
}
func (x fastReflection_GetAuthnDescriptorRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetAuthnDescriptorRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetAuthnDescriptorRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_GetAuthnDescriptorRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetAuthnDescriptorRequest) Type() protoreflect.MessageType {
	return _fastReflection_GetAuthnDescriptorRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetAuthnDescriptorRequest) New() protoreflect.Message {
	return new(fastReflection_GetAuthnDescriptorRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetAuthnDescriptorRequest) Interface() protoreflect.ProtoMessage {
	return (*GetAuthnDescriptorRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetAuthnDescriptorRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_GetAuthnDescriptorRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetAuthnDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetAuthnDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetAuthnDescriptorRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetAuthnDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetAuthnDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetAuthnDescriptorRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetAuthnDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetAuthnDescriptorRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_GetAuthnDescriptorRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetAuthnDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetAuthnDescriptorRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_GetAuthnDescriptorRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetAuthnDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetAuthnDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetAuthnDescriptorRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetAuthnDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetAuthnDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetAuthnDescriptorRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.GetAuthnDescriptorRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetAuthnDescriptorRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetAuthnDescriptorRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_GetAuthnDescriptorRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetAuthnDescriptorRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetAuthnDescriptorRequest)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GetAuthnDescriptorRequest)
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
		x := input.Message.Interface().(*GetAuthnDescriptorRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetAuthnDescriptorRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetAuthnDescriptorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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
	md_GetAuthnDescriptorResponse       protoreflect.MessageDescriptor
	fd_GetAuthnDescriptorResponse_authn protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_GetAuthnDescriptorResponse = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("GetAuthnDescriptorResponse")
	fd_GetAuthnDescriptorResponse_authn = md_GetAuthnDescriptorResponse.Fields().ByName("authn")
}

var _ protoreflect.Message = (*fastReflection_GetAuthnDescriptorResponse)(nil)

type fastReflection_GetAuthnDescriptorResponse GetAuthnDescriptorResponse

func (x *GetAuthnDescriptorResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetAuthnDescriptorResponse)(x)
}

func (x *GetAuthnDescriptorResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetAuthnDescriptorResponse_messageType fastReflection_GetAuthnDescriptorResponse_messageType
var _ protoreflect.MessageType = fastReflection_GetAuthnDescriptorResponse_messageType{}

type fastReflection_GetAuthnDescriptorResponse_messageType struct{}

func (x fastReflection_GetAuthnDescriptorResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetAuthnDescriptorResponse)(nil)
}
func (x fastReflection_GetAuthnDescriptorResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_GetAuthnDescriptorResponse)
}
func (x fastReflection_GetAuthnDescriptorResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetAuthnDescriptorResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetAuthnDescriptorResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_GetAuthnDescriptorResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetAuthnDescriptorResponse) Type() protoreflect.MessageType {
	return _fastReflection_GetAuthnDescriptorResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetAuthnDescriptorResponse) New() protoreflect.Message {
	return new(fastReflection_GetAuthnDescriptorResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetAuthnDescriptorResponse) Interface() protoreflect.ProtoMessage {
	return (*GetAuthnDescriptorResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetAuthnDescriptorResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Authn != nil {
		value := protoreflect.ValueOfMessage(x.Authn.ProtoReflect())
		if !f(fd_GetAuthnDescriptorResponse_authn, value) {
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
func (x *fastReflection_GetAuthnDescriptorResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse.authn":
		return x.Authn != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetAuthnDescriptorResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse.authn":
		x.Authn = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetAuthnDescriptorResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse.authn":
		value := x.Authn
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_GetAuthnDescriptorResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse.authn":
		x.Authn = value.Message().Interface().(*AuthnDescriptor)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_GetAuthnDescriptorResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse.authn":
		if x.Authn == nil {
			x.Authn = new(AuthnDescriptor)
		}
		return protoreflect.ValueOfMessage(x.Authn.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetAuthnDescriptorResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse.authn":
		m := new(AuthnDescriptor)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetAuthnDescriptorResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetAuthnDescriptorResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetAuthnDescriptorResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_GetAuthnDescriptorResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetAuthnDescriptorResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetAuthnDescriptorResponse)
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
		if x.Authn != nil {
			l = options.Size(x.Authn)
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
		x := input.Message.Interface().(*GetAuthnDescriptorResponse)
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
		if x.Authn != nil {
			encoded, err := options.Marshal(x.Authn)
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
		x := input.Message.Interface().(*GetAuthnDescriptorResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetAuthnDescriptorResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetAuthnDescriptorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Authn", wireType)
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
				if x.Authn == nil {
					x.Authn = &AuthnDescriptor{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Authn); err != nil {
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
	md_GetChainDescriptorRequest protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_GetChainDescriptorRequest = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("GetChainDescriptorRequest")
}

var _ protoreflect.Message = (*fastReflection_GetChainDescriptorRequest)(nil)

type fastReflection_GetChainDescriptorRequest GetChainDescriptorRequest

func (x *GetChainDescriptorRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetChainDescriptorRequest)(x)
}

func (x *GetChainDescriptorRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetChainDescriptorRequest_messageType fastReflection_GetChainDescriptorRequest_messageType
var _ protoreflect.MessageType = fastReflection_GetChainDescriptorRequest_messageType{}

type fastReflection_GetChainDescriptorRequest_messageType struct{}

func (x fastReflection_GetChainDescriptorRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetChainDescriptorRequest)(nil)
}
func (x fastReflection_GetChainDescriptorRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_GetChainDescriptorRequest)
}
func (x fastReflection_GetChainDescriptorRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetChainDescriptorRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetChainDescriptorRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_GetChainDescriptorRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetChainDescriptorRequest) Type() protoreflect.MessageType {
	return _fastReflection_GetChainDescriptorRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetChainDescriptorRequest) New() protoreflect.Message {
	return new(fastReflection_GetChainDescriptorRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetChainDescriptorRequest) Interface() protoreflect.ProtoMessage {
	return (*GetChainDescriptorRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetChainDescriptorRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_GetChainDescriptorRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetChainDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetChainDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetChainDescriptorRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetChainDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetChainDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetChainDescriptorRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetChainDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetChainDescriptorRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_GetChainDescriptorRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetChainDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetChainDescriptorRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_GetChainDescriptorRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetChainDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetChainDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetChainDescriptorRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetChainDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetChainDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetChainDescriptorRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.GetChainDescriptorRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetChainDescriptorRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetChainDescriptorRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_GetChainDescriptorRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetChainDescriptorRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetChainDescriptorRequest)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GetChainDescriptorRequest)
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
		x := input.Message.Interface().(*GetChainDescriptorRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetChainDescriptorRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetChainDescriptorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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
	md_GetChainDescriptorResponse       protoreflect.MessageDescriptor
	fd_GetChainDescriptorResponse_chain protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_GetChainDescriptorResponse = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("GetChainDescriptorResponse")
	fd_GetChainDescriptorResponse_chain = md_GetChainDescriptorResponse.Fields().ByName("chain")
}

var _ protoreflect.Message = (*fastReflection_GetChainDescriptorResponse)(nil)

type fastReflection_GetChainDescriptorResponse GetChainDescriptorResponse

func (x *GetChainDescriptorResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetChainDescriptorResponse)(x)
}

func (x *GetChainDescriptorResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetChainDescriptorResponse_messageType fastReflection_GetChainDescriptorResponse_messageType
var _ protoreflect.MessageType = fastReflection_GetChainDescriptorResponse_messageType{}

type fastReflection_GetChainDescriptorResponse_messageType struct{}

func (x fastReflection_GetChainDescriptorResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetChainDescriptorResponse)(nil)
}
func (x fastReflection_GetChainDescriptorResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_GetChainDescriptorResponse)
}
func (x fastReflection_GetChainDescriptorResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetChainDescriptorResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetChainDescriptorResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_GetChainDescriptorResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetChainDescriptorResponse) Type() protoreflect.MessageType {
	return _fastReflection_GetChainDescriptorResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetChainDescriptorResponse) New() protoreflect.Message {
	return new(fastReflection_GetChainDescriptorResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetChainDescriptorResponse) Interface() protoreflect.ProtoMessage {
	return (*GetChainDescriptorResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetChainDescriptorResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Chain != nil {
		value := protoreflect.ValueOfMessage(x.Chain.ProtoReflect())
		if !f(fd_GetChainDescriptorResponse_chain, value) {
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
func (x *fastReflection_GetChainDescriptorResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse.chain":
		return x.Chain != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetChainDescriptorResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse.chain":
		x.Chain = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetChainDescriptorResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse.chain":
		value := x.Chain
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_GetChainDescriptorResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse.chain":
		x.Chain = value.Message().Interface().(*ChainDescriptor)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_GetChainDescriptorResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse.chain":
		if x.Chain == nil {
			x.Chain = new(ChainDescriptor)
		}
		return protoreflect.ValueOfMessage(x.Chain.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetChainDescriptorResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse.chain":
		m := new(ChainDescriptor)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetChainDescriptorResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetChainDescriptorResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetChainDescriptorResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_GetChainDescriptorResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetChainDescriptorResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetChainDescriptorResponse)
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
		if x.Chain != nil {
			l = options.Size(x.Chain)
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
		x := input.Message.Interface().(*GetChainDescriptorResponse)
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
		if x.Chain != nil {
			encoded, err := options.Marshal(x.Chain)
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
		x := input.Message.Interface().(*GetChainDescriptorResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetChainDescriptorResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetChainDescriptorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
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
				if x.Chain == nil {
					x.Chain = &ChainDescriptor{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Chain); err != nil {
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
	md_GetCodecDescriptorRequest protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_GetCodecDescriptorRequest = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("GetCodecDescriptorRequest")
}

var _ protoreflect.Message = (*fastReflection_GetCodecDescriptorRequest)(nil)

type fastReflection_GetCodecDescriptorRequest GetCodecDescriptorRequest

func (x *GetCodecDescriptorRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetCodecDescriptorRequest)(x)
}

func (x *GetCodecDescriptorRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetCodecDescriptorRequest_messageType fastReflection_GetCodecDescriptorRequest_messageType
var _ protoreflect.MessageType = fastReflection_GetCodecDescriptorRequest_messageType{}

type fastReflection_GetCodecDescriptorRequest_messageType struct{}

func (x fastReflection_GetCodecDescriptorRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetCodecDescriptorRequest)(nil)
}
func (x fastReflection_GetCodecDescriptorRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_GetCodecDescriptorRequest)
}
func (x fastReflection_GetCodecDescriptorRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetCodecDescriptorRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetCodecDescriptorRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_GetCodecDescriptorRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetCodecDescriptorRequest) Type() protoreflect.MessageType {
	return _fastReflection_GetCodecDescriptorRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetCodecDescriptorRequest) New() protoreflect.Message {
	return new(fastReflection_GetCodecDescriptorRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetCodecDescriptorRequest) Interface() protoreflect.ProtoMessage {
	return (*GetCodecDescriptorRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetCodecDescriptorRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_GetCodecDescriptorRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetCodecDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetCodecDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetCodecDescriptorRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetCodecDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetCodecDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetCodecDescriptorRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetCodecDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetCodecDescriptorRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_GetCodecDescriptorRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetCodecDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetCodecDescriptorRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_GetCodecDescriptorRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetCodecDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetCodecDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetCodecDescriptorRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetCodecDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetCodecDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetCodecDescriptorRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.GetCodecDescriptorRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetCodecDescriptorRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetCodecDescriptorRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_GetCodecDescriptorRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetCodecDescriptorRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetCodecDescriptorRequest)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GetCodecDescriptorRequest)
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
		x := input.Message.Interface().(*GetCodecDescriptorRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetCodecDescriptorRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetCodecDescriptorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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
	md_GetCodecDescriptorResponse       protoreflect.MessageDescriptor
	fd_GetCodecDescriptorResponse_codec protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_GetCodecDescriptorResponse = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("GetCodecDescriptorResponse")
	fd_GetCodecDescriptorResponse_codec = md_GetCodecDescriptorResponse.Fields().ByName("codec")
}

var _ protoreflect.Message = (*fastReflection_GetCodecDescriptorResponse)(nil)

type fastReflection_GetCodecDescriptorResponse GetCodecDescriptorResponse

func (x *GetCodecDescriptorResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetCodecDescriptorResponse)(x)
}

func (x *GetCodecDescriptorResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetCodecDescriptorResponse_messageType fastReflection_GetCodecDescriptorResponse_messageType
var _ protoreflect.MessageType = fastReflection_GetCodecDescriptorResponse_messageType{}

type fastReflection_GetCodecDescriptorResponse_messageType struct{}

func (x fastReflection_GetCodecDescriptorResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetCodecDescriptorResponse)(nil)
}
func (x fastReflection_GetCodecDescriptorResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_GetCodecDescriptorResponse)
}
func (x fastReflection_GetCodecDescriptorResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetCodecDescriptorResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetCodecDescriptorResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_GetCodecDescriptorResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetCodecDescriptorResponse) Type() protoreflect.MessageType {
	return _fastReflection_GetCodecDescriptorResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetCodecDescriptorResponse) New() protoreflect.Message {
	return new(fastReflection_GetCodecDescriptorResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetCodecDescriptorResponse) Interface() protoreflect.ProtoMessage {
	return (*GetCodecDescriptorResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetCodecDescriptorResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Codec != nil {
		value := protoreflect.ValueOfMessage(x.Codec.ProtoReflect())
		if !f(fd_GetCodecDescriptorResponse_codec, value) {
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
func (x *fastReflection_GetCodecDescriptorResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse.codec":
		return x.Codec != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetCodecDescriptorResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse.codec":
		x.Codec = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetCodecDescriptorResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse.codec":
		value := x.Codec
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_GetCodecDescriptorResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse.codec":
		x.Codec = value.Message().Interface().(*CodecDescriptor)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_GetCodecDescriptorResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse.codec":
		if x.Codec == nil {
			x.Codec = new(CodecDescriptor)
		}
		return protoreflect.ValueOfMessage(x.Codec.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetCodecDescriptorResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse.codec":
		m := new(CodecDescriptor)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetCodecDescriptorResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetCodecDescriptorResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetCodecDescriptorResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_GetCodecDescriptorResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetCodecDescriptorResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetCodecDescriptorResponse)
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
		if x.Codec != nil {
			l = options.Size(x.Codec)
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
		x := input.Message.Interface().(*GetCodecDescriptorResponse)
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
		if x.Codec != nil {
			encoded, err := options.Marshal(x.Codec)
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
		x := input.Message.Interface().(*GetCodecDescriptorResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetCodecDescriptorResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetCodecDescriptorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Codec", wireType)
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
				if x.Codec == nil {
					x.Codec = &CodecDescriptor{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Codec); err != nil {
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
	md_GetConfigurationDescriptorRequest protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_GetConfigurationDescriptorRequest = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("GetConfigurationDescriptorRequest")
}

var _ protoreflect.Message = (*fastReflection_GetConfigurationDescriptorRequest)(nil)

type fastReflection_GetConfigurationDescriptorRequest GetConfigurationDescriptorRequest

func (x *GetConfigurationDescriptorRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetConfigurationDescriptorRequest)(x)
}

func (x *GetConfigurationDescriptorRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetConfigurationDescriptorRequest_messageType fastReflection_GetConfigurationDescriptorRequest_messageType
var _ protoreflect.MessageType = fastReflection_GetConfigurationDescriptorRequest_messageType{}

type fastReflection_GetConfigurationDescriptorRequest_messageType struct{}

func (x fastReflection_GetConfigurationDescriptorRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetConfigurationDescriptorRequest)(nil)
}
func (x fastReflection_GetConfigurationDescriptorRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_GetConfigurationDescriptorRequest)
}
func (x fastReflection_GetConfigurationDescriptorRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetConfigurationDescriptorRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetConfigurationDescriptorRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_GetConfigurationDescriptorRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetConfigurationDescriptorRequest) Type() protoreflect.MessageType {
	return _fastReflection_GetConfigurationDescriptorRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetConfigurationDescriptorRequest) New() protoreflect.Message {
	return new(fastReflection_GetConfigurationDescriptorRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetConfigurationDescriptorRequest) Interface() protoreflect.ProtoMessage {
	return (*GetConfigurationDescriptorRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetConfigurationDescriptorRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_GetConfigurationDescriptorRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetConfigurationDescriptorRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetConfigurationDescriptorRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_GetConfigurationDescriptorRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_GetConfigurationDescriptorRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetConfigurationDescriptorRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetConfigurationDescriptorRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetConfigurationDescriptorRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetConfigurationDescriptorRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_GetConfigurationDescriptorRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetConfigurationDescriptorRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetConfigurationDescriptorRequest)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GetConfigurationDescriptorRequest)
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
		x := input.Message.Interface().(*GetConfigurationDescriptorRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetConfigurationDescriptorRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetConfigurationDescriptorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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
	md_GetConfigurationDescriptorResponse        protoreflect.MessageDescriptor
	fd_GetConfigurationDescriptorResponse_config protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_GetConfigurationDescriptorResponse = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("GetConfigurationDescriptorResponse")
	fd_GetConfigurationDescriptorResponse_config = md_GetConfigurationDescriptorResponse.Fields().ByName("config")
}

var _ protoreflect.Message = (*fastReflection_GetConfigurationDescriptorResponse)(nil)

type fastReflection_GetConfigurationDescriptorResponse GetConfigurationDescriptorResponse

func (x *GetConfigurationDescriptorResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetConfigurationDescriptorResponse)(x)
}

func (x *GetConfigurationDescriptorResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetConfigurationDescriptorResponse_messageType fastReflection_GetConfigurationDescriptorResponse_messageType
var _ protoreflect.MessageType = fastReflection_GetConfigurationDescriptorResponse_messageType{}

type fastReflection_GetConfigurationDescriptorResponse_messageType struct{}

func (x fastReflection_GetConfigurationDescriptorResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetConfigurationDescriptorResponse)(nil)
}
func (x fastReflection_GetConfigurationDescriptorResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_GetConfigurationDescriptorResponse)
}
func (x fastReflection_GetConfigurationDescriptorResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetConfigurationDescriptorResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetConfigurationDescriptorResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_GetConfigurationDescriptorResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetConfigurationDescriptorResponse) Type() protoreflect.MessageType {
	return _fastReflection_GetConfigurationDescriptorResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetConfigurationDescriptorResponse) New() protoreflect.Message {
	return new(fastReflection_GetConfigurationDescriptorResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetConfigurationDescriptorResponse) Interface() protoreflect.ProtoMessage {
	return (*GetConfigurationDescriptorResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetConfigurationDescriptorResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Config != nil {
		value := protoreflect.ValueOfMessage(x.Config.ProtoReflect())
		if !f(fd_GetConfigurationDescriptorResponse_config, value) {
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
func (x *fastReflection_GetConfigurationDescriptorResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse.config":
		return x.Config != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetConfigurationDescriptorResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse.config":
		x.Config = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetConfigurationDescriptorResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse.config":
		value := x.Config
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_GetConfigurationDescriptorResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse.config":
		x.Config = value.Message().Interface().(*ConfigurationDescriptor)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_GetConfigurationDescriptorResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse.config":
		if x.Config == nil {
			x.Config = new(ConfigurationDescriptor)
		}
		return protoreflect.ValueOfMessage(x.Config.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetConfigurationDescriptorResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse.config":
		m := new(ConfigurationDescriptor)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetConfigurationDescriptorResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetConfigurationDescriptorResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetConfigurationDescriptorResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_GetConfigurationDescriptorResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetConfigurationDescriptorResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetConfigurationDescriptorResponse)
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
		if x.Config != nil {
			l = options.Size(x.Config)
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
		x := input.Message.Interface().(*GetConfigurationDescriptorResponse)
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
		if x.Config != nil {
			encoded, err := options.Marshal(x.Config)
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
		x := input.Message.Interface().(*GetConfigurationDescriptorResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetConfigurationDescriptorResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetConfigurationDescriptorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
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
				if x.Config == nil {
					x.Config = &ConfigurationDescriptor{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Config); err != nil {
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
	md_GetQueryServicesDescriptorRequest protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_GetQueryServicesDescriptorRequest = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("GetQueryServicesDescriptorRequest")
}

var _ protoreflect.Message = (*fastReflection_GetQueryServicesDescriptorRequest)(nil)

type fastReflection_GetQueryServicesDescriptorRequest GetQueryServicesDescriptorRequest

func (x *GetQueryServicesDescriptorRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetQueryServicesDescriptorRequest)(x)
}

func (x *GetQueryServicesDescriptorRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetQueryServicesDescriptorRequest_messageType fastReflection_GetQueryServicesDescriptorRequest_messageType
var _ protoreflect.MessageType = fastReflection_GetQueryServicesDescriptorRequest_messageType{}

type fastReflection_GetQueryServicesDescriptorRequest_messageType struct{}

func (x fastReflection_GetQueryServicesDescriptorRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetQueryServicesDescriptorRequest)(nil)
}
func (x fastReflection_GetQueryServicesDescriptorRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_GetQueryServicesDescriptorRequest)
}
func (x fastReflection_GetQueryServicesDescriptorRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetQueryServicesDescriptorRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetQueryServicesDescriptorRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_GetQueryServicesDescriptorRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetQueryServicesDescriptorRequest) Type() protoreflect.MessageType {
	return _fastReflection_GetQueryServicesDescriptorRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetQueryServicesDescriptorRequest) New() protoreflect.Message {
	return new(fastReflection_GetQueryServicesDescriptorRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetQueryServicesDescriptorRequest) Interface() protoreflect.ProtoMessage {
	return (*GetQueryServicesDescriptorRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetQueryServicesDescriptorRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_GetQueryServicesDescriptorRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetQueryServicesDescriptorRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetQueryServicesDescriptorRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_GetQueryServicesDescriptorRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_GetQueryServicesDescriptorRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetQueryServicesDescriptorRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetQueryServicesDescriptorRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetQueryServicesDescriptorRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetQueryServicesDescriptorRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_GetQueryServicesDescriptorRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetQueryServicesDescriptorRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetQueryServicesDescriptorRequest)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GetQueryServicesDescriptorRequest)
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
		x := input.Message.Interface().(*GetQueryServicesDescriptorRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetQueryServicesDescriptorRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetQueryServicesDescriptorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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
	md_GetQueryServicesDescriptorResponse         protoreflect.MessageDescriptor
	fd_GetQueryServicesDescriptorResponse_queries protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_GetQueryServicesDescriptorResponse = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("GetQueryServicesDescriptorResponse")
	fd_GetQueryServicesDescriptorResponse_queries = md_GetQueryServicesDescriptorResponse.Fields().ByName("queries")
}

var _ protoreflect.Message = (*fastReflection_GetQueryServicesDescriptorResponse)(nil)

type fastReflection_GetQueryServicesDescriptorResponse GetQueryServicesDescriptorResponse

func (x *GetQueryServicesDescriptorResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetQueryServicesDescriptorResponse)(x)
}

func (x *GetQueryServicesDescriptorResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[20]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetQueryServicesDescriptorResponse_messageType fastReflection_GetQueryServicesDescriptorResponse_messageType
var _ protoreflect.MessageType = fastReflection_GetQueryServicesDescriptorResponse_messageType{}

type fastReflection_GetQueryServicesDescriptorResponse_messageType struct{}

func (x fastReflection_GetQueryServicesDescriptorResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetQueryServicesDescriptorResponse)(nil)
}
func (x fastReflection_GetQueryServicesDescriptorResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_GetQueryServicesDescriptorResponse)
}
func (x fastReflection_GetQueryServicesDescriptorResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetQueryServicesDescriptorResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetQueryServicesDescriptorResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_GetQueryServicesDescriptorResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetQueryServicesDescriptorResponse) Type() protoreflect.MessageType {
	return _fastReflection_GetQueryServicesDescriptorResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetQueryServicesDescriptorResponse) New() protoreflect.Message {
	return new(fastReflection_GetQueryServicesDescriptorResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetQueryServicesDescriptorResponse) Interface() protoreflect.ProtoMessage {
	return (*GetQueryServicesDescriptorResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetQueryServicesDescriptorResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Queries != nil {
		value := protoreflect.ValueOfMessage(x.Queries.ProtoReflect())
		if !f(fd_GetQueryServicesDescriptorResponse_queries, value) {
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
func (x *fastReflection_GetQueryServicesDescriptorResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse.queries":
		return x.Queries != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetQueryServicesDescriptorResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse.queries":
		x.Queries = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetQueryServicesDescriptorResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse.queries":
		value := x.Queries
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_GetQueryServicesDescriptorResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse.queries":
		x.Queries = value.Message().Interface().(*QueryServicesDescriptor)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_GetQueryServicesDescriptorResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse.queries":
		if x.Queries == nil {
			x.Queries = new(QueryServicesDescriptor)
		}
		return protoreflect.ValueOfMessage(x.Queries.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetQueryServicesDescriptorResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse.queries":
		m := new(QueryServicesDescriptor)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetQueryServicesDescriptorResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetQueryServicesDescriptorResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetQueryServicesDescriptorResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_GetQueryServicesDescriptorResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetQueryServicesDescriptorResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetQueryServicesDescriptorResponse)
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
		if x.Queries != nil {
			l = options.Size(x.Queries)
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
		x := input.Message.Interface().(*GetQueryServicesDescriptorResponse)
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
		if x.Queries != nil {
			encoded, err := options.Marshal(x.Queries)
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
		x := input.Message.Interface().(*GetQueryServicesDescriptorResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetQueryServicesDescriptorResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetQueryServicesDescriptorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Queries", wireType)
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
				if x.Queries == nil {
					x.Queries = &QueryServicesDescriptor{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Queries); err != nil {
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
	md_GetTxDescriptorRequest protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_GetTxDescriptorRequest = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("GetTxDescriptorRequest")
}

var _ protoreflect.Message = (*fastReflection_GetTxDescriptorRequest)(nil)

type fastReflection_GetTxDescriptorRequest GetTxDescriptorRequest

func (x *GetTxDescriptorRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetTxDescriptorRequest)(x)
}

func (x *GetTxDescriptorRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[21]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetTxDescriptorRequest_messageType fastReflection_GetTxDescriptorRequest_messageType
var _ protoreflect.MessageType = fastReflection_GetTxDescriptorRequest_messageType{}

type fastReflection_GetTxDescriptorRequest_messageType struct{}

func (x fastReflection_GetTxDescriptorRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetTxDescriptorRequest)(nil)
}
func (x fastReflection_GetTxDescriptorRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_GetTxDescriptorRequest)
}
func (x fastReflection_GetTxDescriptorRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetTxDescriptorRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetTxDescriptorRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_GetTxDescriptorRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetTxDescriptorRequest) Type() protoreflect.MessageType {
	return _fastReflection_GetTxDescriptorRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetTxDescriptorRequest) New() protoreflect.Message {
	return new(fastReflection_GetTxDescriptorRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetTxDescriptorRequest) Interface() protoreflect.ProtoMessage {
	return (*GetTxDescriptorRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetTxDescriptorRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_GetTxDescriptorRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetTxDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetTxDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetTxDescriptorRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetTxDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetTxDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetTxDescriptorRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetTxDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetTxDescriptorRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_GetTxDescriptorRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetTxDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetTxDescriptorRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_GetTxDescriptorRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetTxDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetTxDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetTxDescriptorRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetTxDescriptorRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetTxDescriptorRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetTxDescriptorRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.GetTxDescriptorRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetTxDescriptorRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetTxDescriptorRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_GetTxDescriptorRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetTxDescriptorRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetTxDescriptorRequest)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GetTxDescriptorRequest)
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
		x := input.Message.Interface().(*GetTxDescriptorRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetTxDescriptorRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetTxDescriptorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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
	md_GetTxDescriptorResponse    protoreflect.MessageDescriptor
	fd_GetTxDescriptorResponse_tx protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_GetTxDescriptorResponse = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("GetTxDescriptorResponse")
	fd_GetTxDescriptorResponse_tx = md_GetTxDescriptorResponse.Fields().ByName("tx")
}

var _ protoreflect.Message = (*fastReflection_GetTxDescriptorResponse)(nil)

type fastReflection_GetTxDescriptorResponse GetTxDescriptorResponse

func (x *GetTxDescriptorResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GetTxDescriptorResponse)(x)
}

func (x *GetTxDescriptorResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[22]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GetTxDescriptorResponse_messageType fastReflection_GetTxDescriptorResponse_messageType
var _ protoreflect.MessageType = fastReflection_GetTxDescriptorResponse_messageType{}

type fastReflection_GetTxDescriptorResponse_messageType struct{}

func (x fastReflection_GetTxDescriptorResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GetTxDescriptorResponse)(nil)
}
func (x fastReflection_GetTxDescriptorResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_GetTxDescriptorResponse)
}
func (x fastReflection_GetTxDescriptorResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GetTxDescriptorResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GetTxDescriptorResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_GetTxDescriptorResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GetTxDescriptorResponse) Type() protoreflect.MessageType {
	return _fastReflection_GetTxDescriptorResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GetTxDescriptorResponse) New() protoreflect.Message {
	return new(fastReflection_GetTxDescriptorResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GetTxDescriptorResponse) Interface() protoreflect.ProtoMessage {
	return (*GetTxDescriptorResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GetTxDescriptorResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Tx != nil {
		value := protoreflect.ValueOfMessage(x.Tx.ProtoReflect())
		if !f(fd_GetTxDescriptorResponse_tx, value) {
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
func (x *fastReflection_GetTxDescriptorResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse.tx":
		return x.Tx != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetTxDescriptorResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse.tx":
		x.Tx = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GetTxDescriptorResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse.tx":
		value := x.Tx
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_GetTxDescriptorResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse.tx":
		x.Tx = value.Message().Interface().(*TxDescriptor)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_GetTxDescriptorResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse.tx":
		if x.Tx == nil {
			x.Tx = new(TxDescriptor)
		}
		return protoreflect.ValueOfMessage(x.Tx.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GetTxDescriptorResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse.tx":
		m := new(TxDescriptor)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GetTxDescriptorResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GetTxDescriptorResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GetTxDescriptorResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_GetTxDescriptorResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GetTxDescriptorResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GetTxDescriptorResponse)
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
		if x.Tx != nil {
			l = options.Size(x.Tx)
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
		x := input.Message.Interface().(*GetTxDescriptorResponse)
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
		if x.Tx != nil {
			encoded, err := options.Marshal(x.Tx)
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
		x := input.Message.Interface().(*GetTxDescriptorResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetTxDescriptorResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GetTxDescriptorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
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
				if x.Tx == nil {
					x.Tx = &TxDescriptor{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Tx); err != nil {
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

var _ protoreflect.List = (*_QueryServicesDescriptor_1_list)(nil)

type _QueryServicesDescriptor_1_list struct {
	list *[]*QueryServiceDescriptor
}

func (x *_QueryServicesDescriptor_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_QueryServicesDescriptor_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_QueryServicesDescriptor_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*QueryServiceDescriptor)
	(*x.list)[i] = concreteValue
}

func (x *_QueryServicesDescriptor_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*QueryServiceDescriptor)
	*x.list = append(*x.list, concreteValue)
}

func (x *_QueryServicesDescriptor_1_list) AppendMutable() protoreflect.Value {
	v := new(QueryServiceDescriptor)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryServicesDescriptor_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_QueryServicesDescriptor_1_list) NewElement() protoreflect.Value {
	v := new(QueryServiceDescriptor)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryServicesDescriptor_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_QueryServicesDescriptor                protoreflect.MessageDescriptor
	fd_QueryServicesDescriptor_query_services protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_QueryServicesDescriptor = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("QueryServicesDescriptor")
	fd_QueryServicesDescriptor_query_services = md_QueryServicesDescriptor.Fields().ByName("query_services")
}

var _ protoreflect.Message = (*fastReflection_QueryServicesDescriptor)(nil)

type fastReflection_QueryServicesDescriptor QueryServicesDescriptor

func (x *QueryServicesDescriptor) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryServicesDescriptor)(x)
}

func (x *QueryServicesDescriptor) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[23]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryServicesDescriptor_messageType fastReflection_QueryServicesDescriptor_messageType
var _ protoreflect.MessageType = fastReflection_QueryServicesDescriptor_messageType{}

type fastReflection_QueryServicesDescriptor_messageType struct{}

func (x fastReflection_QueryServicesDescriptor_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryServicesDescriptor)(nil)
}
func (x fastReflection_QueryServicesDescriptor_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryServicesDescriptor)
}
func (x fastReflection_QueryServicesDescriptor_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryServicesDescriptor
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryServicesDescriptor) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryServicesDescriptor
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryServicesDescriptor) Type() protoreflect.MessageType {
	return _fastReflection_QueryServicesDescriptor_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryServicesDescriptor) New() protoreflect.Message {
	return new(fastReflection_QueryServicesDescriptor)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryServicesDescriptor) Interface() protoreflect.ProtoMessage {
	return (*QueryServicesDescriptor)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryServicesDescriptor) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.QueryServices) != 0 {
		value := protoreflect.ValueOfList(&_QueryServicesDescriptor_1_list{list: &x.QueryServices})
		if !f(fd_QueryServicesDescriptor_query_services, value) {
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
func (x *fastReflection_QueryServicesDescriptor) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.QueryServicesDescriptor.query_services":
		return len(x.QueryServices) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.QueryServicesDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.QueryServicesDescriptor does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryServicesDescriptor) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.QueryServicesDescriptor.query_services":
		x.QueryServices = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.QueryServicesDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.QueryServicesDescriptor does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryServicesDescriptor) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v2alpha1.QueryServicesDescriptor.query_services":
		if len(x.QueryServices) == 0 {
			return protoreflect.ValueOfList(&_QueryServicesDescriptor_1_list{})
		}
		listValue := &_QueryServicesDescriptor_1_list{list: &x.QueryServices}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.QueryServicesDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.QueryServicesDescriptor does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryServicesDescriptor) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.QueryServicesDescriptor.query_services":
		lv := value.List()
		clv := lv.(*_QueryServicesDescriptor_1_list)
		x.QueryServices = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.QueryServicesDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.QueryServicesDescriptor does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryServicesDescriptor) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.QueryServicesDescriptor.query_services":
		if x.QueryServices == nil {
			x.QueryServices = []*QueryServiceDescriptor{}
		}
		value := &_QueryServicesDescriptor_1_list{list: &x.QueryServices}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.QueryServicesDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.QueryServicesDescriptor does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryServicesDescriptor) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.QueryServicesDescriptor.query_services":
		list := []*QueryServiceDescriptor{}
		return protoreflect.ValueOfList(&_QueryServicesDescriptor_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.QueryServicesDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.QueryServicesDescriptor does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryServicesDescriptor) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.QueryServicesDescriptor", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryServicesDescriptor) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryServicesDescriptor) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryServicesDescriptor) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryServicesDescriptor) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryServicesDescriptor)
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
		if len(x.QueryServices) > 0 {
			for _, e := range x.QueryServices {
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
		x := input.Message.Interface().(*QueryServicesDescriptor)
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
		if len(x.QueryServices) > 0 {
			for iNdEx := len(x.QueryServices) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.QueryServices[iNdEx])
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
		x := input.Message.Interface().(*QueryServicesDescriptor)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryServicesDescriptor: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryServicesDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field QueryServices", wireType)
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
				x.QueryServices = append(x.QueryServices, &QueryServiceDescriptor{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.QueryServices[len(x.QueryServices)-1]); err != nil {
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

var _ protoreflect.List = (*_QueryServiceDescriptor_3_list)(nil)

type _QueryServiceDescriptor_3_list struct {
	list *[]*QueryMethodDescriptor
}

func (x *_QueryServiceDescriptor_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_QueryServiceDescriptor_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_QueryServiceDescriptor_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*QueryMethodDescriptor)
	(*x.list)[i] = concreteValue
}

func (x *_QueryServiceDescriptor_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*QueryMethodDescriptor)
	*x.list = append(*x.list, concreteValue)
}

func (x *_QueryServiceDescriptor_3_list) AppendMutable() protoreflect.Value {
	v := new(QueryMethodDescriptor)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryServiceDescriptor_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_QueryServiceDescriptor_3_list) NewElement() protoreflect.Value {
	v := new(QueryMethodDescriptor)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryServiceDescriptor_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_QueryServiceDescriptor           protoreflect.MessageDescriptor
	fd_QueryServiceDescriptor_fullname  protoreflect.FieldDescriptor
	fd_QueryServiceDescriptor_is_module protoreflect.FieldDescriptor
	fd_QueryServiceDescriptor_methods   protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_QueryServiceDescriptor = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("QueryServiceDescriptor")
	fd_QueryServiceDescriptor_fullname = md_QueryServiceDescriptor.Fields().ByName("fullname")
	fd_QueryServiceDescriptor_is_module = md_QueryServiceDescriptor.Fields().ByName("is_module")
	fd_QueryServiceDescriptor_methods = md_QueryServiceDescriptor.Fields().ByName("methods")
}

var _ protoreflect.Message = (*fastReflection_QueryServiceDescriptor)(nil)

type fastReflection_QueryServiceDescriptor QueryServiceDescriptor

func (x *QueryServiceDescriptor) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryServiceDescriptor)(x)
}

func (x *QueryServiceDescriptor) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[24]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryServiceDescriptor_messageType fastReflection_QueryServiceDescriptor_messageType
var _ protoreflect.MessageType = fastReflection_QueryServiceDescriptor_messageType{}

type fastReflection_QueryServiceDescriptor_messageType struct{}

func (x fastReflection_QueryServiceDescriptor_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryServiceDescriptor)(nil)
}
func (x fastReflection_QueryServiceDescriptor_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryServiceDescriptor)
}
func (x fastReflection_QueryServiceDescriptor_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryServiceDescriptor
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryServiceDescriptor) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryServiceDescriptor
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryServiceDescriptor) Type() protoreflect.MessageType {
	return _fastReflection_QueryServiceDescriptor_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryServiceDescriptor) New() protoreflect.Message {
	return new(fastReflection_QueryServiceDescriptor)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryServiceDescriptor) Interface() protoreflect.ProtoMessage {
	return (*QueryServiceDescriptor)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryServiceDescriptor) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Fullname != "" {
		value := protoreflect.ValueOfString(x.Fullname)
		if !f(fd_QueryServiceDescriptor_fullname, value) {
			return
		}
	}
	if x.IsModule != false {
		value := protoreflect.ValueOfBool(x.IsModule)
		if !f(fd_QueryServiceDescriptor_is_module, value) {
			return
		}
	}
	if len(x.Methods) != 0 {
		value := protoreflect.ValueOfList(&_QueryServiceDescriptor_3_list{list: &x.Methods})
		if !f(fd_QueryServiceDescriptor_methods, value) {
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
func (x *fastReflection_QueryServiceDescriptor) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.QueryServiceDescriptor.fullname":
		return x.Fullname != ""
	case "cosmos.base.reflection.v2alpha1.QueryServiceDescriptor.is_module":
		return x.IsModule != false
	case "cosmos.base.reflection.v2alpha1.QueryServiceDescriptor.methods":
		return len(x.Methods) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.QueryServiceDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.QueryServiceDescriptor does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryServiceDescriptor) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.QueryServiceDescriptor.fullname":
		x.Fullname = ""
	case "cosmos.base.reflection.v2alpha1.QueryServiceDescriptor.is_module":
		x.IsModule = false
	case "cosmos.base.reflection.v2alpha1.QueryServiceDescriptor.methods":
		x.Methods = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.QueryServiceDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.QueryServiceDescriptor does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryServiceDescriptor) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v2alpha1.QueryServiceDescriptor.fullname":
		value := x.Fullname
		return protoreflect.ValueOfString(value)
	case "cosmos.base.reflection.v2alpha1.QueryServiceDescriptor.is_module":
		value := x.IsModule
		return protoreflect.ValueOfBool(value)
	case "cosmos.base.reflection.v2alpha1.QueryServiceDescriptor.methods":
		if len(x.Methods) == 0 {
			return protoreflect.ValueOfList(&_QueryServiceDescriptor_3_list{})
		}
		listValue := &_QueryServiceDescriptor_3_list{list: &x.Methods}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.QueryServiceDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.QueryServiceDescriptor does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryServiceDescriptor) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.QueryServiceDescriptor.fullname":
		x.Fullname = value.Interface().(string)
	case "cosmos.base.reflection.v2alpha1.QueryServiceDescriptor.is_module":
		x.IsModule = value.Bool()
	case "cosmos.base.reflection.v2alpha1.QueryServiceDescriptor.methods":
		lv := value.List()
		clv := lv.(*_QueryServiceDescriptor_3_list)
		x.Methods = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.QueryServiceDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.QueryServiceDescriptor does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryServiceDescriptor) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.QueryServiceDescriptor.methods":
		if x.Methods == nil {
			x.Methods = []*QueryMethodDescriptor{}
		}
		value := &_QueryServiceDescriptor_3_list{list: &x.Methods}
		return protoreflect.ValueOfList(value)
	case "cosmos.base.reflection.v2alpha1.QueryServiceDescriptor.fullname":
		panic(fmt.Errorf("field fullname of message cosmos.base.reflection.v2alpha1.QueryServiceDescriptor is not mutable"))
	case "cosmos.base.reflection.v2alpha1.QueryServiceDescriptor.is_module":
		panic(fmt.Errorf("field is_module of message cosmos.base.reflection.v2alpha1.QueryServiceDescriptor is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.QueryServiceDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.QueryServiceDescriptor does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryServiceDescriptor) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.QueryServiceDescriptor.fullname":
		return protoreflect.ValueOfString("")
	case "cosmos.base.reflection.v2alpha1.QueryServiceDescriptor.is_module":
		return protoreflect.ValueOfBool(false)
	case "cosmos.base.reflection.v2alpha1.QueryServiceDescriptor.methods":
		list := []*QueryMethodDescriptor{}
		return protoreflect.ValueOfList(&_QueryServiceDescriptor_3_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.QueryServiceDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.QueryServiceDescriptor does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryServiceDescriptor) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.QueryServiceDescriptor", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryServiceDescriptor) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryServiceDescriptor) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryServiceDescriptor) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryServiceDescriptor) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryServiceDescriptor)
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
		l = len(x.Fullname)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.IsModule {
			n += 2
		}
		if len(x.Methods) > 0 {
			for _, e := range x.Methods {
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
		x := input.Message.Interface().(*QueryServiceDescriptor)
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
		if len(x.Methods) > 0 {
			for iNdEx := len(x.Methods) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Methods[iNdEx])
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
		if x.IsModule {
			i--
			if x.IsModule {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x10
		}
		if len(x.Fullname) > 0 {
			i -= len(x.Fullname)
			copy(dAtA[i:], x.Fullname)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Fullname)))
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
		x := input.Message.Interface().(*QueryServiceDescriptor)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryServiceDescriptor: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryServiceDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Fullname", wireType)
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
				x.Fullname = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field IsModule", wireType)
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
				x.IsModule = bool(v != 0)
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Methods", wireType)
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
				x.Methods = append(x.Methods, &QueryMethodDescriptor{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Methods[len(x.Methods)-1]); err != nil {
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
	md_QueryMethodDescriptor                 protoreflect.MessageDescriptor
	fd_QueryMethodDescriptor_name            protoreflect.FieldDescriptor
	fd_QueryMethodDescriptor_full_query_path protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_init()
	md_QueryMethodDescriptor = File_cosmos_base_reflection_v2alpha1_reflection_proto.Messages().ByName("QueryMethodDescriptor")
	fd_QueryMethodDescriptor_name = md_QueryMethodDescriptor.Fields().ByName("name")
	fd_QueryMethodDescriptor_full_query_path = md_QueryMethodDescriptor.Fields().ByName("full_query_path")
}

var _ protoreflect.Message = (*fastReflection_QueryMethodDescriptor)(nil)

type fastReflection_QueryMethodDescriptor QueryMethodDescriptor

func (x *QueryMethodDescriptor) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryMethodDescriptor)(x)
}

func (x *QueryMethodDescriptor) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[25]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryMethodDescriptor_messageType fastReflection_QueryMethodDescriptor_messageType
var _ protoreflect.MessageType = fastReflection_QueryMethodDescriptor_messageType{}

type fastReflection_QueryMethodDescriptor_messageType struct{}

func (x fastReflection_QueryMethodDescriptor_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryMethodDescriptor)(nil)
}
func (x fastReflection_QueryMethodDescriptor_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryMethodDescriptor)
}
func (x fastReflection_QueryMethodDescriptor_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryMethodDescriptor
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryMethodDescriptor) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryMethodDescriptor
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryMethodDescriptor) Type() protoreflect.MessageType {
	return _fastReflection_QueryMethodDescriptor_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryMethodDescriptor) New() protoreflect.Message {
	return new(fastReflection_QueryMethodDescriptor)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryMethodDescriptor) Interface() protoreflect.ProtoMessage {
	return (*QueryMethodDescriptor)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryMethodDescriptor) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Name != "" {
		value := protoreflect.ValueOfString(x.Name)
		if !f(fd_QueryMethodDescriptor_name, value) {
			return
		}
	}
	if x.FullQueryPath != "" {
		value := protoreflect.ValueOfString(x.FullQueryPath)
		if !f(fd_QueryMethodDescriptor_full_query_path, value) {
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
func (x *fastReflection_QueryMethodDescriptor) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.QueryMethodDescriptor.name":
		return x.Name != ""
	case "cosmos.base.reflection.v2alpha1.QueryMethodDescriptor.full_query_path":
		return x.FullQueryPath != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.QueryMethodDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.QueryMethodDescriptor does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryMethodDescriptor) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.QueryMethodDescriptor.name":
		x.Name = ""
	case "cosmos.base.reflection.v2alpha1.QueryMethodDescriptor.full_query_path":
		x.FullQueryPath = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.QueryMethodDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.QueryMethodDescriptor does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryMethodDescriptor) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v2alpha1.QueryMethodDescriptor.name":
		value := x.Name
		return protoreflect.ValueOfString(value)
	case "cosmos.base.reflection.v2alpha1.QueryMethodDescriptor.full_query_path":
		value := x.FullQueryPath
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.QueryMethodDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.QueryMethodDescriptor does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryMethodDescriptor) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.QueryMethodDescriptor.name":
		x.Name = value.Interface().(string)
	case "cosmos.base.reflection.v2alpha1.QueryMethodDescriptor.full_query_path":
		x.FullQueryPath = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.QueryMethodDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.QueryMethodDescriptor does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryMethodDescriptor) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.QueryMethodDescriptor.name":
		panic(fmt.Errorf("field name of message cosmos.base.reflection.v2alpha1.QueryMethodDescriptor is not mutable"))
	case "cosmos.base.reflection.v2alpha1.QueryMethodDescriptor.full_query_path":
		panic(fmt.Errorf("field full_query_path of message cosmos.base.reflection.v2alpha1.QueryMethodDescriptor is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.QueryMethodDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.QueryMethodDescriptor does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryMethodDescriptor) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v2alpha1.QueryMethodDescriptor.name":
		return protoreflect.ValueOfString("")
	case "cosmos.base.reflection.v2alpha1.QueryMethodDescriptor.full_query_path":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v2alpha1.QueryMethodDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v2alpha1.QueryMethodDescriptor does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryMethodDescriptor) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v2alpha1.QueryMethodDescriptor", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryMethodDescriptor) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryMethodDescriptor) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryMethodDescriptor) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryMethodDescriptor) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryMethodDescriptor)
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
		l = len(x.Name)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.FullQueryPath)
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
		x := input.Message.Interface().(*QueryMethodDescriptor)
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
		if len(x.FullQueryPath) > 0 {
			i -= len(x.FullQueryPath)
			copy(dAtA[i:], x.FullQueryPath)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.FullQueryPath)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Name) > 0 {
			i -= len(x.Name)
			copy(dAtA[i:], x.Name)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Name)))
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
		x := input.Message.Interface().(*QueryMethodDescriptor)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryMethodDescriptor: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryMethodDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field FullQueryPath", wireType)
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
				x.FullQueryPath = string(dAtA[iNdEx:postIndex])
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ReflectionServiceClient is the client API for ReflectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReflectionServiceClient interface {
	// GetAuthnDescriptor returns information on how to authenticate transactions in the application
	// NOTE: this RPC is still experimental and might be subject to breaking changes or removal in
	// future releases of the cosmos-sdk.
	GetAuthnDescriptor(ctx context.Context, in *GetAuthnDescriptorRequest, opts ...grpc.CallOption) (*GetAuthnDescriptorResponse, error)
	// GetChainDescriptor returns the description of the chain
	GetChainDescriptor(ctx context.Context, in *GetChainDescriptorRequest, opts ...grpc.CallOption) (*GetChainDescriptorResponse, error)
	// GetCodecDescriptor returns the descriptor of the codec of the application
	GetCodecDescriptor(ctx context.Context, in *GetCodecDescriptorRequest, opts ...grpc.CallOption) (*GetCodecDescriptorResponse, error)
	// GetConfigurationDescriptor returns the descriptor for the sdk.Config of the application
	GetConfigurationDescriptor(ctx context.Context, in *GetConfigurationDescriptorRequest, opts ...grpc.CallOption) (*GetConfigurationDescriptorResponse, error)
	// GetQueryServicesDescriptor returns the available gRPC queryable services of the application
	GetQueryServicesDescriptor(ctx context.Context, in *GetQueryServicesDescriptorRequest, opts ...grpc.CallOption) (*GetQueryServicesDescriptorResponse, error)
	// GetTxDescriptor returns information on the used transaction object and available msgs that can be used
	GetTxDescriptor(ctx context.Context, in *GetTxDescriptorRequest, opts ...grpc.CallOption) (*GetTxDescriptorResponse, error)
}

type reflectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReflectionServiceClient(cc grpc.ClientConnInterface) ReflectionServiceClient {
	return &reflectionServiceClient{cc}
}

func (c *reflectionServiceClient) GetAuthnDescriptor(ctx context.Context, in *GetAuthnDescriptorRequest, opts ...grpc.CallOption) (*GetAuthnDescriptorResponse, error) {
	out := new(GetAuthnDescriptorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v2alpha1.ReflectionService/GetAuthnDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reflectionServiceClient) GetChainDescriptor(ctx context.Context, in *GetChainDescriptorRequest, opts ...grpc.CallOption) (*GetChainDescriptorResponse, error) {
	out := new(GetChainDescriptorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v2alpha1.ReflectionService/GetChainDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reflectionServiceClient) GetCodecDescriptor(ctx context.Context, in *GetCodecDescriptorRequest, opts ...grpc.CallOption) (*GetCodecDescriptorResponse, error) {
	out := new(GetCodecDescriptorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v2alpha1.ReflectionService/GetCodecDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reflectionServiceClient) GetConfigurationDescriptor(ctx context.Context, in *GetConfigurationDescriptorRequest, opts ...grpc.CallOption) (*GetConfigurationDescriptorResponse, error) {
	out := new(GetConfigurationDescriptorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v2alpha1.ReflectionService/GetConfigurationDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reflectionServiceClient) GetQueryServicesDescriptor(ctx context.Context, in *GetQueryServicesDescriptorRequest, opts ...grpc.CallOption) (*GetQueryServicesDescriptorResponse, error) {
	out := new(GetQueryServicesDescriptorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v2alpha1.ReflectionService/GetQueryServicesDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reflectionServiceClient) GetTxDescriptor(ctx context.Context, in *GetTxDescriptorRequest, opts ...grpc.CallOption) (*GetTxDescriptorResponse, error) {
	out := new(GetTxDescriptorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v2alpha1.ReflectionService/GetTxDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReflectionServiceServer is the server API for ReflectionService service.
// All implementations must embed UnimplementedReflectionServiceServer
// for forward compatibility
type ReflectionServiceServer interface {
	// GetAuthnDescriptor returns information on how to authenticate transactions in the application
	// NOTE: this RPC is still experimental and might be subject to breaking changes or removal in
	// future releases of the cosmos-sdk.
	GetAuthnDescriptor(context.Context, *GetAuthnDescriptorRequest) (*GetAuthnDescriptorResponse, error)
	// GetChainDescriptor returns the description of the chain
	GetChainDescriptor(context.Context, *GetChainDescriptorRequest) (*GetChainDescriptorResponse, error)
	// GetCodecDescriptor returns the descriptor of the codec of the application
	GetCodecDescriptor(context.Context, *GetCodecDescriptorRequest) (*GetCodecDescriptorResponse, error)
	// GetConfigurationDescriptor returns the descriptor for the sdk.Config of the application
	GetConfigurationDescriptor(context.Context, *GetConfigurationDescriptorRequest) (*GetConfigurationDescriptorResponse, error)
	// GetQueryServicesDescriptor returns the available gRPC queryable services of the application
	GetQueryServicesDescriptor(context.Context, *GetQueryServicesDescriptorRequest) (*GetQueryServicesDescriptorResponse, error)
	// GetTxDescriptor returns information on the used transaction object and available msgs that can be used
	GetTxDescriptor(context.Context, *GetTxDescriptorRequest) (*GetTxDescriptorResponse, error)
	mustEmbedUnimplementedReflectionServiceServer()
}

// UnimplementedReflectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReflectionServiceServer struct {
}

func (UnimplementedReflectionServiceServer) GetAuthnDescriptor(context.Context, *GetAuthnDescriptorRequest) (*GetAuthnDescriptorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthnDescriptor not implemented")
}
func (UnimplementedReflectionServiceServer) GetChainDescriptor(context.Context, *GetChainDescriptorRequest) (*GetChainDescriptorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChainDescriptor not implemented")
}
func (UnimplementedReflectionServiceServer) GetCodecDescriptor(context.Context, *GetCodecDescriptorRequest) (*GetCodecDescriptorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCodecDescriptor not implemented")
}
func (UnimplementedReflectionServiceServer) GetConfigurationDescriptor(context.Context, *GetConfigurationDescriptorRequest) (*GetConfigurationDescriptorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfigurationDescriptor not implemented")
}
func (UnimplementedReflectionServiceServer) GetQueryServicesDescriptor(context.Context, *GetQueryServicesDescriptorRequest) (*GetQueryServicesDescriptorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueryServicesDescriptor not implemented")
}
func (UnimplementedReflectionServiceServer) GetTxDescriptor(context.Context, *GetTxDescriptorRequest) (*GetTxDescriptorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxDescriptor not implemented")
}
func (UnimplementedReflectionServiceServer) mustEmbedUnimplementedReflectionServiceServer() {}

// UnsafeReflectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReflectionServiceServer will
// result in compilation errors.
type UnsafeReflectionServiceServer interface {
	mustEmbedUnimplementedReflectionServiceServer()
}

func RegisterReflectionServiceServer(s grpc.ServiceRegistrar, srv ReflectionServiceServer) {
	s.RegisterService(&ReflectionService_ServiceDesc, srv)
}

func _ReflectionService_GetAuthnDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthnDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).GetAuthnDescriptor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v2alpha1.ReflectionService/GetAuthnDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).GetAuthnDescriptor(ctx, req.(*GetAuthnDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReflectionService_GetChainDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChainDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).GetChainDescriptor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v2alpha1.ReflectionService/GetChainDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).GetChainDescriptor(ctx, req.(*GetChainDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReflectionService_GetCodecDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCodecDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).GetCodecDescriptor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v2alpha1.ReflectionService/GetCodecDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).GetCodecDescriptor(ctx, req.(*GetCodecDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReflectionService_GetConfigurationDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigurationDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).GetConfigurationDescriptor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v2alpha1.ReflectionService/GetConfigurationDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).GetConfigurationDescriptor(ctx, req.(*GetConfigurationDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReflectionService_GetQueryServicesDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQueryServicesDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).GetQueryServicesDescriptor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v2alpha1.ReflectionService/GetQueryServicesDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).GetQueryServicesDescriptor(ctx, req.(*GetQueryServicesDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReflectionService_GetTxDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).GetTxDescriptor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v2alpha1.ReflectionService/GetTxDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).GetTxDescriptor(ctx, req.(*GetTxDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReflectionService_ServiceDesc is the grpc.ServiceDesc for ReflectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReflectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.base.reflection.v2alpha1.ReflectionService",
	HandlerType: (*ReflectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthnDescriptor",
			Handler:    _ReflectionService_GetAuthnDescriptor_Handler,
		},
		{
			MethodName: "GetChainDescriptor",
			Handler:    _ReflectionService_GetChainDescriptor_Handler,
		},
		{
			MethodName: "GetCodecDescriptor",
			Handler:    _ReflectionService_GetCodecDescriptor_Handler,
		},
		{
			MethodName: "GetConfigurationDescriptor",
			Handler:    _ReflectionService_GetConfigurationDescriptor_Handler,
		},
		{
			MethodName: "GetQueryServicesDescriptor",
			Handler:    _ReflectionService_GetQueryServicesDescriptor_Handler,
		},
		{
			MethodName: "GetTxDescriptor",
			Handler:    _ReflectionService_GetTxDescriptor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/base/reflection/v2alpha1/reflection.proto",
}

// Since: cosmos-sdk 0.43

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.19.1
// source: cosmos/base/reflection/v2alpha1/reflection.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AppDescriptor describes a cosmos-sdk based application
type AppDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// AuthnDescriptor provides information on how to authenticate transactions on the application
	// NOTE: experimental and subject to change in future releases.
	Authn *AuthnDescriptor `protobuf:"bytes,1,opt,name=authn,proto3" json:"authn,omitempty"`
	// chain provides the chain descriptor
	Chain *ChainDescriptor `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	// codec provides metadata information regarding codec related types
	Codec *CodecDescriptor `protobuf:"bytes,3,opt,name=codec,proto3" json:"codec,omitempty"`
	// configuration provides metadata information regarding the sdk.Config type
	Configuration *ConfigurationDescriptor `protobuf:"bytes,4,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// query_services provides metadata information regarding the available queriable endpoints
	QueryServices *QueryServicesDescriptor `protobuf:"bytes,5,opt,name=query_services,json=queryServices,proto3" json:"query_services,omitempty"`
	// tx provides metadata information regarding how to send transactions to the given application
	Tx *TxDescriptor `protobuf:"bytes,6,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (x *AppDescriptor) Reset() {
	*x = AppDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppDescriptor) ProtoMessage() {}

// Deprecated: Use AppDescriptor.ProtoReflect.Descriptor instead.
func (*AppDescriptor) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{0}
}

func (x *AppDescriptor) GetAuthn() *AuthnDescriptor {
	if x != nil {
		return x.Authn
	}
	return nil
}

func (x *AppDescriptor) GetChain() *ChainDescriptor {
	if x != nil {
		return x.Chain
	}
	return nil
}

func (x *AppDescriptor) GetCodec() *CodecDescriptor {
	if x != nil {
		return x.Codec
	}
	return nil
}

func (x *AppDescriptor) GetConfiguration() *ConfigurationDescriptor {
	if x != nil {
		return x.Configuration
	}
	return nil
}

func (x *AppDescriptor) GetQueryServices() *QueryServicesDescriptor {
	if x != nil {
		return x.QueryServices
	}
	return nil
}

func (x *AppDescriptor) GetTx() *TxDescriptor {
	if x != nil {
		return x.Tx
	}
	return nil
}

// TxDescriptor describes the accepted transaction type
type TxDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// fullname is the protobuf fullname of the raw transaction type (for instance the tx.Tx type)
	// it is not meant to support polymorphism of transaction types, it is supposed to be used by
	// reflection clients to understand if they can handle a specific transaction type in an application.
	Fullname string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	// msgs lists the accepted application messages (sdk.Msg)
	Msgs []*MsgDescriptor `protobuf:"bytes,2,rep,name=msgs,proto3" json:"msgs,omitempty"`
}

func (x *TxDescriptor) Reset() {
	*x = TxDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxDescriptor) ProtoMessage() {}

// Deprecated: Use TxDescriptor.ProtoReflect.Descriptor instead.
func (*TxDescriptor) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{1}
}

func (x *TxDescriptor) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *TxDescriptor) GetMsgs() []*MsgDescriptor {
	if x != nil {
		return x.Msgs
	}
	return nil
}

// AuthnDescriptor provides information on how to sign transactions without relying
// on the online RPCs GetTxMetadata and CombineUnsignedTxAndSignatures
type AuthnDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sign_modes defines the supported signature algorithm
	SignModes []*SigningModeDescriptor `protobuf:"bytes,1,rep,name=sign_modes,json=signModes,proto3" json:"sign_modes,omitempty"`
}

func (x *AuthnDescriptor) Reset() {
	*x = AuthnDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthnDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthnDescriptor) ProtoMessage() {}

// Deprecated: Use AuthnDescriptor.ProtoReflect.Descriptor instead.
func (*AuthnDescriptor) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{2}
}

func (x *AuthnDescriptor) GetSignModes() []*SigningModeDescriptor {
	if x != nil {
		return x.SignModes
	}
	return nil
}

// SigningModeDescriptor provides information on a signing flow of the application
// NOTE(fdymylja): here we could go as far as providing an entire flow on how
// to sign a message given a SigningModeDescriptor, but it's better to think about
// this another time
type SigningModeDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name defines the unique name of the signing mode
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// number is the unique int32 identifier for the sign_mode enum
	Number int32 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	// authn_info_provider_method_fullname defines the fullname of the method to call to get
	// the metadata required to authenticate using the provided sign_modes
	AuthnInfoProviderMethodFullname string `protobuf:"bytes,3,opt,name=authn_info_provider_method_fullname,json=authnInfoProviderMethodFullname,proto3" json:"authn_info_provider_method_fullname,omitempty"`
}

func (x *SigningModeDescriptor) Reset() {
	*x = SigningModeDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningModeDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningModeDescriptor) ProtoMessage() {}

// Deprecated: Use SigningModeDescriptor.ProtoReflect.Descriptor instead.
func (*SigningModeDescriptor) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{3}
}

func (x *SigningModeDescriptor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SigningModeDescriptor) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *SigningModeDescriptor) GetAuthnInfoProviderMethodFullname() string {
	if x != nil {
		return x.AuthnInfoProviderMethodFullname
	}
	return ""
}

// ChainDescriptor describes chain information of the application
type ChainDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the chain id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ChainDescriptor) Reset() {
	*x = ChainDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChainDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChainDescriptor) ProtoMessage() {}

// Deprecated: Use ChainDescriptor.ProtoReflect.Descriptor instead.
func (*ChainDescriptor) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{4}
}

func (x *ChainDescriptor) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// CodecDescriptor describes the registered interfaces and provides metadata information on the types
type CodecDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// interfaces is a list of the registerted interfaces descriptors
	Interfaces []*InterfaceDescriptor `protobuf:"bytes,1,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
}

func (x *CodecDescriptor) Reset() {
	*x = CodecDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodecDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodecDescriptor) ProtoMessage() {}

// Deprecated: Use CodecDescriptor.ProtoReflect.Descriptor instead.
func (*CodecDescriptor) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{5}
}

func (x *CodecDescriptor) GetInterfaces() []*InterfaceDescriptor {
	if x != nil {
		return x.Interfaces
	}
	return nil
}

// InterfaceDescriptor describes the implementation of an interface
type InterfaceDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// fullname is the name of the interface
	Fullname string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	// interface_accepting_messages contains information regarding the proto messages which contain the interface as
	// google.protobuf.Any field
	InterfaceAcceptingMessages []*InterfaceAcceptingMessageDescriptor `protobuf:"bytes,2,rep,name=interface_accepting_messages,json=interfaceAcceptingMessages,proto3" json:"interface_accepting_messages,omitempty"`
	// interface_implementers is a list of the descriptors of the interface implementers
	InterfaceImplementers []*InterfaceImplementerDescriptor `protobuf:"bytes,3,rep,name=interface_implementers,json=interfaceImplementers,proto3" json:"interface_implementers,omitempty"`
}

func (x *InterfaceDescriptor) Reset() {
	*x = InterfaceDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterfaceDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfaceDescriptor) ProtoMessage() {}

// Deprecated: Use InterfaceDescriptor.ProtoReflect.Descriptor instead.
func (*InterfaceDescriptor) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{6}
}

func (x *InterfaceDescriptor) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *InterfaceDescriptor) GetInterfaceAcceptingMessages() []*InterfaceAcceptingMessageDescriptor {
	if x != nil {
		return x.InterfaceAcceptingMessages
	}
	return nil
}

func (x *InterfaceDescriptor) GetInterfaceImplementers() []*InterfaceImplementerDescriptor {
	if x != nil {
		return x.InterfaceImplementers
	}
	return nil
}

// InterfaceImplementerDescriptor describes an interface implementer
type InterfaceImplementerDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// fullname is the protobuf queryable name of the interface implementer
	Fullname string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	// type_url defines the type URL used when marshalling the type as any
	// this is required so we can provide type safe google.protobuf.Any marshalling and
	// unmarshalling, making sure that we don't accept just 'any' type
	// in our interface fields
	TypeUrl string `protobuf:"bytes,2,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
}

func (x *InterfaceImplementerDescriptor) Reset() {
	*x = InterfaceImplementerDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterfaceImplementerDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfaceImplementerDescriptor) ProtoMessage() {}

// Deprecated: Use InterfaceImplementerDescriptor.ProtoReflect.Descriptor instead.
func (*InterfaceImplementerDescriptor) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{7}
}

func (x *InterfaceImplementerDescriptor) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *InterfaceImplementerDescriptor) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

// InterfaceAcceptingMessageDescriptor describes a protobuf message which contains
// an interface represented as a google.protobuf.Any
type InterfaceAcceptingMessageDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// fullname is the protobuf fullname of the type containing the interface
	Fullname string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	// field_descriptor_names is a list of the protobuf name (not fullname) of the field
	// which contains the interface as google.protobuf.Any (the interface is the same, but
	// it can be in multiple fields of the same proto message)
	FieldDescriptorNames []string `protobuf:"bytes,2,rep,name=field_descriptor_names,json=fieldDescriptorNames,proto3" json:"field_descriptor_names,omitempty"`
}

func (x *InterfaceAcceptingMessageDescriptor) Reset() {
	*x = InterfaceAcceptingMessageDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterfaceAcceptingMessageDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfaceAcceptingMessageDescriptor) ProtoMessage() {}

// Deprecated: Use InterfaceAcceptingMessageDescriptor.ProtoReflect.Descriptor instead.
func (*InterfaceAcceptingMessageDescriptor) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{8}
}

func (x *InterfaceAcceptingMessageDescriptor) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *InterfaceAcceptingMessageDescriptor) GetFieldDescriptorNames() []string {
	if x != nil {
		return x.FieldDescriptorNames
	}
	return nil
}

// ConfigurationDescriptor contains metadata information on the sdk.Config
type ConfigurationDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// bech32_account_address_prefix is the account address prefix
	Bech32AccountAddressPrefix string `protobuf:"bytes,1,opt,name=bech32_account_address_prefix,json=bech32AccountAddressPrefix,proto3" json:"bech32_account_address_prefix,omitempty"`
}

func (x *ConfigurationDescriptor) Reset() {
	*x = ConfigurationDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigurationDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigurationDescriptor) ProtoMessage() {}

// Deprecated: Use ConfigurationDescriptor.ProtoReflect.Descriptor instead.
func (*ConfigurationDescriptor) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{9}
}

func (x *ConfigurationDescriptor) GetBech32AccountAddressPrefix() string {
	if x != nil {
		return x.Bech32AccountAddressPrefix
	}
	return ""
}

// MsgDescriptor describes a cosmos-sdk message that can be delivered with a transaction
type MsgDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// msg_type_url contains the TypeURL of a sdk.Msg.
	MsgTypeUrl string `protobuf:"bytes,1,opt,name=msg_type_url,json=msgTypeUrl,proto3" json:"msg_type_url,omitempty"`
}

func (x *MsgDescriptor) Reset() {
	*x = MsgDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDescriptor) ProtoMessage() {}

// Deprecated: Use MsgDescriptor.ProtoReflect.Descriptor instead.
func (*MsgDescriptor) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{10}
}

func (x *MsgDescriptor) GetMsgTypeUrl() string {
	if x != nil {
		return x.MsgTypeUrl
	}
	return ""
}

// GetAuthnDescriptorRequest is the request used for the GetAuthnDescriptor RPC
type GetAuthnDescriptorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAuthnDescriptorRequest) Reset() {
	*x = GetAuthnDescriptorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthnDescriptorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthnDescriptorRequest) ProtoMessage() {}

// Deprecated: Use GetAuthnDescriptorRequest.ProtoReflect.Descriptor instead.
func (*GetAuthnDescriptorRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{11}
}

// GetAuthnDescriptorResponse is the response returned by the GetAuthnDescriptor RPC
type GetAuthnDescriptorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// authn describes how to authenticate to the application when sending transactions
	Authn *AuthnDescriptor `protobuf:"bytes,1,opt,name=authn,proto3" json:"authn,omitempty"`
}

func (x *GetAuthnDescriptorResponse) Reset() {
	*x = GetAuthnDescriptorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthnDescriptorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthnDescriptorResponse) ProtoMessage() {}

// Deprecated: Use GetAuthnDescriptorResponse.ProtoReflect.Descriptor instead.
func (*GetAuthnDescriptorResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{12}
}

func (x *GetAuthnDescriptorResponse) GetAuthn() *AuthnDescriptor {
	if x != nil {
		return x.Authn
	}
	return nil
}

// GetChainDescriptorRequest is the request used for the GetChainDescriptor RPC
type GetChainDescriptorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetChainDescriptorRequest) Reset() {
	*x = GetChainDescriptorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChainDescriptorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChainDescriptorRequest) ProtoMessage() {}

// Deprecated: Use GetChainDescriptorRequest.ProtoReflect.Descriptor instead.
func (*GetChainDescriptorRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{13}
}

// GetChainDescriptorResponse is the response returned by the GetChainDescriptor RPC
type GetChainDescriptorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// chain describes application chain information
	Chain *ChainDescriptor `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
}

func (x *GetChainDescriptorResponse) Reset() {
	*x = GetChainDescriptorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChainDescriptorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChainDescriptorResponse) ProtoMessage() {}

// Deprecated: Use GetChainDescriptorResponse.ProtoReflect.Descriptor instead.
func (*GetChainDescriptorResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{14}
}

func (x *GetChainDescriptorResponse) GetChain() *ChainDescriptor {
	if x != nil {
		return x.Chain
	}
	return nil
}

// GetCodecDescriptorRequest is the request used for the GetCodecDescriptor RPC
type GetCodecDescriptorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCodecDescriptorRequest) Reset() {
	*x = GetCodecDescriptorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCodecDescriptorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCodecDescriptorRequest) ProtoMessage() {}

// Deprecated: Use GetCodecDescriptorRequest.ProtoReflect.Descriptor instead.
func (*GetCodecDescriptorRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{15}
}

// GetCodecDescriptorResponse is the response returned by the GetCodecDescriptor RPC
type GetCodecDescriptorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// codec describes the application codec such as registered interfaces and implementations
	Codec *CodecDescriptor `protobuf:"bytes,1,opt,name=codec,proto3" json:"codec,omitempty"`
}

func (x *GetCodecDescriptorResponse) Reset() {
	*x = GetCodecDescriptorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCodecDescriptorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCodecDescriptorResponse) ProtoMessage() {}

// Deprecated: Use GetCodecDescriptorResponse.ProtoReflect.Descriptor instead.
func (*GetCodecDescriptorResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{16}
}

func (x *GetCodecDescriptorResponse) GetCodec() *CodecDescriptor {
	if x != nil {
		return x.Codec
	}
	return nil
}

// GetConfigurationDescriptorRequest is the request used for the GetConfigurationDescriptor RPC
type GetConfigurationDescriptorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetConfigurationDescriptorRequest) Reset() {
	*x = GetConfigurationDescriptorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigurationDescriptorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigurationDescriptorRequest) ProtoMessage() {}

// Deprecated: Use GetConfigurationDescriptorRequest.ProtoReflect.Descriptor instead.
func (*GetConfigurationDescriptorRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{17}
}

// GetConfigurationDescriptorResponse is the response returned by the GetConfigurationDescriptor RPC
type GetConfigurationDescriptorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// config describes the application's sdk.Config
	Config *ConfigurationDescriptor `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *GetConfigurationDescriptorResponse) Reset() {
	*x = GetConfigurationDescriptorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigurationDescriptorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigurationDescriptorResponse) ProtoMessage() {}

// Deprecated: Use GetConfigurationDescriptorResponse.ProtoReflect.Descriptor instead.
func (*GetConfigurationDescriptorResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{18}
}

func (x *GetConfigurationDescriptorResponse) GetConfig() *ConfigurationDescriptor {
	if x != nil {
		return x.Config
	}
	return nil
}

// GetQueryServicesDescriptorRequest is the request used for the GetQueryServicesDescriptor RPC
type GetQueryServicesDescriptorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetQueryServicesDescriptorRequest) Reset() {
	*x = GetQueryServicesDescriptorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQueryServicesDescriptorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQueryServicesDescriptorRequest) ProtoMessage() {}

// Deprecated: Use GetQueryServicesDescriptorRequest.ProtoReflect.Descriptor instead.
func (*GetQueryServicesDescriptorRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{19}
}

// GetQueryServicesDescriptorResponse is the response returned by the GetQueryServicesDescriptor RPC
type GetQueryServicesDescriptorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// queries provides information on the available queryable services
	Queries *QueryServicesDescriptor `protobuf:"bytes,1,opt,name=queries,proto3" json:"queries,omitempty"`
}

func (x *GetQueryServicesDescriptorResponse) Reset() {
	*x = GetQueryServicesDescriptorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[20]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQueryServicesDescriptorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQueryServicesDescriptorResponse) ProtoMessage() {}

// Deprecated: Use GetQueryServicesDescriptorResponse.ProtoReflect.Descriptor instead.
func (*GetQueryServicesDescriptorResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{20}
}

func (x *GetQueryServicesDescriptorResponse) GetQueries() *QueryServicesDescriptor {
	if x != nil {
		return x.Queries
	}
	return nil
}

// GetTxDescriptorRequest is the request used for the GetTxDescriptor RPC
type GetTxDescriptorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTxDescriptorRequest) Reset() {
	*x = GetTxDescriptorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[21]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxDescriptorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxDescriptorRequest) ProtoMessage() {}

// Deprecated: Use GetTxDescriptorRequest.ProtoReflect.Descriptor instead.
func (*GetTxDescriptorRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{21}
}

// GetTxDescriptorResponse is the response returned by the GetTxDescriptor RPC
type GetTxDescriptorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tx provides information on msgs that can be forwarded to the application
	// alongside the accepted transaction protobuf type
	Tx *TxDescriptor `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (x *GetTxDescriptorResponse) Reset() {
	*x = GetTxDescriptorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[22]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxDescriptorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxDescriptorResponse) ProtoMessage() {}

// Deprecated: Use GetTxDescriptorResponse.ProtoReflect.Descriptor instead.
func (*GetTxDescriptorResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{22}
}

func (x *GetTxDescriptorResponse) GetTx() *TxDescriptor {
	if x != nil {
		return x.Tx
	}
	return nil
}

// QueryServicesDescriptor contains the list of cosmos-sdk queriable services
type QueryServicesDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// query_services is a list of cosmos-sdk QueryServiceDescriptor
	QueryServices []*QueryServiceDescriptor `protobuf:"bytes,1,rep,name=query_services,json=queryServices,proto3" json:"query_services,omitempty"`
}

func (x *QueryServicesDescriptor) Reset() {
	*x = QueryServicesDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[23]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryServicesDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryServicesDescriptor) ProtoMessage() {}

// Deprecated: Use QueryServicesDescriptor.ProtoReflect.Descriptor instead.
func (*QueryServicesDescriptor) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{23}
}

func (x *QueryServicesDescriptor) GetQueryServices() []*QueryServiceDescriptor {
	if x != nil {
		return x.QueryServices
	}
	return nil
}

// QueryServiceDescriptor describes a cosmos-sdk queryable service
type QueryServiceDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// fullname is the protobuf fullname of the service descriptor
	Fullname string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	// is_module describes if this service is actually exposed by an application's module
	IsModule bool `protobuf:"varint,2,opt,name=is_module,json=isModule,proto3" json:"is_module,omitempty"`
	// methods provides a list of query service methods
	Methods []*QueryMethodDescriptor `protobuf:"bytes,3,rep,name=methods,proto3" json:"methods,omitempty"`
}

func (x *QueryServiceDescriptor) Reset() {
	*x = QueryServiceDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[24]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryServiceDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryServiceDescriptor) ProtoMessage() {}

// Deprecated: Use QueryServiceDescriptor.ProtoReflect.Descriptor instead.
func (*QueryServiceDescriptor) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{24}
}

func (x *QueryServiceDescriptor) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *QueryServiceDescriptor) GetIsModule() bool {
	if x != nil {
		return x.IsModule
	}
	return false
}

func (x *QueryServiceDescriptor) GetMethods() []*QueryMethodDescriptor {
	if x != nil {
		return x.Methods
	}
	return nil
}

// QueryMethodDescriptor describes a queryable method of a query service
// no other info is provided beside method name and tendermint queryable path
// because it would be redundant with the grpc reflection service
type QueryMethodDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the protobuf name (not fullname) of the method
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// full_query_path is the path that can be used to query
	// this method via tendermint abci.Query
	FullQueryPath string `protobuf:"bytes,2,opt,name=full_query_path,json=fullQueryPath,proto3" json:"full_query_path,omitempty"`
}

func (x *QueryMethodDescriptor) Reset() {
	*x = QueryMethodDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[25]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMethodDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMethodDescriptor) ProtoMessage() {}

// Deprecated: Use QueryMethodDescriptor.ProtoReflect.Descriptor instead.
func (*QueryMethodDescriptor) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP(), []int{25}
}

func (x *QueryMethodDescriptor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *QueryMethodDescriptor) GetFullQueryPath() string {
	if x != nil {
		return x.FullQueryPath
	}
	return ""
}

var File_cosmos_base_reflection_v2alpha1_reflection_proto protoreflect.FileDescriptor

var file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDesc = []byte{
	0x0a, 0x30, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x72, 0x65,
	0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe7, 0x03, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x12, 0x46, 0x0a, 0x05, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x12, 0x46, 0x0a, 0x05, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x12, 0x46, 0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x12, 0x5e, 0x0a, 0x0d, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5f, 0x0a, 0x0e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x02,
	0x74, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x78, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x02, 0x74, 0x78, 0x22, 0x6e, 0x0a, 0x0c, 0x54,
	0x78, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x6d, 0x73, 0x67, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x04, 0x6d, 0x73, 0x67, 0x73, 0x22, 0x68, 0x0a, 0x0f, 0x41,
	0x75, 0x74, 0x68, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x55,
	0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x4d, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x91, 0x01, 0x0a, 0x15, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e,
	0x67, 0x4d, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x23, 0x61,
	0x75, 0x74, 0x68, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x46, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x21, 0x0a, 0x0f, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x67, 0x0a, 0x0f,
	0x43, 0x6f, 0x64, 0x65, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12,
	0x54, 0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x73, 0x22, 0xb2, 0x02, 0x0a, 0x13, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x1c, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6e,
	0x67, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x44, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72,
	0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x76, 0x0a, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6d,
	0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x52, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6d,
	0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x22, 0x57, 0x0a, 0x1e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65,
	0x55, 0x72, 0x6c, 0x22, 0x77, 0x0a, 0x23, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75,
	0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75,
	0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x5c, 0x0a, 0x17,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x41, 0x0a, 0x1d, 0x62, 0x65, 0x63, 0x68, 0x33,
	0x32, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a,
	0x62, 0x65, 0x63, 0x68, 0x33, 0x32, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x31, 0x0a, 0x0d, 0x4d, 0x73,
	0x67, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0c, 0x6d,
	0x73, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x1b, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x64, 0x0a, 0x1a, 0x47, 0x65,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x05, 0x61, 0x75, 0x74, 0x68,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x61, 0x75, 0x74, 0x68, 0x6e,
	0x22, 0x1b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x64, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x05, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x22, 0x1b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x64, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46,
	0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52,
	0x05, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x22, 0x23, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x76, 0x0a, 0x22, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x50, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0x23, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x78, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52,
	0x0a, 0x07, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x38, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65,
	0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x78, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x58, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x54, 0x78, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x78, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x52, 0x02, 0x74, 0x78, 0x22, 0x79, 0x0a, 0x17, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x12, 0x5e, 0x0a, 0x0e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x52, 0x0d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x22, 0xa3, 0x01, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x50, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x07,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x22, 0x53, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66,
	0x75, 0x6c, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x74, 0x68, 0x32, 0xa7, 0x0a, 0x0a,
	0x11, 0x52, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xcb, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x3a, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x3c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x36, 0x12, 0x34, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6e,
	0x12, 0xcb, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x3a, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x3c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x36, 0x12, 0x34, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0xcb,
	0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x3a, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x63,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x3b, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x36, 0x12, 0x34, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x12, 0xeb, 0x01, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x42, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x43, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65,
	0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x44, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3e, 0x12, 0x3c, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x72, 0x65, 0x66, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x70,
	0x70, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xec, 0x01, 0x0a, 0x1a, 0x47,
	0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x42, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x45, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3f, 0x12, 0x3d, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0xca, 0x01, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x54, 0x78, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x37, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x78, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x78, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x44, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3e, 0x12, 0x3c, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2f, 0x74, 0x78, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x42, 0xaa, 0x02, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0f,
	0x52, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f,
	0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x3b, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x42, 0x52, 0xaa, 0x02, 0x1f, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02,
	0x1f, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x42, 0x61, 0x73, 0x65, 0x5c, 0x52, 0x65, 0x66,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xe2, 0x02, 0x2b, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x42, 0x61, 0x73, 0x65, 0x5c, 0x52,
	0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x22, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x42, 0x61, 0x73, 0x65, 0x3a, 0x3a, 0x52,
	0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescOnce sync.Once
	file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescData = file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDesc
)

func file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescGZIP() []byte {
	file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescOnce.Do(func() {
		file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescData)
	})
	return file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDescData
}

var file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes = make([]protoimpl.MessageInfo, 26)
var file_cosmos_base_reflection_v2alpha1_reflection_proto_goTypes = []interface{}{
	(*AppDescriptor)(nil),                       // 0: cosmos.base.reflection.v2alpha1.AppDescriptor
	(*TxDescriptor)(nil),                        // 1: cosmos.base.reflection.v2alpha1.TxDescriptor
	(*AuthnDescriptor)(nil),                     // 2: cosmos.base.reflection.v2alpha1.AuthnDescriptor
	(*SigningModeDescriptor)(nil),               // 3: cosmos.base.reflection.v2alpha1.SigningModeDescriptor
	(*ChainDescriptor)(nil),                     // 4: cosmos.base.reflection.v2alpha1.ChainDescriptor
	(*CodecDescriptor)(nil),                     // 5: cosmos.base.reflection.v2alpha1.CodecDescriptor
	(*InterfaceDescriptor)(nil),                 // 6: cosmos.base.reflection.v2alpha1.InterfaceDescriptor
	(*InterfaceImplementerDescriptor)(nil),      // 7: cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor
	(*InterfaceAcceptingMessageDescriptor)(nil), // 8: cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor
	(*ConfigurationDescriptor)(nil),             // 9: cosmos.base.reflection.v2alpha1.ConfigurationDescriptor
	(*MsgDescriptor)(nil),                       // 10: cosmos.base.reflection.v2alpha1.MsgDescriptor
	(*GetAuthnDescriptorRequest)(nil),           // 11: cosmos.base.reflection.v2alpha1.GetAuthnDescriptorRequest
	(*GetAuthnDescriptorResponse)(nil),          // 12: cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse
	(*GetChainDescriptorRequest)(nil),           // 13: cosmos.base.reflection.v2alpha1.GetChainDescriptorRequest
	(*GetChainDescriptorResponse)(nil),          // 14: cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse
	(*GetCodecDescriptorRequest)(nil),           // 15: cosmos.base.reflection.v2alpha1.GetCodecDescriptorRequest
	(*GetCodecDescriptorResponse)(nil),          // 16: cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse
	(*GetConfigurationDescriptorRequest)(nil),   // 17: cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorRequest
	(*GetConfigurationDescriptorResponse)(nil),  // 18: cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse
	(*GetQueryServicesDescriptorRequest)(nil),   // 19: cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorRequest
	(*GetQueryServicesDescriptorResponse)(nil),  // 20: cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse
	(*GetTxDescriptorRequest)(nil),              // 21: cosmos.base.reflection.v2alpha1.GetTxDescriptorRequest
	(*GetTxDescriptorResponse)(nil),             // 22: cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse
	(*QueryServicesDescriptor)(nil),             // 23: cosmos.base.reflection.v2alpha1.QueryServicesDescriptor
	(*QueryServiceDescriptor)(nil),              // 24: cosmos.base.reflection.v2alpha1.QueryServiceDescriptor
	(*QueryMethodDescriptor)(nil),               // 25: cosmos.base.reflection.v2alpha1.QueryMethodDescriptor
}
var file_cosmos_base_reflection_v2alpha1_reflection_proto_depIdxs = []int32{
	2,  // 0: cosmos.base.reflection.v2alpha1.AppDescriptor.authn:type_name -> cosmos.base.reflection.v2alpha1.AuthnDescriptor
	4,  // 1: cosmos.base.reflection.v2alpha1.AppDescriptor.chain:type_name -> cosmos.base.reflection.v2alpha1.ChainDescriptor
	5,  // 2: cosmos.base.reflection.v2alpha1.AppDescriptor.codec:type_name -> cosmos.base.reflection.v2alpha1.CodecDescriptor
	9,  // 3: cosmos.base.reflection.v2alpha1.AppDescriptor.configuration:type_name -> cosmos.base.reflection.v2alpha1.ConfigurationDescriptor
	23, // 4: cosmos.base.reflection.v2alpha1.AppDescriptor.query_services:type_name -> cosmos.base.reflection.v2alpha1.QueryServicesDescriptor
	1,  // 5: cosmos.base.reflection.v2alpha1.AppDescriptor.tx:type_name -> cosmos.base.reflection.v2alpha1.TxDescriptor
	10, // 6: cosmos.base.reflection.v2alpha1.TxDescriptor.msgs:type_name -> cosmos.base.reflection.v2alpha1.MsgDescriptor
	3,  // 7: cosmos.base.reflection.v2alpha1.AuthnDescriptor.sign_modes:type_name -> cosmos.base.reflection.v2alpha1.SigningModeDescriptor
	6,  // 8: cosmos.base.reflection.v2alpha1.CodecDescriptor.interfaces:type_name -> cosmos.base.reflection.v2alpha1.InterfaceDescriptor
	8,  // 9: cosmos.base.reflection.v2alpha1.InterfaceDescriptor.interface_accepting_messages:type_name -> cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor
	7,  // 10: cosmos.base.reflection.v2alpha1.InterfaceDescriptor.interface_implementers:type_name -> cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor
	2,  // 11: cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse.authn:type_name -> cosmos.base.reflection.v2alpha1.AuthnDescriptor
	4,  // 12: cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse.chain:type_name -> cosmos.base.reflection.v2alpha1.ChainDescriptor
	5,  // 13: cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse.codec:type_name -> cosmos.base.reflection.v2alpha1.CodecDescriptor
	9,  // 14: cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse.config:type_name -> cosmos.base.reflection.v2alpha1.ConfigurationDescriptor
	23, // 15: cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse.queries:type_name -> cosmos.base.reflection.v2alpha1.QueryServicesDescriptor
	1,  // 16: cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse.tx:type_name -> cosmos.base.reflection.v2alpha1.TxDescriptor
	24, // 17: cosmos.base.reflection.v2alpha1.QueryServicesDescriptor.query_services:type_name -> cosmos.base.reflection.v2alpha1.QueryServiceDescriptor
	25, // 18: cosmos.base.reflection.v2alpha1.QueryServiceDescriptor.methods:type_name -> cosmos.base.reflection.v2alpha1.QueryMethodDescriptor
	11, // 19: cosmos.base.reflection.v2alpha1.ReflectionService.GetAuthnDescriptor:input_type -> cosmos.base.reflection.v2alpha1.GetAuthnDescriptorRequest
	13, // 20: cosmos.base.reflection.v2alpha1.ReflectionService.GetChainDescriptor:input_type -> cosmos.base.reflection.v2alpha1.GetChainDescriptorRequest
	15, // 21: cosmos.base.reflection.v2alpha1.ReflectionService.GetCodecDescriptor:input_type -> cosmos.base.reflection.v2alpha1.GetCodecDescriptorRequest
	17, // 22: cosmos.base.reflection.v2alpha1.ReflectionService.GetConfigurationDescriptor:input_type -> cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorRequest
	19, // 23: cosmos.base.reflection.v2alpha1.ReflectionService.GetQueryServicesDescriptor:input_type -> cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorRequest
	21, // 24: cosmos.base.reflection.v2alpha1.ReflectionService.GetTxDescriptor:input_type -> cosmos.base.reflection.v2alpha1.GetTxDescriptorRequest
	12, // 25: cosmos.base.reflection.v2alpha1.ReflectionService.GetAuthnDescriptor:output_type -> cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse
	14, // 26: cosmos.base.reflection.v2alpha1.ReflectionService.GetChainDescriptor:output_type -> cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse
	16, // 27: cosmos.base.reflection.v2alpha1.ReflectionService.GetCodecDescriptor:output_type -> cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse
	18, // 28: cosmos.base.reflection.v2alpha1.ReflectionService.GetConfigurationDescriptor:output_type -> cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse
	20, // 29: cosmos.base.reflection.v2alpha1.ReflectionService.GetQueryServicesDescriptor:output_type -> cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse
	22, // 30: cosmos.base.reflection.v2alpha1.ReflectionService.GetTxDescriptor:output_type -> cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse
	25, // [25:31] is the sub-list for method output_type
	19, // [19:25] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_cosmos_base_reflection_v2alpha1_reflection_proto_init() }
func file_cosmos_base_reflection_v2alpha1_reflection_proto_init() {
	if File_cosmos_base_reflection_v2alpha1_reflection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppDescriptor); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxDescriptor); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthnDescriptor); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigningModeDescriptor); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChainDescriptor); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodecDescriptor); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterfaceDescriptor); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterfaceImplementerDescriptor); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterfaceAcceptingMessageDescriptor); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigurationDescriptor); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDescriptor); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthnDescriptorRequest); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthnDescriptorResponse); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChainDescriptorRequest); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChainDescriptorResponse); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCodecDescriptorRequest); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCodecDescriptorResponse); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigurationDescriptorRequest); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigurationDescriptorResponse); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[19].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQueryServicesDescriptorRequest); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[20].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQueryServicesDescriptorResponse); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[21].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTxDescriptorRequest); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[22].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTxDescriptorResponse); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[23].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryServicesDescriptor); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[24].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryServiceDescriptor); i {
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
		file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes[25].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMethodDescriptor); i {
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
			RawDescriptor: file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   26,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_base_reflection_v2alpha1_reflection_proto_goTypes,
		DependencyIndexes: file_cosmos_base_reflection_v2alpha1_reflection_proto_depIdxs,
		MessageInfos:      file_cosmos_base_reflection_v2alpha1_reflection_proto_msgTypes,
	}.Build()
	File_cosmos_base_reflection_v2alpha1_reflection_proto = out.File
	file_cosmos_base_reflection_v2alpha1_reflection_proto_rawDesc = nil
	file_cosmos_base_reflection_v2alpha1_reflection_proto_goTypes = nil
	file_cosmos_base_reflection_v2alpha1_reflection_proto_depIdxs = nil
}
