package games

import "time"

type NintendoResponce struct {
	// ResponseHeader struct {
	// 	Status int `json:"status"`
	// 	QTime  int `json:"QTime"`
	// 	Params struct {
	// 		Q     string `json:"q"`
	// 		Bf    string `json:"bf"`
	// 		Start string `json:"start"`
	// 		Fq    string `json:"fq"`
	// 		Sort  string `json:"sort"`
	// 		Rows  string `json:"rows"`
	// 		Wt    string `json:"wt"`
	// 		Bq    string `json:"bq"`
	// 	} `json:"params"`
	// } `json:"responseHeader"`
	Response struct {
		NumFound int     `json:"numFound"`
		Start    int     `json:"start"`
		Docs     []*Game `json:"docs"`
	} `json:"response"`
}

type Game struct {
	FsID                       string      `json:"fs_id" storm:"id"`
	ChangeDate                 time.Time   `json:"change_date" storm:"index"`
	URL                        string      `json:"url"`
	DatesReleasedDts           []time.Time `json:"dates_released_dts" storm:"index"`
	PlayModeTvModeB            bool        `json:"play_mode_tv_mode_b"`
	DeprioritiseB              bool        `json:"deprioritise_b"`
	ImageURL                   string      `json:"image_url"`
	PaidSubscriptionRequiredB  bool        `json:"paid_subscription_required_b"`
	CloudSavesB                bool        `json:"cloud_saves_b"`
	Priority                   time.Time   `json:"priority" storm:"index"`
	DigitalVersionB            bool        `json:"digital_version_b"`
	PlayModeTabletopModeB      bool        `json:"play_mode_tabletop_mode_b"`
	Publisher                  string      `json:"publisher" storm:"index"`
	Excerpt                    string      `json:"excerpt" storm:"index"`
	DateFrom                   time.Time   `json:"date_from" storm:"index"`
	LanguageAvailability       []string    `json:"language_availability"`
	ProductCatalogDescriptionS string      `json:"product_catalog_description_s"`
	PriceDiscountPercentageF   float32     `json:"price_discount_percentage_f"`
	Title                      string      `json:"title"`
	PlayersTo                  int         `json:"players_to"`
	PlayableOnTxt              []string    `json:"playable_on_txt"`
	HitsI                      int         `json:"hits_i"`
	SwitchGameVoucherB         bool        `json:"switch_game_voucher_b"`
	GameCategory               []string    `json:"game_category"`
	PriceRegularF              float32     `json:"price_regular_f" storm:"index"`
	EshopRemovedB              bool        `json:"eshop_removed_b"`
	AgeRatingType              string      `json:"age_rating_type"`
	PriceSortingF              float32     `json:"price_sorting_f"`
	PriceLowestF               float32     `json:"price_lowest_f" storm:"index"`
	AgeRatingValue             string      `json:"age_rating_value"`
	Version                    int64       `json:"_version_"`
	Popularity                 int         `json:"popularity" storm:"index"`
	GameCategoriesTxt          []string    `json:"game_categories_txt"`
	// PrettyGameCategoriesTxt    []string    `json:"pretty_game_categories_txt"`
	// Type                      string      `json:"type"`
	// PgS                       string      `json:"pg_s"`
	// OriginallyForT            string      `json:"originally_for_t"`
	// ImageURLSqS               string      `json:"image_url_sq_s"`
	// TitleExtrasTxt            []string    `json:"title_extras_txt"`
	// ClubNintendo              bool        `json:"club_nintendo"`
	// PrettyDateS               string      `json:"pretty_date_s"`
	// NsuidTxt                  []string    `json:"nsuid_txt"`
	// ImageURLH2X1S             string      `json:"image_url_h2x1_s"`
	// SystemType                []string    `json:"system_type"`
	// AgeRatingSortingI         int         `json:"age_rating_sorting_i"`
	// SystemNamesTxt                   []string `json:"system_names_txt"`
	// PrettyAgeratingS                 string      `json:"pretty_agerating_s"`
	// PriceHasDiscountB                bool        `json:"price_has_discount_b"`
	// PhysicalVersionB                 bool        `json:"physical_version_b"`
	// SortingTitle               string  `json:"sorting_title"`
	// WishlistEmailSquareImageURLS     string      `json:"wishlist_email_square_image_url_s"`
	// WishlistEmailBanner460WImageURLS string `json:"wishlist_email_banner460w_image_url_s"`
}

type (
	GameFollow struct {
		GameId string
		Users  map[int]*UserFollow
	}
	UserFollow struct {
		Showed bool
	}
	GamesList struct {
		Games []*Game
	}
	GamesPagineted interface{}
	Paginated      struct {
		Pages [][]*Game
	}
	GamesDataInterface interface {
		GetAll() []*Game
		GetByID(int) *Game
	}
	Difference struct {
		Field    string
		OldValue interface{}
		NewValue interface{}
	}
)
