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

func easyjson15d5d517DecodeHeshInternalPkgDomain(in *jlexer.Lexer, out *RecordUpdateImageResponse) {
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
		case "diaryId":
			out.DiaryId = uint64(in.Uint64())
		case "creatingdate":
			out.CreatingDate = string(in.String())
		case "images":
			if in.IsNull() {
				in.Skip()
				out.Images = nil
			} else {
				in.Delim('[')
				if out.Images == nil {
					if !in.IsDelim(']') {
						out.Images = make([]RecordImageInfo, 0, 1)
					} else {
						out.Images = []RecordImageInfo{}
					}
				} else {
					out.Images = (out.Images)[:0]
				}
				for !in.IsDelim(']') {
					var v1 RecordImageInfo
					(v1).UnmarshalEasyJSON(in)
					out.Images = append(out.Images, v1)
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
func easyjson15d5d517EncodeHeshInternalPkgDomain(out *jwriter.Writer, in RecordUpdateImageResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"diaryId\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.DiaryId))
	}
	{
		const prefix string = ",\"creatingdate\":"
		out.RawString(prefix)
		out.String(string(in.CreatingDate))
	}
	{
		const prefix string = ",\"images\":"
		out.RawString(prefix)
		if in.Images == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Images {
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
func (v RecordUpdateImageResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson15d5d517EncodeHeshInternalPkgDomain(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RecordUpdateImageResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson15d5d517EncodeHeshInternalPkgDomain(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RecordUpdateImageResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson15d5d517DecodeHeshInternalPkgDomain(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RecordUpdateImageResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson15d5d517DecodeHeshInternalPkgDomain(l, v)
}
func easyjson15d5d517DecodeHeshInternalPkgDomain1(in *jlexer.Lexer, out *RecordImageInfo) {
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
		case "imagename":
			out.ImageName = string(in.String())
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]string, 0, 4)
					} else {
						out.Tags = []string{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Tags = append(out.Tags, v4)
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
func easyjson15d5d517EncodeHeshInternalPkgDomain1(out *jwriter.Writer, in RecordImageInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"imagename\":"
		out.RawString(prefix[1:])
		out.String(string(in.ImageName))
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Tags {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RecordImageInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson15d5d517EncodeHeshInternalPkgDomain1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RecordImageInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson15d5d517EncodeHeshInternalPkgDomain1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RecordImageInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson15d5d517DecodeHeshInternalPkgDomain1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RecordImageInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson15d5d517DecodeHeshInternalPkgDomain1(l, v)
}
func easyjson15d5d517DecodeHeshInternalPkgDomain2(in *jlexer.Lexer, out *MedicRecordUpdateTextResponse) {
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
		case "diaryid":
			out.DiaryId = uint64(in.Uint64())
		case "id":
			out.Id = uint64(in.Uint64())
		case "creatingdate":
			out.CreatingDate = string(in.String())
		case "basicinfo":
			(out.BasicInfo).UnmarshalEasyJSON(in)
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
func easyjson15d5d517EncodeHeshInternalPkgDomain2(out *jwriter.Writer, in MedicRecordUpdateTextResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"diaryid\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.DiaryId))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"creatingdate\":"
		out.RawString(prefix)
		out.String(string(in.CreatingDate))
	}
	{
		const prefix string = ",\"basicinfo\":"
		out.RawString(prefix)
		(in.BasicInfo).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MedicRecordUpdateTextResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson15d5d517EncodeHeshInternalPkgDomain2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MedicRecordUpdateTextResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson15d5d517EncodeHeshInternalPkgDomain2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MedicRecordUpdateTextResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson15d5d517DecodeHeshInternalPkgDomain2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MedicRecordUpdateTextResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson15d5d517DecodeHeshInternalPkgDomain2(l, v)
}
func easyjson15d5d517DecodeHeshInternalPkgDomain3(in *jlexer.Lexer, out *MedicRecordUpdateImageRequest) {
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
		case "images":
			if in.IsNull() {
				in.Skip()
				out.Images = nil
			} else {
				in.Delim('[')
				if out.Images == nil {
					if !in.IsDelim(']') {
						out.Images = make([]RecordImageInfo, 0, 1)
					} else {
						out.Images = []RecordImageInfo{}
					}
				} else {
					out.Images = (out.Images)[:0]
				}
				for !in.IsDelim(']') {
					var v7 RecordImageInfo
					(v7).UnmarshalEasyJSON(in)
					out.Images = append(out.Images, v7)
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
func easyjson15d5d517EncodeHeshInternalPkgDomain3(out *jwriter.Writer, in MedicRecordUpdateImageRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"images\":"
		out.RawString(prefix[1:])
		if in.Images == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Images {
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
func (v MedicRecordUpdateImageRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson15d5d517EncodeHeshInternalPkgDomain3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MedicRecordUpdateImageRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson15d5d517EncodeHeshInternalPkgDomain3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MedicRecordUpdateImageRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson15d5d517DecodeHeshInternalPkgDomain3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MedicRecordUpdateImageRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson15d5d517DecodeHeshInternalPkgDomain3(l, v)
}
func easyjson15d5d517DecodeHeshInternalPkgDomain4(in *jlexer.Lexer, out *MedicRecordCreateResponse) {
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
		case "diaryid":
			out.DiaryId = uint64(in.Uint64())
		case "id":
			out.Id = uint64(in.Uint64())
		case "creatingdate":
			out.CreatingDate = string(in.String())
		case "basicinfo":
			(out.BasicInfo).UnmarshalEasyJSON(in)
		case "imagelist":
			if in.IsNull() {
				in.Skip()
				out.ImageList = nil
			} else {
				in.Delim('[')
				if out.ImageList == nil {
					if !in.IsDelim(']') {
						out.ImageList = make([]RecordImageInfo, 0, 1)
					} else {
						out.ImageList = []RecordImageInfo{}
					}
				} else {
					out.ImageList = (out.ImageList)[:0]
				}
				for !in.IsDelim(']') {
					var v10 RecordImageInfo
					(v10).UnmarshalEasyJSON(in)
					out.ImageList = append(out.ImageList, v10)
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
func easyjson15d5d517EncodeHeshInternalPkgDomain4(out *jwriter.Writer, in MedicRecordCreateResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"diaryid\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.DiaryId))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"creatingdate\":"
		out.RawString(prefix)
		out.String(string(in.CreatingDate))
	}
	{
		const prefix string = ",\"basicinfo\":"
		out.RawString(prefix)
		(in.BasicInfo).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"imagelist\":"
		out.RawString(prefix)
		if in.ImageList == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.ImageList {
				if v11 > 0 {
					out.RawByte(',')
				}
				(v12).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MedicRecordCreateResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson15d5d517EncodeHeshInternalPkgDomain4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MedicRecordCreateResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson15d5d517EncodeHeshInternalPkgDomain4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MedicRecordCreateResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson15d5d517DecodeHeshInternalPkgDomain4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MedicRecordCreateResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson15d5d517DecodeHeshInternalPkgDomain4(l, v)
}
func easyjson15d5d517DecodeHeshInternalPkgDomain5(in *jlexer.Lexer, out *MedicRecordCreateRequest) {
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
		case "basicinfo":
			(out.BasicInfo).UnmarshalEasyJSON(in)
		case "images":
			if in.IsNull() {
				in.Skip()
				out.Images = nil
			} else {
				in.Delim('[')
				if out.Images == nil {
					if !in.IsDelim(']') {
						out.Images = make([]RecordImageInfo, 0, 1)
					} else {
						out.Images = []RecordImageInfo{}
					}
				} else {
					out.Images = (out.Images)[:0]
				}
				for !in.IsDelim(']') {
					var v13 RecordImageInfo
					(v13).UnmarshalEasyJSON(in)
					out.Images = append(out.Images, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "audio":
			if in.IsNull() {
				in.Skip()
				out.Auido = nil
			} else {
				in.Delim('[')
				if out.Auido == nil {
					if !in.IsDelim(']') {
						out.Auido = make([]string, 0, 4)
					} else {
						out.Auido = []string{}
					}
				} else {
					out.Auido = (out.Auido)[:0]
				}
				for !in.IsDelim(']') {
					var v14 string
					v14 = string(in.String())
					out.Auido = append(out.Auido, v14)
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
func easyjson15d5d517EncodeHeshInternalPkgDomain5(out *jwriter.Writer, in MedicRecordCreateRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"basicinfo\":"
		out.RawString(prefix[1:])
		(in.BasicInfo).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"images\":"
		out.RawString(prefix)
		if in.Images == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v15, v16 := range in.Images {
				if v15 > 0 {
					out.RawByte(',')
				}
				(v16).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"audio\":"
		out.RawString(prefix)
		if in.Auido == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v17, v18 := range in.Auido {
				if v17 > 0 {
					out.RawByte(',')
				}
				out.String(string(v18))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MedicRecordCreateRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson15d5d517EncodeHeshInternalPkgDomain5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MedicRecordCreateRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson15d5d517EncodeHeshInternalPkgDomain5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MedicRecordCreateRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson15d5d517DecodeHeshInternalPkgDomain5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MedicRecordCreateRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson15d5d517DecodeHeshInternalPkgDomain5(l, v)
}
func easyjson15d5d517DecodeHeshInternalPkgDomain6(in *jlexer.Lexer, out *MedicRecordBasicInfo) {
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
		case "treatment":
			out.Treatment = string(in.String())
		case "recommendations":
			out.Recommendations = string(in.String())
		case "details":
			out.Details = string(in.String())
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
func easyjson15d5d517EncodeHeshInternalPkgDomain6(out *jwriter.Writer, in MedicRecordBasicInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"treatment\":"
		out.RawString(prefix)
		out.String(string(in.Treatment))
	}
	{
		const prefix string = ",\"recommendations\":"
		out.RawString(prefix)
		out.String(string(in.Recommendations))
	}
	{
		const prefix string = ",\"details\":"
		out.RawString(prefix)
		out.String(string(in.Details))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MedicRecordBasicInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson15d5d517EncodeHeshInternalPkgDomain6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MedicRecordBasicInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson15d5d517EncodeHeshInternalPkgDomain6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MedicRecordBasicInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson15d5d517DecodeHeshInternalPkgDomain6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MedicRecordBasicInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson15d5d517DecodeHeshInternalPkgDomain6(l, v)
}
