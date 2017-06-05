// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: github.com/overcyn/mochi/pb/text.proto

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

@class MochiPBColor;
@class MochiPBFont;
@class MochiPBPoint;
@class MochiPBText;
@class MochiPBTextStyle;

NS_ASSUME_NONNULL_BEGIN

#pragma mark - Enum MochiPBTextAlignment

typedef GPB_ENUM(MochiPBTextAlignment) {
  /**
   * Value used if any message's field encounters a value that is not defined
   * by this enum. The message will also have C functions to get/set the rawValue
   * of the field.
   **/
  MochiPBTextAlignment_GPBUnrecognizedEnumeratorValue = kGPBUnrecognizedEnumeratorValue,
  MochiPBTextAlignment_TextAlignmentLeft = 0,
  MochiPBTextAlignment_TextAlignmentRight = 1,
  MochiPBTextAlignment_TextAlignmentCenter = 2,
  MochiPBTextAlignment_TextAlignmentJustified = 3,
};

GPBEnumDescriptor *MochiPBTextAlignment_EnumDescriptor(void);

/**
 * Checks to see if the given value is defined by the enum or was not known at
 * the time this source was generated.
 **/
BOOL MochiPBTextAlignment_IsValidValue(int32_t value);

#pragma mark - Enum MochiPBStrikethroughStyle

typedef GPB_ENUM(MochiPBStrikethroughStyle) {
  /**
   * Value used if any message's field encounters a value that is not defined
   * by this enum. The message will also have C functions to get/set the rawValue
   * of the field.
   **/
  MochiPBStrikethroughStyle_GPBUnrecognizedEnumeratorValue = kGPBUnrecognizedEnumeratorValue,
  MochiPBStrikethroughStyle_StrikethroughStyleNone = 0,
  MochiPBStrikethroughStyle_StrikethroughStyleSingle = 1,
  MochiPBStrikethroughStyle_StrikethroughStyleDouble = 2,
  MochiPBStrikethroughStyle_StrikethroughStyleThick = 3,
  MochiPBStrikethroughStyle_StrikethroughStyleDotted = 4,
  MochiPBStrikethroughStyle_StrikethroughStyleDashed = 5,
};

GPBEnumDescriptor *MochiPBStrikethroughStyle_EnumDescriptor(void);

/**
 * Checks to see if the given value is defined by the enum or was not known at
 * the time this source was generated.
 **/
BOOL MochiPBStrikethroughStyle_IsValidValue(int32_t value);

#pragma mark - Enum MochiPBUnderlineStyle

typedef GPB_ENUM(MochiPBUnderlineStyle) {
  /**
   * Value used if any message's field encounters a value that is not defined
   * by this enum. The message will also have C functions to get/set the rawValue
   * of the field.
   **/
  MochiPBUnderlineStyle_GPBUnrecognizedEnumeratorValue = kGPBUnrecognizedEnumeratorValue,
  MochiPBUnderlineStyle_UndrelineStyleNone = 0,
  MochiPBUnderlineStyle_UndrelineStyleSingle = 1,
  MochiPBUnderlineStyle_UndrelineStyleDouble = 2,
  MochiPBUnderlineStyle_UndrelineStyleThick = 3,
  MochiPBUnderlineStyle_UndrelineStyleDotted = 4,
  MochiPBUnderlineStyle_UndrelineStyleDashed = 5,
};

GPBEnumDescriptor *MochiPBUnderlineStyle_EnumDescriptor(void);

/**
 * Checks to see if the given value is defined by the enum or was not known at
 * the time this source was generated.
 **/
BOOL MochiPBUnderlineStyle_IsValidValue(int32_t value);

#pragma mark - Enum MochiPBTextWrap

typedef GPB_ENUM(MochiPBTextWrap) {
  /**
   * Value used if any message's field encounters a value that is not defined
   * by this enum. The message will also have C functions to get/set the rawValue
   * of the field.
   **/
  MochiPBTextWrap_GPBUnrecognizedEnumeratorValue = kGPBUnrecognizedEnumeratorValue,
  MochiPBTextWrap_TextWrapNone = 0,
  MochiPBTextWrap_TextWrapWord = 1,
  MochiPBTextWrap_TextWrapCharacter = 2,
};

GPBEnumDescriptor *MochiPBTextWrap_EnumDescriptor(void);

/**
 * Checks to see if the given value is defined by the enum or was not known at
 * the time this source was generated.
 **/
BOOL MochiPBTextWrap_IsValidValue(int32_t value);

#pragma mark - Enum MochiPBTruncation

typedef GPB_ENUM(MochiPBTruncation) {
  /**
   * Value used if any message's field encounters a value that is not defined
   * by this enum. The message will also have C functions to get/set the rawValue
   * of the field.
   **/
  MochiPBTruncation_GPBUnrecognizedEnumeratorValue = kGPBUnrecognizedEnumeratorValue,
  MochiPBTruncation_TruncationNone = 0,
  MochiPBTruncation_TruncationStart = 1,
  MochiPBTruncation_TruncationMiddle = 2,
  MochiPBTruncation_TruncationEnd = 3,
};

GPBEnumDescriptor *MochiPBTruncation_EnumDescriptor(void);

/**
 * Checks to see if the given value is defined by the enum or was not known at
 * the time this source was generated.
 **/
BOOL MochiPBTruncation_IsValidValue(int32_t value);

#pragma mark - MochiPBTextRoot

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
@interface MochiPBTextRoot : GPBRootObject
@end

#pragma mark - MochiPBSizeFunc

typedef GPB_ENUM(MochiPBSizeFunc_FieldNumber) {
  MochiPBSizeFunc_FieldNumber_Text = 1,
  MochiPBSizeFunc_FieldNumber_MinSize = 2,
  MochiPBSizeFunc_FieldNumber_MaxSize = 3,
};

@interface MochiPBSizeFunc : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) MochiPBText *text;
/** Test to see if @c text has been set. */
@property(nonatomic, readwrite) BOOL hasText;

@property(nonatomic, readwrite, strong, null_resettable) MochiPBPoint *minSize;
/** Test to see if @c minSize has been set. */
@property(nonatomic, readwrite) BOOL hasMinSize;

@property(nonatomic, readwrite, strong, null_resettable) MochiPBPoint *maxSize;
/** Test to see if @c maxSize has been set. */
@property(nonatomic, readwrite) BOOL hasMaxSize;

@end

#pragma mark - MochiPBText

typedef GPB_ENUM(MochiPBText_FieldNumber) {
  MochiPBText_FieldNumber_Style = 1,
  MochiPBText_FieldNumber_Text = 2,
};

@interface MochiPBText : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) MochiPBTextStyle *style;
/** Test to see if @c style has been set. */
@property(nonatomic, readwrite) BOOL hasStyle;

@property(nonatomic, readwrite, copy, null_resettable) NSString *text;

@end

#pragma mark - MochiPBFont

typedef GPB_ENUM(MochiPBFont_FieldNumber) {
  MochiPBFont_FieldNumber_Family = 1,
  MochiPBFont_FieldNumber_Face = 2,
  MochiPBFont_FieldNumber_Size = 3,
};

@interface MochiPBFont : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *family;

@property(nonatomic, readwrite, copy, null_resettable) NSString *face;

@property(nonatomic, readwrite) double size;

@end

#pragma mark - MochiPBColor

typedef GPB_ENUM(MochiPBColor_FieldNumber) {
  MochiPBColor_FieldNumber_Red = 1,
  MochiPBColor_FieldNumber_Blue = 2,
  MochiPBColor_FieldNumber_Green = 3,
  MochiPBColor_FieldNumber_Alpha = 4,
};

@interface MochiPBColor : GPBMessage

@property(nonatomic, readwrite) uint32_t red;

@property(nonatomic, readwrite) uint32_t blue;

@property(nonatomic, readwrite) uint32_t green;

@property(nonatomic, readwrite) uint32_t alpha;

@end

#pragma mark - MochiPBTextStyle

typedef GPB_ENUM(MochiPBTextStyle_FieldNumber) {
  MochiPBTextStyle_FieldNumber_TextAlignment = 2,
  MochiPBTextStyle_FieldNumber_StrikethroughStyle = 4,
  MochiPBTextStyle_FieldNumber_StrikethroughColor = 6,
  MochiPBTextStyle_FieldNumber_UnderlineStyle = 8,
  MochiPBTextStyle_FieldNumber_UnderlineColor = 10,
  MochiPBTextStyle_FieldNumber_Font = 12,
  MochiPBTextStyle_FieldNumber_Hyphenation = 14,
  MochiPBTextStyle_FieldNumber_LineHeightMultiple = 16,
  MochiPBTextStyle_FieldNumber_MaxLines = 18,
  MochiPBTextStyle_FieldNumber_TextColor = 20,
  MochiPBTextStyle_FieldNumber_Wrap = 22,
  MochiPBTextStyle_FieldNumber_Truncation = 24,
  MochiPBTextStyle_FieldNumber_TruncationString = 26,
};

@interface MochiPBTextStyle : GPBMessage

@property(nonatomic, readwrite) MochiPBTextAlignment textAlignment;

@property(nonatomic, readwrite) MochiPBStrikethroughStyle strikethroughStyle;

@property(nonatomic, readwrite, strong, null_resettable) MochiPBColor *strikethroughColor;
/** Test to see if @c strikethroughColor has been set. */
@property(nonatomic, readwrite) BOOL hasStrikethroughColor;

@property(nonatomic, readwrite) MochiPBUnderlineStyle underlineStyle;

@property(nonatomic, readwrite, strong, null_resettable) MochiPBColor *underlineColor;
/** Test to see if @c underlineColor has been set. */
@property(nonatomic, readwrite) BOOL hasUnderlineColor;

@property(nonatomic, readwrite, strong, null_resettable) MochiPBFont *font;
/** Test to see if @c font has been set. */
@property(nonatomic, readwrite) BOOL hasFont;

@property(nonatomic, readwrite) double hyphenation;

@property(nonatomic, readwrite) double lineHeightMultiple;

@property(nonatomic, readwrite) int64_t maxLines;

@property(nonatomic, readwrite, strong, null_resettable) MochiPBColor *textColor;
/** Test to see if @c textColor has been set. */
@property(nonatomic, readwrite) BOOL hasTextColor;

@property(nonatomic, readwrite) MochiPBTextWrap wrap;

@property(nonatomic, readwrite) MochiPBTruncation truncation;

@property(nonatomic, readwrite, copy, null_resettable) NSString *truncationString;

@end

/**
 * Fetches the raw value of a @c MochiPBTextStyle's @c textAlignment property, even
 * if the value was not defined by the enum at the time the code was generated.
 **/
int32_t MochiPBTextStyle_TextAlignment_RawValue(MochiPBTextStyle *message);
/**
 * Sets the raw value of an @c MochiPBTextStyle's @c textAlignment property, allowing
 * it to be set to a value that was not defined by the enum at the time the code
 * was generated.
 **/
void SetMochiPBTextStyle_TextAlignment_RawValue(MochiPBTextStyle *message, int32_t value);

/**
 * Fetches the raw value of a @c MochiPBTextStyle's @c strikethroughStyle property, even
 * if the value was not defined by the enum at the time the code was generated.
 **/
int32_t MochiPBTextStyle_StrikethroughStyle_RawValue(MochiPBTextStyle *message);
/**
 * Sets the raw value of an @c MochiPBTextStyle's @c strikethroughStyle property, allowing
 * it to be set to a value that was not defined by the enum at the time the code
 * was generated.
 **/
void SetMochiPBTextStyle_StrikethroughStyle_RawValue(MochiPBTextStyle *message, int32_t value);

/**
 * Fetches the raw value of a @c MochiPBTextStyle's @c underlineStyle property, even
 * if the value was not defined by the enum at the time the code was generated.
 **/
int32_t MochiPBTextStyle_UnderlineStyle_RawValue(MochiPBTextStyle *message);
/**
 * Sets the raw value of an @c MochiPBTextStyle's @c underlineStyle property, allowing
 * it to be set to a value that was not defined by the enum at the time the code
 * was generated.
 **/
void SetMochiPBTextStyle_UnderlineStyle_RawValue(MochiPBTextStyle *message, int32_t value);

/**
 * Fetches the raw value of a @c MochiPBTextStyle's @c wrap property, even
 * if the value was not defined by the enum at the time the code was generated.
 **/
int32_t MochiPBTextStyle_Wrap_RawValue(MochiPBTextStyle *message);
/**
 * Sets the raw value of an @c MochiPBTextStyle's @c wrap property, allowing
 * it to be set to a value that was not defined by the enum at the time the code
 * was generated.
 **/
void SetMochiPBTextStyle_Wrap_RawValue(MochiPBTextStyle *message, int32_t value);

/**
 * Fetches the raw value of a @c MochiPBTextStyle's @c truncation property, even
 * if the value was not defined by the enum at the time the code was generated.
 **/
int32_t MochiPBTextStyle_Truncation_RawValue(MochiPBTextStyle *message);
/**
 * Sets the raw value of an @c MochiPBTextStyle's @c truncation property, allowing
 * it to be set to a value that was not defined by the enum at the time the code
 * was generated.
 **/
void SetMochiPBTextStyle_Truncation_RawValue(MochiPBTextStyle *message, int32_t value);

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)