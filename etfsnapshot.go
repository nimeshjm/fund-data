package main

type EtfSnapshot struct {
	Dbgtime             string `json:"dbgtime"`
	ID                  string `json:"Id"`
	InceptionDate       string `json:"InceptionDate"`
	Isin                string `json:"Isin"`
	Sedol               string `json:"Sedol"`
	InvestmentType      string `json:"InvestmentType"`
	HoldingType         string `json:"HoldingType"`
	BrandingCompanyID   string `json:"BrandingCompanyId"`
	BrandingCompanyName string `json:"BrandingCompanyName"`
	Type                string `json:"Type"`
	Name                string `json:"Name"`
	LegalName           string `json:"LegalName"`
	FundAttributes      struct {
		UCITS string `json:"UCITS"`
	} `json:"FundAttributes"`
	FundNetAssetValues []struct {
		CurrencyID    string  `json:"CurrencyId"`
		DayEndDate    string  `json:"DayEndDate"`
		DayEndValue   float64 `json:"DayEndValue"`
		MonthEndDate  string  `json:"MonthEndDate"`
		MonthEndValue float64 `json:"MonthEndValue"`
	} `json:"FundNetAssetValues"`
	NetAssetValues []struct {
		CurrencyID    string  `json:"CurrencyId"`
		DayEndDate    string  `json:"DayEndDate"`
		DayEndValue   float64 `json:"DayEndValue"`
		MonthEndDate  string  `json:"MonthEndDate"`
		MonthEndValue float64 `json:"MonthEndValue"`
	} `json:"NetAssetValues"`
	Domicile         string  `json:"Domicile"`
	RealtimeID       string  `json:"RealtimeId"`
	NetChange        float64 `json:"NetChange"`
	NetChangePercent float64 `json:"NetChangePercent"`
	OngoingCharge    string  `json:"OngoingCharge"`
	CollectedSRRI    struct {
		Date string `json:"Date"`
		Rank int    `json:"Rank"`
	} `json:"CollectedSRRI"`
	DividendDistributionFrequency string  `json:"DividendDistributionFrequency"`
	OpenPrice                     float64 `json:"OpenPrice"`
	HighPrice                     float64 `json:"HighPrice"`
	LowPrice                      float64 `json:"LowPrice"`
	PreviousClose                 float64 `json:"PreviousClose"`
	Rolling52WeekHigh             float64 `json:"Rolling52WeekHigh"`
	Rolling52WeekHighDate         string  `json:"Rolling52WeekHighDate"`
	Rolling52WeekLow              float64 `json:"Rolling52WeekLow"`
	Rolling52WeekLowDate          string  `json:"Rolling52WeekLowDate"`
	Calendar52WeekHigh            float64 `json:"Calendar52WeekHigh"`
	Calendar52WeekHighDate        string  `json:"Calendar52WeekHighDate"`
	Calendar52WeekLow             float64 `json:"Calendar52WeekLow"`
	Calendar52WeekLowDate         string  `json:"Calendar52WeekLowDate"`
	YieldHistory                  struct {
		Type    string  `json:"Type"`
		Value   float64 `json:"Value"`
		EndDate string  `json:"EndDate"`
	} `json:"YieldHistory"`
	Custom struct {
		CustomCategoryIDName                  string  `json:"CustomCategoryIdName"`
		CustomText1                           string  `json:"CustomText1"`
		CustomAdditionalBuyFee                float64 `json:"CustomAdditionalBuyFee"`
		CustomValue2                          float64 `json:"CustomValue2"`
		CustomMinimumAdditionalPurchaseAmount float64 `json:"CustomMinimumAdditionalPurchaseAmount"`
		CustomMinimumPurchaseAmount           float64 `json:"CustomMinimumPurchaseAmount"`
		CustomSellFee                         float64 `json:"CustomSellFee"`
		CustomHistoricYield                   float64 `json:"CustomHistoricYield"`
	} `json:"Custom"`
	TradingCurrency struct {
		ID string `json:"Id"`
	} `json:"TradingCurrency"`
	Currency struct {
		ID string `json:"Id"`
	} `json:"Currency"`
	IMASector struct {
		ID   string `json:"Id"`
		Name string `json:"Name"`
	} `json:"IMASector"`
	CategoryBroadAssetClass struct {
		ID   string `json:"Id"`
		Name string `json:"Name"`
	} `json:"CategoryBroadAssetClass"`
	RiskAndRating []struct {
		Date                   string  `json:"Date"`
		TimePeriod             string  `json:"TimePeriod"`
		RatingValue            int     `json:"RatingValue"`
		RiskAdjustedReturn     float64 `json:"RiskAdjustedReturn"`
		RiskRatingValue        int     `json:"RiskRatingValue"`
		RiskScore              float64 `json:"RiskScore"`
		CategoryRank           float64 `json:"CategoryRank,omitempty"`
		PerformanceRatingValue int     `json:"PerformanceRatingValue"`
	} `json:"RiskAndRating"`
	PurchaseDetails []struct {
		Currency struct {
			ID string `json:"Id"`
		} `json:"Currency"`
		InitialInvestment struct {
			Unit  int     `json:"Unit"`
			Value float64 `json:"Value"`
		} `json:"InitialInvestment"`
		SubsequentInvestment struct {
			Unit  int     `json:"Unit"`
			Value float64 `json:"Value"`
		} `json:"SubsequentInvestment"`
		FrontLoad struct {
			Unit  int     `json:"Unit"`
			Value float64 `json:"Value"`
		} `json:"FrontLoad"`
		RedemptionFee struct {
			Unit  int     `json:"Unit"`
			Value float64 `json:"Value"`
		} `json:"RedemptionFee"`
	} `json:"PurchaseDetails"`
	LastPrice struct {
		Date           string  `json:"Date"`
		MarketDate     string  `json:"MarketDate"`
		MarketTime     string  `json:"MarketTime"`
		MarketTimeZone string  `json:"MarketTimeZone"`
		Value          float64 `json:"Value"`
		Currency       struct {
			ID string `json:"Id"`
		} `json:"Currency"`
	} `json:"LastPrice"`
	LegalStructure      string  `json:"LegalStructure"`
	FiscalYearEndMonth  int     `json:"FiscalYearEndMonth"`
	InvestmentStrategy  string  `json:"InvestmentStrategy"`
	ManagementFee       float64 `json:"ManagementFee"`
	ActualManagementFee float64 `json:"ActualManagementFee"`
	ProviderCompany     struct {
		Name         string `json:"Name"`
		AddressLine1 string `json:"AddressLine1"`
		AddressLine2 string `json:"AddressLine2"`
		Phone        string `json:"Phone"`
		City         string `json:"City"`
		Country      string `json:"Country"`
		PostalCode   string `json:"PostalCode"`
		Homepage     string `json:"Homepage"`
	} `json:"ProviderCompany"`
	AdministratorCompany struct {
		Name         string `json:"Name"`
		AddressLine1 string `json:"AddressLine1"`
		AddressLine2 string `json:"AddressLine2"`
		Phone        string `json:"Phone"`
		City         string `json:"City"`
		Country      string `json:"Country"`
		PostalCode   string `json:"PostalCode"`
		Homepage     string `json:"Homepage"`
	} `json:"AdministratorCompany"`
	ManagerList []struct {
		DisplayPreference        float64 `json:"DisplayPreference"`
		GivenName                string  `json:"GivenName"`
		FamilyName               string  `json:"FamilyName"`
		StartDate                string  `json:"StartDate"`
		ManagerProvidedBiography string  `json:"ManagerProvidedBiography"`
	} `json:"ManagerList"`
	TrailingPerformance []struct {
		CurrencyID string `json:"CurrencyId"`
		Date       string `json:"Date"`
		Type       string `json:"Type"`
		ReturnType string `json:"ReturnType"`
		Return     []struct {
			Value      float64 `json:"Value"`
			Date       string  `json:"Date"`
			TimePeriod string  `json:"TimePeriod"`
			Annualized bool    `json:"Annualized,omitempty"`
		} `json:"Return"`
	} `json:"TrailingPerformance"`
	RiskStatistics []struct {
		CurrencyID     string `json:"CurrencyId"`
		CurrencyOption string `json:"CurrencyOption"`
		Date           string `json:"Date"`
		Type           string `json:"Type"`
		Alphas         []struct {
			Value      float64 `json:"Value"`
			Frequency  string  `json:"Frequency"`
			TimePeriod string  `json:"TimePeriod"`
		} `json:"Alphas"`
		Betas []struct {
			Value      float64 `json:"Value"`
			Frequency  string  `json:"Frequency"`
			TimePeriod string  `json:"TimePeriod"`
		} `json:"Betas"`
		RSquareds []struct {
			Value      float64 `json:"Value"`
			Frequency  string  `json:"Frequency"`
			TimePeriod string  `json:"TimePeriod"`
		} `json:"RSquareds"`
		InformationRatios []struct {
			Value      float64 `json:"Value"`
			Frequency  string  `json:"Frequency"`
			TimePeriod string  `json:"TimePeriod"`
		} `json:"InformationRatios"`
		TrackingErrors []struct {
			Value      float64 `json:"Value"`
			Frequency  string  `json:"Frequency"`
			TimePeriod string  `json:"TimePeriod"`
		} `json:"TrackingErrors"`
		ArithmeticMeans []struct {
			Value      float64 `json:"Value"`
			Frequency  string  `json:"Frequency"`
			TimePeriod string  `json:"TimePeriod"`
		} `json:"ArithmeticMeans,omitempty"`
		StandardDeviations []struct {
			Value      float64 `json:"Value"`
			Frequency  string  `json:"Frequency"`
			TimePeriod string  `json:"TimePeriod"`
		} `json:"StandardDeviations"`
		SharpeRatios []struct {
			Frequency  string  `json:"Frequency"`
			TimePeriod string  `json:"TimePeriod"`
			Value      float64 `json:"Value"`
		} `json:"SharpeRatios"`
		MaximumReturns []struct {
			Value      float64 `json:"Value"`
			Frequency  string  `json:"Frequency"`
			TimePeriod string  `json:"TimePeriod"`
			Date       string  `json:"Date"`
		} `json:"MaximumReturns"`
		MinimumReturns []struct {
			Value      float64 `json:"Value"`
			Frequency  string  `json:"Frequency"`
			TimePeriod string  `json:"TimePeriod"`
			Date       string  `json:"Date"`
		} `json:"MinimumReturns"`
		NumberOfPositives []struct {
			Count      float64 `json:"Count"`
			Frequency  string  `json:"Frequency"`
			TimePeriod string  `json:"TimePeriod"`
		} `json:"NumberOfPositives"`
		NumberOfNegatives []struct {
			Count      float64 `json:"Count"`
			Frequency  string  `json:"Frequency"`
			TimePeriod string  `json:"TimePeriod"`
		} `json:"NumberOfNegatives"`
	} `json:"RiskStatistics"`
	HistoricalPerformanceSeries []struct {
		CurrencyID string `json:"CurrencyId"`
		StartDate  string `json:"StartDate"`
		Frequency  string `json:"Frequency"`
		TimePeriod string `json:"TimePeriod"`
		ReturnType string `json:"ReturnType"`
		Return     []struct {
			Value      float64 `json:"Value"`
			Date       string  `json:"Date"`
			TimePeriod string  `json:"TimePeriod"`
		} `json:"Return"`
	} `json:"HistoricalPerformanceSeries"`
	GrowthOf10K []struct {
		CurrencyID     string  `json:"CurrencyId"`
		NumberOfYears  float64 `json:"NumberOfYears"`
		StartValue     float64 `json:"StartValue"`
		HistoryDetails []struct {
			EndDate string  `json:"EndDate"`
			Value   float64 `json:"Value"`
		} `json:"HistoryDetails"`
		PeriodCalculation string `json:"PeriodCalculation"`
	} `json:"GrowthOf10K"`
	Benchmark []struct {
		ID                  string `json:"Id"`
		GlobalAssetClassID  string `json:"GlobalAssetClassId,omitempty"`
		Type                string `json:"Type"`
		Name                string `json:"Name"`
		TrailingPerformance []struct {
			CurrencyID string `json:"CurrencyId"`
			Date       string `json:"Date"`
			Type       string `json:"Type"`
			ReturnType string `json:"ReturnType"`
			Return     []struct {
				Value      float64 `json:"Value"`
				Annualized bool    `json:"Annualized,omitempty"`
				Date       string  `json:"Date"`
				TimePeriod string  `json:"TimePeriod"`
			} `json:"Return"`
		} `json:"TrailingPerformance,omitempty"`
		HistoricalPerformanceSeries []struct {
			CurrencyID string `json:"CurrencyId"`
			StartDate  string `json:"StartDate"`
			Frequency  string `json:"Frequency"`
			TimePeriod string `json:"TimePeriod"`
			ReturnType string `json:"ReturnType"`
			Return     []struct {
				Value      float64 `json:"Value"`
				Date       string  `json:"Date"`
				TimePeriod string  `json:"TimePeriod"`
			} `json:"Return"`
		} `json:"HistoricalPerformanceSeries,omitempty"`
		GrowthOf10K []struct {
			CurrencyID     string  `json:"CurrencyId"`
			NumberOfYears  float64 `json:"NumberOfYears"`
			StartValue     float64 `json:"StartValue"`
			HistoryDetails []struct {
				EndDate string  `json:"EndDate"`
				Value   float64 `json:"Value"`
			} `json:"HistoryDetails"`
			PeriodCalculation string `json:"PeriodCalculation"`
		} `json:"GrowthOf10K,omitempty"`
		Portfolios []struct {
			Date                  string `json:"Date"`
			SuppressionExpiration string `json:"SuppressionExpiration"`
			Holdings              string `json:"Holdings"`
			ID                    int    `json:"Id"`
			TotalMarketValues     []struct {
				Value        float64 `json:"Value"`
				CurrencyID   string  `json:"CurrencyId"`
				SalePosition string  `json:"SalePosition"`
			} `json:"TotalMarketValues"`
			AssetAllocations []struct {
				Type            string `json:"Type"`
				SalePosition    string `json:"SalePosition"`
				BreakdownValues []struct {
					Value float64 `json:"Value,omitempty"`
					Type  string  `json:"Type"`
				} `json:"BreakdownValues"`
			} `json:"AssetAllocations"`
			AssetAllocationBreakdown []struct {
				SalePosition    string `json:"SalePosition"`
				BreakdownValues []struct {
					Value float64 `json:"Value,omitempty"`
					Type  string  `json:"Type"`
				} `json:"BreakdownValues"`
				NotClassified float64 `json:"NotClassified,omitempty"`
			} `json:"AssetAllocationBreakdown"`
			SpecialBreakdown []struct {
				NotClassified   float64 `json:"NotClassified"`
				SalePosition    string  `json:"SalePosition"`
				BreakdownValues []struct {
					Value float64 `json:"Value,omitempty"`
					Type  string  `json:"Type"`
				} `json:"BreakdownValues"`
			} `json:"SpecialBreakdown"`
			RegionalExposure []struct {
				NotClassified   float64 `json:"NotClassified"`
				SalePosition    string  `json:"SalePosition"`
				BreakdownValues []struct {
					Value float64 `json:"Value,omitempty"`
					Type  string  `json:"Type"`
				} `json:"BreakdownValues"`
			} `json:"RegionalExposure"`
			MarketCapitalBreakdown []struct {
				NotClassified   float64 `json:"NotClassified"`
				SalePosition    string  `json:"SalePosition"`
				BreakdownValues []struct {
					Value float64 `json:"Value"`
					Type  string  `json:"Type"`
				} `json:"BreakdownValues"`
			} `json:"MarketCapitalBreakdown"`
			StyleBoxBreakdown []struct {
				NotClassified   float64 `json:"NotClassified"`
				SalePosition    string  `json:"SalePosition"`
				BreakdownValues []struct {
					Value float64 `json:"Value"`
					Type  string  `json:"Type"`
				} `json:"BreakdownValues"`
			} `json:"StyleBoxBreakdown"`
			GlobalStockSectorBreakdown []struct {
				NotClassified   float64 `json:"NotClassified"`
				SalePosition    string  `json:"SalePosition"`
				BreakdownValues []struct {
					Value float64 `json:"Value"`
					Type  string  `json:"Type"`
				} `json:"BreakdownValues"`
			} `json:"GlobalStockSectorBreakdown"`
			IndustryGroupBreakdown []struct {
				NotClassified   float64 `json:"NotClassified"`
				SalePosition    string  `json:"SalePosition"`
				BreakdownValues []struct {
					Value float64 `json:"Value,omitempty"`
					Type  string  `json:"Type"`
				} `json:"BreakdownValues"`
			} `json:"IndustryGroupBreakdown"`
			CashBreakdown []struct {
				NotClassified   float64 `json:"NotClassified"`
				SalePosition    string  `json:"SalePosition"`
				BreakdownValues []struct {
					Value float64 `json:"Value"`
					Type  string  `json:"Type"`
				} `json:"BreakdownValues"`
			} `json:"CashBreakdown"`
			GlobalBondSectorBreakdownLevel1 []struct {
				NotClassified   float64 `json:"NotClassified"`
				SalePosition    string  `json:"SalePosition"`
				BreakdownValues []struct {
					Value float64 `json:"Value"`
					Type  string  `json:"Type"`
				} `json:"BreakdownValues"`
			} `json:"GlobalBondSectorBreakdownLevel1"`
			GlobalBondSectorBreakdownLevel2 []struct {
				NotClassified   float64 `json:"NotClassified"`
				SalePosition    string  `json:"SalePosition"`
				BreakdownValues []struct {
					Value float64 `json:"Value"`
					Type  string  `json:"Type"`
				} `json:"BreakdownValues"`
			} `json:"GlobalBondSectorBreakdownLevel2"`
			GlobalBondSectorBreakdownLevel3 []struct {
				NotClassified   float64 `json:"NotClassified"`
				SalePosition    string  `json:"SalePosition"`
				BreakdownValues []struct {
					Value float64 `json:"Value,omitempty"`
					Type  string  `json:"Type"`
				} `json:"BreakdownValues"`
			} `json:"GlobalBondSectorBreakdownLevel3"`
			MaturityRange []struct {
				NotClassified   float64 `json:"NotClassified"`
				SalePosition    string  `json:"SalePosition"`
				BreakdownValues []struct {
					Value float64 `json:"Value,omitempty"`
					Type  string  `json:"Type"`
				} `json:"BreakdownValues"`
			} `json:"MaturityRange"`
			CouponRange []struct {
				NotClassified   float64 `json:"NotClassified"`
				SalePosition    string  `json:"SalePosition"`
				BreakdownValues []struct {
					Value float64 `json:"Value,omitempty"`
					Type  string  `json:"Type"`
				} `json:"BreakdownValues"`
			} `json:"CouponRange"`
			CreditQualityBreakdown []struct {
				SalePosition    string `json:"SalePosition"`
				BreakdownValues []struct {
					Value float64 `json:"Value,omitempty"`
					Type  string  `json:"Type"`
				} `json:"BreakdownValues"`
				NotClassified float64 `json:"NotClassified,omitempty"`
			} `json:"CreditQualityBreakdown"`
			CountryExposure []struct {
				NotClassified   float64 `json:"NotClassified"`
				SalePosition    string  `json:"SalePosition"`
				BreakdownValues []struct {
					Value float64 `json:"Value"`
					Type  string  `json:"Type"`
				} `json:"BreakdownValues"`
			} `json:"CountryExposure"`
			EquityStatistics struct {
				StyleBox                      int     `json:"StyleBox"`
				StyleScore                    float64 `json:"StyleScore"`
				SizeScore                     float64 `json:"SizeScore"`
				ProspectiveEarningsYield      float64 `json:"ProspectiveEarningsYield"`
				ProspectiveBookValueYield     float64 `json:"ProspectiveBookValueYield"`
				ProspectiveRevenueYield       float64 `json:"ProspectiveRevenueYield"`
				ProspectiveCashFlowYield      float64 `json:"ProspectiveCashFlowYield"`
				ProspectiveDividendYield      float64 `json:"ProspectiveDividendYield"`
				Forecasted5YearEarningsGrowth float64 `json:"Forecasted5YearEarningsGrowth"`
				ForecastedEarningsGrowth      float64 `json:"ForecastedEarningsGrowth"`
				ForecastedBookValueGrowth     float64 `json:"ForecastedBookValueGrowth"`
				ForecastedRevenueGrowth       float64 `json:"ForecastedRevenueGrowth"`
				ForecastedCashFlowGrowth      float64 `json:"ForecastedCashFlowGrowth"`
				MarketCapital                 float64 `json:"MarketCapital"`
				DTC                           float64 `json:"DTC"`
				ROA                           float64 `json:"ROA"`
				ROE                           float64 `json:"ROE"`
				NetMargin                     float64 `json:"NetMargin"`
				EarningsYield                 float64 `json:"EarningsYield"`
				BookValueYield                float64 `json:"BookValueYield"`
				RevenueYield                  float64 `json:"RevenueYield"`
				CashFlowYield                 float64 `json:"CashFlowYield"`
				Past3YearEarningsGrowth       float64 `json:"Past3YearEarningsGrowth"`
				EconomicMoatScore             float64 `json:"EconomicMoatScore"`
				FreeCashFlowYield             float64 `json:"FreeCashFlowYield"`
				ROIC                          float64 `json:"ROIC"`
				FinancialHealthGrade          float64 `json:"FinancialHealthGrade"`
				ProfitabilityGrade            float64 `json:"ProfitabilityGrade"`
				GrowthGrade                   float64 `json:"GrowthGrade"`
			} `json:"EquityStatistics"`
			PortfolioHoldings []struct {
				ID                  string  `json:"Id"`
				DetailHoldingTypeID string  `json:"DetailHoldingTypeId"`
				ExternalName        string  `json:"ExternalName"`
				PerformanceID       string  `json:"PerformanceId,omitempty"`
				ISIN                string  `json:"ISIN"`
				CurrencyID          string  `json:"CurrencyId"`
				CUSIP               string  `json:"CUSIP,omitempty"`
				CountryID           string  `json:"CountryId"`
				SecurityName        string  `json:"SecurityName"`
				Weighting           float64 `json:"Weighting"`
				IndustryID          int     `json:"IndustryId,omitempty"`
				MarketValue         float64 `json:"MarketValue"`
				GlobalSectorID      string  `json:"GlobalSectorId,omitempty"`
				NumberOfShare       int     `json:"NumberOfShare"`
				ShareChange         int     `json:"ShareChange"`
			} `json:"PortfolioHoldings"`
			HoldingAggregates []struct {
				NumberOfHolding int    `json:"NumberOfHolding"`
				SalePosition    string `json:"SalePosition"`
				HoldingsType    string `json:"HoldingsType,omitempty"`
			} `json:"HoldingAggregates"`
		} `json:"Portfolios,omitempty"`
	} `json:"Benchmark"`
	FundBenchmark []struct {
		ID        string  `json:"Id"`
		Name      string  `json:"Name"`
		Weighting float64 `json:"Weighting"`
	} `json:"FundBenchmark"`
	Portfolios []struct {
		Date                  string `json:"Date"`
		SuppressionExpiration string `json:"SuppressionExpiration"`
		Holdings              string `json:"Holdings"`
		ID                    int    `json:"Id"`
		TotalMarketValues     []struct {
			Value        float64 `json:"Value"`
			CurrencyID   string  `json:"CurrencyId"`
			SalePosition string  `json:"SalePosition"`
		} `json:"TotalMarketValues"`
		AssetAllocations []struct {
			Type            string `json:"Type"`
			SalePosition    string `json:"SalePosition"`
			BreakdownValues []struct {
				Value float64 `json:"Value,omitempty"`
				Type  string  `json:"Type"`
			} `json:"BreakdownValues"`
		} `json:"AssetAllocations"`
		AssetAllocationBreakdown []struct {
			SalePosition    string `json:"SalePosition"`
			BreakdownValues []struct {
				Value float64 `json:"Value"`
				Type  string  `json:"Type"`
			} `json:"BreakdownValues"`
			NotClassified float64 `json:"NotClassified,omitempty"`
		} `json:"AssetAllocationBreakdown"`
		SpecialBreakdown []struct {
			NotClassified   float64 `json:"NotClassified"`
			SalePosition    string  `json:"SalePosition"`
			BreakdownValues []struct {
				Value float64 `json:"Value,omitempty"`
				Type  string  `json:"Type"`
			} `json:"BreakdownValues"`
		} `json:"SpecialBreakdown"`
		RegionalExposure []struct {
			NotClassified   float64 `json:"NotClassified"`
			SalePosition    string  `json:"SalePosition"`
			BreakdownValues []struct {
				Type  string  `json:"Type"`
				Value float64 `json:"Value,omitempty"`
			} `json:"BreakdownValues"`
		} `json:"RegionalExposure"`
		MarketCapitalBreakdown []struct {
			NotClassified   float64 `json:"NotClassified"`
			SalePosition    string  `json:"SalePosition"`
			BreakdownValues []struct {
				Value float64 `json:"Value,omitempty"`
				Type  string  `json:"Type"`
			} `json:"BreakdownValues"`
		} `json:"MarketCapitalBreakdown"`
		StyleBoxBreakdown []struct {
			NotClassified   float64 `json:"NotClassified"`
			SalePosition    string  `json:"SalePosition"`
			BreakdownValues []struct {
				Value float64 `json:"Value,omitempty"`
				Type  string  `json:"Type"`
			} `json:"BreakdownValues"`
		} `json:"StyleBoxBreakdown"`
		GlobalStockSectorBreakdown []struct {
			NotClassified   float64 `json:"NotClassified"`
			SalePosition    string  `json:"SalePosition"`
			BreakdownValues []struct {
				Value float64 `json:"Value"`
				Type  string  `json:"Type"`
			} `json:"BreakdownValues"`
		} `json:"GlobalStockSectorBreakdown"`
		IndustryGroupBreakdown []struct {
			NotClassified   float64 `json:"NotClassified"`
			SalePosition    string  `json:"SalePosition"`
			BreakdownValues []struct {
				Value float64 `json:"Value"`
				Type  string  `json:"Type"`
			} `json:"BreakdownValues"`
		} `json:"IndustryGroupBreakdown"`
		CashBreakdown []struct {
			NotClassified   float64 `json:"NotClassified"`
			SalePosition    string  `json:"SalePosition"`
			BreakdownValues []struct {
				Value float64 `json:"Value"`
				Type  string  `json:"Type"`
			} `json:"BreakdownValues"`
		} `json:"CashBreakdown"`
		GlobalBondSectorBreakdownLevel1 []struct {
			NotClassified   float64 `json:"NotClassified"`
			SalePosition    string  `json:"SalePosition"`
			BreakdownValues []struct {
				Value float64 `json:"Value"`
				Type  string  `json:"Type"`
			} `json:"BreakdownValues"`
		} `json:"GlobalBondSectorBreakdownLevel1"`
		GlobalBondSectorBreakdownLevel2 []struct {
			NotClassified   float64 `json:"NotClassified"`
			SalePosition    string  `json:"SalePosition"`
			BreakdownValues []struct {
				Value float64 `json:"Value"`
				Type  string  `json:"Type"`
			} `json:"BreakdownValues"`
		} `json:"GlobalBondSectorBreakdownLevel2"`
		GlobalBondSectorBreakdownLevel3 []struct {
			NotClassified   float64 `json:"NotClassified"`
			SalePosition    string  `json:"SalePosition"`
			BreakdownValues []struct {
				Value float64 `json:"Value"`
				Type  string  `json:"Type"`
			} `json:"BreakdownValues"`
		} `json:"GlobalBondSectorBreakdownLevel3"`
		MaturityRange []struct {
			NotClassified   float64 `json:"NotClassified"`
			SalePosition    string  `json:"SalePosition"`
			BreakdownValues []struct {
				Value float64 `json:"Value,omitempty"`
				Type  string  `json:"Type"`
			} `json:"BreakdownValues"`
		} `json:"MaturityRange"`
		CountryExposure []struct {
			NotClassified   float64 `json:"NotClassified"`
			SalePosition    string  `json:"SalePosition"`
			BreakdownValues []struct {
				Value float64 `json:"Value"`
				Type  string  `json:"Type"`
			} `json:"BreakdownValues"`
		} `json:"CountryExposure"`
		EquityStatistics struct {
			StyleBox                      int     `json:"StyleBox"`
			StyleScore                    float64 `json:"StyleScore"`
			SizeScore                     float64 `json:"SizeScore"`
			ProspectiveEarningsYield      float64 `json:"ProspectiveEarningsYield"`
			ProspectiveBookValueYield     float64 `json:"ProspectiveBookValueYield"`
			ProspectiveRevenueYield       float64 `json:"ProspectiveRevenueYield"`
			ProspectiveCashFlowYield      float64 `json:"ProspectiveCashFlowYield"`
			ProspectiveDividendYield      float64 `json:"ProspectiveDividendYield"`
			Forecasted5YearEarningsGrowth float64 `json:"Forecasted5YearEarningsGrowth"`
			ForecastedEarningsGrowth      float64 `json:"ForecastedEarningsGrowth"`
			ForecastedBookValueGrowth     float64 `json:"ForecastedBookValueGrowth"`
			ForecastedRevenueGrowth       float64 `json:"ForecastedRevenueGrowth"`
			ForecastedCashFlowGrowth      float64 `json:"ForecastedCashFlowGrowth"`
			MarketCapital                 float64 `json:"MarketCapital"`
			DTC                           float64 `json:"DTC"`
			ROA                           float64 `json:"ROA"`
			ROE                           float64 `json:"ROE"`
			NetMargin                     float64 `json:"NetMargin"`
			EarningsYield                 float64 `json:"EarningsYield"`
			BookValueYield                float64 `json:"BookValueYield"`
			RevenueYield                  float64 `json:"RevenueYield"`
			CashFlowYield                 float64 `json:"CashFlowYield"`
			Past3YearEarningsGrowth       float64 `json:"Past3YearEarningsGrowth"`
			EconomicMoatScore             float64 `json:"EconomicMoatScore"`
			FreeCashFlowYield             float64 `json:"FreeCashFlowYield"`
			ROIC                          float64 `json:"ROIC"`
			FinancialHealthGrade          float64 `json:"FinancialHealthGrade"`
			ProfitabilityGrade            float64 `json:"ProfitabilityGrade"`
			GrowthGrade                   float64 `json:"GrowthGrade"`
		} `json:"EquityStatistics"`
		PortfolioHoldings []struct {
			ID                  string  `json:"Id"`
			ExternalID          string  `json:"ExternalId"`
			DetailHoldingTypeID string  `json:"DetailHoldingTypeId"`
			ExternalName        string  `json:"ExternalName"`
			PerformanceID       string  `json:"PerformanceId,omitempty"`
			ISIN                string  `json:"ISIN"`
			CurrencyID          string  `json:"CurrencyId"`
			CUSIP               string  `json:"CUSIP"`
			CountryID           string  `json:"CountryId"`
			SecurityName        string  `json:"SecurityName"`
			Weighting           float64 `json:"Weighting"`
			IndustryID          int     `json:"IndustryId"`
			MarketValue         float64 `json:"MarketValue"`
			GlobalSectorID      string  `json:"GlobalSectorId"`
			NumberOfShare       int     `json:"NumberOfShare"`
			ShareChange         int     `json:"ShareChange,omitempty"`
		} `json:"PortfolioHoldings"`
		HoldingAggregates []struct {
			NumberOfHolding int    `json:"NumberOfHolding,omitempty"`
			SalePosition    string `json:"SalePosition"`
			HoldingsType    string `json:"HoldingsType,omitempty"`
		} `json:"HoldingAggregates"`
	} `json:"Portfolios"`
	DistributionHistory []struct {
		ExDate            string  `json:"ExDate"`
		PayDate           string  `json:"PayDate"`
		ReinvestmentPrice float64 `json:"ReinvestmentPrice"`
		Amount            float64 `json:"Amount"`
	} `json:"DistributionHistory"`
}
