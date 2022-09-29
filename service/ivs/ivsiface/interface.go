// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package ivsiface provides an interface to enable mocking the Amazon Interactive Video Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package ivsiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ivs"
)

// IVSAPI provides an interface to enable mocking the
// ivs.IVS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Interactive Video Service.
//    func myFunc(svc ivsiface.IVSAPI) bool {
//        // Make svc.BatchGetChannel request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := ivs.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockIVSClient struct {
//        ivsiface.IVSAPI
//    }
//    func (m *mockIVSClient) BatchGetChannel(input *ivs.BatchGetChannelInput) (*ivs.BatchGetChannelOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockIVSClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type IVSAPI interface {
	BatchGetChannel(*ivs.BatchGetChannelInput) (*ivs.BatchGetChannelOutput, error)
	BatchGetChannelWithContext(aws.Context, *ivs.BatchGetChannelInput, ...request.Option) (*ivs.BatchGetChannelOutput, error)
	BatchGetChannelRequest(*ivs.BatchGetChannelInput) (*request.Request, *ivs.BatchGetChannelOutput)

	BatchGetStreamKey(*ivs.BatchGetStreamKeyInput) (*ivs.BatchGetStreamKeyOutput, error)
	BatchGetStreamKeyWithContext(aws.Context, *ivs.BatchGetStreamKeyInput, ...request.Option) (*ivs.BatchGetStreamKeyOutput, error)
	BatchGetStreamKeyRequest(*ivs.BatchGetStreamKeyInput) (*request.Request, *ivs.BatchGetStreamKeyOutput)

	CreateChannel(*ivs.CreateChannelInput) (*ivs.CreateChannelOutput, error)
	CreateChannelWithContext(aws.Context, *ivs.CreateChannelInput, ...request.Option) (*ivs.CreateChannelOutput, error)
	CreateChannelRequest(*ivs.CreateChannelInput) (*request.Request, *ivs.CreateChannelOutput)

	CreateRecordingConfiguration(*ivs.CreateRecordingConfigurationInput) (*ivs.CreateRecordingConfigurationOutput, error)
	CreateRecordingConfigurationWithContext(aws.Context, *ivs.CreateRecordingConfigurationInput, ...request.Option) (*ivs.CreateRecordingConfigurationOutput, error)
	CreateRecordingConfigurationRequest(*ivs.CreateRecordingConfigurationInput) (*request.Request, *ivs.CreateRecordingConfigurationOutput)

	CreateStreamKey(*ivs.CreateStreamKeyInput) (*ivs.CreateStreamKeyOutput, error)
	CreateStreamKeyWithContext(aws.Context, *ivs.CreateStreamKeyInput, ...request.Option) (*ivs.CreateStreamKeyOutput, error)
	CreateStreamKeyRequest(*ivs.CreateStreamKeyInput) (*request.Request, *ivs.CreateStreamKeyOutput)

	DeleteChannel(*ivs.DeleteChannelInput) (*ivs.DeleteChannelOutput, error)
	DeleteChannelWithContext(aws.Context, *ivs.DeleteChannelInput, ...request.Option) (*ivs.DeleteChannelOutput, error)
	DeleteChannelRequest(*ivs.DeleteChannelInput) (*request.Request, *ivs.DeleteChannelOutput)

	DeletePlaybackKeyPair(*ivs.DeletePlaybackKeyPairInput) (*ivs.DeletePlaybackKeyPairOutput, error)
	DeletePlaybackKeyPairWithContext(aws.Context, *ivs.DeletePlaybackKeyPairInput, ...request.Option) (*ivs.DeletePlaybackKeyPairOutput, error)
	DeletePlaybackKeyPairRequest(*ivs.DeletePlaybackKeyPairInput) (*request.Request, *ivs.DeletePlaybackKeyPairOutput)

	DeleteRecordingConfiguration(*ivs.DeleteRecordingConfigurationInput) (*ivs.DeleteRecordingConfigurationOutput, error)
	DeleteRecordingConfigurationWithContext(aws.Context, *ivs.DeleteRecordingConfigurationInput, ...request.Option) (*ivs.DeleteRecordingConfigurationOutput, error)
	DeleteRecordingConfigurationRequest(*ivs.DeleteRecordingConfigurationInput) (*request.Request, *ivs.DeleteRecordingConfigurationOutput)

	DeleteStreamKey(*ivs.DeleteStreamKeyInput) (*ivs.DeleteStreamKeyOutput, error)
	DeleteStreamKeyWithContext(aws.Context, *ivs.DeleteStreamKeyInput, ...request.Option) (*ivs.DeleteStreamKeyOutput, error)
	DeleteStreamKeyRequest(*ivs.DeleteStreamKeyInput) (*request.Request, *ivs.DeleteStreamKeyOutput)

	GetChannel(*ivs.GetChannelInput) (*ivs.GetChannelOutput, error)
	GetChannelWithContext(aws.Context, *ivs.GetChannelInput, ...request.Option) (*ivs.GetChannelOutput, error)
	GetChannelRequest(*ivs.GetChannelInput) (*request.Request, *ivs.GetChannelOutput)

	GetPlaybackKeyPair(*ivs.GetPlaybackKeyPairInput) (*ivs.GetPlaybackKeyPairOutput, error)
	GetPlaybackKeyPairWithContext(aws.Context, *ivs.GetPlaybackKeyPairInput, ...request.Option) (*ivs.GetPlaybackKeyPairOutput, error)
	GetPlaybackKeyPairRequest(*ivs.GetPlaybackKeyPairInput) (*request.Request, *ivs.GetPlaybackKeyPairOutput)

	GetRecordingConfiguration(*ivs.GetRecordingConfigurationInput) (*ivs.GetRecordingConfigurationOutput, error)
	GetRecordingConfigurationWithContext(aws.Context, *ivs.GetRecordingConfigurationInput, ...request.Option) (*ivs.GetRecordingConfigurationOutput, error)
	GetRecordingConfigurationRequest(*ivs.GetRecordingConfigurationInput) (*request.Request, *ivs.GetRecordingConfigurationOutput)

	GetStream(*ivs.GetStreamInput) (*ivs.GetStreamOutput, error)
	GetStreamWithContext(aws.Context, *ivs.GetStreamInput, ...request.Option) (*ivs.GetStreamOutput, error)
	GetStreamRequest(*ivs.GetStreamInput) (*request.Request, *ivs.GetStreamOutput)

	GetStreamKey(*ivs.GetStreamKeyInput) (*ivs.GetStreamKeyOutput, error)
	GetStreamKeyWithContext(aws.Context, *ivs.GetStreamKeyInput, ...request.Option) (*ivs.GetStreamKeyOutput, error)
	GetStreamKeyRequest(*ivs.GetStreamKeyInput) (*request.Request, *ivs.GetStreamKeyOutput)

	ImportPlaybackKeyPair(*ivs.ImportPlaybackKeyPairInput) (*ivs.ImportPlaybackKeyPairOutput, error)
	ImportPlaybackKeyPairWithContext(aws.Context, *ivs.ImportPlaybackKeyPairInput, ...request.Option) (*ivs.ImportPlaybackKeyPairOutput, error)
	ImportPlaybackKeyPairRequest(*ivs.ImportPlaybackKeyPairInput) (*request.Request, *ivs.ImportPlaybackKeyPairOutput)

	ListChannels(*ivs.ListChannelsInput) (*ivs.ListChannelsOutput, error)
	ListChannelsWithContext(aws.Context, *ivs.ListChannelsInput, ...request.Option) (*ivs.ListChannelsOutput, error)
	ListChannelsRequest(*ivs.ListChannelsInput) (*request.Request, *ivs.ListChannelsOutput)

	ListChannelsPages(*ivs.ListChannelsInput, func(*ivs.ListChannelsOutput, bool) bool) error
	ListChannelsPagesWithContext(aws.Context, *ivs.ListChannelsInput, func(*ivs.ListChannelsOutput, bool) bool, ...request.Option) error

	ListPlaybackKeyPairs(*ivs.ListPlaybackKeyPairsInput) (*ivs.ListPlaybackKeyPairsOutput, error)
	ListPlaybackKeyPairsWithContext(aws.Context, *ivs.ListPlaybackKeyPairsInput, ...request.Option) (*ivs.ListPlaybackKeyPairsOutput, error)
	ListPlaybackKeyPairsRequest(*ivs.ListPlaybackKeyPairsInput) (*request.Request, *ivs.ListPlaybackKeyPairsOutput)

	ListPlaybackKeyPairsPages(*ivs.ListPlaybackKeyPairsInput, func(*ivs.ListPlaybackKeyPairsOutput, bool) bool) error
	ListPlaybackKeyPairsPagesWithContext(aws.Context, *ivs.ListPlaybackKeyPairsInput, func(*ivs.ListPlaybackKeyPairsOutput, bool) bool, ...request.Option) error

	ListRecordingConfigurations(*ivs.ListRecordingConfigurationsInput) (*ivs.ListRecordingConfigurationsOutput, error)
	ListRecordingConfigurationsWithContext(aws.Context, *ivs.ListRecordingConfigurationsInput, ...request.Option) (*ivs.ListRecordingConfigurationsOutput, error)
	ListRecordingConfigurationsRequest(*ivs.ListRecordingConfigurationsInput) (*request.Request, *ivs.ListRecordingConfigurationsOutput)

	ListRecordingConfigurationsPages(*ivs.ListRecordingConfigurationsInput, func(*ivs.ListRecordingConfigurationsOutput, bool) bool) error
	ListRecordingConfigurationsPagesWithContext(aws.Context, *ivs.ListRecordingConfigurationsInput, func(*ivs.ListRecordingConfigurationsOutput, bool) bool, ...request.Option) error

	ListStreamKeys(*ivs.ListStreamKeysInput) (*ivs.ListStreamKeysOutput, error)
	ListStreamKeysWithContext(aws.Context, *ivs.ListStreamKeysInput, ...request.Option) (*ivs.ListStreamKeysOutput, error)
	ListStreamKeysRequest(*ivs.ListStreamKeysInput) (*request.Request, *ivs.ListStreamKeysOutput)

	ListStreamKeysPages(*ivs.ListStreamKeysInput, func(*ivs.ListStreamKeysOutput, bool) bool) error
	ListStreamKeysPagesWithContext(aws.Context, *ivs.ListStreamKeysInput, func(*ivs.ListStreamKeysOutput, bool) bool, ...request.Option) error

	ListStreams(*ivs.ListStreamsInput) (*ivs.ListStreamsOutput, error)
	ListStreamsWithContext(aws.Context, *ivs.ListStreamsInput, ...request.Option) (*ivs.ListStreamsOutput, error)
	ListStreamsRequest(*ivs.ListStreamsInput) (*request.Request, *ivs.ListStreamsOutput)

	ListStreamsPages(*ivs.ListStreamsInput, func(*ivs.ListStreamsOutput, bool) bool) error
	ListStreamsPagesWithContext(aws.Context, *ivs.ListStreamsInput, func(*ivs.ListStreamsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*ivs.ListTagsForResourceInput) (*ivs.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *ivs.ListTagsForResourceInput, ...request.Option) (*ivs.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*ivs.ListTagsForResourceInput) (*request.Request, *ivs.ListTagsForResourceOutput)

	PutMetadata(*ivs.PutMetadataInput) (*ivs.PutMetadataOutput, error)
	PutMetadataWithContext(aws.Context, *ivs.PutMetadataInput, ...request.Option) (*ivs.PutMetadataOutput, error)
	PutMetadataRequest(*ivs.PutMetadataInput) (*request.Request, *ivs.PutMetadataOutput)

	StopStream(*ivs.StopStreamInput) (*ivs.StopStreamOutput, error)
	StopStreamWithContext(aws.Context, *ivs.StopStreamInput, ...request.Option) (*ivs.StopStreamOutput, error)
	StopStreamRequest(*ivs.StopStreamInput) (*request.Request, *ivs.StopStreamOutput)

	TagResource(*ivs.TagResourceInput) (*ivs.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *ivs.TagResourceInput, ...request.Option) (*ivs.TagResourceOutput, error)
	TagResourceRequest(*ivs.TagResourceInput) (*request.Request, *ivs.TagResourceOutput)

	UntagResource(*ivs.UntagResourceInput) (*ivs.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *ivs.UntagResourceInput, ...request.Option) (*ivs.UntagResourceOutput, error)
	UntagResourceRequest(*ivs.UntagResourceInput) (*request.Request, *ivs.UntagResourceOutput)

	UpdateChannel(*ivs.UpdateChannelInput) (*ivs.UpdateChannelOutput, error)
	UpdateChannelWithContext(aws.Context, *ivs.UpdateChannelInput, ...request.Option) (*ivs.UpdateChannelOutput, error)
	UpdateChannelRequest(*ivs.UpdateChannelInput) (*request.Request, *ivs.UpdateChannelOutput)
}

var _ IVSAPI = (*ivs.IVS)(nil)
