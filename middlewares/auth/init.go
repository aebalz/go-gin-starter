package auth

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	gooauth2gorm "github.com/nouhoum/go-oauth2-gorm/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, r *gin.Engine) {
	// Auto migrate the models for the OAuth2 package
	// err := db.AutoMigrate(&gooauth2gorm.Token{}, &gooauth2gorm.Client{})
	// if err != nil {
	// 	log.Fatalf("failed to migrate database: %v", err)
	// }

	// Create the OAuth2 manager
	manager := manage.NewDefaultManager()

	authConfig := &gooauth2gorm.Config{
		DBType: gooauth2gorm.PostgresSQL,
	}

	// Use GORM token store
	tokenStore := gooauth2gorm.NewTokenStoreWithDB(authConfig, db, "", 0)
	manager.MapTokenStorage(tokenStore)

	// Use GORM client store
	clientStore := gooauth2gorm.NewClientStoreWithDB(authConfig, db, "")
	// manager.MapClientStorage(clientStore)
	token, err := clientStore.GetByID(context.Background(), "YGvMeedZPKeByrnZfypTaV1UKOz9YpQANvh4DEXL")
	fmt.Println("token >>>> ", token)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Initialize the oauth2 service
	// ginserver.InitServer(manager)
	// ginserver.SetAllowGetAccessRequest(true)
	// ginserver.SetClientInfoHandler(server.ClientFormHandler)

	// Create an OAuth2 server instance
	srv := server.NewServer(server.NewConfig(), manager)

	auth := r.Group("/oauth2")
	{
		auth.POST("/token", func(c *gin.Context) {
			err := srv.HandleTokenRequest(c.Writer, c.Request)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
			}
		})

		auth.GET("/authorize", func(c *gin.Context) {
			err := srv.HandleAuthorizeRequest(c.Writer, c.Request)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
			}
		})
	}

}
