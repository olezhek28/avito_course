// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: service.proto

package event_v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on EventInfo with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *EventInfo) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetTitle()) < 1 {
		return EventInfoValidationError{
			field:  "Title",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetStartDate() == nil {
		return EventInfoValidationError{
			field:  "StartDate",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetEndDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventInfoValidationError{
				field:  "EndDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetNotificationIntervalMin()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventInfoValidationError{
				field:  "NotificationIntervalMin",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDescription()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventInfoValidationError{
				field:  "Description",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetOwnerId() <= 0 {
		return EventInfoValidationError{
			field:  "OwnerId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// EventInfoValidationError is the validation error returned by
// EventInfo.Validate if the designated constraints aren't met.
type EventInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventInfoValidationError) ErrorName() string { return "EventInfoValidationError" }

// Error satisfies the builtin error interface
func (e EventInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEventInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventInfoValidationError{}

// Validate checks the field values on Event with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Event) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if v, ok := interface{}(m.GetEventInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventValidationError{
				field:  "EventInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// EventValidationError is the validation error returned by Event.Validate if
// the designated constraints aren't met.
type EventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventValidationError) ErrorName() string { return "EventValidationError" }

// Error satisfies the builtin error interface
func (e EventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventValidationError{}

// Validate checks the field values on CreateEventRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateEventRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetEventInfo() == nil {
		return CreateEventRequestValidationError{
			field:  "EventInfo",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetEventInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateEventRequestValidationError{
				field:  "EventInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateEventRequestValidationError is the validation error returned by
// CreateEventRequest.Validate if the designated constraints aren't met.
type CreateEventRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateEventRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateEventRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateEventRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateEventRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateEventRequestValidationError) ErrorName() string {
	return "CreateEventRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateEventRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateEventRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateEventRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateEventRequestValidationError{}

// Validate checks the field values on UpdateEventRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateEventRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateEventRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdateEventInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateEventRequestValidationError{
				field:  "UpdateEventInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateEventRequestValidationError is the validation error returned by
// UpdateEventRequest.Validate if the designated constraints aren't met.
type UpdateEventRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateEventRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateEventRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateEventRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateEventRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateEventRequestValidationError) ErrorName() string {
	return "UpdateEventRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateEventRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateEventRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateEventRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateEventRequestValidationError{}

// Validate checks the field values on DeleteEventRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteEventRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeleteEventRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DeleteEventRequestValidationError is the validation error returned by
// DeleteEventRequest.Validate if the designated constraints aren't met.
type DeleteEventRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteEventRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteEventRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteEventRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteEventRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteEventRequestValidationError) ErrorName() string {
	return "DeleteEventRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteEventRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteEventRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteEventRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteEventRequestValidationError{}

// Validate checks the field values on GetEventListForDayRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetEventListForDayRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetEventListForDayRequestValidationError{
				field:  "Date",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetEventListForDayRequestValidationError is the validation error returned by
// GetEventListForDayRequest.Validate if the designated constraints aren't met.
type GetEventListForDayRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEventListForDayRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEventListForDayRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEventListForDayRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEventListForDayRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEventListForDayRequestValidationError) ErrorName() string {
	return "GetEventListForDayRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetEventListForDayRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEventListForDayRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEventListForDayRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEventListForDayRequestValidationError{}

// Validate checks the field values on GetEventListForDayResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetEventListForDayResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetEventListForDayResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetEventListForDayResponseValidationError is the validation error returned
// by GetEventListForDayResponse.Validate if the designated constraints aren't met.
type GetEventListForDayResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEventListForDayResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEventListForDayResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEventListForDayResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEventListForDayResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEventListForDayResponseValidationError) ErrorName() string {
	return "GetEventListForDayResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetEventListForDayResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEventListForDayResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEventListForDayResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEventListForDayResponseValidationError{}

// Validate checks the field values on GetEventListForWeekRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetEventListForWeekRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetWeekStart()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetEventListForWeekRequestValidationError{
				field:  "WeekStart",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetEventListForWeekRequestValidationError is the validation error returned
// by GetEventListForWeekRequest.Validate if the designated constraints aren't met.
type GetEventListForWeekRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEventListForWeekRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEventListForWeekRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEventListForWeekRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEventListForWeekRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEventListForWeekRequestValidationError) ErrorName() string {
	return "GetEventListForWeekRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetEventListForWeekRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEventListForWeekRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEventListForWeekRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEventListForWeekRequestValidationError{}

// Validate checks the field values on GetEventListForWeekResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetEventListForWeekResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetEventListForWeekResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetEventListForWeekResponseValidationError is the validation error returned
// by GetEventListForWeekResponse.Validate if the designated constraints
// aren't met.
type GetEventListForWeekResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEventListForWeekResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEventListForWeekResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEventListForWeekResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEventListForWeekResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEventListForWeekResponseValidationError) ErrorName() string {
	return "GetEventListForWeekResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetEventListForWeekResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEventListForWeekResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEventListForWeekResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEventListForWeekResponseValidationError{}

// Validate checks the field values on GetEventListForMonthRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetEventListForMonthRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetMonthStart()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetEventListForMonthRequestValidationError{
				field:  "MonthStart",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetEventListForMonthRequestValidationError is the validation error returned
// by GetEventListForMonthRequest.Validate if the designated constraints
// aren't met.
type GetEventListForMonthRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEventListForMonthRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEventListForMonthRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEventListForMonthRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEventListForMonthRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEventListForMonthRequestValidationError) ErrorName() string {
	return "GetEventListForMonthRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetEventListForMonthRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEventListForMonthRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEventListForMonthRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEventListForMonthRequestValidationError{}

// Validate checks the field values on GetEventListForMonthResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetEventListForMonthResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetEventListForMonthResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetEventListForMonthResponseValidationError is the validation error returned
// by GetEventListForMonthResponse.Validate if the designated constraints
// aren't met.
type GetEventListForMonthResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEventListForMonthResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEventListForMonthResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEventListForMonthResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEventListForMonthResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEventListForMonthResponseValidationError) ErrorName() string {
	return "GetEventListForMonthResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetEventListForMonthResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEventListForMonthResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEventListForMonthResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEventListForMonthResponseValidationError{}

// Validate checks the field values on UpdateEventRequest_UpdateEventInfo with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *UpdateEventRequest_UpdateEventInfo) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTitle()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateEventRequest_UpdateEventInfoValidationError{
				field:  "Title",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetStartDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateEventRequest_UpdateEventInfoValidationError{
				field:  "StartDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetEndDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateEventRequest_UpdateEventInfoValidationError{
				field:  "EndDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetNotificationIntervalMin()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateEventRequest_UpdateEventInfoValidationError{
				field:  "NotificationIntervalMin",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDescription()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateEventRequest_UpdateEventInfoValidationError{
				field:  "Description",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetOwnerId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateEventRequest_UpdateEventInfoValidationError{
				field:  "OwnerId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateEventRequest_UpdateEventInfoValidationError is the validation error
// returned by UpdateEventRequest_UpdateEventInfo.Validate if the designated
// constraints aren't met.
type UpdateEventRequest_UpdateEventInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateEventRequest_UpdateEventInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateEventRequest_UpdateEventInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateEventRequest_UpdateEventInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateEventRequest_UpdateEventInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateEventRequest_UpdateEventInfoValidationError) ErrorName() string {
	return "UpdateEventRequest_UpdateEventInfoValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateEventRequest_UpdateEventInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateEventRequest_UpdateEventInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateEventRequest_UpdateEventInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateEventRequest_UpdateEventInfoValidationError{}

// Validate checks the field values on GetEventListForDayResponse_Result with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *GetEventListForDayResponse_Result) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetEvents() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetEventListForDayResponse_ResultValidationError{
					field:  fmt.Sprintf("Events[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetEventListForDayResponse_ResultValidationError is the validation error
// returned by GetEventListForDayResponse_Result.Validate if the designated
// constraints aren't met.
type GetEventListForDayResponse_ResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEventListForDayResponse_ResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEventListForDayResponse_ResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEventListForDayResponse_ResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEventListForDayResponse_ResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEventListForDayResponse_ResultValidationError) ErrorName() string {
	return "GetEventListForDayResponse_ResultValidationError"
}

// Error satisfies the builtin error interface
func (e GetEventListForDayResponse_ResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEventListForDayResponse_Result.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEventListForDayResponse_ResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEventListForDayResponse_ResultValidationError{}

// Validate checks the field values on GetEventListForWeekResponse_Result with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *GetEventListForWeekResponse_Result) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetEvents() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetEventListForWeekResponse_ResultValidationError{
					field:  fmt.Sprintf("Events[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetEventListForWeekResponse_ResultValidationError is the validation error
// returned by GetEventListForWeekResponse_Result.Validate if the designated
// constraints aren't met.
type GetEventListForWeekResponse_ResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEventListForWeekResponse_ResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEventListForWeekResponse_ResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEventListForWeekResponse_ResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEventListForWeekResponse_ResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEventListForWeekResponse_ResultValidationError) ErrorName() string {
	return "GetEventListForWeekResponse_ResultValidationError"
}

// Error satisfies the builtin error interface
func (e GetEventListForWeekResponse_ResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEventListForWeekResponse_Result.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEventListForWeekResponse_ResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEventListForWeekResponse_ResultValidationError{}

// Validate checks the field values on GetEventListForMonthResponse_Result with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *GetEventListForMonthResponse_Result) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetEvents() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetEventListForMonthResponse_ResultValidationError{
					field:  fmt.Sprintf("Events[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetEventListForMonthResponse_ResultValidationError is the validation error
// returned by GetEventListForMonthResponse_Result.Validate if the designated
// constraints aren't met.
type GetEventListForMonthResponse_ResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEventListForMonthResponse_ResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEventListForMonthResponse_ResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEventListForMonthResponse_ResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEventListForMonthResponse_ResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEventListForMonthResponse_ResultValidationError) ErrorName() string {
	return "GetEventListForMonthResponse_ResultValidationError"
}

// Error satisfies the builtin error interface
func (e GetEventListForMonthResponse_ResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEventListForMonthResponse_Result.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEventListForMonthResponse_ResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEventListForMonthResponse_ResultValidationError{}
