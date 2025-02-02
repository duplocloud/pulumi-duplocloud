// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package duplocloud

import (
	"context"
	"reflect"

	"errors"
	"github.com/duplocloud/pulumi-duplocloud/sdk/go/duplocloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `AwsSqsQueue` manages a SQS queue in Duplo.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/duplocloud/pulumi-duplocloud/sdk/go/duplocloud"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myapp, err := duplocloud.NewTenant(ctx, "myapp", &duplocloud.TenantArgs{
//				AccountName: pulumi.String("myapp"),
//				PlanId:      pulumi.String("default"),
//			})
//			if err != nil {
//				return err
//			}
//			sqsQueue, err := duplocloud.NewAwsSqsQueue(ctx, "sqs_queue", &duplocloud.AwsSqsQueueArgs{
//				TenantId:                  myapp.TenantId,
//				Name:                      pulumi.String("duplo_queue"),
//				FifoQueue:                 pulumi.Bool(true),
//				MessageRetentionSeconds:   pulumi.Int(345600),
//				VisibilityTimeoutSeconds:  pulumi.Int(30),
//				ContentBasedDeduplication: pulumi.Bool(true),
//				DelaySeconds:              pulumi.Int(10),
//			})
//			if err != nil {
//				return err
//			}
//			// SQS queue with dead letter queue configuration
//			_, err = duplocloud.NewAwsSqsQueue(ctx, "sqs_queue_with_dlq", &duplocloud.AwsSqsQueueArgs{
//				TenantId:                  myapp.TenantId,
//				Name:                      pulumi.String("duplo_queue"),
//				FifoQueue:                 pulumi.Bool(true),
//				MessageRetentionSeconds:   pulumi.Int(345600),
//				VisibilityTimeoutSeconds:  pulumi.Int(30),
//				ContentBasedDeduplication: pulumi.Bool(true),
//				DelaySeconds:              pulumi.Int(10),
//				DeadLetterQueueConfiguration: &duplocloud.AwsSqsQueueDeadLetterQueueConfigurationArgs{
//					TargetSqsDlqName:          sqsQueue.Fullname,
//					MaxMessageReceiveAttempts: pulumi.Int(5),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Example: Importing an existing AWS SQS Queue
//
//   - *TENANT_ID* is the tenant GUID
//
//   - *URL* The URL for the created Amazon SQS queue.
//
// #
//
// ```sh
// $ pulumi import duplocloud:index/awsSqsQueue:AwsSqsQueue sqs_queue *TENANT_ID*/*URL*
// ```
type AwsSqsQueue struct {
	pulumi.CustomResourceState

	// The ARN of the SQS queue.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Enables content-based deduplication for FIFO queues.
	ContentBasedDeduplication pulumi.BoolOutput `pulumi:"contentBasedDeduplication"`
	// SQS configuration for the SQS resource
	DeadLetterQueueConfiguration AwsSqsQueueDeadLetterQueueConfigurationPtrOutput `pulumi:"deadLetterQueueConfiguration"`
	// Specifies whether message deduplication occurs at the message group or queue level. Valid values are `messageGroup` and `queue`.
	DeduplicationScope pulumi.StringOutput `pulumi:"deduplicationScope"`
	// Postpone the delivery of new messages to consumers for a number of seconds seconds range [0-900]
	DelaySeconds pulumi.IntOutput `pulumi:"delaySeconds"`
	// Boolean designating a FIFO queue. If not set, it defaults to `false` making it standard.
	FifoQueue pulumi.BoolPtrOutput `pulumi:"fifoQueue"`
	// Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are `perQueue` (default) and `perMessageGroupId`.
	FifoThroughputLimit pulumi.StringOutput `pulumi:"fifoThroughputLimit"`
	// The full name of the SQS queue.
	Fullname pulumi.StringOutput `pulumi:"fullname"`
	// The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days).
	MessageRetentionSeconds pulumi.IntOutput `pulumi:"messageRetentionSeconds"`
	// The name of the queue. Queue names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and have up to 80 characters long which is inclusive of duplo prefix (duploservices-{tenant_name}-) appended by the system.
	Name pulumi.StringOutput `pulumi:"name"`
	// The GUID of the tenant that the SQS queue will be created in.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The URL for the created Amazon SQS queue.
	Url pulumi.StringOutput `pulumi:"url"`
	// The visibility timeout for the queue. An integer from 0 to 43200 (12 hours).
	VisibilityTimeoutSeconds pulumi.IntOutput `pulumi:"visibilityTimeoutSeconds"`
}

// NewAwsSqsQueue registers a new resource with the given unique name, arguments, and options.
func NewAwsSqsQueue(ctx *pulumi.Context,
	name string, args *AwsSqsQueueArgs, opts ...pulumi.ResourceOption) (*AwsSqsQueue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AwsSqsQueue
	err := ctx.RegisterResource("duplocloud:index/awsSqsQueue:AwsSqsQueue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAwsSqsQueue gets an existing AwsSqsQueue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAwsSqsQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AwsSqsQueueState, opts ...pulumi.ResourceOption) (*AwsSqsQueue, error) {
	var resource AwsSqsQueue
	err := ctx.ReadResource("duplocloud:index/awsSqsQueue:AwsSqsQueue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AwsSqsQueue resources.
type awsSqsQueueState struct {
	// The ARN of the SQS queue.
	Arn *string `pulumi:"arn"`
	// Enables content-based deduplication for FIFO queues.
	ContentBasedDeduplication *bool `pulumi:"contentBasedDeduplication"`
	// SQS configuration for the SQS resource
	DeadLetterQueueConfiguration *AwsSqsQueueDeadLetterQueueConfiguration `pulumi:"deadLetterQueueConfiguration"`
	// Specifies whether message deduplication occurs at the message group or queue level. Valid values are `messageGroup` and `queue`.
	DeduplicationScope *string `pulumi:"deduplicationScope"`
	// Postpone the delivery of new messages to consumers for a number of seconds seconds range [0-900]
	DelaySeconds *int `pulumi:"delaySeconds"`
	// Boolean designating a FIFO queue. If not set, it defaults to `false` making it standard.
	FifoQueue *bool `pulumi:"fifoQueue"`
	// Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are `perQueue` (default) and `perMessageGroupId`.
	FifoThroughputLimit *string `pulumi:"fifoThroughputLimit"`
	// The full name of the SQS queue.
	Fullname *string `pulumi:"fullname"`
	// The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days).
	MessageRetentionSeconds *int `pulumi:"messageRetentionSeconds"`
	// The name of the queue. Queue names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and have up to 80 characters long which is inclusive of duplo prefix (duploservices-{tenant_name}-) appended by the system.
	Name *string `pulumi:"name"`
	// The GUID of the tenant that the SQS queue will be created in.
	TenantId *string `pulumi:"tenantId"`
	// The URL for the created Amazon SQS queue.
	Url *string `pulumi:"url"`
	// The visibility timeout for the queue. An integer from 0 to 43200 (12 hours).
	VisibilityTimeoutSeconds *int `pulumi:"visibilityTimeoutSeconds"`
}

type AwsSqsQueueState struct {
	// The ARN of the SQS queue.
	Arn pulumi.StringPtrInput
	// Enables content-based deduplication for FIFO queues.
	ContentBasedDeduplication pulumi.BoolPtrInput
	// SQS configuration for the SQS resource
	DeadLetterQueueConfiguration AwsSqsQueueDeadLetterQueueConfigurationPtrInput
	// Specifies whether message deduplication occurs at the message group or queue level. Valid values are `messageGroup` and `queue`.
	DeduplicationScope pulumi.StringPtrInput
	// Postpone the delivery of new messages to consumers for a number of seconds seconds range [0-900]
	DelaySeconds pulumi.IntPtrInput
	// Boolean designating a FIFO queue. If not set, it defaults to `false` making it standard.
	FifoQueue pulumi.BoolPtrInput
	// Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are `perQueue` (default) and `perMessageGroupId`.
	FifoThroughputLimit pulumi.StringPtrInput
	// The full name of the SQS queue.
	Fullname pulumi.StringPtrInput
	// The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days).
	MessageRetentionSeconds pulumi.IntPtrInput
	// The name of the queue. Queue names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and have up to 80 characters long which is inclusive of duplo prefix (duploservices-{tenant_name}-) appended by the system.
	Name pulumi.StringPtrInput
	// The GUID of the tenant that the SQS queue will be created in.
	TenantId pulumi.StringPtrInput
	// The URL for the created Amazon SQS queue.
	Url pulumi.StringPtrInput
	// The visibility timeout for the queue. An integer from 0 to 43200 (12 hours).
	VisibilityTimeoutSeconds pulumi.IntPtrInput
}

func (AwsSqsQueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*awsSqsQueueState)(nil)).Elem()
}

type awsSqsQueueArgs struct {
	// Enables content-based deduplication for FIFO queues.
	ContentBasedDeduplication *bool `pulumi:"contentBasedDeduplication"`
	// SQS configuration for the SQS resource
	DeadLetterQueueConfiguration *AwsSqsQueueDeadLetterQueueConfiguration `pulumi:"deadLetterQueueConfiguration"`
	// Specifies whether message deduplication occurs at the message group or queue level. Valid values are `messageGroup` and `queue`.
	DeduplicationScope *string `pulumi:"deduplicationScope"`
	// Postpone the delivery of new messages to consumers for a number of seconds seconds range [0-900]
	DelaySeconds *int `pulumi:"delaySeconds"`
	// Boolean designating a FIFO queue. If not set, it defaults to `false` making it standard.
	FifoQueue *bool `pulumi:"fifoQueue"`
	// Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are `perQueue` (default) and `perMessageGroupId`.
	FifoThroughputLimit *string `pulumi:"fifoThroughputLimit"`
	// The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days).
	MessageRetentionSeconds *int `pulumi:"messageRetentionSeconds"`
	// The name of the queue. Queue names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and have up to 80 characters long which is inclusive of duplo prefix (duploservices-{tenant_name}-) appended by the system.
	Name *string `pulumi:"name"`
	// The GUID of the tenant that the SQS queue will be created in.
	TenantId string `pulumi:"tenantId"`
	// The visibility timeout for the queue. An integer from 0 to 43200 (12 hours).
	VisibilityTimeoutSeconds *int `pulumi:"visibilityTimeoutSeconds"`
}

// The set of arguments for constructing a AwsSqsQueue resource.
type AwsSqsQueueArgs struct {
	// Enables content-based deduplication for FIFO queues.
	ContentBasedDeduplication pulumi.BoolPtrInput
	// SQS configuration for the SQS resource
	DeadLetterQueueConfiguration AwsSqsQueueDeadLetterQueueConfigurationPtrInput
	// Specifies whether message deduplication occurs at the message group or queue level. Valid values are `messageGroup` and `queue`.
	DeduplicationScope pulumi.StringPtrInput
	// Postpone the delivery of new messages to consumers for a number of seconds seconds range [0-900]
	DelaySeconds pulumi.IntPtrInput
	// Boolean designating a FIFO queue. If not set, it defaults to `false` making it standard.
	FifoQueue pulumi.BoolPtrInput
	// Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are `perQueue` (default) and `perMessageGroupId`.
	FifoThroughputLimit pulumi.StringPtrInput
	// The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days).
	MessageRetentionSeconds pulumi.IntPtrInput
	// The name of the queue. Queue names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and have up to 80 characters long which is inclusive of duplo prefix (duploservices-{tenant_name}-) appended by the system.
	Name pulumi.StringPtrInput
	// The GUID of the tenant that the SQS queue will be created in.
	TenantId pulumi.StringInput
	// The visibility timeout for the queue. An integer from 0 to 43200 (12 hours).
	VisibilityTimeoutSeconds pulumi.IntPtrInput
}

func (AwsSqsQueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*awsSqsQueueArgs)(nil)).Elem()
}

type AwsSqsQueueInput interface {
	pulumi.Input

	ToAwsSqsQueueOutput() AwsSqsQueueOutput
	ToAwsSqsQueueOutputWithContext(ctx context.Context) AwsSqsQueueOutput
}

func (*AwsSqsQueue) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsSqsQueue)(nil)).Elem()
}

func (i *AwsSqsQueue) ToAwsSqsQueueOutput() AwsSqsQueueOutput {
	return i.ToAwsSqsQueueOutputWithContext(context.Background())
}

func (i *AwsSqsQueue) ToAwsSqsQueueOutputWithContext(ctx context.Context) AwsSqsQueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsSqsQueueOutput)
}

// AwsSqsQueueArrayInput is an input type that accepts AwsSqsQueueArray and AwsSqsQueueArrayOutput values.
// You can construct a concrete instance of `AwsSqsQueueArrayInput` via:
//
//	AwsSqsQueueArray{ AwsSqsQueueArgs{...} }
type AwsSqsQueueArrayInput interface {
	pulumi.Input

	ToAwsSqsQueueArrayOutput() AwsSqsQueueArrayOutput
	ToAwsSqsQueueArrayOutputWithContext(context.Context) AwsSqsQueueArrayOutput
}

type AwsSqsQueueArray []AwsSqsQueueInput

func (AwsSqsQueueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AwsSqsQueue)(nil)).Elem()
}

func (i AwsSqsQueueArray) ToAwsSqsQueueArrayOutput() AwsSqsQueueArrayOutput {
	return i.ToAwsSqsQueueArrayOutputWithContext(context.Background())
}

func (i AwsSqsQueueArray) ToAwsSqsQueueArrayOutputWithContext(ctx context.Context) AwsSqsQueueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsSqsQueueArrayOutput)
}

// AwsSqsQueueMapInput is an input type that accepts AwsSqsQueueMap and AwsSqsQueueMapOutput values.
// You can construct a concrete instance of `AwsSqsQueueMapInput` via:
//
//	AwsSqsQueueMap{ "key": AwsSqsQueueArgs{...} }
type AwsSqsQueueMapInput interface {
	pulumi.Input

	ToAwsSqsQueueMapOutput() AwsSqsQueueMapOutput
	ToAwsSqsQueueMapOutputWithContext(context.Context) AwsSqsQueueMapOutput
}

type AwsSqsQueueMap map[string]AwsSqsQueueInput

func (AwsSqsQueueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AwsSqsQueue)(nil)).Elem()
}

func (i AwsSqsQueueMap) ToAwsSqsQueueMapOutput() AwsSqsQueueMapOutput {
	return i.ToAwsSqsQueueMapOutputWithContext(context.Background())
}

func (i AwsSqsQueueMap) ToAwsSqsQueueMapOutputWithContext(ctx context.Context) AwsSqsQueueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsSqsQueueMapOutput)
}

type AwsSqsQueueOutput struct{ *pulumi.OutputState }

func (AwsSqsQueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsSqsQueue)(nil)).Elem()
}

func (o AwsSqsQueueOutput) ToAwsSqsQueueOutput() AwsSqsQueueOutput {
	return o
}

func (o AwsSqsQueueOutput) ToAwsSqsQueueOutputWithContext(ctx context.Context) AwsSqsQueueOutput {
	return o
}

// The ARN of the SQS queue.
func (o AwsSqsQueueOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsSqsQueue) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Enables content-based deduplication for FIFO queues.
func (o AwsSqsQueueOutput) ContentBasedDeduplication() pulumi.BoolOutput {
	return o.ApplyT(func(v *AwsSqsQueue) pulumi.BoolOutput { return v.ContentBasedDeduplication }).(pulumi.BoolOutput)
}

// SQS configuration for the SQS resource
func (o AwsSqsQueueOutput) DeadLetterQueueConfiguration() AwsSqsQueueDeadLetterQueueConfigurationPtrOutput {
	return o.ApplyT(func(v *AwsSqsQueue) AwsSqsQueueDeadLetterQueueConfigurationPtrOutput {
		return v.DeadLetterQueueConfiguration
	}).(AwsSqsQueueDeadLetterQueueConfigurationPtrOutput)
}

// Specifies whether message deduplication occurs at the message group or queue level. Valid values are `messageGroup` and `queue`.
func (o AwsSqsQueueOutput) DeduplicationScope() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsSqsQueue) pulumi.StringOutput { return v.DeduplicationScope }).(pulumi.StringOutput)
}

// Postpone the delivery of new messages to consumers for a number of seconds seconds range [0-900]
func (o AwsSqsQueueOutput) DelaySeconds() pulumi.IntOutput {
	return o.ApplyT(func(v *AwsSqsQueue) pulumi.IntOutput { return v.DelaySeconds }).(pulumi.IntOutput)
}

// Boolean designating a FIFO queue. If not set, it defaults to `false` making it standard.
func (o AwsSqsQueueOutput) FifoQueue() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AwsSqsQueue) pulumi.BoolPtrOutput { return v.FifoQueue }).(pulumi.BoolPtrOutput)
}

// Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are `perQueue` (default) and `perMessageGroupId`.
func (o AwsSqsQueueOutput) FifoThroughputLimit() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsSqsQueue) pulumi.StringOutput { return v.FifoThroughputLimit }).(pulumi.StringOutput)
}

// The full name of the SQS queue.
func (o AwsSqsQueueOutput) Fullname() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsSqsQueue) pulumi.StringOutput { return v.Fullname }).(pulumi.StringOutput)
}

// The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days).
func (o AwsSqsQueueOutput) MessageRetentionSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v *AwsSqsQueue) pulumi.IntOutput { return v.MessageRetentionSeconds }).(pulumi.IntOutput)
}

// The name of the queue. Queue names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and have up to 80 characters long which is inclusive of duplo prefix (duploservices-{tenant_name}-) appended by the system.
func (o AwsSqsQueueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsSqsQueue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The GUID of the tenant that the SQS queue will be created in.
func (o AwsSqsQueueOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsSqsQueue) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The URL for the created Amazon SQS queue.
func (o AwsSqsQueueOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsSqsQueue) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The visibility timeout for the queue. An integer from 0 to 43200 (12 hours).
func (o AwsSqsQueueOutput) VisibilityTimeoutSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v *AwsSqsQueue) pulumi.IntOutput { return v.VisibilityTimeoutSeconds }).(pulumi.IntOutput)
}

type AwsSqsQueueArrayOutput struct{ *pulumi.OutputState }

func (AwsSqsQueueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AwsSqsQueue)(nil)).Elem()
}

func (o AwsSqsQueueArrayOutput) ToAwsSqsQueueArrayOutput() AwsSqsQueueArrayOutput {
	return o
}

func (o AwsSqsQueueArrayOutput) ToAwsSqsQueueArrayOutputWithContext(ctx context.Context) AwsSqsQueueArrayOutput {
	return o
}

func (o AwsSqsQueueArrayOutput) Index(i pulumi.IntInput) AwsSqsQueueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AwsSqsQueue {
		return vs[0].([]*AwsSqsQueue)[vs[1].(int)]
	}).(AwsSqsQueueOutput)
}

type AwsSqsQueueMapOutput struct{ *pulumi.OutputState }

func (AwsSqsQueueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AwsSqsQueue)(nil)).Elem()
}

func (o AwsSqsQueueMapOutput) ToAwsSqsQueueMapOutput() AwsSqsQueueMapOutput {
	return o
}

func (o AwsSqsQueueMapOutput) ToAwsSqsQueueMapOutputWithContext(ctx context.Context) AwsSqsQueueMapOutput {
	return o
}

func (o AwsSqsQueueMapOutput) MapIndex(k pulumi.StringInput) AwsSqsQueueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AwsSqsQueue {
		return vs[0].(map[string]*AwsSqsQueue)[vs[1].(string)]
	}).(AwsSqsQueueOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AwsSqsQueueInput)(nil)).Elem(), &AwsSqsQueue{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsSqsQueueArrayInput)(nil)).Elem(), AwsSqsQueueArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsSqsQueueMapInput)(nil)).Elem(), AwsSqsQueueMap{})
	pulumi.RegisterOutputType(AwsSqsQueueOutput{})
	pulumi.RegisterOutputType(AwsSqsQueueArrayOutput{})
	pulumi.RegisterOutputType(AwsSqsQueueMapOutput{})
}
