// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codegururevieweriface provides an interface to enable mocking the Amazon CodeGuru Reviewer service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package codegururevieweriface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codegurureviewer"
)

// CodeGuruReviewerAPI provides an interface to enable mocking the
// codegurureviewer.CodeGuruReviewer service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon CodeGuru Reviewer.
//	func myFunc(svc codegururevieweriface.CodeGuruReviewerAPI) bool {
//	    // Make svc.AssociateRepository request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := codegurureviewer.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockCodeGuruReviewerClient struct {
//	    codegururevieweriface.CodeGuruReviewerAPI
//	}
//	func (m *mockCodeGuruReviewerClient) AssociateRepository(input *codegurureviewer.AssociateRepositoryInput) (*codegurureviewer.AssociateRepositoryOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockCodeGuruReviewerClient{}
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
type CodeGuruReviewerAPI interface {
	AssociateRepository(*codegurureviewer.AssociateRepositoryInput) (*codegurureviewer.AssociateRepositoryOutput, error)
	AssociateRepositoryWithContext(aws.Context, *codegurureviewer.AssociateRepositoryInput, ...request.Option) (*codegurureviewer.AssociateRepositoryOutput, error)
	AssociateRepositoryRequest(*codegurureviewer.AssociateRepositoryInput) (*request.Request, *codegurureviewer.AssociateRepositoryOutput)

	CreateCodeReview(*codegurureviewer.CreateCodeReviewInput) (*codegurureviewer.CreateCodeReviewOutput, error)
	CreateCodeReviewWithContext(aws.Context, *codegurureviewer.CreateCodeReviewInput, ...request.Option) (*codegurureviewer.CreateCodeReviewOutput, error)
	CreateCodeReviewRequest(*codegurureviewer.CreateCodeReviewInput) (*request.Request, *codegurureviewer.CreateCodeReviewOutput)

	DescribeCodeReview(*codegurureviewer.DescribeCodeReviewInput) (*codegurureviewer.DescribeCodeReviewOutput, error)
	DescribeCodeReviewWithContext(aws.Context, *codegurureviewer.DescribeCodeReviewInput, ...request.Option) (*codegurureviewer.DescribeCodeReviewOutput, error)
	DescribeCodeReviewRequest(*codegurureviewer.DescribeCodeReviewInput) (*request.Request, *codegurureviewer.DescribeCodeReviewOutput)

	DescribeRecommendationFeedback(*codegurureviewer.DescribeRecommendationFeedbackInput) (*codegurureviewer.DescribeRecommendationFeedbackOutput, error)
	DescribeRecommendationFeedbackWithContext(aws.Context, *codegurureviewer.DescribeRecommendationFeedbackInput, ...request.Option) (*codegurureviewer.DescribeRecommendationFeedbackOutput, error)
	DescribeRecommendationFeedbackRequest(*codegurureviewer.DescribeRecommendationFeedbackInput) (*request.Request, *codegurureviewer.DescribeRecommendationFeedbackOutput)

	DescribeRepositoryAssociation(*codegurureviewer.DescribeRepositoryAssociationInput) (*codegurureviewer.DescribeRepositoryAssociationOutput, error)
	DescribeRepositoryAssociationWithContext(aws.Context, *codegurureviewer.DescribeRepositoryAssociationInput, ...request.Option) (*codegurureviewer.DescribeRepositoryAssociationOutput, error)
	DescribeRepositoryAssociationRequest(*codegurureviewer.DescribeRepositoryAssociationInput) (*request.Request, *codegurureviewer.DescribeRepositoryAssociationOutput)

	DisassociateRepository(*codegurureviewer.DisassociateRepositoryInput) (*codegurureviewer.DisassociateRepositoryOutput, error)
	DisassociateRepositoryWithContext(aws.Context, *codegurureviewer.DisassociateRepositoryInput, ...request.Option) (*codegurureviewer.DisassociateRepositoryOutput, error)
	DisassociateRepositoryRequest(*codegurureviewer.DisassociateRepositoryInput) (*request.Request, *codegurureviewer.DisassociateRepositoryOutput)

	ListCodeReviews(*codegurureviewer.ListCodeReviewsInput) (*codegurureviewer.ListCodeReviewsOutput, error)
	ListCodeReviewsWithContext(aws.Context, *codegurureviewer.ListCodeReviewsInput, ...request.Option) (*codegurureviewer.ListCodeReviewsOutput, error)
	ListCodeReviewsRequest(*codegurureviewer.ListCodeReviewsInput) (*request.Request, *codegurureviewer.ListCodeReviewsOutput)

	ListCodeReviewsPages(*codegurureviewer.ListCodeReviewsInput, func(*codegurureviewer.ListCodeReviewsOutput, bool) bool) error
	ListCodeReviewsPagesWithContext(aws.Context, *codegurureviewer.ListCodeReviewsInput, func(*codegurureviewer.ListCodeReviewsOutput, bool) bool, ...request.Option) error

	ListRecommendationFeedback(*codegurureviewer.ListRecommendationFeedbackInput) (*codegurureviewer.ListRecommendationFeedbackOutput, error)
	ListRecommendationFeedbackWithContext(aws.Context, *codegurureviewer.ListRecommendationFeedbackInput, ...request.Option) (*codegurureviewer.ListRecommendationFeedbackOutput, error)
	ListRecommendationFeedbackRequest(*codegurureviewer.ListRecommendationFeedbackInput) (*request.Request, *codegurureviewer.ListRecommendationFeedbackOutput)

	ListRecommendationFeedbackPages(*codegurureviewer.ListRecommendationFeedbackInput, func(*codegurureviewer.ListRecommendationFeedbackOutput, bool) bool) error
	ListRecommendationFeedbackPagesWithContext(aws.Context, *codegurureviewer.ListRecommendationFeedbackInput, func(*codegurureviewer.ListRecommendationFeedbackOutput, bool) bool, ...request.Option) error

	ListRecommendations(*codegurureviewer.ListRecommendationsInput) (*codegurureviewer.ListRecommendationsOutput, error)
	ListRecommendationsWithContext(aws.Context, *codegurureviewer.ListRecommendationsInput, ...request.Option) (*codegurureviewer.ListRecommendationsOutput, error)
	ListRecommendationsRequest(*codegurureviewer.ListRecommendationsInput) (*request.Request, *codegurureviewer.ListRecommendationsOutput)

	ListRecommendationsPages(*codegurureviewer.ListRecommendationsInput, func(*codegurureviewer.ListRecommendationsOutput, bool) bool) error
	ListRecommendationsPagesWithContext(aws.Context, *codegurureviewer.ListRecommendationsInput, func(*codegurureviewer.ListRecommendationsOutput, bool) bool, ...request.Option) error

	ListRepositoryAssociations(*codegurureviewer.ListRepositoryAssociationsInput) (*codegurureviewer.ListRepositoryAssociationsOutput, error)
	ListRepositoryAssociationsWithContext(aws.Context, *codegurureviewer.ListRepositoryAssociationsInput, ...request.Option) (*codegurureviewer.ListRepositoryAssociationsOutput, error)
	ListRepositoryAssociationsRequest(*codegurureviewer.ListRepositoryAssociationsInput) (*request.Request, *codegurureviewer.ListRepositoryAssociationsOutput)

	ListRepositoryAssociationsPages(*codegurureviewer.ListRepositoryAssociationsInput, func(*codegurureviewer.ListRepositoryAssociationsOutput, bool) bool) error
	ListRepositoryAssociationsPagesWithContext(aws.Context, *codegurureviewer.ListRepositoryAssociationsInput, func(*codegurureviewer.ListRepositoryAssociationsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*codegurureviewer.ListTagsForResourceInput) (*codegurureviewer.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *codegurureviewer.ListTagsForResourceInput, ...request.Option) (*codegurureviewer.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*codegurureviewer.ListTagsForResourceInput) (*request.Request, *codegurureviewer.ListTagsForResourceOutput)

	PutRecommendationFeedback(*codegurureviewer.PutRecommendationFeedbackInput) (*codegurureviewer.PutRecommendationFeedbackOutput, error)
	PutRecommendationFeedbackWithContext(aws.Context, *codegurureviewer.PutRecommendationFeedbackInput, ...request.Option) (*codegurureviewer.PutRecommendationFeedbackOutput, error)
	PutRecommendationFeedbackRequest(*codegurureviewer.PutRecommendationFeedbackInput) (*request.Request, *codegurureviewer.PutRecommendationFeedbackOutput)

	TagResource(*codegurureviewer.TagResourceInput) (*codegurureviewer.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *codegurureviewer.TagResourceInput, ...request.Option) (*codegurureviewer.TagResourceOutput, error)
	TagResourceRequest(*codegurureviewer.TagResourceInput) (*request.Request, *codegurureviewer.TagResourceOutput)

	UntagResource(*codegurureviewer.UntagResourceInput) (*codegurureviewer.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *codegurureviewer.UntagResourceInput, ...request.Option) (*codegurureviewer.UntagResourceOutput, error)
	UntagResourceRequest(*codegurureviewer.UntagResourceInput) (*request.Request, *codegurureviewer.UntagResourceOutput)

	WaitUntilCodeReviewCompleted(*codegurureviewer.DescribeCodeReviewInput) error
	WaitUntilCodeReviewCompletedWithContext(aws.Context, *codegurureviewer.DescribeCodeReviewInput, ...request.WaiterOption) error

	WaitUntilRepositoryAssociationSucceeded(*codegurureviewer.DescribeRepositoryAssociationInput) error
	WaitUntilRepositoryAssociationSucceededWithContext(aws.Context, *codegurureviewer.DescribeRepositoryAssociationInput, ...request.WaiterOption) error
}

var _ CodeGuruReviewerAPI = (*codegurureviewer.CodeGuruReviewer)(nil)
