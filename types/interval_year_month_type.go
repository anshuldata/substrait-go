package types

import (
	"fmt"

	"github.com/substrait-io/substrait-go/proto"
)

// IntervalYearToMonthType this is used to represent a type of interval which represents YearToMonth.
type IntervalYearToMonthType struct {
	typeVariationRef uint32
	nullability      Nullability
}

// NewIntervalYearToMonthType creates a type of new interval YearToMonth.
// Created type has nullability as Nullable
func NewIntervalYearToMonthType() IntervalYearToMonthType {
	return IntervalYearToMonthType{
		nullability: NullabilityNullable,
	}
}

func (IntervalYearToMonthType) isRootRef() {}
func (m IntervalYearToMonthType) WithNullability(n Nullability) Type {
	return m.withNullability(n)
}

func (m IntervalYearToMonthType) withNullability(n Nullability) IntervalYearToMonthType {
	return IntervalYearToMonthType{
		nullability: n,
	}
}

func (m IntervalYearToMonthType) GetType() Type                     { return m }
func (m IntervalYearToMonthType) GetNullability() Nullability       { return m.nullability }
func (m IntervalYearToMonthType) GetTypeVariationReference() uint32 { return m.typeVariationRef }
func (m IntervalYearToMonthType) Equals(rhs Type) bool {
	if o, ok := rhs.(IntervalYearToMonthType); ok {
		return o == m
	}
	return false
}

func (m IntervalYearToMonthType) ToProtoFuncArg() *proto.FunctionArgument {
	return &proto.FunctionArgument{
		ArgType: &proto.FunctionArgument_Type{Type: m.ToProto()},
	}
}

func (m IntervalYearToMonthType) ToProto() *proto.Type {
	return &proto.Type{Kind: &proto.Type_IntervalYear_{
		IntervalYear: &proto.Type_IntervalYear{
			Nullability:            m.nullability,
			TypeVariationReference: m.typeVariationRef}}}
}

func (IntervalYearToMonthType) ShortString() string { return "intrytm" }
func (m IntervalYearToMonthType) String() string {
	return fmt.Sprintf("intervalyeartomonth%s", strNullable(m))
}
