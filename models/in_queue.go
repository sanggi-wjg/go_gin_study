package models

import "time"

type InQueue struct {
	Idx               uint   `gorm:"primaryKey;autoIncrement;"`
	StockCd           string `gorm:"Column:stockCd;size:100;not null;"`
	StockType         string `gorm:"Column:stockType;size:4;not null;"`
	StockTypeName     string `gorm:"Column:stockTypeName;size:255;not null;"`
	PartnerId         string `gorm:"Column:partnerId;size:255;not null;"`
	PartnerName       string `gorm:"Column:partnerName;size:255;not null;"`
	InOrderDate       string `gorm:"Column:inOrderDate;size:100;default:null;"`
	InOrderURL        string `gorm:"Column:inOrderURL;size:255;default:null;"`
	InShippingCompany string `gorm:"Column:inShippingCompany;size:100;default:null;"`
	InWaybillNo       string `gorm:"Column:inWaybillNo;size:255;default:null;"`

	PackageCd            string `gorm:"Column:packageCd;size:100;not null;"`
	ProductOwner         string `gorm:"Column:productOwner;size:100;not null;default:JAMY;"`
	ProductGroup         string `gorm:"Column:productGroup;size:100;not null;"`
	ProductGroupName     string `gorm:"Column:productGroupName;size:255;not null;"`
	ProductGroupNameCh   string `gorm:"Column:productGroupNameCh;size:255;default:null;"`
	PartnerProductCd     string `gorm:"Column:partnerProductCd;size:100;default:null;"`
	PartnerProductOption string `gorm:"Column:partnerProductOption;size:255;default:null;"`

	ProductItemCd      string  `gorm:"Column:productItemCd;size:100;default:null;"`
	ProductCd          string  `gorm:"Column:productCd;size:100;not null;"`
	ProductName        string  `gorm:"Column:productName;size:255;not null;"`
	ProductNameCh      string  `gorm:"Column:productNameCh;size:255;default:null;"`
	ProductNameEn      string  `gorm:"Column:productNameEn;size:255;not null;"`
	ProductOption      string  `gorm:"Column:productOption;size:255;not null;"`
	ProductWeight      float32 `gorm:"Column:productWeight;default:0.0;"`
	ProductImageURL    string  `gorm:"Column:productImageURL;size:255;not null;"`
	ProductSize        string  `gorm:"Column:productSize;size:255;not null;"`
	ProductUnitPrice   string  `gorm:"Column:productUnitPrice;size:255;not null;default:0"`
	ProductNature      string  `gorm:"Column:productNature;size:255;not null;"`
	ProductBrandName   string  `gorm:"Column:productBrandName;size:255;not null;"`
	ProductVendorName  string  `gorm:"Column:productVendorName;size:255;not null;"`
	ProductVendorPrice string  `gorm:"Column:productVendorPrice;size:100;not null;default:0"`
	ProductHSCode      string  `gorm:"Column:productHSCode;size:255;not null;"`
	ProductQuantity    uint    `gorm:"Column:productQuantity;not null;default:1"`

	PartnerUserType      string `gorm:"Column:partnerUserType;size:2;not null;"`
	PartnerUserTypeName  string `gorm:"Column:partnerUserTypeName;size:255;not null;"`
	TransferComapany     string `gorm:"Column:transferComapany;size:100;not null;"`
	TransferComapanyName string `gorm:"Column:transferComapanyName;size:255;not null;"`

	Result     string `gorm:"Column:result;size:255;default:null;"`
	ResultDate uint   `gorm:"Column:resultDate;size:255;default:null;"`
	RegDate    uint   `gorm:"Column:regDate;size:255;not null;"`

	Remark    string `gorm:"Column:remark;default:null"`
	ExtraData string `gorm:"Column:extraData;default:null"`

	BrandProductItemCd          string    `gorm:"Column:brandProductItemCd;size:100;;default:null"`
	BrandProductItemBarcode     string    `gorm:"Column:brandProductItemBarcode;size:100;;default:null"`
	ProductOptionKr             string    `gorm:"Column:productOptionKr;size:255;;default:."`
	PriorityTime                time.Time `gorm:"Column:priorityTime;size:100;;default:null"`
	BrandProductItemBarcodeFlag string    `gorm:"Column:brandProductItemBarcodeFlag;size:1;;default:null"`
	IsCosmeticsOrder            string    `gorm:"Column:isCosmeticsOrder;size:1;;default:null"`
	StockBatchNo                string    `gorm:"Column:stockBatchNo;size:100;;default:null"`
}
