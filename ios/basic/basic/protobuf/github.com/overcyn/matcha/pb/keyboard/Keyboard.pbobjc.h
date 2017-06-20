// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: github.com/overcyn/matcha/pb/keyboard/keyboard.proto

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

NS_ASSUME_NONNULL_BEGIN

#pragma mark - Enum MatchaKeyboardPBType

typedef GPB_ENUM(MatchaKeyboardPBType) {
  /**
   * Value used if any message's field encounters a value that is not defined
   * by this enum. The message will also have C functions to get/set the rawValue
   * of the field.
   **/
  MatchaKeyboardPBType_GPBUnrecognizedEnumeratorValue = kGPBUnrecognizedEnumeratorValue,
  MatchaKeyboardPBType_DefaultType = 0,
  MatchaKeyboardPBType_NumberType = 1,
  MatchaKeyboardPBType_NumberPunctuationType = 2,
  MatchaKeyboardPBType_DecimalType = 3,
  MatchaKeyboardPBType_PhoneType = 4,
  MatchaKeyboardPBType_AsciiType = 5,
  MatchaKeyboardPBType_EmailType = 6,
  MatchaKeyboardPBType_URLType = 7,
  MatchaKeyboardPBType_WebSearchType = 8,
  MatchaKeyboardPBType_NamePhoneType = 9,
};

GPBEnumDescriptor *MatchaKeyboardPBType_EnumDescriptor(void);

/**
 * Checks to see if the given value is defined by the enum or was not known at
 * the time this source was generated.
 **/
BOOL MatchaKeyboardPBType_IsValidValue(int32_t value);

#pragma mark - Enum MatchaKeyboardPBAppearance

typedef GPB_ENUM(MatchaKeyboardPBAppearance) {
  /**
   * Value used if any message's field encounters a value that is not defined
   * by this enum. The message will also have C functions to get/set the rawValue
   * of the field.
   **/
  MatchaKeyboardPBAppearance_GPBUnrecognizedEnumeratorValue = kGPBUnrecognizedEnumeratorValue,
  MatchaKeyboardPBAppearance_DefaultAppearance = 0,
  MatchaKeyboardPBAppearance_LightAppearance = 1,
  MatchaKeyboardPBAppearance_DarkAppearance = 2,
};

GPBEnumDescriptor *MatchaKeyboardPBAppearance_EnumDescriptor(void);

/**
 * Checks to see if the given value is defined by the enum or was not known at
 * the time this source was generated.
 **/
BOOL MatchaKeyboardPBAppearance_IsValidValue(int32_t value);

#pragma mark - Enum MatchaKeyboardPBReturnType

typedef GPB_ENUM(MatchaKeyboardPBReturnType) {
  /**
   * Value used if any message's field encounters a value that is not defined
   * by this enum. The message will also have C functions to get/set the rawValue
   * of the field.
   **/
  MatchaKeyboardPBReturnType_GPBUnrecognizedEnumeratorValue = kGPBUnrecognizedEnumeratorValue,
  MatchaKeyboardPBReturnType_DefaultReturnType = 0,
  MatchaKeyboardPBReturnType_GoReturnType = 1,
  MatchaKeyboardPBReturnType_GoogleReturnType = 2,
  MatchaKeyboardPBReturnType_JoinReturnType = 3,
  MatchaKeyboardPBReturnType_NextReturnType = 4,
  MatchaKeyboardPBReturnType_RouteReturnType = 5,
  MatchaKeyboardPBReturnType_SearchReturnType = 6,
  MatchaKeyboardPBReturnType_SendReturnType = 7,
  MatchaKeyboardPBReturnType_YahooReturnType = 8,
  MatchaKeyboardPBReturnType_DoneReturnType = 9,
  MatchaKeyboardPBReturnType_EmergencyCallReturnType = 10,
  MatchaKeyboardPBReturnType_ContinueReturnType = 11,
};

GPBEnumDescriptor *MatchaKeyboardPBReturnType_EnumDescriptor(void);

/**
 * Checks to see if the given value is defined by the enum or was not known at
 * the time this source was generated.
 **/
BOOL MatchaKeyboardPBReturnType_IsValidValue(int32_t value);

#pragma mark - MatchaKeyboardPBKeyboardRoot

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
@interface MatchaKeyboardPBKeyboardRoot : GPBRootObject
@end

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
