// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: promotion/promotion_service.proto

package promotion

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_promotion_promotion_service_proto protoreflect.FileDescriptor

const file_promotion_promotion_service_proto_rawDesc = "" +
	"\n" +
	"!promotion/promotion_service.proto\x12\tpromotion\x1a\x1dpromotion/promotion_msg.proto2i\n" +
	"\x10PromotionService\x12U\n" +
	"\x0eApplyPromotion\x12 .promotion.ApplyPromotionRequest\x1a!.promotion.ApplyPromotionResponseB Z\x1eztf-backend/proto/pb/promotionb\x06proto3"

var file_promotion_promotion_service_proto_goTypes = []any{
	(*ApplyPromotionRequest)(nil),  // 0: promotion.ApplyPromotionRequest
	(*ApplyPromotionResponse)(nil), // 1: promotion.ApplyPromotionResponse
}
var file_promotion_promotion_service_proto_depIdxs = []int32{
	0, // 0: promotion.PromotionService.ApplyPromotion:input_type -> promotion.ApplyPromotionRequest
	1, // 1: promotion.PromotionService.ApplyPromotion:output_type -> promotion.ApplyPromotionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_promotion_promotion_service_proto_init() }
func file_promotion_promotion_service_proto_init() {
	if File_promotion_promotion_service_proto != nil {
		return
	}
	file_promotion_promotion_msg_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_promotion_promotion_service_proto_rawDesc), len(file_promotion_promotion_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_promotion_promotion_service_proto_goTypes,
		DependencyIndexes: file_promotion_promotion_service_proto_depIdxs,
	}.Build()
	File_promotion_promotion_service_proto = out.File
	file_promotion_promotion_service_proto_goTypes = nil
	file_promotion_promotion_service_proto_depIdxs = nil
}
