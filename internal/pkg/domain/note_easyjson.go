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

func easyjson5e0d3bb8DecodeHeshInternalPkgDomain(in *jlexer.Lexer, out *NoteInListInfo) {
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
		case "basicnoteinfo":
			(out.BasicNoteInfo).UnmarshalEasyJSON(in)
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
func easyjson5e0d3bb8EncodeHeshInternalPkgDomain(out *jwriter.Writer, in NoteInListInfo) {
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
		const prefix string = ",\"basicnoteinfo\":"
		out.RawString(prefix)
		(in.BasicNoteInfo).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NoteInListInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5e0d3bb8EncodeHeshInternalPkgDomain(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NoteInListInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5e0d3bb8EncodeHeshInternalPkgDomain(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NoteInListInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5e0d3bb8DecodeHeshInternalPkgDomain(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NoteInListInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5e0d3bb8DecodeHeshInternalPkgDomain(l, v)
}
func easyjson5e0d3bb8DecodeHeshInternalPkgDomain1(in *jlexer.Lexer, out *GetNoteResponse) {
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
		case "ismedicrecord":
			out.IsMedicRecord = bool(in.Bool())
		case "recordid":
			out.RecordId = uint64(in.Uint64())
		case "notelist":
			if in.IsNull() {
				in.Skip()
				out.NoteList = nil
			} else {
				in.Delim('[')
				if out.NoteList == nil {
					if !in.IsDelim(']') {
						out.NoteList = make([]NoteInListInfo, 0, 1)
					} else {
						out.NoteList = []NoteInListInfo{}
					}
				} else {
					out.NoteList = (out.NoteList)[:0]
				}
				for !in.IsDelim(']') {
					var v1 NoteInListInfo
					(v1).UnmarshalEasyJSON(in)
					out.NoteList = append(out.NoteList, v1)
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
func easyjson5e0d3bb8EncodeHeshInternalPkgDomain1(out *jwriter.Writer, in GetNoteResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ismedicrecord\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.IsMedicRecord))
	}
	{
		const prefix string = ",\"recordid\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.RecordId))
	}
	{
		const prefix string = ",\"notelist\":"
		out.RawString(prefix)
		if in.NoteList == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.NoteList {
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
func (v GetNoteResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5e0d3bb8EncodeHeshInternalPkgDomain1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetNoteResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5e0d3bb8EncodeHeshInternalPkgDomain1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetNoteResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5e0d3bb8DecodeHeshInternalPkgDomain1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetNoteResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5e0d3bb8DecodeHeshInternalPkgDomain1(l, v)
}
func easyjson5e0d3bb8DecodeHeshInternalPkgDomain2(in *jlexer.Lexer, out *BasicNoteInfo) {
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
		case "text":
			out.Text = string(in.String())
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
func easyjson5e0d3bb8EncodeHeshInternalPkgDomain2(out *jwriter.Writer, in BasicNoteInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix[1:])
		out.String(string(in.Text))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BasicNoteInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5e0d3bb8EncodeHeshInternalPkgDomain2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BasicNoteInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5e0d3bb8EncodeHeshInternalPkgDomain2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BasicNoteInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5e0d3bb8DecodeHeshInternalPkgDomain2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BasicNoteInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5e0d3bb8DecodeHeshInternalPkgDomain2(l, v)
}
