// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
// Use this pattern to construct endpoints for identity domains: `https://<domainURL>/admin/v1/`. See Finding an Identity Domain URL (https://docs.oracle.com/en-us/iaas/Content/Identity/api-getstarted/locate-identity-domain-url.htm) to locate the domain URL you need.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AppExtensionEnterpriseAppApp This extension defines the Enterprise App related attributes.
type AppExtensionEnterpriseAppApp struct {

	// A list of AppResources of this App.
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsCompositeKey: [value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	AppResources []AppAppResources `mandatory:"false" json:"appResources"`

	DenyAuthzPolicy *AppDenyAuthzPolicy `mandatory:"false" json:"denyAuthzPolicy"`

	AllowAuthzPolicy *AppAllowAuthzPolicy `mandatory:"false" json:"allowAuthzPolicy"`

	// Allow Authz policy decision expiry time in seconds.
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 3600
	//  - idcsMinValue: 0
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	AllowAuthzDecisionTTL *int `mandatory:"false" json:"allowAuthzDecisionTTL"`

	// Deny Authz policy decision expiry time in seconds.
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 3600
	//  - idcsMinValue: 0
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	DenyAuthzDecisionTTL *int `mandatory:"false" json:"denyAuthzDecisionTTL"`
}

func (m AppExtensionEnterpriseAppApp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppExtensionEnterpriseAppApp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
