// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: proto/v1/job.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type JobServiceStartTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkflowMeta  *WorkflowMeta          `protobuf:"bytes,1,opt,name=workflow_meta,json=workflowMeta,proto3" json:"workflow_meta,omitempty"`
	JobMeta       *JobMeta               `protobuf:"bytes,2,opt,name=job_meta,json=jobMeta,proto3" json:"job_meta,omitempty"`
	Data          *StartTaskData         `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Id            string                 `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobServiceStartTaskRequest) Reset() {
	*x = JobServiceStartTaskRequest{}
	mi := &file_proto_v1_job_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobServiceStartTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobServiceStartTaskRequest) ProtoMessage() {}

func (x *JobServiceStartTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_job_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobServiceStartTaskRequest.ProtoReflect.Descriptor instead.
func (*JobServiceStartTaskRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_job_proto_rawDescGZIP(), []int{0}
}

func (x *JobServiceStartTaskRequest) GetWorkflowMeta() *WorkflowMeta {
	if x != nil {
		return x.WorkflowMeta
	}
	return nil
}

func (x *JobServiceStartTaskRequest) GetJobMeta() *JobMeta {
	if x != nil {
		return x.JobMeta
	}
	return nil
}

func (x *JobServiceStartTaskRequest) GetData() *StartTaskData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *JobServiceStartTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type JobServiceStartTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskMeta      *TaskMeta              `protobuf:"bytes,1,opt,name=task_meta,json=taskMeta,proto3" json:"task_meta,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobServiceStartTaskResponse) Reset() {
	*x = JobServiceStartTaskResponse{}
	mi := &file_proto_v1_job_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobServiceStartTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobServiceStartTaskResponse) ProtoMessage() {}

func (x *JobServiceStartTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_job_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobServiceStartTaskResponse.ProtoReflect.Descriptor instead.
func (*JobServiceStartTaskResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_job_proto_rawDescGZIP(), []int{1}
}

func (x *JobServiceStartTaskResponse) GetTaskMeta() *TaskMeta {
	if x != nil {
		return x.TaskMeta
	}
	return nil
}

type JobServiceStopTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkflowMeta  *WorkflowMeta          `protobuf:"bytes,1,opt,name=workflow_meta,json=workflowMeta,proto3" json:"workflow_meta,omitempty"`
	JobMeta       *JobMeta               `protobuf:"bytes,2,opt,name=job_meta,json=jobMeta,proto3" json:"job_meta,omitempty"`
	TaskMeta      *TaskMeta              `protobuf:"bytes,3,opt,name=task_meta,json=taskMeta,proto3" json:"task_meta,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobServiceStopTaskRequest) Reset() {
	*x = JobServiceStopTaskRequest{}
	mi := &file_proto_v1_job_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobServiceStopTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobServiceStopTaskRequest) ProtoMessage() {}

func (x *JobServiceStopTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_job_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobServiceStopTaskRequest.ProtoReflect.Descriptor instead.
func (*JobServiceStopTaskRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_job_proto_rawDescGZIP(), []int{2}
}

func (x *JobServiceStopTaskRequest) GetWorkflowMeta() *WorkflowMeta {
	if x != nil {
		return x.WorkflowMeta
	}
	return nil
}

func (x *JobServiceStopTaskRequest) GetJobMeta() *JobMeta {
	if x != nil {
		return x.JobMeta
	}
	return nil
}

func (x *JobServiceStopTaskRequest) GetTaskMeta() *TaskMeta {
	if x != nil {
		return x.TaskMeta
	}
	return nil
}

type JobServiceStopTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobServiceStopTaskResponse) Reset() {
	*x = JobServiceStopTaskResponse{}
	mi := &file_proto_v1_job_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobServiceStopTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobServiceStopTaskResponse) ProtoMessage() {}

func (x *JobServiceStopTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_job_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobServiceStopTaskResponse.ProtoReflect.Descriptor instead.
func (*JobServiceStopTaskResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_job_proto_rawDescGZIP(), []int{3}
}

type JobServiceStartStepRequest struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	WorkflowJwt        string                 `protobuf:"bytes,1,opt,name=workflow_jwt,json=workflowJwt,proto3" json:"workflow_jwt,omitempty"`
	JobMeta            *JobMeta               `protobuf:"bytes,2,opt,name=job_meta,json=jobMeta,proto3" json:"job_meta,omitempty"`
	TaskMeta           *TaskMeta              `protobuf:"bytes,3,opt,name=task_meta,json=taskMeta,proto3,oneof" json:"task_meta,omitempty"`
	Id                 string                 `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	PresignedOutputUrl string                 `protobuf:"bytes,5,opt,name=presigned_output_url,json=presignedOutputUrl,proto3" json:"presigned_output_url,omitempty"`
	// Types that are valid to be assigned to Data:
	//
	//	*JobServiceStartStepRequest_ExecData
	Data          isJobServiceStartStepRequest_Data `protobuf_oneof:"data"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobServiceStartStepRequest) Reset() {
	*x = JobServiceStartStepRequest{}
	mi := &file_proto_v1_job_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobServiceStartStepRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobServiceStartStepRequest) ProtoMessage() {}

func (x *JobServiceStartStepRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_job_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobServiceStartStepRequest.ProtoReflect.Descriptor instead.
func (*JobServiceStartStepRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_job_proto_rawDescGZIP(), []int{4}
}

func (x *JobServiceStartStepRequest) GetWorkflowJwt() string {
	if x != nil {
		return x.WorkflowJwt
	}
	return ""
}

func (x *JobServiceStartStepRequest) GetJobMeta() *JobMeta {
	if x != nil {
		return x.JobMeta
	}
	return nil
}

func (x *JobServiceStartStepRequest) GetTaskMeta() *TaskMeta {
	if x != nil {
		return x.TaskMeta
	}
	return nil
}

func (x *JobServiceStartStepRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JobServiceStartStepRequest) GetPresignedOutputUrl() string {
	if x != nil {
		return x.PresignedOutputUrl
	}
	return ""
}

func (x *JobServiceStartStepRequest) GetData() isJobServiceStartStepRequest_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *JobServiceStartStepRequest) GetExecData() *ExecInfo {
	if x != nil {
		if x, ok := x.Data.(*JobServiceStartStepRequest_ExecData); ok {
			return x.ExecData
		}
	}
	return nil
}

type isJobServiceStartStepRequest_Data interface {
	isJobServiceStartStepRequest_Data()
}

type JobServiceStartStepRequest_ExecData struct {
	ExecData *ExecInfo `protobuf:"bytes,6,opt,name=exec_data,json=execData,proto3,oneof"`
}

func (*JobServiceStartStepRequest_ExecData) isJobServiceStartStepRequest_Data() {}

type JobServiceStartStepResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Payload:
	//
	//	*JobServiceStartStepResponse_Exec
	Payload       isJobServiceStartStepResponse_Payload `protobuf_oneof:"payload"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobServiceStartStepResponse) Reset() {
	*x = JobServiceStartStepResponse{}
	mi := &file_proto_v1_job_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobServiceStartStepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobServiceStartStepResponse) ProtoMessage() {}

func (x *JobServiceStartStepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_job_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobServiceStartStepResponse.ProtoReflect.Descriptor instead.
func (*JobServiceStartStepResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_job_proto_rawDescGZIP(), []int{5}
}

func (x *JobServiceStartStepResponse) GetPayload() isJobServiceStartStepResponse_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *JobServiceStartStepResponse) GetExec() *StepExecPayload {
	if x != nil {
		if x, ok := x.Payload.(*JobServiceStartStepResponse_Exec); ok {
			return x.Exec
		}
	}
	return nil
}

type isJobServiceStartStepResponse_Payload interface {
	isJobServiceStartStepResponse_Payload()
}

type JobServiceStartStepResponse_Exec struct {
	Exec *StepExecPayload `protobuf:"bytes,1,opt,name=exec,proto3,oneof"`
}

func (*JobServiceStartStepResponse_Exec) isJobServiceStartStepResponse_Payload() {}

type JobServiceStopStepRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobServiceStopStepRequest) Reset() {
	*x = JobServiceStopStepRequest{}
	mi := &file_proto_v1_job_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobServiceStopStepRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobServiceStopStepRequest) ProtoMessage() {}

func (x *JobServiceStopStepRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_job_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobServiceStopStepRequest.ProtoReflect.Descriptor instead.
func (*JobServiceStopStepRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_job_proto_rawDescGZIP(), []int{6}
}

type JobServiceStopStepResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobServiceStopStepResponse) Reset() {
	*x = JobServiceStopStepResponse{}
	mi := &file_proto_v1_job_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobServiceStopStepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobServiceStopStepResponse) ProtoMessage() {}

func (x *JobServiceStopStepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_job_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobServiceStopStepResponse.ProtoReflect.Descriptor instead.
func (*JobServiceStopStepResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_job_proto_rawDescGZIP(), []int{7}
}

type JobServiceCreateJobVolumeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkflowMeta  *WorkflowMeta          `protobuf:"bytes,1,opt,name=workflow_meta,json=workflowMeta,proto3" json:"workflow_meta,omitempty"`
	JobMeta       *JobMeta               `protobuf:"bytes,2,opt,name=job_meta,json=jobMeta,proto3" json:"job_meta,omitempty"`
	Host          *string                `protobuf:"bytes,3,opt,name=host,proto3,oneof" json:"host,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobServiceCreateJobVolumeRequest) Reset() {
	*x = JobServiceCreateJobVolumeRequest{}
	mi := &file_proto_v1_job_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobServiceCreateJobVolumeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobServiceCreateJobVolumeRequest) ProtoMessage() {}

func (x *JobServiceCreateJobVolumeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_job_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobServiceCreateJobVolumeRequest.ProtoReflect.Descriptor instead.
func (*JobServiceCreateJobVolumeRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_job_proto_rawDescGZIP(), []int{8}
}

func (x *JobServiceCreateJobVolumeRequest) GetWorkflowMeta() *WorkflowMeta {
	if x != nil {
		return x.WorkflowMeta
	}
	return nil
}

func (x *JobServiceCreateJobVolumeRequest) GetJobMeta() *JobMeta {
	if x != nil {
		return x.JobMeta
	}
	return nil
}

func (x *JobServiceCreateJobVolumeRequest) GetHost() string {
	if x != nil && x.Host != nil {
		return *x.Host
	}
	return ""
}

type JobServiceCreateJobVolumeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Source        string                 `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobServiceCreateJobVolumeResponse) Reset() {
	*x = JobServiceCreateJobVolumeResponse{}
	mi := &file_proto_v1_job_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobServiceCreateJobVolumeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobServiceCreateJobVolumeResponse) ProtoMessage() {}

func (x *JobServiceCreateJobVolumeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_job_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobServiceCreateJobVolumeResponse.ProtoReflect.Descriptor instead.
func (*JobServiceCreateJobVolumeResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_job_proto_rawDescGZIP(), []int{9}
}

func (x *JobServiceCreateJobVolumeResponse) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type JobServicePingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobServicePingRequest) Reset() {
	*x = JobServicePingRequest{}
	mi := &file_proto_v1_job_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobServicePingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobServicePingRequest) ProtoMessage() {}

func (x *JobServicePingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_job_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobServicePingRequest.ProtoReflect.Descriptor instead.
func (*JobServicePingRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_job_proto_rawDescGZIP(), []int{10}
}

type JobServicePingResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobServicePingResponse) Reset() {
	*x = JobServicePingResponse{}
	mi := &file_proto_v1_job_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobServicePingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobServicePingResponse) ProtoMessage() {}

func (x *JobServicePingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_job_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobServicePingResponse.ProtoReflect.Descriptor instead.
func (*JobServicePingResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_job_proto_rawDescGZIP(), []int{11}
}

var File_proto_v1_job_proto protoreflect.FileDescriptor

var file_proto_v1_job_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x15,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x1a, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x74,
	0x61, 0x12, 0x2c, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4a,
	0x6f, 0x62, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4d, 0x65, 0x74, 0x61, 0x12,
	0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4e, 0x0a, 0x1b,
	0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x22, 0xb7, 0x01, 0x0a,
	0x19, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x6f, 0x70, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0d, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x07, 0x6a, 0x6f,
	0x62, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x74, 0x61,
	0x73, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x22, 0x1c, 0x0a, 0x1a, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0xae, 0x02, 0x0a, 0x1a, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f,
	0x6a, 0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x4a, 0x77, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x07, 0x6a, 0x6f, 0x62,
	0x4d, 0x65, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x01, 0x52, 0x08, 0x74,
	0x61, 0x73, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x72,
	0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x72, 0x65, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x31, 0x0a, 0x09,
	0x65, 0x78, 0x65, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x49,
	0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x08, 0x65, 0x78, 0x65, 0x63, 0x44, 0x61, 0x74, 0x61, 0x42,
	0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x59, 0x0a, 0x1b, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x65, 0x78, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x65, 0x70, 0x45, 0x78, 0x65, 0x63, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x00, 0x52,
	0x04, 0x65, 0x78, 0x65, 0x63, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x22, 0x1b, 0x0a, 0x19, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74,
	0x6f, 0x70, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1c, 0x0a,
	0x1a, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x6f, 0x70, 0x53,
	0x74, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xaf, 0x01, 0x0a, 0x20,
	0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4a, 0x6f, 0x62, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3b, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x2c, 0x0a,
	0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x22, 0x3b, 0x0a,
	0x21, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4a, 0x6f, 0x62, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x4a, 0x6f,
	0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa7, 0x04,
	0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x09,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x08, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f,
	0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x6f,
	0x70, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a,
	0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x65, 0x70, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x65, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x55, 0x0a, 0x08, 0x53, 0x74, 0x6f,
	0x70, 0x53, 0x74, 0x65, 0x70, 0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x6f, 0x70, 0x53,
	0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x74, 0x6f, 0x70, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x6a, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x12, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4a,
	0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a,
	0x6f, 0x62, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x04,
	0x50, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1_job_proto_rawDescOnce sync.Once
	file_proto_v1_job_proto_rawDescData = file_proto_v1_job_proto_rawDesc
)

func file_proto_v1_job_proto_rawDescGZIP() []byte {
	file_proto_v1_job_proto_rawDescOnce.Do(func() {
		file_proto_v1_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_job_proto_rawDescData)
	})
	return file_proto_v1_job_proto_rawDescData
}

var file_proto_v1_job_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_v1_job_proto_goTypes = []any{
	(*JobServiceStartTaskRequest)(nil),        // 0: proto.v1.JobServiceStartTaskRequest
	(*JobServiceStartTaskResponse)(nil),       // 1: proto.v1.JobServiceStartTaskResponse
	(*JobServiceStopTaskRequest)(nil),         // 2: proto.v1.JobServiceStopTaskRequest
	(*JobServiceStopTaskResponse)(nil),        // 3: proto.v1.JobServiceStopTaskResponse
	(*JobServiceStartStepRequest)(nil),        // 4: proto.v1.JobServiceStartStepRequest
	(*JobServiceStartStepResponse)(nil),       // 5: proto.v1.JobServiceStartStepResponse
	(*JobServiceStopStepRequest)(nil),         // 6: proto.v1.JobServiceStopStepRequest
	(*JobServiceStopStepResponse)(nil),        // 7: proto.v1.JobServiceStopStepResponse
	(*JobServiceCreateJobVolumeRequest)(nil),  // 8: proto.v1.JobServiceCreateJobVolumeRequest
	(*JobServiceCreateJobVolumeResponse)(nil), // 9: proto.v1.JobServiceCreateJobVolumeResponse
	(*JobServicePingRequest)(nil),             // 10: proto.v1.JobServicePingRequest
	(*JobServicePingResponse)(nil),            // 11: proto.v1.JobServicePingResponse
	(*WorkflowMeta)(nil),                      // 12: proto.v1.WorkflowMeta
	(*JobMeta)(nil),                           // 13: proto.v1.JobMeta
	(*StartTaskData)(nil),                     // 14: proto.v1.StartTaskData
	(*TaskMeta)(nil),                          // 15: proto.v1.TaskMeta
	(*ExecInfo)(nil),                          // 16: proto.v1.ExecInfo
	(*StepExecPayload)(nil),                   // 17: proto.v1.StepExecPayload
}
var file_proto_v1_job_proto_depIdxs = []int32{
	12, // 0: proto.v1.JobServiceStartTaskRequest.workflow_meta:type_name -> proto.v1.WorkflowMeta
	13, // 1: proto.v1.JobServiceStartTaskRequest.job_meta:type_name -> proto.v1.JobMeta
	14, // 2: proto.v1.JobServiceStartTaskRequest.data:type_name -> proto.v1.StartTaskData
	15, // 3: proto.v1.JobServiceStartTaskResponse.task_meta:type_name -> proto.v1.TaskMeta
	12, // 4: proto.v1.JobServiceStopTaskRequest.workflow_meta:type_name -> proto.v1.WorkflowMeta
	13, // 5: proto.v1.JobServiceStopTaskRequest.job_meta:type_name -> proto.v1.JobMeta
	15, // 6: proto.v1.JobServiceStopTaskRequest.task_meta:type_name -> proto.v1.TaskMeta
	13, // 7: proto.v1.JobServiceStartStepRequest.job_meta:type_name -> proto.v1.JobMeta
	15, // 8: proto.v1.JobServiceStartStepRequest.task_meta:type_name -> proto.v1.TaskMeta
	16, // 9: proto.v1.JobServiceStartStepRequest.exec_data:type_name -> proto.v1.ExecInfo
	17, // 10: proto.v1.JobServiceStartStepResponse.exec:type_name -> proto.v1.StepExecPayload
	12, // 11: proto.v1.JobServiceCreateJobVolumeRequest.workflow_meta:type_name -> proto.v1.WorkflowMeta
	13, // 12: proto.v1.JobServiceCreateJobVolumeRequest.job_meta:type_name -> proto.v1.JobMeta
	0,  // 13: proto.v1.JobService.StartTask:input_type -> proto.v1.JobServiceStartTaskRequest
	2,  // 14: proto.v1.JobService.StopTask:input_type -> proto.v1.JobServiceStopTaskRequest
	4,  // 15: proto.v1.JobService.StartStep:input_type -> proto.v1.JobServiceStartStepRequest
	6,  // 16: proto.v1.JobService.StopStep:input_type -> proto.v1.JobServiceStopStepRequest
	8,  // 17: proto.v1.JobService.CreateJobVolume:input_type -> proto.v1.JobServiceCreateJobVolumeRequest
	10, // 18: proto.v1.JobService.Ping:input_type -> proto.v1.JobServicePingRequest
	1,  // 19: proto.v1.JobService.StartTask:output_type -> proto.v1.JobServiceStartTaskResponse
	3,  // 20: proto.v1.JobService.StopTask:output_type -> proto.v1.JobServiceStopTaskResponse
	5,  // 21: proto.v1.JobService.StartStep:output_type -> proto.v1.JobServiceStartStepResponse
	7,  // 22: proto.v1.JobService.StopStep:output_type -> proto.v1.JobServiceStopStepResponse
	9,  // 23: proto.v1.JobService.CreateJobVolume:output_type -> proto.v1.JobServiceCreateJobVolumeResponse
	11, // 24: proto.v1.JobService.Ping:output_type -> proto.v1.JobServicePingResponse
	19, // [19:25] is the sub-list for method output_type
	13, // [13:19] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_proto_v1_job_proto_init() }
func file_proto_v1_job_proto_init() {
	if File_proto_v1_job_proto != nil {
		return
	}
	file_proto_v1_shared_proto_init()
	file_proto_v1_job_proto_msgTypes[4].OneofWrappers = []any{
		(*JobServiceStartStepRequest_ExecData)(nil),
	}
	file_proto_v1_job_proto_msgTypes[5].OneofWrappers = []any{
		(*JobServiceStartStepResponse_Exec)(nil),
	}
	file_proto_v1_job_proto_msgTypes[8].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_v1_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1_job_proto_goTypes,
		DependencyIndexes: file_proto_v1_job_proto_depIdxs,
		MessageInfos:      file_proto_v1_job_proto_msgTypes,
	}.Build()
	File_proto_v1_job_proto = out.File
	file_proto_v1_job_proto_rawDesc = nil
	file_proto_v1_job_proto_goTypes = nil
	file_proto_v1_job_proto_depIdxs = nil
}
