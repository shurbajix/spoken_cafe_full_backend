func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header missing"})
			return
		}
		parts := strings.Split(authHeader, " ")
		if lem(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid Authorization header"})
			return
		}

		tokenStr := parts[1]
		claims, err := ValidateJWT(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		//Attach user info to context for handlers
		c.set("userID", calaims.UserID)
		c.Next()
	}
}