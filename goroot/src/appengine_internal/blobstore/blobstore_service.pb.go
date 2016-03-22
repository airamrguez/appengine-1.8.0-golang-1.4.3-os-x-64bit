// Code generated by protoc-gen-go.
// DO NOT EDIT!

/*
Package appengine is a generated protocol buffer package.

It is generated from these files:
	appengine_internal/blobstore

It has these top-level messages:
	BlobstoreServiceError
	CreateUploadURLRequest
	CreateUploadURLResponse
	DeleteBlobRequest
	FetchDataRequest
	FetchDataResponse
	CloneBlobRequest
	CloneBlobResponse
	DecodeBlobKeyRequest
	DecodeBlobKeyResponse
	CreateEncodedGoogleStorageKeyRequest
	CreateEncodedGoogleStorageKeyResponse
*/
package appengine

import proto "appengine_internal/github.com/golang/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type BlobstoreServiceError_ErrorCode int32

const (
	BlobstoreServiceError_OK                        BlobstoreServiceError_ErrorCode = 0
	BlobstoreServiceError_INTERNAL_ERROR            BlobstoreServiceError_ErrorCode = 1
	BlobstoreServiceError_URL_TOO_LONG              BlobstoreServiceError_ErrorCode = 2
	BlobstoreServiceError_PERMISSION_DENIED         BlobstoreServiceError_ErrorCode = 3
	BlobstoreServiceError_BLOB_NOT_FOUND            BlobstoreServiceError_ErrorCode = 4
	BlobstoreServiceError_DATA_INDEX_OUT_OF_RANGE   BlobstoreServiceError_ErrorCode = 5
	BlobstoreServiceError_BLOB_FETCH_SIZE_TOO_LARGE BlobstoreServiceError_ErrorCode = 6
	BlobstoreServiceError_ARGUMENT_OUT_OF_RANGE     BlobstoreServiceError_ErrorCode = 8
	BlobstoreServiceError_INVALID_BLOB_KEY          BlobstoreServiceError_ErrorCode = 9
)

var BlobstoreServiceError_ErrorCode_name = map[int32]string{
	0: "OK",
	1: "INTERNAL_ERROR",
	2: "URL_TOO_LONG",
	3: "PERMISSION_DENIED",
	4: "BLOB_NOT_FOUND",
	5: "DATA_INDEX_OUT_OF_RANGE",
	6: "BLOB_FETCH_SIZE_TOO_LARGE",
	8: "ARGUMENT_OUT_OF_RANGE",
	9: "INVALID_BLOB_KEY",
}
var BlobstoreServiceError_ErrorCode_value = map[string]int32{
	"OK":                        0,
	"INTERNAL_ERROR":            1,
	"URL_TOO_LONG":              2,
	"PERMISSION_DENIED":         3,
	"BLOB_NOT_FOUND":            4,
	"DATA_INDEX_OUT_OF_RANGE":   5,
	"BLOB_FETCH_SIZE_TOO_LARGE": 6,
	"ARGUMENT_OUT_OF_RANGE":     8,
	"INVALID_BLOB_KEY":          9,
}

func (x BlobstoreServiceError_ErrorCode) Enum() *BlobstoreServiceError_ErrorCode {
	p := new(BlobstoreServiceError_ErrorCode)
	*p = x
	return p
}
func (x BlobstoreServiceError_ErrorCode) String() string {
	return proto.EnumName(BlobstoreServiceError_ErrorCode_name, int32(x))
}
func (x *BlobstoreServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BlobstoreServiceError_ErrorCode_value, data, "BlobstoreServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = BlobstoreServiceError_ErrorCode(value)
	return nil
}

type BlobstoreServiceError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *BlobstoreServiceError) Reset()         { *m = BlobstoreServiceError{} }
func (m *BlobstoreServiceError) String() string { return proto.CompactTextString(m) }
func (*BlobstoreServiceError) ProtoMessage()    {}

type CreateUploadURLRequest struct {
	SuccessPath               *string `protobuf:"bytes,1,req,name=success_path" json:"success_path,omitempty"`
	MaxUploadSizeBytes        *int64  `protobuf:"varint,2,opt,name=max_upload_size_bytes" json:"max_upload_size_bytes,omitempty"`
	MaxUploadSizePerBlobBytes *int64  `protobuf:"varint,3,opt,name=max_upload_size_per_blob_bytes" json:"max_upload_size_per_blob_bytes,omitempty"`
	GsBucketName              *string `protobuf:"bytes,4,opt,name=gs_bucket_name" json:"gs_bucket_name,omitempty"`
	XXX_unrecognized          []byte  `json:"-"`
}

func (m *CreateUploadURLRequest) Reset()         { *m = CreateUploadURLRequest{} }
func (m *CreateUploadURLRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUploadURLRequest) ProtoMessage()    {}

func (m *CreateUploadURLRequest) GetSuccessPath() string {
	if m != nil && m.SuccessPath != nil {
		return *m.SuccessPath
	}
	return ""
}

func (m *CreateUploadURLRequest) GetMaxUploadSizeBytes() int64 {
	if m != nil && m.MaxUploadSizeBytes != nil {
		return *m.MaxUploadSizeBytes
	}
	return 0
}

func (m *CreateUploadURLRequest) GetMaxUploadSizePerBlobBytes() int64 {
	if m != nil && m.MaxUploadSizePerBlobBytes != nil {
		return *m.MaxUploadSizePerBlobBytes
	}
	return 0
}

func (m *CreateUploadURLRequest) GetGsBucketName() string {
	if m != nil && m.GsBucketName != nil {
		return *m.GsBucketName
	}
	return ""
}

type CreateUploadURLResponse struct {
	Url              *string `protobuf:"bytes,1,req,name=url" json:"url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateUploadURLResponse) Reset()         { *m = CreateUploadURLResponse{} }
func (m *CreateUploadURLResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUploadURLResponse) ProtoMessage()    {}

func (m *CreateUploadURLResponse) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

type DeleteBlobRequest struct {
	BlobKey          []string `protobuf:"bytes,1,rep,name=blob_key" json:"blob_key,omitempty"`
	Token            *string  `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DeleteBlobRequest) Reset()         { *m = DeleteBlobRequest{} }
func (m *DeleteBlobRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBlobRequest) ProtoMessage()    {}

func (m *DeleteBlobRequest) GetBlobKey() []string {
	if m != nil {
		return m.BlobKey
	}
	return nil
}

func (m *DeleteBlobRequest) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

type FetchDataRequest struct {
	BlobKey          *string `protobuf:"bytes,1,req,name=blob_key" json:"blob_key,omitempty"`
	StartIndex       *int64  `protobuf:"varint,2,req,name=start_index" json:"start_index,omitempty"`
	EndIndex         *int64  `protobuf:"varint,3,req,name=end_index" json:"end_index,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FetchDataRequest) Reset()         { *m = FetchDataRequest{} }
func (m *FetchDataRequest) String() string { return proto.CompactTextString(m) }
func (*FetchDataRequest) ProtoMessage()    {}

func (m *FetchDataRequest) GetBlobKey() string {
	if m != nil && m.BlobKey != nil {
		return *m.BlobKey
	}
	return ""
}

func (m *FetchDataRequest) GetStartIndex() int64 {
	if m != nil && m.StartIndex != nil {
		return *m.StartIndex
	}
	return 0
}

func (m *FetchDataRequest) GetEndIndex() int64 {
	if m != nil && m.EndIndex != nil {
		return *m.EndIndex
	}
	return 0
}

type FetchDataResponse struct {
	Data             []byte `protobuf:"bytes,1000,req,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FetchDataResponse) Reset()         { *m = FetchDataResponse{} }
func (m *FetchDataResponse) String() string { return proto.CompactTextString(m) }
func (*FetchDataResponse) ProtoMessage()    {}

func (m *FetchDataResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type CloneBlobRequest struct {
	BlobKey          []byte `protobuf:"bytes,1,req,name=blob_key" json:"blob_key,omitempty"`
	MimeType         []byte `protobuf:"bytes,2,req,name=mime_type" json:"mime_type,omitempty"`
	TargetAppId      []byte `protobuf:"bytes,3,req,name=target_app_id" json:"target_app_id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CloneBlobRequest) Reset()         { *m = CloneBlobRequest{} }
func (m *CloneBlobRequest) String() string { return proto.CompactTextString(m) }
func (*CloneBlobRequest) ProtoMessage()    {}

func (m *CloneBlobRequest) GetBlobKey() []byte {
	if m != nil {
		return m.BlobKey
	}
	return nil
}

func (m *CloneBlobRequest) GetMimeType() []byte {
	if m != nil {
		return m.MimeType
	}
	return nil
}

func (m *CloneBlobRequest) GetTargetAppId() []byte {
	if m != nil {
		return m.TargetAppId
	}
	return nil
}

type CloneBlobResponse struct {
	BlobKey          []byte `protobuf:"bytes,1,req,name=blob_key" json:"blob_key,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CloneBlobResponse) Reset()         { *m = CloneBlobResponse{} }
func (m *CloneBlobResponse) String() string { return proto.CompactTextString(m) }
func (*CloneBlobResponse) ProtoMessage()    {}

func (m *CloneBlobResponse) GetBlobKey() []byte {
	if m != nil {
		return m.BlobKey
	}
	return nil
}

type DecodeBlobKeyRequest struct {
	BlobKey          []string `protobuf:"bytes,1,rep,name=blob_key" json:"blob_key,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DecodeBlobKeyRequest) Reset()         { *m = DecodeBlobKeyRequest{} }
func (m *DecodeBlobKeyRequest) String() string { return proto.CompactTextString(m) }
func (*DecodeBlobKeyRequest) ProtoMessage()    {}

func (m *DecodeBlobKeyRequest) GetBlobKey() []string {
	if m != nil {
		return m.BlobKey
	}
	return nil
}

type DecodeBlobKeyResponse struct {
	Decoded          []string `protobuf:"bytes,1,rep,name=decoded" json:"decoded,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DecodeBlobKeyResponse) Reset()         { *m = DecodeBlobKeyResponse{} }
func (m *DecodeBlobKeyResponse) String() string { return proto.CompactTextString(m) }
func (*DecodeBlobKeyResponse) ProtoMessage()    {}

func (m *DecodeBlobKeyResponse) GetDecoded() []string {
	if m != nil {
		return m.Decoded
	}
	return nil
}

type CreateEncodedGoogleStorageKeyRequest struct {
	Filename         *string `protobuf:"bytes,1,req,name=filename" json:"filename,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateEncodedGoogleStorageKeyRequest) Reset()         { *m = CreateEncodedGoogleStorageKeyRequest{} }
func (m *CreateEncodedGoogleStorageKeyRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEncodedGoogleStorageKeyRequest) ProtoMessage()    {}

func (m *CreateEncodedGoogleStorageKeyRequest) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

type CreateEncodedGoogleStorageKeyResponse struct {
	BlobKey          *string `protobuf:"bytes,1,req,name=blob_key" json:"blob_key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateEncodedGoogleStorageKeyResponse) Reset()         { *m = CreateEncodedGoogleStorageKeyResponse{} }
func (m *CreateEncodedGoogleStorageKeyResponse) String() string { return proto.CompactTextString(m) }
func (*CreateEncodedGoogleStorageKeyResponse) ProtoMessage()    {}

func (m *CreateEncodedGoogleStorageKeyResponse) GetBlobKey() string {
	if m != nil && m.BlobKey != nil {
		return *m.BlobKey
	}
	return ""
}

func init() {
	proto.RegisterEnum("appengine.BlobstoreServiceError_ErrorCode", BlobstoreServiceError_ErrorCode_name, BlobstoreServiceError_ErrorCode_value)
}
