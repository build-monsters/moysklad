package moysklad

import (
	"github.com/goccy/go-json"
	"github.com/google/uuid"
)

type Payment struct {
	AccountID      *uuid.UUID            `json:"accountId,omitempty"`      // ID учетной записи
	Agent          *Counterparty         `json:"agent,omitempty"`          // Метаданные контрагента
	Applicable     *bool                 `json:"applicable,omitempty"`     // Отметка о проведении
	Attributes     Slice[AttributeValue] `json:"attributes,omitempty"`     // Коллекция метаданных доп. полей. Поля объекта
	Code           *string               `json:"code,omitempty"`           // Код выданного
	Contract       *Contract             `json:"contract,omitempty"`       // Метаданные договора
	Created        *Timestamp            `json:"created,omitempty"`        // Дата создания
	Deleted        *Timestamp            `json:"deleted,omitempty"`        // Момент последнего удаления
	Description    *string               `json:"description,omitempty"`    // Комментарий
	ExternalCode   *string               `json:"externalCode,omitempty"`   // Внешний код
	Files          *MetaArray[File]      `json:"files,omitempty"`          // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group          *Group                `json:"group,omitempty"`          // Отдел сотрудника
	ID             *uuid.UUID            `json:"id,omitempty"`             // ID сущности
	Meta           Meta                  `json:"meta"`                     // Метаданные
	Moment         *Timestamp            `json:"moment,omitempty"`         // Дата документа
	Name           *string               `json:"name,omitempty"`           // Наименование
	Organization   *Organization         `json:"organization,omitempty"`   // Метаданные юрлица
	Owner          *Employee             `json:"owner,omitempty"`          // Владелец (Сотрудник)
	PaymentPurpose *string               `json:"paymentPurpose,omitempty"` // Основание
	Printed        *bool                 `json:"printed,omitempty"`        // Напечатан ли документ
	Project        *Project              `json:"project,omitempty"`        // Метаданные проекта
	Published      *bool                 `json:"published,omitempty"`      // Опубликован ли документ
	Rate           *Rate                 `json:"rate,omitempty"`           // Валюта
	SalesChannel   *SalesChannel         `json:"salesChannel,omitempty"`   // Метаданные канала продаж
	Shared         *bool                 `json:"shared,omitempty"`         // Общий доступ
	State          *State                `json:"state,omitempty"`          // Метаданные статуса
	Sum            *float64              `json:"sum,omitempty"`            // Сумма
	SyncID         *uuid.UUID            `json:"syncId,omitempty"`         // ID синхронизации. После заполнения недоступен для изменения
	Updated        *Timestamp            `json:"updated,omitempty"`        // Момент последнего обновления
	VatSum         *float64              `json:"vatSum,omitempty"`         // Сумма включая НДС
	LinkedSum      *float64              `json:"linkedSum,omitempty"`      // Сумма, оплаченная по документу из этого платежа
	Operations     Operations            `json:"operations,omitempty"`     // Массив ссылок на связанные операции в формате Метаданных
	raw            []byte                // сырые данные
}

func (payment Payment) GetAccountID() uuid.UUID {
	return Deref(payment.AccountID)
}

func (payment Payment) GetAgent() Counterparty {
	return Deref(payment.Agent)
}

func (payment Payment) GetApplicable() bool {
	return Deref(payment.Applicable)
}

func (payment Payment) GetAttributes() Slice[AttributeValue] {
	return payment.Attributes
}

func (payment Payment) GetCode() string {
	return Deref(payment.Code)
}

func (payment Payment) GetContract() Contract {
	return Deref(payment.Contract)
}

func (payment Payment) GetCreated() Timestamp {
	return Deref(payment.Created)
}

func (payment Payment) GetDeleted() Timestamp {
	return Deref(payment.Deleted)
}

func (payment Payment) GetDescription() string {
	return Deref(payment.Description)
}

func (payment Payment) GetExternalCode() string {
	return Deref(payment.ExternalCode)
}

func (payment Payment) GetFiles() MetaArray[File] {
	return Deref(payment.Files)
}

func (payment Payment) GetGroup() Group {
	return Deref(payment.Group)
}

func (payment Payment) GetID() uuid.UUID {
	return Deref(payment.ID)
}

func (payment Payment) GetMeta() Meta {
	return payment.Meta
}

func (payment Payment) GetMoment() Timestamp {
	return Deref(payment.Moment)
}

func (payment Payment) GetName() string {
	return Deref(payment.Name)
}

func (payment Payment) GetOrganization() Organization {
	return Deref(payment.Organization)
}

func (payment Payment) GetOwner() Employee {
	return Deref(payment.Owner)
}

func (payment Payment) GetPaymentPurpose() string {
	return Deref(payment.PaymentPurpose)
}

func (payment Payment) GetPrinted() bool {
	return Deref(payment.Printed)
}

func (payment Payment) GetProject() Project {
	return Deref(payment.Project)
}

func (payment Payment) GetPublished() bool {
	return Deref(payment.Published)
}

func (payment Payment) GetRate() Rate {
	return Deref(payment.Rate)
}

func (payment Payment) GetSalesChannel() SalesChannel {
	return Deref(payment.SalesChannel)
}

func (payment Payment) GetShared() bool {
	return Deref(payment.Shared)
}

func (payment Payment) GetState() State {
	return Deref(payment.State)
}

func (payment Payment) GetSum() float64 {
	return Deref(payment.Sum)
}

func (payment Payment) GetSyncID() uuid.UUID {
	return Deref(payment.SyncID)
}

func (payment Payment) GetUpdated() Timestamp {
	return Deref(payment.Updated)
}

func (payment Payment) GetVatSum() float64 {
	return Deref(payment.VatSum)
}

func (payment Payment) GetLinkedSum() float64 {
	return Deref(payment.LinkedSum)
}

func (payment Payment) GetOperations() Operations {
	return payment.Operations
}

// MetaType удовлетворяет интерфейсу MetaTyper
func (payment Payment) MetaType() MetaType {
	return payment.Meta.GetType()
}

// Raw удовлетворяет интерфейсу RawMetaTyper
func (payment Payment) Raw() []byte {
	return payment.raw
}

func (payment Payment) String() string {
	return Stringify(payment.Meta)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (payment *Payment) UnmarshalJSON(data []byte) (err error) {
	type alias Payment
	var t alias

	if err = json.Unmarshal(data, &t); err != nil {
		return err
	}

	t.raw = data
	*payment = Payment(t)
	return nil
}

// AsCashIn десериализует сырые данные в тип *CashIn
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (payment *Payment) AsCashIn() *CashIn {
	return unmarshalAsType[CashIn](payment)
}

// AsCashOut десериализует сырые данные в тип *CashOut
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (payment *Payment) AsCashOut() *CashOut {
	return unmarshalAsType[CashOut](payment)
}

// AsPaymentIn десериализует сырые данные в тип *PaymentIn
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (payment *Payment) AsPaymentIn() *PaymentIn {
	return unmarshalAsType[PaymentIn](payment)
}

// AsPaymentOut десериализует сырые данные в тип *PaymentOut
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (payment *Payment) AsPaymentOut() *PaymentOut {
	return unmarshalAsType[PaymentOut](payment)
}

func NewPayments() Slice[Payment] {
	return NewSlice[Payment]()
}
