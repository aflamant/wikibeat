// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeImagesRequest
type DescribeImagesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// Scopes the images by users with explicit launch permissions. Specify an AWS
	// account ID, self (the sender of the request), or all (public AMIs).
	ExecutableUsers []string `locationName:"ExecutableBy" locationNameList:"ExecutableBy" type:"list"`

	// The filters.
	//
	//    * architecture - The image architecture (i386 | x86_64).
	//
	//    * block-device-mapping.delete-on-termination - A Boolean value that indicates
	//    whether the Amazon EBS volume is deleted on instance termination.
	//
	//    * block-device-mapping.device-name - The device name specified in the
	//    block device mapping (for example, /dev/sdh or xvdh).
	//
	//    * block-device-mapping.snapshot-id - The ID of the snapshot used for the
	//    EBS volume.
	//
	//    * block-device-mapping.volume-size - The volume size of the EBS volume,
	//    in GiB.
	//
	//    * block-device-mapping.volume-type - The volume type of the EBS volume
	//    (gp2 | io1 | st1 | sc1 | standard).
	//
	//    * block-device-mapping.encrypted - A Boolean that indicates whether the
	//    EBS volume is encrypted.
	//
	//    * description - The description of the image (provided during image creation).
	//
	//    * ena-support - A Boolean that indicates whether enhanced networking with
	//    ENA is enabled.
	//
	//    * hypervisor - The hypervisor type (ovm | xen).
	//
	//    * image-id - The ID of the image.
	//
	//    * image-type - The image type (machine | kernel | ramdisk).
	//
	//    * is-public - A Boolean that indicates whether the image is public.
	//
	//    * kernel-id - The kernel ID.
	//
	//    * manifest-location - The location of the image manifest.
	//
	//    * name - The name of the AMI (provided during image creation).
	//
	//    * owner-alias - String value from an Amazon-maintained list (amazon |
	//    aws-marketplace | microsoft) of snapshot owners. Not to be confused with
	//    the user-configured AWS account alias, which is set from the IAM console.
	//
	//    * owner-id - The AWS account ID of the image owner.
	//
	//    * platform - The platform. To only list Windows-based AMIs, use windows.
	//
	//    * product-code - The product code.
	//
	//    * product-code.type - The type of the product code (devpay | marketplace).
	//
	//    * ramdisk-id - The RAM disk ID.
	//
	//    * root-device-name - The device name of the root device volume (for example,
	//    /dev/sda1).
	//
	//    * root-device-type - The type of the root device volume (ebs | instance-store).
	//
	//    * state - The state of the image (available | pending | failed).
	//
	//    * state-reason-code - The reason code for the state change.
	//
	//    * state-reason-message - The message for the state change.
	//
	//    * sriov-net-support - A value of simple indicates that enhanced networking
	//    with the Intel 82599 VF interface is enabled.
	//
	//    * tag:<key> - The key/value combination of a tag assigned to the resource.
	//    Use the tag key in the filter name and the tag value as the filter value.
	//    For example, to find all resources that have a tag with the key Owner
	//    and the value TeamA, specify tag:Owner for the filter name and TeamA for
	//    the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. Use this filter
	//    to find all resources assigned a tag with a specific key, regardless of
	//    the tag value.
	//
	//    * virtualization-type - The virtualization type (paravirtual | hvm).
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The image IDs.
	//
	// Default: Describes all images available to you.
	ImageIds []string `locationName:"ImageId" locationNameList:"ImageId" type:"list"`

	// Filters the images by the owner. Specify an AWS account ID, self (owner is
	// the sender of the request), or an AWS owner alias (valid values are amazon
	// | aws-marketplace | microsoft). Omitting this option returns all images for
	// which you have launch permissions, regardless of ownership.
	Owners []string `locationName:"Owner" locationNameList:"Owner" type:"list"`
}

// String returns the string representation
func (s DescribeImagesInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeImagesResult
type DescribeImagesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the images.
	Images []Image `locationName:"imagesSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeImagesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeImages = "DescribeImages"

// DescribeImagesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified images (AMIs, AKIs, and ARIs) available to you or
// all of the images available to you.
//
// The images available to you include public images, private images that you
// own, and private images owned by other AWS accounts for which you have explicit
// launch permissions.
//
// Recently deregistered images appear in the returned results for a short interval
// and then return empty results. After all instances that reference a deregistered
// AMI are terminated, specifying the ID of the image results in an error indicating
// that the AMI ID cannot be found.
//
//    // Example sending a request using DescribeImagesRequest.
//    req := client.DescribeImagesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeImages
func (c *Client) DescribeImagesRequest(input *DescribeImagesInput) DescribeImagesRequest {
	op := &aws.Operation{
		Name:       opDescribeImages,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeImagesInput{}
	}

	req := c.newRequest(op, input, &DescribeImagesOutput{})
	return DescribeImagesRequest{Request: req, Input: input, Copy: c.DescribeImagesRequest}
}

// DescribeImagesRequest is the request type for the
// DescribeImages API operation.
type DescribeImagesRequest struct {
	*aws.Request
	Input *DescribeImagesInput
	Copy  func(*DescribeImagesInput) DescribeImagesRequest
}

// Send marshals and sends the DescribeImages API request.
func (r DescribeImagesRequest) Send(ctx context.Context) (*DescribeImagesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeImagesResponse{
		DescribeImagesOutput: r.Request.Data.(*DescribeImagesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeImagesResponse is the response type for the
// DescribeImages API operation.
type DescribeImagesResponse struct {
	*DescribeImagesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeImages request.
func (r *DescribeImagesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
