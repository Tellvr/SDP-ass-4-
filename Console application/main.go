package main

import (
  "fmt"
  "math/rand"
  "time"
)

// Sport interfaces
type Sport interface {
	GetRecommendation() string
  }
  
  type Football struct{}
  
  func (f Football) GetRecommendation() string {
	return "Cleats, shin guards, ball" 
  }
// Observer patterns
type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)  
	Notify()
  }
  
  type Observer interface {
	Update()
  }
  
  type ShopSubject struct {
	observers []Observer
  }
  
  // Subject methods 
  
  type ShopObserver struct{}
  
  func (o ShopObserver) Update() {
	fmt.Println("Notification sent")
  }
// Decorator patterns
type Color interface {
  GetColor() string
}

// Color types

// Color factory
type ColorFactory struct{}

func (f ColorFactory) GetColor(name string) Color {
  switch name {
  case "red":
    return &Red{}
  case "blue": 
    return &Blue{}  
  default:
    return nil
  }
}


// Recommendation decorator
type RecommendationDecorator struct {
	Sport 
	Size string
	Color
  }
  
  func (r RecommendationDecorator) GetRecommendation() string {
	return fmt.Sprint(
	  r.sport.GetRecommendation(),
	  " size ", r.size,
	  " color ", r.color.GetColor()) 
  }
  // Cart
type Cart struct {
	items []string
  } 
  
  func (c *Cart) AddItem(item string) {
	c.items = append(c.items, item)
  }
  
  func (c *Cart) RemoveItem(item string) {
	// removal logic
  }
  
  func (c *Cart) DeleteAll() {
	c.items = []string{}
  }
  
  func (c *Cart) PrintItems() {
	// print items
  }
  // Item model
type Item struct {
	Name string
	Stock int
  }
  // Sport factory
type SportFactory struct{}

func (f SportFactory) GetSport(name string) Sport {
  switch name {
  case "football":
    return &Football{}
  case "basketball":
    return &Basketball{}
  // etc
  default:
    return nil
  }
}

// Get initial sport
factory := SportFactory{}
sport := factory.GetSport("football")
// Color factory
type ColorFactory struct{}

func (f ColorFactory) GetColor(name string) Color {
  switch name {
  case "red":
    return &Red{}
  case "blue": 
    return &Blue{}  
  default:
    return nil
  }
}

// Get color 
colorFactory := ColorFactory{}
color := colorFactory.GetColor("red")
// Recommendation service
type RecommendationService struct {
	SportFactory 
	ColorFactory
  }
  
  func (s *RecommendationService) GetRecommendation(
	sportName string, 
	size string,
	colorName string) Recommendation {
  
	sport := s.SportFactory.GetSport(sportName)
	color := s.ColorFactory.GetColor(colorName)
  
	return RecommendationDecorator{
	  Sport: sport,
	  Size: size, 
	  Color: color,
	}
  }
  
  // Usage
  service := RecommendationService{}
  rec := service.GetRecommendation("football", "large", "red")
  

func main() {

  rand.Seed(time.Now().UnixNano())

  var items [3]Item

  // Initialize random stock

  rand.Seed(time.Now().UnixNano())

  for i := range items {
    items[i].Stock = rand.Intn(10)
  }

  // Print available stock

  for i, item := range items {
    fmt.Printf("%s: %d\n", item.Name, item.Stock) 
  }

 // Check stock before adding item 
if items[itemIndex].Stock == 0 {
	fmt.Println("Out of stock")
  }
}

var choice string

for {
	    // Sport selection
fmt.Println("Select sport:")
var sport Sport

// Generate recommendation
recommendation := RecommendationDecorator{sport: sport}

  switch choice {

  case "view cart":
	// Print cart items
	cart.PrintItems()

  case "add item":
	// Prompt for item  
	var item string  
	fmt.Scan(&item)

	// Add to cart
	cart.AddItem(item)

  case "change sport": 
	// Prompt for new sport
	var sport Sport 

	// Regenerate recommendation

  case "delete item":
	// Prompt for item to delete
	var item string

	// Delete from cart  
	cart.DeleteItem(item)

  case "checkout":
	// Process payment  
	// Empty cart
	cart.DeleteAll()

  case "quit":
	return
  }

}
