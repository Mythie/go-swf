package swftype

import "github.com/mythie/go-swf/pkg/swfreader"

type FilterList struct {
	NumberOfFilters uint8
	Filters         []*Filter
}

func ParseFilterList(reader swfreader.Reader) (*FilterList, error) {
	numberOfFilters, err := reader.ReadUInt8()

	if err != nil {
		return nil, err
	}

	filters := make([]*Filter, numberOfFilters)

	for i := range filters {
		filters[i], err = ParseFilter(reader)

		if err != nil {
			return nil, err
		}
	}

	return &FilterList{
		NumberOfFilters: numberOfFilters,
		Filters:         filters,
	}, nil
}
