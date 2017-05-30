// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: github.com/overcyn/mochi/pb/touch.proto

// This CPP symbol can be defined to use imports that match up to the framework
// imports needed when using CocoaPods.
#if !defined(GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS)
 #define GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS 0
#endif

#if GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS
 #import <Protobuf/GPBProtocolBuffers.h>
#else
 #import "GPBProtocolBuffers.h"
#endif

#if GOOGLE_PROTOBUF_OBJC_VERSION < 30002
#error This file was generated by a newer version of protoc which is incompatible with your Protocol Buffer library sources.
#endif
#if 30002 < GOOGLE_PROTOBUF_OBJC_MIN_SUPPORTED_VERSION
#error This file was generated by an older version of protoc which is incompatible with your Protocol Buffer library sources.
#endif

// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

CF_EXTERN_C_BEGIN

@class GPBAny;
@class GPBDuration;
@class GPBTimestamp;
@class MochiPBPoint;
@class MochiPBRecognizer;

NS_ASSUME_NONNULL_BEGIN

#pragma mark - Enum MochiPBEventKind

typedef GPB_ENUM(MochiPBEventKind) {
  /**
   * Value used if any message's field encounters a value that is not defined
   * by this enum. The message will also have C functions to get/set the rawValue
   * of the field.
   **/
  MochiPBEventKind_GPBUnrecognizedEnumeratorValue = kGPBUnrecognizedEnumeratorValue,
  MochiPBEventKind_EventKindBegan = 0,
  MochiPBEventKind_EventKindChanged = 1,
  MochiPBEventKind_EventKindCancelled = 2,
  MochiPBEventKind_EventKindEnded = 3,
};

GPBEnumDescriptor *MochiPBEventKind_EnumDescriptor(void);

/**
 * Checks to see if the given value is defined by the enum or was not known at
 * the time this source was generated.
 **/
BOOL MochiPBEventKind_IsValidValue(int32_t value);

#pragma mark - MochiPBTouchRoot

/**
 * Exposes the extension registry for this file.
 *
 * The base class provides:
 * @code
 *   + (GPBExtensionRegistry *)extensionRegistry;
 * @endcode
 * which is a @c GPBExtensionRegistry that includes all the extensions defined by
 * this file and all files that it depends on.
 **/
@interface MochiPBTouchRoot : GPBRootObject
@end

#pragma mark - MochiPBRecognizer

typedef GPB_ENUM(MochiPBRecognizer_FieldNumber) {
  MochiPBRecognizer_FieldNumber_Id_p = 1,
  MochiPBRecognizer_FieldNumber_Recognizer = 3,
};

@interface MochiPBRecognizer : GPBMessage

@property(nonatomic, readwrite) int64_t id_p;

@property(nonatomic, readwrite, strong, null_resettable) GPBAny *recognizer;
/** Test to see if @c recognizer has been set. */
@property(nonatomic, readwrite) BOOL hasRecognizer;

@end

#pragma mark - MochiPBRecognizerList

typedef GPB_ENUM(MochiPBRecognizerList_FieldNumber) {
  MochiPBRecognizerList_FieldNumber_RecognizersArray = 1,
};

@interface MochiPBRecognizerList : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<MochiPBRecognizer*> *recognizersArray;
/** The number of items in @c recognizersArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger recognizersArray_Count;

@end

#pragma mark - MochiPBTapRecognizer

typedef GPB_ENUM(MochiPBTapRecognizer_FieldNumber) {
  MochiPBTapRecognizer_FieldNumber_Count = 1,
  MochiPBTapRecognizer_FieldNumber_RecognizedFunc = 2,
};

@interface MochiPBTapRecognizer : GPBMessage

@property(nonatomic, readwrite) int64_t count;

@property(nonatomic, readwrite) int64_t recognizedFunc;

@end

#pragma mark - MochiPBTapEvent

typedef GPB_ENUM(MochiPBTapEvent_FieldNumber) {
  MochiPBTapEvent_FieldNumber_Timestamp = 1,
  MochiPBTapEvent_FieldNumber_Position = 2,
};

@interface MochiPBTapEvent : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) GPBTimestamp *timestamp;
/** Test to see if @c timestamp has been set. */
@property(nonatomic, readwrite) BOOL hasTimestamp;

@property(nonatomic, readwrite, strong, null_resettable) MochiPBPoint *position;
/** Test to see if @c position has been set. */
@property(nonatomic, readwrite) BOOL hasPosition;

@end

#pragma mark - MochiPBPressRecognizer

typedef GPB_ENUM(MochiPBPressRecognizer_FieldNumber) {
  MochiPBPressRecognizer_FieldNumber_MinDuration = 1,
  MochiPBPressRecognizer_FieldNumber_FuncId = 2,
};

@interface MochiPBPressRecognizer : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) GPBDuration *minDuration;
/** Test to see if @c minDuration has been set. */
@property(nonatomic, readwrite) BOOL hasMinDuration;

@property(nonatomic, readwrite) int64_t funcId;

@end

#pragma mark - MochiPBPressEvent

typedef GPB_ENUM(MochiPBPressEvent_FieldNumber) {
  MochiPBPressEvent_FieldNumber_Timestamp = 1,
  MochiPBPressEvent_FieldNumber_Position = 2,
  MochiPBPressEvent_FieldNumber_Kind = 3,
  MochiPBPressEvent_FieldNumber_Duration = 4,
};

@interface MochiPBPressEvent : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) GPBTimestamp *timestamp;
/** Test to see if @c timestamp has been set. */
@property(nonatomic, readwrite) BOOL hasTimestamp;

@property(nonatomic, readwrite, strong, null_resettable) MochiPBPoint *position;
/** Test to see if @c position has been set. */
@property(nonatomic, readwrite) BOOL hasPosition;

@property(nonatomic, readwrite) MochiPBEventKind kind;

@property(nonatomic, readwrite, strong, null_resettable) GPBDuration *duration;
/** Test to see if @c duration has been set. */
@property(nonatomic, readwrite) BOOL hasDuration;

@end

/**
 * Fetches the raw value of a @c MochiPBPressEvent's @c kind property, even
 * if the value was not defined by the enum at the time the code was generated.
 **/
int32_t MochiPBPressEvent_Kind_RawValue(MochiPBPressEvent *message);
/**
 * Sets the raw value of an @c MochiPBPressEvent's @c kind property, allowing
 * it to be set to a value that was not defined by the enum at the time the code
 * was generated.
 **/
void SetMochiPBPressEvent_Kind_RawValue(MochiPBPressEvent *message, int32_t value);

#pragma mark - MochiPBPanRecognizer

@interface MochiPBPanRecognizer : GPBMessage

@end

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
