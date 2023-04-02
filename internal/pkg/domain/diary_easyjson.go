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

func easyjson1ddc3ff7DecodeHeshInternalPkgDomain(in *jlexer.Lexer, out *RecordUpdateRequest) {
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
		case "RecordCreateRequest":
			(out.RecordCreateRequest).UnmarshalEasyJSON(in)
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
func easyjson1ddc3ff7EncodeHeshInternalPkgDomain(out *jwriter.Writer, in RecordUpdateRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"RecordCreateRequest\":"
		out.RawString(prefix)
		(in.RecordCreateRequest).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RecordUpdateRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RecordUpdateRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RecordUpdateRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RecordUpdateRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain(l, v)
}
func easyjson1ddc3ff7DecodeHeshInternalPkgDomain1(in *jlexer.Lexer, out *RecordCreateResponse) {
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
		case "diaryid":
			out.DiaryId = uint64(in.Uint64())
		case "creatingdate":
			out.CreatingDate = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "area":
			out.Area = float64(in.Float64())
		case "characteristics":
			(out.Characteristics).UnmarshalEasyJSON(in)
		case "imagelist":
			if in.IsNull() {
				in.Skip()
				out.ImageList = nil
			} else {
				in.Delim('[')
				if out.ImageList == nil {
					if !in.IsDelim(']') {
						out.ImageList = make([]ImageInfo, 0, 1)
					} else {
						out.ImageList = []ImageInfo{}
					}
				} else {
					out.ImageList = (out.ImageList)[:0]
				}
				for !in.IsDelim(']') {
					var v1 ImageInfo
					(v1).UnmarshalEasyJSON(in)
					out.ImageList = append(out.ImageList, v1)
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
func easyjson1ddc3ff7EncodeHeshInternalPkgDomain1(out *jwriter.Writer, in RecordCreateResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"diaryid\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.DiaryId))
	}
	{
		const prefix string = ",\"creatingdate\":"
		out.RawString(prefix)
		out.String(string(in.CreatingDate))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"area\":"
		out.RawString(prefix)
		out.Float64(float64(in.Area))
	}
	{
		const prefix string = ",\"characteristics\":"
		out.RawString(prefix)
		(in.Characteristics).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"imagelist\":"
		out.RawString(prefix)
		if in.ImageList == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.ImageList {
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
func (v RecordCreateResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RecordCreateResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RecordCreateResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RecordCreateResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain1(l, v)
}
func easyjson1ddc3ff7DecodeHeshInternalPkgDomain2(in *jlexer.Lexer, out *RecordCreateRequest) {
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
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "characteristics":
			(out.Characteristics).UnmarshalEasyJSON(in)
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
func easyjson1ddc3ff7EncodeHeshInternalPkgDomain2(out *jwriter.Writer, in RecordCreateRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"characteristics\":"
		out.RawString(prefix)
		(in.Characteristics).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RecordCreateRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RecordCreateRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RecordCreateRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RecordCreateRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain2(l, v)
}
func easyjson1ddc3ff7DecodeHeshInternalPkgDomain3(in *jlexer.Lexer, out *ImageInfoUsecase) {
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
		case "name":
			out.Name = string(in.String())
		case "area":
			out.Area = float64(in.Float64())
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
func easyjson1ddc3ff7EncodeHeshInternalPkgDomain3(out *jwriter.Writer, in ImageInfoUsecase) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"area\":"
		out.RawString(prefix)
		out.Float64(float64(in.Area))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ImageInfoUsecase) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ImageInfoUsecase) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ImageInfoUsecase) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ImageInfoUsecase) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain3(l, v)
}
func easyjson1ddc3ff7DecodeHeshInternalPkgDomain4(in *jlexer.Lexer, out *ImageInfo) {
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
		case "recordid":
			out.RecordId = uint64(in.Uint64())
		case "name":
			out.Name = string(in.String())
		case "area":
			out.Area = float64(in.Float64())
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
func easyjson1ddc3ff7EncodeHeshInternalPkgDomain4(out *jwriter.Writer, in ImageInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"recordid\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.RecordId))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"area\":"
		out.RawString(prefix)
		out.Float64(float64(in.Area))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ImageInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ImageInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ImageInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ImageInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain4(l, v)
}
func easyjson1ddc3ff7DecodeHeshInternalPkgDomain5(in *jlexer.Lexer, out *DiaryUpdateResponse) {
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
		case "title":
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
func easyjson1ddc3ff7EncodeHeshInternalPkgDomain5(out *jwriter.Writer, in DiaryUpdateResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"title\":"
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
func (v DiaryUpdateResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DiaryUpdateResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DiaryUpdateResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DiaryUpdateResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain5(l, v)
}
func easyjson1ddc3ff7DecodeHeshInternalPkgDomain6(in *jlexer.Lexer, out *DiaryUpdateRequest) {
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
		case "title":
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
func easyjson1ddc3ff7EncodeHeshInternalPkgDomain6(out *jwriter.Writer, in DiaryUpdateRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"title\":"
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
func (v DiaryUpdateRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DiaryUpdateRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DiaryUpdateRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DiaryUpdateRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain6(l, v)
}
func easyjson1ddc3ff7DecodeHeshInternalPkgDomain7(in *jlexer.Lexer, out *DiaryResponse) {
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
		case "diary":
			(out.Diary).UnmarshalEasyJSON(in)
		case "records":
			if in.IsNull() {
				in.Skip()
				out.RecordsList = nil
			} else {
				in.Delim('[')
				if out.RecordsList == nil {
					if !in.IsDelim(']') {
						out.RecordsList = make([]RecordCreateResponse, 0, 0)
					} else {
						out.RecordsList = []RecordCreateResponse{}
					}
				} else {
					out.RecordsList = (out.RecordsList)[:0]
				}
				for !in.IsDelim(']') {
					var v4 RecordCreateResponse
					(v4).UnmarshalEasyJSON(in)
					out.RecordsList = append(out.RecordsList, v4)
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
func easyjson1ddc3ff7EncodeHeshInternalPkgDomain7(out *jwriter.Writer, in DiaryResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"diary\":"
		out.RawString(prefix[1:])
		(in.Diary).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"records\":"
		out.RawString(prefix)
		if in.RecordsList == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.RecordsList {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DiaryResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DiaryResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DiaryResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DiaryResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain7(l, v)
}
func easyjson1ddc3ff7DecodeHeshInternalPkgDomain8(in *jlexer.Lexer, out *DiaryListResponse) {
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
						out.DiaryList = make([]DiaryInList, 0, 0)
					} else {
						out.DiaryList = []DiaryInList{}
					}
				} else {
					out.DiaryList = (out.DiaryList)[:0]
				}
				for !in.IsDelim(']') {
					var v7 DiaryInList
					(v7).UnmarshalEasyJSON(in)
					out.DiaryList = append(out.DiaryList, v7)
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
func easyjson1ddc3ff7EncodeHeshInternalPkgDomain8(out *jwriter.Writer, in DiaryListResponse) {
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
			for v8, v9 := range in.DiaryList {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DiaryListResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DiaryListResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DiaryListResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DiaryListResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain8(l, v)
}
func easyjson1ddc3ff7DecodeHeshInternalPkgDomain9(in *jlexer.Lexer, out *DiaryInList) {
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
		case "title":
			out.Title = string(in.String())
		case "medicid":
			out.MedicId = uint32(in.Uint32())
		case "medicname":
			out.MedicName = string(in.String())
		case "patientid":
			out.PatientId = uint32(in.Uint32())
		case "patientname":
			out.PatientName = string(in.String())
		case "creatingdate":
			out.CreatingDate = string(in.String())
		case "objectively":
			out.Objectively = string(in.String())
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
func easyjson1ddc3ff7EncodeHeshInternalPkgDomain9(out *jwriter.Writer, in DiaryInList) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"medicid\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.MedicId))
	}
	{
		const prefix string = ",\"medicname\":"
		out.RawString(prefix)
		out.String(string(in.MedicName))
	}
	{
		const prefix string = ",\"patientid\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.PatientId))
	}
	{
		const prefix string = ",\"patientname\":"
		out.RawString(prefix)
		out.String(string(in.PatientName))
	}
	{
		const prefix string = ",\"creatingdate\":"
		out.RawString(prefix)
		out.String(string(in.CreatingDate))
	}
	{
		const prefix string = ",\"objectively\":"
		out.RawString(prefix)
		out.String(string(in.Objectively))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DiaryInList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DiaryInList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DiaryInList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DiaryInList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain9(l, v)
}
func easyjson1ddc3ff7DecodeHeshInternalPkgDomain10(in *jlexer.Lexer, out *DiaryCreateResponse) {
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
		case "medicid":
			out.MedicId = uint32(in.Uint32())
		case "patientid":
			out.PatientId = uint32(in.Uint32())
		case "creatingdate":
			out.CreatingDate = string(in.String())
		case "diarybasicinfo":
			(out.DiaryBasicInfo).UnmarshalEasyJSON(in)
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
func easyjson1ddc3ff7EncodeHeshInternalPkgDomain10(out *jwriter.Writer, in DiaryCreateResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
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
		const prefix string = ",\"diarybasicinfo\":"
		out.RawString(prefix)
		(in.DiaryBasicInfo).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DiaryCreateResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DiaryCreateResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DiaryCreateResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DiaryCreateResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain10(l, v)
}
func easyjson1ddc3ff7DecodeHeshInternalPkgDomain11(in *jlexer.Lexer, out *DiaryCreateRequest) {
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
		case "diarybasicinfo":
			(out.DiaryBasicInfo).UnmarshalEasyJSON(in)
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
func easyjson1ddc3ff7EncodeHeshInternalPkgDomain11(out *jwriter.Writer, in DiaryCreateRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"diarybasicinfo\":"
		out.RawString(prefix[1:])
		(in.DiaryBasicInfo).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DiaryCreateRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DiaryCreateRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DiaryCreateRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DiaryCreateRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain11(l, v)
}
func easyjson1ddc3ff7DecodeHeshInternalPkgDomain12(in *jlexer.Lexer, out *DiaryBasicInfo) {
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
		case "title":
			out.Title = string(in.String())
		case "complaints":
			out.Complaints = string(in.String())
		case "anamnesis":
			out.Anamnesis = string(in.String())
		case "objectively":
			out.Objectively = string(in.String())
		case "diagnosis":
			out.Diagnosis = string(in.String())
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
func easyjson1ddc3ff7EncodeHeshInternalPkgDomain12(out *jwriter.Writer, in DiaryBasicInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"complaints\":"
		out.RawString(prefix)
		out.String(string(in.Complaints))
	}
	{
		const prefix string = ",\"anamnesis\":"
		out.RawString(prefix)
		out.String(string(in.Anamnesis))
	}
	{
		const prefix string = ",\"objectively\":"
		out.RawString(prefix)
		out.String(string(in.Objectively))
	}
	{
		const prefix string = ",\"diagnosis\":"
		out.RawString(prefix)
		out.String(string(in.Diagnosis))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DiaryBasicInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DiaryBasicInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DiaryBasicInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DiaryBasicInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain12(l, v)
}
func easyjson1ddc3ff7DecodeHeshInternalPkgDomain13(in *jlexer.Lexer, out *Characteristics) {
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
		case "dryness":
			out.Dryness = uint8(in.Uint8())
		case "edema":
			out.Edema = uint8(in.Uint8())
		case "itching":
			out.Itching = uint8(in.Uint8())
		case "pain":
			out.Pain = uint8(in.Uint8())
		case "peeling":
			out.Peeling = uint8(in.Uint8())
		case "redness":
			out.Redness = uint8(in.Uint8())
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
func easyjson1ddc3ff7EncodeHeshInternalPkgDomain13(out *jwriter.Writer, in Characteristics) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"dryness\":"
		out.RawString(prefix[1:])
		out.Uint8(uint8(in.Dryness))
	}
	{
		const prefix string = ",\"edema\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.Edema))
	}
	{
		const prefix string = ",\"itching\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.Itching))
	}
	{
		const prefix string = ",\"pain\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.Pain))
	}
	{
		const prefix string = ",\"peeling\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.Peeling))
	}
	{
		const prefix string = ",\"redness\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.Redness))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Characteristics) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain13(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Characteristics) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1ddc3ff7EncodeHeshInternalPkgDomain13(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Characteristics) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain13(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Characteristics) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1ddc3ff7DecodeHeshInternalPkgDomain13(l, v)
}
