package main

import (
	"fmt" // Importa o pacote fmt para formatação de strings
	"net/http" // Importa o pacote net/http para funcionalidades HTTP (Pacote padrão do Go para HTTP requests/responses.)
	"os" // Importa o pacote os para interagir com variáveis de ambiente

	"github.com/gin-contrib/cors" // Importa o middleware CORS para o Gin
	"github.com/gin-gonic/gin" // Importa o framework web Gin

	// Importa os pacotes de rotas 

	// "workplace/api/routes"
)

func main() {
	// Define o modo do Gin (release ou debug)
	switch os.Getenv("APP_ENV") {
		case "production":
			gin.SetMode(gin.ReleaseMode)
		default:
			gin.SetMode(gin.DebugMode)
	}

	// Cria a instância do servidor
	app := gin.Default()

	// Configura o CORS (similar ao app.use(cors()))
	app.Use(cors.Default())

	// Rotas públicas
	app.GET("/", func(c *gin.Context) {
		protocol := "http"
		if c.Request.TLS != nil {
			protocol = "https"
		}
		host := c.Request.Host
		rootDomain := fmt.Sprintf("%s://%s", protocol, host)

		c.JSON(http.StatusOK, gin.H{
			"status_server": "ok",
			"dominio_raiz":  rootDomain,
			"env":    os.Getenv("APP_ENV"),
			"atualizacao":   "14/09/2024 - 18:42",
			"rotas": gin.H{
				"GET - teste": fmt.Sprintf("%s/api/teste", rootDomain),
			},
		})
	})

	// Rotas das outras entidades
	// routes.RegisterUsuarioRoutes(app)
	// routes.RegisterInstituicoesRoutes(app)
	// routes.RegisterInstituicaoUsuarioRoutes(app)

	// Porta do servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Inicia o servidor
	fmt.Printf("Sistema inicializado: Acesso em http://localhost:%s\n", port)
	app.Run(":" + port)
}
