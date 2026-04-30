package middleware

import (
	"context"
	"fmt"
	"hta-platform/global"
	"net/http"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

var jwks *keyfunc.JWKS

// InitAuth0 initializes the JWKS keyfunc for Auth0 token validation.
// Must be called once at startup after config is loaded.
func InitAuth0() error {
	jwksURL := fmt.Sprintf("https://%s/.well-known/jwks.json", global.ConfigValue.Auth0Domain)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	jwks, err = keyfunc.Get(jwksURL, keyfunc.Options{
		Ctx:              ctx,
		RefreshInterval:  time.Hour,
		RefreshRateLimit: 5 * time.Minute,
	})
	if err != nil {
		return fmt.Errorf("failed to initialize JWKS: %w", err)
	}

	global.Logger.Info("Auth0 JWKS initialized", zap.String("domain", global.ConfigValue.Auth0Domain))
	return nil
}

// Auth0Guard is a Gin middleware that validates Auth0 JWT tokens.
func Auth0Guard() gin.HandlerFunc {
	issuer := fmt.Sprintf("https://%s/", global.ConfigValue.Auth0Domain)
	audience := global.ConfigValue.Auth0Audience

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or malformed token"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, jwks.Keyfunc,
			jwt.WithIssuer(issuer),
			jwt.WithExpirationRequired(),
		)
		if err != nil || !token.Valid {
			global.Logger.Debug("Token validation failed", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
			return
		}

		// Validate audience
		if !containsAudience(claims["aud"], audience) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid audience"})
			return
		}

		// Set user claims in context for downstream handlers
		c.Set("user", claims)
		c.Set("user_id", claims["sub"])

		c.Next()
	}
}

func containsAudience(aud interface{}, expected string) bool {
	switch v := aud.(type) {
	case string:
		return v == expected
	case []interface{}:
		for _, a := range v {
			if s, ok := a.(string); ok && s == expected {
				return true
			}
		}
	}
	return false
}
