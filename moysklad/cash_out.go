package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CashOut Расходный ордер.
// Ключевое слово: cashout
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-rashodnyj-order
type CashOut struct {
	Name           *string               `json:"name,omitempty"`
	Deleted        *Timestamp            `json:"deleted,omitempty"`
	Applicable     *bool                 `json:"applicable,omitempty"`
	AccountID      *uuid.UUID            `json:"accountId,omitempty"`
	Code           *string               `json:"code,omitempty"`
	Contract       *Contract             `json:"contract,omitempty"`
	Created        *Timestamp            `json:"created,omitempty"`
	Organization   *Organization         `json:"organization,omitempty"`
	Description    *string               `json:"description,omitempty"`
	ExpenseItem    *ExpenseItem          `json:"expenseItem,omitempty"`
	ExternalCode   *string               `json:"externalCode,omitempty"`
	Files          *MetaArray[File]      `json:"files,omitempty"`
	Group          *Group                `json:"group,omitempty"`
	Owner          *Employee             `json:"owner,omitempty"`
	Meta           *Meta                 `json:"meta,omitempty"`
	Moment         *Timestamp            `json:"moment,omitempty"`
	Operations     Operations            `json:"operations,omitempty"`
	Agent          *Counterparty         `json:"agent,omitempty"`
	ID             *uuid.UUID            `json:"id,omitempty"`
	PaymentPurpose *string               `json:"paymentPurpose,omitempty"`
	Printed        *bool                 `json:"printed,omitempty"`
	Project        *Project              `json:"project,omitempty"`
	Published      *bool                 `json:"published,omitempty"`
	Rate           *Rate                 `json:"rate,omitempty"`
	SalesChannel   *SalesChannel         `json:"salesChannel,omitempty"`
	Shared         *bool                 `json:"shared,omitempty"`
	State          *State                `json:"state,omitempty"`
	Sum            *float64              `json:"sum,omitempty"`
	SyncID         *uuid.UUID            `json:"syncId,omitempty"`
	Updated        *Timestamp            `json:"updated,omitempty"`
	VatSum         *float64              `json:"vatSum,omitempty"`
	FactureOut     *FactureOut           `json:"factureOut,omitempty"`
	Attributes     Slice[AttributeValue] `json:"attributes,omitempty"`
}

func (cashOut CashOut) Clean() *CashOut {
	return &CashOut{Meta: cashOut.Meta}
}

func (cashOut CashOut) GetName() string {
	return Deref(cashOut.Name)
}

func (cashOut CashOut) GetDeleted() Timestamp {
	return Deref(cashOut.Deleted)
}

func (cashOut CashOut) GetApplicable() bool {
	return Deref(cashOut.Applicable)
}

func (cashOut CashOut) GetAccountID() uuid.UUID {
	return Deref(cashOut.AccountID)
}

func (cashOut CashOut) GetCode() string {
	return Deref(cashOut.Code)
}

func (cashOut CashOut) GetContract() Contract {
	return Deref(cashOut.Contract)
}

func (cashOut CashOut) GetCreated() Timestamp {
	return Deref(cashOut.Created)
}

func (cashOut CashOut) GetOrganization() Organization {
	return Deref(cashOut.Organization)
}

func (cashOut CashOut) GetDescription() string {
	return Deref(cashOut.Description)
}

func (cashOut CashOut) GetExpenseItem() ExpenseItem {
	return Deref(cashOut.ExpenseItem)
}

func (cashOut CashOut) GetExternalCode() string {
	return Deref(cashOut.ExternalCode)
}

func (cashOut CashOut) GetFiles() MetaArray[File] {
	return Deref(cashOut.Files)
}

func (cashOut CashOut) GetGroup() Group {
	return Deref(cashOut.Group)
}

func (cashOut CashOut) GetOwner() Employee {
	return Deref(cashOut.Owner)
}

func (cashOut CashOut) GetMeta() Meta {
	return Deref(cashOut.Meta)
}

func (cashOut CashOut) GetMoment() Timestamp {
	return Deref(cashOut.Moment)
}

func (cashOut CashOut) GetOperations() Operations {
	return cashOut.Operations
}

func (cashOut CashOut) GetAgent() Counterparty {
	return Deref(cashOut.Agent)
}

func (cashOut CashOut) GetID() uuid.UUID {
	return Deref(cashOut.ID)
}

func (cashOut CashOut) GetPaymentPurpose() string {
	return Deref(cashOut.PaymentPurpose)
}

func (cashOut CashOut) GetPrinted() bool {
	return Deref(cashOut.Printed)
}

func (cashOut CashOut) GetProject() Project {
	return Deref(cashOut.Project)
}

func (cashOut CashOut) GetPublished() bool {
	return Deref(cashOut.Published)
}

func (cashOut CashOut) GetRate() Rate {
	return Deref(cashOut.Rate)
}

func (cashOut CashOut) GetSalesChannel() SalesChannel {
	return Deref(cashOut.SalesChannel)
}

func (cashOut CashOut) GetShared() bool {
	return Deref(cashOut.Shared)
}

func (cashOut CashOut) GetState() State {
	return Deref(cashOut.State)
}

func (cashOut CashOut) GetSum() float64 {
	return Deref(cashOut.Sum)
}

func (cashOut CashOut) GetSyncID() uuid.UUID {
	return Deref(cashOut.SyncID)
}

func (cashOut CashOut) GetUpdated() Timestamp {
	return Deref(cashOut.Updated)
}

func (cashOut CashOut) GetVatSum() float64 {
	return Deref(cashOut.VatSum)
}

func (cashOut CashOut) GetFactureOut() FactureOut {
	return Deref(cashOut.FactureOut)
}

func (cashOut CashOut) GetAttributes() Slice[AttributeValue] {
	return cashOut.Attributes
}

func (cashOut *CashOut) SetName(name string) *CashOut {
	cashOut.Name = &name
	return cashOut
}

func (cashOut *CashOut) SetApplicable(applicable bool) *CashOut {
	cashOut.Applicable = &applicable
	return cashOut
}

func (cashOut *CashOut) SetCode(code string) *CashOut {
	cashOut.Code = &code
	return cashOut
}

func (cashOut *CashOut) SetContract(contract *Contract) *CashOut {
	cashOut.Contract = contract.Clean()
	return cashOut
}

func (cashOut *CashOut) SetOrganization(organization *Organization) *CashOut {
	cashOut.Organization = organization.Clean()
	return cashOut
}

func (cashOut *CashOut) SetDescription(description string) *CashOut {
	cashOut.Description = &description
	return cashOut
}

func (cashOut *CashOut) SetExpenseItem(expenseItem *ExpenseItem) *CashOut {
	cashOut.ExpenseItem = expenseItem.Clean()
	return cashOut
}

func (cashOut *CashOut) SetExternalCode(externalCode string) *CashOut {
	cashOut.ExternalCode = &externalCode
	return cashOut
}

func (cashOut *CashOut) SetFiles(files Slice[File]) *CashOut {
	cashOut.Files = NewMetaArrayRows(files)
	return cashOut
}

func (cashOut *CashOut) SetGroup(group *Group) *CashOut {
	cashOut.Group = group.Clean()
	return cashOut
}

func (cashOut *CashOut) SetOwner(owner *Employee) *CashOut {
	cashOut.Owner = owner.Clean()
	return cashOut
}

func (cashOut *CashOut) SetMeta(meta *Meta) *CashOut {
	cashOut.Meta = meta
	return cashOut
}

func (cashOut *CashOut) SetMoment(moment *Timestamp) *CashOut {
	cashOut.Moment = moment
	return cashOut
}

func (cashOut *CashOut) SetOperations(operations Operations) *CashOut {
	cashOut.Operations = operations
	return cashOut
}

func (cashOut *CashOut) SetAgent(agent *Counterparty) *CashOut {
	cashOut.Agent = agent.Clean()
	return cashOut
}

func (cashOut *CashOut) SetPaymentPurpose(paymentPurpose string) *CashOut {
	cashOut.PaymentPurpose = &paymentPurpose
	return cashOut
}

func (cashOut *CashOut) SetProject(project *Project) *CashOut {
	cashOut.Project = project.Clean()
	return cashOut
}

func (cashOut *CashOut) SetRate(rate *Rate) *CashOut {
	cashOut.Rate = rate
	return cashOut
}

func (cashOut *CashOut) SetSalesChannel(salesChannel *SalesChannel) *CashOut {
	cashOut.SalesChannel = salesChannel.Clean()
	return cashOut
}

func (cashOut *CashOut) SetShared(shared bool) *CashOut {
	cashOut.Shared = &shared
	return cashOut
}

func (cashOut *CashOut) SetState(state *State) *CashOut {
	cashOut.State = state.Clean()
	return cashOut
}

func (cashOut *CashOut) SetSum(sum float64) *CashOut {
	cashOut.Sum = &sum
	return cashOut
}

func (cashOut *CashOut) SetSyncID(syncID uuid.UUID) *CashOut {
	cashOut.SyncID = &syncID
	return cashOut
}

func (cashOut *CashOut) SetVatSum(vatSum float64) *CashOut {
	cashOut.VatSum = &vatSum
	return cashOut
}

func (cashOut *CashOut) SetFactureOut(factureOut *FactureOut) *CashOut {
	cashOut.FactureOut = factureOut.Clean()
	return cashOut
}

func (cashOut *CashOut) SetAttributes(attributes Slice[AttributeValue]) *CashOut {
	cashOut.Attributes = attributes
	return cashOut
}

func (cashOut CashOut) String() string {
	return Stringify(cashOut)
}

func (cashOut CashOut) MetaType() MetaType {
	return MetaTypeCashOut
}

func (cashOut CashOut) AsPayment() *Payment {
	return &Payment{Meta: cashOut.GetMeta()}
}

// CashOutTemplateArg
// Документ: Расходный ордер (cashout)
// Основание, на котором он может быть создан:
// - Возврат покупателя (salesreturn)
// - Приемка (supply)
// - Счет поставщика (invoicein)
// - Заказ поставщику (purchaseorder)
// - Выданный отчет комиссионера (commissionreportout)
//type CashOutTemplateArg struct {
//	SalesReturn         *MetaWrapper `json:"salesReturn,omitempty"`
//	Supply              *MetaWrapper `json:"supply,omitempty"`
//	InvoiceIn           *MetaWrapper `json:"invoiceIn,omitempty"`
//	PurchaseOrder       *MetaWrapper `json:"purchaseOrder,omitempty"`
//	CommissionReportOut *MetaWrapper `json:"commissionReportOut,omitempty"`
//}

// CashOutService cashout
// Сервис для работы с расходными ордерами.
type CashOutService interface {
	GetList(ctx context.Context, params *Params) (*List[CashOut], *resty.Response, error)
	Create(ctx context.Context, cashOut *CashOut, params *Params) (*CashOut, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, cashOutList Slice[CashOut], params *Params) (*Slice[CashOut], *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteMany(ctx context.Context, cashOutList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList Slice[Attribute]) (*Slice[Attribute], *resty.Response, error)
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	//Template(ctx context.Context) (*CashOut, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*CashOut, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, cashOut *CashOut, params *Params) (*CashOut, *resty.Response, error)
	GetPublications(ctx context.Context, id uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id uuid.UUID, template Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*CashOut, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetStateByID(ctx context.Context, id uuid.UUID) (*State, *resty.Response, error)
	CreateState(ctx context.Context, state *State) (*State, *resty.Response, error)
	UpdateState(ctx context.Context, id uuid.UUID, state *State) (*State, *resty.Response, error)
	CreateOrUpdateStates(ctx context.Context, states Slice[State]) (*Slice[State], *resty.Response, error)
	DeleteState(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetFiles(ctx context.Context, id uuid.UUID) (*MetaArray[File], *resty.Response, error)
	CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error)
	UpdateFiles(ctx context.Context, id uuid.UUID, files Slice[File]) (*Slice[File], *resty.Response, error)
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)
	DeleteFiles(ctx context.Context, id uuid.UUID, files *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
}

func NewCashOutService(client *Client) CashOutService {
	e := NewEndpoint(client, "entity/cashout")
	return newMainService[CashOut, any, MetaAttributesSharedStatesWrapper, any](e)
}
