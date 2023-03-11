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

func easyjson1ddc3ff7DecodeHeshInternalPkgDomain(in *jlexer.Lexer, out *DiaryListResponse) {
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
		case "diarylist":
			if in.IsNull() {
				in.Skip()
				out.DiaryList = nil
			} else {
				in.Delim('[')
				if out.DiaryList == nil {
					if !in.IsDelim(']') {
						out.DiaryList = make([]DiaryCreatingResponse, 0, 0)
					} else {
						out.DiaryList = []DiaryCreatingResponse{}
					}
				} else {
					out.DiaryList = (out.DiaryList)[:0]
				}
				for !in.IsDelim(']') {
					var v1 DiaryCreatingResponse
					(v1).UnmarshalEasyJSON(in)
					out.DiaryList = append(out.DiaryList, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson1ddc3ff7EncodeHeshInternalPkgDomain(out *jwriter.Writer, in DiaryListResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"diarylist\":"
		out.RawString(prefix[1:])
		if in.DiaryList == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.DiaryList {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DiaryListResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DiaryListResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DiaryListResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DiaryListResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain(l, v)
}
func easyjson1ddc3ff7DecodeHeshInternalPkgDomain1(in *jlexer.Lexer, out *DiaryCreatingResponse) {
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
		case "category":
			out.Category = uint32(in.Uint32())
		case "medicid":
			out.MedicId = uint32(in.Uint32())
		case "patientid":
			out.PatientId = uint32(in.Uint32())
		case "creatingdate":
			out.CreatingDate = string(in.String())
		case "name":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
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
func easyjson1ddc3ff7EncodeHeshInternalPkgDomain1(out *jwriter.Writer, in DiaryCreatingResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"category\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.Category))
	}
	{
		const prefix string = ",\"medicid\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.MedicId))
	}
	{
		const prefix string = ",\"patientid\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.PatientId))
	}
	{
		const prefix string = ",\"creatingdate\":"
		out.RawString(prefix)
		out.String(string(in.CreatingDate))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DiaryCreatingResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DiaryCreatingResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DiaryCreatingResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DiaryCreatingResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain1(l, v)
}
func easyjson1ddc3ff7DecodeHeshInternalPkgDomain2(in *jlexer.Lexer, out *DiaryCreatingRequest) {
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
		case "category":
			out.Category = uint32(in.Uint32())
		case "medicid":
			out.MedicId = uint32(in.Uint32())
		case "patientid":
			out.PatientId = uint32(in.Uint32())
		case "name":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
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
func easyjson1ddc3ff7EncodeHeshInternalPkgDomain2(out *jwriter.Writer, in DiaryCreatingRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"category\":"
		out.RawString(prefix[1:])
		out.Uint32(uint32(in.Category))
	}
	{
		const prefix string = ",\"medicid\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.MedicId))
	}
	{
		const prefix string = ",\"patientid\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.PatientId))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DiaryCreatingRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DiaryCreatingRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DiaryCreatingRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DiaryCreatingRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain2(l, v)
}
