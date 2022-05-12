package games

import "time"

type Game struct {
	FsID                             string      `json:"fs_id"`
	ChangeDate                       time.Time   `json:"change_date"`
	URL                              string      `json:"url"`
	Type                             string      `json:"type"`
	DatesReleasedDts                 []time.Time `json:"dates_released_dts"`
	ClubNintendo                     bool        `json:"club_nintendo"`
	PrettyDateS                      string      `json:"pretty_date_s"`
	PlayModeTvModeB                  bool        `json:"play_mode_tv_mode_b"`
	ImageURLSqS                      string      `json:"image_url_sq_s"`
	DeprioritiseB                    bool        `json:"deprioritise_b"`
	PgS                              string      `json:"pg_s"`
	ImageURL                         string      `json:"image_url"`
	OriginallyForT                   string      `json:"originally_for_t"`
	PaidSubscriptionRequiredB        bool        `json:"paid_subscription_required_b"`
	CloudSavesB                      bool        `json:"cloud_saves_b"`
	Priority                         time.Time   `json:"priority"`
	DigitalVersionB                  bool        `json:"digital_version_b"`
	TitleExtrasTxt                   []string    `json:"title_extras_txt"`
	ImageURLH2X1S                    string      `json:"image_url_h2x1_s"`
	SystemType                       []string    `json:"system_type"`
	AgeRatingSortingI                int         `json:"age_rating_sorting_i"`
	GameCategoriesTxt                []string    `json:"game_categories_txt"`
	PlayModeTabletopModeB            bool        `json:"play_mode_tabletop_mode_b"`
	Publisher                        string      `json:"publisher"`
	Excerpt                          string      `json:"excerpt"`
	NsuidTxt                         []string    `json:"nsuid_txt"`
	DateFrom                         time.Time   `json:"date_from"`
	LanguageAvailability             []string    `json:"language_availability"`
	PriceHasDiscountB                bool        `json:"price_has_discount_b"`
	ProductCatalogDescriptionS       string      `json:"product_catalog_description_s"`
	PriceDiscountPercentageF         float32         `json:"price_discount_percentage_f"`
	Title                            string      `json:"title"`
	SortingTitle                     string      `json:"sorting_title"`
	WishlistEmailSquareImageURLS     string      `json:"wishlist_email_square_image_url_s"`
	PlayersTo                        int         `json:"players_to"`
	WishlistEmailBanner640WImageURLS string      `json:"wishlist_email_banner640w_image_url_s"`
	PlayableOnTxt                    []string    `json:"playable_on_txt"`
	HitsI                            int         `json:"hits_i"`
	PrettyGameCategoriesTxt          []string    `json:"pretty_game_categories_txt"`
	SwitchGameVoucherB               bool        `json:"switch_game_voucher_b"`
	GameCategory                     []string    `json:"game_category"`
	SystemNamesTxt                   []string    `json:"system_names_txt"`
	PrettyAgeratingS                 string      `json:"pretty_agerating_s"`
	PriceRegularF                    int         `json:"price_regular_f"`
	EshopRemovedB                    bool        `json:"eshop_removed_b"`
	AgeRatingType                    string      `json:"age_rating_type"`
	PriceSortingF                    int         `json:"price_sorting_f"`
	PriceLowestF                     int         `json:"price_lowest_f"`
	AgeRatingValue                   string      `json:"age_rating_value"`
	PhysicalVersionB                 bool        `json:"physical_version_b"`
	WishlistEmailBanner460WImageURLS string      `json:"wishlist_email_banner460w_image_url_s"`
	Version                          int64       `json:"_version_"`
	Popularity                       int         `json:"popularity"`
}

type Games []*Game
