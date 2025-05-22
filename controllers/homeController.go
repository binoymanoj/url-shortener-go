package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeController(c *gin.Context) {
	html := `<!DOCTYPE html>
	<html lang="en">
	<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>URL Shortener</title>
  <script src="https://unpkg.com/htmx.org@1.9.10"></script>
  <style>
  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }

  body {
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 20px;
  }

  .container {
    background: white;
    padding: 40px;
    border-radius: 20px;
    box-shadow: 0 20px 40px rgba(0,0,0,0.1);
    width: 100%;
    max-width: 500px;
  }

  h1 {
    text-align: center;
    color: #333;
    font-size: 2.5rem;
    font-weight: 700;
  }

  .form-group {
    margin-bottom: 20px;
  }

  label {
    display: block;
    margin-bottom: 8px;
    color: #555;
    font-weight: 500;
  }

  .username {
    margin-bottom: 30px;
    text-align: right;
    font-size: 14px;
  }

  .username a {
    color: #333;
  }

  input[type="url"] {
    width: 100%;
    padding: 15px;
    border: 2px solid #e1e5e9;
    border-radius: 10px;
    font-size: 16px;
    transition: border-color 0.3s ease;
  }

  input[type="url"]:focus {
    outline: none;
    border-color: #667eea;
  }

  .btn {
    width: 100%;
    padding: 15px;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
    border: none;
    border-radius: 10px;
    font-size: 16px;
    font-weight: 600;
    cursor: pointer;
    transition: transform 0.2s ease;
  }

  .btn:hover {
    transform: translateY(-2px);
  }

  .btn:disabled {
    opacity: 0.6;
    cursor: not-allowed;
    transform: none;
  }

  .result {
    margin-top: 30px;
    padding: 20px;
    background: #f8f9fa;
    border-radius: 10px;
    border-left: 4px solid #28a745;
  }

  .result h3 {
    color: #28a745;
    margin-bottom: 10px;
  }

  .url-display {
    display: flex;
    align-items: center;
    gap: 10px;
    margin-top: 10px;
  }

  .short-url {
    flex: 1;
    padding: 10px;
    background: white;
    border: 1px solid #dee2e6;
    border-radius: 5px;
    font-family: monospace;
    font-size: 14px;
  }

  .copy-btn {
    padding: 10px 15px;
    background: #28a745;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    font-size: 14px;
  }

  .copy-btn:hover {
    background: #218838;
  }

  .error {
    margin-top: 20px;
    padding: 15px;
    background: #f8d7da;
    color: #721c24;
    border-radius: 10px;
    border-left: 4px solid #dc3545;
  }

  .stats {
    margin-top: 10px;
    font-size: 14px;
    color: #6c757d;
  }

  .htmx-indicator {
    display: none;
  }

  .htmx-request .htmx-indicator {
    display: inline;
  }

  .htmx-request .btn {
    opacity: 0.6;
  }
  </style>
	</head>
	<body>
  <div class="container">
  <h1>🔗 URL Shortener</h1>
  <p class="username">by <a href="https://github.com/binoymanoj" target="_blank">binoy_manoj</a></p>

  <form hx-post="/shrtnurl" hx-target="#result" hx-indicator="#loading">
  <div class="form-group">
  <label for="url">Enter URL to shorten:</label>
  <input 
  type="url" 
  id="url" 
  name="url" 
  placeholder="https://example.com/very/long/url"
  required
  >
  </div>
  <button type="submit" class="btn">
  <span class="htmx-indicator" id="loading">⏳ Shortening...</span>
  <span class="htmx-indicator:not(.htmx-request)">✨</span>
  Shorten URL
  </button>
  </form>

  <div id="result"></div>
  </div>

  <script>
  function copyToClipboard(text) {
    navigator.clipboard.writeText(text).then(function() {
      // Show success feedback
      const btn = event.target;
      const originalText = btn.textContent;
      btn.textContent = '✅ Copied!';
      btn.style.background = '#28a745';

      setTimeout(() => {
        btn.textContent = originalText;
        btn.style.background = '#28a745';
      }, 2000);
    });
  }
  </script>
	</body>
	</html>`

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}
