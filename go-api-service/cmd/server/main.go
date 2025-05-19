package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Post represents the structure of a post from the JSONPlaceholder API
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	// Initialize HTTP server
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/posts", handlePosts)
	http.HandleFunc("/posts/", handleSinglePost)

	// Set up server with timeouts (good SRE practice)
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go API Service!\n")
	fmt.Fprintf(w, "Available endpoints:\n")
	fmt.Fprintf(w, "- /posts - Get all posts\n")
	fmt.Fprintf(w, "- /posts/{id} - Get a specific post\n")
}

func handlePosts(w http.ResponseWriter, r *http.Request) {
	posts, err := fetchPosts()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching posts: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func handleSinglePost(w http.ResponseWriter, r *http.Request) {
	// Simple path parsing to get ID
	path := r.URL.Path
	if path == "/posts/" {
		http.Error(w, "Please provide a post ID", http.StatusBadRequest)
		return
	}

	id := path[len("/posts/"):]
	post, err := fetchSinglePost(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching post: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func fetchPosts() ([]Post, error) {
	// Create an HTTP client with timeouts (good SRE practice)
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var posts []Post
	if err := json.Unmarshal(body, &posts); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return posts, nil
}

func fetchSinglePost(id string) (Post, error) {
	var post Post

	// Create an HTTP client with timeouts (good SRE practice)
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%s", id))
	if err != nil {
		return post, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return post, fmt.Errorf("API returned status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return post, fmt.Errorf("error reading response body: %w", err)
	}

	if err := json.Unmarshal(body, &post); err != nil {
		return post, fmt.Errorf("error parsing JSON: %w", err)
	}

	return post, nil
}
