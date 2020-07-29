package flag

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/spf13/pflag"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func BoolWrapperVar(fs *pflag.FlagSet, p **wrappers.BoolValue, name, usage string) {
	v := fs.Bool(name, false, usage)
	f := fs.Lookup(name)
	f.Value = &wrapperValue{f.Value, func() { *p = wrapperspb.Bool(*v) }, "BoolWrapper"}
}

func BoolWrapperSliceVar(fs *pflag.FlagSet, p *[]*wrappers.BoolValue, name, usage string) {
	v := fs.BoolSlice(name, nil, usage)
	f := fs.Lookup(name)
	var changed bool
	set := func() {
		out := make([]*wrappers.BoolValue, len(*v))
		for i, item := range *v {
			out[i] = wrapperspb.Bool(item)
		}
		if !changed {
			*p = out
			changed = true
		} else {
			*p = append(*p, out...)
		}
	}
	f.Value = &wrapperValue{f.Value, set, "BoolWrapperSlice"}
}

func Int32WrapperVar(fs *pflag.FlagSet, p **wrappers.Int32Value, name, usage string) {
	v := fs.Int32(name, 0, usage)
	f := fs.Lookup(name)
	f.Value = &wrapperValue{f.Value, func() { *p = wrapperspb.Int32(*v) }, "Int32Wrapper"}
}

func Int32WrapperSliceVar(fs *pflag.FlagSet, p *[]*wrappers.Int32Value, name, usage string) {
	v := fs.Int32Slice(name, nil, usage)
	f := fs.Lookup(name)
	var changed bool
	set := func() {
		out := make([]*wrappers.Int32Value, len(*v))
		for i, item := range *v {
			out[i] = wrapperspb.Int32(item)
		}
		if !changed {
			*p = out
			changed = true
		} else {
			*p = append(*p, out...)
		}
	}
	f.Value = &wrapperValue{f.Value, set, "Int32WrapperSlice"}
}

func Int64WrapperVar(fs *pflag.FlagSet, p **wrappers.Int64Value, name, usage string) {
	v := fs.Int64(name, 0, usage)
	f := fs.Lookup(name)
	f.Value = &wrapperValue{f.Value, func() { *p = wrapperspb.Int64(*v) }, "Int64Wrapper"}
}

func Int64WrapperSliceVar(fs *pflag.FlagSet, p *[]*wrappers.Int64Value, name, usage string) {
	v := fs.Int64Slice(name, nil, usage)
	f := fs.Lookup(name)
	var changed bool
	set := func() {
		out := make([]*wrappers.Int64Value, len(*v))
		for i, item := range *v {
			out[i] = wrapperspb.Int64(item)
		}
		if !changed {
			*p = out
			changed = true
		} else {
			*p = append(*p, out...)
		}
	}
	f.Value = &wrapperValue{f.Value, set, "Int64WrapperSlice"}
}

func UInt32WrapperVar(fs *pflag.FlagSet, p **wrappers.UInt32Value, name, usage string) {
	v := fs.Uint32(name, 0, usage)
	f := fs.Lookup(name)
	f.Value = &wrapperValue{f.Value, func() { *p = wrapperspb.UInt32(*v) }, "UInt32Wrapper"}
}

func UInt32WrapperSliceVar(fs *pflag.FlagSet, p *[]*wrappers.UInt32Value, name, usage string) {
	var v []uint32
	Uint32SliceVar(fs, &v, name, usage)
	f := fs.Lookup(name)
	var changed bool
	set := func() {
		out := make([]*wrappers.UInt32Value, len(v))
		for i, item := range v {
			out[i] = wrapperspb.UInt32(item)
		}
		if !changed {
			*p = out
			changed = true
		} else {
			*p = append(*p, out...)
		}
	}
	f.Value = &wrapperValue{f.Value, set, "UInt32WrapperSlice"}
}

func UInt64WrapperVar(fs *pflag.FlagSet, p **wrappers.UInt64Value, name, usage string) {
	v := fs.Uint64(name, 0, usage)
	f := fs.Lookup(name)
	f.Value = &wrapperValue{f.Value, func() { *p = wrapperspb.UInt64(*v) }, "UInt64Wrapper"}
}

func UInt64WrapperSliceVar(fs *pflag.FlagSet, p *[]*wrappers.UInt64Value, name, usage string) {
	var v []uint64
	Uint64SliceVar(fs, &v, name, usage)
	f := fs.Lookup(name)
	var changed bool
	set := func() {
		out := make([]*wrappers.UInt64Value, len(v))
		for i, item := range v {
			out[i] = wrapperspb.UInt64(item)
		}
		if !changed {
			*p = out
			changed = true
		} else {
			*p = append(*p, out...)
		}
	}
	f.Value = &wrapperValue{f.Value, set, "UInt64WrapperSlice"}
}

func FloatWrapperVar(fs *pflag.FlagSet, p **wrappers.FloatValue, name, usage string) {
	v := fs.Float32(name, 0, usage)
	f := fs.Lookup(name)
	f.Value = &wrapperValue{f.Value, func() { *p = wrapperspb.Float(*v) }, "FloatWrapper"}
}

func FloatWrapperSliceVar(fs *pflag.FlagSet, p *[]*wrappers.FloatValue, name, usage string) {
	v := fs.Float32Slice(name, nil, usage)
	f := fs.Lookup(name)
	var changed bool
	set := func() {
		out := make([]*wrappers.FloatValue, len(*v))
		for i, item := range *v {
			out[i] = wrapperspb.Float(item)
		}
		if !changed {
			*p = out
			changed = true
		} else {
			*p = append(*p, out...)
		}
	}
	f.Value = &wrapperValue{f.Value, set, "FloatWrapperSlice"}
}

func DoubleWrapperVar(fs *pflag.FlagSet, p **wrappers.DoubleValue, name, usage string) {
	v := fs.Float64(name, 0, usage)
	f := fs.Lookup(name)
	f.Value = &wrapperValue{f.Value, func() { *p = wrapperspb.Double(*v) }, "DoubleWrapper"}
}

func DoubleWrapperSliceVar(fs *pflag.FlagSet, p *[]*wrappers.DoubleValue, name, usage string) {
	v := fs.Float64Slice(name, nil, usage)
	f := fs.Lookup(name)
	var changed bool
	set := func() {
		out := make([]*wrappers.DoubleValue, len(*v))
		for i, item := range *v {
			out[i] = wrapperspb.Double(item)
		}
		if !changed {
			*p = out
			changed = true
		} else {
			*p = append(*p, out...)
		}
	}
	f.Value = &wrapperValue{f.Value, set, "DoubleWrapperSlice"}
}

func StringWrapperVar(fs *pflag.FlagSet, p **wrappers.StringValue, name, usage string) {
	v := fs.String(name, "", usage)
	f := fs.Lookup(name)
	f.Value = &wrapperValue{f.Value, func() { *p = wrapperspb.String(*v) }, "StringWrapper"}
}

func StringWrapperSliceVar(fs *pflag.FlagSet, p *[]*wrappers.StringValue, name, usage string) {
	v := fs.StringSlice(name, nil, usage)
	f := fs.Lookup(name)
	var changed bool
	set := func() {
		out := make([]*wrappers.StringValue, len(*v))
		for i, item := range *v {
			out[i] = wrapperspb.String(item)
		}
		if !changed {
			*p = out
			changed = true
		} else {
			*p = append(*p, out...)
		}
	}
	f.Value = &wrapperValue{f.Value, set, "StringWrapperSlice"}
}

func BytesBase64WrapperVar(fs *pflag.FlagSet, p **wrappers.BytesValue, name, usage string) {
	v := fs.BytesBase64(name, nil, usage)
	f := fs.Lookup(name)
	f.Value = &wrapperValue{f.Value, func() { *p = wrapperspb.Bytes(*v) }, "BytesBase64Wrapper"}
}

func BytesBase64WrapperSliceVar(fs *pflag.FlagSet, p *[]*wrappers.BytesValue, name, usage string) {
	var v [][]byte
	BytesBase64SliceVar(fs, &v, name, usage)
	f := fs.Lookup(name)
	var changed bool
	set := func() {
		out := make([]*wrappers.BytesValue, len(v))
		for i, item := range v {
			out[i] = wrapperspb.Bytes(item)
		}
		if !changed {
			*p = out
			changed = true
		} else {
			*p = append(*p, out...)
		}
	}
	f.Value = &wrapperValue{f.Value, set, "BytesBase64WrapperSlice"}
}

type wrapperValue struct {
	pflag.Value
	set func()
	typ string
}

func (v *wrapperValue) Set(s string) error {
	if err := v.Value.Set(s); err != nil {
		return err
	}
	v.set()
	return nil
}

func (v *wrapperValue) Type() string { return v.typ }

func (*wrapperValue) String() string { return "<nil>" }