// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: cart.proto

package proto

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on CartItem with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CartItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CartItem with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CartItemMultiError, or nil
// if none found.
func (m *CartItem) ValidateAll() error {
	return m.validate(true)
}

func (m *CartItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProductId

	// no validation rules for Quantity

	// no validation rules for Price

	if len(errors) > 0 {
		return CartItemMultiError(errors)
	}

	return nil
}

// CartItemMultiError is an error wrapping multiple validation errors returned
// by CartItem.ValidateAll() if the designated constraints aren't met.
type CartItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CartItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CartItemMultiError) AllErrors() []error { return m }

// CartItemValidationError is the validation error returned by
// CartItem.Validate if the designated constraints aren't met.
type CartItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CartItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CartItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CartItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CartItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CartItemValidationError) ErrorName() string { return "CartItemValidationError" }

// Error satisfies the builtin error interface
func (e CartItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCartItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CartItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CartItemValidationError{}

// Validate checks the field values on Cart with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Cart) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Cart with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CartMultiError, or nil if none found.
func (m *Cart) ValidateAll() error {
	return m.validate(true)
}

func (m *Cart) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CartValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CartValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CartValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for TotalPrice

	if len(errors) > 0 {
		return CartMultiError(errors)
	}

	return nil
}

// CartMultiError is an error wrapping multiple validation errors returned by
// Cart.ValidateAll() if the designated constraints aren't met.
type CartMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CartMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CartMultiError) AllErrors() []error { return m }

// CartValidationError is the validation error returned by Cart.Validate if the
// designated constraints aren't met.
type CartValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CartValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CartValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CartValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CartValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CartValidationError) ErrorName() string { return "CartValidationError" }

// Error satisfies the builtin error interface
func (e CartValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCart.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CartValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CartValidationError{}

// Validate checks the field values on CartUserId with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CartUserId) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CartUserId with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CartUserIdMultiError, or
// nil if none found.
func (m *CartUserId) ValidateAll() error {
	return m.validate(true)
}

func (m *CartUserId) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	if len(errors) > 0 {
		return CartUserIdMultiError(errors)
	}

	return nil
}

// CartUserIdMultiError is an error wrapping multiple validation errors
// returned by CartUserId.ValidateAll() if the designated constraints aren't met.
type CartUserIdMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CartUserIdMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CartUserIdMultiError) AllErrors() []error { return m }

// CartUserIdValidationError is the validation error returned by
// CartUserId.Validate if the designated constraints aren't met.
type CartUserIdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CartUserIdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CartUserIdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CartUserIdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CartUserIdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CartUserIdValidationError) ErrorName() string { return "CartUserIdValidationError" }

// Error satisfies the builtin error interface
func (e CartUserIdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCartUserId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CartUserIdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CartUserIdValidationError{}

// Validate checks the field values on CartItemRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CartItemRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CartItemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CartItemRequestMultiError, or nil if none found.
func (m *CartItemRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CartItemRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for ProductId

	if len(errors) > 0 {
		return CartItemRequestMultiError(errors)
	}

	return nil
}

// CartItemRequestMultiError is an error wrapping multiple validation errors
// returned by CartItemRequest.ValidateAll() if the designated constraints
// aren't met.
type CartItemRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CartItemRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CartItemRequestMultiError) AllErrors() []error { return m }

// CartItemRequestValidationError is the validation error returned by
// CartItemRequest.Validate if the designated constraints aren't met.
type CartItemRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CartItemRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CartItemRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CartItemRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CartItemRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CartItemRequestValidationError) ErrorName() string { return "CartItemRequestValidationError" }

// Error satisfies the builtin error interface
func (e CartItemRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCartItemRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CartItemRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CartItemRequestValidationError{}

// Validate checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Empty) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EmptyMultiError, or nil if none found.
func (m *Empty) ValidateAll() error {
	return m.validate(true)
}

func (m *Empty) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EmptyMultiError(errors)
	}

	return nil
}

// EmptyMultiError is an error wrapping multiple validation errors returned by
// Empty.ValidateAll() if the designated constraints aren't met.
type EmptyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmptyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmptyMultiError) AllErrors() []error { return m }

// EmptyValidationError is the validation error returned by Empty.Validate if
// the designated constraints aren't met.
type EmptyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyValidationError) ErrorName() string { return "EmptyValidationError" }

// Error satisfies the builtin error interface
func (e EmptyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmpty.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyValidationError{}
