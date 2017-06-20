// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: github.com/overcyn/matcha/pb/view/tabscreen/tabscreen.proto

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

 #import "github.com/overcyn/matcha/pb/view/tabscreen/Tabscreen.pbobjc.h"
 #import "github.com/overcyn/matcha/pb/Image.pbobjc.h"
// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

#pragma mark - MatchaTabScreenPBTabscreenRoot

@implementation MatchaTabScreenPBTabscreenRoot

// No extensions in the file and none of the imports (direct or indirect)
// defined extensions, so no need to generate +extensionRegistry.

@end

#pragma mark - MatchaTabScreenPBTabscreenRoot_FileDescriptor

static GPBFileDescriptor *MatchaTabScreenPBTabscreenRoot_FileDescriptor(void) {
  // This is called by +initialize so there is no need to worry
  // about thread safety of the singleton.
  static GPBFileDescriptor *descriptor = NULL;
  if (!descriptor) {
    GPB_DEBUG_CHECK_RUNTIME_VERSIONS();
    descriptor = [[GPBFileDescriptor alloc] initWithPackage:@"matcha.view.tabscreen"
                                                 objcPrefix:@"MatchaTabScreenPB"
                                                     syntax:GPBFileSyntaxProto3];
  }
  return descriptor;
}

#pragma mark - MatchaTabScreenPBChildView

@implementation MatchaTabScreenPBChildView

@dynamic id_p;
@dynamic title;
@dynamic hasIcon, icon;
@dynamic hasSelectedIcon, selectedIcon;
@dynamic badge;

typedef struct MatchaTabScreenPBChildView__storage_ {
  uint32_t _has_storage_[1];
  NSString *title;
  MatchaPBImageOrResource *icon;
  MatchaPBImageOrResource *selectedIcon;
  NSString *badge;
  int64_t id_p;
} MatchaTabScreenPBChildView__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "id_p",
        .dataTypeSpecific.className = NULL,
        .number = MatchaTabScreenPBChildView_FieldNumber_Id_p,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaTabScreenPBChildView__storage_, id_p),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "title",
        .dataTypeSpecific.className = NULL,
        .number = MatchaTabScreenPBChildView_FieldNumber_Title,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(MatchaTabScreenPBChildView__storage_, title),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "icon",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaPBImageOrResource),
        .number = MatchaTabScreenPBChildView_FieldNumber_Icon,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(MatchaTabScreenPBChildView__storage_, icon),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "selectedIcon",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaPBImageOrResource),
        .number = MatchaTabScreenPBChildView_FieldNumber_SelectedIcon,
        .hasIndex = 3,
        .offset = (uint32_t)offsetof(MatchaTabScreenPBChildView__storage_, selectedIcon),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "badge",
        .dataTypeSpecific.className = NULL,
        .number = MatchaTabScreenPBChildView_FieldNumber_Badge,
        .hasIndex = 4,
        .offset = (uint32_t)offsetof(MatchaTabScreenPBChildView__storage_, badge),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaTabScreenPBChildView class]
                                     rootClass:[MatchaTabScreenPBTabscreenRoot class]
                                          file:MatchaTabScreenPBTabscreenRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaTabScreenPBChildView__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\001\004\014\000";
    [localDescriptor setupExtraTextInfo:extraTextFormatInfo];
#endif  // !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MatchaTabScreenPBView

@implementation MatchaTabScreenPBView

@dynamic screensArray, screensArray_Count;
@dynamic selectedIndex;

typedef struct MatchaTabScreenPBView__storage_ {
  uint32_t _has_storage_[1];
  NSMutableArray *screensArray;
  int64_t selectedIndex;
} MatchaTabScreenPBView__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "screensArray",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaTabScreenPBChildView),
        .number = MatchaTabScreenPBView_FieldNumber_ScreensArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(MatchaTabScreenPBView__storage_, screensArray),
        .flags = GPBFieldRepeated,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "selectedIndex",
        .dataTypeSpecific.className = NULL,
        .number = MatchaTabScreenPBView_FieldNumber_SelectedIndex,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaTabScreenPBView__storage_, selectedIndex),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaTabScreenPBView class]
                                     rootClass:[MatchaTabScreenPBTabscreenRoot class]
                                          file:MatchaTabScreenPBTabscreenRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaTabScreenPBView__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\001\002\r\000";
    [localDescriptor setupExtraTextInfo:extraTextFormatInfo];
#endif  // !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MatchaTabScreenPBEvent

@implementation MatchaTabScreenPBEvent

@dynamic selectedIndex;

typedef struct MatchaTabScreenPBEvent__storage_ {
  uint32_t _has_storage_[1];
  int64_t selectedIndex;
} MatchaTabScreenPBEvent__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "selectedIndex",
        .dataTypeSpecific.className = NULL,
        .number = MatchaTabScreenPBEvent_FieldNumber_SelectedIndex,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaTabScreenPBEvent__storage_, selectedIndex),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaTabScreenPBEvent class]
                                     rootClass:[MatchaTabScreenPBTabscreenRoot class]
                                          file:MatchaTabScreenPBTabscreenRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaTabScreenPBEvent__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\001\001\r\000";
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
