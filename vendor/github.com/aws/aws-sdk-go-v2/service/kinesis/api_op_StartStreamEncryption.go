// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type StartStreamEncryptionInput struct {
	_ struct{} `type:"structure"`

	// The encryption type to use. The only valid value is KMS.
	//
	// EncryptionType is a required field
	EncryptionType EncryptionType `type:"string" required:"true" enum:"true"`

	// The GUID for the customer-managed AWS KMS key to use for encryption. This
	// value can be a globally unique identifier, a fully specified Amazon Resource
	// Name (ARN) to either an alias or a key, or an alias name prefixed by "alias/".You
	// can also use a master key owned by Kinesis Data Streams by specifying the
	// alias aws/kinesis.
	//
	//    * Key ARN example: arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012
	//
	//    * Alias ARN example: arn:aws:kms:us-east-1:123456789012:alias/MyAliasName
	//
	//    * Globally unique key ID example: 12345678-1234-1234-1234-123456789012
	//
	//    * Alias name example: alias/MyAliasName
	//
	//    * Master key owned by Kinesis Data Streams: alias/aws/kinesis
	//
	// KeyId is a required field
	KeyId *string `min:"1" type:"string" required:"true"`

	// The name of the stream for which to start encrypting records.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartStreamEncryptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartStreamEncryptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartStreamEncryptionInput"}
	if len(s.EncryptionType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("EncryptionType"))
	}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}

	if s.StreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamName"))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartStreamEncryptionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StartStreamEncryptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartStreamEncryption = "StartStreamEncryption"

// StartStreamEncryptionRequest returns a request value for making API operation for
// Amazon Kinesis.
//
// Enables or updates server-side encryption using an AWS KMS key for a specified
// stream.
//
// Starting encryption is an asynchronous operation. Upon receiving the request,
// Kinesis Data Streams returns immediately and sets the status of the stream
// to UPDATING. After the update is complete, Kinesis Data Streams sets the
// status of the stream back to ACTIVE. Updating or applying encryption normally
// takes a few seconds to complete, but it can take minutes. You can continue
// to read and write data to your stream while its status is UPDATING. Once
// the status of the stream is ACTIVE, encryption begins for records written
// to the stream.
//
// API Limits: You can successfully apply a new AWS KMS key for server-side
// encryption 25 times in a rolling 24-hour period.
//
// Note: It can take up to 5 seconds after the stream is in an ACTIVE status
// before all records written to the stream are encrypted. After you enable
// encryption, you can verify that encryption is applied by inspecting the API
// response from PutRecord or PutRecords.
//
//    // Example sending a request using StartStreamEncryptionRequest.
//    req := client.StartStreamEncryptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/StartStreamEncryption
func (c *Client) StartStreamEncryptionRequest(input *StartStreamEncryptionInput) StartStreamEncryptionRequest {
	op := &aws.Operation{
		Name:       opStartStreamEncryption,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartStreamEncryptionInput{}
	}

	req := c.newRequest(op, input, &StartStreamEncryptionOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return StartStreamEncryptionRequest{Request: req, Input: input, Copy: c.StartStreamEncryptionRequest}
}

// StartStreamEncryptionRequest is the request type for the
// StartStreamEncryption API operation.
type StartStreamEncryptionRequest struct {
	*aws.Request
	Input *StartStreamEncryptionInput
	Copy  func(*StartStreamEncryptionInput) StartStreamEncryptionRequest
}

// Send marshals and sends the StartStreamEncryption API request.
func (r StartStreamEncryptionRequest) Send(ctx context.Context) (*StartStreamEncryptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartStreamEncryptionResponse{
		StartStreamEncryptionOutput: r.Request.Data.(*StartStreamEncryptionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartStreamEncryptionResponse is the response type for the
// StartStreamEncryption API operation.
type StartStreamEncryptionResponse struct {
	*StartStreamEncryptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartStreamEncryption request.
func (r *StartStreamEncryptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
