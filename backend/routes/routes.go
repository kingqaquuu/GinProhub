package routes
import(
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoadRoutes(router *gin.Engine,db *gorm.DB){
	RegisterAuthRoutes(router,db)
	
	
}