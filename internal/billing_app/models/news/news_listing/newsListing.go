package news_listing

import (
	"backend/internal/database"
	"database/sql"
	"fmt"
	"github.com/Conight/go-googletrans"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

type NewsListingEntity struct {
	db *database.StorageDb
}

func NewNewsListEntity(db *database.StorageDb) *NewsListingEntity {
	return &NewsListingEntity{
		db: db,
	}
}

type NewsListing struct {
	//ID uint `gorm:"primarykey"`
	Id        uint   `json:"-" gorm:"id, primaryKey"`
	Slug      string `json:"slug"`
	FullTitle string `json:"full_title"`
	ShortDesc string `json:"short_desc"`
	FullDesc  string `json:"full_desc"`

	FullTitleEn string `json:"full_title_en"`
	FullDescEn  string `json:"full_desc_en"`
	ShortDescEn string `json:"short_desc_en"`

	CreatedAt time.Time `json:"created_at"`
}

func TranslateLocLang(textToTranslate string) string {

	t := translator.New()
	translatedText, err := t.Translate(textToTranslate, "ru", "en")
	if err != nil {
		fmt.Println("Failed to translate:", err)
		return ""
	} else {
		fmt.Printf("Original Text: %s\n", textToTranslate)
		fmt.Printf("Translated Text (English): %s\n", translatedText)
		//
		return translatedText.Text
	}
}

// TranslateToEnglish Translate Text Content
func (n *NewsListing) TranslateToEnglish() *NewsListing {
	n.FullTitleEn = TranslateLocLang(n.FullTitle)
	n.ShortDescEn = TranslateLocLang(n.ShortDesc)
	n.FullDescEn = TranslateLocLang(n.FullDesc)
	return n
	//n.FullTitleEn = fullTitleEn
	//n.ShortDescEn = shortDescEn
	//n.FullDescEn = fullDescEn
}

func (nle NewsListingEntity) GetAllListingNews() []NewsListing {
	db, err := nle.db.GetDB()
	if err != nil {
		return []NewsListing{}
	}
	var newsListingAll []NewsListing
	// Fetch services
	if db.Select("*").
		Order("created_at desc").
		Find(&newsListingAll).Error != nil {
		return newsListingAll
	} else {
		return newsListingAll
	}
}

func (nle NewsListingEntity) GetBySlugListingNews(slugName string) (NewsListing, error) {
	var news NewsListing
	db, err := nle.db.GetDB()
	if err != nil {
		return NewsListing{}, err
	}

	db.
		Select("*").
		Where("slug = @slug", sql.Named("slug", slugName)).
		First(&news)

	if db.Error != nil {
		return news, db.Error
	} else {
		return news, nil
	}
}

// CreateListingNews TODO: Логику перенести в Handlers (Service)
func (nle NewsListingEntity) CreateListingNews(c *fiber.Ctx) error {
	db, err := nle.db.GetDB()
	if err != nil {
		return err
	}

	news := NewsListing{
		Slug:      uuid.New().String(),
		FullTitle: time.Now().String() + "Обновление тарифов",
		ShortDesc: time.Now().String() + "Мы характерно обновили наши тарифы на услуги",
		FullDesc:  time.Now().String() + "С нашими новыми тарифами вы получите доступ к еще большему функционалу и удобству. Наши услуги стали более доступными и гибкими, чтобы удовлетворить ваши потребности. Переходите на новые тарифы и погрузитесь в мир бесконечных возможностей, предлагаемых нами для вас!",
	}

	// Вызываем метод для перевода на английский
	newsAndTranslate := news.TranslateToEnglish()

	db.Select("*").Create(&newsAndTranslate)

	if db.Error != nil {
		return c.JSON(fiber.Map{
			"success": false,
		})
	} else {
		// Теперь у вас есть оригинальный текст и его английский перевод
		fmt.Println("Original Short Desc:", news.ShortDesc)
		fmt.Println("Translated Short Desc (English):", newsAndTranslate.ShortDescEn)
		fmt.Println("Original Full Desc:", news.FullDesc)
		fmt.Println("Translated Full Desc (English):", newsAndTranslate.FullDescEn)
		//
		return c.JSON(fiber.Map{
			"success": true,
		})
	}
}
