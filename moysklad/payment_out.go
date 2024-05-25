package moysklad

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// PaymentOut Исходящий платеж.
// Ключевое слово: paymentout
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-ishodqschij-platezh
type PaymentOut struct {
	AccountID           *uuid.UUID       `json:"accountId,omitempty"`           // ID учетной записи
	Agent               *Counterparty    `json:"agent,omitempty"`               // Метаданные контрагента
	AgentAccount        *AgentAccount    `json:"agentAccount,omitempty"`        // Метаданные счета контрагента
	Applicable          *bool            `json:"applicable,omitempty"`          // Отметка о проведении
	Attributes          *Attributes      `json:"attributes,omitempty"`          // Коллекция метаданных доп. полей. Поля объекта
	Code                *string          `json:"code,omitempty"`                // Код выданного
	Contract            *Contract        `json:"contract,omitempty"`            // Метаданные договора
	Created             *Timestamp       `json:"created,omitempty"`             // Дата создания
	Deleted             *Timestamp       `json:"deleted,omitempty"`             // Момент последнего удаления
	Description         *string          `json:"description,omitempty"`         // Комментарий
	ExpenseItem         *ExpenseItem     `json:"expenseItem,omitempty"`         // Метаданные Статьи расходов
	ExternalCode        *string          `json:"externalCode,omitempty"`        // Внешний код
	Files               *Files           `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group               *Group           `json:"group,omitempty"`               // Отдел сотрудника
	ID                  *uuid.UUID       `json:"id,omitempty"`                  // ID сущности
	Meta                *Meta            `json:"meta,omitempty"`                // Метаданные
	Moment              *Timestamp       `json:"moment,omitempty"`              // Дата документа
	Name                *string          `json:"name,omitempty"`                // Наименование
	Organization        *Organization    `json:"organization,omitempty"`        // Метаданные юрлица
	OrganizationAccount *AgentAccount    `json:"organizationAccount,omitempty"` // Метаданные счета юрлица
	Owner               *Employee        `json:"owner,omitempty"`               // Владелец (Сотрудник)
	PaymentPurpose      *string          `json:"paymentPurpose,omitempty"`      // Назначение платежа
	Printed             *bool            `json:"printed,omitempty"`             // Напечатан ли документ
	Project             *Project         `json:"project,omitempty"`             // Проект
	Published           *bool            `json:"published,omitempty"`           // Опубликован ли документ
	Rate                *Rate            `json:"rate,omitempty"`                // Валюта
	SalesChannel        *SalesChannel    `json:"salesChannel,omitempty"`        // Метаданные канала продаж
	Shared              *bool            `json:"shared,omitempty"`              // Общий доступ
	State               *State           `json:"state,omitempty"`               // Метаданные статуса
	Sum                 *decimal.Decimal `json:"sum,omitempty"`                 // Сумма
	SyncID              *uuid.UUID       `json:"syncId,omitempty"`              // ID синхронизации. После заполнения недоступен для изменения
	Updated             *Timestamp       `json:"updated,omitempty"`             // Момент последнего обновления
	VatSum              *decimal.Decimal `json:"vatSum,omitempty"`              // Сумма включая НДС
	FactureIn           *FactureIn       `json:"factureIn,omitempty"`           // Ссылка на Счет-фактуру
	Operations          *Operations      `json:"operations,omitempty"`          // Массив ссылок на связанные операции в формате Метаданных
}

func (p PaymentOut) String() string {
	return Stringify(p)
}

func (p PaymentOut) MetaType() MetaType {
	return MetaTypePaymentOut
}

// BindDocuments Привязка платежей к документам.
// Необходимо передать *Meta документов, к которым необходимо привязать платёж.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-obschie-swedeniq-priwqzka-platezhej-k-dokumentam
func (p *PaymentOut) BindDocuments(documentsMeta ...*Meta) *PaymentOut {
	if p.Operations == nil {
		p.Operations = new(Operations)
	}

	for _, meta := range documentsMeta {
		*p.Operations = append(*p.Operations, Operation{Meta: Deref(meta)})
	}

	return p
}

// PaymentOutTemplateArg
// Документ: Исходящий платеж (paymentout)
// Основание, на котором он может быть создан:
// - Возврат покупателя (salesreturn)
// - Приемка (supply)
// - Счет поставщика (invoicein)
// - Заказ поставщику (purchaseorder)
// - Выданный отчет комиссионера (commissionreportout)
type PaymentOutTemplateArg struct {
	SalesReturn         *MetaWrapper `json:"salesReturn,omitempty"`
	Supply              *MetaWrapper `json:"supply,omitempty"`
	InvoiceIn           *MetaWrapper `json:"invoiceIn,omitempty"`
	PurchaseOrder       *MetaWrapper `json:"purchaseOrder,omitempty"`
	CommissionReportOut *MetaWrapper `json:"commissionReportOut,omitempty"`
}
