package middleware

import (
    "gin-prohub/utils/jwt"
    "gin-prohub/utils/response"
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 获取 Authorization 头部信息
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            response.Error(c, http.StatusUnauthorized, "未提供认证信息")
            c.Abort()
            return
        }

        // 解析 Bearer Token
        parts := strings.SplitN(authHeader, " ", 2)
        if !(len(parts) == 2 && parts[0] == "Bearer") {
            response.Error(c, http.StatusUnauthorized, "认证信息格式错误")
            c.Abort()
            return
        }

        // 验证 Token
        claims, err := jwt.ParseToken(parts[1])
        if err != nil {
            response.Error(c, http.StatusUnauthorized, "无效的Token")
            c.Abort()
            return
        }

        // 将用户信息存储在上下文中
        c.Set("userID", claims.UserId)
        c.Set("username", claims.Username)
        c.Next()
    }
}