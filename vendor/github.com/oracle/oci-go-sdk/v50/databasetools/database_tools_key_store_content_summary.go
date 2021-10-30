// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Database Tools APIs to manage Connections and Private Endpoints.
//

package databasetools

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// DatabaseToolsKeyStoreContentSummary The key store content.
type DatabaseToolsKeyStoreContentSummary interface {
}

type databasetoolskeystorecontentsummary struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorecontentsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorecontentsummary databasetoolskeystorecontentsummary
	s := struct {
		Model Unmarshalerdatabasetoolskeystorecontentsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorecontentsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStoreContentSecretIdSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m databasetoolskeystorecontentsummary) String() string {
	return common.PointerString(m)
}

// DatabaseToolsKeyStoreContentSummaryValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStoreContentSummaryValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStoreContentSummaryValueTypeEnum
const (
	DatabaseToolsKeyStoreContentSummaryValueTypeSecretid DatabaseToolsKeyStoreContentSummaryValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStoreContentSummaryValueType = map[string]DatabaseToolsKeyStoreContentSummaryValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStoreContentSummaryValueTypeSecretid,
}

// GetDatabaseToolsKeyStoreContentSummaryValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStoreContentSummaryValueTypeEnum
func GetDatabaseToolsKeyStoreContentSummaryValueTypeEnumValues() []DatabaseToolsKeyStoreContentSummaryValueTypeEnum {
	values := make([]DatabaseToolsKeyStoreContentSummaryValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStoreContentSummaryValueType {
		values = append(values, v)
	}
	return values
}
