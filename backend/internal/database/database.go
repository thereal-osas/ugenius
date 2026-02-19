package database

import (
	"fmt"
	"log"
	"time"

	"github.com/ugenius/backend/internal/config"
	"github.com/ugenius/backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := cfg.DSN()

	logLevel := logger.Silent
	if cfg.SSLMode == "disable" {
		logLevel = logger.Info
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// Connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	DB = db
	log.Println("Database connected successfully")
	return db, nil
}

func Migrate(db *gorm.DB) error {
	log.Println("Running database migrations...")

	err := db.AutoMigrate(
		// Core models
		&models.Campus{},
		&models.User{},
		&models.RefreshToken{},

		// Reading hours
		&models.ReadingHour{},
		&models.AdminReview{},

		// Goals and achievements
		&models.StudyGoal{},
		&models.Achievement{},

		// Notifications
		&models.Notification{},
		&models.NotificationPreferences{},

		// Study groups
		&models.StudyGroup{},
		&models.StudyGroupMember{},
		&models.StudySession{},

		// Scholarships
		&models.Scholarship{},
		&models.ScholarshipApplication{},

		// Waitlist
		&models.WaitlistRegistration{},

		// Events
		&models.Event{},
		&models.EventRegistration{},

		// Gallery
		&models.Gallery{},
	)
	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Database migrations completed successfully")
	return nil
}

func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// SeedCampuses populates the campuses table with Nigerian universities and polytechnics
func SeedCampuses(db *gorm.DB) error {
	campuses := []models.Campus{
		// Federal Universities
		{Name: "University of Lagos", Code: "UNILAG", Location: "Akoka, Yaba", State: "Lagos", Country: "Nigeria", IsActive: true},
		{Name: "University of Ibadan", Code: "UI", Location: "Ibadan", State: "Oyo", Country: "Nigeria", IsActive: true},
		{Name: "University of Nigeria, Nsukka", Code: "UNN", Location: "Nsukka", State: "Enugu", Country: "Nigeria", IsActive: true},
		{Name: "Obafemi Awolowo University", Code: "OAU", Location: "Ile-Ife", State: "Osun", Country: "Nigeria", IsActive: true},
		{Name: "University of Benin", Code: "UNIBEN", Location: "Benin City", State: "Edo", Country: "Nigeria", IsActive: true},
		{Name: "University of Ilorin", Code: "UNILORIN", Location: "Ilorin", State: "Kwara", Country: "Nigeria", IsActive: true},
		{Name: "University of Port Harcourt", Code: "UNIPORT", Location: "Port Harcourt", State: "Rivers", Country: "Nigeria", IsActive: true},
		{Name: "Federal University of Technology Minna", Code: "FUTMINNA", Location: "Minna", State: "Niger", Country: "Nigeria", IsActive: true},
		{Name: "Federal University of Technology Owerri", Code: "FUTO", Location: "Owerri", State: "Imo", Country: "Nigeria", IsActive: true},
		{Name: "University of Abuja", Code: "UNIABUJA", Location: "Abuja", State: "FCT", Country: "Nigeria", IsActive: true},
		{Name: "University of Calabar", Code: "UNICAL", Location: "Calabar", State: "Cross River", Country: "Nigeria", IsActive: true},
		{Name: "University of Uyo", Code: "UNIUYO", Location: "Uyo", State: "Akwa Ibom", Country: "Nigeria", IsActive: true},
		{Name: "Federal University of Agriculture, Abeokuta", Code: "FUNAAB", Location: "Abeokuta", State: "Ogun", Country: "Nigeria", IsActive: true},
		
		// State Universities
		{Name: "Olabisi Onabanjo University", Code: "OOU", Location: "Ago-Iwoye", State: "Ogun", Country: "Nigeria", IsActive: true},
		{Name: "Lagos State University", Code: "LASU", Location: "Ojo", State: "Lagos", Country: "Nigeria", IsActive: true},
		{Name: "Ambrose Alli University", Code: "AAU", Location: "Ekpoma", State: "Edo", Country: "Nigeria", IsActive: true},
		{Name: "Delta State University", Code: "DELSU", Location: "Abraka", State: "Delta", Country: "Nigeria", IsActive: true},
		{Name: "Rivers State University", Code: "RSU", Location: "Port Harcourt", State: "Rivers", Country: "Nigeria", IsActive: true},
		{Name: "Benue State University, Makurdi", Code: "BSU", Location: "Makurdi", State: "Benue", Country: "Nigeria", IsActive: true},
		{Name: "Ignatius Ajuru University of Education", Code: "IAUE", Location: "Port Harcourt", State: "Rivers", Country: "Nigeria", IsActive: true},
		{Name: "Southern Delta University, Ozoro", Code: "SDU", Location: "Ozoro", State: "Delta", Country: "Nigeria", IsActive: true},
		
		// Private Universities
		{Name: "Igbinedion University", Code: "IUO", Location: "Okada", State: "Edo", Country: "Nigeria", IsActive: true},
		{Name: "Novena Private University", Code: "NPU", Location: "Ogume", State: "Delta", Country: "Nigeria", IsActive: true},
		{Name: "Ambrose Ali College of Medicine, Ekpoma", Code: "AACME", Location: "Ekpoma", State: "Edo", Country: "Nigeria", IsActive: true},
		
		// International Universities
		{Name: "University of Texas Arlington", Code: "UTA", Location: "Arlington, Texas", State: "Texas", Country: "USA", IsActive: true},
		{Name: "Tarrant Community College - Southeast", Code: "TCCSE", Location: "Arlington, Texas", State: "Texas", Country: "USA", IsActive: true},
		{Name: "Tarrant Community College Trinity River", Code: "TCCTR", Location: "Fort Worth, Texas", State: "Texas", Country: "USA", IsActive: true},
		
		// Polytechnics and Colleges
		{Name: "Yaba College of Technology", Code: "YABATECH", Location: "Yaba", State: "Lagos", Country: "Nigeria", IsActive: true},
		{Name: "Federal Polytechnic Auchi", Code: "AUCHIPOLY", Location: "Auchi", State: "Edo", Country: "Nigeria", IsActive: true},
		{Name: "Lagos State Polytechnic", Code: "LASPOTECH", Location: "Ikorodu", State: "Lagos", Country: "Nigeria", IsActive: true},
		{Name: "Kwara State Polytechnic", Code: "KWARAPOLY", Location: "Ilorin", State: "Kwara", Country: "Nigeria", IsActive: true},
		{Name: "Lagos City Polytechnic", Code: "LCP", Location: "Ikeja", State: "Lagos", Country: "Nigeria", IsActive: true},
		{Name: "Polytechnic Ibadan", Code: "POLYIBADAN", Location: "Ibadan", State: "Oyo", Country: "Nigeria", IsActive: true},
	}

	for _, campus := range campuses {
		// Check if campus already exists by code
		var existing models.Campus
		result := db.Where("code = ?", campus.Code).First(&existing)
		if result.Error == nil {
			// Campus already exists, skip
			continue
		}
		// Create the campus
		if err := db.Create(&campus).Error; err != nil {
			log.Printf("Failed to seed campus %s: %v", campus.Code, err)
			continue
		}
		log.Printf("Seeded campus: %s (%s)", campus.Name, campus.Code)
	}

	log.Println("Campus seeding completed")
	return nil
}

// SeedEvents populates the events table with sample upcoming events
func SeedEvents(db *gorm.DB) error {
	// Check if events already exist
	var count int64
	db.Model(&models.Event{}).Count(&count)
	if count > 0 {
		log.Println("Events already seeded, skipping...")
		return nil
	}

	// Get a system user ID for created_by (use first admin or create placeholder)
	var adminUser models.User
	if err := db.Where("role = ?", models.RoleSuperAdmin).First(&adminUser).Error; err != nil {
		// No admin exists yet, skip seeding events
		log.Println("No admin user found, skipping event seeding...")
		return nil
	}

	now := time.Now()
	events := []models.Event{
		{
			Title:       "Academic Excellence Workshop",
			Description: "Join us for an intensive workshop on study techniques, time management, and exam preparation strategies. Learn from top-performing students and mentors.",
			Type:        models.EventTypeWorkshop,
			Status:      models.EventStatusUpcoming,
			StartTime:   now.AddDate(0, 0, 7).Add(10 * time.Hour), // 1 week from now at 10 AM
			EndTime:     now.AddDate(0, 0, 7).Add(14 * time.Hour), // 4 hours duration
			Location:    "Online",
			VirtualLink: "https://meet.google.com/ugenius-workshop",
			ImageURL:    "https://images.unsplash.com/photo-1524178232363-1fb2b075b655?w=800&h=400&fit=crop&q=80",
			CreatedByID: adminUser.ID,
			IsFeatured:  true,
		},
		{
			Title:       "First Class Achievers Seminar",
			Description: "Hear from students who achieved first-class honors and learn their secrets to academic success. Q&A session included.",
			Type:        models.EventTypeSeminar,
			Status:      models.EventStatusUpcoming,
			StartTime:   now.AddDate(0, 0, 14).Add(14 * time.Hour), // 2 weeks from now at 2 PM
			EndTime:     now.AddDate(0, 0, 14).Add(17 * time.Hour), // 3 hours duration
			Location:    "Online",
			VirtualLink: "https://zoom.us/ugenius-seminar",
			ImageURL:    "https://images.unsplash.com/photo-1541339907198-e08756dedf3f?w=800&h=400&fit=crop&q=80",
			CreatedByID: adminUser.ID,
			IsFeatured:  true,
		},
		{
			Title:       "Study Group Networking Meetup",
			Description: "Connect with fellow U-Genius members, form study groups, and build lasting academic partnerships.",
			Type:        models.EventTypeMeetup,
			Status:      models.EventStatusUpcoming,
			StartTime:   now.AddDate(0, 0, 21).Add(16 * time.Hour), // 3 weeks from now at 4 PM
			EndTime:     now.AddDate(0, 0, 21).Add(19 * time.Hour), // 3 hours duration
			Location:    "Online",
			VirtualLink: "https://meet.google.com/ugenius-meetup",
			ImageURL:    "https://images.unsplash.com/photo-1456513080510-7bf3a84b82f8?w=800&h=400&fit=crop&q=80",
			CreatedByID: adminUser.ID,
			IsFeatured:  true,
		},
	}

	for _, event := range events {
		if err := db.Create(&event).Error; err != nil {
			log.Printf("Failed to seed event %s: %v", event.Title, err)
			continue
		}
		log.Printf("Seeded event: %s", event.Title)
	}

	log.Println("Event seeding completed")
	return nil
}
