package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// InvoiceOut Счет покупателю.
// Ключевое слово: invoiceout
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-pokupatelu
type InvoiceOut struct {
	PayedSum             *Decimal                    `json:"payedSum,omitempty"`
	VatEnabled           *bool                       `json:"vatEnabled,omitempty"`
	AgentAccount         *AgentAccount               `json:"agentAccount,omitempty"`
	Applicable           *bool                       `json:"applicable,omitempty"`
	Demands              *Demands                    `json:"demands,omitempty"`
	Code                 *string                     `json:"code,omitempty"`
	OrganizationAccount  *AgentAccount               `json:"organizationAccount,omitempty"`
	Created              *Timestamp                  `json:"created,omitempty"`
	Deleted              *Timestamp                  `json:"deleted,omitempty"`
	Description          *string                     `json:"description,omitempty"`
	ExternalCode         *string                     `json:"externalCode,omitempty"`
	Files                *Files                      `json:"files,omitempty"`
	Group                *Group                      `json:"group,omitempty"`
	ID                   *uuid.UUID                  `json:"id,omitempty"`
	Meta                 *Meta                       `json:"meta,omitempty"`
	Moment               *Timestamp                  `json:"moment,omitempty"`
	Name                 *string                     `json:"name,omitempty"`
	AccountID            *uuid.UUID                  `json:"accountId,omitempty"`
	Contract             *Contract                   `json:"contract,omitempty"`
	Agent                *Counterparty               `json:"agent,omitempty"`
	Organization         *Organization               `json:"organization,omitempty"`
	PaymentPlannedMoment *Timestamp                  `json:"paymentPlannedMoment,omitempty"`
	Positions            *Positions[InvoicePosition] `json:"positions,omitempty"`
	Printed              *bool                       `json:"printed,omitempty"`
	Project              *Project                    `json:"project,omitempty"`
	Published            *bool                       `json:"published,omitempty"`
	Rate                 *Rate                       `json:"rate,omitempty"`
	Shared               *bool                       `json:"shared,omitempty"`
	ShippedSum           *Decimal                    `json:"shippedSum,omitempty"`
	State                *State                      `json:"state,omitempty"`
	Store                *Store                      `json:"store,omitempty"`
	Sum                  *Decimal                    `json:"sum,omitempty"`
	SyncID               *uuid.UUID                  `json:"syncId,omitempty"`
	Updated              *Timestamp                  `json:"updated,omitempty"`
	Owner                *Employee                   `json:"owner,omitempty"`
	VatIncluded          *bool                       `json:"vatIncluded,omitempty"`
	VatSum               *Decimal                    `json:"vatSum,omitempty"`
	CustomerOrder        *CustomerOrder              `json:"customerOrder,omitempty"`
	Payments             *Payments                   `json:"payments,omitempty"`
	Attributes           Attributes                  `json:"attributes,omitempty"`
}

func (i InvoiceOut) String() string {
	return Stringify(i)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (i InvoiceOut) GetMeta() Meta {
	return Deref(i.Meta)
}

func (i InvoiceOut) MetaType() MetaType {
	return MetaTypeInvoiceOut
}

type InvoicesOut = Slice[InvoiceOut]

// InvoiceOutPosition Позиция Счета покупателю.
// Ключевое слово: invoiceposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-pokupatelu-scheta-pokupatelqm-pozicii-scheta-pokupatelu
type InvoiceOutPosition struct {
	InvoicePosition
}

func (i InvoiceOutPosition) String() string {
	return Stringify(i)
}

func (i InvoiceOutPosition) MetaType() MetaType {
	return MetaTypeInvoicePosition
}

// InvoiceOutTemplateArg
// Документ: Cчет покупателю (invoiceout)
// Основание, на котором он может быть создан:
// - Заказ покупателя (customerorder)
type InvoiceOutTemplateArg struct {
	CustomerOrder *MetaWrapper `json:"customerOrder,omitempty"`
}

// InvoiceOutService
// Сервис для работы со счетами покупателей.
type InvoiceOutService interface {
	GetList(ctx context.Context, params *Params) (*List[InvoiceOut], *resty.Response, error)
	Create(ctx context.Context, invoiceOut *InvoiceOut, params *Params) (*InvoiceOut, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, invoiceOutList []*InvoiceOut, params *Params) (*[]InvoiceOut, *resty.Response, error)
	DeleteMany(ctx context.Context, invoiceOutList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*InvoiceOut, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, invoiceOut *InvoiceOut, params *Params) (*InvoiceOut, *resty.Response, error)
	//endpointTemplate[InvoiceOut]
	//endpointTemplateBasedOn[InvoiceOut, InvoiceOutTemplateArg]
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[InvoiceOutPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*InvoiceOutPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *InvoiceOutPosition, params *Params) (*InvoiceOutPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *InvoiceOutPosition) (*InvoiceOutPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*InvoiceOutPosition) (*[]InvoiceOutPosition, *resty.Response, error)
	DeletePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (bool, *resty.Response, error)
	GetPositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error)
	CreateOrUpdatePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*[]TrackingCode, *resty.Response, error)
	DeletePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*DeleteManyResponse, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*InvoiceOut, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewInvoiceOutService(client *Client) InvoiceOutService {
	e := NewEndpoint(client, "entity/invoiceout")
	return newMainService[InvoiceOut, InvoiceOutPosition, MetadataAttributeSharedStates, any](e)
}
