package clinical

import (
	"fmt"
	"gitlab.slade360emr.com/go/base"
	"io"
	"log"
	"strconv"
	"time"
)

// FHIRAddress definition: an address expressed using postal conventions (as opposed to gps or other location definition formats).  this data type may be used to convey addresses for use in delivering mail as well as for visiting locations which might not be valid for mail delivery.  there are a variety of postal address formats defined around the world.
type FHIRAddress struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The purpose of this address.
	Use *AddressUseEnum `json:"use,omitempty"`

	// Distinguishes between physical addresses (those you can visit) and mailing addresses (e.g. PO Boxes and care-of addresses). Most addresses are both.
	Type *AddressTypeEnum `json:"type,omitempty"`

	// Specifies the entire address as it should be displayed e.g. on a postal label. This may be provided instead of or as well as the specific parts.
	Text string `json:"text,omitempty"`

	// This component contains the house number, apartment number, street name, street direction,  P.O. Box number, delivery hints, and similar address information.
	Line []*string `json:"line,omitempty"`

	// The name of the city, town, suburb, village or other community or delivery center.
	City *string `json:"city,omitempty"`

	// The name of the administrative area (county).
	District *string `json:"district,omitempty"`

	// Sub-unit of a country with limited sovereignty in a federally organized country. A code may be used if codes are in common use (e.g. US 2 letter state codes).
	State *string `json:"state,omitempty"`

	// A postal code designating a region defined by the postal service.
	PostalCode *base.Code `json:"postalCode,omitempty"`

	// Country - a nation as commonly understood or generally accepted.
	Country *string `json:"country,omitempty"`

	// Time period when address was/is in use.
	Period *FHIRPeriod `json:"period,omitempty"`
}

// FHIRAddressInput is the input type for Address
type FHIRAddressInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The purpose of this address.
	Use *AddressUseEnum `json:"use,omitempty"`

	// Distinguishes between physical addresses (those you can visit) and mailing addresses (e.g. PO Boxes and care-of addresses). Most addresses are both.
	Type *AddressTypeEnum `json:"type,omitempty"`

	// Specifies the entire address as it should be displayed e.g. on a postal label. This may be provided instead of or as well as the specific parts.
	Text string `json:"text,omitempty"`

	// This component contains the house number, apartment number, street name, street direction,  P.O. Box number, delivery hints, and similar address information.
	Line *string `json:"line,omitempty"`

	// The name of the city, town, suburb, village or other community or delivery center.
	City *string `json:"city,omitempty"`

	// The name of the administrative area (county).
	District *string `json:"district,omitempty"`

	// Sub-unit of a country with limited sovereignty in a federally organized country. A code may be used if codes are in common use (e.g. US 2 letter state codes).
	State *string `json:"state,omitempty"`

	// A postal code designating a region defined by the postal service.
	PostalCode *base.Code `json:"postalCode,omitempty"`

	// Country - a nation as commonly understood or generally accepted.
	Country *string `json:"country,omitempty"`

	// Time period when address was/is in use.
	Period *FHIRPeriodInput `json:"period,omitempty"`
}

// FHIRAge definition: a duration of time during which an organism (or a process) has existed.
type FHIRAge struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Value *base.Decimal `json:"value,omitempty"`

	// How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Comparator *AgeComparatorEnum `json:"comparator,omitempty"`

	// A human-readable form of the unit.
	Unit *string `json:"unit,omitempty"`

	// The identification of the system that provides the coded form of the unit.
	System *base.URI `json:"system,omitempty"`

	// A computer processable form of the unit in some unit representation system.
	Code *base.Code `json:"code,omitempty"`
}

// FHIRAgeInput is the input type for Age
type FHIRAgeInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Value *base.Decimal `json:"value,omitempty"`

	// How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Comparator *AgeComparatorEnum `json:"comparator,omitempty"`

	// A human-readable form of the unit.
	Unit *string `json:"unit,omitempty"`

	// The identification of the system that provides the coded form of the unit.
	System *base.URI `json:"system,omitempty"`

	// A computer processable form of the unit in some unit representation system.
	Code *base.Code `json:"code,omitempty"`
}

// FHIRAnnotation definition: a  text note which also  contains information about who made the statement and when.
type FHIRAnnotation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The individual responsible for making the annotation.
	AuthorReference *FHIRReference `json:"authorReference,omitempty"`

	// The individual responsible for making the annotation.
	AuthorString *string `json:"authorString,omitempty"`

	// Indicates when this particular annotation was made.
	Time *time.Time `json:"time,omitempty"`

	// The text of the annotation in markdown format.
	Text *base.Markdown `json:"text,omitempty"`
}

// FHIRAnnotationInput is the input type for Annotation
type FHIRAnnotationInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The individual responsible for making the annotation.
	AuthorReference *FHIRReferenceInput `json:"authorReference,omitempty"`

	// The individual responsible for making the annotation.
	AuthorString *string `json:"authorString,omitempty"`

	// Indicates when this particular annotation was made.
	Time *base.DateTime `json:"time,omitempty"`

	// The text of the annotation in markdown format.
	Text *base.Markdown `json:"text,omitempty"`
}

// FHIRAttachment definition: for referring to data content defined in other formats.
type FHIRAttachment struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// Identifies the type of the data in the attachment and allows a method to be chosen to interpret or render the data. Includes mime type parameters such as charset where appropriate.
	ContentType *base.Code `json:"contentType,omitempty"`

	// The human language of the content. The value can be any valid value according to BCP 47.
	Language *base.Code `json:"language,omitempty"`

	// The actual data of the attachment - a sequence of bytes, base64 encoded.
	Data *base.Base64Binary `json:"data,omitempty"`

	// A location where the data can be accessed.
	URL *base.URL `json:"url,omitempty"`

	// The number of bytes of data that make up this attachment (before base64 encoding, if that is done).
	Size *int `json:"size,omitempty"`

	// The calculated hash of the data using SHA-1. Represented using base64.
	Hash *base.Base64Binary `json:"hash,omitempty"`

	// A label or set of text to display in place of the data.
	Title *string `json:"title,omitempty"`

	// The date that the attachment was first created.
	Creation *base.DateTime `json:"creation,omitempty"`
}

// FHIRAttachmentInput is the input type for Attachment
type FHIRAttachmentInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// Identifies the type of the data in the attachment and allows a method to be chosen to interpret or render the data. Includes mime type parameters such as charset where appropriate.
	ContentType *base.Code `json:"contentType,omitempty"`

	// The human language of the content. The value can be any valid value according to BCP 47.
	Language *base.Code `json:"language,omitempty"`

	// The actual data of the attachment - a sequence of bytes, base64 encoded.
	Data *base.Base64Binary `json:"data,omitempty"`

	// A location where the data can be accessed.
	URL *base.URL `json:"url,omitempty"`

	// The number of bytes of data that make up this attachment (before base64 encoding, if that is done).
	Size *int `json:"size,omitempty"`

	// The calculated hash of the data using SHA-1. Represented using base64.
	Hash *base.Base64Binary `json:"hash,omitempty"`

	// A label or set of text to display in place of the data.
	Title *string `json:"title,omitempty"`

	// The date that the attachment was first created.
	Creation *base.DateTime `json:"creation,omitempty"`
}

// FHIRCodeableConcept definition: a concept that may be defined by a formal reference to a terminology or ontology or may be provided by text.
type FHIRCodeableConcept struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// A reference to a code defined by a terminology system.
	Coding []*FHIRCoding `json:"coding,omitempty"`

	// A human language representation of the concept as seen/selected/uttered by the user who entered the data and/or which represents the intended meaning of the user.
	Text string `json:"text,omitempty"`
}

// FHIRCodeableConceptInput is the input type for CodeableConcept
type FHIRCodeableConceptInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// A reference to a code defined by a terminology system.
	Coding []*FHIRCodingInput `json:"coding,omitempty"`

	// A human language representation of the concept as seen/selected/uttered by the user who entered the data and/or which represents the intended meaning of the user.
	Text string `json:"text,omitempty"`
}

// FHIRCoding definition: a reference to a code defined by a terminology system.
type FHIRCoding struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The identification of the code system that defines the meaning of the symbol in the code.
	System *base.URI `json:"system,omitempty"`

	// The version of the code system which was used when choosing this code. Note that a well-maintained code system does not need the version reported, because the meaning of codes is consistent across versions. However this cannot consistently be assured, and when the meaning is not guaranteed to be consistent, the version SHOULD be exchanged.
	Version *string `json:"version,omitempty"`

	// A symbol in syntax defined by the system. The symbol may be a predefined code or an expression in a syntax defined by the coding system (e.g. post-coordination).
	Code base.Code `json:"code,omitempty"`

	// A representation of the meaning of the code in the system, following the rules of the system.
	Display string `json:"display,omitempty"`

	// Indicates that this coding was chosen by a user directly - e.g. off a pick list of available items (codes or displays).
	UserSelected *bool `json:"userSelected,omitempty"`
}

// FHIRCodingInput is the input type for Coding
type FHIRCodingInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The identification of the code system that defines the meaning of the symbol in the code.
	System *base.URI `json:"system,omitempty"`

	// The version of the code system which was used when choosing this code. Note that a well-maintained code system does not need the version reported, because the meaning of codes is consistent across versions. However this cannot consistently be assured, and when the meaning is not guaranteed to be consistent, the version SHOULD be exchanged.
	Version *string `json:"version,omitempty"`

	// A symbol in syntax defined by the system. The symbol may be a predefined code or an expression in a syntax defined by the coding system (e.g. post-coordination).
	Code base.Code `json:"code,omitempty"`

	// A representation of the meaning of the code in the system, following the rules of the system.
	Display string `json:"display,omitempty"`

	// Indicates that this coding was chosen by a user directly - e.g. off a pick list of available items (codes or displays).
	UserSelected *bool `json:"userSelected,omitempty"`
}

// FHIRContactDetail definition: specifies contact information for a person or organization.
type FHIRContactDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The name of an individual to contact.
	Name *string `json:"name,omitempty"`

	// The contact details for the individual (if a name was provided) or the organization.
	Telecom []*FHIRContactPoint `json:"telecom,omitempty"`
}

// FHIRContactDetailInput is the input type for ContactDetail
type FHIRContactDetailInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The name of an individual to contact.
	Name *string `json:"name,omitempty"`

	// The contact details for the individual (if a name was provided) or the organization.
	Telecom []*FHIRContactPointInput `json:"telecom,omitempty"`
}

// FHIRContactPoint definition: details for all kinds of technology mediated contact points for a person or organization, including telephone, email, etc.
type FHIRContactPoint struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// Telecommunications form for contact point - what communications system is required to make use of the contact.
	System *ContactPointSystemEnum `json:"system,omitempty"`

	// The actual contact point details, in a form that is meaningful to the designated communication system (i.e. phone number or email address).
	Value *string `json:"value,omitempty"`

	// Identifies the purpose for the contact point.
	Use *ContactPointUseEnum `json:"use,omitempty"`

	// Specifies a preferred order in which to use a set of contacts. ContactPoints with lower rank values are more preferred than those with higher rank values.
	Rank *int64 `json:"rank,omitempty"`

	// Time period when the contact point was/is in use.
	Period *FHIRPeriod `json:"period,omitempty"`
}

// FHIRContactPointInput is the input type for ContactPoint
type FHIRContactPointInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// Telecommunications form for contact point - what communications system is required to make use of the contact.
	System *ContactPointSystemEnum `json:"system,omitempty"`

	// The actual contact point details, in a form that is meaningful to the designated communication system (i.e. phone number or email address).
	Value *string `json:"value,omitempty"`

	// Identifies the purpose for the contact point.
	Use *ContactPointUseEnum `json:"use,omitempty"`

	// Specifies a preferred order in which to use a set of contacts. ContactPoints with lower rank values are more preferred than those with higher rank values.
	Rank *int64 `json:"rank,omitempty"`

	// Time period when the contact point was/is in use.
	Period *FHIRPeriodInput `json:"period,omitempty"`
}

// FHIRContributor definition: a contributor to the content of a knowledge asset, including authors, editors, reviewers, and endorsers.
type FHIRContributor struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The type of contributor.
	Type *ContributorTypeEnum `json:"type,omitempty"`

	// The name of the individual or organization responsible for the contribution.
	Name *string `json:"name,omitempty"`

	// Contact details to assist a user in finding and communicating with the contributor.
	Contact []*FHIRContactDetail `json:"contact,omitempty"`
}

// FHIRContributorInput is the input type for Contributor
type FHIRContributorInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The type of contributor.
	Type *ContributorTypeEnum `json:"type,omitempty"`

	// The name of the individual or organization responsible for the contribution.
	Name *string `json:"name,omitempty"`

	// Contact details to assist a user in finding and communicating with the contributor.
	Contact []*FHIRContactDetailInput `json:"contact,omitempty"`
}

// FHIRCount definition: a measured amount (or an amount that can potentially be measured). note that measured amounts include amounts that are not precisely quantified, including amounts involving arbitrary units and floating currencies.
type FHIRCount struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Value *base.Decimal `json:"value,omitempty"`

	// How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Comparator *CountComparatorEnum `json:"comparator,omitempty"`

	// A human-readable form of the unit.
	Unit *string `json:"unit,omitempty"`

	// The identification of the system that provides the coded form of the unit.
	System *base.URI `json:"system,omitempty"`

	// A computer processable form of the unit in some unit representation system.
	Code *base.Code `json:"code,omitempty"`
}

// FHIRCountInput is the input type for Count
type FHIRCountInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Value *base.Decimal `json:"value,omitempty"`

	// How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Comparator *CountComparatorEnum `json:"comparator,omitempty"`

	// A human-readable form of the unit.
	Unit *string `json:"unit,omitempty"`

	// The identification of the system that provides the coded form of the unit.
	System *base.URI `json:"system,omitempty"`

	// A computer processable form of the unit in some unit representation system.
	Code *base.Code `json:"code,omitempty"`
}

// FHIRDataRequirement definition: describes a required data item for evaluation in terms of the type of data, and optional code or date-based filters of the data.
type FHIRDataRequirement struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The type of the required data, specified as the type name of a resource. For profiles, this value is set to the type of the base resource of the profile.
	Type *base.Code `json:"type,omitempty"`

	// The profile of the required data, specified as the uri of the profile definition.
	Profile []*base.Canonical `json:"profile,omitempty"`

	// The intended subjects of the data requirement. If this element is not provided, a Patient subject is assumed.
	SubjectCodeableConcept *base.Code `json:"subjectCodeableConcept,omitempty"`

	// The intended subjects of the data requirement. If this element is not provided, a Patient subject is assumed.
	SubjectReference *FHIRReference `json:"subjectReference,omitempty"`

	//     Indicates that specific elements of the type are referenced by the knowledge module and must be supported by the consumer in order to obtain an effective evaluation. This does not mean that a value is required for this element, only that the consuming system must understand the element and be able to provide values for it if they are available.
	//
	// The value of mustSupport SHALL be a FHIRPath resolveable on the type of the DataRequirement. The path SHALL consist only of identifiers, constant indexers, and .resolve() (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details).
	MustSupport []*string `json:"mustSupport,omitempty"`

	// Code filters specify additional constraints on the data, specifying the value set of interest for a particular element of the data. Each code filter defines an additional constraint on the data, i.e. code filters are AND'ed, not OR'ed.
	CodeFilter *base.Code `json:"codeFilter,omitempty"`

	// Date filters specify additional constraints on the data in terms of the applicable date range for specific elements. Each date filter specifies an additional constraint on the data, i.e. date filters are AND'ed, not OR'ed.
	DateFilter *base.Date `json:"dateFilter,omitempty"`

	// Specifies a maximum number of results that are required (uses the _count search parameter).
	Limit *string `json:"limit,omitempty"`

	// Specifies the order of the results to be returned.
	Sort []*FHIRDatarequirementSort `json:"sort,omitempty"`
}

// FHIRDataRequirementInput is the input type for DataRequirement
type FHIRDataRequirementInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The type of the required data, specified as the type name of a resource. For profiles, this value is set to the type of the base resource of the profile.
	Type *base.Code `json:"type,omitempty"`

	// The profile of the required data, specified as the uri of the profile definition.
	Profile *base.Canonical `json:"profile,omitempty"`

	// The intended subjects of the data requirement. If this element is not provided, a Patient subject is assumed.
	SubjectCodeableConcept *base.Code `json:"subjectCodeableConcept,omitempty"`

	// The intended subjects of the data requirement. If this element is not provided, a Patient subject is assumed.
	SubjectReference *FHIRReferenceInput `json:"subjectReference,omitempty"`

	//     Indicates that specific elements of the type are referenced by the knowledge module and must be supported by the consumer in order to obtain an effective evaluation. This does not mean that a value is required for this element, only that the consuming system must understand the element and be able to provide values for it if they are available.
	//
	// The value of mustSupport SHALL be a FHIRPath resolveable on the type of the DataRequirement. The path SHALL consist only of identifiers, constant indexers, and .resolve() (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details).
	MustSupport *string `json:"mustSupport,omitempty"`

	// Code filters specify additional constraints on the data, specifying the value set of interest for a particular element of the data. Each code filter defines an additional constraint on the data, i.e. code filters are AND'ed, not OR'ed.
	CodeFilter *base.Code `json:"codeFilter,omitempty"`

	// Date filters specify additional constraints on the data in terms of the applicable date range for specific elements. Each date filter specifies an additional constraint on the data, i.e. date filters are AND'ed, not OR'ed.
	DateFilter *base.Date `json:"dateFilter,omitempty"`

	// Specifies a maximum number of results that are required (uses the _count search parameter).
	Limit *string `json:"limit,omitempty"`

	// Specifies the order of the results to be returned.
	Sort []*FHIRDatarequirementSortInput `json:"sort,omitempty"`
}

// FHIRDatarequirementCodefilter definition: describes a required data item for evaluation in terms of the type of data, and optional code or date-based filters of the data.
type FHIRDatarequirementCodefilter struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The code-valued attribute of the filter. The specified path SHALL be a FHIRPath resolveable on the specified type of the DataRequirement, and SHALL consist only of identifiers, constant indexers, and .resolve(). The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details). Note that the index must be an integer constant. The path must resolve to an element of type code, Coding, or CodeableConcept.
	Path *string `json:"path,omitempty"`

	// A token parameter that refers to a search parameter defined on the specified type of the DataRequirement, and which searches on elements of type code, Coding, or CodeableConcept.
	SearchParam *string `json:"searchParam,omitempty"`

	// The valueset for the code filter. The valueSet and code elements are additive. If valueSet is specified, the filter will return only those data items for which the value of the code-valued element specified in the path is a member of the specified valueset.
	ValueSet *base.Canonical `json:"valueSet,omitempty"`

	// The codes for the code filter. If values are given, the filter will return only those data items for which the code-valued attribute specified by the path has a value that is one of the specified codes. If codes are specified in addition to a value set, the filter returns items matching a code in the value set or one of the specified codes.
	Code *base.Code `json:"code,omitempty"`
}

// FHIRDatarequirementCodefilterInput is the input type for DatarequirementCodefilter
type FHIRDatarequirementCodefilterInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The code-valued attribute of the filter. The specified path SHALL be a FHIRPath resolveable on the specified type of the DataRequirement, and SHALL consist only of identifiers, constant indexers, and .resolve(). The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details). Note that the index must be an integer constant. The path must resolve to an element of type code, Coding, or CodeableConcept.
	Path *string `json:"path,omitempty"`

	// A token parameter that refers to a search parameter defined on the specified type of the DataRequirement, and which searches on elements of type code, Coding, or CodeableConcept.
	SearchParam *string `json:"searchParam,omitempty"`

	// The valueset for the code filter. The valueSet and code elements are additive. If valueSet is specified, the filter will return only those data items for which the value of the code-valued element specified in the path is a member of the specified valueset.
	ValueSet *base.Canonical `json:"valueSet,omitempty"`

	// The codes for the code filter. If values are given, the filter will return only those data items for which the code-valued attribute specified by the path has a value that is one of the specified codes. If codes are specified in addition to a value set, the filter returns items matching a code in the value set or one of the specified codes.
	Code *base.Code `json:"code,omitempty"`
}

// FHIRDatarequirementDatefilter definition: describes a required data item for evaluation in terms of the type of data, and optional code or date-based filters of the data.
type FHIRDatarequirementDatefilter struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The date-valued attribute of the filter. The specified path SHALL be a FHIRPath resolveable on the specified type of the DataRequirement, and SHALL consist only of identifiers, constant indexers, and .resolve(). The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details). Note that the index must be an integer constant. The path must resolve to an element of type date, dateTime, Period, Schedule, or Timing.
	Path *string `json:"path,omitempty"`

	// A date parameter that refers to a search parameter defined on the specified type of the DataRequirement, and which searches on elements of type date, dateTime, Period, Schedule, or Timing.
	SearchParam *string `json:"searchParam,omitempty"`

	// The value of the filter. If period is specified, the filter will return only those data items that fall within the bounds determined by the Period, inclusive of the period boundaries. If dateTime is specified, the filter will return only those data items that are equal to the specified dateTime. If a Duration is specified, the filter will return only those data items that fall within Duration before now.
	ValueDateTime *base.Date `json:"valueDateTime,omitempty"`

	// The value of the filter. If period is specified, the filter will return only those data items that fall within the bounds determined by the Period, inclusive of the period boundaries. If dateTime is specified, the filter will return only those data items that are equal to the specified dateTime. If a Duration is specified, the filter will return only those data items that fall within Duration before now.
	ValuePeriod *FHIRPeriod `json:"valuePeriod,omitempty"`

	// The value of the filter. If period is specified, the filter will return only those data items that fall within the bounds determined by the Period, inclusive of the period boundaries. If dateTime is specified, the filter will return only those data items that are equal to the specified dateTime. If a Duration is specified, the filter will return only those data items that fall within Duration before now.
	ValueDuration *FHIRDuration `json:"valueDuration,omitempty"`
}

// FHIRDatarequirementDatefilterInput is the input type for DatarequirementDatefilter
type FHIRDatarequirementDatefilterInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The date-valued attribute of the filter. The specified path SHALL be a FHIRPath resolveable on the specified type of the DataRequirement, and SHALL consist only of identifiers, constant indexers, and .resolve(). The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details). Note that the index must be an integer constant. The path must resolve to an element of type date, dateTime, Period, Schedule, or Timing.
	Path *string `json:"path,omitempty"`

	// A date parameter that refers to a search parameter defined on the specified type of the DataRequirement, and which searches on elements of type date, dateTime, Period, Schedule, or Timing.
	SearchParam *string `json:"searchParam,omitempty"`

	// The value of the filter. If period is specified, the filter will return only those data items that fall within the bounds determined by the Period, inclusive of the period boundaries. If dateTime is specified, the filter will return only those data items that are equal to the specified dateTime. If a Duration is specified, the filter will return only those data items that fall within Duration before now.
	ValueDateTime *base.Date `json:"valueDateTime,omitempty"`

	// The value of the filter. If period is specified, the filter will return only those data items that fall within the bounds determined by the Period, inclusive of the period boundaries. If dateTime is specified, the filter will return only those data items that are equal to the specified dateTime. If a Duration is specified, the filter will return only those data items that fall within Duration before now.
	ValuePeriod *FHIRPeriodInput `json:"valuePeriod,omitempty"`

	// The value of the filter. If period is specified, the filter will return only those data items that fall within the bounds determined by the Period, inclusive of the period boundaries. If dateTime is specified, the filter will return only those data items that are equal to the specified dateTime. If a Duration is specified, the filter will return only those data items that fall within Duration before now.
	ValueDuration *FHIRDurationInput `json:"valueDuration,omitempty"`
}

// FHIRDatarequirementSort definition: describes a required data item for evaluation in terms of the type of data, and optional code or date-based filters of the data.
type FHIRDatarequirementSort struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The attribute of the sort. The specified path must be resolvable from the type of the required data. The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements. Note that the index must be an integer constant.
	Path *string `json:"path,omitempty"`

	// The direction of the sort, ascending or descending.
	Direction *DataRequirementSortDirectionEnum `json:"direction,omitempty"`
}

// FHIRDatarequirementSortInput is the input type for DatarequirementSort
type FHIRDatarequirementSortInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The attribute of the sort. The specified path must be resolvable from the type of the required data. The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements. Note that the index must be an integer constant.
	Path *string `json:"path,omitempty"`

	// The direction of the sort, ascending or descending.
	Direction *DataRequirementSortDirectionEnum `json:"direction,omitempty"`
}

// FHIRDistance definition: a length - a value with a unit that is a physical distance.
type FHIRDistance struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Value *base.Decimal `json:"value,omitempty"`

	// How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Comparator *DistanceComparatorEnum `json:"comparator,omitempty"`

	// A human-readable form of the unit.
	Unit *string `json:"unit,omitempty"`

	// The identification of the system that provides the coded form of the unit.
	System *base.URI `json:"system,omitempty"`

	// A computer processable form of the unit in some unit representation system.
	Code *base.Code `json:"code,omitempty"`
}

// FHIRDistanceInput is the input type for Distance
type FHIRDistanceInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Value *base.Decimal `json:"value,omitempty"`

	// How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Comparator *DistanceComparatorEnum `json:"comparator,omitempty"`

	// A human-readable form of the unit.
	Unit *string `json:"unit,omitempty"`

	// The identification of the system that provides the coded form of the unit.
	System *base.URI `json:"system,omitempty"`

	// A computer processable form of the unit in some unit representation system.
	Code *base.Code `json:"code,omitempty"`
}

// FHIRDosage definition: indicates how the medication is/was taken or should be taken by the patient.
type FHIRDosage struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// Indicates the order in which the dosage instructions should be applied or interpreted.
	Sequence *string `json:"sequence,omitempty"`

	// Free text dosage instructions e.g. SIG.
	Text *string `json:"text,omitempty"`

	// Supplemental instructions to the patient on how to take the medication  (e.g. "with meals" or"take half to one hour before food") or warnings for the patient about the medication (e.g. "may cause drowsiness" or "avoid exposure of skin to direct sunlight or sunlamps").
	AdditionalInstruction []*FHIRCodeableConcept `json:"additionalInstruction,omitempty"`

	// Instructions in terms that are understood by the patient or consumer.
	PatientInstruction *string `json:"patientInstruction,omitempty"`

	// When medication should be administered.
	Timing *FHIRTiming `json:"timing,omitempty"`

	// Indicates whether the Medication is only taken when needed within a specific dosing schedule (Boolean option), or it indicates the precondition for taking the Medication (CodeableConcept).
	AsNeededBoolean *bool `json:"asNeededBoolean,omitempty"`

	// Indicates whether the Medication is only taken when needed within a specific dosing schedule (Boolean option), or it indicates the precondition for taking the Medication (CodeableConcept).
	AsNeededCodeableConcept *base.Code `json:"asNeededCodeableConcept,omitempty"`

	// Body site to administer to.
	Site *FHIRCodeableConcept `json:"site,omitempty"`

	// How drug should enter body.
	Route *FHIRCodeableConcept `json:"route,omitempty"`

	// Technique for administering medication.
	Method *FHIRCodeableConcept `json:"method,omitempty"`

	// The amount of medication administered.
	DoseAndRate []*FHIRDosageDoseandrate `json:"doseAndRate,omitempty"`

	// Upper limit on medication per unit of time.
	MaxDosePerPeriod *FHIRRatio `json:"maxDosePerPeriod,omitempty"`

	// Upper limit on medication per administration.
	MaxDosePerAdministration *FHIRQuantity `json:"maxDosePerAdministration,omitempty"`

	// Upper limit on medication per lifetime of the patient.
	MaxDosePerLifetime *FHIRQuantity `json:"maxDosePerLifetime,omitempty"`
}

// FHIRDosageDoseandrate definition: indicates how the medication is/was taken or should be taken by the patient.
type FHIRDosageDoseandrate struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The kind of dose or rate specified, for example, ordered or calculated.
	Type *FHIRCodeableConcept `json:"type,omitempty"`

	// Amount of medication per dose.
	DoseRange *FHIRRange `json:"doseRange,omitempty"`

	// Amount of medication per dose.
	DoseQuantity *FHIRQuantity `json:"doseQuantity,omitempty"`

	// Amount of medication per unit of time.
	RateRatio *FHIRRatio `json:"rateRatio,omitempty"`

	// Amount of medication per unit of time.
	RateRange *FHIRRange `json:"rateRange,omitempty"`

	// Amount of medication per unit of time.
	RateQuantity *FHIRQuantity `json:"rateQuantity,omitempty"`
}

// FHIRDosageDoseandrateInput is the input type for DosageDoseandrate
type FHIRDosageDoseandrateInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The kind of dose or rate specified, for example, ordered or calculated.
	Type *FHIRCodeableConceptInput `json:"type,omitempty"`

	// Amount of medication per dose.
	DoseRange *FHIRRangeInput `json:"doseRange,omitempty"`

	// Amount of medication per dose.
	DoseQuantity *FHIRQuantityInput `json:"doseQuantity,omitempty"`

	// Amount of medication per unit of time.
	RateRatio *FHIRRatioInput `json:"rateRatio,omitempty"`

	// Amount of medication per unit of time.
	RateRange *FHIRRangeInput `json:"rateRange,omitempty"`

	// Amount of medication per unit of time.
	RateQuantity *FHIRQuantityInput `json:"rateQuantity,omitempty"`
}

// FHIRDosageInput is the input type for Dosage
type FHIRDosageInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// Indicates the order in which the dosage instructions should be applied or interpreted.
	Sequence *string `json:"sequence,omitempty"`

	// Free text dosage instructions e.g. SIG.
	Text *string `json:"text,omitempty"`

	// Supplemental instructions to the patient on how to take the medication  (e.g. "with meals" or"take half to one hour before food") or warnings for the patient about the medication (e.g. "may cause drowsiness" or "avoid exposure of skin to direct sunlight or sunlamps").
	AdditionalInstruction []*FHIRCodeableConceptInput `json:"additionalInstruction,omitempty"`

	// Instructions in terms that are understood by the patient or consumer.
	PatientInstruction *string `json:"patientInstruction,omitempty"`

	// When medication should be administered.
	Timing *FHIRTimingInput `json:"timing,omitempty"`

	// Indicates whether the Medication is only taken when needed within a specific dosing schedule (Boolean option), or it indicates the precondition for taking the Medication (CodeableConcept).
	AsNeededBoolean *bool `json:"asNeededBoolean,omitempty"`

	// Indicates whether the Medication is only taken when needed within a specific dosing schedule (Boolean option), or it indicates the precondition for taking the Medication (CodeableConcept).
	AsNeededCodeableConcept *base.Code `json:"asNeededCodeableConcept,omitempty"`

	// Body site to administer to.
	Site *FHIRCodeableConceptInput `json:"site,omitempty"`

	// How drug should enter body.
	Route *FHIRCodeableConceptInput `json:"route,omitempty"`

	// Technique for administering medication.
	Method *FHIRCodeableConceptInput `json:"method,omitempty"`

	// The amount of medication administered.
	DoseAndRate []*FHIRDosageDoseandrateInput `json:"doseAndRate,omitempty"`

	// Upper limit on medication per unit of time.
	MaxDosePerPeriod *FHIRRatioInput `json:"maxDosePerPeriod,omitempty"`

	// Upper limit on medication per administration.
	MaxDosePerAdministration *FHIRQuantityInput `json:"maxDosePerAdministration,omitempty"`

	// Upper limit on medication per lifetime of the patient.
	MaxDosePerLifetime *FHIRQuantityInput `json:"maxDosePerLifetime,omitempty"`
}

// FHIRDuration definition: a length of time.
type FHIRDuration struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Value *base.Decimal `json:"value,omitempty"`

	// How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Comparator *DurationComparatorEnum `json:"comparator,omitempty"`

	// A human-readable form of the unit.
	Unit *string `json:"unit,omitempty"`

	// The identification of the system that provides the coded form of the unit.
	System *base.URI `json:"system,omitempty"`

	// A computer processable form of the unit in some unit representation system.
	Code *base.Code `json:"code,omitempty"`
}

// FHIRDurationInput is the input type for Duration
type FHIRDurationInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Value *base.Decimal `json:"value,omitempty"`

	// How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Comparator *DurationComparatorEnum `json:"comparator,omitempty"`

	// A human-readable form of the unit.
	Unit *string `json:"unit,omitempty"`

	// The identification of the system that provides the coded form of the unit.
	System *base.URI `json:"system,omitempty"`

	// A computer processable form of the unit in some unit representation system.
	Code *base.Code `json:"code,omitempty"`
}

// FHIRHumanName definition: a human's name with the ability to identify parts and usage.
type FHIRHumanName struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// Identifies the purpose for this name.
	Use HumanNameUseEnum `json:"use,omitempty"`

	// Specifies the entire name as it should be displayed e.g. on an application UI. This may be provided instead of or as well as the specific parts.
	Text string `json:"text,omitempty"`

	// The part of a name that links to the genealogy. In some cultures (e.g. Eritrea) the family name of a son is the first name of his father.
	Family *string `json:"family,omitempty"`

	// Given name.
	Given []*string `json:"given,omitempty"`

	// Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the start of the name.
	Prefix []*string `json:"prefix,omitempty"`

	// Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the end of the name.
	Suffix []*string `json:"suffix,omitempty"`

	// Indicates the period of time when this name was valid for the named person.
	Period *FHIRPeriod `json:"period,omitempty"`
}

// FHIRHumanNameInput is the input type for HumanName
type FHIRHumanNameInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// Identifies the purpose for this name.
	Use HumanNameUseEnum `json:"use,omitempty"`

	// Specifies the entire name as it should be displayed e.g. on an application UI. This may be provided instead of or as well as the specific parts.
	Text string `json:"text,omitempty"`

	// The part of a name that links to the genealogy. In some cultures (e.g. Eritrea) the family name of a son is the first name of his father.
	Family *string `json:"family,omitempty"`

	// Given name.
	Given *string `json:"given,omitempty"`

	// Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the start of the name.
	Prefix *string `json:"prefix,omitempty"`

	// Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the end of the name.
	Suffix *string `json:"suffix,omitempty"`

	// Indicates the period of time when this name was valid for the named person.
	Period *FHIRPeriodInput `json:"period,omitempty"`
}

// FHIRIdentifier definition: an identifier - identifies some entity uniquely and unambiguously. typically this is used for business identifiers.
type FHIRIdentifier struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The purpose of this identifier.
	Use IdentifierUseEnum `json:"use,omitempty"`

	// A coded type for the identifier that can be used to determine which identifier to use for a specific purpose.
	Type FHIRCodeableConcept `json:"type,omitempty"`

	// Establishes the namespace for the value - that is, a URL that describes a set values that are unique.
	System *base.URI `json:"system,omitempty"`

	// The portion of the identifier typically relevant to the user and which is unique within the context of the system.
	Value string `json:"value,omitempty"`

	// Time period during which identifier is/was valid for use.
	Period *FHIRPeriod `json:"period,omitempty"`

	// Organization that issued/manages the identifier.
	Assigner *FHIRReference `json:"assigner,omitempty"`
}

// FHIRIdentifierInput is the input type for Identifier
type FHIRIdentifierInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The purpose of this identifier.
	Use IdentifierUseEnum `json:"use,omitempty"`

	// A coded type for the identifier that can be used to determine which identifier to use for a specific purpose.
	Type FHIRCodeableConceptInput `json:"type,omitempty"`

	// Establishes the namespace for the value - that is, a URL that describes a set values that are unique.
	System *base.URI `json:"system,omitempty"`

	// The portion of the identifier typically relevant to the user and which is unique within the context of the system.
	Value string `json:"value,omitempty"`

	// Time period during which identifier is/was valid for use.
	Period *FHIRPeriodInput `json:"period,omitempty"`

	// Organization that issued/manages the identifier.
	Assigner *FHIRReferenceInput `json:"assigner,omitempty"`
}

// FHIRMoney definition: an amount of economic utility in some recognized currency.
type FHIRMoney struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// Numerical value (with implicit precision).
	Value *base.Decimal `json:"value,omitempty"`

	// ISO 4217 Currency Code.
	Currency *base.Code `json:"currency,omitempty"`
}

// FHIRMoneyInput is the input type for Money
type FHIRMoneyInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// Numerical value (with implicit precision).
	Value *base.Decimal `json:"value,omitempty"`

	// ISO 4217 Currency Code.
	Currency *base.Code `json:"currency,omitempty"`
}

// FHIRNarrative definition: a human-readable summary of the resource conveying the essential clinical and business information for the resource.
type FHIRNarrative struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The status of the narrative - whether it's entirely generated (from just the defined data or the extensions too), or whether a human authored it and it may contain additional data.
	Status *NarrativeStatusEnum `json:"status,omitempty"`

	// The actual narrative content, a stripped down version of XHTML.
	Div base.XHTML `json:"div,omitempty"`
}

// FHIRNarrativeInput is the input type for Narrative
type FHIRNarrativeInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The status of the narrative - whether it's entirely generated (from just the defined data or the extensions too), or whether a human authored it and it may contain additional data.
	Status *NarrativeStatusEnum `json:"status,omitempty"`

	// The actual narrative content, a stripped down version of XHTML.
	Div base.XHTML `json:"div,omitempty"`
}

// FHIRPeriod definition: a time period defined by a start and end date and optionally time.
type FHIRPeriod struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The start of the period. The boundary is inclusive.
	Start base.DateTime `json:"start,omitempty"`

	// The end of the period. If the end of the period is missing, it means no end was known or planned at the time the instance was created. The start may be in the past, and the end date in the future, which means that period is expected/planned to end at that time.
	End base.DateTime `json:"end,omitempty"`
}

// FHIRPeriodInput is the input type for Period
type FHIRPeriodInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The start of the period. The boundary is inclusive.
	Start base.DateTime `json:"start,omitempty"`

	// The end of the period. If the end of the period is missing, it means no end was known or planned at the time the instance was created. The start may be in the past, and the end date in the future, which means that period is expected/planned to end at that time.
	End base.DateTime `json:"end,omitempty"`
}

// FHIRQuantity definition: a measured amount (or an amount that can potentially be measured). note that measured amounts include amounts that are not precisely quantified, including amounts involving arbitrary units and floating currencies.
type FHIRQuantity struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Value float64 `json:"value"`

	// How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Comparator *QuantityComparatorEnum `json:"comparator,omitempty"`

	// A human-readable form of the unit.
	Unit string `json:"unit"`

	// The identification of the system that provides the coded form of the unit.
	System base.URI `json:"system"`

	// A computer processable form of the unit in some unit representation system.
	Code base.Code `json:"code"`
}

// FHIRQuantityInput is the input type for Quantity
type FHIRQuantityInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Value float64 `json:"value"`

	// How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Comparator *QuantityComparatorEnum `json:"comparator,omitempty"`

	// A human-readable form of the unit.
	Unit string `json:"unit"`

	// The identification of the system that provides the coded form of the unit.
	System base.URI `json:"system"`

	// A computer processable form of the unit in some unit representation system.
	Code base.Code `json:"code"`
}

// FHIRRange definition: a set of ordered quantities defined by a low and high limit.
type FHIRRange struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The low limit. The boundary is inclusive.
	Low FHIRQuantity `json:"low,omitempty"`

	// The high limit. The boundary is inclusive.
	High FHIRQuantity `json:"high,omitempty"`
}

// FHIRRangeInput is the input type for Range
type FHIRRangeInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The low limit. The boundary is inclusive.
	Low FHIRQuantityInput `json:"low,omitempty"`

	// The high limit. The boundary is inclusive.
	High FHIRQuantityInput `json:"high,omitempty"`
}

// FHIRRatio definition: a relationship of two quantity values - expressed as a numerator and a denominator.
type FHIRRatio struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The value of the numerator.
	Numerator FHIRQuantity `json:"numerator,omitempty"`

	// The value of the denominator.
	Denominator FHIRQuantity `json:"denominator,omitempty"`
}

// FHIRRatioInput is the input type for Ratio
type FHIRRatioInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The value of the numerator.
	Numerator FHIRQuantityInput `json:"numerator,omitempty"`

	// The value of the denominator.
	Denominator FHIRQuantityInput `json:"denominator,omitempty"`
}

// FHIRReference definition: a reference from one resource to another.
type FHIRReference struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// A reference to a location at which the other resource is found. The reference may be a relative reference, in which case it is relative to the service base URL, or an absolute URL that resolves to the location where the resource is found. The reference may be version specific or not. If the reference is not to a FHIR RESTful server, then it should be assumed to be version specific. Internal fragment references (start with '#') refer to contained resources.
	Reference *string `json:"reference,omitempty"`

	//     The expected type of the target of the reference. If both Reference.type and Reference.reference are populated and Reference.reference is a FHIR URL, both SHALL be consistent.
	//
	// The type is the Canonical URL of Resource Definition that is the type this reference refers to. References are URLs that are relative to http://hl7.org/fhir/StructureDefinition/ e.g. "Patient" is a reference to http://hl7.org/fhir/StructureDefinition/Patient. Absolute URLs are only allowed for logical models (and can only be used in references in logical models, not resources).
	Type *base.URI `json:"type,omitempty"`

	// An identifier for the target resource. This is used when there is no way to reference the other resource directly, either because the entity it represents is not available through a FHIR server, or because there is no way for the author of the resource to convert a known identifier to an actual location. There is no requirement that a Reference.identifier point to something that is actually exposed as a FHIR instance, but it SHALL point to a business concept that would be expected to be exposed as a FHIR instance, and that instance would need to be of a FHIR resource type allowed by the reference.
	Identifier *FHIRIdentifier `json:"identifier,omitempty"`

	// Plain text narrative that identifies the resource in addition to the resource reference.
	Display string `json:"display,omitempty"`
}

// FHIRReferenceInput is the input type for Reference
type FHIRReferenceInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// A reference to a location at which the other resource is found. The reference may be a relative reference, in which case it is relative to the service base URL, or an absolute URL that resolves to the location where the resource is found. The reference may be version specific or not. If the reference is not to a FHIR RESTful server, then it should be assumed to be version specific. Internal fragment references (start with '#') refer to contained resources.
	Reference *string `json:"reference,omitempty"`

	//     The expected type of the target of the reference. If both Reference.type and Reference.reference are populated and Reference.reference is a FHIR URL, both SHALL be consistent.
	//
	// The type is the Canonical URL of Resource Definition that is the type this reference refers to. References are URLs that are relative to http://hl7.org/fhir/StructureDefinition/ e.g. "Patient" is a reference to http://hl7.org/fhir/StructureDefinition/Patient. Absolute URLs are only allowed for logical models (and can only be used in references in logical models, not resources).
	Type *base.URI `json:"type,omitempty"`

	// An identifier for the target resource. This is used when there is no way to reference the other resource directly, either because the entity it represents is not available through a FHIR server, or because there is no way for the author of the resource to convert a known identifier to an actual location. There is no requirement that a Reference.identifier point to something that is actually exposed as a FHIR instance, but it SHALL point to a business concept that would be expected to be exposed as a FHIR instance, and that instance would need to be of a FHIR resource type allowed by the reference.
	Identifier *FHIRIdentifierInput `json:"identifier,omitempty"`

	// Plain text narrative that identifies the resource in addition to the resource reference.
	Display string `json:"display,omitempty"`
}

// FHIRSampledData definition: a series of measurements taken by a device, with upper and lower limits. there may be more than one dimension in the data.
type FHIRSampledData struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The base quantity that a measured value of zero represents. In addition, this provides the units of the entire measurement series.
	Origin *FHIRQuantity `json:"origin,omitempty"`

	// The length of time between sampling times, measured in milliseconds.
	Period *base.Decimal `json:"period,omitempty"`

	// A correction factor that is applied to the sampled data points before they are added to the origin.
	Factor *base.Decimal `json:"factor,omitempty"`

	// The lower limit of detection of the measured points. This is needed if any of the data points have the value "L" (lower than detection limit).
	LowerLimit *base.Decimal `json:"lowerLimit,omitempty"`

	// The upper limit of detection of the measured points. This is needed if any of the data points have the value "U" (higher than detection limit).
	UpperLimit *base.Decimal `json:"upperLimit,omitempty"`

	// The number of sample points at each time point. If this value is greater than one, then the dimensions will be interlaced - all the sample points for a point in time will be recorded at once.
	Dimensions *string `json:"dimensions,omitempty"`

	// A series of data points which are decimal values separated by a single space (character u20). The special values "E" (error), "L" (below detection limit) and "U" (above detection limit) can also be used in place of a decimal value.
	Data *string `json:"data,omitempty"`
}

// FHIRSampledDataInput is the input type for SampledData
type FHIRSampledDataInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// The base quantity that a measured value of zero represents. In addition, this provides the units of the entire measurement series.
	Origin *FHIRQuantityInput `json:"origin,omitempty"`

	// The length of time between sampling times, measured in milliseconds.
	Period *base.Decimal `json:"period,omitempty"`

	// A correction factor that is applied to the sampled data points before they are added to the origin.
	Factor *base.Decimal `json:"factor,omitempty"`

	// The lower limit of detection of the measured points. This is needed if any of the data points have the value "L" (lower than detection limit).
	LowerLimit *base.Decimal `json:"lowerLimit,omitempty"`

	// The upper limit of detection of the measured points. This is needed if any of the data points have the value "U" (higher than detection limit).
	UpperLimit *base.Decimal `json:"upperLimit,omitempty"`

	// The number of sample points at each time point. If this value is greater than one, then the dimensions will be interlaced - all the sample points for a point in time will be recorded at once.
	Dimensions *string `json:"dimensions,omitempty"`

	// A series of data points which are decimal values separated by a single space (character u20). The special values "E" (error), "L" (below detection limit) and "U" (above detection limit) can also be used in place of a decimal value.
	Data *string `json:"data,omitempty"`
}

// FHIRSignature definition: a signature along with supporting context. the signature may be a digital signature that is cryptographic in nature, or some other signature acceptable to the domain. this other signature may be as simple as a graphical image representing a hand-written signature, or a signature ceremony different signature approaches have different utilities.
type FHIRSignature struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// An indication of the reason that the entity signed this document. This may be explicitly included as part of the signature information and can be used when determining accountability for various actions concerning the document.
	Type []*FHIRCoding `json:"type,omitempty"`

	// When the digital signature was signed.
	When *base.Instant `json:"when,omitempty"`

	// A reference to an application-usable description of the identity that signed  (e.g. the signature used their private key).
	Who *FHIRReference `json:"who,omitempty"`

	// A reference to an application-usable description of the identity that is represented by the signature.
	OnBehalfOf *FHIRReference `json:"onBehalfOf,omitempty"`

	// A mime type that indicates the technical format of the target resources signed by the signature.
	TargetFormat *base.Code `json:"targetFormat,omitempty"`

	// A mime type that indicates the technical format of the signature. Important mime types are application/signature+xml for X ML DigSig, application/jose for JWS, and image/* for a graphical image of a signature, etc.
	SigFormat *base.Code `json:"sigFormat,omitempty"`

	// The base64 encoding of the Signature content. When signature is not recorded electronically this element would be empty.
	Data *base.Base64Binary `json:"data,omitempty"`
}

// FHIRSignatureInput is the input type for Signature
type FHIRSignatureInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// An indication of the reason that the entity signed this document. This may be explicitly included as part of the signature information and can be used when determining accountability for various actions concerning the document.
	Type []*FHIRCodingInput `json:"type,omitempty"`

	// When the digital signature was signed.
	When *base.Instant `json:"when,omitempty"`

	// A reference to an application-usable description of the identity that signed  (e.g. the signature used their private key).
	Who *FHIRReferenceInput `json:"who,omitempty"`

	// A reference to an application-usable description of the identity that is represented by the signature.
	OnBehalfOf *FHIRReferenceInput `json:"onBehalfOf,omitempty"`

	// A mime type that indicates the technical format of the target resources signed by the signature.
	TargetFormat *base.Code `json:"targetFormat,omitempty"`

	// A mime type that indicates the technical format of the signature. Important mime types are application/signature+xml for X ML DigSig, application/jose for JWS, and image/* for a graphical image of a signature, etc.
	SigFormat *base.Code `json:"sigFormat,omitempty"`

	// The base64 encoding of the Signature content. When signature is not recorded electronically this element would be empty.
	Data *base.Base64Binary `json:"data,omitempty"`
}

// FHIRTiming definition: specifies an event that may occur multiple times. timing schedules are used to record when things are planned, expected or requested to occur. the most common usage is in dosage instructions for medications. they are also used when planning care of various kinds, and may be used for reporting the schedule to which past regular activities were carried out.
type FHIRTiming struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// Identifies specific times when the event occurs.
	Event []*base.DateTime `json:"event,omitempty"`

	// A set of rules that describe when the event is scheduled.
	Repeat *FHIRTimingRepeat `json:"repeat,omitempty"`

	// A code for the timing schedule (or just text in code.text). Some codes such as BID are ubiquitous, but many institutions define their own additional codes. If a code is provided, the code is understood to be a complete statement of whatever is specified in the structured timing data, and either the code or the data may be used to interpret the Timing, with the exception that .repeat.bounds still applies over the code (and is not contained in the code).
	Code base.Code `json:"code,omitempty"`
}

// FHIRTimingInput is the input type for Timing
type FHIRTimingInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// Identifies specific times when the event occurs.
	Event *base.DateTime `json:"event,omitempty"`

	// A set of rules that describe when the event is scheduled.
	Repeat *FHIRTimingRepeatInput `json:"repeat,omitempty"`

	// A code for the timing schedule (or just text in code.text). Some codes such as BID are ubiquitous, but many institutions define their own additional codes. If a code is provided, the code is understood to be a complete statement of whatever is specified in the structured timing data, and either the code or the data may be used to interpret the Timing, with the exception that .repeat.bounds still applies over the code (and is not contained in the code).
	Code base.Code `json:"code,omitempty"`
}

// FHIRTimingRepeat definition: specifies an event that may occur multiple times. timing schedules are used to record when things are planned, expected or requested to occur. the most common usage is in dosage instructions for medications. they are also used when planning care of various kinds, and may be used for reporting the schedule to which past regular activities were carried out.
type FHIRTimingRepeat struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// Either a duration for the length of the timing schedule, a range of possible length, or outer bounds for start and/or end limits of the timing schedule.
	BoundsDuration *FHIRDuration `json:"boundsDuration,omitempty"`

	// Either a duration for the length of the timing schedule, a range of possible length, or outer bounds for start and/or end limits of the timing schedule.
	BoundsRange *FHIRRange `json:"boundsRange,omitempty"`

	// Either a duration for the length of the timing schedule, a range of possible length, or outer bounds for start and/or end limits of the timing schedule.
	BoundsPeriod *FHIRPeriod `json:"boundsPeriod,omitempty"`

	// A total count of the desired number of repetitions across the duration of the entire timing specification. If countMax is present, this element indicates the lower bound of the allowed range of count values.
	Count *string `json:"count,omitempty"`

	// If present, indicates that the count is a range - so to perform the action between [count] and [countMax] times.
	CountMax *string `json:"countMax,omitempty"`

	// How long this thing happens for when it happens. If durationMax is present, this element indicates the lower bound of the allowed range of the duration.
	Duration *base.Decimal `json:"duration,omitempty"`

	// If present, indicates that the duration is a range - so to perform the action between [duration] and [durationMax] time length.
	DurationMax *base.Decimal `json:"durationMax,omitempty"`

	// The units of time for the duration, in UCUM units.
	DurationUnit *TimingRepeatDurationUnitEnum `json:"durationUnit,omitempty"`

	// The number of times to repeat the action within the specified period. If frequencyMax is present, this element indicates the lower bound of the allowed range of the frequency.
	Frequency *string `json:"frequency,omitempty"`

	// If present, indicates that the frequency is a range - so to repeat between [frequency] and [frequencyMax] times within the period or period range.
	FrequencyMax *string `json:"frequencyMax,omitempty"`

	// Indicates the duration of time over which repetitions are to occur; e.g. to express "3 times per day", 3 would be the frequency and "1 day" would be the period. If periodMax is present, this element indicates the lower bound of the allowed range of the period length.
	Period *base.Decimal `json:"period,omitempty"`

	// If present, indicates that the period is a range from [period] to [periodMax], allowing expressing concepts such as "do this once every 3-5 days.
	PeriodMax *base.Decimal `json:"periodMax,omitempty"`

	// The units of time for the period in UCUM units.
	PeriodUnit *TimingRepeatPeriodUnitEnum `json:"periodUnit,omitempty"`

	// If one or more days of week is provided, then the action happens only on the specified day(s).
	DayOfWeek []*base.Code `json:"dayOfWeek,omitempty"`

	// Specified time of day for action to take place.
	TimeOfDay *time.Time `json:"timeOfDay,omitempty"`

	// An approximate time period during the day, potentially linked to an event of daily living that indicates when the action should occur.
	When *TimingRepeatWhenEnum `json:"when,omitempty"`

	// The number of minutes from the event. If the event code does not indicate whether the minutes is before or after the event, then the offset is assumed to be after the event.
	Offset *int `json:"offset,omitempty"`
}

// FHIRTimingRepeatInput is the input type for TimingRepeat
type FHIRTimingRepeatInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// Either a duration for the length of the timing schedule, a range of possible length, or outer bounds for start and/or end limits of the timing schedule.
	BoundsDuration *FHIRDurationInput `json:"boundsDuration,omitempty"`

	// Either a duration for the length of the timing schedule, a range of possible length, or outer bounds for start and/or end limits of the timing schedule.
	BoundsRange *FHIRRangeInput `json:"boundsRange,omitempty"`

	// Either a duration for the length of the timing schedule, a range of possible length, or outer bounds for start and/or end limits of the timing schedule.
	BoundsPeriod *FHIRPeriodInput `json:"boundsPeriod,omitempty"`

	// A total count of the desired number of repetitions across the duration of the entire timing specification. If countMax is present, this element indicates the lower bound of the allowed range of count values.
	Count *string `json:"count,omitempty"`

	// If present, indicates that the count is a range - so to perform the action between [count] and [countMax] times.
	CountMax *string `json:"countMax,omitempty"`

	// How long this thing happens for when it happens. If durationMax is present, this element indicates the lower bound of the allowed range of the duration.
	Duration *base.Decimal `json:"duration,omitempty"`

	// If present, indicates that the duration is a range - so to perform the action between [duration] and [durationMax] time length.
	DurationMax *base.Decimal `json:"durationMax,omitempty"`

	// The units of time for the duration, in UCUM units.
	DurationUnit *TimingRepeatDurationUnitEnum `json:"durationUnit,omitempty"`

	// The number of times to repeat the action within the specified period. If frequencyMax is present, this element indicates the lower bound of the allowed range of the frequency.
	Frequency *string `json:"frequency,omitempty"`

	// If present, indicates that the frequency is a range - so to repeat between [frequency] and [frequencyMax] times within the period or period range.
	FrequencyMax *string `json:"frequencyMax,omitempty"`

	// Indicates the duration of time over which repetitions are to occur; e.g. to express "3 times per day", 3 would be the frequency and "1 day" would be the period. If periodMax is present, this element indicates the lower bound of the allowed range of the period length.
	Period *base.Decimal `json:"period,omitempty"`

	// If present, indicates that the period is a range from [period] to [periodMax], allowing expressing concepts such as "do this once every 3-5 days.
	PeriodMax *base.Decimal `json:"periodMax,omitempty"`

	// The units of time for the period in UCUM units.
	PeriodUnit *TimingRepeatPeriodUnitEnum `json:"periodUnit,omitempty"`

	// If one or more days of week is provided, then the action happens only on the specified day(s).
	DayOfWeek *base.Code `json:"dayOfWeek,omitempty"`

	// Specified time of day for action to take place.
	TimeOfDay *time.Time `json:"timeOfDay,omitempty"`

	// An approximate time period during the day, potentially linked to an event of daily living that indicates when the action should occur.
	When *TimingRepeatWhenEnum `json:"when,omitempty"`

	// The number of minutes from the event. If the event code does not indicate whether the minutes is before or after the event, then the offset is assumed to be after the event.
	Offset *int `json:"offset,omitempty"`
}

// FHIRUsageContext definition: specifies clinical/business/etc. metadata that can be used to retrieve, index and/or categorize an artifact. this metadata can either be specific to the applicable population (e.g., age category, drg) or the specific context of care (e.g., venue, care setting, provider of care).
type FHIRUsageContext struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// A code that identifies the type of context being specified by this usage context.
	Code base.Code `json:"code,omitempty"`

	// A value that defines the context specified in this context of use. The interpretation of the value is defined by the code.
	ValueCodeableConcept *base.Code `json:"valueCodeableConcept,omitempty"`

	// A value that defines the context specified in this context of use. The interpretation of the value is defined by the code.
	ValueQuantity *FHIRQuantity `json:"valueQuantity,omitempty"`

	// A value that defines the context specified in this context of use. The interpretation of the value is defined by the code.
	ValueRange *FHIRRange `json:"valueRange,omitempty"`

	// A value that defines the context specified in this context of use. The interpretation of the value is defined by the code.
	ValueReference *FHIRReference `json:"valueReference,omitempty"`
}

// FHIRUsageContextInput is the input type for UsageContext
type FHIRUsageContextInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	ID *string `json:"id,omitempty"`

	// A code that identifies the type of context being specified by this usage context.
	Code base.Code `json:"code,omitempty"`

	// A value that defines the context specified in this context of use. The interpretation of the value is defined by the code.
	ValueCodeableConcept *base.Code `json:"valueCodeableConcept,omitempty"`

	// A value that defines the context specified in this context of use. The interpretation of the value is defined by the code.
	ValueQuantity *FHIRQuantityInput `json:"valueQuantity,omitempty"`

	// A value that defines the context specified in this context of use. The interpretation of the value is defined by the code.
	ValueRange *FHIRRangeInput `json:"valueRange,omitempty"`

	// A value that defines the context specified in this context of use. The interpretation of the value is defined by the code.
	ValueReference *FHIRReferenceInput `json:"valueReference,omitempty"`
}

// AddressTypeEnum is a FHIR enum
type AddressTypeEnum string

const (
	// AddressTypeEnumPostal ...
	AddressTypeEnumPostal AddressTypeEnum = "postal"
	// AddressTypeEnumPhysical ...
	AddressTypeEnumPhysical AddressTypeEnum = "physical"
	// AddressTypeEnumBoth ...
	AddressTypeEnumBoth AddressTypeEnum = "both"
)

// AllAddressTypeEnum ...
var AllAddressTypeEnum = []AddressTypeEnum{
	AddressTypeEnumPostal,
	AddressTypeEnumPhysical,
	AddressTypeEnumBoth,
}

// IsValid ...
func (e AddressTypeEnum) IsValid() bool {
	switch e {
	case AddressTypeEnumPostal, AddressTypeEnumPhysical, AddressTypeEnumBoth:
		return true
	}
	return false
}

// String ...
func (e AddressTypeEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e *AddressTypeEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AddressTypeEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AddressTypeEnum", str)
	}
	return nil
}

// MarshalGQL writes the address type enum to the supplied writer as a quoted string
func (e AddressTypeEnum) MarshalGQL(w io.Writer) {
	_, err := fmt.Fprint(w, strconv.Quote(e.String()))
	if err != nil {
		log.Printf("%v\n", err)
	}
}

// AddressUseEnum is a FHIR enum
type AddressUseEnum string

const (
	// AddressUseEnumHome ...
	AddressUseEnumHome AddressUseEnum = "home"
	// AddressUseEnumWork ...
	AddressUseEnumWork AddressUseEnum = "work"
	// AddressUseEnumTemp ...
	AddressUseEnumTemp AddressUseEnum = "temp"
	// AddressUseEnumOld ...
	AddressUseEnumOld AddressUseEnum = "old"
	// AddressUseEnumBilling ...
	AddressUseEnumBilling AddressUseEnum = "billing"
)

// AllAddressUseEnum ...
var AllAddressUseEnum = []AddressUseEnum{
	AddressUseEnumHome,
	AddressUseEnumWork,
	AddressUseEnumTemp,
	AddressUseEnumOld,
	AddressUseEnumBilling,
}

// IsValid ...
func (e AddressUseEnum) IsValid() bool {
	switch e {
	case AddressUseEnumHome, AddressUseEnumWork, AddressUseEnumTemp, AddressUseEnumOld, AddressUseEnumBilling:
		return true
	}
	return false
}

// String ...
func (e AddressUseEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e *AddressUseEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AddressUseEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AddressUseEnum", str)
	}
	return nil
}

// MarshalGQL writes the address use enum to the supplied writer as a quoted string
func (e AddressUseEnum) MarshalGQL(w io.Writer) {
	_, err := fmt.Fprint(w, strconv.Quote(e.String()))
	if err != nil {
		log.Printf("%v\n", err)
	}
}

// AgeComparatorEnum is a FHIR enum
type AgeComparatorEnum string

const (
	// AgeComparatorEnumLessThan ...
	AgeComparatorEnumLessThan AgeComparatorEnum = "less_than"
	// AgeComparatorEnumLessThanOrEqualTo ...
	AgeComparatorEnumLessThanOrEqualTo AgeComparatorEnum = "less_than_or_equal_to"
	// AgeComparatorEnumGreaterThanOrEqualTo ...
	AgeComparatorEnumGreaterThanOrEqualTo AgeComparatorEnum = "greater_than_or_equal_to"
	// AgeComparatorEnumGreaterThan ...
	AgeComparatorEnumGreaterThan AgeComparatorEnum = "greater_than"
)

// AllAgeComparatorEnum ...
var AllAgeComparatorEnum = []AgeComparatorEnum{
	AgeComparatorEnumLessThan,
	AgeComparatorEnumLessThanOrEqualTo,
	AgeComparatorEnumGreaterThanOrEqualTo,
	AgeComparatorEnumGreaterThan,
}

// IsValid ...
func (e AgeComparatorEnum) IsValid() bool {
	switch e {
	case AgeComparatorEnumLessThan, AgeComparatorEnumLessThanOrEqualTo, AgeComparatorEnumGreaterThanOrEqualTo, AgeComparatorEnumGreaterThan:
		return true
	}
	return false
}

// String ...
func (e AgeComparatorEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e *AgeComparatorEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AgeComparatorEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AgeComparatorEnum", str)
	}
	return nil
}

// MarshalGQL writes the age comparator to the supplied writer as a quoted string
func (e AgeComparatorEnum) MarshalGQL(w io.Writer) {
	_, err := fmt.Fprint(w, strconv.Quote(e.String()))
	if err != nil {
		log.Printf("%v\n", err)
	}
}

// ContactPointSystemEnum is a FHIR enum
type ContactPointSystemEnum string

const (
	// ContactPointSystemEnumPhone ...
	ContactPointSystemEnumPhone ContactPointSystemEnum = "phone"
	// ContactPointSystemEnumFax ...
	ContactPointSystemEnumFax ContactPointSystemEnum = "fax"
	// ContactPointSystemEnumEmail ...
	ContactPointSystemEnumEmail ContactPointSystemEnum = "email"
	// ContactPointSystemEnumPager ...
	ContactPointSystemEnumPager ContactPointSystemEnum = "pager"
	// ContactPointSystemEnumURL ...
	ContactPointSystemEnumURL ContactPointSystemEnum = "url"
	// ContactPointSystemEnumSms ...
	ContactPointSystemEnumSms ContactPointSystemEnum = "sms"
	// ContactPointSystemEnumOther ...
	ContactPointSystemEnumOther ContactPointSystemEnum = "other"
)

// AllContactPointSystemEnum ...
var AllContactPointSystemEnum = []ContactPointSystemEnum{
	ContactPointSystemEnumPhone,
	ContactPointSystemEnumFax,
	ContactPointSystemEnumEmail,
	ContactPointSystemEnumPager,
	ContactPointSystemEnumURL,
	ContactPointSystemEnumSms,
	ContactPointSystemEnumOther,
}

// IsValid ...
func (e ContactPointSystemEnum) IsValid() bool {
	switch e {
	case ContactPointSystemEnumPhone, ContactPointSystemEnumFax, ContactPointSystemEnumEmail, ContactPointSystemEnumPager, ContactPointSystemEnumURL, ContactPointSystemEnumSms, ContactPointSystemEnumOther:
		return true
	}
	return false
}

// String ...
func (e ContactPointSystemEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e *ContactPointSystemEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ContactPointSystemEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContactPointSystemEnum", str)
	}
	return nil
}

// MarshalGQL writes the given enum to the supplied writer as a quoted string
func (e ContactPointSystemEnum) MarshalGQL(w io.Writer) {
	_, err := fmt.Fprint(w, strconv.Quote(e.String()))
	if err != nil {
		log.Printf("%v\n", err)
	}
}

// ContactPointUseEnum is a FHIR enum
type ContactPointUseEnum string

const (
	// ContactPointUseEnumHome ...
	ContactPointUseEnumHome ContactPointUseEnum = "home"
	// ContactPointUseEnumWork ...
	ContactPointUseEnumWork ContactPointUseEnum = "work"
	// ContactPointUseEnumTemp ...
	ContactPointUseEnumTemp ContactPointUseEnum = "temp"
	// ContactPointUseEnumOld ...
	ContactPointUseEnumOld ContactPointUseEnum = "old"
	// ContactPointUseEnumMobile ...
	ContactPointUseEnumMobile ContactPointUseEnum = "mobile"
)

// AllContactPointUseEnum ...
var AllContactPointUseEnum = []ContactPointUseEnum{
	ContactPointUseEnumHome,
	ContactPointUseEnumWork,
	ContactPointUseEnumTemp,
	ContactPointUseEnumOld,
	ContactPointUseEnumMobile,
}

// IsValid ...
func (e ContactPointUseEnum) IsValid() bool {
	switch e {
	case ContactPointUseEnumHome, ContactPointUseEnumWork, ContactPointUseEnumTemp, ContactPointUseEnumOld, ContactPointUseEnumMobile:
		return true
	}
	return false
}

// String ...
func (e ContactPointUseEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e *ContactPointUseEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ContactPointUseEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContactPointUseEnum", str)
	}
	return nil
}

// MarshalGQL writes the contact point use to the supplied writer as a quoted string
func (e ContactPointUseEnum) MarshalGQL(w io.Writer) {
	_, err := fmt.Fprint(w, strconv.Quote(e.String()))
	if err != nil {
		log.Printf("%v\n", err)
	}
}

// ContributorTypeEnum is a FHIR enum
type ContributorTypeEnum string

const (
	// ContributorTypeEnumAuthor ...
	ContributorTypeEnumAuthor ContributorTypeEnum = "author"
	// ContributorTypeEnumEditor ...
	ContributorTypeEnumEditor ContributorTypeEnum = "editor"
	// ContributorTypeEnumReviewer ...
	ContributorTypeEnumReviewer ContributorTypeEnum = "reviewer"
	// ContributorTypeEnumEndorser ...
	ContributorTypeEnumEndorser ContributorTypeEnum = "endorser"
)

// AllContributorTypeEnum ...
var AllContributorTypeEnum = []ContributorTypeEnum{
	ContributorTypeEnumAuthor,
	ContributorTypeEnumEditor,
	ContributorTypeEnumReviewer,
	ContributorTypeEnumEndorser,
}

// IsValid ...
func (e ContributorTypeEnum) IsValid() bool {
	switch e {
	case ContributorTypeEnumAuthor, ContributorTypeEnumEditor, ContributorTypeEnumReviewer, ContributorTypeEnumEndorser:
		return true
	}
	return false
}

// String ...
func (e ContributorTypeEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e *ContributorTypeEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ContributorTypeEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContributorTypeEnum", str)
	}
	return nil
}

// MarshalGQL writes the contributor type to the supplied writer as a quoted string
func (e ContributorTypeEnum) MarshalGQL(w io.Writer) {
	_, err := fmt.Fprint(w, strconv.Quote(e.String()))
	if err != nil {
		log.Printf("%v\n", err)
	}
}

// CountComparatorEnum is a FHIR enum
type CountComparatorEnum string

const (
	// CountComparatorEnumLessThan ...
	CountComparatorEnumLessThan CountComparatorEnum = "less_than"
	// CountComparatorEnumLessThanOrEqualTo ...
	CountComparatorEnumLessThanOrEqualTo CountComparatorEnum = "less_than_or_equal_to"
	// CountComparatorEnumGreaterThanOrEqualTo ...
	CountComparatorEnumGreaterThanOrEqualTo CountComparatorEnum = "greater_than_or_equal_to"
	// CountComparatorEnumGreaterThan ...
	CountComparatorEnumGreaterThan CountComparatorEnum = "greater_than"
)

// AllCountComparatorEnum ...
var AllCountComparatorEnum = []CountComparatorEnum{
	CountComparatorEnumLessThan,
	CountComparatorEnumLessThanOrEqualTo,
	CountComparatorEnumGreaterThanOrEqualTo,
	CountComparatorEnumGreaterThan,
}

// IsValid ...
func (e CountComparatorEnum) IsValid() bool {
	switch e {
	case CountComparatorEnumLessThan, CountComparatorEnumLessThanOrEqualTo, CountComparatorEnumGreaterThanOrEqualTo, CountComparatorEnumGreaterThan:
		return true
	}
	return false
}

// String ...
func (e CountComparatorEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e *CountComparatorEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CountComparatorEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CountComparatorEnum", str)
	}
	return nil
}

// MarshalGQL writes the count comparator to the supplied writer as a quoted string
func (e CountComparatorEnum) MarshalGQL(w io.Writer) {
	_, err := fmt.Fprint(w, strconv.Quote(e.String()))
	if err != nil {
		log.Printf("%v\n", err)
	}
}

// DataRequirementSortDirectionEnum is a FHIR enum
type DataRequirementSortDirectionEnum string

const (
	// DataRequirementSortDirectionEnumAscending ...
	DataRequirementSortDirectionEnumAscending DataRequirementSortDirectionEnum = "ascending"
	// DataRequirementSortDirectionEnumDescending ...
	DataRequirementSortDirectionEnumDescending DataRequirementSortDirectionEnum = "descending"
)

// AllDataRequirementSortDirectionEnum ...
var AllDataRequirementSortDirectionEnum = []DataRequirementSortDirectionEnum{
	DataRequirementSortDirectionEnumAscending,
	DataRequirementSortDirectionEnumDescending,
}

// IsValid ...
func (e DataRequirementSortDirectionEnum) IsValid() bool {
	switch e {
	case DataRequirementSortDirectionEnumAscending, DataRequirementSortDirectionEnumDescending:
		return true
	}
	return false
}

// String ...
func (e DataRequirementSortDirectionEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e *DataRequirementSortDirectionEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DataRequirementSortDirectionEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DataRequirement_SortDirectionEnum", str)
	}
	return nil
}

// MarshalGQL writes the given enum to the supplied writer as a quoted string
func (e DataRequirementSortDirectionEnum) MarshalGQL(w io.Writer) {
	_, err := fmt.Fprint(w, strconv.Quote(e.String()))
	if err != nil {
		log.Printf("%v\n", err)
	}
}

// DistanceComparatorEnum is a FHIR enum
type DistanceComparatorEnum string

const (
	// DistanceComparatorEnumLessThan ...
	DistanceComparatorEnumLessThan DistanceComparatorEnum = "less_than"
	// DistanceComparatorEnumLessThanOrEqualTo ...
	DistanceComparatorEnumLessThanOrEqualTo DistanceComparatorEnum = "less_than_or_equal_to"
	// DistanceComparatorEnumGreaterThanOrEqualTo ...
	DistanceComparatorEnumGreaterThanOrEqualTo DistanceComparatorEnum = "greater_than_or_equal_to"
	// DistanceComparatorEnumGreaterThan ...
	DistanceComparatorEnumGreaterThan DistanceComparatorEnum = "greater_than"
)

// AllDistanceComparatorEnum ...
var AllDistanceComparatorEnum = []DistanceComparatorEnum{
	DistanceComparatorEnumLessThan,
	DistanceComparatorEnumLessThanOrEqualTo,
	DistanceComparatorEnumGreaterThanOrEqualTo,
	DistanceComparatorEnumGreaterThan,
}

// IsValid ...
func (e DistanceComparatorEnum) IsValid() bool {
	switch e {
	case DistanceComparatorEnumLessThan, DistanceComparatorEnumLessThanOrEqualTo, DistanceComparatorEnumGreaterThanOrEqualTo, DistanceComparatorEnumGreaterThan:
		return true
	}
	return false
}

// String ...
func (e DistanceComparatorEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e *DistanceComparatorEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DistanceComparatorEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DistanceComparatorEnum", str)
	}
	return nil
}

// MarshalGQL writes the distance comparator to the supplied writer as a quoted string
func (e DistanceComparatorEnum) MarshalGQL(w io.Writer) {
	_, err := fmt.Fprint(w, strconv.Quote(e.String()))
	if err != nil {
		log.Printf("%v\n", err)
	}
}

// DurationComparatorEnum is a FHIR enum
type DurationComparatorEnum string

const (
	// DurationComparatorEnumLessThan ...
	DurationComparatorEnumLessThan DurationComparatorEnum = "less_than"
	// DurationComparatorEnumLessThanOrEqualTo ...
	DurationComparatorEnumLessThanOrEqualTo DurationComparatorEnum = "less_than_or_equal_to"
	// DurationComparatorEnumGreaterThanOrEqualTo ...
	DurationComparatorEnumGreaterThanOrEqualTo DurationComparatorEnum = "greater_than_or_equal_to"
	// DurationComparatorEnumGreaterThan ...
	DurationComparatorEnumGreaterThan DurationComparatorEnum = "greater_than"
)

// AllDurationComparatorEnum ...
var AllDurationComparatorEnum = []DurationComparatorEnum{
	DurationComparatorEnumLessThan,
	DurationComparatorEnumLessThanOrEqualTo,
	DurationComparatorEnumGreaterThanOrEqualTo,
	DurationComparatorEnumGreaterThan,
}

// IsValid ...
func (e DurationComparatorEnum) IsValid() bool {
	switch e {
	case DurationComparatorEnumLessThan, DurationComparatorEnumLessThanOrEqualTo, DurationComparatorEnumGreaterThanOrEqualTo, DurationComparatorEnumGreaterThan:
		return true
	}
	return false
}

// String ...
func (e DurationComparatorEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e *DurationComparatorEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DurationComparatorEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DurationComparatorEnum", str)
	}
	return nil
}

// MarshalGQL writes the duration comparator to the supplied writer as a quoted string
func (e DurationComparatorEnum) MarshalGQL(w io.Writer) {
	_, err := fmt.Fprint(w, strconv.Quote(e.String()))
	if err != nil {
		log.Printf("%v\n", err)
	}
}

// HumanNameUseEnum is a FHIR enum
type HumanNameUseEnum string

const (
	// HumanNameUseEnumUsual ...
	HumanNameUseEnumUsual HumanNameUseEnum = "usual"
	// HumanNameUseEnumOfficial ...
	HumanNameUseEnumOfficial HumanNameUseEnum = "official"
	// HumanNameUseEnumTemp ...
	HumanNameUseEnumTemp HumanNameUseEnum = "temp"
	// HumanNameUseEnumNickname ...
	HumanNameUseEnumNickname HumanNameUseEnum = "nickname"
	// HumanNameUseEnumAnonymous ...
	HumanNameUseEnumAnonymous HumanNameUseEnum = "anonymous"
	// HumanNameUseEnumOld ...
	HumanNameUseEnumOld HumanNameUseEnum = "old"
	// HumanNameUseEnumMaiden ...
	HumanNameUseEnumMaiden HumanNameUseEnum = "maiden"
)

// AllHumanNameUseEnum ...
var AllHumanNameUseEnum = []HumanNameUseEnum{
	HumanNameUseEnumUsual,
	HumanNameUseEnumOfficial,
	HumanNameUseEnumTemp,
	HumanNameUseEnumNickname,
	HumanNameUseEnumAnonymous,
	HumanNameUseEnumOld,
	HumanNameUseEnumMaiden,
}

// IsValid ...
func (e HumanNameUseEnum) IsValid() bool {
	switch e {
	case HumanNameUseEnumUsual, HumanNameUseEnumOfficial, HumanNameUseEnumTemp, HumanNameUseEnumNickname, HumanNameUseEnumAnonymous, HumanNameUseEnumOld, HumanNameUseEnumMaiden:
		return true
	}
	return false
}

// String ...
func (e HumanNameUseEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e *HumanNameUseEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = HumanNameUseEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid HumanNameUseEnum", str)
	}
	return nil
}

// MarshalGQL writes the given enum to the supplied writer as a quoted string
func (e HumanNameUseEnum) MarshalGQL(w io.Writer) {
	_, err := fmt.Fprint(w, strconv.Quote(e.String()))
	if err != nil {
		log.Printf("%v\n", err)
	}
}

// IdentifierUseEnum is a FHIR enum
type IdentifierUseEnum string

const (
	// IdentifierUseEnumUsual ...
	IdentifierUseEnumUsual IdentifierUseEnum = "usual"
	// IdentifierUseEnumOfficial ...
	IdentifierUseEnumOfficial IdentifierUseEnum = "official"
	// IdentifierUseEnumTemp ...
	IdentifierUseEnumTemp IdentifierUseEnum = "temp"
	// IdentifierUseEnumSecondary ...
	IdentifierUseEnumSecondary IdentifierUseEnum = "secondary"
	// IdentifierUseEnumOld ...
	IdentifierUseEnumOld IdentifierUseEnum = "old"
)

// AllIdentifierUseEnum ...
var AllIdentifierUseEnum = []IdentifierUseEnum{
	IdentifierUseEnumUsual,
	IdentifierUseEnumOfficial,
	IdentifierUseEnumTemp,
	IdentifierUseEnumSecondary,
	IdentifierUseEnumOld,
}

// IsValid ...
func (e IdentifierUseEnum) IsValid() bool {
	switch e {
	case IdentifierUseEnumUsual, IdentifierUseEnumOfficial, IdentifierUseEnumTemp, IdentifierUseEnumSecondary, IdentifierUseEnumOld:
		return true
	}
	return false
}

// String ...
func (e IdentifierUseEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e *IdentifierUseEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IdentifierUseEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IdentifierUseEnum", str)
	}
	return nil
}

// MarshalGQL writes the identifier use to the supplied writer as a quoted string
func (e IdentifierUseEnum) MarshalGQL(w io.Writer) {
	_, err := fmt.Fprint(w, strconv.Quote(e.String()))
	if err != nil {
		log.Printf("%v\n", err)
	}
}

// NarrativeStatusEnum is a FHIR enum
type NarrativeStatusEnum string

const (
	// NarrativeStatusEnumGenerated ...
	NarrativeStatusEnumGenerated NarrativeStatusEnum = "generated"
	// NarrativeStatusEnumExtensions ...
	NarrativeStatusEnumExtensions NarrativeStatusEnum = "extensions"
	// NarrativeStatusEnumAdditional ...
	NarrativeStatusEnumAdditional NarrativeStatusEnum = "additional"
	// NarrativeStatusEnumEmpty ...
	NarrativeStatusEnumEmpty NarrativeStatusEnum = "empty"
)

// AllNarrativeStatusEnum ...
var AllNarrativeStatusEnum = []NarrativeStatusEnum{
	NarrativeStatusEnumGenerated,
	NarrativeStatusEnumExtensions,
	NarrativeStatusEnumAdditional,
	NarrativeStatusEnumEmpty,
}

// IsValid ...
func (e NarrativeStatusEnum) IsValid() bool {
	switch e {
	case NarrativeStatusEnumGenerated, NarrativeStatusEnumExtensions, NarrativeStatusEnumAdditional, NarrativeStatusEnumEmpty:
		return true
	}
	return false
}

// String ...
func (e NarrativeStatusEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e *NarrativeStatusEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NarrativeStatusEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NarrativeStatusEnum", str)
	}
	return nil
}

// MarshalGQL writes the given enum to the supplied writer as a quoted string
func (e NarrativeStatusEnum) MarshalGQL(w io.Writer) {
	_, err := fmt.Fprint(w, strconv.Quote(e.String()))
	if err != nil {
		log.Printf("%v\n", err)
	}
}

// QuantityComparatorEnum is a FHIR enum
type QuantityComparatorEnum string

const (
	// QuantityComparatorEnumLessThan ...
	QuantityComparatorEnumLessThan QuantityComparatorEnum = "less_than"
	// QuantityComparatorEnumLessThanOrEqualTo ...
	QuantityComparatorEnumLessThanOrEqualTo QuantityComparatorEnum = "less_than_or_equal_to"
	// QuantityComparatorEnumGreaterThanOrEqualTo ...
	QuantityComparatorEnumGreaterThanOrEqualTo QuantityComparatorEnum = "greater_than_or_equal_to"
	// QuantityComparatorEnumGreaterThan ...
	QuantityComparatorEnumGreaterThan QuantityComparatorEnum = "greater_than"
)

// AllQuantityComparatorEnum ...
var AllQuantityComparatorEnum = []QuantityComparatorEnum{
	QuantityComparatorEnumLessThan,
	QuantityComparatorEnumLessThanOrEqualTo,
	QuantityComparatorEnumGreaterThanOrEqualTo,
	QuantityComparatorEnumGreaterThan,
}

// IsValid ...
func (e QuantityComparatorEnum) IsValid() bool {
	switch e {
	case QuantityComparatorEnumLessThan, QuantityComparatorEnumLessThanOrEqualTo, QuantityComparatorEnumGreaterThanOrEqualTo, QuantityComparatorEnumGreaterThan:
		return true
	}
	return false
}

// String ...
func (e QuantityComparatorEnum) String() string {
	switch e {
	case QuantityComparatorEnumLessThan:
		return "<"
	case QuantityComparatorEnumLessThanOrEqualTo:
		return "<="
	case QuantityComparatorEnumGreaterThanOrEqualTo:
		return ">="
	case QuantityComparatorEnumGreaterThan:
		return ">"
	default:
		return string(e)
	}
}

// UnmarshalGQL ...
func (e *QuantityComparatorEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = QuantityComparatorEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid QuantityComparatorEnum", str)
	}
	return nil
}

// MarshalGQL writes the quality comparator to the supplied writer as a quoted string
func (e QuantityComparatorEnum) MarshalGQL(w io.Writer) {
	_, err := fmt.Fprint(w, strconv.Quote(e.String()))
	if err != nil {
		log.Printf("%v\n", err)
	}
}

// TimingRepeatDurationUnitEnum is a FHIR enum
type TimingRepeatDurationUnitEnum string

const (
	// TimingRepeatDurationUnitEnumS ...
	TimingRepeatDurationUnitEnumS TimingRepeatDurationUnitEnum = "s"
	// TimingRepeatDurationUnitEnumMin ...
	TimingRepeatDurationUnitEnumMin TimingRepeatDurationUnitEnum = "min"
	// TimingRepeatDurationUnitEnumH ...
	TimingRepeatDurationUnitEnumH TimingRepeatDurationUnitEnum = "h"
	// TimingRepeatDurationUnitEnumD ...
	TimingRepeatDurationUnitEnumD TimingRepeatDurationUnitEnum = "d"
	// TimingRepeatDurationUnitEnumWk ...
	TimingRepeatDurationUnitEnumWk TimingRepeatDurationUnitEnum = "wk"
	// TimingRepeatDurationUnitEnumMo ...
	TimingRepeatDurationUnitEnumMo TimingRepeatDurationUnitEnum = "mo"
	// TimingRepeatDurationUnitEnumA ...
	TimingRepeatDurationUnitEnumA TimingRepeatDurationUnitEnum = "a"
)

// AllTimingRepeatDurationUnitEnum ...
var AllTimingRepeatDurationUnitEnum = []TimingRepeatDurationUnitEnum{
	TimingRepeatDurationUnitEnumS,
	TimingRepeatDurationUnitEnumMin,
	TimingRepeatDurationUnitEnumH,
	TimingRepeatDurationUnitEnumD,
	TimingRepeatDurationUnitEnumWk,
	TimingRepeatDurationUnitEnumMo,
	TimingRepeatDurationUnitEnumA,
}

// IsValid ...
func (e TimingRepeatDurationUnitEnum) IsValid() bool {
	switch e {
	case TimingRepeatDurationUnitEnumS, TimingRepeatDurationUnitEnumMin, TimingRepeatDurationUnitEnumH, TimingRepeatDurationUnitEnumD, TimingRepeatDurationUnitEnumWk, TimingRepeatDurationUnitEnumMo, TimingRepeatDurationUnitEnumA:
		return true
	}
	return false
}

// String ...
func (e TimingRepeatDurationUnitEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e *TimingRepeatDurationUnitEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TimingRepeatDurationUnitEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Timing_RepeatDurationUnitEnum", str)
	}
	return nil
}

// MarshalGQL writes the timing repeat duration to the supplied writer as a quoted string
func (e TimingRepeatDurationUnitEnum) MarshalGQL(w io.Writer) {
	_, err := fmt.Fprint(w, strconv.Quote(e.String()))
	if err != nil {
		log.Printf("%v\n", err)
	}
}

// TimingRepeatPeriodUnitEnum is a FHIR enum
type TimingRepeatPeriodUnitEnum string

const (
	// TimingRepeatPeriodUnitEnumS ...
	TimingRepeatPeriodUnitEnumS TimingRepeatPeriodUnitEnum = "s"
	// TimingRepeatPeriodUnitEnumMin ...
	TimingRepeatPeriodUnitEnumMin TimingRepeatPeriodUnitEnum = "min"
	// TimingRepeatPeriodUnitEnumH ...
	TimingRepeatPeriodUnitEnumH TimingRepeatPeriodUnitEnum = "h"
	// TimingRepeatPeriodUnitEnumD ...
	TimingRepeatPeriodUnitEnumD TimingRepeatPeriodUnitEnum = "d"
	// TimingRepeatPeriodUnitEnumWk ...
	TimingRepeatPeriodUnitEnumWk TimingRepeatPeriodUnitEnum = "wk"
	// TimingRepeatPeriodUnitEnumMo ...
	TimingRepeatPeriodUnitEnumMo TimingRepeatPeriodUnitEnum = "mo"
	// TimingRepeatPeriodUnitEnumA ...
	TimingRepeatPeriodUnitEnumA TimingRepeatPeriodUnitEnum = "a"
)

// AllTimingRepeatPeriodUnitEnum ...
var AllTimingRepeatPeriodUnitEnum = []TimingRepeatPeriodUnitEnum{
	TimingRepeatPeriodUnitEnumS,
	TimingRepeatPeriodUnitEnumMin,
	TimingRepeatPeriodUnitEnumH,
	TimingRepeatPeriodUnitEnumD,
	TimingRepeatPeriodUnitEnumWk,
	TimingRepeatPeriodUnitEnumMo,
	TimingRepeatPeriodUnitEnumA,
}

// IsValid ...
func (e TimingRepeatPeriodUnitEnum) IsValid() bool {
	switch e {
	case TimingRepeatPeriodUnitEnumS, TimingRepeatPeriodUnitEnumMin, TimingRepeatPeriodUnitEnumH, TimingRepeatPeriodUnitEnumD, TimingRepeatPeriodUnitEnumWk, TimingRepeatPeriodUnitEnumMo, TimingRepeatPeriodUnitEnumA:
		return true
	}
	return false
}

// String ...
func (e TimingRepeatPeriodUnitEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e *TimingRepeatPeriodUnitEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TimingRepeatPeriodUnitEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Timing_RepeatPeriodUnitEnum", str)
	}
	return nil
}

// MarshalGQL writes the timing repeat period to the supplied writer as a quoted string
func (e TimingRepeatPeriodUnitEnum) MarshalGQL(w io.Writer) {
	_, err := fmt.Fprint(w, strconv.Quote(e.String()))
	if err != nil {
		log.Printf("%v\n", err)
	}
}

// TimingRepeatWhenEnum is a FHIR enum
type TimingRepeatWhenEnum string

const (
	// TimingRepeatWhenEnumMorn ...
	TimingRepeatWhenEnumMorn TimingRepeatWhenEnum = "MORN"
	// TimingRepeatWhenEnumMornEarly ...
	TimingRepeatWhenEnumMornEarly TimingRepeatWhenEnum = "MORN_early"
	// TimingRepeatWhenEnumMornLate ...
	TimingRepeatWhenEnumMornLate TimingRepeatWhenEnum = "MORN_late"
	// TimingRepeatWhenEnumNoon ...
	TimingRepeatWhenEnumNoon TimingRepeatWhenEnum = "NOON"
	// TimingRepeatWhenEnumAft ...
	TimingRepeatWhenEnumAft TimingRepeatWhenEnum = "AFT"
	// TimingRepeatWhenEnumAftEarly ...
	TimingRepeatWhenEnumAftEarly TimingRepeatWhenEnum = "AFT_early"
	// TimingRepeatWhenEnumAftLate ...
	TimingRepeatWhenEnumAftLate TimingRepeatWhenEnum = "AFT_late"
	// TimingRepeatWhenEnumEve ...
	TimingRepeatWhenEnumEve TimingRepeatWhenEnum = "EVE"
	// TimingRepeatWhenEnumEveEarly ...
	TimingRepeatWhenEnumEveEarly TimingRepeatWhenEnum = "EVE_early"
	// TimingRepeatWhenEnumEveLate ...
	TimingRepeatWhenEnumEveLate TimingRepeatWhenEnum = "EVE_late"
	// TimingRepeatWhenEnumNight ...
	TimingRepeatWhenEnumNight TimingRepeatWhenEnum = "NIGHT"
	// TimingRepeatWhenEnumPhs ...
	TimingRepeatWhenEnumPhs TimingRepeatWhenEnum = "PHS"
	// TimingRepeatWhenEnumHs ...
	TimingRepeatWhenEnumHs TimingRepeatWhenEnum = "HS"
	// TimingRepeatWhenEnumWake ...
	TimingRepeatWhenEnumWake TimingRepeatWhenEnum = "WAKE"
	// TimingRepeatWhenEnumC ...
	TimingRepeatWhenEnumC TimingRepeatWhenEnum = "C"
	// TimingRepeatWhenEnumCm ...
	TimingRepeatWhenEnumCm TimingRepeatWhenEnum = "CM"
	// TimingRepeatWhenEnumCd ...
	TimingRepeatWhenEnumCd TimingRepeatWhenEnum = "CD"
	// TimingRepeatWhenEnumCv ...
	TimingRepeatWhenEnumCv TimingRepeatWhenEnum = "CV"
	// TimingRepeatWhenEnumAc ...
	TimingRepeatWhenEnumAc TimingRepeatWhenEnum = "AC"
	// TimingRepeatWhenEnumAcm ...
	TimingRepeatWhenEnumAcm TimingRepeatWhenEnum = "ACM"
	// TimingRepeatWhenEnumAcd ...
	TimingRepeatWhenEnumAcd TimingRepeatWhenEnum = "ACD"
	// TimingRepeatWhenEnumAcv ...
	TimingRepeatWhenEnumAcv TimingRepeatWhenEnum = "ACV"
	// TimingRepeatWhenEnumPc ...
	TimingRepeatWhenEnumPc TimingRepeatWhenEnum = "PC"
	// TimingRepeatWhenEnumPcm ...
	TimingRepeatWhenEnumPcm TimingRepeatWhenEnum = "PCM"
	// TimingRepeatWhenEnumPcd ...
	TimingRepeatWhenEnumPcd TimingRepeatWhenEnum = "PCD"
	// TimingRepeatWhenEnumPcv ...
	TimingRepeatWhenEnumPcv TimingRepeatWhenEnum = "PCV"
)

// AllTimingRepeatWhenEnum ...
var AllTimingRepeatWhenEnum = []TimingRepeatWhenEnum{
	TimingRepeatWhenEnumMorn,
	TimingRepeatWhenEnumMornEarly,
	TimingRepeatWhenEnumMornLate,
	TimingRepeatWhenEnumNoon,
	TimingRepeatWhenEnumAft,
	TimingRepeatWhenEnumAftEarly,
	TimingRepeatWhenEnumAftLate,
	TimingRepeatWhenEnumEve,
	TimingRepeatWhenEnumEveEarly,
	TimingRepeatWhenEnumEveLate,
	TimingRepeatWhenEnumNight,
	TimingRepeatWhenEnumPhs,
	TimingRepeatWhenEnumHs,
	TimingRepeatWhenEnumWake,
	TimingRepeatWhenEnumC,
	TimingRepeatWhenEnumCm,
	TimingRepeatWhenEnumCd,
	TimingRepeatWhenEnumCv,
	TimingRepeatWhenEnumAc,
	TimingRepeatWhenEnumAcm,
	TimingRepeatWhenEnumAcd,
	TimingRepeatWhenEnumAcv,
	TimingRepeatWhenEnumPc,
	TimingRepeatWhenEnumPcm,
	TimingRepeatWhenEnumPcd,
	TimingRepeatWhenEnumPcv,
}

// IsValid ...
func (e TimingRepeatWhenEnum) IsValid() bool {
	switch e {
	case TimingRepeatWhenEnumMorn, TimingRepeatWhenEnumMornEarly, TimingRepeatWhenEnumMornLate, TimingRepeatWhenEnumNoon, TimingRepeatWhenEnumAft, TimingRepeatWhenEnumAftEarly, TimingRepeatWhenEnumAftLate, TimingRepeatWhenEnumEve, TimingRepeatWhenEnumEveEarly, TimingRepeatWhenEnumEveLate, TimingRepeatWhenEnumNight, TimingRepeatWhenEnumPhs, TimingRepeatWhenEnumHs, TimingRepeatWhenEnumWake, TimingRepeatWhenEnumC, TimingRepeatWhenEnumCm, TimingRepeatWhenEnumCd, TimingRepeatWhenEnumCv, TimingRepeatWhenEnumAc, TimingRepeatWhenEnumAcm, TimingRepeatWhenEnumAcd, TimingRepeatWhenEnumAcv, TimingRepeatWhenEnumPc, TimingRepeatWhenEnumPcm, TimingRepeatWhenEnumPcd, TimingRepeatWhenEnumPcv:
		return true
	}
	return false
}

// String ...
func (e TimingRepeatWhenEnum) String() string {
	return string(e)
}

// UnmarshalGQL ...
func (e *TimingRepeatWhenEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TimingRepeatWhenEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Timing_RepeatWhenEnum", str)
	}
	return nil
}

// MarshalGQL writes when timings repeat to the supplied writer as a quoted string
func (e TimingRepeatWhenEnum) MarshalGQL(w io.Writer) {
	_, err := fmt.Fprint(w, strconv.Quote(e.String()))
	if err != nil {
		log.Printf("%v\n", err)
	}
}