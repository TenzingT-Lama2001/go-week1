// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: blog_service.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_blog_service_proto protoreflect.FileDescriptor

var file_blog_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x13, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xaf, 0x02,
	0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x00, 0x12,
	0x2f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x15, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x00,
	0x12, 0x2d, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x12,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x2f, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x15, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x15,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x2e, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_blog_service_proto_goTypes = []interface{}{
	(*Empty)(nil),              // 0: pb.Empty
	(*GetPostRequest)(nil),     // 1: pb.GetPostRequest
	(*CreatePostRequest)(nil),  // 2: pb.CreatePostRequest
	(*UpdatePostRequest)(nil),  // 3: pb.UpdatePostRequest
	(*PostExistsRequest)(nil),  // 4: pb.PostExistsRequest
	(*PostList)(nil),           // 5: pb.PostList
	(*Post)(nil),               // 6: pb.Post
	(*PostExistsResponse)(nil), // 7: pb.PostExistsResponse
}
var file_blog_service_proto_depIdxs = []int32{
	0, // 0: pb.BlogService.GetPosts:input_type -> pb.Empty
	1, // 1: pb.BlogService.GetPost:input_type -> pb.GetPostRequest
	2, // 2: pb.BlogService.CreatePost:input_type -> pb.CreatePostRequest
	1, // 3: pb.BlogService.DeletePost:input_type -> pb.GetPostRequest
	3, // 4: pb.BlogService.UpdatePost:input_type -> pb.UpdatePostRequest
	4, // 5: pb.BlogService.PostExists:input_type -> pb.PostExistsRequest
	5, // 6: pb.BlogService.GetPosts:output_type -> pb.PostList
	6, // 7: pb.BlogService.GetPost:output_type -> pb.Post
	6, // 8: pb.BlogService.CreatePost:output_type -> pb.Post
	0, // 9: pb.BlogService.DeletePost:output_type -> pb.Empty
	6, // 10: pb.BlogService.UpdatePost:output_type -> pb.Post
	7, // 11: pb.BlogService.PostExists:output_type -> pb.PostExistsResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_blog_service_proto_init() }
func file_blog_service_proto_init() {
	if File_blog_service_proto != nil {
		return
	}
	file_blog_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blog_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blog_service_proto_goTypes,
		DependencyIndexes: file_blog_service_proto_depIdxs,
	}.Build()
	File_blog_service_proto = out.File
	file_blog_service_proto_rawDesc = nil
	file_blog_service_proto_goTypes = nil
	file_blog_service_proto_depIdxs = nil
}