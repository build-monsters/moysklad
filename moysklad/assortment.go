package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
)

// Assortment Ассортимент.
//
// Ключевое слово: assortment
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-assortiment
type Assortment Slice[AssortmentPosition]

// MetaType возвращает тип сущности.
func (Assortment) MetaType() MetaType {
	return MetaTypeAssortment
}

// AssortmentPosition представляет позицию ассортимента.
//
// Создать позицию можно с помощью NewAssortmentPosition, передав в качестве аргумента объект,
// удовлетворяющий интерфейсу AssortmentType.
type AssortmentPosition struct {
	Meta         Meta           `json:"meta"`                   // Метаданные сущности
	Code         string         `json:"code,omitempty"`         // Код сущности
	Description  string         `json:"description,omitempty"`  // Комментарий сущности
	ExternalCode string         `json:"externalCode,omitempty"` // Внешний код сущности
	Name         string         `json:"name,omitempty"`         // Наименование сущности
	Barcodes     Slice[Barcode] `json:"barcodes,omitempty"`     // Штрихкоды
	raw          []byte         // сырые данные для последующей конвертации в нужный тип
	AccountID    uuid.UUID      `json:"accountId,omitempty"` // ID учетной записи
	ID           uuid.UUID      `json:"id,omitempty"`        // ID сущности
}

// AssortmentType описывает типы, которые входят в состав ассортимента.
//
// Возможные типы:
//   - Product 		– Товар
//   - Variant 		– Модификация
//   - Bundle 		– Комплект
//   - Service 		– Услуга
//   - Consignment 	– Серия
type AssortmentType interface {
	Product | Variant | Bundle | Service | Consignment
	MetaOwner
}

// AsAssortmentInterface описывает необходимый метод AsAssortment
type AsAssortmentInterface interface {
	// AsAssortment возвращает указатель на [AssortmentPosition]
	AsAssortment() *AssortmentPosition
}

// NewAssortmentPosition принимает в качестве аргумента объект, удовлетворяющий интерфейсу [AssortmentType].
//
// Возвращает [AssortmentPosition] с заполненным полем Meta.
func NewAssortmentPosition[T AsAssortmentInterface](entity T) *AssortmentPosition {
	return entity.AsAssortment()
}

// String реализует интерфейс [fmt.Stringer].
func (assortmentPosition *AssortmentPosition) String() string {
	return Stringify(assortmentPosition)
}

// MetaType возвращает тип сущности.
func (assortmentPosition AssortmentPosition) MetaType() MetaType {
	return assortmentPosition.Meta.GetType()
}

// GetMeta возвращает Метаданные сущности.
func (assortmentPosition AssortmentPosition) GetMeta() Meta {
	return assortmentPosition.Meta
}

// Raw реализует интерфейс RawMetaTyper
func (assortmentPosition AssortmentPosition) Raw() []byte {
	return assortmentPosition.raw
}

// UnmarshalJSON реализует интерфейс json.Unmarshaler
func (assortmentPosition *AssortmentPosition) UnmarshalJSON(data []byte) error {
	type alias AssortmentPosition
	var t alias
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	t.raw = data
	*assortmentPosition = AssortmentPosition(t)
	return nil
}

// AsProduct пытается привести объект к типу [Product].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает:
//   - указатель на [Product].
//   - nil в случае неудачи.
func (assortmentPosition *AssortmentPosition) AsProduct() *Product {
	return UnmarshalAsType[Product](assortmentPosition)
}

// AsVariant пытается привести объект к типу [Variant].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает:
//   - указатель на [Variant].
//   - nil в случае неудачи.
func (assortmentPosition *AssortmentPosition) AsVariant() *Variant {
	return UnmarshalAsType[Variant](assortmentPosition)
}

// AsBundle пытается привести объект к типу [Bundle].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает:
//   - указатель на [Bundle].
//   - nil в случае неудачи.
func (assortmentPosition *AssortmentPosition) AsBundle() *Bundle {
	return UnmarshalAsType[Bundle](assortmentPosition)
}

// AsService пытается привести объект к типу [Service].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает:
//   - указатель на [Service].
//   - nil в случае неудачи.
func (assortmentPosition *AssortmentPosition) AsService() *Service {
	return UnmarshalAsType[Service](assortmentPosition)
}

// AsConsignment пытается привести объект к типу [Consignment].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает: [Consignment]
//   - указатель на .
//   - nil в случае неудачи.
func (assortmentPosition *AssortmentPosition) AsConsignment() *Consignment {
	return UnmarshalAsType[Consignment](assortmentPosition)
}

// FilterBundle фильтрует позиции по типу [Bundle] (Комплект)
func (assortment Assortment) FilterBundle() Slice[Bundle] {
	return filterType[Bundle](assortment)
}

// FilterProduct фильтрует позиции по типу [Product] (Товар)
func (assortment Assortment) FilterProduct() Slice[Product] {
	return filterType[Product](assortment)
}

// FilterVariant фильтрует позиции по типу [Variant] (Модификация)
func (assortment Assortment) FilterVariant() Slice[Variant] {
	return filterType[Variant](assortment)
}

// FilterConsignment фильтрует позиции по типу [Consignment] (Серия)
func (assortment Assortment) FilterConsignment() Slice[Consignment] {
	return filterType[Consignment](assortment)
}

// FilterService фильтрует позиции по типу [Service] (Услуга)
func (assortment Assortment) FilterService() Slice[Service] {
	return filterType[Service](assortment)
}

// AssortmentSettings Настройки справочника.
//
// Ключевое слово: assortmentsettings
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-assortiment-nastrojki-sprawochnika
type AssortmentSettings struct {
	Meta            *Meta            `json:"meta,omitempty"`            // Метаданные Настроек справочника
	BarcodeRules    *BarcodeRules    `json:"barcodeRules,omitempty"`    // Настройки правил штрихкодов для сущностей справочника
	UniqueCodeRules *UniqueCodeRules `json:"uniqueCodeRules,omitempty"` // Настройки уникальности кода для сущностей справочника
	CreatedShared   *bool            `json:"createdShared,omitempty"`   // Создавать новые документы с меткой «Общий»
}

// GetMeta возвращает Метаданные Настроек справочника.
func (assortmentSettings AssortmentSettings) GetMeta() Meta {
	return Deref(assortmentSettings.Meta)
}

// GetBarcodeRules возвращает Настройки правил штрихкодов для сущностей справочника.
func (assortmentSettings AssortmentSettings) GetBarcodeRules() BarcodeRules {
	return Deref(assortmentSettings.BarcodeRules)
}

// GetUniqueCodeRules возвращает Настройки уникальности кода для сущностей справочника.
func (assortmentSettings AssortmentSettings) GetUniqueCodeRules() UniqueCodeRules {
	return Deref(assortmentSettings.UniqueCodeRules)
}

// GetCreatedShared возвращает true, если новые документы создаются с пометкой «Общий».
func (assortmentSettings AssortmentSettings) GetCreatedShared() bool {
	return Deref(assortmentSettings.CreatedShared)
}

// SetBarcodeRules устанавливает Настройки правил штрихкодов для сущностей справочника.
func (assortmentSettings *AssortmentSettings) SetBarcodeRules(barcodeRules *BarcodeRules) *AssortmentSettings {
	assortmentSettings.BarcodeRules = barcodeRules
	return assortmentSettings
}

// SetUniqueCodeRules устанавливает Настройки уникальности кода для сущностей справочника.
func (assortmentSettings *AssortmentSettings) SetUniqueCodeRules(uniqueCodeRules *UniqueCodeRules) *AssortmentSettings {
	assortmentSettings.UniqueCodeRules = uniqueCodeRules
	return assortmentSettings
}

// SetCreatedShared устанавливает значение создания новых документов с пометкой «Общий».
func (assortmentSettings *AssortmentSettings) SetCreatedShared(createdShared bool) *AssortmentSettings {
	assortmentSettings.CreatedShared = &createdShared
	return assortmentSettings
}

// String реализует интерфейс [fmt.Stringer].
func (assortmentSettings AssortmentSettings) String() string {
	return Stringify(assortmentSettings)
}

// MetaType возвращает тип сущности.
func (AssortmentSettings) MetaType() MetaType {
	return MetaTypeAssortmentSettings
}

// BarcodeRules Настройки правил штрихкодов для сущностей справочника.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-assortiment-atributy-wlozhennyh-suschnostej-nastrojki-prawil-shtrihkodow-dlq-suschnostej-sprawochnika
type BarcodeRules struct {
	FillEAN13Barcode    *bool `json:"fillEAN13Barcode,omitempty"`    // Автоматически создавать штрихкод EAN13 для новых товаров, комплектов, модификаций и услуг
	WeightBarcode       *bool `json:"weightBarcode,omitempty"`       // Использовать префиксы штрихкодов для весовых товаров
	WeightBarcodePrefix *int  `json:"weightBarcodePrefix,omitempty"` // Префикс штрихкодов для весовых товаров. Возможные значения: число формата X или XX
}

// GetFillEAN13Barcode возвращает true, если штрихкод EAN13 для новых товаров, комплектов, модификаций и услуг создаётся автоматически.
func (barcodeRules BarcodeRules) GetFillEAN13Barcode() bool {
	return Deref(barcodeRules.FillEAN13Barcode)
}

// GetWeightBarcode возвращает true, если используются префиксы штрихкодов для весовых товаров.
func (barcodeRules BarcodeRules) GetWeightBarcode() bool {
	return Deref(barcodeRules.WeightBarcode)
}

// GetWeightBarcodePrefix возвращает Префикс штрихкодов для весовых товаров. Возможные значения: число формата X или XX.
func (barcodeRules BarcodeRules) GetWeightBarcodePrefix() int {
	return Deref(barcodeRules.WeightBarcodePrefix)
}

// SetFillEAN13Barcode устанавливает значение автоматического создания штрихкода EAN13 для новых товаров, комплектов, модификаций и услуг.
func (barcodeRules *BarcodeRules) SetFillEAN13Barcode(fillEAN13Barcode bool) *BarcodeRules {
	barcodeRules.FillEAN13Barcode = &fillEAN13Barcode
	return barcodeRules
}

// SetWeightBarcode устанавливает значение использования префиксов штрихкодов для весовых товаров.
func (barcodeRules *BarcodeRules) SetWeightBarcode(weightBarcode bool) *BarcodeRules {
	barcodeRules.WeightBarcode = &weightBarcode
	return barcodeRules
}

// SetWeightBarcodePrefix устанавливает Префикс штрихкодов для весовых товаров.
func (barcodeRules *BarcodeRules) SetWeightBarcodePrefix(weightBarcodePrefix int) *BarcodeRules {
	barcodeRules.WeightBarcodePrefix = &weightBarcodePrefix
	return barcodeRules
}

// String реализует интерфейс [fmt.Stringer].
func (barcodeRules BarcodeRules) String() string {
	return Stringify(barcodeRules)
}

// AssortmentResponse объект ответа на запрос получения ассортимента.
type AssortmentResponse struct {
	Context Context        `json:"context,omitempty"` // Информация о сотруднике, выполнившем запрос
	Rows    Assortment     `json:"rows,omitempty"`    // Список товаров, услуг, комплектов, модификаций и серий
	Meta    MetaCollection `json:"meta,omitempty"`    // Информация о контексте запроса
}

// String реализует интерфейс [fmt.Stringer].
func (assortmentResponse AssortmentResponse) String() string {
	return Stringify(assortmentResponse)
}

// PaymentItem Признак предмета расчета.
//
// Возможные варианты:
//   - PaymentItemGood                – Товар
//   - PaymentItemExcisableGood       – Подакцизный товар
//   - PaymentItemCompoundPaymentItem – Составной предмет расчета
//   - PaymentItemAnotherPaymentItem  – Иной предмет расчета
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-suschnosti-priznak-predmeta-rascheta
type PaymentItem string

const (
	PaymentItemGood                PaymentItem = "GOOD"                  // Товар
	PaymentItemExcisableGood       PaymentItem = "EXCISABLE_GOOD"        // Подакцизный товар
	PaymentItemCompoundPaymentItem PaymentItem = "COMPOUND_PAYMENT_ITEM" // Составной предмет расчета
	PaymentItemAnotherPaymentItem  PaymentItem = "ANOTHER_PAYMENT_ITEM"  // Иной предмет расчета
)

// TrackingType Тип маркируемой продукции.
//
// Возможные варианты:
//   - TrackingTypeBeerAlcohol 		– Пиво и слабоалкогольная продукция
//   - TrackingTypeElectronics 		– Фотокамеры и лампы-вспышки
//   - TrackingTypeFoodSupplement 	– Биологически активные добавки к пище
//   - TrackingTypeClothes 			– Тип маркировки "Одежда"
//   - TrackingTypeLinens 			– Тип маркировки "Постельное белье"
//   - TrackingTypeMedicalDevices 	– Медизделия и кресла-коляски
//   - TrackingTypeMilk 			– Молочная продукция
//   - TrackingTypeNcp 				– Никотиносодержащая продукция
//   - TrackingTypeNotTracked 		– Без маркировки
//   - TrackingTypeOtp 				– Альтернативная табачная продукция
//   - TrackingTypePerfumery 		– Духи и туалетная вода
//   - TrackingTypeSanitizer 		– Антисептики
//   - TrackingTypeShoes 			– Тип маркировки "Обувь"
//   - TrackingTypeTires 			– Шины и покрышки
//   - TrackingTypeTobacco 			– Тип маркировки "Табак"
//   - TrackingTypeWater 			– Упакованная вода
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-suschnosti-tip-markiruemoj-produkcii
type TrackingType string

const (
	TrackingTypeBeerAlcohol    TrackingType = "BEER_ALCOHOL"    // Пиво и слабоалкогольная продукция
	TrackingTypeElectronics    TrackingType = "ELECTRONICS"     // Фотокамеры и лампы-вспышки
	TrackingTypeFoodSupplement TrackingType = "FOOD_SUPPLEMENT" // Биологически активные добавки к пище
	TrackingTypeClothes        TrackingType = "LP_CLOTHES"      // Тип маркировки "Одежда"
	TrackingTypeLinens         TrackingType = "LP_LINENS"       // Тип маркировки "Постельное белье"
	TrackingTypeMedicalDevices TrackingType = "MEDICAL_DEVICES" // Медизделия и кресла-коляски
	TrackingTypeMilk           TrackingType = "MILK"            // Молочная продукция
	TrackingTypeNcp            TrackingType = "NCP"             // Никотиносодержащая продукция
	TrackingTypeNotTracked     TrackingType = "NOT_TRACKED"     // Без маркировки
	TrackingTypeOtp            TrackingType = "OTP"             // Альтернативная табачная продукция
	TrackingTypePerfumery      TrackingType = "PERFUMERY"       // Духи и туалетная вода
	TrackingTypeSanitizer      TrackingType = "SANITIZER"       // Антисептики
	TrackingTypeShoes          TrackingType = "SHOES"           // Тип маркировки "Обувь"
	TrackingTypeTires          TrackingType = "TIRES"           // Шины и покрышки
	TrackingTypeTobacco        TrackingType = "TOBACCO"         // Тип маркировки "Табак"
	TrackingTypeWater          TrackingType = "WATER"           // Упакованная вода
)

// TaxSystem Код системы налогообложения.
//
// Возможные варианты:
//   - TaxSystemGeneral                 – ОСН
//   - TaxSystemSimplifiedIncome        – УСН. Доход
//   - TaxSystemSimplifiedIncomeOutcome – УСН. Доход-Расход
//   - TaxSystemUnifiedAgricultural     – ЕСХН
//   - TaxSystemPresumptive             – ЕНВД
//   - TaxSystemPatentBased             – Патент
//   - TaxSystemSameAsGroup             – Совпадает с группой
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-suschnosti-kod-sistemy-nalogooblozheniq
type TaxSystem string

const (
	TaxSystemGeneral                 TaxSystem = "GENERAL_TAX_SYSTEM"                   // ОСН
	TaxSystemSimplifiedIncome        TaxSystem = "SIMPLIFIED_TAX_SYSTEM_INCOME"         // УСН. Доход
	TaxSystemSimplifiedIncomeOutcome TaxSystem = "SIMPLIFIED_TAX_SYSTEM_INCOME_OUTCOME" // УСН. Доход-Расход
	TaxSystemUnifiedAgricultural     TaxSystem = "UNIFIED_AGRICULTURAL_TAX"             // ЕСХН
	TaxSystemPresumptive             TaxSystem = "PRESUMPTIVE_TAX_SYSTEM"               // ЕНВД
	TaxSystemPatentBased             TaxSystem = "PATENT_BASED"                         // Патент
	TaxSystemSameAsGroup             TaxSystem = "TAX_SYSTEM_SAME_AS_GROUP"             // Совпадает с группой
)

type assortmentService struct {
	Endpoint
}

func (service *assortmentService) Get(ctx context.Context, params ...*Params) (*AssortmentResponse, *resty.Response, error) {
	return NewRequestBuilder[AssortmentResponse](service.client, service.uri).SetParams(params...).Get(ctx)
}

func (service *assortmentService) GetAsync(ctx context.Context, params ...*Params) (AsyncResultService[AssortmentResponse], *resty.Response, error) {
	p := new(Params)
	if len(params) > 0 {
		p = params[0]
	}
	p.withAsync()
	_, resp, err := NewRequestBuilder[any](service.client, service.uri).SetParams(p).Get(ctx)
	if err != nil {
		return nil, resp, nil
	}
	async := NewAsyncResultService[AssortmentResponse](service.client, resp)
	return async, resp, err
}

func (service *assortmentService) DeleteMany(ctx context.Context, entities ...AsAssortmentInterface) (*DeleteManyResponse, *resty.Response, error) {
	var mw = make([]MetaWrapper, 0, len(entities))
	for _, entity := range entities {
		mw = append(mw, (entity).AsAssortment().GetMeta().Wrap())
	}
	return NewRequestBuilder[DeleteManyResponse](service.client, service.uri).Post(ctx, mw)
}

func (service *assortmentService) GetSettings(ctx context.Context) (*AssortmentSettings, *resty.Response, error) {
	path := fmt.Sprintf("%s/settings", service.uri)
	return NewRequestBuilder[AssortmentSettings](service.client, path).Get(ctx)
}

func (service *assortmentService) UpdateSettings(ctx context.Context, settings *AssortmentSettings) (*AssortmentSettings, *resty.Response, error) {
	path := fmt.Sprintf("%s/settings", service.uri)
	return NewRequestBuilder[AssortmentSettings](service.client, path).Put(ctx, settings)
}

// AssortmentService
// Сервис для работы с ассортиментом.
type AssortmentService interface {
	// Get выполняет запрос на получение всех товаров, услуг, комплектов, модификаций и серий в виде списка.
	// Принимает контекст context.Context и опционально объект параметров запроса Params.
	// Возвращает объект AssortmentResponse.
	Get(ctx context.Context, params ...*Params) (*AssortmentResponse, *resty.Response, error)

	// GetAsync выполняет асинхронный запрос на получение всех товаров, услуг, комплектов, модификаций и серий в виде списка.
	// Принимает контекст context.Context и опционально объект параметров запроса Params.
	// Возвращает готовый сервис AsyncResultService для обработки данного запроса.
	GetAsync(ctx context.Context, params ...*Params) (AsyncResultService[AssortmentResponse], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление позиций в Ассортименте.
	// Принимает контекст context.Context и множество объектов, реализующих интерфейс AsAssortmentInterface.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...AsAssortmentInterface) (*DeleteManyResponse, *resty.Response, error)

	// GetSettings выполняет запрос на получение настроек справочника ассортимента.
	// Принимает контекст context.Context.
	// Возвращает объект AssortmentSettings.
	GetSettings(ctx context.Context) (*AssortmentSettings, *resty.Response, error)

	// UpdateSettings выполняет запрос на изменение метаданных справочника ассортимента.
	// Принимает контекст context.Context и объект AssortmentSettings.
	// Возвращает обновлённый объект AssortmentSettings.
	UpdateSettings(ctx context.Context, settings *AssortmentSettings) (*AssortmentSettings, *resty.Response, error)
}

// NewAssortmentService возвращает сервис для работы с ассортиментом.
func NewAssortmentService(client *Client) AssortmentService {
	e := NewEndpoint(client, "entity/assortment")
	return &assortmentService{e}
}
