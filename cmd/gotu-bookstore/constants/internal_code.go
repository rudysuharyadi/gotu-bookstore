package constants

import (
	"gotu-bookstore/pkg/resfmt/base_error"
)

// This InternalCode is an extension from base_error. Each service need to use their own unique prefix and code, so that
// the error can be easily traced

const InternalCodePrefix = "GOTU-IC"

const (
	IC0001 base_error.InternalCode = InternalCodePrefix + "0001"
	IC0002 base_error.InternalCode = InternalCodePrefix + "0002"
	IC0003 base_error.InternalCode = InternalCodePrefix + "0003"
	IC0004 base_error.InternalCode = InternalCodePrefix + "0004"
	IC0005 base_error.InternalCode = InternalCodePrefix + "0005"
	IC0006 base_error.InternalCode = InternalCodePrefix + "0006"
	IC0007 base_error.InternalCode = InternalCodePrefix + "0007"
	IC0008 base_error.InternalCode = InternalCodePrefix + "0008"
	IC0009 base_error.InternalCode = InternalCodePrefix + "0009"
	IC0010 base_error.InternalCode = InternalCodePrefix + "0010"
	IC0011 base_error.InternalCode = InternalCodePrefix + "0011"
	IC0012 base_error.InternalCode = InternalCodePrefix + "0012"
	IC0013 base_error.InternalCode = InternalCodePrefix + "0013"
	IC0014 base_error.InternalCode = InternalCodePrefix + "0014"
	IC0015 base_error.InternalCode = InternalCodePrefix + "0015"
	IC0016 base_error.InternalCode = InternalCodePrefix + "0016"
	IC0017 base_error.InternalCode = InternalCodePrefix + "0017"
	IC0018 base_error.InternalCode = InternalCodePrefix + "0018"
	IC0019 base_error.InternalCode = InternalCodePrefix + "0019"
	IC0020 base_error.InternalCode = InternalCodePrefix + "0020"
	IC0021 base_error.InternalCode = InternalCodePrefix + "0021"
	IC0022 base_error.InternalCode = InternalCodePrefix + "0022"
	IC0023 base_error.InternalCode = InternalCodePrefix + "0023"
	IC0024 base_error.InternalCode = InternalCodePrefix + "0024"
	IC0025 base_error.InternalCode = InternalCodePrefix + "0025"
	IC0026 base_error.InternalCode = InternalCodePrefix + "0026"
	IC0027 base_error.InternalCode = InternalCodePrefix + "0027"
	IC0028 base_error.InternalCode = InternalCodePrefix + "0028"
	IC0029 base_error.InternalCode = InternalCodePrefix + "0029"
	IC0030 base_error.InternalCode = InternalCodePrefix + "0030"
)

var ErrorMessages = map[base_error.InternalCode]string{
	IC0001: "Error when bind payload into struct",
	IC0002: "Error when generate uuid",
	IC0003: "ApiKey not found",
	IC0004: "Unable to verify token",
	IC0005: "Authorization header not found",
	IC0006: "Unauthorized token",
	IC0007: "Recover from panic",
	IC0008: "Route not found",
	IC0009: "Invalid request",
	IC0010: "Error when create user data",
	IC0011: "Error when get user data",
	IC0012: "Error when generate token",
	IC0013: "Error when invalidate token",
	IC0014: "Error when get book data",
	IC0015: "Error when generate uuid",
	IC0016: "Error when add item to cart",
	IC0017: "Error when get shopping cart data",
	IC0018: "Error when get shopping cart item data",
	IC0019: "Error when update item on the cart",
	IC0020: "Error when clear shopping cart",
	IC0021: "Error when checkout",
	IC0022: "Error when delete item on the cart",
	IC0023: "Error when get data transactions",
}
