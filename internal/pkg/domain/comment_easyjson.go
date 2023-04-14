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

func easyjsonE9abebc9DecodeHeshInternalPkgDomain(in *jlexer.Lexer, out *CommentCreateResponse) {
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
		case "basiccommentinfo":
			(out.BasicCommentInfo).UnmarshalEasyJSON(in)
		case "authorismedic":
			out.AuthorIsMedic = bool(in.Bool())
		case "isreaded":
			out.IsReaded = bool(in.Bool())
		case "creatingdate":
			out.CreatingDate = string(in.String())
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
func easyjsonE9abebc9EncodeHeshInternalPkgDomain(out *jwriter.Writer, in CommentCreateResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"basiccommentinfo\":"
		out.RawString(prefix)
		(in.BasicCommentInfo).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"authorismedic\":"
		out.RawString(prefix)
		out.Bool(bool(in.AuthorIsMedic))
	}
	{
		const prefix string = ",\"isreaded\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsReaded))
	}
	{
		const prefix string = ",\"creatingdate\":"
		out.RawString(prefix)
		out.String(string(in.CreatingDate))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CommentCreateResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE9abebc9EncodeHeshInternalPkgDomain(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CommentCreateResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE9abebc9EncodeHeshInternalPkgDomain(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CommentCreateResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE9abebc9DecodeHeshInternalPkgDomain(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CommentCreateResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE9abebc9DecodeHeshInternalPkgDomain(l, v)
}
func easyjsonE9abebc9DecodeHeshInternalPkgDomain1(in *jlexer.Lexer, out *BasicCommentInfo) {
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
func easyjsonE9abebc9EncodeHeshInternalPkgDomain1(out *jwriter.Writer, in BasicCommentInfo) {
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
func (v BasicCommentInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE9abebc9EncodeHeshInternalPkgDomain1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BasicCommentInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE9abebc9EncodeHeshInternalPkgDomain1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BasicCommentInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE9abebc9DecodeHeshInternalPkgDomain1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BasicCommentInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE9abebc9DecodeHeshInternalPkgDomain1(l, v)
}
