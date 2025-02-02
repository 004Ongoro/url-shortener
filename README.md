# **Go URL Shortener**  

A simple **URL shortener** built using **Go**. It takes long URLs and generates shorter, unique links that redirect to the original URL.  

## **Features**  
✅ Shorten long URLs  
✅ Redirect to original URLs  
✅ Store shortened URLs in a file (persistence)  
✅ Simple web-based UI  
✅ Deployed on **Render**  

## **Getting Started**  

### **1. Clone the Repository**  
```sh
git clone https://github.com/004Ongoro/url-shortener.git
cd url-shortener
```

### **2. Install Dependencies**  
Ensure you have Go installed. Then run:  
```sh
go mod tidy
```

### **3. Run the Server**  
```sh
go run main.go
```

### **4. Shorten a URL**  
Send a **POST** request to:  
```
http://localhost:8080/api/shorten
```
With a JSON body:  
```json
{
  "url": "https://example.com"
}
```
Response:  
```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

### **5. Access the Shortened URL**  
Open in a browser:  
```
http://localhost:8080/abc123
```
It will **redirect** to the original URL.

## **Deployment**  

This project is deployed on **Render**. To deploy manually:  
1. Push to GitHub.  
2. Connect to Render.  
3. Set the build command:  
   ```sh
   go build -o main .
   ```
4. Set the start command:  
   ```sh
   ./main
   ```
5. Deploy!

## **Upcoming Features**  
- ✅ Analytics (track link clicks)  
- ✅ Database storage (instead of a file)  
- ✅ Custom short URLs  

## **License**  
MIT License © 2025 [George Ongoro](https://github.com/004Ongoro)