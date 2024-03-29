// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package case_item

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC803d3e7DecodeBetterGoJsonCaseItem(in *jlexer.Lexer, out *Item) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "impTime":
			out.Time = int(in.Int())
		case "cTime":
			out.CTime = int(in.Int())
		case "iid":
			out.Iid = string(in.String())
		case "fps":
			if in.IsNull() {
				in.Skip()
				out.Fps = nil
			} else {
				in.Delim('[')
				if out.Fps == nil {
					if !in.IsDelim(']') {
						out.Fps = make([]string, 0, 4)
					} else {
						out.Fps = []string{}
					}
				} else {
					out.Fps = (out.Fps)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Fps = append(out.Fps, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "fp_title_hash":
			out.FpTitleHash = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC803d3e7EncodeBetterGoJsonCaseItem(out *jwriter.Writer, in Item) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"impTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Time))
	}
	{
		const prefix string = ",\"cTime\":"
		out.RawString(prefix)
		out.Int(int(in.CTime))
	}
	{
		const prefix string = ",\"iid\":"
		out.RawString(prefix)
		out.String(string(in.Iid))
	}
	{
		const prefix string = ",\"fps\":"
		out.RawString(prefix)
		if in.Fps == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Fps {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	if in.FpTitleHash != "" {
		const prefix string = ",\"fp_title_hash\":"
		out.RawString(prefix)
		out.String(string(in.FpTitleHash))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Item) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC803d3e7EncodeBetterGoJsonCaseItem(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Item) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC803d3e7EncodeBetterGoJsonCaseItem(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Item) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC803d3e7DecodeBetterGoJsonCaseItem(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Item) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC803d3e7DecodeBetterGoJsonCaseItem(l, v)
}
