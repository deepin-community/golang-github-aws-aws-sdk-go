// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package medialiveiface provides an interface to enable mocking the AWS Elemental MediaLive service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package medialiveiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/medialive"
)

// MediaLiveAPI provides an interface to enable mocking the
// medialive.MediaLive service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Elemental MediaLive.
//	func myFunc(svc medialiveiface.MediaLiveAPI) bool {
//	    // Make svc.AcceptInputDeviceTransfer request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := medialive.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockMediaLiveClient struct {
//	    medialiveiface.MediaLiveAPI
//	}
//	func (m *mockMediaLiveClient) AcceptInputDeviceTransfer(input *medialive.AcceptInputDeviceTransferInput) (*medialive.AcceptInputDeviceTransferOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockMediaLiveClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type MediaLiveAPI interface {
	AcceptInputDeviceTransfer(*medialive.AcceptInputDeviceTransferInput) (*medialive.AcceptInputDeviceTransferOutput, error)
	AcceptInputDeviceTransferWithContext(aws.Context, *medialive.AcceptInputDeviceTransferInput, ...request.Option) (*medialive.AcceptInputDeviceTransferOutput, error)
	AcceptInputDeviceTransferRequest(*medialive.AcceptInputDeviceTransferInput) (*request.Request, *medialive.AcceptInputDeviceTransferOutput)

	BatchDelete(*medialive.BatchDeleteInput) (*medialive.BatchDeleteOutput, error)
	BatchDeleteWithContext(aws.Context, *medialive.BatchDeleteInput, ...request.Option) (*medialive.BatchDeleteOutput, error)
	BatchDeleteRequest(*medialive.BatchDeleteInput) (*request.Request, *medialive.BatchDeleteOutput)

	BatchStart(*medialive.BatchStartInput) (*medialive.BatchStartOutput, error)
	BatchStartWithContext(aws.Context, *medialive.BatchStartInput, ...request.Option) (*medialive.BatchStartOutput, error)
	BatchStartRequest(*medialive.BatchStartInput) (*request.Request, *medialive.BatchStartOutput)

	BatchStop(*medialive.BatchStopInput) (*medialive.BatchStopOutput, error)
	BatchStopWithContext(aws.Context, *medialive.BatchStopInput, ...request.Option) (*medialive.BatchStopOutput, error)
	BatchStopRequest(*medialive.BatchStopInput) (*request.Request, *medialive.BatchStopOutput)

	BatchUpdateSchedule(*medialive.BatchUpdateScheduleInput) (*medialive.BatchUpdateScheduleOutput, error)
	BatchUpdateScheduleWithContext(aws.Context, *medialive.BatchUpdateScheduleInput, ...request.Option) (*medialive.BatchUpdateScheduleOutput, error)
	BatchUpdateScheduleRequest(*medialive.BatchUpdateScheduleInput) (*request.Request, *medialive.BatchUpdateScheduleOutput)

	CancelInputDeviceTransfer(*medialive.CancelInputDeviceTransferInput) (*medialive.CancelInputDeviceTransferOutput, error)
	CancelInputDeviceTransferWithContext(aws.Context, *medialive.CancelInputDeviceTransferInput, ...request.Option) (*medialive.CancelInputDeviceTransferOutput, error)
	CancelInputDeviceTransferRequest(*medialive.CancelInputDeviceTransferInput) (*request.Request, *medialive.CancelInputDeviceTransferOutput)

	ClaimDevice(*medialive.ClaimDeviceInput) (*medialive.ClaimDeviceOutput, error)
	ClaimDeviceWithContext(aws.Context, *medialive.ClaimDeviceInput, ...request.Option) (*medialive.ClaimDeviceOutput, error)
	ClaimDeviceRequest(*medialive.ClaimDeviceInput) (*request.Request, *medialive.ClaimDeviceOutput)

	CreateChannel(*medialive.CreateChannelInput) (*medialive.CreateChannelOutput, error)
	CreateChannelWithContext(aws.Context, *medialive.CreateChannelInput, ...request.Option) (*medialive.CreateChannelOutput, error)
	CreateChannelRequest(*medialive.CreateChannelInput) (*request.Request, *medialive.CreateChannelOutput)

	CreateInput(*medialive.CreateInputInput) (*medialive.CreateInputOutput, error)
	CreateInputWithContext(aws.Context, *medialive.CreateInputInput, ...request.Option) (*medialive.CreateInputOutput, error)
	CreateInputRequest(*medialive.CreateInputInput) (*request.Request, *medialive.CreateInputOutput)

	CreateInputSecurityGroup(*medialive.CreateInputSecurityGroupInput) (*medialive.CreateInputSecurityGroupOutput, error)
	CreateInputSecurityGroupWithContext(aws.Context, *medialive.CreateInputSecurityGroupInput, ...request.Option) (*medialive.CreateInputSecurityGroupOutput, error)
	CreateInputSecurityGroupRequest(*medialive.CreateInputSecurityGroupInput) (*request.Request, *medialive.CreateInputSecurityGroupOutput)

	CreateMultiplex(*medialive.CreateMultiplexInput) (*medialive.CreateMultiplexOutput, error)
	CreateMultiplexWithContext(aws.Context, *medialive.CreateMultiplexInput, ...request.Option) (*medialive.CreateMultiplexOutput, error)
	CreateMultiplexRequest(*medialive.CreateMultiplexInput) (*request.Request, *medialive.CreateMultiplexOutput)

	CreateMultiplexProgram(*medialive.CreateMultiplexProgramInput) (*medialive.CreateMultiplexProgramOutput, error)
	CreateMultiplexProgramWithContext(aws.Context, *medialive.CreateMultiplexProgramInput, ...request.Option) (*medialive.CreateMultiplexProgramOutput, error)
	CreateMultiplexProgramRequest(*medialive.CreateMultiplexProgramInput) (*request.Request, *medialive.CreateMultiplexProgramOutput)

	CreatePartnerInput(*medialive.CreatePartnerInputInput) (*medialive.CreatePartnerInputOutput, error)
	CreatePartnerInputWithContext(aws.Context, *medialive.CreatePartnerInputInput, ...request.Option) (*medialive.CreatePartnerInputOutput, error)
	CreatePartnerInputRequest(*medialive.CreatePartnerInputInput) (*request.Request, *medialive.CreatePartnerInputOutput)

	CreateTags(*medialive.CreateTagsInput) (*medialive.CreateTagsOutput, error)
	CreateTagsWithContext(aws.Context, *medialive.CreateTagsInput, ...request.Option) (*medialive.CreateTagsOutput, error)
	CreateTagsRequest(*medialive.CreateTagsInput) (*request.Request, *medialive.CreateTagsOutput)

	DeleteChannel(*medialive.DeleteChannelInput) (*medialive.DeleteChannelOutput, error)
	DeleteChannelWithContext(aws.Context, *medialive.DeleteChannelInput, ...request.Option) (*medialive.DeleteChannelOutput, error)
	DeleteChannelRequest(*medialive.DeleteChannelInput) (*request.Request, *medialive.DeleteChannelOutput)

	DeleteInput(*medialive.DeleteInputInput) (*medialive.DeleteInputOutput, error)
	DeleteInputWithContext(aws.Context, *medialive.DeleteInputInput, ...request.Option) (*medialive.DeleteInputOutput, error)
	DeleteInputRequest(*medialive.DeleteInputInput) (*request.Request, *medialive.DeleteInputOutput)

	DeleteInputSecurityGroup(*medialive.DeleteInputSecurityGroupInput) (*medialive.DeleteInputSecurityGroupOutput, error)
	DeleteInputSecurityGroupWithContext(aws.Context, *medialive.DeleteInputSecurityGroupInput, ...request.Option) (*medialive.DeleteInputSecurityGroupOutput, error)
	DeleteInputSecurityGroupRequest(*medialive.DeleteInputSecurityGroupInput) (*request.Request, *medialive.DeleteInputSecurityGroupOutput)

	DeleteMultiplex(*medialive.DeleteMultiplexInput) (*medialive.DeleteMultiplexOutput, error)
	DeleteMultiplexWithContext(aws.Context, *medialive.DeleteMultiplexInput, ...request.Option) (*medialive.DeleteMultiplexOutput, error)
	DeleteMultiplexRequest(*medialive.DeleteMultiplexInput) (*request.Request, *medialive.DeleteMultiplexOutput)

	DeleteMultiplexProgram(*medialive.DeleteMultiplexProgramInput) (*medialive.DeleteMultiplexProgramOutput, error)
	DeleteMultiplexProgramWithContext(aws.Context, *medialive.DeleteMultiplexProgramInput, ...request.Option) (*medialive.DeleteMultiplexProgramOutput, error)
	DeleteMultiplexProgramRequest(*medialive.DeleteMultiplexProgramInput) (*request.Request, *medialive.DeleteMultiplexProgramOutput)

	DeleteReservation(*medialive.DeleteReservationInput) (*medialive.DeleteReservationOutput, error)
	DeleteReservationWithContext(aws.Context, *medialive.DeleteReservationInput, ...request.Option) (*medialive.DeleteReservationOutput, error)
	DeleteReservationRequest(*medialive.DeleteReservationInput) (*request.Request, *medialive.DeleteReservationOutput)

	DeleteSchedule(*medialive.DeleteScheduleInput) (*medialive.DeleteScheduleOutput, error)
	DeleteScheduleWithContext(aws.Context, *medialive.DeleteScheduleInput, ...request.Option) (*medialive.DeleteScheduleOutput, error)
	DeleteScheduleRequest(*medialive.DeleteScheduleInput) (*request.Request, *medialive.DeleteScheduleOutput)

	DeleteTags(*medialive.DeleteTagsInput) (*medialive.DeleteTagsOutput, error)
	DeleteTagsWithContext(aws.Context, *medialive.DeleteTagsInput, ...request.Option) (*medialive.DeleteTagsOutput, error)
	DeleteTagsRequest(*medialive.DeleteTagsInput) (*request.Request, *medialive.DeleteTagsOutput)

	DescribeAccountConfiguration(*medialive.DescribeAccountConfigurationInput) (*medialive.DescribeAccountConfigurationOutput, error)
	DescribeAccountConfigurationWithContext(aws.Context, *medialive.DescribeAccountConfigurationInput, ...request.Option) (*medialive.DescribeAccountConfigurationOutput, error)
	DescribeAccountConfigurationRequest(*medialive.DescribeAccountConfigurationInput) (*request.Request, *medialive.DescribeAccountConfigurationOutput)

	DescribeChannel(*medialive.DescribeChannelInput) (*medialive.DescribeChannelOutput, error)
	DescribeChannelWithContext(aws.Context, *medialive.DescribeChannelInput, ...request.Option) (*medialive.DescribeChannelOutput, error)
	DescribeChannelRequest(*medialive.DescribeChannelInput) (*request.Request, *medialive.DescribeChannelOutput)

	DescribeInput(*medialive.DescribeInputInput) (*medialive.DescribeInputOutput, error)
	DescribeInputWithContext(aws.Context, *medialive.DescribeInputInput, ...request.Option) (*medialive.DescribeInputOutput, error)
	DescribeInputRequest(*medialive.DescribeInputInput) (*request.Request, *medialive.DescribeInputOutput)

	DescribeInputDevice(*medialive.DescribeInputDeviceInput) (*medialive.DescribeInputDeviceOutput, error)
	DescribeInputDeviceWithContext(aws.Context, *medialive.DescribeInputDeviceInput, ...request.Option) (*medialive.DescribeInputDeviceOutput, error)
	DescribeInputDeviceRequest(*medialive.DescribeInputDeviceInput) (*request.Request, *medialive.DescribeInputDeviceOutput)

	DescribeInputDeviceThumbnail(*medialive.DescribeInputDeviceThumbnailInput) (*medialive.DescribeInputDeviceThumbnailOutput, error)
	DescribeInputDeviceThumbnailWithContext(aws.Context, *medialive.DescribeInputDeviceThumbnailInput, ...request.Option) (*medialive.DescribeInputDeviceThumbnailOutput, error)
	DescribeInputDeviceThumbnailRequest(*medialive.DescribeInputDeviceThumbnailInput) (*request.Request, *medialive.DescribeInputDeviceThumbnailOutput)

	DescribeInputSecurityGroup(*medialive.DescribeInputSecurityGroupInput) (*medialive.DescribeInputSecurityGroupOutput, error)
	DescribeInputSecurityGroupWithContext(aws.Context, *medialive.DescribeInputSecurityGroupInput, ...request.Option) (*medialive.DescribeInputSecurityGroupOutput, error)
	DescribeInputSecurityGroupRequest(*medialive.DescribeInputSecurityGroupInput) (*request.Request, *medialive.DescribeInputSecurityGroupOutput)

	DescribeMultiplex(*medialive.DescribeMultiplexInput) (*medialive.DescribeMultiplexOutput, error)
	DescribeMultiplexWithContext(aws.Context, *medialive.DescribeMultiplexInput, ...request.Option) (*medialive.DescribeMultiplexOutput, error)
	DescribeMultiplexRequest(*medialive.DescribeMultiplexInput) (*request.Request, *medialive.DescribeMultiplexOutput)

	DescribeMultiplexProgram(*medialive.DescribeMultiplexProgramInput) (*medialive.DescribeMultiplexProgramOutput, error)
	DescribeMultiplexProgramWithContext(aws.Context, *medialive.DescribeMultiplexProgramInput, ...request.Option) (*medialive.DescribeMultiplexProgramOutput, error)
	DescribeMultiplexProgramRequest(*medialive.DescribeMultiplexProgramInput) (*request.Request, *medialive.DescribeMultiplexProgramOutput)

	DescribeOffering(*medialive.DescribeOfferingInput) (*medialive.DescribeOfferingOutput, error)
	DescribeOfferingWithContext(aws.Context, *medialive.DescribeOfferingInput, ...request.Option) (*medialive.DescribeOfferingOutput, error)
	DescribeOfferingRequest(*medialive.DescribeOfferingInput) (*request.Request, *medialive.DescribeOfferingOutput)

	DescribeReservation(*medialive.DescribeReservationInput) (*medialive.DescribeReservationOutput, error)
	DescribeReservationWithContext(aws.Context, *medialive.DescribeReservationInput, ...request.Option) (*medialive.DescribeReservationOutput, error)
	DescribeReservationRequest(*medialive.DescribeReservationInput) (*request.Request, *medialive.DescribeReservationOutput)

	DescribeSchedule(*medialive.DescribeScheduleInput) (*medialive.DescribeScheduleOutput, error)
	DescribeScheduleWithContext(aws.Context, *medialive.DescribeScheduleInput, ...request.Option) (*medialive.DescribeScheduleOutput, error)
	DescribeScheduleRequest(*medialive.DescribeScheduleInput) (*request.Request, *medialive.DescribeScheduleOutput)

	DescribeSchedulePages(*medialive.DescribeScheduleInput, func(*medialive.DescribeScheduleOutput, bool) bool) error
	DescribeSchedulePagesWithContext(aws.Context, *medialive.DescribeScheduleInput, func(*medialive.DescribeScheduleOutput, bool) bool, ...request.Option) error

	DescribeThumbnails(*medialive.DescribeThumbnailsInput) (*medialive.DescribeThumbnailsOutput, error)
	DescribeThumbnailsWithContext(aws.Context, *medialive.DescribeThumbnailsInput, ...request.Option) (*medialive.DescribeThumbnailsOutput, error)
	DescribeThumbnailsRequest(*medialive.DescribeThumbnailsInput) (*request.Request, *medialive.DescribeThumbnailsOutput)

	ListChannels(*medialive.ListChannelsInput) (*medialive.ListChannelsOutput, error)
	ListChannelsWithContext(aws.Context, *medialive.ListChannelsInput, ...request.Option) (*medialive.ListChannelsOutput, error)
	ListChannelsRequest(*medialive.ListChannelsInput) (*request.Request, *medialive.ListChannelsOutput)

	ListChannelsPages(*medialive.ListChannelsInput, func(*medialive.ListChannelsOutput, bool) bool) error
	ListChannelsPagesWithContext(aws.Context, *medialive.ListChannelsInput, func(*medialive.ListChannelsOutput, bool) bool, ...request.Option) error

	ListInputDeviceTransfers(*medialive.ListInputDeviceTransfersInput) (*medialive.ListInputDeviceTransfersOutput, error)
	ListInputDeviceTransfersWithContext(aws.Context, *medialive.ListInputDeviceTransfersInput, ...request.Option) (*medialive.ListInputDeviceTransfersOutput, error)
	ListInputDeviceTransfersRequest(*medialive.ListInputDeviceTransfersInput) (*request.Request, *medialive.ListInputDeviceTransfersOutput)

	ListInputDeviceTransfersPages(*medialive.ListInputDeviceTransfersInput, func(*medialive.ListInputDeviceTransfersOutput, bool) bool) error
	ListInputDeviceTransfersPagesWithContext(aws.Context, *medialive.ListInputDeviceTransfersInput, func(*medialive.ListInputDeviceTransfersOutput, bool) bool, ...request.Option) error

	ListInputDevices(*medialive.ListInputDevicesInput) (*medialive.ListInputDevicesOutput, error)
	ListInputDevicesWithContext(aws.Context, *medialive.ListInputDevicesInput, ...request.Option) (*medialive.ListInputDevicesOutput, error)
	ListInputDevicesRequest(*medialive.ListInputDevicesInput) (*request.Request, *medialive.ListInputDevicesOutput)

	ListInputDevicesPages(*medialive.ListInputDevicesInput, func(*medialive.ListInputDevicesOutput, bool) bool) error
	ListInputDevicesPagesWithContext(aws.Context, *medialive.ListInputDevicesInput, func(*medialive.ListInputDevicesOutput, bool) bool, ...request.Option) error

	ListInputSecurityGroups(*medialive.ListInputSecurityGroupsInput) (*medialive.ListInputSecurityGroupsOutput, error)
	ListInputSecurityGroupsWithContext(aws.Context, *medialive.ListInputSecurityGroupsInput, ...request.Option) (*medialive.ListInputSecurityGroupsOutput, error)
	ListInputSecurityGroupsRequest(*medialive.ListInputSecurityGroupsInput) (*request.Request, *medialive.ListInputSecurityGroupsOutput)

	ListInputSecurityGroupsPages(*medialive.ListInputSecurityGroupsInput, func(*medialive.ListInputSecurityGroupsOutput, bool) bool) error
	ListInputSecurityGroupsPagesWithContext(aws.Context, *medialive.ListInputSecurityGroupsInput, func(*medialive.ListInputSecurityGroupsOutput, bool) bool, ...request.Option) error

	ListInputs(*medialive.ListInputsInput) (*medialive.ListInputsOutput, error)
	ListInputsWithContext(aws.Context, *medialive.ListInputsInput, ...request.Option) (*medialive.ListInputsOutput, error)
	ListInputsRequest(*medialive.ListInputsInput) (*request.Request, *medialive.ListInputsOutput)

	ListInputsPages(*medialive.ListInputsInput, func(*medialive.ListInputsOutput, bool) bool) error
	ListInputsPagesWithContext(aws.Context, *medialive.ListInputsInput, func(*medialive.ListInputsOutput, bool) bool, ...request.Option) error

	ListMultiplexPrograms(*medialive.ListMultiplexProgramsInput) (*medialive.ListMultiplexProgramsOutput, error)
	ListMultiplexProgramsWithContext(aws.Context, *medialive.ListMultiplexProgramsInput, ...request.Option) (*medialive.ListMultiplexProgramsOutput, error)
	ListMultiplexProgramsRequest(*medialive.ListMultiplexProgramsInput) (*request.Request, *medialive.ListMultiplexProgramsOutput)

	ListMultiplexProgramsPages(*medialive.ListMultiplexProgramsInput, func(*medialive.ListMultiplexProgramsOutput, bool) bool) error
	ListMultiplexProgramsPagesWithContext(aws.Context, *medialive.ListMultiplexProgramsInput, func(*medialive.ListMultiplexProgramsOutput, bool) bool, ...request.Option) error

	ListMultiplexes(*medialive.ListMultiplexesInput) (*medialive.ListMultiplexesOutput, error)
	ListMultiplexesWithContext(aws.Context, *medialive.ListMultiplexesInput, ...request.Option) (*medialive.ListMultiplexesOutput, error)
	ListMultiplexesRequest(*medialive.ListMultiplexesInput) (*request.Request, *medialive.ListMultiplexesOutput)

	ListMultiplexesPages(*medialive.ListMultiplexesInput, func(*medialive.ListMultiplexesOutput, bool) bool) error
	ListMultiplexesPagesWithContext(aws.Context, *medialive.ListMultiplexesInput, func(*medialive.ListMultiplexesOutput, bool) bool, ...request.Option) error

	ListOfferings(*medialive.ListOfferingsInput) (*medialive.ListOfferingsOutput, error)
	ListOfferingsWithContext(aws.Context, *medialive.ListOfferingsInput, ...request.Option) (*medialive.ListOfferingsOutput, error)
	ListOfferingsRequest(*medialive.ListOfferingsInput) (*request.Request, *medialive.ListOfferingsOutput)

	ListOfferingsPages(*medialive.ListOfferingsInput, func(*medialive.ListOfferingsOutput, bool) bool) error
	ListOfferingsPagesWithContext(aws.Context, *medialive.ListOfferingsInput, func(*medialive.ListOfferingsOutput, bool) bool, ...request.Option) error

	ListReservations(*medialive.ListReservationsInput) (*medialive.ListReservationsOutput, error)
	ListReservationsWithContext(aws.Context, *medialive.ListReservationsInput, ...request.Option) (*medialive.ListReservationsOutput, error)
	ListReservationsRequest(*medialive.ListReservationsInput) (*request.Request, *medialive.ListReservationsOutput)

	ListReservationsPages(*medialive.ListReservationsInput, func(*medialive.ListReservationsOutput, bool) bool) error
	ListReservationsPagesWithContext(aws.Context, *medialive.ListReservationsInput, func(*medialive.ListReservationsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*medialive.ListTagsForResourceInput) (*medialive.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *medialive.ListTagsForResourceInput, ...request.Option) (*medialive.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*medialive.ListTagsForResourceInput) (*request.Request, *medialive.ListTagsForResourceOutput)

	PurchaseOffering(*medialive.PurchaseOfferingInput) (*medialive.PurchaseOfferingOutput, error)
	PurchaseOfferingWithContext(aws.Context, *medialive.PurchaseOfferingInput, ...request.Option) (*medialive.PurchaseOfferingOutput, error)
	PurchaseOfferingRequest(*medialive.PurchaseOfferingInput) (*request.Request, *medialive.PurchaseOfferingOutput)

	RebootInputDevice(*medialive.RebootInputDeviceInput) (*medialive.RebootInputDeviceOutput, error)
	RebootInputDeviceWithContext(aws.Context, *medialive.RebootInputDeviceInput, ...request.Option) (*medialive.RebootInputDeviceOutput, error)
	RebootInputDeviceRequest(*medialive.RebootInputDeviceInput) (*request.Request, *medialive.RebootInputDeviceOutput)

	RejectInputDeviceTransfer(*medialive.RejectInputDeviceTransferInput) (*medialive.RejectInputDeviceTransferOutput, error)
	RejectInputDeviceTransferWithContext(aws.Context, *medialive.RejectInputDeviceTransferInput, ...request.Option) (*medialive.RejectInputDeviceTransferOutput, error)
	RejectInputDeviceTransferRequest(*medialive.RejectInputDeviceTransferInput) (*request.Request, *medialive.RejectInputDeviceTransferOutput)

	StartChannel(*medialive.StartChannelInput) (*medialive.StartChannelOutput, error)
	StartChannelWithContext(aws.Context, *medialive.StartChannelInput, ...request.Option) (*medialive.StartChannelOutput, error)
	StartChannelRequest(*medialive.StartChannelInput) (*request.Request, *medialive.StartChannelOutput)

	StartInputDevice(*medialive.StartInputDeviceInput) (*medialive.StartInputDeviceOutput, error)
	StartInputDeviceWithContext(aws.Context, *medialive.StartInputDeviceInput, ...request.Option) (*medialive.StartInputDeviceOutput, error)
	StartInputDeviceRequest(*medialive.StartInputDeviceInput) (*request.Request, *medialive.StartInputDeviceOutput)

	StartInputDeviceMaintenanceWindow(*medialive.StartInputDeviceMaintenanceWindowInput) (*medialive.StartInputDeviceMaintenanceWindowOutput, error)
	StartInputDeviceMaintenanceWindowWithContext(aws.Context, *medialive.StartInputDeviceMaintenanceWindowInput, ...request.Option) (*medialive.StartInputDeviceMaintenanceWindowOutput, error)
	StartInputDeviceMaintenanceWindowRequest(*medialive.StartInputDeviceMaintenanceWindowInput) (*request.Request, *medialive.StartInputDeviceMaintenanceWindowOutput)

	StartMultiplex(*medialive.StartMultiplexInput) (*medialive.StartMultiplexOutput, error)
	StartMultiplexWithContext(aws.Context, *medialive.StartMultiplexInput, ...request.Option) (*medialive.StartMultiplexOutput, error)
	StartMultiplexRequest(*medialive.StartMultiplexInput) (*request.Request, *medialive.StartMultiplexOutput)

	StopChannel(*medialive.StopChannelInput) (*medialive.StopChannelOutput, error)
	StopChannelWithContext(aws.Context, *medialive.StopChannelInput, ...request.Option) (*medialive.StopChannelOutput, error)
	StopChannelRequest(*medialive.StopChannelInput) (*request.Request, *medialive.StopChannelOutput)

	StopInputDevice(*medialive.StopInputDeviceInput) (*medialive.StopInputDeviceOutput, error)
	StopInputDeviceWithContext(aws.Context, *medialive.StopInputDeviceInput, ...request.Option) (*medialive.StopInputDeviceOutput, error)
	StopInputDeviceRequest(*medialive.StopInputDeviceInput) (*request.Request, *medialive.StopInputDeviceOutput)

	StopMultiplex(*medialive.StopMultiplexInput) (*medialive.StopMultiplexOutput, error)
	StopMultiplexWithContext(aws.Context, *medialive.StopMultiplexInput, ...request.Option) (*medialive.StopMultiplexOutput, error)
	StopMultiplexRequest(*medialive.StopMultiplexInput) (*request.Request, *medialive.StopMultiplexOutput)

	TransferInputDevice(*medialive.TransferInputDeviceInput) (*medialive.TransferInputDeviceOutput, error)
	TransferInputDeviceWithContext(aws.Context, *medialive.TransferInputDeviceInput, ...request.Option) (*medialive.TransferInputDeviceOutput, error)
	TransferInputDeviceRequest(*medialive.TransferInputDeviceInput) (*request.Request, *medialive.TransferInputDeviceOutput)

	UpdateAccountConfiguration(*medialive.UpdateAccountConfigurationInput) (*medialive.UpdateAccountConfigurationOutput, error)
	UpdateAccountConfigurationWithContext(aws.Context, *medialive.UpdateAccountConfigurationInput, ...request.Option) (*medialive.UpdateAccountConfigurationOutput, error)
	UpdateAccountConfigurationRequest(*medialive.UpdateAccountConfigurationInput) (*request.Request, *medialive.UpdateAccountConfigurationOutput)

	UpdateChannel(*medialive.UpdateChannelInput) (*medialive.UpdateChannelOutput, error)
	UpdateChannelWithContext(aws.Context, *medialive.UpdateChannelInput, ...request.Option) (*medialive.UpdateChannelOutput, error)
	UpdateChannelRequest(*medialive.UpdateChannelInput) (*request.Request, *medialive.UpdateChannelOutput)

	UpdateChannelClass(*medialive.UpdateChannelClassInput) (*medialive.UpdateChannelClassOutput, error)
	UpdateChannelClassWithContext(aws.Context, *medialive.UpdateChannelClassInput, ...request.Option) (*medialive.UpdateChannelClassOutput, error)
	UpdateChannelClassRequest(*medialive.UpdateChannelClassInput) (*request.Request, *medialive.UpdateChannelClassOutput)

	UpdateInput(*medialive.UpdateInputInput) (*medialive.UpdateInputOutput, error)
	UpdateInputWithContext(aws.Context, *medialive.UpdateInputInput, ...request.Option) (*medialive.UpdateInputOutput, error)
	UpdateInputRequest(*medialive.UpdateInputInput) (*request.Request, *medialive.UpdateInputOutput)

	UpdateInputDevice(*medialive.UpdateInputDeviceInput) (*medialive.UpdateInputDeviceOutput, error)
	UpdateInputDeviceWithContext(aws.Context, *medialive.UpdateInputDeviceInput, ...request.Option) (*medialive.UpdateInputDeviceOutput, error)
	UpdateInputDeviceRequest(*medialive.UpdateInputDeviceInput) (*request.Request, *medialive.UpdateInputDeviceOutput)

	UpdateInputSecurityGroup(*medialive.UpdateInputSecurityGroupInput) (*medialive.UpdateInputSecurityGroupOutput, error)
	UpdateInputSecurityGroupWithContext(aws.Context, *medialive.UpdateInputSecurityGroupInput, ...request.Option) (*medialive.UpdateInputSecurityGroupOutput, error)
	UpdateInputSecurityGroupRequest(*medialive.UpdateInputSecurityGroupInput) (*request.Request, *medialive.UpdateInputSecurityGroupOutput)

	UpdateMultiplex(*medialive.UpdateMultiplexInput) (*medialive.UpdateMultiplexOutput, error)
	UpdateMultiplexWithContext(aws.Context, *medialive.UpdateMultiplexInput, ...request.Option) (*medialive.UpdateMultiplexOutput, error)
	UpdateMultiplexRequest(*medialive.UpdateMultiplexInput) (*request.Request, *medialive.UpdateMultiplexOutput)

	UpdateMultiplexProgram(*medialive.UpdateMultiplexProgramInput) (*medialive.UpdateMultiplexProgramOutput, error)
	UpdateMultiplexProgramWithContext(aws.Context, *medialive.UpdateMultiplexProgramInput, ...request.Option) (*medialive.UpdateMultiplexProgramOutput, error)
	UpdateMultiplexProgramRequest(*medialive.UpdateMultiplexProgramInput) (*request.Request, *medialive.UpdateMultiplexProgramOutput)

	UpdateReservation(*medialive.UpdateReservationInput) (*medialive.UpdateReservationOutput, error)
	UpdateReservationWithContext(aws.Context, *medialive.UpdateReservationInput, ...request.Option) (*medialive.UpdateReservationOutput, error)
	UpdateReservationRequest(*medialive.UpdateReservationInput) (*request.Request, *medialive.UpdateReservationOutput)

	WaitUntilChannelCreated(*medialive.DescribeChannelInput) error
	WaitUntilChannelCreatedWithContext(aws.Context, *medialive.DescribeChannelInput, ...request.WaiterOption) error

	WaitUntilChannelDeleted(*medialive.DescribeChannelInput) error
	WaitUntilChannelDeletedWithContext(aws.Context, *medialive.DescribeChannelInput, ...request.WaiterOption) error

	WaitUntilChannelRunning(*medialive.DescribeChannelInput) error
	WaitUntilChannelRunningWithContext(aws.Context, *medialive.DescribeChannelInput, ...request.WaiterOption) error

	WaitUntilChannelStopped(*medialive.DescribeChannelInput) error
	WaitUntilChannelStoppedWithContext(aws.Context, *medialive.DescribeChannelInput, ...request.WaiterOption) error

	WaitUntilInputAttached(*medialive.DescribeInputInput) error
	WaitUntilInputAttachedWithContext(aws.Context, *medialive.DescribeInputInput, ...request.WaiterOption) error

	WaitUntilInputDeleted(*medialive.DescribeInputInput) error
	WaitUntilInputDeletedWithContext(aws.Context, *medialive.DescribeInputInput, ...request.WaiterOption) error

	WaitUntilInputDetached(*medialive.DescribeInputInput) error
	WaitUntilInputDetachedWithContext(aws.Context, *medialive.DescribeInputInput, ...request.WaiterOption) error

	WaitUntilMultiplexCreated(*medialive.DescribeMultiplexInput) error
	WaitUntilMultiplexCreatedWithContext(aws.Context, *medialive.DescribeMultiplexInput, ...request.WaiterOption) error

	WaitUntilMultiplexDeleted(*medialive.DescribeMultiplexInput) error
	WaitUntilMultiplexDeletedWithContext(aws.Context, *medialive.DescribeMultiplexInput, ...request.WaiterOption) error

	WaitUntilMultiplexRunning(*medialive.DescribeMultiplexInput) error
	WaitUntilMultiplexRunningWithContext(aws.Context, *medialive.DescribeMultiplexInput, ...request.WaiterOption) error

	WaitUntilMultiplexStopped(*medialive.DescribeMultiplexInput) error
	WaitUntilMultiplexStoppedWithContext(aws.Context, *medialive.DescribeMultiplexInput, ...request.WaiterOption) error
}

var _ MediaLiveAPI = (*medialive.MediaLive)(nil)
