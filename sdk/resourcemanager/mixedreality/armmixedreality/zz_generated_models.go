//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmixedreality

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"time"
)

// AccountKeyRegenerateRequest - Request for account key regeneration
type AccountKeyRegenerateRequest struct {
	// Serial of key to be regenerated
	Serial *Serial `json:"serial,omitempty"`
}

// AccountKeys - Developer Keys of account
type AccountKeys struct {
	// READ-ONLY; value of primary key.
	PrimaryKey *string `json:"primaryKey,omitempty" azure:"ro"`

	// READ-ONLY; value of secondary key.
	SecondaryKey *string `json:"secondaryKey,omitempty" azure:"ro"`
}

// CheckNameAvailabilityRequest - Check Name Availability Request
type CheckNameAvailabilityRequest struct {
	// REQUIRED; Resource Name To Verify
	Name *string `json:"name,omitempty"`

	// REQUIRED; Fully qualified resource type which includes provider namespace
	Type *string `json:"type,omitempty"`
}

// CheckNameAvailabilityResponse - Check Name Availability Response
type CheckNameAvailabilityResponse struct {
	// REQUIRED; if name Available
	NameAvailable *bool `json:"nameAvailable,omitempty"`

	// detail message
	Message *string `json:"message,omitempty"`

	// Resource Name To Verify
	Reason *NameUnavailableReason `json:"reason,omitempty"`
}

// CloudError - An Error response.
// Implements the error and azcore.HTTPResponse interfaces.
type CloudError struct {
	raw string
	// An Error response.
	InnerError *CloudErrorBody `json:"error,omitempty"`
}

// Error implements the error interface for type CloudError.
// The contents of the error text are not contractual and subject to change.
func (e CloudError) Error() string {
	return e.raw
}

// CloudErrorBody - An error response from Azure.
type CloudErrorBody struct {
	// An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
	Code *string `json:"code,omitempty"`

	// A list of additional details about the error.
	Details []*CloudErrorBody `json:"details,omitempty"`

	// A message describing the error, intended to be suitable for displaying in a user interface.
	Message *string `json:"message,omitempty"`

	// The target of the particular error. For example, the name of the property in error.
	Target *string `json:"target,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type CloudErrorBody.
func (c CloudErrorBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", c.Code)
	populate(objectMap, "details", c.Details)
	populate(objectMap, "message", c.Message)
	populate(objectMap, "target", c.Target)
	return json.Marshal(objectMap)
}

// Identity for the resource.
type Identity struct {
	// The identity type.
	Type *string `json:"type,omitempty"`

	// READ-ONLY; The principal ID of resource identity.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; The tenant ID of resource.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// LogSpecification - Specifications of the Log for Azure Monitoring
type LogSpecification struct {
	// Blob duration of the log
	BlobDuration *string `json:"blobDuration,omitempty"`

	// Localized friendly display name of the log
	DisplayName *string `json:"displayName,omitempty"`

	// Name of the log
	Name *string `json:"name,omitempty"`
}

// MetricDimension - Specifications of the Dimension of metrics
type MetricDimension struct {
	// Localized friendly display name of the dimension
	DisplayName *string `json:"displayName,omitempty"`

	// Internal name of the dimension.
	InternalName *string `json:"internalName,omitempty"`

	// Name of the dimension
	Name *string `json:"name,omitempty"`

	// Flag to indicate export for Shoebox
	ToBeExportedForShoebox *bool `json:"toBeExportedForShoebox,omitempty"`
}

// MetricSpecification - Specifications of the Metrics for Azure Monitoring
type MetricSpecification struct {
	// Only provide one value for this field. Valid values: Average, Minimum, Maximum, Total, Count.
	AggregationType *string `json:"aggregationType,omitempty"`

	// Metric category
	Category *string `json:"category,omitempty"`

	// Dimensions of the metric
	Dimensions []*MetricDimension `json:"dimensions,omitempty"`

	// Localized friendly description of the metric
	DisplayDescription *string `json:"displayDescription,omitempty"`

	// Localized friendly display name of the metric
	DisplayName *string `json:"displayName,omitempty"`

	// Flag to indicate use of regional Mdm accounts
	EnableRegionalMdmAccount *bool `json:"enableRegionalMdmAccount,omitempty"`

	// Flag to determine is Zero is returned for time duration where no metric is emitted
	FillGapWithZero *bool `json:"fillGapWithZero,omitempty"`

	// Internal metric name.
	InternalMetricName *string `json:"internalMetricName,omitempty"`

	// Locked aggregation type of the metric
	LockedAggregationType *string `json:"lockedAggregationType,omitempty"`

	// Metric filter regex pattern
	MetricFilterPattern *string `json:"metricFilterPattern,omitempty"`

	// Name of the metric
	Name *string `json:"name,omitempty"`

	// Source mdm account
	SourceMdmAccount *string `json:"sourceMdmAccount,omitempty"`

	// Source mdm namespace
	SourceMdmNamespace *string `json:"sourceMdmNamespace,omitempty"`

	// Supported aggregation types. Valid values: Average, Minimum, Maximum, Total, Count.
	SupportedAggregationTypes []*string `json:"supportedAggregationTypes,omitempty"`

	// Supported time grains. Valid values: PT1M, PT5M, PT15M, PT30M, PT1H, PT6H, PT12H, P1D
	SupportedTimeGrainTypes []*string `json:"supportedTimeGrainTypes,omitempty"`

	// Unit that makes sense for the metric
	Unit *string `json:"unit,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type MetricSpecification.
func (m MetricSpecification) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "aggregationType", m.AggregationType)
	populate(objectMap, "category", m.Category)
	populate(objectMap, "dimensions", m.Dimensions)
	populate(objectMap, "displayDescription", m.DisplayDescription)
	populate(objectMap, "displayName", m.DisplayName)
	populate(objectMap, "enableRegionalMdmAccount", m.EnableRegionalMdmAccount)
	populate(objectMap, "fillGapWithZero", m.FillGapWithZero)
	populate(objectMap, "internalMetricName", m.InternalMetricName)
	populate(objectMap, "lockedAggregationType", m.LockedAggregationType)
	populate(objectMap, "metricFilterPattern", m.MetricFilterPattern)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "sourceMdmAccount", m.SourceMdmAccount)
	populate(objectMap, "sourceMdmNamespace", m.SourceMdmNamespace)
	populate(objectMap, "supportedAggregationTypes", m.SupportedAggregationTypes)
	populate(objectMap, "supportedTimeGrainTypes", m.SupportedTimeGrainTypes)
	populate(objectMap, "unit", m.Unit)
	return json.Marshal(objectMap)
}

// MixedRealityAccountProperties - Common Properties shared by Mixed Reality Accounts
type MixedRealityAccountProperties struct {
	// The name of the storage account associated with this accountId
	StorageAccountName *string `json:"storageAccountName,omitempty"`

	// READ-ONLY; Correspond domain name of certain Spatial Anchors Account
	AccountDomain *string `json:"accountDomain,omitempty" azure:"ro"`

	// READ-ONLY; unique id of certain account.
	AccountID *string `json:"accountId,omitempty" azure:"ro"`
}

// MixedRealityClientCheckNameAvailabilityLocalOptions contains the optional parameters for the MixedRealityClient.CheckNameAvailabilityLocal method.
type MixedRealityClientCheckNameAvailabilityLocalOptions struct {
	// placeholder for future optional parameters
}

// ObjectAnchorsAccount Response.
type ObjectAnchorsAccount struct {
	TrackedResource
	Identity *ObjectAnchorsAccountIdentity `json:"identity,omitempty"`

	// The kind of account, if supported
	Kind *SKU `json:"kind,omitempty"`

	// The plan associated with this account
	Plan *Identity `json:"plan,omitempty"`

	// Property bag.
	Properties *MixedRealityAccountProperties `json:"properties,omitempty"`

	// The sku associated with this account
	SKU *SKU `json:"sku,omitempty"`

	// READ-ONLY; The system metadata related to an object anchors account.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ObjectAnchorsAccount.
func (o ObjectAnchorsAccount) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	o.TrackedResource.marshalInternal(objectMap)
	populate(objectMap, "identity", o.Identity)
	populate(objectMap, "kind", o.Kind)
	populate(objectMap, "plan", o.Plan)
	populate(objectMap, "properties", o.Properties)
	populate(objectMap, "sku", o.SKU)
	populate(objectMap, "systemData", o.SystemData)
	return json.Marshal(objectMap)
}

type ObjectAnchorsAccountIdentity struct {
	Identity
}

// ObjectAnchorsAccountPage - Result of the request to get resource collection. It contains a list of resources and a URL link to get the next set of results.
type ObjectAnchorsAccountPage struct {
	// URL to get the next set of resource list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of resources supported by the Resource Provider.
	Value []*ObjectAnchorsAccount `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ObjectAnchorsAccountPage.
func (o ObjectAnchorsAccountPage) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// ObjectAnchorsAccountsCreateOptions contains the optional parameters for the ObjectAnchorsAccounts.Create method.
type ObjectAnchorsAccountsCreateOptions struct {
	// placeholder for future optional parameters
}

// ObjectAnchorsAccountsDeleteOptions contains the optional parameters for the ObjectAnchorsAccounts.Delete method.
type ObjectAnchorsAccountsDeleteOptions struct {
	// placeholder for future optional parameters
}

// ObjectAnchorsAccountsGetOptions contains the optional parameters for the ObjectAnchorsAccounts.Get method.
type ObjectAnchorsAccountsGetOptions struct {
	// placeholder for future optional parameters
}

// ObjectAnchorsAccountsListByResourceGroupOptions contains the optional parameters for the ObjectAnchorsAccounts.ListByResourceGroup method.
type ObjectAnchorsAccountsListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ObjectAnchorsAccountsListBySubscriptionOptions contains the optional parameters for the ObjectAnchorsAccounts.ListBySubscription method.
type ObjectAnchorsAccountsListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// ObjectAnchorsAccountsListKeysOptions contains the optional parameters for the ObjectAnchorsAccounts.ListKeys method.
type ObjectAnchorsAccountsListKeysOptions struct {
	// placeholder for future optional parameters
}

// ObjectAnchorsAccountsRegenerateKeysOptions contains the optional parameters for the ObjectAnchorsAccounts.RegenerateKeys method.
type ObjectAnchorsAccountsRegenerateKeysOptions struct {
	// placeholder for future optional parameters
}

// ObjectAnchorsAccountsUpdateOptions contains the optional parameters for the ObjectAnchorsAccounts.Update method.
type ObjectAnchorsAccountsUpdateOptions struct {
	// placeholder for future optional parameters
}

// Operation - REST API operation
type Operation struct {
	// The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// Whether or not this is a data plane operation
	IsDataAction *bool `json:"isDataAction,omitempty"`

	// Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`

	// The origin
	Origin *string `json:"origin,omitempty"`

	// Properties of the operation
	Properties *OperationProperties `json:"properties,omitempty"`
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// REQUIRED; Description of operation
	Description *string `json:"description,omitempty"`

	// REQUIRED; Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`

	// REQUIRED; Service provider: Microsoft.ResourceProvider
	Provider *string `json:"provider,omitempty"`

	// REQUIRED; Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string `json:"resource,omitempty"`
}

// OperationPage - Result of the request to list Resource Provider operations. It contains a list of operations and a URL link to get the next set of results.
type OperationPage struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of operations supported by the Resource Provider.
	Value []*Operation `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type OperationPage.
func (o OperationPage) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// OperationProperties - Operation properties.
type OperationProperties struct {
	// Service specification.
	ServiceSpecification *ServiceSpecification `json:"serviceSpecification,omitempty"`
}

// OperationsListOptions contains the optional parameters for the Operations.List method.
type OperationsListOptions struct {
	// placeholder for future optional parameters
}

// RemoteRenderingAccount Response.
type RemoteRenderingAccount struct {
	TrackedResource
	// The identity associated with this account
	Identity *Identity `json:"identity,omitempty"`

	// The kind of account, if supported
	Kind *SKU `json:"kind,omitempty"`

	// The plan associated with this account
	Plan *Identity `json:"plan,omitempty"`

	// Property bag.
	Properties *MixedRealityAccountProperties `json:"properties,omitempty"`

	// The sku associated with this account
	SKU *SKU `json:"sku,omitempty"`

	// READ-ONLY; System metadata for this account
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type RemoteRenderingAccount.
func (r RemoteRenderingAccount) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	r.TrackedResource.marshalInternal(objectMap)
	populate(objectMap, "identity", r.Identity)
	populate(objectMap, "kind", r.Kind)
	populate(objectMap, "plan", r.Plan)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "sku", r.SKU)
	populate(objectMap, "systemData", r.SystemData)
	return json.Marshal(objectMap)
}

// RemoteRenderingAccountPage - Result of the request to get resource collection. It contains a list of resources and a URL link to get the next set of
// results.
type RemoteRenderingAccountPage struct {
	// URL to get the next set of resource list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of resources supported by the Resource Provider.
	Value []*RemoteRenderingAccount `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type RemoteRenderingAccountPage.
func (r RemoteRenderingAccountPage) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// RemoteRenderingAccountsCreateOptions contains the optional parameters for the RemoteRenderingAccounts.Create method.
type RemoteRenderingAccountsCreateOptions struct {
	// placeholder for future optional parameters
}

// RemoteRenderingAccountsDeleteOptions contains the optional parameters for the RemoteRenderingAccounts.Delete method.
type RemoteRenderingAccountsDeleteOptions struct {
	// placeholder for future optional parameters
}

// RemoteRenderingAccountsGetOptions contains the optional parameters for the RemoteRenderingAccounts.Get method.
type RemoteRenderingAccountsGetOptions struct {
	// placeholder for future optional parameters
}

// RemoteRenderingAccountsListByResourceGroupOptions contains the optional parameters for the RemoteRenderingAccounts.ListByResourceGroup method.
type RemoteRenderingAccountsListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// RemoteRenderingAccountsListBySubscriptionOptions contains the optional parameters for the RemoteRenderingAccounts.ListBySubscription method.
type RemoteRenderingAccountsListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// RemoteRenderingAccountsListKeysOptions contains the optional parameters for the RemoteRenderingAccounts.ListKeys method.
type RemoteRenderingAccountsListKeysOptions struct {
	// placeholder for future optional parameters
}

// RemoteRenderingAccountsRegenerateKeysOptions contains the optional parameters for the RemoteRenderingAccounts.RegenerateKeys method.
type RemoteRenderingAccountsRegenerateKeysOptions struct {
	// placeholder for future optional parameters
}

// RemoteRenderingAccountsUpdateOptions contains the optional parameters for the RemoteRenderingAccounts.Update method.
type RemoteRenderingAccountsUpdateOptions struct {
	// placeholder for future optional parameters
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	r.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

func (r Resource) marshalInternal(objectMap map[string]interface{}) {
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "type", r.Type)
}

// SKU - The resource model definition representing SKU
type SKU struct {
	// REQUIRED; The name of the SKU. Ex - P3. It is typically a letter+number code
	Name *string `json:"name,omitempty"`

	// If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible for the resource this may be omitted.
	Capacity *int32 `json:"capacity,omitempty"`

	// If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family *string `json:"family,omitempty"`

	// The SKU size. When the name field is the combination of tier and some other value, this would be the standalone code.
	Size *string `json:"size,omitempty"`

	// This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not required on a PUT.
	Tier *SKUTier `json:"tier,omitempty"`
}

// ServiceSpecification - Service specification payload
type ServiceSpecification struct {
	// Specifications of the Log for Azure Monitoring
	LogSpecifications []*LogSpecification `json:"logSpecifications,omitempty"`

	// Specifications of the Metrics for Azure Monitoring
	MetricSpecifications []*MetricSpecification `json:"metricSpecifications,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ServiceSpecification.
func (s ServiceSpecification) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "logSpecifications", s.LogSpecifications)
	populate(objectMap, "metricSpecifications", s.MetricSpecifications)
	return json.Marshal(objectMap)
}

// SpatialAnchorsAccount Response.
type SpatialAnchorsAccount struct {
	TrackedResource
	// The identity associated with this account
	Identity *Identity `json:"identity,omitempty"`

	// The kind of account, if supported
	Kind *SKU `json:"kind,omitempty"`

	// The plan associated with this account
	Plan *Identity `json:"plan,omitempty"`

	// Property bag.
	Properties *MixedRealityAccountProperties `json:"properties,omitempty"`

	// The sku associated with this account
	SKU *SKU `json:"sku,omitempty"`

	// READ-ONLY; System metadata for this account
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type SpatialAnchorsAccount.
func (s SpatialAnchorsAccount) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	s.TrackedResource.marshalInternal(objectMap)
	populate(objectMap, "identity", s.Identity)
	populate(objectMap, "kind", s.Kind)
	populate(objectMap, "plan", s.Plan)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "sku", s.SKU)
	populate(objectMap, "systemData", s.SystemData)
	return json.Marshal(objectMap)
}

// SpatialAnchorsAccountPage - Result of the request to get resource collection. It contains a list of resources and a URL link to get the next set of results.
type SpatialAnchorsAccountPage struct {
	// URL to get the next set of resource list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of resources supported by the Resource Provider.
	Value []*SpatialAnchorsAccount `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type SpatialAnchorsAccountPage.
func (s SpatialAnchorsAccountPage) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// SpatialAnchorsAccountsCreateOptions contains the optional parameters for the SpatialAnchorsAccounts.Create method.
type SpatialAnchorsAccountsCreateOptions struct {
	// placeholder for future optional parameters
}

// SpatialAnchorsAccountsDeleteOptions contains the optional parameters for the SpatialAnchorsAccounts.Delete method.
type SpatialAnchorsAccountsDeleteOptions struct {
	// placeholder for future optional parameters
}

// SpatialAnchorsAccountsGetOptions contains the optional parameters for the SpatialAnchorsAccounts.Get method.
type SpatialAnchorsAccountsGetOptions struct {
	// placeholder for future optional parameters
}

// SpatialAnchorsAccountsListByResourceGroupOptions contains the optional parameters for the SpatialAnchorsAccounts.ListByResourceGroup method.
type SpatialAnchorsAccountsListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// SpatialAnchorsAccountsListBySubscriptionOptions contains the optional parameters for the SpatialAnchorsAccounts.ListBySubscription method.
type SpatialAnchorsAccountsListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// SpatialAnchorsAccountsListKeysOptions contains the optional parameters for the SpatialAnchorsAccounts.ListKeys method.
type SpatialAnchorsAccountsListKeysOptions struct {
	// placeholder for future optional parameters
}

// SpatialAnchorsAccountsRegenerateKeysOptions contains the optional parameters for the SpatialAnchorsAccounts.RegenerateKeys method.
type SpatialAnchorsAccountsRegenerateKeysOptions struct {
	// placeholder for future optional parameters
}

// SpatialAnchorsAccountsUpdateOptions contains the optional parameters for the SpatialAnchorsAccounts.Update method.
type SpatialAnchorsAccountsUpdateOptions struct {
	// placeholder for future optional parameters
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags' and a 'location'
type TrackedResource struct {
	Resource
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	t.marshalInternal(objectMap)
	return json.Marshal(objectMap)
}

func (t TrackedResource) marshalInternal(objectMap map[string]interface{}) {
	t.Resource.marshalInternal(objectMap)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "tags", t.Tags)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}