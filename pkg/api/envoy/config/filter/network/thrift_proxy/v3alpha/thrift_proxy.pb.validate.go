// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/network/thrift_proxy/v3alpha/thrift_proxy.proto

package envoy_config_filter_network_thrift_proxy_v3alpha

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on ThriftProxy with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ThriftProxy) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := TransportType_name[int32(m.GetTransport())]; !ok {
		return ThriftProxyValidationError{
			field:  "Transport",
			reason: "value must be one of the defined enum values",
		}
	}

	if _, ok := ProtocolType_name[int32(m.GetProtocol())]; !ok {
		return ThriftProxyValidationError{
			field:  "Protocol",
			reason: "value must be one of the defined enum values",
		}
	}

	if len(m.GetStatPrefix()) < 1 {
		return ThriftProxyValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 bytes",
		}
	}

	{
		tmp := m.GetRouteConfig()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ThriftProxyValidationError{
					field:  "RouteConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	for idx, item := range m.GetThriftFilters() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ThriftProxyValidationError{
						field:  fmt.Sprintf("ThriftFilters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// ThriftProxyValidationError is the validation error returned by
// ThriftProxy.Validate if the designated constraints aren't met.
type ThriftProxyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ThriftProxyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ThriftProxyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ThriftProxyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ThriftProxyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ThriftProxyValidationError) ErrorName() string { return "ThriftProxyValidationError" }

// Error satisfies the builtin error interface
func (e ThriftProxyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sThriftProxy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ThriftProxyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ThriftProxyValidationError{}

// Validate checks the field values on ThriftFilter with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ThriftFilter) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return ThriftFilterValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
	}

	switch m.ConfigType.(type) {

	case *ThriftFilter_Config:

		{
			tmp := m.GetConfig()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ThriftFilterValidationError{
						field:  "Config",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *ThriftFilter_TypedConfig:

		{
			tmp := m.GetTypedConfig()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ThriftFilterValidationError{
						field:  "TypedConfig",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// ThriftFilterValidationError is the validation error returned by
// ThriftFilter.Validate if the designated constraints aren't met.
type ThriftFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ThriftFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ThriftFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ThriftFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ThriftFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ThriftFilterValidationError) ErrorName() string { return "ThriftFilterValidationError" }

// Error satisfies the builtin error interface
func (e ThriftFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sThriftFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ThriftFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ThriftFilterValidationError{}

// Validate checks the field values on ThriftProtocolOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ThriftProtocolOptions) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := TransportType_name[int32(m.GetTransport())]; !ok {
		return ThriftProtocolOptionsValidationError{
			field:  "Transport",
			reason: "value must be one of the defined enum values",
		}
	}

	if _, ok := ProtocolType_name[int32(m.GetProtocol())]; !ok {
		return ThriftProtocolOptionsValidationError{
			field:  "Protocol",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// ThriftProtocolOptionsValidationError is the validation error returned by
// ThriftProtocolOptions.Validate if the designated constraints aren't met.
type ThriftProtocolOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ThriftProtocolOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ThriftProtocolOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ThriftProtocolOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ThriftProtocolOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ThriftProtocolOptionsValidationError) ErrorName() string {
	return "ThriftProtocolOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e ThriftProtocolOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sThriftProtocolOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ThriftProtocolOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ThriftProtocolOptionsValidationError{}
