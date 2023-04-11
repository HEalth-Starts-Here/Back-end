// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package domain

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

func easyjsonCb88da03DecodeHeshInternalPkgDomain(in *jlexer.Lexer, out *ImageQualityAssesment) {
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
		case "assesment":
			out.Assesment = bool(in.Bool())
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
func easyjsonCb88da03EncodeHeshInternalPkgDomain(out *jwriter.Writer, in ImageQualityAssesment) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"assesment\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Assesment))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ImageQualityAssesment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCb88da03EncodeHeshInternalPkgDomain(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ImageQualityAssesment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCb88da03EncodeHeshInternalPkgDomain(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ImageQualityAssesment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCb88da03DecodeHeshInternalPkgDomain(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ImageQualityAssesment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCb88da03DecodeHeshInternalPkgDomain(l, v)
}
func easyjsonCb88da03DecodeHeshInternalPkgDomain1(in *jlexer.Lexer, out *DiarisationResponse) {
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
		case "id":
			out.Id = uint64(in.Uint64())
		case "creatingdate":
			out.CreatingDate = string(in.String())
		case "medicrecordid":
			out.MedicRecordId = uint64(in.Uint64())
		case "DiarisationInfo":
			(out.DiarisationInfo).UnmarshalEasyJSON(in)
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
func easyjsonCb88da03EncodeHeshInternalPkgDomain1(out *jwriter.Writer, in DiarisationResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"creatingdate\":"
		out.RawString(prefix)
		out.String(string(in.CreatingDate))
	}
	{
		const prefix string = ",\"medicrecordid\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.MedicRecordId))
	}
	{
		const prefix string = ",\"DiarisationInfo\":"
		out.RawString(prefix)
		(in.DiarisationInfo).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DiarisationResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCb88da03EncodeHeshInternalPkgDomain1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DiarisationResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCb88da03EncodeHeshInternalPkgDomain1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DiarisationResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCb88da03DecodeHeshInternalPkgDomain1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DiarisationResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCb88da03DecodeHeshInternalPkgDomain1(l, v)
}
func easyjsonCb88da03DecodeHeshInternalPkgDomain2(in *jlexer.Lexer, out *DiarisationInfo) {
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
		case "diarisation":
			out.Diarisation = string(in.String())
		case "filename":
			out.Filename = string(in.String())
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
func easyjsonCb88da03EncodeHeshInternalPkgDomain2(out *jwriter.Writer, in DiarisationInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"diarisation\":"
		out.RawString(prefix[1:])
		out.String(string(in.Diarisation))
	}
	{
		const prefix string = ",\"filename\":"
		out.RawString(prefix)
		out.String(string(in.Filename))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DiarisationInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCb88da03EncodeHeshInternalPkgDomain2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DiarisationInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCb88da03EncodeHeshInternalPkgDomain2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DiarisationInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCb88da03DecodeHeshInternalPkgDomain2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DiarisationInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCb88da03DecodeHeshInternalPkgDomain2(l, v)
}
func easyjsonCb88da03DecodeHeshInternalPkgDomain3(in *jlexer.Lexer, out *DetermineAreaResponse) {
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
		case "area":
			out.Area = int32(in.Int32())
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
func easyjsonCb88da03EncodeHeshInternalPkgDomain3(out *jwriter.Writer, in DetermineAreaResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"area\":"
		out.RawString(prefix[1:])
		out.Int32(int32(in.Area))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DetermineAreaResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCb88da03EncodeHeshInternalPkgDomain3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DetermineAreaResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCb88da03EncodeHeshInternalPkgDomain3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DetermineAreaResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCb88da03DecodeHeshInternalPkgDomain3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DetermineAreaResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCb88da03DecodeHeshInternalPkgDomain3(l, v)
}