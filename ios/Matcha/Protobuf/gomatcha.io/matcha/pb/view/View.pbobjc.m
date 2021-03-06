// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: gomatcha.io/matcha/pb/view/view.proto

// This CPP symbol can be defined to use imports that match up to the framework
// imports needed when using CocoaPods.
#if !defined(GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS)
 #define GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS 0
#endif

#if GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS
 #import <Protobuf/GPBProtocolBuffers_RuntimeSupport.h>
#else
 #import "GPBProtocolBuffers_RuntimeSupport.h"
#endif

#if GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS
 #import <Protobuf/Any.pbobjc.h>
#else
 #import "google/protobuf/Any.pbobjc.h"
#endif

 #import "gomatcha.io/matcha/pb/view/View.pbobjc.h"
 #import "gomatcha.io/matcha/pb/paint/Paint.pbobjc.h"
// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

#pragma mark - MatchaViewPBViewRoot

@implementation MatchaViewPBViewRoot

// No extensions in the file and none of the imports (direct or indirect)
// defined extensions, so no need to generate +extensionRegistry.

@end

#pragma mark - MatchaViewPBViewRoot_FileDescriptor

static GPBFileDescriptor *MatchaViewPBViewRoot_FileDescriptor(void) {
  // This is called by +initialize so there is no need to worry
  // about thread safety of the singleton.
  static GPBFileDescriptor *descriptor = NULL;
  if (!descriptor) {
    GPB_DEBUG_CHECK_RUNTIME_VERSIONS();
    descriptor = [[GPBFileDescriptor alloc] initWithPackage:@"matcha.view"
                                                 objcPrefix:@"MatchaViewPB"
                                                     syntax:GPBFileSyntaxProto3];
  }
  return descriptor;
}

#pragma mark - MatchaViewPBBuildNode

@implementation MatchaViewPBBuildNode

@dynamic id_p;
@dynamic buildId;
@dynamic bridgeName;
@dynamic hasBridgeValue, bridgeValue;
@dynamic values, values_Count;
@dynamic childrenArray, childrenArray_Count;

typedef struct MatchaViewPBBuildNode__storage_ {
  uint32_t _has_storage_[1];
  NSString *bridgeName;
  GPBAny *bridgeValue;
  NSMutableDictionary *values;
  GPBInt64Array *childrenArray;
  int64_t id_p;
  int64_t buildId;
} MatchaViewPBBuildNode__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "id_p",
        .dataTypeSpecific.className = NULL,
        .number = MatchaViewPBBuildNode_FieldNumber_Id_p,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaViewPBBuildNode__storage_, id_p),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "buildId",
        .dataTypeSpecific.className = NULL,
        .number = MatchaViewPBBuildNode_FieldNumber_BuildId,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(MatchaViewPBBuildNode__storage_, buildId),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "bridgeName",
        .dataTypeSpecific.className = NULL,
        .number = MatchaViewPBBuildNode_FieldNumber_BridgeName,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(MatchaViewPBBuildNode__storage_, bridgeName),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "bridgeValue",
        .dataTypeSpecific.className = GPBStringifySymbol(GPBAny),
        .number = MatchaViewPBBuildNode_FieldNumber_BridgeValue,
        .hasIndex = 3,
        .offset = (uint32_t)offsetof(MatchaViewPBBuildNode__storage_, bridgeValue),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "values",
        .dataTypeSpecific.className = GPBStringifySymbol(GPBAny),
        .number = MatchaViewPBBuildNode_FieldNumber_Values,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(MatchaViewPBBuildNode__storage_, values),
        .flags = GPBFieldMapKeyString,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "childrenArray",
        .dataTypeSpecific.className = NULL,
        .number = MatchaViewPBBuildNode_FieldNumber_ChildrenArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(MatchaViewPBBuildNode__storage_, childrenArray),
        .flags = (GPBFieldFlags)(GPBFieldRepeated | GPBFieldPacked),
        .dataType = GPBDataTypeInt64,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaViewPBBuildNode class]
                                     rootClass:[MatchaViewPBViewRoot class]
                                          file:MatchaViewPBViewRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaViewPBBuildNode__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\003\002\007\000\003\n\000\004\013\000";
    [localDescriptor setupExtraTextInfo:extraTextFormatInfo];
#endif  // !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MatchaViewPBLayoutPaintNode

@implementation MatchaViewPBLayoutPaintNode

@dynamic id_p;
@dynamic layoutId;
@dynamic paintId;
@dynamic minx;
@dynamic miny;
@dynamic maxx;
@dynamic maxy;
@dynamic zIndex;
@dynamic childOrderArray, childOrderArray_Count;
@dynamic hasPaintStyle, paintStyle;

typedef struct MatchaViewPBLayoutPaintNode__storage_ {
  uint32_t _has_storage_[1];
  GPBInt64Array *childOrderArray;
  MatchaPaintPBStyle *paintStyle;
  int64_t id_p;
  int64_t layoutId;
  int64_t paintId;
  double minx;
  double miny;
  double maxx;
  double maxy;
  int64_t zIndex;
} MatchaViewPBLayoutPaintNode__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "id_p",
        .dataTypeSpecific.className = NULL,
        .number = MatchaViewPBLayoutPaintNode_FieldNumber_Id_p,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaViewPBLayoutPaintNode__storage_, id_p),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "layoutId",
        .dataTypeSpecific.className = NULL,
        .number = MatchaViewPBLayoutPaintNode_FieldNumber_LayoutId,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(MatchaViewPBLayoutPaintNode__storage_, layoutId),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "paintId",
        .dataTypeSpecific.className = NULL,
        .number = MatchaViewPBLayoutPaintNode_FieldNumber_PaintId,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(MatchaViewPBLayoutPaintNode__storage_, paintId),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "minx",
        .dataTypeSpecific.className = NULL,
        .number = MatchaViewPBLayoutPaintNode_FieldNumber_Minx,
        .hasIndex = 3,
        .offset = (uint32_t)offsetof(MatchaViewPBLayoutPaintNode__storage_, minx),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeDouble,
      },
      {
        .name = "miny",
        .dataTypeSpecific.className = NULL,
        .number = MatchaViewPBLayoutPaintNode_FieldNumber_Miny,
        .hasIndex = 4,
        .offset = (uint32_t)offsetof(MatchaViewPBLayoutPaintNode__storage_, miny),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeDouble,
      },
      {
        .name = "maxx",
        .dataTypeSpecific.className = NULL,
        .number = MatchaViewPBLayoutPaintNode_FieldNumber_Maxx,
        .hasIndex = 5,
        .offset = (uint32_t)offsetof(MatchaViewPBLayoutPaintNode__storage_, maxx),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeDouble,
      },
      {
        .name = "maxy",
        .dataTypeSpecific.className = NULL,
        .number = MatchaViewPBLayoutPaintNode_FieldNumber_Maxy,
        .hasIndex = 6,
        .offset = (uint32_t)offsetof(MatchaViewPBLayoutPaintNode__storage_, maxy),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeDouble,
      },
      {
        .name = "zIndex",
        .dataTypeSpecific.className = NULL,
        .number = MatchaViewPBLayoutPaintNode_FieldNumber_ZIndex,
        .hasIndex = 7,
        .offset = (uint32_t)offsetof(MatchaViewPBLayoutPaintNode__storage_, zIndex),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "childOrderArray",
        .dataTypeSpecific.className = NULL,
        .number = MatchaViewPBLayoutPaintNode_FieldNumber_ChildOrderArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(MatchaViewPBLayoutPaintNode__storage_, childOrderArray),
        .flags = (GPBFieldFlags)(GPBFieldRepeated | GPBFieldPacked | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "paintStyle",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaPaintPBStyle),
        .number = MatchaViewPBLayoutPaintNode_FieldNumber_PaintStyle,
        .hasIndex = 8,
        .offset = (uint32_t)offsetof(MatchaViewPBLayoutPaintNode__storage_, paintStyle),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeMessage,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaViewPBLayoutPaintNode class]
                                     rootClass:[MatchaViewPBViewRoot class]
                                          file:MatchaViewPBViewRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaViewPBLayoutPaintNode__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\005\002\010\000\003\007\000\010\006\000\t\000childOrder\000\n\n\000";
    [localDescriptor setupExtraTextInfo:extraTextFormatInfo];
#endif  // !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MatchaViewPBRoot

@implementation MatchaViewPBRoot

@dynamic layoutPaintNodes, layoutPaintNodes_Count;
@dynamic buildNodes, buildNodes_Count;
@dynamic middleware, middleware_Count;

typedef struct MatchaViewPBRoot__storage_ {
  uint32_t _has_storage_[1];
  GPBInt64ObjectDictionary *layoutPaintNodes;
  GPBInt64ObjectDictionary *buildNodes;
  NSMutableDictionary *middleware;
} MatchaViewPBRoot__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "layoutPaintNodes",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaViewPBLayoutPaintNode),
        .number = MatchaViewPBRoot_FieldNumber_LayoutPaintNodes,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(MatchaViewPBRoot__storage_, layoutPaintNodes),
        .flags = (GPBFieldFlags)(GPBFieldMapKeyInt64 | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "buildNodes",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaViewPBBuildNode),
        .number = MatchaViewPBRoot_FieldNumber_BuildNodes,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(MatchaViewPBRoot__storage_, buildNodes),
        .flags = (GPBFieldFlags)(GPBFieldMapKeyInt64 | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "middleware",
        .dataTypeSpecific.className = GPBStringifySymbol(GPBAny),
        .number = MatchaViewPBRoot_FieldNumber_Middleware,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(MatchaViewPBRoot__storage_, middleware),
        .flags = GPBFieldMapKeyString,
        .dataType = GPBDataTypeMessage,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaViewPBRoot class]
                                     rootClass:[MatchaViewPBViewRoot class]
                                          file:MatchaViewPBViewRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaViewPBRoot__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\002\002\020\000\003\n\000";
    [localDescriptor setupExtraTextInfo:extraTextFormatInfo];
#endif  // !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end


#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
