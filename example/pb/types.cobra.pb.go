// Code generated by protoc-gen-cobra. DO NOT EDIT.

package pb

import (
	client "github.com/NathanBaulch/protoc-gen-cobra/client"
	flag "github.com/NathanBaulch/protoc-gen-cobra/flag"
	iocodec "github.com/NathanBaulch/protoc-gen-cobra/iocodec"
	proto "github.com/golang/protobuf/proto"
	cobra "github.com/spf13/cobra"
	pflag "github.com/spf13/pflag"
	grpc "google.golang.org/grpc"
	strconv "strconv"
	strings "strings"
)

func TypesClientCommand(cfgs ...*client.Config) *cobra.Command {
	cfg := client.DefaultConfig
	if len(cfgs) > 0 {
		cfg = cfgs[0]
	}
	cmd := &cobra.Command{
		Use:   "types",
		Short: "Types service client",
		Long:  "",
	}
	cfg.BindFlags(cmd.PersistentFlags())
	d := &client.Dialer{Config: cfg}
	cmd.AddCommand(
		_TypesEchoCommand(d),
	)
	return cmd
}

func _TypesEchoCommand(d *client.Dialer) *cobra.Command {
	req := &Sound{}

	cmd := &cobra.Command{
		Use:   "echo",
		Short: "Echo RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if d.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), d.EnvVarPrefix); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), d.EnvVarPrefix, "TYPES", "ECHO"); err != nil {
					return err
				}
			}
			return d.RoundTrip(cmd.Context(), func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewTypesClient(cc)
				v := &Sound{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Echo(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().Float64Var(&req.Double, "double", 0, "")
	cmd.PersistentFlags().Float32Var(&req.Float, "float", 0, "")
	cmd.PersistentFlags().Int32Var(&req.Int32, "int32", 0, "")
	cmd.PersistentFlags().Int64Var(&req.Int64, "int64", 0, "")
	cmd.PersistentFlags().Uint32Var(&req.Uint32, "uint32", 0, "")
	cmd.PersistentFlags().Uint64Var(&req.Uint64, "uint64", 0, "")
	cmd.PersistentFlags().Int32Var(&req.Sint32, "sint32", 0, "")
	cmd.PersistentFlags().Int64Var(&req.Sint64, "sint64", 0, "")
	cmd.PersistentFlags().Uint32Var(&req.Fixed32, "fixed32", 0, "")
	cmd.PersistentFlags().Uint64Var(&req.Fixed64, "fixed64", 0, "")
	cmd.PersistentFlags().Int32Var(&req.Sfixed32, "sfixed32", 0, "")
	cmd.PersistentFlags().Int64Var(&req.Sfixed64, "sfixed64", 0, "")
	cmd.PersistentFlags().BoolVar(&req.Bool, "bool", false, "")
	cmd.PersistentFlags().StringVar(&req.String_, "string_", "", "")
	cmd.PersistentFlags().BytesBase64Var(&req.Bytes, "bytes", nil, "")
	_Sound_NestedEnumVar(cmd.PersistentFlags(), &req.NestedEnum, "nestedenum", "")
	_GlobalEnumVar(cmd.PersistentFlags(), &req.GlobalEnum, "globalenum", "")
	cmd.PersistentFlags().Float64SliceVar(&req.ListDouble, "listdouble", nil, "")
	cmd.PersistentFlags().Float32SliceVar(&req.ListFloat, "listfloat", nil, "")
	cmd.PersistentFlags().Int32SliceVar(&req.ListInt32, "listint32", nil, "")
	cmd.PersistentFlags().Int64SliceVar(&req.ListInt64, "listint64", nil, "")
	flag.Uint32SliceVar(cmd.PersistentFlags(), &req.ListUint32, "listuint32", "")
	flag.Uint64SliceVar(cmd.PersistentFlags(), &req.ListUint64, "listuint64", "")
	cmd.PersistentFlags().Int32SliceVar(&req.ListSint32, "listsint32", nil, "")
	cmd.PersistentFlags().Int64SliceVar(&req.ListSint64, "listsint64", nil, "")
	flag.Uint32SliceVar(cmd.PersistentFlags(), &req.ListFixed32, "listfixed32", "")
	flag.Uint64SliceVar(cmd.PersistentFlags(), &req.ListFixed64, "listfixed64", "")
	cmd.PersistentFlags().Int32SliceVar(&req.ListSfixed32, "listsfixed32", nil, "")
	cmd.PersistentFlags().Int64SliceVar(&req.ListSfixed64, "listsfixed64", nil, "")
	cmd.PersistentFlags().BoolSliceVar(&req.ListBool, "listbool", nil, "")
	cmd.PersistentFlags().StringSliceVar(&req.ListString, "liststring", nil, "")
	flag.BytesBase64SliceVar(cmd.PersistentFlags(), &req.ListBytes, "listbytes", "")
	_Sound_NestedEnumSliceVar(cmd.PersistentFlags(), &req.ListNestedEnum, "listnestedenum", "")
	_GlobalEnumSliceVar(cmd.PersistentFlags(), &req.ListGlobalEnum, "listglobalenum", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseFloat64, "string=float64", &req.MapStringDouble, "mapstringdouble", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseFloat32, "string=float32", &req.MapStringFloat, "mapstringfloat", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseInt32, "string=int32", &req.MapStringInt32, "mapstringint32", "")
	cmd.PersistentFlags().StringToInt64Var(&req.MapStringInt64, "mapstringint64", nil, "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseUint32, "string=uint32", &req.MapStringUint32, "mapstringuint32", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseUint64, "string=uint64", &req.MapStringUint64, "mapstringuint64", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseInt32, "string=int32", &req.MapStringSint32, "mapstringsint32", "")
	cmd.PersistentFlags().StringToInt64Var(&req.MapStringSint64, "mapstringsint64", nil, "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseUint32, "string=uint32", &req.MapStringFixed32, "mapstringfixed32", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseUint64, "string=uint64", &req.MapStringFixed64, "mapstringfixed64", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseInt32, "string=int32", &req.MapStringSfixed32, "mapstringsfixed32", "")
	cmd.PersistentFlags().StringToInt64Var(&req.MapStringSfixed64, "mapstringsfixed64", nil, "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseBool, "string=bool", &req.MapStringBool, "mapstringbool", "")
	cmd.PersistentFlags().StringToStringVar(&req.MapStringString, "mapstringstring", nil, "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseBytesBase64, "string=bytesBase64", &req.MapStringBytes, "mapstringbytes", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, _Sound_NestedEnumParse, "string=Sound_NestedEnum", &req.MapStringNestedEnum, "mapstringnestedenum", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, _GlobalEnumParse, "string=GlobalEnum", &req.MapStringGlobalEnum, "mapstringglobalenum", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseInt32, flag.ParseString, "int32=string", &req.MapInt32String, "mapint32string", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseInt64, flag.ParseString, "int64=string", &req.MapInt64String, "mapint64string", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseUint32, flag.ParseString, "uint32=string", &req.MapUint32String, "mapuint32string", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseUint64, flag.ParseString, "uint64=string", &req.MapUint64String, "mapuint64string", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseInt32, flag.ParseString, "int32=string", &req.MapSint32String, "mapsint32string", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseInt64, flag.ParseString, "int64=string", &req.MapSint64String, "mapsint64string", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseUint32, flag.ParseString, "uint32=string", &req.MapFixed32String, "mapfixed32string", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseUint64, flag.ParseString, "uint64=string", &req.MapFixed64String, "mapfixed64string", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseInt32, flag.ParseString, "int32=string", &req.MapSfixed32String, "mapsfixed32string", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseInt64, flag.ParseString, "int64=string", &req.MapSfixed64String, "mapsfixed64string", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseBool, flag.ParseString, "bool=string", &req.MapBoolString, "mapboolstring", "")
	flag.TimestampVar(cmd.PersistentFlags(), &req.Timestamp, "timestamp", "")
	flag.DurationVar(cmd.PersistentFlags(), &req.Duration, "duration", "")
	flag.BoolWrapperVar(cmd.PersistentFlags(), &req.WrapperBool, "wrapperbool", "")
	flag.BytesBase64WrapperVar(cmd.PersistentFlags(), &req.WrapperBytes, "wrapperbytes", "")
	flag.DoubleWrapperVar(cmd.PersistentFlags(), &req.WrapperDouble, "wrapperdouble", "")
	flag.FloatWrapperVar(cmd.PersistentFlags(), &req.WrapperFloat, "wrapperfloat", "")
	flag.Int32WrapperVar(cmd.PersistentFlags(), &req.WrapperInt32, "wrapperint32", "")
	flag.Int64WrapperVar(cmd.PersistentFlags(), &req.WrapperInt64, "wrapperint64", "")
	flag.StringWrapperVar(cmd.PersistentFlags(), &req.WrapperString, "wrapperstring", "")
	flag.UInt32WrapperVar(cmd.PersistentFlags(), &req.WrapperUint32, "wrapperuint32", "")
	flag.UInt64WrapperVar(cmd.PersistentFlags(), &req.WrapperUint64, "wrapperuint64", "")
	flag.TimestampSliceVar(cmd.PersistentFlags(), &req.ListTimestamp, "listtimestamp", "")
	flag.DurationSliceVar(cmd.PersistentFlags(), &req.ListDuration, "listduration", "")
	flag.BoolWrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperBool, "listwrapperbool", "")
	flag.BytesBase64WrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperBytes, "listwrapperbytes", "")
	flag.DoubleWrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperDouble, "listwrapperdouble", "")
	flag.FloatWrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperFloat, "listwrapperfloat", "")
	flag.Int32WrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperInt32, "listwrapperint32", "")
	flag.Int64WrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperInt64, "listwrapperint64", "")
	flag.StringWrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperString, "listwrapperstring", "")
	flag.UInt32WrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperUint32, "listwrapperuint32", "")
	flag.UInt64WrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperUint64, "listwrapperuint64", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseTimestamp, "string=timestamp", &req.MapStringTimestamp, "mapstringtimestamp", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseDuration, "string=duration", &req.MapStringDuration, "mapstringduration", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseBoolWrapper, "string=bool", &req.MapStringWrapperBool, "mapstringwrapperbool", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseBytesBase64Wrapper, "string=bytesBase64", &req.MapStringWrapperBytes, "mapstringwrapperbytes", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseDoubleWrapper, "string=float64", &req.MapStringWrapperDouble, "mapstringwrapperdouble", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseFloatWrapper, "string=float32", &req.MapStringWrapperFloat, "mapstringwrapperfloat", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseInt32Wrapper, "string=int32", &req.MapStringWrapperInt32, "mapstringwrapperint32", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseInt64Wrapper, "string=int64", &req.MapStringWrapperInt64, "mapstringwrapperint64", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseStringWrapper, "string=string", &req.MapStringWrapperString, "mapstringwrapperstring", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseUInt32Wrapper, "string=uint32", &req.MapStringWrapperUint32, "mapstringwrapperuint32", "")
	flag.ReflectMapVar(cmd.PersistentFlags(), flag.ParseString, flag.ParseUInt64Wrapper, "string=uint64", &req.MapStringWrapperUint64, "mapstringwrapperuint64", "")

	return cmd
}

type _GlobalEnumValue GlobalEnum

func _GlobalEnumVar(fs *pflag.FlagSet, p *GlobalEnum, name, usage string) {
	fs.Var((*_GlobalEnumValue)(p), name, usage)
}

func (v *_GlobalEnumValue) Set(val string) error {
	if e, err := parseGlobalEnum(val); err != nil {
		return err
	} else {
		*v = _GlobalEnumValue(e)
		return nil
	}
}

func (v *_GlobalEnumValue) Type() string { return "GlobalEnum" }

func (v *_GlobalEnumValue) String() string { return (GlobalEnum)(*v).String() }

type _GlobalEnumSliceValue struct {
	value   *[]GlobalEnum
	changed bool
}

func _GlobalEnumSliceVar(fs *pflag.FlagSet, p *[]GlobalEnum, name, usage string) {
	fs.Var(&_GlobalEnumSliceValue{value: p}, name, usage)
}

func (s *_GlobalEnumSliceValue) Set(val string) error {
	ss := strings.Split(val, ",")
	out := make([]GlobalEnum, len(ss))
	for i, s := range ss {
		var err error
		if out[i], err = parseGlobalEnum(s); err != nil {
			return err
		}
	}
	if !s.changed {
		*s.value = out
		s.changed = true
	} else {
		*s.value = append(*s.value, out...)
	}
	return nil
}

func (s *_GlobalEnumSliceValue) Type() string { return "GlobalEnumSlice" }

func (s *_GlobalEnumSliceValue) String() string { return "[]" }

func _GlobalEnumParse(val string) (interface{}, error) {
	return parseGlobalEnum(val)
}

func parseGlobalEnum(s string) (GlobalEnum, error) {
	if i, ok := GlobalEnum_value[s]; ok {
		return GlobalEnum(i), nil
	} else if i, err := strconv.ParseInt(s, 0, 32); err == nil {
		return GlobalEnum(i), nil
	} else {
		return 0, err
	}
}

type _Sound_NestedEnumValue Sound_NestedEnum

func _Sound_NestedEnumVar(fs *pflag.FlagSet, p *Sound_NestedEnum, name, usage string) {
	fs.Var((*_Sound_NestedEnumValue)(p), name, usage)
}

func (v *_Sound_NestedEnumValue) Set(val string) error {
	if e, err := parseSound_NestedEnum(val); err != nil {
		return err
	} else {
		*v = _Sound_NestedEnumValue(e)
		return nil
	}
}

func (v *_Sound_NestedEnumValue) Type() string { return "Sound_NestedEnum" }

func (v *_Sound_NestedEnumValue) String() string { return (Sound_NestedEnum)(*v).String() }

type _Sound_NestedEnumSliceValue struct {
	value   *[]Sound_NestedEnum
	changed bool
}

func _Sound_NestedEnumSliceVar(fs *pflag.FlagSet, p *[]Sound_NestedEnum, name, usage string) {
	fs.Var(&_Sound_NestedEnumSliceValue{value: p}, name, usage)
}

func (s *_Sound_NestedEnumSliceValue) Set(val string) error {
	ss := strings.Split(val, ",")
	out := make([]Sound_NestedEnum, len(ss))
	for i, s := range ss {
		var err error
		if out[i], err = parseSound_NestedEnum(s); err != nil {
			return err
		}
	}
	if !s.changed {
		*s.value = out
		s.changed = true
	} else {
		*s.value = append(*s.value, out...)
	}
	return nil
}

func (s *_Sound_NestedEnumSliceValue) Type() string { return "Sound_NestedEnumSlice" }

func (s *_Sound_NestedEnumSliceValue) String() string { return "[]" }

func _Sound_NestedEnumParse(val string) (interface{}, error) {
	return parseSound_NestedEnum(val)
}

func parseSound_NestedEnum(s string) (Sound_NestedEnum, error) {
	if i, ok := Sound_NestedEnum_value[s]; ok {
		return Sound_NestedEnum(i), nil
	} else if i, err := strconv.ParseInt(s, 0, 32); err == nil {
		return Sound_NestedEnum(i), nil
	} else {
		return 0, err
	}
}
