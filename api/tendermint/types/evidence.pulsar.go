package types

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
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var (
	md_Evidence                              protoreflect.MessageDescriptor
	fd_Evidence_duplicate_vote_evidence      protoreflect.FieldDescriptor
	fd_Evidence_light_client_attack_evidence protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_evidence_proto_init()
	md_Evidence = File_tendermint_types_evidence_proto.Messages().ByName("Evidence")
	fd_Evidence_duplicate_vote_evidence = md_Evidence.Fields().ByName("duplicate_vote_evidence")
	fd_Evidence_light_client_attack_evidence = md_Evidence.Fields().ByName("light_client_attack_evidence")
}

var _ protoreflect.Message = (*fastReflection_Evidence)(nil)

type fastReflection_Evidence Evidence

func (x *Evidence) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Evidence)(x)
}

func (x *Evidence) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_evidence_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Evidence_messageType fastReflection_Evidence_messageType
var _ protoreflect.MessageType = fastReflection_Evidence_messageType{}

type fastReflection_Evidence_messageType struct{}

func (x fastReflection_Evidence_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Evidence)(nil)
}
func (x fastReflection_Evidence_messageType) New() protoreflect.Message {
	return new(fastReflection_Evidence)
}
func (x fastReflection_Evidence_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Evidence
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Evidence) Descriptor() protoreflect.MessageDescriptor {
	return md_Evidence
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Evidence) Type() protoreflect.MessageType {
	return _fastReflection_Evidence_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Evidence) New() protoreflect.Message {
	return new(fastReflection_Evidence)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Evidence) Interface() protoreflect.ProtoMessage {
	return (*Evidence)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Evidence) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Sum != nil {
		switch o := x.Sum.(type) {
		case *Evidence_DuplicateVoteEvidence:
			v := o.DuplicateVoteEvidence
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_Evidence_duplicate_vote_evidence, value) {
				return
			}
		case *Evidence_LightClientAttackEvidence:
			v := o.LightClientAttackEvidence
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_Evidence_light_client_attack_evidence, value) {
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
func (x *fastReflection_Evidence) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.Evidence.duplicate_vote_evidence":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*Evidence_DuplicateVoteEvidence); ok {
			return true
		} else {
			return false
		}
	case "tendermint.types.Evidence.light_client_attack_evidence":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*Evidence_LightClientAttackEvidence); ok {
			return true
		} else {
			return false
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Evidence"))
		}
		panic(fmt.Errorf("message tendermint.types.Evidence does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Evidence) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.Evidence.duplicate_vote_evidence":
		x.Sum = nil
	case "tendermint.types.Evidence.light_client_attack_evidence":
		x.Sum = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Evidence"))
		}
		panic(fmt.Errorf("message tendermint.types.Evidence does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Evidence) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.Evidence.duplicate_vote_evidence":
		if x.Sum == nil {
			return protoreflect.ValueOfMessage((*DuplicateVoteEvidence)(nil).ProtoReflect())
		} else if v, ok := x.Sum.(*Evidence_DuplicateVoteEvidence); ok {
			return protoreflect.ValueOfMessage(v.DuplicateVoteEvidence.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*DuplicateVoteEvidence)(nil).ProtoReflect())
		}
	case "tendermint.types.Evidence.light_client_attack_evidence":
		if x.Sum == nil {
			return protoreflect.ValueOfMessage((*LightClientAttackEvidence)(nil).ProtoReflect())
		} else if v, ok := x.Sum.(*Evidence_LightClientAttackEvidence); ok {
			return protoreflect.ValueOfMessage(v.LightClientAttackEvidence.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*LightClientAttackEvidence)(nil).ProtoReflect())
		}
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Evidence"))
		}
		panic(fmt.Errorf("message tendermint.types.Evidence does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Evidence) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.Evidence.duplicate_vote_evidence":
		cv := value.Message().Interface().(*DuplicateVoteEvidence)
		x.Sum = &Evidence_DuplicateVoteEvidence{DuplicateVoteEvidence: cv}
	case "tendermint.types.Evidence.light_client_attack_evidence":
		cv := value.Message().Interface().(*LightClientAttackEvidence)
		x.Sum = &Evidence_LightClientAttackEvidence{LightClientAttackEvidence: cv}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Evidence"))
		}
		panic(fmt.Errorf("message tendermint.types.Evidence does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Evidence) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.Evidence.duplicate_vote_evidence":
		if x.Sum == nil {
			value := &DuplicateVoteEvidence{}
			oneofValue := &Evidence_DuplicateVoteEvidence{DuplicateVoteEvidence: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Sum.(type) {
		case *Evidence_DuplicateVoteEvidence:
			return protoreflect.ValueOfMessage(m.DuplicateVoteEvidence.ProtoReflect())
		default:
			value := &DuplicateVoteEvidence{}
			oneofValue := &Evidence_DuplicateVoteEvidence{DuplicateVoteEvidence: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	case "tendermint.types.Evidence.light_client_attack_evidence":
		if x.Sum == nil {
			value := &LightClientAttackEvidence{}
			oneofValue := &Evidence_LightClientAttackEvidence{LightClientAttackEvidence: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Sum.(type) {
		case *Evidence_LightClientAttackEvidence:
			return protoreflect.ValueOfMessage(m.LightClientAttackEvidence.ProtoReflect())
		default:
			value := &LightClientAttackEvidence{}
			oneofValue := &Evidence_LightClientAttackEvidence{LightClientAttackEvidence: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Evidence"))
		}
		panic(fmt.Errorf("message tendermint.types.Evidence does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Evidence) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.Evidence.duplicate_vote_evidence":
		value := &DuplicateVoteEvidence{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.Evidence.light_client_attack_evidence":
		value := &LightClientAttackEvidence{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Evidence"))
		}
		panic(fmt.Errorf("message tendermint.types.Evidence does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Evidence) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "tendermint.types.Evidence.sum":
		if x.Sum == nil {
			return nil
		}
		switch x.Sum.(type) {
		case *Evidence_DuplicateVoteEvidence:
			return x.Descriptor().Fields().ByName("duplicate_vote_evidence")
		case *Evidence_LightClientAttackEvidence:
			return x.Descriptor().Fields().ByName("light_client_attack_evidence")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.Evidence", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Evidence) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Evidence) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Evidence) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Evidence) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Evidence)
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
		case *Evidence_DuplicateVoteEvidence:
			if x == nil {
				break
			}
			l = options.Size(x.DuplicateVoteEvidence)
			n += 1 + l + runtime.Sov(uint64(l))
		case *Evidence_LightClientAttackEvidence:
			if x == nil {
				break
			}
			l = options.Size(x.LightClientAttackEvidence)
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
		x := input.Message.Interface().(*Evidence)
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
		case *Evidence_DuplicateVoteEvidence:
			encoded, err := options.Marshal(x.DuplicateVoteEvidence)
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
		case *Evidence_LightClientAttackEvidence:
			encoded, err := options.Marshal(x.LightClientAttackEvidence)
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
		x := input.Message.Interface().(*Evidence)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Evidence: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Evidence: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DuplicateVoteEvidence", wireType)
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
				v := &DuplicateVoteEvidence{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Sum = &Evidence_DuplicateVoteEvidence{v}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LightClientAttackEvidence", wireType)
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
				v := &LightClientAttackEvidence{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Sum = &Evidence_LightClientAttackEvidence{v}
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
	md_DuplicateVoteEvidence                    protoreflect.MessageDescriptor
	fd_DuplicateVoteEvidence_vote_a             protoreflect.FieldDescriptor
	fd_DuplicateVoteEvidence_vote_b             protoreflect.FieldDescriptor
	fd_DuplicateVoteEvidence_total_voting_power protoreflect.FieldDescriptor
	fd_DuplicateVoteEvidence_validator_power    protoreflect.FieldDescriptor
	fd_DuplicateVoteEvidence_timestamp          protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_evidence_proto_init()
	md_DuplicateVoteEvidence = File_tendermint_types_evidence_proto.Messages().ByName("DuplicateVoteEvidence")
	fd_DuplicateVoteEvidence_vote_a = md_DuplicateVoteEvidence.Fields().ByName("vote_a")
	fd_DuplicateVoteEvidence_vote_b = md_DuplicateVoteEvidence.Fields().ByName("vote_b")
	fd_DuplicateVoteEvidence_total_voting_power = md_DuplicateVoteEvidence.Fields().ByName("total_voting_power")
	fd_DuplicateVoteEvidence_validator_power = md_DuplicateVoteEvidence.Fields().ByName("validator_power")
	fd_DuplicateVoteEvidence_timestamp = md_DuplicateVoteEvidence.Fields().ByName("timestamp")
}

var _ protoreflect.Message = (*fastReflection_DuplicateVoteEvidence)(nil)

type fastReflection_DuplicateVoteEvidence DuplicateVoteEvidence

func (x *DuplicateVoteEvidence) ProtoReflect() protoreflect.Message {
	return (*fastReflection_DuplicateVoteEvidence)(x)
}

func (x *DuplicateVoteEvidence) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_evidence_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_DuplicateVoteEvidence_messageType fastReflection_DuplicateVoteEvidence_messageType
var _ protoreflect.MessageType = fastReflection_DuplicateVoteEvidence_messageType{}

type fastReflection_DuplicateVoteEvidence_messageType struct{}

func (x fastReflection_DuplicateVoteEvidence_messageType) Zero() protoreflect.Message {
	return (*fastReflection_DuplicateVoteEvidence)(nil)
}
func (x fastReflection_DuplicateVoteEvidence_messageType) New() protoreflect.Message {
	return new(fastReflection_DuplicateVoteEvidence)
}
func (x fastReflection_DuplicateVoteEvidence_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_DuplicateVoteEvidence
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_DuplicateVoteEvidence) Descriptor() protoreflect.MessageDescriptor {
	return md_DuplicateVoteEvidence
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_DuplicateVoteEvidence) Type() protoreflect.MessageType {
	return _fastReflection_DuplicateVoteEvidence_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_DuplicateVoteEvidence) New() protoreflect.Message {
	return new(fastReflection_DuplicateVoteEvidence)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_DuplicateVoteEvidence) Interface() protoreflect.ProtoMessage {
	return (*DuplicateVoteEvidence)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_DuplicateVoteEvidence) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.VoteA != nil {
		value := protoreflect.ValueOfMessage(x.VoteA.ProtoReflect())
		if !f(fd_DuplicateVoteEvidence_vote_a, value) {
			return
		}
	}
	if x.VoteB != nil {
		value := protoreflect.ValueOfMessage(x.VoteB.ProtoReflect())
		if !f(fd_DuplicateVoteEvidence_vote_b, value) {
			return
		}
	}
	if x.TotalVotingPower != int64(0) {
		value := protoreflect.ValueOfInt64(x.TotalVotingPower)
		if !f(fd_DuplicateVoteEvidence_total_voting_power, value) {
			return
		}
	}
	if x.ValidatorPower != int64(0) {
		value := protoreflect.ValueOfInt64(x.ValidatorPower)
		if !f(fd_DuplicateVoteEvidence_validator_power, value) {
			return
		}
	}
	if x.Timestamp != nil {
		value := protoreflect.ValueOfMessage(x.Timestamp.ProtoReflect())
		if !f(fd_DuplicateVoteEvidence_timestamp, value) {
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
func (x *fastReflection_DuplicateVoteEvidence) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.DuplicateVoteEvidence.vote_a":
		return x.VoteA != nil
	case "tendermint.types.DuplicateVoteEvidence.vote_b":
		return x.VoteB != nil
	case "tendermint.types.DuplicateVoteEvidence.total_voting_power":
		return x.TotalVotingPower != int64(0)
	case "tendermint.types.DuplicateVoteEvidence.validator_power":
		return x.ValidatorPower != int64(0)
	case "tendermint.types.DuplicateVoteEvidence.timestamp":
		return x.Timestamp != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.DuplicateVoteEvidence"))
		}
		panic(fmt.Errorf("message tendermint.types.DuplicateVoteEvidence does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DuplicateVoteEvidence) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.DuplicateVoteEvidence.vote_a":
		x.VoteA = nil
	case "tendermint.types.DuplicateVoteEvidence.vote_b":
		x.VoteB = nil
	case "tendermint.types.DuplicateVoteEvidence.total_voting_power":
		x.TotalVotingPower = int64(0)
	case "tendermint.types.DuplicateVoteEvidence.validator_power":
		x.ValidatorPower = int64(0)
	case "tendermint.types.DuplicateVoteEvidence.timestamp":
		x.Timestamp = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.DuplicateVoteEvidence"))
		}
		panic(fmt.Errorf("message tendermint.types.DuplicateVoteEvidence does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_DuplicateVoteEvidence) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.DuplicateVoteEvidence.vote_a":
		value := x.VoteA
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.DuplicateVoteEvidence.vote_b":
		value := x.VoteB
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.DuplicateVoteEvidence.total_voting_power":
		value := x.TotalVotingPower
		return protoreflect.ValueOfInt64(value)
	case "tendermint.types.DuplicateVoteEvidence.validator_power":
		value := x.ValidatorPower
		return protoreflect.ValueOfInt64(value)
	case "tendermint.types.DuplicateVoteEvidence.timestamp":
		value := x.Timestamp
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.DuplicateVoteEvidence"))
		}
		panic(fmt.Errorf("message tendermint.types.DuplicateVoteEvidence does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_DuplicateVoteEvidence) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.DuplicateVoteEvidence.vote_a":
		x.VoteA = value.Message().Interface().(*Vote)
	case "tendermint.types.DuplicateVoteEvidence.vote_b":
		x.VoteB = value.Message().Interface().(*Vote)
	case "tendermint.types.DuplicateVoteEvidence.total_voting_power":
		x.TotalVotingPower = value.Int()
	case "tendermint.types.DuplicateVoteEvidence.validator_power":
		x.ValidatorPower = value.Int()
	case "tendermint.types.DuplicateVoteEvidence.timestamp":
		x.Timestamp = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.DuplicateVoteEvidence"))
		}
		panic(fmt.Errorf("message tendermint.types.DuplicateVoteEvidence does not contain field %s", fd.FullName()))
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
func (x *fastReflection_DuplicateVoteEvidence) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.DuplicateVoteEvidence.vote_a":
		if x.VoteA == nil {
			x.VoteA = new(Vote)
		}
		return protoreflect.ValueOfMessage(x.VoteA.ProtoReflect())
	case "tendermint.types.DuplicateVoteEvidence.vote_b":
		if x.VoteB == nil {
			x.VoteB = new(Vote)
		}
		return protoreflect.ValueOfMessage(x.VoteB.ProtoReflect())
	case "tendermint.types.DuplicateVoteEvidence.timestamp":
		if x.Timestamp == nil {
			x.Timestamp = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.Timestamp.ProtoReflect())
	case "tendermint.types.DuplicateVoteEvidence.total_voting_power":
		panic(fmt.Errorf("field total_voting_power of message tendermint.types.DuplicateVoteEvidence is not mutable"))
	case "tendermint.types.DuplicateVoteEvidence.validator_power":
		panic(fmt.Errorf("field validator_power of message tendermint.types.DuplicateVoteEvidence is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.DuplicateVoteEvidence"))
		}
		panic(fmt.Errorf("message tendermint.types.DuplicateVoteEvidence does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_DuplicateVoteEvidence) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.DuplicateVoteEvidence.vote_a":
		m := new(Vote)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.DuplicateVoteEvidence.vote_b":
		m := new(Vote)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.DuplicateVoteEvidence.total_voting_power":
		return protoreflect.ValueOfInt64(int64(0))
	case "tendermint.types.DuplicateVoteEvidence.validator_power":
		return protoreflect.ValueOfInt64(int64(0))
	case "tendermint.types.DuplicateVoteEvidence.timestamp":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.DuplicateVoteEvidence"))
		}
		panic(fmt.Errorf("message tendermint.types.DuplicateVoteEvidence does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_DuplicateVoteEvidence) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.DuplicateVoteEvidence", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_DuplicateVoteEvidence) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DuplicateVoteEvidence) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_DuplicateVoteEvidence) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_DuplicateVoteEvidence) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*DuplicateVoteEvidence)
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
		if x.VoteA != nil {
			l = options.Size(x.VoteA)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.VoteB != nil {
			l = options.Size(x.VoteB)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.TotalVotingPower != 0 {
			n += 1 + runtime.Sov(uint64(x.TotalVotingPower))
		}
		if x.ValidatorPower != 0 {
			n += 1 + runtime.Sov(uint64(x.ValidatorPower))
		}
		if x.Timestamp != nil {
			l = options.Size(x.Timestamp)
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
		x := input.Message.Interface().(*DuplicateVoteEvidence)
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
		if x.Timestamp != nil {
			encoded, err := options.Marshal(x.Timestamp)
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
		if x.ValidatorPower != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ValidatorPower))
			i--
			dAtA[i] = 0x20
		}
		if x.TotalVotingPower != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.TotalVotingPower))
			i--
			dAtA[i] = 0x18
		}
		if x.VoteB != nil {
			encoded, err := options.Marshal(x.VoteB)
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
		if x.VoteA != nil {
			encoded, err := options.Marshal(x.VoteA)
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
		x := input.Message.Interface().(*DuplicateVoteEvidence)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DuplicateVoteEvidence: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DuplicateVoteEvidence: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field VoteA", wireType)
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
				if x.VoteA == nil {
					x.VoteA = &Vote{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.VoteA); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field VoteB", wireType)
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
				if x.VoteB == nil {
					x.VoteB = &Vote{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.VoteB); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TotalVotingPower", wireType)
				}
				x.TotalVotingPower = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.TotalVotingPower |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorPower", wireType)
				}
				x.ValidatorPower = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ValidatorPower |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
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
				if x.Timestamp == nil {
					x.Timestamp = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Timestamp); err != nil {
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

var _ protoreflect.List = (*_LightClientAttackEvidence_3_list)(nil)

type _LightClientAttackEvidence_3_list struct {
	list *[]*Validator
}

func (x *_LightClientAttackEvidence_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_LightClientAttackEvidence_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_LightClientAttackEvidence_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Validator)
	(*x.list)[i] = concreteValue
}

func (x *_LightClientAttackEvidence_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Validator)
	*x.list = append(*x.list, concreteValue)
}

func (x *_LightClientAttackEvidence_3_list) AppendMutable() protoreflect.Value {
	v := new(Validator)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_LightClientAttackEvidence_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_LightClientAttackEvidence_3_list) NewElement() protoreflect.Value {
	v := new(Validator)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_LightClientAttackEvidence_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_LightClientAttackEvidence                      protoreflect.MessageDescriptor
	fd_LightClientAttackEvidence_conflicting_block    protoreflect.FieldDescriptor
	fd_LightClientAttackEvidence_common_height        protoreflect.FieldDescriptor
	fd_LightClientAttackEvidence_byzantine_validators protoreflect.FieldDescriptor
	fd_LightClientAttackEvidence_total_voting_power   protoreflect.FieldDescriptor
	fd_LightClientAttackEvidence_timestamp            protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_evidence_proto_init()
	md_LightClientAttackEvidence = File_tendermint_types_evidence_proto.Messages().ByName("LightClientAttackEvidence")
	fd_LightClientAttackEvidence_conflicting_block = md_LightClientAttackEvidence.Fields().ByName("conflicting_block")
	fd_LightClientAttackEvidence_common_height = md_LightClientAttackEvidence.Fields().ByName("common_height")
	fd_LightClientAttackEvidence_byzantine_validators = md_LightClientAttackEvidence.Fields().ByName("byzantine_validators")
	fd_LightClientAttackEvidence_total_voting_power = md_LightClientAttackEvidence.Fields().ByName("total_voting_power")
	fd_LightClientAttackEvidence_timestamp = md_LightClientAttackEvidence.Fields().ByName("timestamp")
}

var _ protoreflect.Message = (*fastReflection_LightClientAttackEvidence)(nil)

type fastReflection_LightClientAttackEvidence LightClientAttackEvidence

func (x *LightClientAttackEvidence) ProtoReflect() protoreflect.Message {
	return (*fastReflection_LightClientAttackEvidence)(x)
}

func (x *LightClientAttackEvidence) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_evidence_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_LightClientAttackEvidence_messageType fastReflection_LightClientAttackEvidence_messageType
var _ protoreflect.MessageType = fastReflection_LightClientAttackEvidence_messageType{}

type fastReflection_LightClientAttackEvidence_messageType struct{}

func (x fastReflection_LightClientAttackEvidence_messageType) Zero() protoreflect.Message {
	return (*fastReflection_LightClientAttackEvidence)(nil)
}
func (x fastReflection_LightClientAttackEvidence_messageType) New() protoreflect.Message {
	return new(fastReflection_LightClientAttackEvidence)
}
func (x fastReflection_LightClientAttackEvidence_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_LightClientAttackEvidence
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_LightClientAttackEvidence) Descriptor() protoreflect.MessageDescriptor {
	return md_LightClientAttackEvidence
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_LightClientAttackEvidence) Type() protoreflect.MessageType {
	return _fastReflection_LightClientAttackEvidence_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_LightClientAttackEvidence) New() protoreflect.Message {
	return new(fastReflection_LightClientAttackEvidence)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_LightClientAttackEvidence) Interface() protoreflect.ProtoMessage {
	return (*LightClientAttackEvidence)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_LightClientAttackEvidence) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ConflictingBlock != nil {
		value := protoreflect.ValueOfMessage(x.ConflictingBlock.ProtoReflect())
		if !f(fd_LightClientAttackEvidence_conflicting_block, value) {
			return
		}
	}
	if x.CommonHeight != int64(0) {
		value := protoreflect.ValueOfInt64(x.CommonHeight)
		if !f(fd_LightClientAttackEvidence_common_height, value) {
			return
		}
	}
	if len(x.ByzantineValidators) != 0 {
		value := protoreflect.ValueOfList(&_LightClientAttackEvidence_3_list{list: &x.ByzantineValidators})
		if !f(fd_LightClientAttackEvidence_byzantine_validators, value) {
			return
		}
	}
	if x.TotalVotingPower != int64(0) {
		value := protoreflect.ValueOfInt64(x.TotalVotingPower)
		if !f(fd_LightClientAttackEvidence_total_voting_power, value) {
			return
		}
	}
	if x.Timestamp != nil {
		value := protoreflect.ValueOfMessage(x.Timestamp.ProtoReflect())
		if !f(fd_LightClientAttackEvidence_timestamp, value) {
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
func (x *fastReflection_LightClientAttackEvidence) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.LightClientAttackEvidence.conflicting_block":
		return x.ConflictingBlock != nil
	case "tendermint.types.LightClientAttackEvidence.common_height":
		return x.CommonHeight != int64(0)
	case "tendermint.types.LightClientAttackEvidence.byzantine_validators":
		return len(x.ByzantineValidators) != 0
	case "tendermint.types.LightClientAttackEvidence.total_voting_power":
		return x.TotalVotingPower != int64(0)
	case "tendermint.types.LightClientAttackEvidence.timestamp":
		return x.Timestamp != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.LightClientAttackEvidence"))
		}
		panic(fmt.Errorf("message tendermint.types.LightClientAttackEvidence does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_LightClientAttackEvidence) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.LightClientAttackEvidence.conflicting_block":
		x.ConflictingBlock = nil
	case "tendermint.types.LightClientAttackEvidence.common_height":
		x.CommonHeight = int64(0)
	case "tendermint.types.LightClientAttackEvidence.byzantine_validators":
		x.ByzantineValidators = nil
	case "tendermint.types.LightClientAttackEvidence.total_voting_power":
		x.TotalVotingPower = int64(0)
	case "tendermint.types.LightClientAttackEvidence.timestamp":
		x.Timestamp = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.LightClientAttackEvidence"))
		}
		panic(fmt.Errorf("message tendermint.types.LightClientAttackEvidence does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_LightClientAttackEvidence) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.LightClientAttackEvidence.conflicting_block":
		value := x.ConflictingBlock
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.LightClientAttackEvidence.common_height":
		value := x.CommonHeight
		return protoreflect.ValueOfInt64(value)
	case "tendermint.types.LightClientAttackEvidence.byzantine_validators":
		if len(x.ByzantineValidators) == 0 {
			return protoreflect.ValueOfList(&_LightClientAttackEvidence_3_list{})
		}
		listValue := &_LightClientAttackEvidence_3_list{list: &x.ByzantineValidators}
		return protoreflect.ValueOfList(listValue)
	case "tendermint.types.LightClientAttackEvidence.total_voting_power":
		value := x.TotalVotingPower
		return protoreflect.ValueOfInt64(value)
	case "tendermint.types.LightClientAttackEvidence.timestamp":
		value := x.Timestamp
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.LightClientAttackEvidence"))
		}
		panic(fmt.Errorf("message tendermint.types.LightClientAttackEvidence does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_LightClientAttackEvidence) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.LightClientAttackEvidence.conflicting_block":
		x.ConflictingBlock = value.Message().Interface().(*LightBlock)
	case "tendermint.types.LightClientAttackEvidence.common_height":
		x.CommonHeight = value.Int()
	case "tendermint.types.LightClientAttackEvidence.byzantine_validators":
		lv := value.List()
		clv := lv.(*_LightClientAttackEvidence_3_list)
		x.ByzantineValidators = *clv.list
	case "tendermint.types.LightClientAttackEvidence.total_voting_power":
		x.TotalVotingPower = value.Int()
	case "tendermint.types.LightClientAttackEvidence.timestamp":
		x.Timestamp = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.LightClientAttackEvidence"))
		}
		panic(fmt.Errorf("message tendermint.types.LightClientAttackEvidence does not contain field %s", fd.FullName()))
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
func (x *fastReflection_LightClientAttackEvidence) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.LightClientAttackEvidence.conflicting_block":
		if x.ConflictingBlock == nil {
			x.ConflictingBlock = new(LightBlock)
		}
		return protoreflect.ValueOfMessage(x.ConflictingBlock.ProtoReflect())
	case "tendermint.types.LightClientAttackEvidence.byzantine_validators":
		if x.ByzantineValidators == nil {
			x.ByzantineValidators = []*Validator{}
		}
		value := &_LightClientAttackEvidence_3_list{list: &x.ByzantineValidators}
		return protoreflect.ValueOfList(value)
	case "tendermint.types.LightClientAttackEvidence.timestamp":
		if x.Timestamp == nil {
			x.Timestamp = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.Timestamp.ProtoReflect())
	case "tendermint.types.LightClientAttackEvidence.common_height":
		panic(fmt.Errorf("field common_height of message tendermint.types.LightClientAttackEvidence is not mutable"))
	case "tendermint.types.LightClientAttackEvidence.total_voting_power":
		panic(fmt.Errorf("field total_voting_power of message tendermint.types.LightClientAttackEvidence is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.LightClientAttackEvidence"))
		}
		panic(fmt.Errorf("message tendermint.types.LightClientAttackEvidence does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_LightClientAttackEvidence) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.LightClientAttackEvidence.conflicting_block":
		m := new(LightBlock)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.LightClientAttackEvidence.common_height":
		return protoreflect.ValueOfInt64(int64(0))
	case "tendermint.types.LightClientAttackEvidence.byzantine_validators":
		list := []*Validator{}
		return protoreflect.ValueOfList(&_LightClientAttackEvidence_3_list{list: &list})
	case "tendermint.types.LightClientAttackEvidence.total_voting_power":
		return protoreflect.ValueOfInt64(int64(0))
	case "tendermint.types.LightClientAttackEvidence.timestamp":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.LightClientAttackEvidence"))
		}
		panic(fmt.Errorf("message tendermint.types.LightClientAttackEvidence does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_LightClientAttackEvidence) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.LightClientAttackEvidence", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_LightClientAttackEvidence) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_LightClientAttackEvidence) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_LightClientAttackEvidence) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_LightClientAttackEvidence) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*LightClientAttackEvidence)
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
		if x.ConflictingBlock != nil {
			l = options.Size(x.ConflictingBlock)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.CommonHeight != 0 {
			n += 1 + runtime.Sov(uint64(x.CommonHeight))
		}
		if len(x.ByzantineValidators) > 0 {
			for _, e := range x.ByzantineValidators {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.TotalVotingPower != 0 {
			n += 1 + runtime.Sov(uint64(x.TotalVotingPower))
		}
		if x.Timestamp != nil {
			l = options.Size(x.Timestamp)
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
		x := input.Message.Interface().(*LightClientAttackEvidence)
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
		if x.Timestamp != nil {
			encoded, err := options.Marshal(x.Timestamp)
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
		if x.TotalVotingPower != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.TotalVotingPower))
			i--
			dAtA[i] = 0x20
		}
		if len(x.ByzantineValidators) > 0 {
			for iNdEx := len(x.ByzantineValidators) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.ByzantineValidators[iNdEx])
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
		if x.CommonHeight != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.CommonHeight))
			i--
			dAtA[i] = 0x10
		}
		if x.ConflictingBlock != nil {
			encoded, err := options.Marshal(x.ConflictingBlock)
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
		x := input.Message.Interface().(*LightClientAttackEvidence)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: LightClientAttackEvidence: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: LightClientAttackEvidence: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ConflictingBlock", wireType)
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
				if x.ConflictingBlock == nil {
					x.ConflictingBlock = &LightBlock{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ConflictingBlock); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CommonHeight", wireType)
				}
				x.CommonHeight = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.CommonHeight |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ByzantineValidators", wireType)
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
				x.ByzantineValidators = append(x.ByzantineValidators, &Validator{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ByzantineValidators[len(x.ByzantineValidators)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TotalVotingPower", wireType)
				}
				x.TotalVotingPower = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.TotalVotingPower |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
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
				if x.Timestamp == nil {
					x.Timestamp = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Timestamp); err != nil {
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

var _ protoreflect.List = (*_EvidenceList_1_list)(nil)

type _EvidenceList_1_list struct {
	list *[]*Evidence
}

func (x *_EvidenceList_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_EvidenceList_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_EvidenceList_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Evidence)
	(*x.list)[i] = concreteValue
}

func (x *_EvidenceList_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Evidence)
	*x.list = append(*x.list, concreteValue)
}

func (x *_EvidenceList_1_list) AppendMutable() protoreflect.Value {
	v := new(Evidence)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_EvidenceList_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_EvidenceList_1_list) NewElement() protoreflect.Value {
	v := new(Evidence)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_EvidenceList_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_EvidenceList          protoreflect.MessageDescriptor
	fd_EvidenceList_evidence protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_evidence_proto_init()
	md_EvidenceList = File_tendermint_types_evidence_proto.Messages().ByName("EvidenceList")
	fd_EvidenceList_evidence = md_EvidenceList.Fields().ByName("evidence")
}

var _ protoreflect.Message = (*fastReflection_EvidenceList)(nil)

type fastReflection_EvidenceList EvidenceList

func (x *EvidenceList) ProtoReflect() protoreflect.Message {
	return (*fastReflection_EvidenceList)(x)
}

func (x *EvidenceList) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_evidence_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_EvidenceList_messageType fastReflection_EvidenceList_messageType
var _ protoreflect.MessageType = fastReflection_EvidenceList_messageType{}

type fastReflection_EvidenceList_messageType struct{}

func (x fastReflection_EvidenceList_messageType) Zero() protoreflect.Message {
	return (*fastReflection_EvidenceList)(nil)
}
func (x fastReflection_EvidenceList_messageType) New() protoreflect.Message {
	return new(fastReflection_EvidenceList)
}
func (x fastReflection_EvidenceList_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_EvidenceList
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_EvidenceList) Descriptor() protoreflect.MessageDescriptor {
	return md_EvidenceList
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_EvidenceList) Type() protoreflect.MessageType {
	return _fastReflection_EvidenceList_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_EvidenceList) New() protoreflect.Message {
	return new(fastReflection_EvidenceList)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_EvidenceList) Interface() protoreflect.ProtoMessage {
	return (*EvidenceList)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_EvidenceList) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Evidence) != 0 {
		value := protoreflect.ValueOfList(&_EvidenceList_1_list{list: &x.Evidence})
		if !f(fd_EvidenceList_evidence, value) {
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
func (x *fastReflection_EvidenceList) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.EvidenceList.evidence":
		return len(x.Evidence) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.EvidenceList"))
		}
		panic(fmt.Errorf("message tendermint.types.EvidenceList does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EvidenceList) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.EvidenceList.evidence":
		x.Evidence = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.EvidenceList"))
		}
		panic(fmt.Errorf("message tendermint.types.EvidenceList does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_EvidenceList) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.EvidenceList.evidence":
		if len(x.Evidence) == 0 {
			return protoreflect.ValueOfList(&_EvidenceList_1_list{})
		}
		listValue := &_EvidenceList_1_list{list: &x.Evidence}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.EvidenceList"))
		}
		panic(fmt.Errorf("message tendermint.types.EvidenceList does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_EvidenceList) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.EvidenceList.evidence":
		lv := value.List()
		clv := lv.(*_EvidenceList_1_list)
		x.Evidence = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.EvidenceList"))
		}
		panic(fmt.Errorf("message tendermint.types.EvidenceList does not contain field %s", fd.FullName()))
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
func (x *fastReflection_EvidenceList) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.EvidenceList.evidence":
		if x.Evidence == nil {
			x.Evidence = []*Evidence{}
		}
		value := &_EvidenceList_1_list{list: &x.Evidence}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.EvidenceList"))
		}
		panic(fmt.Errorf("message tendermint.types.EvidenceList does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_EvidenceList) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.EvidenceList.evidence":
		list := []*Evidence{}
		return protoreflect.ValueOfList(&_EvidenceList_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.EvidenceList"))
		}
		panic(fmt.Errorf("message tendermint.types.EvidenceList does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_EvidenceList) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.EvidenceList", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_EvidenceList) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_EvidenceList) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_EvidenceList) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_EvidenceList) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*EvidenceList)
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
		if len(x.Evidence) > 0 {
			for _, e := range x.Evidence {
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
		x := input.Message.Interface().(*EvidenceList)
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
		if len(x.Evidence) > 0 {
			for iNdEx := len(x.Evidence) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Evidence[iNdEx])
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
		x := input.Message.Interface().(*EvidenceList)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EvidenceList: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: EvidenceList: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Evidence", wireType)
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
				x.Evidence = append(x.Evidence, &Evidence{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Evidence[len(x.Evidence)-1]); err != nil {
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
// source: tendermint/types/evidence.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Evidence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Sum:
	//	*Evidence_DuplicateVoteEvidence
	//	*Evidence_LightClientAttackEvidence
	Sum isEvidence_Sum `protobuf_oneof:"sum"`
}

func (x *Evidence) Reset() {
	*x = Evidence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_evidence_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Evidence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Evidence) ProtoMessage() {}

// Deprecated: Use Evidence.ProtoReflect.Descriptor instead.
func (*Evidence) Descriptor() ([]byte, []int) {
	return file_tendermint_types_evidence_proto_rawDescGZIP(), []int{0}
}

func (x *Evidence) GetSum() isEvidence_Sum {
	if x != nil {
		return x.Sum
	}
	return nil
}

func (x *Evidence) GetDuplicateVoteEvidence() *DuplicateVoteEvidence {
	if x, ok := x.GetSum().(*Evidence_DuplicateVoteEvidence); ok {
		return x.DuplicateVoteEvidence
	}
	return nil
}

func (x *Evidence) GetLightClientAttackEvidence() *LightClientAttackEvidence {
	if x, ok := x.GetSum().(*Evidence_LightClientAttackEvidence); ok {
		return x.LightClientAttackEvidence
	}
	return nil
}

type isEvidence_Sum interface {
	isEvidence_Sum()
}

type Evidence_DuplicateVoteEvidence struct {
	DuplicateVoteEvidence *DuplicateVoteEvidence `protobuf:"bytes,1,opt,name=duplicate_vote_evidence,json=duplicateVoteEvidence,proto3,oneof"`
}

type Evidence_LightClientAttackEvidence struct {
	LightClientAttackEvidence *LightClientAttackEvidence `protobuf:"bytes,2,opt,name=light_client_attack_evidence,json=lightClientAttackEvidence,proto3,oneof"`
}

func (*Evidence_DuplicateVoteEvidence) isEvidence_Sum() {}

func (*Evidence_LightClientAttackEvidence) isEvidence_Sum() {}

// DuplicateVoteEvidence contains evidence of a validator signed two conflicting votes.
type DuplicateVoteEvidence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoteA            *Vote                  `protobuf:"bytes,1,opt,name=vote_a,json=voteA,proto3" json:"vote_a,omitempty"`
	VoteB            *Vote                  `protobuf:"bytes,2,opt,name=vote_b,json=voteB,proto3" json:"vote_b,omitempty"`
	TotalVotingPower int64                  `protobuf:"varint,3,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"total_voting_power,omitempty"`
	ValidatorPower   int64                  `protobuf:"varint,4,opt,name=validator_power,json=validatorPower,proto3" json:"validator_power,omitempty"`
	Timestamp        *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *DuplicateVoteEvidence) Reset() {
	*x = DuplicateVoteEvidence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_evidence_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DuplicateVoteEvidence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DuplicateVoteEvidence) ProtoMessage() {}

// Deprecated: Use DuplicateVoteEvidence.ProtoReflect.Descriptor instead.
func (*DuplicateVoteEvidence) Descriptor() ([]byte, []int) {
	return file_tendermint_types_evidence_proto_rawDescGZIP(), []int{1}
}

func (x *DuplicateVoteEvidence) GetVoteA() *Vote {
	if x != nil {
		return x.VoteA
	}
	return nil
}

func (x *DuplicateVoteEvidence) GetVoteB() *Vote {
	if x != nil {
		return x.VoteB
	}
	return nil
}

func (x *DuplicateVoteEvidence) GetTotalVotingPower() int64 {
	if x != nil {
		return x.TotalVotingPower
	}
	return 0
}

func (x *DuplicateVoteEvidence) GetValidatorPower() int64 {
	if x != nil {
		return x.ValidatorPower
	}
	return 0
}

func (x *DuplicateVoteEvidence) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

// LightClientAttackEvidence contains evidence of a set of validators attempting to mislead a light client.
type LightClientAttackEvidence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConflictingBlock    *LightBlock            `protobuf:"bytes,1,opt,name=conflicting_block,json=conflictingBlock,proto3" json:"conflicting_block,omitempty"`
	CommonHeight        int64                  `protobuf:"varint,2,opt,name=common_height,json=commonHeight,proto3" json:"common_height,omitempty"`
	ByzantineValidators []*Validator           `protobuf:"bytes,3,rep,name=byzantine_validators,json=byzantineValidators,proto3" json:"byzantine_validators,omitempty"`
	TotalVotingPower    int64                  `protobuf:"varint,4,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"total_voting_power,omitempty"`
	Timestamp           *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *LightClientAttackEvidence) Reset() {
	*x = LightClientAttackEvidence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_evidence_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LightClientAttackEvidence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LightClientAttackEvidence) ProtoMessage() {}

// Deprecated: Use LightClientAttackEvidence.ProtoReflect.Descriptor instead.
func (*LightClientAttackEvidence) Descriptor() ([]byte, []int) {
	return file_tendermint_types_evidence_proto_rawDescGZIP(), []int{2}
}

func (x *LightClientAttackEvidence) GetConflictingBlock() *LightBlock {
	if x != nil {
		return x.ConflictingBlock
	}
	return nil
}

func (x *LightClientAttackEvidence) GetCommonHeight() int64 {
	if x != nil {
		return x.CommonHeight
	}
	return 0
}

func (x *LightClientAttackEvidence) GetByzantineValidators() []*Validator {
	if x != nil {
		return x.ByzantineValidators
	}
	return nil
}

func (x *LightClientAttackEvidence) GetTotalVotingPower() int64 {
	if x != nil {
		return x.TotalVotingPower
	}
	return 0
}

func (x *LightClientAttackEvidence) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type EvidenceList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Evidence []*Evidence `protobuf:"bytes,1,rep,name=evidence,proto3" json:"evidence,omitempty"`
}

func (x *EvidenceList) Reset() {
	*x = EvidenceList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_evidence_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvidenceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvidenceList) ProtoMessage() {}

// Deprecated: Use EvidenceList.ProtoReflect.Descriptor instead.
func (*EvidenceList) Descriptor() ([]byte, []int) {
	return file_tendermint_types_evidence_proto_rawDescGZIP(), []int{3}
}

func (x *EvidenceList) GetEvidence() []*Evidence {
	if x != nil {
		return x.Evidence
	}
	return nil
}

var File_tendermint_types_evidence_proto protoreflect.FileDescriptor

var file_tendermint_types_evidence_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x74, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x01, 0x0a, 0x08, 0x45,
	0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x17, 0x64, 0x75, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x75, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63,
	0x65, 0x48, 0x00, 0x52, 0x15, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x56, 0x6f,
	0x74, 0x65, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x1c, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x6b, 0x5f, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x6b, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x48, 0x00, 0x52,
	0x19, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x6b, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x73, 0x75,
	0x6d, 0x22, 0x90, 0x02, 0x0a, 0x15, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x56,
	0x6f, 0x74, 0x65, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x76,
	0x6f, 0x74, 0x65, 0x5f, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56,
	0x6f, 0x74, 0x65, 0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x41, 0x12, 0x2d, 0x0a, 0x06, 0x76, 0x6f,
	0x74, 0x65, 0x5f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x42, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x6f, 0x74, 0x69,
	0x6e, 0x67, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x6f, 0x77, 0x65, 0x72,
	0x12, 0x42, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x08, 0xc8, 0xde, 0x1f, 0x00, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0xcd, 0x02, 0x0a, 0x19, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x49, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x69, 0x6e,
	0x67, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x10, 0x63, 0x6f, 0x6e,
	0x66, 0x6c, 0x69, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x4e, 0x0a, 0x14, 0x62, 0x79, 0x7a, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x5f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x13, 0x62,
	0x79, 0x7a, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x76, 0x6f, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x77, 0x65, 0x72,
	0x12, 0x42, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x08, 0xc8, 0xde, 0x1f, 0x00, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0x4c, 0x0a, 0x0c, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x08, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x08, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x42, 0xb9, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x0d, 0x45, 0x76, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0xa2,
	0x02, 0x03, 0x54, 0x54, 0x58, 0xaa, 0x02, 0x10, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0xca, 0x02, 0x10, 0x54, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x74, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0xe2, 0x02, 0x1c, 0x54, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x54, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tendermint_types_evidence_proto_rawDescOnce sync.Once
	file_tendermint_types_evidence_proto_rawDescData = file_tendermint_types_evidence_proto_rawDesc
)

func file_tendermint_types_evidence_proto_rawDescGZIP() []byte {
	file_tendermint_types_evidence_proto_rawDescOnce.Do(func() {
		file_tendermint_types_evidence_proto_rawDescData = protoimpl.X.CompressGZIP(file_tendermint_types_evidence_proto_rawDescData)
	})
	return file_tendermint_types_evidence_proto_rawDescData
}

var file_tendermint_types_evidence_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tendermint_types_evidence_proto_goTypes = []interface{}{
	(*Evidence)(nil),                  // 0: tendermint.types.Evidence
	(*DuplicateVoteEvidence)(nil),     // 1: tendermint.types.DuplicateVoteEvidence
	(*LightClientAttackEvidence)(nil), // 2: tendermint.types.LightClientAttackEvidence
	(*EvidenceList)(nil),              // 3: tendermint.types.EvidenceList
	(*Vote)(nil),                      // 4: tendermint.types.Vote
	(*timestamppb.Timestamp)(nil),     // 5: google.protobuf.Timestamp
	(*LightBlock)(nil),                // 6: tendermint.types.LightBlock
	(*Validator)(nil),                 // 7: tendermint.types.Validator
}
var file_tendermint_types_evidence_proto_depIdxs = []int32{
	1, // 0: tendermint.types.Evidence.duplicate_vote_evidence:type_name -> tendermint.types.DuplicateVoteEvidence
	2, // 1: tendermint.types.Evidence.light_client_attack_evidence:type_name -> tendermint.types.LightClientAttackEvidence
	4, // 2: tendermint.types.DuplicateVoteEvidence.vote_a:type_name -> tendermint.types.Vote
	4, // 3: tendermint.types.DuplicateVoteEvidence.vote_b:type_name -> tendermint.types.Vote
	5, // 4: tendermint.types.DuplicateVoteEvidence.timestamp:type_name -> google.protobuf.Timestamp
	6, // 5: tendermint.types.LightClientAttackEvidence.conflicting_block:type_name -> tendermint.types.LightBlock
	7, // 6: tendermint.types.LightClientAttackEvidence.byzantine_validators:type_name -> tendermint.types.Validator
	5, // 7: tendermint.types.LightClientAttackEvidence.timestamp:type_name -> google.protobuf.Timestamp
	0, // 8: tendermint.types.EvidenceList.evidence:type_name -> tendermint.types.Evidence
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_tendermint_types_evidence_proto_init() }
func file_tendermint_types_evidence_proto_init() {
	if File_tendermint_types_evidence_proto != nil {
		return
	}
	file_tendermint_types_types_proto_init()
	file_tendermint_types_validator_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tendermint_types_evidence_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Evidence); i {
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
		file_tendermint_types_evidence_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DuplicateVoteEvidence); i {
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
		file_tendermint_types_evidence_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LightClientAttackEvidence); i {
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
		file_tendermint_types_evidence_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvidenceList); i {
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
	file_tendermint_types_evidence_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Evidence_DuplicateVoteEvidence)(nil),
		(*Evidence_LightClientAttackEvidence)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tendermint_types_evidence_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tendermint_types_evidence_proto_goTypes,
		DependencyIndexes: file_tendermint_types_evidence_proto_depIdxs,
		MessageInfos:      file_tendermint_types_evidence_proto_msgTypes,
	}.Build()
	File_tendermint_types_evidence_proto = out.File
	file_tendermint_types_evidence_proto_rawDesc = nil
	file_tendermint_types_evidence_proto_goTypes = nil
	file_tendermint_types_evidence_proto_depIdxs = nil
}
